package api_requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type Chunk struct {
	Text string `json:"text"`
}

type Chunks struct {
	Chunks []Chunk `json:"chunks"`
}

type ResultObject struct {
	Text   string    `json:"text"`
	Vector []float64 `json:"vector"`
}

type Result struct {
	Data []ResultObject `json:"data"`
}

type GetResponses struct {
	mux    sync.Mutex
	result Result
}

func CreateDataForRag() {
	chunks, err := readChunksFile()
	fmt.Println(len(chunks.Chunks))
	if err != nil {
		fmt.Println("Error reading chunks file:", err)
		return
	}

	getResponses := GetResponses{mux: sync.Mutex{}, result: Result{Data: make([]ResultObject, 0)}}
	var i int
	wg := new(sync.WaitGroup)
	poolSize := 50
	for i < 8000 {
		wg.Add(poolSize)
		for j := i; j < i+poolSize; j++ {
			if j >= len(chunks.Chunks) {
				wg.Done()
				continue
			}
			go func(text string) {
				defer wg.Done()
				vector, err := TextEmbedding(text)
				if err != nil {
					return
				}
				resObj := ResultObject{Text: text, Vector: vector}
				getResponses.mux.Lock()
				getResponses.result.Data = append(getResponses.result.Data, resObj)
				getResponses.mux.Unlock()
			}(chunks.Chunks[j].Text)
		}
		wg.Wait()
		i += poolSize
		fmt.Println(i)

	}

	//for i, chunk := range chunks.Chunks {
	//	vector, err := TextEmbedding(chunk.Text)
	//	fmt.Println(i, vector)
	//	if err != nil {
	//		continue
	//	}
	//	resObj := ResultObject{Text: chunk.Text, Vector: vector}
	//	res.Data = append(res.Data, resObj)
	//	if i == 10 {
	//		break
	//	}
	//}

	file, err := os.Create("data.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(getResponses.result)
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
