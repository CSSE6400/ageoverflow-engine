package service

type Request struct {
	ID       string `json:"id"`
	Content  string `json:"content"`
	Metadata string `json:"metadata"`
}
