package api_requests

type EmbeddingsRequest struct {
}

type EmbeddingsResponse struct {
}

func Embedding(tokens []int) ([]int, error) {
	//request := EmbeddingsRequest{}
	//data, err := json.Marshal(request)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//
	//resp, err := http.Post(embeddingsUrl, "application/json", bytes.NewBuffer(data))
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

	return nil, nil
}
