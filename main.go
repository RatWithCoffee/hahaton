package main

import (
	"hahaton/api_requests"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	api_requests.CreateDataForRag()
}
