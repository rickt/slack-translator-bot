// http://play.golang.org/p/SQjlJEcHF1

package slacktranslatorbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"net/http"
	"os"
	"strings"
)

// globals
var (
	env envVars
)

// types
type GoogleTranslateAPIResponse struct {
	Data struct {
		Translations []struct {
			TranslatedText string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

// helper function to do a case-insensitive search
func ciContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}

// helper func to read the api response body (an io.ReadCloser) and output it as simple string
func getBody(resp *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String()
}

// there's no main() func in google app engine
func init() {
	// get runtime options from the app.yaml
	env.APIKey = os.Getenv("GOOGLE_TRANSLATE_API_KEY")
	env.BaseURL = os.Getenv("GOOGLE_TRANSLATE_BASEURL")
	env.VerifyToken = os.Getenv("SLACK_VERIFY_TOKEN")
	// setup http handlers
	http.HandleFunc("/", handler_redirect)
	http.HandleFunc("/translate/en_ja", handler_translate)
	http.HandleFunc("/translate/ja_en", handler_translate)
}

// assume default url is en --> ja by redirecting / to /en_ja
func handler_redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/translate/en_ja", 302)
}

func handler_translate(w http.ResponseWriter, r *http.Request) {
	// what is our from/to?
	var emoji_from, emoji_to, from, to, texttobetranslated string
	// this could be handled much better, but whatever this is just demo code
	if ciContains(r.RequestURI, "en_ja") {
		from = "en"
		to = "ja"
		emoji_from = ":uk:"
		emoji_to = ":jp:"
	}
	if ciContains(r.RequestURI, "ja_en") {
		from = "ja"
		to = "en"
		emoji_from = ":jp:"
		emoji_to = ":uk:"
	}
	// create a google app engine context
	ctx := appengine.NewContext(r)
	// what is our OG query from slack?
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "\n\nError parsing form, err=%s, rsp.Body= %s", err.Error(), r.Body)
		return
	}
	defer r.Body.Close()
	// build the url! an example call to translate API:
	// https://www.googleapis.com/language/translate/v2?q=QUERY&target=ja&format=text&source=en&key=KEY
	texttobetranslated = r.URL.Query().Get("text")
	apiurl := fmt.Sprintf(env.BaseURL + "?q=" + texttobetranslated + "&target=" + to + "&source=" + from + "&format=html&key=" + env.APIKey)
	// create an http.Client and make the call to the Google Translate API
	client := urlfetch.Client(ctx)
	response, err := client.Get(apiurl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the API response
	jstr := getBody(response)
	// unmarshal the JSON string into a nice struct
	var jsonresponse GoogleTranslateAPIResponse
	err = json.Unmarshal([]byte(jstr), &jsonresponse)
	if err != nil {
		fmt.Fprintf(w, "\n\nError unmarshalling JSON, err=%s, response.Body= %s", err.Error(), r.Body)
		return
	}
	// print out the "from"
	fmt.Fprintf(w, "*Original* %s = %s\n\n", emoji_from, texttobetranslated)
	// print out the "to"
	for _, z := range jsonresponse.Data.Translations {
		fmt.Fprintf(w, "*Translated* %s = %s\n\n", emoji_to, z.TranslatedText)
	}
	return
}
