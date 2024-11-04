package main

import (
	"fmt" // format - println
	"strings" // reader
	"io" // read all?
	"encoding/json" // marshall / unmarshall
	"errors"
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	_"github.com/aws/aws-lambda-go/lambda"
	"context"
	// "context" //?
)

//
// type PaymentInfo struct {
//     ID            string  `json:"id"`
//     Value         float64 `json:"value"`
//     // add others?
// }
//
// type WebhookRequest struct {
//     // Defina a estrutura de dados que você espera receber
//     EventType string `json:"event_type"`
//     Data      string `json:"data"`
// }

// lambda func:
func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    var webhookReq WebhookRequest

    // Decodificar o corpo do webhook
    err := json.Unmarshal([]byte(req.Body), &webhookReq)
    if err != nil { return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
    }

    // Processar o webhook (exemplo simples)
    fmt.Printf("Recebido evento: %s com dados: %v\n", webhookReq.Event, webhookReq.Payment)

	// switch webhookReq.Event {
	// case "PAYMENT_CREATED":
	// 	createPayment(webhookReq.Payment)
	// case "PAYMENT_RECEIVED":
	// 	receivePayment(webhookReq.Payment)
	// case "PAYMENT_OVERDUE":
	// 	attPayment(webhookReq.Payment)
	// case "PAYMENT_DELETED":
	// 	removePayment(webhookReq.Payment)
	// default:
	// 	fmt.Printf("Este evento não é aceito: %s\n", webhookReq.Event)
	// }

    // Retornar uma resposta
    return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       "Webhook recebido com sucesso",
    }, nil
}
/*
func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var payload PaymentINFO
	// unmarshal reads from a slice of bytes
	// decoder reads json from a reader.io
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Payload inválido", http.StatusBadRequest)
		return
	}


	response := map[string]bool{"received": true}
	json.NewEncoder(w).Encode(response)
	// retornar resposta?

	// fmt.Printf("Webhook recebido: %s\n", payload.Message)
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Webhook recebido!"))
}
*/

// WIP:
func createWebhook() () {
	endpoint := "https://sandbox.asaas.com/api/v3/webhooks"
	payload := strings.NewReader("{\"name\":\"CASH IN\",\"url\":\"https://lazy-motorcycle-68.webhook.cool\",\"enabled\":true,\"interrupted\":false,\"apiVersion\":3,\"sendType\":\"SEQUENTIALLY\",\"events\":[\"PAYMENT_RECEIVED\",\"PAYMENT_CREATED\",\"PAYMENT_OVERDUE\"],\"email\":\"jou.42.rio@gmail.com\"}")
	res, _ := newRequest("POST", endpoint, payload)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
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

// func createStaticQRcode() {
// 	endpoint := "https://sandbox.asaas.com/api/v3/pix/qrCodes/static"
// 	payload := strings.NewReader("{\"addressKey\":\"cbabe4f8-65b8-4dae-9cfb-1577fccf7cd2\",\"value\":100,\"format\":\"ALL\",\"expirationSeconds\":120,\"allowsMultiplePayments\":false}")
// 	newRequest("POST", endpoint, payload)
// }
//
// func createPixKey() {
// 	endpoint := "https://sandbox.asaas.com/api/v3/pix/addressKeys"
// 	payload := strings.NewReader("{\"type\":\"EVP\"}")
// 	newRequest("POST", endpoint, payload)
// }
