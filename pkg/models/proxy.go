package models

type ProxyRequest struct {
    Method  string            `json:"method" example:"GET"`
    URL     string            `json:"url" example:"http://google.com"`
	Headers map[string]string `json:"headers" example:"Authorization:Basic bG9naW46cGFzc3dvcmQ="`
}

type ProxyResponse struct {
	ID      string              `json:"id"`
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int                 `json:"method"`
}
