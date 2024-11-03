package api_requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DetokenizeRequest struct {
	Model  string `json:"model"`
	Tokens []int  `json:"tokens"`
}

type DetokenizeResponse struct {
}

func Detokenize(tokens []int) (string, error) {
	request := DetokenizeRequest{Model: modelName, Tokens: tokens}
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	resp, err := http.Post(detokenizeUrl, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("detokenize err, %d", resp.StatusCode)
		return "", errors.New("")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}
