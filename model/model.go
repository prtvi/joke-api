package model

type Joke struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Punchline string `json:"punchline"`
	Setup     string `json:"setup"`
}

type Response struct {
	ResponseMsg string `json:"responseMsg"`
	StatusCode  int    `json:"statusCode"`
	Joke        Joke   `json:"joke"`
}

type ErrResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
