package slackdictionarybot

// struct for basic JSON request to be sent to Oxford Dictionary API
type OxfordRequest struct {
	Type	string	`json:"Accept"`
	App_ID	string	`json:"app_id"`
	App_Key	string	`json:"app_key"`
}

// struct for Oxford Dictionary API reply
type OxfordReply struct {
	Metadata struct {
		Provider string `json:"provider"`
	} `json:"metadata"`
	Results []struct {
		ID             string `json:"id"`
		Language       string `json:"language"`
		LexicalEntries []struct {
			Entries []struct {
				Etymologies         []string `json:"etymologies"`
				GrammaticalFeatures []struct {
					Text string `json:"text"`
					Type string `json:"type"`
				} `json:"grammaticalFeatures"`
				HomographNumber string `json:"homographNumber"`
				Senses          []struct {
					Definitions []string `json:"definitions"`
					Domains     []string `json:"domains"`
					Examples    []struct {
						Text string `json:"text"`
					} `json:"examples"`
					ID               string   `json:"id"`
					ShortDefinitions []string `json:"short_definitions"`
					Subsenses        []struct {
						Definitions []string `json:"definitions"`
						Examples    []struct {
							Text string `json:"text"`
						} `json:"examples"`
						ID               string   `json:"id"`
						ShortDefinitions []string `json:"short_definitions"`
						ThesaurusLinks   []struct {
							EntryID string `json:"entry_id"`
							SenseID string `json:"sense_id"`
						} `json:"thesaurusLinks"`
						Registers []string `json:"registers,omitempty"`
					} `json:"subsenses"`
					ThesaurusLinks []struct {
						EntryID string `json:"entry_id"`
						SenseID string `json:"sense_id"`
					} `json:"thesaurusLinks"`
				} `json:"senses"`
			} `json:"entries"`
			Language        string `json:"language"`
			LexicalCategory string `json:"lexicalCategory"`
			Pronunciations  []struct {
				AudioFile        string   `json:"audioFile"`
				Dialects         []string `json:"dialects"`
				PhoneticNotation string   `json:"phoneticNotation"`
				PhoneticSpelling string   `json:"phoneticSpelling"`
			} `json:"pronunciations"`
			Text string `json:"text"`
		} `json:"lexicalEntries"`
		Type string `json:"type"`
		Word string `json:"word"`
	} `json:"results"`
}

// struct for runtime environment variables
type envVars struct {
	APIBaseURL  string
	APPID	      string
	APPKey      string
	APIPath     string
	VerifyToken string
}

// struct for forming a slack request
type slackRequest struct {
	Token       string `schema:"token"`
	TeamID      string `schema:"team_id"`
	TeamDomain  string `schema:"team_domain"`
	ChannelID   string `schema:"channel_id"`
	ServiceID   string `schema:"service_id"`
	ChannelName string `schema:"channel_name"`
	Timestamp   string `schema:"timestamp"`
	UserID      string `schema:"user_id"`
	UserName    string `schema:"user_name"`
	Text        string `schema:"text"`
	TriggerWord string `schema:"trigger_word"`
}

// structs for slack inbound webhook message
type Payload struct {
	Channel      string       `json:"channel"`
	Username     string       `json:"username"`
	Text         string       `json:"text"`
	ResponseType string       `json:"response_type"`
	Icon_emoji   string       `json:"icon_emoji"`
	Unfurl_links bool         `json:"unfurl_links"`
	Attachments  []Attachment `json:"attachments"`
}
type Attachment struct {
	Fallback   string  `json:"fallback"`
	Pretext    string  `json:"pretext"`
	Color      string  `json:"color"`
	AuthorName string  `json:"author_name"`
	AuthorLink string  `json:"author_link"`
	AuthorIcon string  `json:"author_icon"`
	Title      string  `json:"title"`
	TitleLink  string  `json:"title_link"`
	Text       string  `json:"text"`
	Fields     []Field `json:"fields"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// EOF
