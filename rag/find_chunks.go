package rag

import (
	"hahaton/api_requests"
	"math"
)

//type ResultObject struct {
//	Tokens    []int `json:"tokens"`
//	Vector    []int `json:"vector"`
//	VectorLen int   `json:"vector_len"`
//}

//
//type Result struct {
//	Count int            `json:"count"`
//	Data  []ResultObject `json:"data"`
//}

const delta = 0.1

const maxLenOfArrVectors = 10

func FindChunksTxt(msg string) ([]string, error) {
	tokens, err := api_requests.Tokenize(msg)
	if err != nil {
		return nil, err
	}
	vector, err := api_requests.Embedding(tokens)
	var dbData []api_requests.ResultObject
	var cos float64
	bestVectors := make([]api_requests.ResultObject, 0)
	texts := make([]string, 0)
	for _, data := range dbData {
		cos = findCos(vector, data.Vector)
		if 1-cos < delta {
			bestVectors = append(bestVectors, data)
			text, err := api_requests.Detokenize(data.Vector)
			if err == nil {
				texts = append(texts, text)
			}
		}
		if len(bestVectors) >= maxLenOfArrVectors {
			break
		}

	}
	return texts, nil
}

func findCos(a []int, b []int) float64 {
	var dotProduct, normA, normB float64
	maxLen := max(len(a), len(b))
	for i := 0; i < maxLen; i++ {
		dotProduct += float64(a[i] * b[i])
		if i < len(a) {
			normA += float64(a[i] * a[i])
		}
		if i < len(b) {
			normB += float64(b[i] * b[i])
		}
	}

	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}
