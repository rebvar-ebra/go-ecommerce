package main
import (
	"encoding/json"
	"log"
	"net/http"
)
type Message struct {
	Text string `json:"message"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	response :=Message{Text: "Welcome to Go E-Commerce API!"}
	w.Header().Set("Content-Type","application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil{
		log.Println("Error encoding response:", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}