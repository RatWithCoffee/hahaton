package api_requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Chunk struct {
	Text string `json:"text"`
}

type Chunks struct {
	Chunks []Chunk `json:"chunks"`
}

type ResultObject struct {
	//Tokens    []int `json:"tokens"`
	Text      string `json:"text"`
	Vector    []int  `json:"vector"`
	VectorLen int    `json:"vector_len"`
}

type Result struct {
	Count int            `json:"count"`
	Data  []ResultObject `json:"data"`
}

func CreateDataForRag() {
	chunks, err := readChunksFile()
	if err != nil {
		fmt.Println("Error reading chunks file:", err)
		return
	}
	res := Result{Data: make([]ResultObject, 0)}
	for _, chunk := range chunks.Chunks {
		//tokens, err := Tokenize(chunk.Text)
		//if err != nil {
		//	continue
		//}
		vector, err := TextEmbedding(chunk.Text)
		if err != nil {
			continue
		}
		resObj := ResultObject{Text: chunk.Text, Vector: vector, VectorLen: len(vector)}
		res.Data = append(res.Data, resObj)
	}
	res.Count = len(res.Data)

	file, err := os.Create("data.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(res)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}

func readChunksFile() (Chunks, error) {
	var chunks Chunks

	file, err := os.Open(chunksFileName)
	if err != nil {
		return chunks, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return chunks, err
	}

	err = json.Unmarshal(data, &chunks)
	if err != nil {
		return chunks, err
	}

	return chunks, nil
}
