package grpc_client

type News struct {
	Id       string `json:"id"`
	Category string `json:"category"`
	Datetime string `json:"datetime"`
	Headline string `json:"headline"`
	Image    string `json:"image"`
	Related  string `json:"related"`
	Resource string `json:"resource"`
	Summary  string `json:"summary"`
	Url      string `json:"url"`
}
