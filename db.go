package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var paymentsTable string = "payments"

type DB struct {
	Svc *dynamodb.DynamoDB
}

func addItem(svc *dynamodb.DynamoDB, payment *Payment) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(paymentsTable),
		Item: map[string]*dynamodb.AttributeValue{
			"age":  {N: aws.String("30")},
		},
	}
	_, err := svc.PutItem(input)
	fmt.Println("Item added successfully")
	return err
}

func createDBconnection() *DB {
	// start aws session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1"),
	})
	if err != nil {
		fmt.Println("erro ao criar sessao")
	}

	return &DB{Svc: dynamodb.New(sess)}
}
