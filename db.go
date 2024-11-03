package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var tableName string = "payments"

type operation int
const (
	add	operation = iota
	remove
	update
)

func addItem(svc *dynamodb.DynamoDB, info *PaymentInfo) error {
	item := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(info.id),},
		"payload": {
			S: aws.String(info.payload),},
		"": {
			S: aws.String(info),
		},
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	}

	_, err := svc.PutItem(input)
	return err
}

// TEST:
func updateTable(op operation, info *PaymentInfo) {
	// start aws session
	sess := session.Must(session.NewSession())

	// novo cliente dynamodb
	svc := dynamodb.New(sess)

	// Usar a função para adicionar itens
	if op == add {
		if err := addItem(svc, info); err != nil {
			log.Fatalf("Erro ao adicionar item: %s", err) // mudar para outra func
		}
	}
}
