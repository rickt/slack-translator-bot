package slackdictionarybot

import (
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"net/http"
	"net/url"
	"os"
)

// globals
var (
	env envVars
)

// there's no main() func in google app engine
func init() {
	// get runtime options from the app.yaml
	env.APIBaseURL = os.Getenv("OXFORD_DICTIONARY_API_BASEURL")
	env.APPID = os.Getenv("OXFORD_DICTIONARY_APP_ID")
	env.APPKey = os.Getenv("OXFORD_DICTIONARY_APP_KEY")
	env.APIPath = os.Getenv("OXFORD_DICTIONARY_API_PATH")
	env.VerifyToken = os.Getenv("SLACK_VERIFY_TOKEN")
	// setup http handlers
	http.HandleFunc("/", handler_redirect)
	http.HandleFunc("/lookup", handler_lookup)
}

// redirect root requests to /lookup
func handler_redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/lookup", 302)
}

func handler_lookup(w http.ResponseWriter, r *http.Request) {
	// vars
	var word string
	var apiurl *url.URL
	// create a google app engine context
	ctx := appengine.NewContext(r)
	// what is our OG query from slack?
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "\n\nError parsing form, err=%s, rsp.Body= %s", err.Error(), r.Body)
		return
	}
	defer r.Body.Close()
	// what is the word to lookup?
	word = r.URL.Query().Get("text")
	// build the url! an example call to the Oxford Dictionary API:
	// https://od-api.oxforddictionaries.com:443/api/v1/entries/en/<word>
	apiurl, err = url.Parse(env.APIBaseURL)
	if err != nil {
		fmt.Fprintf(w, "\n\nError parsing URL = %s\n", err.Error())
		return
	}
	apiurl.Path += env.APIPath
	apiurl.Path += word
	// create an http.Client 
	client := urlfetch.Client(ctx)
	// create an http NewRequest so we can add the headers that the API requires
	req, err := http.NewRequest("GET", apiurl.String(), nil)
	if err != nil {
		fmt.Fprintf(w, "\n\nError creating http.NewRequest = %s\n", err.Error())
		return
	}
	// add the headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("app_id", env.APPID)
	req.Header.Add("app_key", env.APPKey)
	// make the request
	response, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "\n\nError making http.NewRequest = %s\n", err.Error())
		return
	}
	if response.StatusCode != 200 {
		fmt.Fprintf(w, "Error: no such word `%s`\n", word)
		return
	}
	// request was successful, so decode the response body into our nice JSON struct
	var odr OxfordReply
	json.NewDecoder(response.Body).Decode(&odr)
	fmt.Fprintf(w, "The Oxford Dictionary definition of `%s` is:\n", word)
	fmt.Fprintf(w, "```%s```\n", odr.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions[0])
	return
}

