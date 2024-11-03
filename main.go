package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	//api_requests.CreateDataForRag()
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "some-user", "12345", "dataset")

	db, err := sql.Open("postgres", psqlconn)
	PanicIfErr(err)
	err = db.Ping()
	PanicIfErr(err)

	mux := http.NewServeMux()
	mux.HandleFunc("/test", Ok)
	mux.HandleFunc("POST /message", HandlePrompt)

	err = http.ListenAndServe(":8080", mux)
	fmt.Println(err)

}

func Ok(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func HandlePrompt(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	var msg Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//_, err = rag.FindChunksTxt(msg.Message)
	// запрос к чату

	w.Write([]byte("Умный ответ от нейронки"))
}
