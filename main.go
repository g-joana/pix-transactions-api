package main

import (
	"fmt" // format - println
	"os"
	"net/http"
	// "github.com/joho/godotenv"
)

var apiKey = os.Getenv("ASAASKEY")

type Customer struct {
    Object      string `json:"object"`
    Id          string `json:"id"`
    DateCreated string `json:"dateCreated"`
    Name        string `json:"name"`
}

type Payment struct {
    Object      string `json:"object"`
    Id          string `json:"id"`
    DateCreated string `json:"dateCreated"`
    Name        string `json:"name"`
}

// fix fields:
type QRcode struct {
    Object      string `json:"object"`
    Id          string `json:"id"`
    DateCreated string `json:"dateCreated"`
    Name        string `json:"name"`
}

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
	w.Write([]byte("vc acessou a rota /webhook"));

	// createQRcode()
	// createPixKey()
	// customer, err := createCustomer()
	// if err != nil {
	// 	fmt.Println("\nErro: nao criou customer\n", err) //?
	// 	return
	// }
	// payment, err := createPayment(customer.Id)
	// if err != nil {
	// 	fmt.Println("\nErro: nao criou pagamento\n", err) //?
	// 	return
	// }
	// fmt.Println("")
	// fmt.Println(payment.Id)
	// getQRcode(payment.Id)
	// fmt.Println("")
	// createWebhook()
}
