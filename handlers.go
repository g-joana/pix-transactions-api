package main

import (
	"fmt" // format - println
	"strings" // reader
	"io" // read all?
	"encoding/json" // marshall / unmarshall
	"errors"
)

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

func createPayment(customerId string) (*Payment, error) {

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
	var payment = Payment{}
	err := json.Unmarshal([]byte(body), &payment)
	if err != nil {
		fmt.Println("\nErro: deu ruim no customer\n", err) //?
		return nil, errors.New("customer deu ruim")
	}
	return &payment, nil
}

func getQRcode(paymentId string) () {
	
	endpoint := "https://sandbox.asaas.com/api/v3/payments/" + paymentId + "/pixQrCode"
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
