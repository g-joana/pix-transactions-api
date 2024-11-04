package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ConvertPaymentToDynamoDBItem(payment Payment) map[string]*dynamodb.AttributeValue {
	item := map[string]*dynamodb.AttributeValue{
		"Object":       {S: aws.String(payment.Object)},
		"ID":           {S: aws.String(payment.ID)},
		"Customer":     {S: aws.String(payment.Customer)},
		"DateCreated":  {S: aws.String(payment.DateCreated)},
		"DueDate":      {S: aws.String(payment.DueDate)},
		"Installment":  {S: payment.Installment},
		"Subscription": {S: payment.Subscription},
		"PaymentLink":  {S: payment.PaymentLink},
		"Value":        {N: aws.String(fmt.Sprintf("%.2f", payment.Value))},
		"NetValue":     {N: aws.String(fmt.Sprintf("%.2f", payment.NetValue))},
		"BillingType":  {S: aws.String(payment.BillingType)},
		"Status":       {S: aws.String(payment.Status)},
		"Description":  {S: aws.String(payment.Description)},
		"DaysAfterDueDateToRegistrationCancellation": {
			N: aws.String("0.00"),
// N: payment.DaysAfterDueDateToRegistrationCancellation,
		},
		"ExternalReference":     {S: aws.String(payment.ExternalReference)},
		"CanBePaidAfterDueDate": {BOOL: aws.Bool(payment.CanBePaidAfterDueDate)},
		"PixTransaction":        {S: payment.PixTransaction},
		"PixQrCodeId":           {S: payment.PixQrCodeId},
		"OriginalValue":         {N: aws.String("0.00")},
		"InterestValue":         {N: aws.String("0.00")},
// {N: payment.InterestValue},
		"OriginalDueDate":       {S: aws.String(payment.OriginalDueDate)},
		"PaymentDate":           {S: payment.PaymentDate},
		"ClientPaymentDate":     {S: payment.ClientPaymentDate},
		"InstallmentNumber":     {N: aws.String("0.00")},
// {N: payment.InstallmentNumber},
		"TransactionReceiptUrl": {S: payment.TransactionReceiptUrl},
		"NossoNumero":           {S: aws.String(payment.NossoNumero)},
		"InvoiceUrl":            {S: aws.String(payment.InvoiceUrl)},
		"BankSlipUrl":           {S: aws.String(payment.BankSlipUrl)},
		"InvoiceNumber":         {S: aws.String(payment.InvoiceNumber)},
		"Discount":              ConvertDiscountToDynamoDBItem(payment.Discount),
		"Fine":                  ConvertFineToDynamoDBItem(payment.Fine),
		"Interest":              ConvertInterestToDynamoDBItem(payment.Interest),
		"Deleted":               {BOOL: aws.Bool(payment.Deleted)},
		"PostalService":         {BOOL: aws.Bool(payment.PostalService)},
		"Anticipated":           {BOOL: aws.Bool(payment.Anticipated)},
		"Anticipable":           {BOOL: aws.Bool(payment.Anticipable)},
		"CreditDate":            {S: aws.String(payment.CreditDate)},
		"EstimatedCreditDate":   {S: aws.String(payment.EstimatedCreditDate)},
		"Refunds":               {L: ConvertRefundsToDynamoDBItem(payment.Refunds)},
		"Split":                 {L: ConvertSplitsToDynamoDBItem(payment.Split)},
		"Chargeback":            ConvertChargebackToDynamoDBItem(payment.Chargeback),
	}

	return item
}

func ConvertDiscountToDynamoDBItem(discount Discount) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"Value":            {N: aws.String(fmt.Sprintf("%.2f", discount.Value))},
			"DueDateLimitDays": {N: aws.String(fmt.Sprintf("%d", discount.DueDateLimitDays))},
			"Type":             {S: aws.String(discount.Type)},
		},
	}
}

func ConvertFineToDynamoDBItem(fine Fine) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"Value": {N: aws.String(fmt.Sprintf("%.2f", fine.Value))},
		},
	}
}

func ConvertInterestToDynamoDBItem(interest Interest) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"Value": {N: aws.String(fmt.Sprintf("%.2f", interest.Value))},
		},
	}
}

func ConvertRefundsToDynamoDBItem(refunds []Refund) []*dynamodb.AttributeValue {
	var items []*dynamodb.AttributeValue
	for _, refund := range refunds {
		item := map[string]*dynamodb.AttributeValue{
			"DateCreated":           {S: aws.String(refund.DateCreated)},
			"Status":                {S: aws.String(refund.Status)},
			"Value":                 {N: aws.String(fmt.Sprintf("%.2f", refund.Value))},
			"EndToEndIdentifier":    {S: refund.EndToEndIdentifier},
			"Description":           {S: refund.Description},
			"EffectiveDate":         {S: aws.String(refund.EffectiveDate)},
			"TransactionReceiptUrl": {S: refund.TransactionReceiptUrl},
			"RefundedSplits":        {L: ConvertRefundedSplitsToDynamoDBItem(refund.RefundedSplits)},
		}
		items = append(items, &dynamodb.AttributeValue{M: item})
	}
	return items
}

func ConvertRefundedSplitsToDynamoDBItem(refundedSplits []RefundedSplit) []*dynamodb.AttributeValue {
	var items []*dynamodb.AttributeValue
	for _, split := range refundedSplits {
		item := map[string]*dynamodb.AttributeValue{
			"ID":    {S: aws.String(split.ID)},
			"Value": {N: aws.String(fmt.Sprintf("%.2f", split.Value))},
			"Done":  {BOOL: aws.Bool(split.Done)},
		}
		items = append(items, &dynamodb.AttributeValue{M: item})
	}
	return items
}

func ConvertSplitsToDynamoDBItem(splits []Split) []*dynamodb.AttributeValue {
	var items []*dynamodb.AttributeValue
	for _, split := range splits {
		item := map[string]*dynamodb.AttributeValue{
			"ID":                 {S: aws.String(split.ID)},
			"WalletId":           {S: aws.String(split.WalletId)},
			"FixedValue":         {N: aws.String(fmt.Sprintf("%.2f", split.FixedValue))},
			"PercentualValue":    {N: aws.String("0.00")},
// {N: split.PercentualValue},
			"TotalValue":         {N: aws.String(fmt.Sprintf("%.2f", split.TotalValue))},
			"CancellationReason": {S: aws.String(split.CancellationReason)},
			"Status":             {S: aws.String(split.Status)},
			"ExternalReference":  {S: split.ExternalReference},
			"Description":        {S: split.Description},
		}
		items = append(items, &dynamodb.AttributeValue{M: item})
	}
	return items
}

func ConvertChargebackToDynamoDBItem(chargeback Chargeback) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"Status": {S: aws.String(chargeback.Status)},
			"Reason": {S: aws.String(chargeback.Reason)},
		},
	}
}
