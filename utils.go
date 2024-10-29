package main

import (
	"fmt" // format - println
	"net/http"
	"strings" // reader
)

const (
    UserAgent      = "pixpayments"
    Accept         = "application/json"
    ContentType    = "application/json"
)

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
