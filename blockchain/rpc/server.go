package main

import (
	"encoding/json"
	"net/http"
	"polaris/blockchain"
)

var chain *blockchain.Blockchain

func Run() {
	http.HandleFunc("/blockchain/get", blockChainGetHandler)
	http.HandleFunc("/blockchain/write", blockChainWriteHandler)
	http.ListenAndServe(":8080", nil)
}

func blockChainGetHandler(w http.ResponseWriter, _ *http.Request) {
	bytes, err := json.Marshal(chain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}

func blockChainWriteHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	chain.SendData(data)
	blockChainGetHandler(w, r)
}

// for example:
// curl http://localhost:8080/blockchain/get
// curl http://localhost:8080/blockchain/write?data=Send 1 BTC to 1302岁的龙猫
func main() {
	chain = blockchain.NewBlockChain()
	Run()
}
