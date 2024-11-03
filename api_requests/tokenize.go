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
	AddSpecialTokens bool    `json:"add_special_tokens"`
	Inputs           string  `json:"inputs"`
	PromptName       *string `json:"prompt_name"`
}

type TokenizeResponse [][]TokenizeResponseObj

type TokenizeResponseObj struct {
	Id      int    `json:"id"`
	Special bool   `json:"special"`
	Start   int    `json:"start"`
	Stop    int    `json:"stop"`
	Text    string `json:"text"`
}

func Tokenize(chunk string) ([]int, error) {
	request := TokenizeRequest{AddSpecialTokens: true, Inputs: chunk}
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

	var response TokenizeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("tokenize err: %v", err)
		return nil, err
	}
	tokens := make([]int, len(response[0]))
	for _, obj := range response[0] {
		tokens = append(tokens, obj.Id)
	}
	return tokens, nil
}
