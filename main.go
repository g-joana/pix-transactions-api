package main

import (
	"fmt" // format - println
	"os" // getenv
	"net/http"
	"strings" // reader
	"io" // read all?
	"encoding/json" // marshall / unmarshall
	"errors"
	// "github.com/joho/godotenv"
)

const (
    UserAgent      = "pixpayments"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/main", router)
	http.ListenAndServe(":8080", mux)
	fmt.Println("server on")
}

func router(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("vc acessou a rota /main"));

	// createQRcode()
	// createPixKey()
	customer, err := createCustomer()
	if err != nil {
		fmt.Println("\nErro: nao criou customer\n", err) //?
		return
	}
	charge, err := createCharge(customer.Id)
	if err != nil {
		fmt.Println("\nErro: nao criou charge\n", err) //?
		return
	}
	fmt.Println("")
	fmt.Println(charge.Id)
	getQRcode(charge.Id)
	fmt.Println("")
}

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
	return res, nil
}

func createCustomer() (*Customer, error) {
	
	endpoint := "https://sandbox.asaas.com/api/v3/customers"
	payload := strings.NewReader("{\"name\":\"Cliente\",\"cpfCnpj\":\"483.035.160-86\"}")
	res, _ := newRequest("POST", endpoint, payload)

	var customer = Customer{}
	defer res.Body.Close()
	body, e := io.ReadAll(res.Body)
	if e != nil {
		fmt.Println("\nErro: deu ruim lendo body\n", e) //?
		return nil, errors.New("body")
	}
	fmt.Println(string(body))
	err := json.Unmarshal([]byte(body), &customer)
	if err != nil {
		fmt.Println("\nErro: deu ruim no customer\n", err) //?
		return nil, errors.New("customer deu ruim")
	}
	return &customer, nil
}

func createCharge(customerId string) (*Charge, error) {

	endpoint := "https://sandbox.asaas.com/api/v3/payments"
	payload := strings.NewReader("{\"billingType\":\"PIX\",\"customer\":" + customerId + ",\"value\":100,\"dueDate\":\"2025-01-01\"}")
	res, _ := newRequest("POST", endpoint, payload)

	defer res.Body.Close()
	body, e := io.ReadAll(res.Body)
	if e != nil {
		fmt.Println("\nErro: deu ruim lendo body\n", e) //?
		return nil, errors.New("body")
	}
	fmt.Println("")
	fmt.Println(string(body))
	var charge = Charge{}
	err := json.Unmarshal([]byte(body), &charge)
	if err != nil {
		fmt.Println("\nErro: deu ruim no customer\n", err) //?
		return nil, errors.New("customer deu ruim")
	}
	return &charge, nil
}

func getQRcode(chargeId string) () {
	
	endpoint := "https://sandbox.asaas.com/api/v3/payments/" + chargeId + "/pixQrCode"
	payload := strings.NewReader("")
	res, err := newRequest("GET", endpoint, payload)
	if err != nil {
		fmt.Println("\nErro: deu ruim lendo body\n", err) //?
		// return nil, errors.New("body")
		return
	}
	defer res.Body.Close()
	body, e := io.ReadAll(res.Body)
	if e != nil {
		fmt.Println("\nErro: deu ruim lendo body\n", e) //?
		// return nil, errors.New("body")
		return
	}
	fmt.Println("")
	fmt.Println(string(body))
	// var QRcode = QRcode{}
	// err := json.Unmarshal([]byte(body), &QRcode)
	// if err != nil {
	// 	fmt.Println("\nErro: deu ruim no customer\n", err) //?
	// 	// return nil, errors.New("customer deu ruim")
	// }
	// return &charge, nil
	return
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
