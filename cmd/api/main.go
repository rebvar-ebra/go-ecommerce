package main

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"strings"
)

type Message struct {
	Text string `json:"message"`
}
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

var products = []Product{
	{ID: 1, Name: "Laptop", Price: 899.99, Stock: 20},
	{ID: 2, Name: "Mouse", Price: 19.99, Stock: 50},
	{ID: 3, Name: "Keyboard", Price: 5.99, Stock: 200},
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := Message{Text: "Welcome to Go E-Commerce API!"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error encoding response:", err)
	}
}
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Println("Error encoding products:", err)
	}
}
func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock int     `json:"stock"`
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1<<30) // Limit request body to 1MB
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Disallow unknown fields in the JSON payload
	  if err := decoder.Decode(&newProduct); err != nil {
        log.Println("JSON decode error:", err)
        http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
        return
    }
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		http.Error(w, "Request body must contain one valid JSON object", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(newProduct.Name) == "" {
		http.Error(w, "Product name is required", http.StatusBadRequest)
		return
	}
	if math.IsNaN(newProduct.Price) || math.IsInf(newProduct.Price, 0) || newProduct.Price < 0 {
		http.Error(w, "Product price must be a non-negative number", http.StatusBadRequest)
		return

	}
	if newProduct.Stock < 0 {
		http.Error(w, "Product stock must be a non-negative integer", http.StatusBadRequest)
		return
	}
	product := Product{
		ID:    len(products) + 1,
		Name:  strings.TrimSpace(newProduct.Name),
		Price: newProduct.Price,
		Stock: newProduct.Stock,
	}
	products = append(products, product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		log.Println("Error encoding product:", err)
	}
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET/", homeHandler)
	mux.HandleFunc("GET /products", getProducts)
	mux.HandleFunc("POST /products", createProduct)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
