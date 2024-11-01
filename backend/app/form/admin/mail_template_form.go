package admin

type MailTemplateForm struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Type     int    `json:"type"`
	Subject  string `json:"subject"`
	Contents string `json:"contents"`
}
