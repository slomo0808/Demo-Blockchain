package main

import (
	"Demo-Blockchain/core"
	"encoding/json"
	"io"
	"net/http"
)

var blockchain *core.Blockchain

func run() {
	http.HandleFunc("/blockchain/get", blockchainGetHandler)
	http.HandleFunc("/blockchain/write", blockchainWriteHandler)
	http.ListenAndServe(":8888", nil)
}

func blockchainGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(blockchain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "application:json")
	io.WriteString(w, string(bytes))
}

func blockchainWriteHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	blockchain.SendDate(data)
	blockchainGetHandler(w, r)
}

func main() {
	blockchain = core.NewBlockChain()
	run()
}