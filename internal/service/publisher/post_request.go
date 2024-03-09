package publisher

type postRequest struct {
	ChatId                string `json:"chat_id"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
	Text                  string `json:"text"`
	Media                 string `json:"media"`
	Photo                 string `json:"photo"`
	Video                 string `json:"video"`
	Caption               string `json:"caption"`
}

type PostMediaField struct {
	Type     string `json:"type"`
	ParseMod string `json:"parse_mode"`
	Caption  string `json:"caption"`
	Media    string `json:"media"`
}
