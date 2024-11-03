package api_requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TokenizeRequest struct {
	Model            string `json:"model"`
	Prompt           string `json:"prompt"`
	AddSpecialTokens bool   `json:"add_special_tokens"`
}

type TokenizeResponse struct {
	Count       int   `json:"count"`
	MaxModelLen int   `json:"max_model_len"`
	Tokens      []int `json:"tokens"`
}

func Tokenize(chunk string) ([]int, error) {
	var response TokenizeResponse
	var request TokenizeRequest
	request = TokenizeRequest{Model: modelName, Prompt: chunk, AddSpecialTokens: true}
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("tokenize err: %v", err)
		return nil, err
	}

	resp, err := http.Post(tokenizeUrl, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("tokenize err: %v", err)

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("tokenize err: %v", err)

		return nil, errors.New("")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("tokenize err: %v", err)

		return nil, err
	}

	err = json.Unmarshal(body, &response)
	tokens := response.Tokens
	if err != nil {
		fmt.Printf("tokenize err: %v", err)

		return nil, err
	}
	return tokens, nil
}
