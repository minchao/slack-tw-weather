package slack

type Attachment struct {
	Text     string `json:"text,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}
