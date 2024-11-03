package api_requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EmbedAllRequest struct {
	Inputs            string  `json:"inputs"`
	PromptName        *string `json:"prompt_name"`
	Truncate          bool    `json:"truncate"`
	TruncateDirection string  `json:"truncation_direction"`
}

type EmbedAllResponse [][][]int

func TextEmbedding(text string) ([]int, error) {
	request := EmbedAllRequest{Inputs: text, TruncateDirection: "Right"}
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp, err := http.Post(embedAllUrl, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("embedding err: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("embedding err: %v", err)
		return nil, errors.New("")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("embedding err: %v", err)
		return nil, err
	}

	var vector EmbedAllResponse
	err = json.Unmarshal(body, &vector)
	if err != nil {
		fmt.Printf("embedding err: %v", err)
		return nil, err
	}

	return vector[0][0], nil
}

//func Embedding(tokens []int) ([]int, error) {
//request := EmbedAllRequest{}
//data, err := json.Marshal(request)
//if err != nil {
//	fmt.Println(err)
//	return nil, err
//}
//
//resp, err := http.Post(embedAllUrl, "application/json", bytes.NewBuffer(data))
//if err != nil {
//	fmt.Println(err)
//	return nil, err
//}
//defer resp.Body.Close()
//
//if resp.StatusCode != http.StatusOK {
//	fmt.Println(resp.StatusCode)
//}
//body, err := ioutil.ReadAll(resp.Body)
//if err != nil {
//	fmt.Println(err)
//	return nil, err
//}
//
//var vector []int
//err = json.Unmarshal(body, &vector)
//if err != nil {
//	fmt.Println("Error decoding JSON:", err)
//	return nil, err
//}
//
//return nil, nil
//}
