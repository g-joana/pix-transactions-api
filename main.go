package main

import (
	"fmt" // format - println
	"net/http"
	"os"
	"log"

	// "github.com/aws/aws-sdk-go/aws/request"
	// "github.com/joho/godotenv"
)

var apiKey = os.Getenv("ASAASKEY")

func main() {
	// apiURL := "https://sandbox.asaas.com/api/v3"
	// err := godotenv.Load()

	// start server
	// STUDY:
	mux := http.NewServeMux()
	mux.HandleFunc("/webhook", router)
	http.ListenAndServe(":8080", mux)
	fmt.Println("server on")
}

func router(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("vc acessou a rota /webhook"))

	db := createDBconnection()
	if err := addItem(db.Svc, nil); err != nil {
		log.Fatalf("Erro ao adicionar item: %s", err) // mudar para outra func
	}
}
