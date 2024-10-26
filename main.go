package main

import (
	"fmt" // format - println
	"os" // getenv
	"net/http"
	"strings" // reader
	"io" // read all?
	"encoding/json" // marshall / unmarshall
	"errors"
)

const (
    UserAgent      = "convem"
    Accept         = "application/json"
    ContentType    = "application/json"
)

var apiKey = os.Getenv("ASAASKEY")

type Customer struct {
    Object      string `json:"object"`
    Id          string `json:"id"`
    DateCreated string `json:"dateCreated"`
    Name        string `json:"name"`
}

type Charge struct {
    Object      string `json:"object"`
    Id          string `json:"id"`
    DateCreated string `json:"dateCreated"`
    Name        string `json:"name"`
}

func main() {
	// apiURL := "https://sandbox.asaas.com/api/v3"

	// start server
	mux := http.NewServeMux()
	mux.HandleFunc("/main", router)
	http.ListenAndServe(":8080", mux)
	fmt.Println("server on")
}

func router(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("vc acessou a rota /main"));

	// createQRcode()
	// createPixKey()
	// createCustomer()
}

//TEST:
func newRequest(method string, endpoint string, payload *strings.Reader) (*http.Response, error) {

	req, err := http.NewRequest(method, endpoint, payload)
	if err != nil {
		fmt.Println("\nErro ao criar requisição:\n", err)
		return nil, err
	}

	// Header.Add appends, while Set over-writes
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("accept", Accept)	
	req.Header.Set("content-type", ContentType)
	req.Header.Set("access_token", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("\nErro: sem resposta do cliente\n", err) //?
		return nil, err
	}
	fmt.Printf("Response: \n%v\n",res)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("\nErro: body deu ruim\n", err) //?
		return nil, err
	}
	fmt.Println(string(body))
	return res, nil
}

//TEST:
func printResBody(res *http.Response) error {
	
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("\nErro ao ler body\n", err) //?
		return errors.New("erro leitura do body da req")
	}
	fmt.Println(string(body))
	return nil
}

//TEST:
func createCustomer() (*Customer, error) {
	
	endpoint := "https://sandbox.asaas.com/api/v3/customers"
	payload := strings.NewReader("{\"name\":\"Cliente\",\"cpfCnpj\":\"483.035.160-86\"}")
	res, _ := newRequest("POST", endpoint, payload)
	var customer *Customer
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	err := json.Unmarshal([]byte(body), customer)
	if err != nil {
		fmt.Println("\nErro: deu ruim no customer\n", err) //?
		return nil, errors.New("customer deu ruim")
	}
	return customer, nil
}

//IN PROGRESS:
func createCharge(customerId string) Charge {

	endpoint := "https://sandbox.asaas.com/api/v3/payments"
	payload := strings.NewReader("{\"billingType\":\"PIX\",\"customer\":" + customerId + ",\"value\":100,\"dueDate\":\"2025-01-01\"}")
	newRequest("POST", endpoint, payload)
}

func createQRcode() {

}

func createStaticQRcode() {
	endpoint := "https://sandbox.asaas.com/api/v3/pix/qrCodes/static"
	payload := strings.NewReader("{\"addressKey\":\"cbabe4f8-65b8-4dae-9cfb-1577fccf7cd2\",\"value\":100,\"format\":\"ALL\",\"expirationSeconds\":120,\"allowsMultiplePayments\":false}")
	newRequest("POST", endpoint, payload)
}

func createPixKey() {
	endpoint := "https://sandbox.asaas.com/api/v3/pix/addressKeys"
	payload := strings.NewReader("{\"type\":\"EVP\"}")
	newRequest("POST", endpoint, payload)
}
