package main


type operation int
const (
	add	operation = iota
	remove
	update
)

type Customer struct {
	Object              string `json:"object"`
	ID                  string `json:"id"`
	DateCreated         string `json:"dateCreated"`
	Name                string `json:"name"`
	Email               string `json:"email,omitempty"`
	Phone               string `json:"phone,omitempty"`
	MobilePhone         string `json:"mobilePhone,omitempty"`
	Address             string `json:"address,omitempty"`
	AddressNumber       string `json:"addressNumber,omitempty"`
	Complement          string `json:"complement,omitempty"`
	Province            string `json:"province,omitempty"`
	City                string `json:"city,omitempty"`
	CityName            string `json:"cityName,omitempty"`
	State               string `json:"state,omitempty"`
	Country             string `json:"country,omitempty"`
	PostalCode          string `json:"postalCode,omitempty"`
	CPFCNPJ		    string `json:"cpfCnpj"`
	PersonType          string `json:"personType"`
	Deleted             bool   `json:"deleted"`
	AdditionalEmails    string `json:"additionalEmails,omitempty"`
	ExternalReference    string `json:"externalReference,omitempty"`
	NotificationDisabled bool   `json:"notificationDisabled"`
	Observations        string `json:"observations,omitempty"`
	ForeignCustomer     bool   `json:"foreignCustomer"`
}

type Payment struct {
	Object                              string         `json:"object"`
	ID                                  string         `json:"id"`
	Customer                            string         `json:"customer"`
	DateCreated                         string         `json:"dateCreated"`
	DueDate                             string         `json:"dueDate"`
	Installment                         *string        `json:"installment,omitempty"`
	Subscription                        *string        `json:"subscription,omitempty"`
	PaymentLink                         *string        `json:"paymentLink,omitempty"`
	Value                               float64        `json:"value"`
	NetValue                            float64        `json:"netValue"`
	BillingType                         string         `json:"billingType"`
	Status                              string         `json:"status"`
	Description                         string         `json:"description"`
	DaysAfterDueDateToRegistrationCancellation *int  `json:"daysAfterDueDateToRegistrationCancellation,omitempty"`
	ExternalReference                   string         `json:"externalReference"`
	CanBePaidAfterDueDate              bool           `json:"canBePaidAfterDueDate"`
	PixTransaction                      *string        `json:"pixTransaction,omitempty"`
	PixQrCodeId                        *string        `json:"pixQrCodeId,omitempty"`
	OriginalValue                       *float64       `json:"originalValue,omitempty"`
	InterestValue                       *float64       `json:"interestValue,omitempty"`
	OriginalDueDate                    string         `json:"originalDueDate"`
	PaymentDate                         *string        `json:"paymentDate,omitempty"`
	ClientPaymentDate                   *string        `json:"clientPaymentDate,omitempty"`
	InstallmentNumber                   *int           `json:"installmentNumber,omitempty"`
	TransactionReceiptUrl               *string        `json:"transactionReceiptUrl,omitempty"`
	NossoNumero                         string         `json:"nossoNumero"`
	InvoiceUrl                          string         `json:"invoiceUrl"`
	BankSlipUrl                         string         `json:"bankSlipUrl"`
	InvoiceNumber                       string         `json:"invoiceNumber"`
	Discount                            Discount       `json:"discount"`
	Fine                                Fine           `json:"fine"`
	Interest                             Interest      `json:"interest"`
	Deleted                             bool           `json:"deleted"`
	PostalService                       bool           `json:"postalService"`
	Anticipated                         bool           `json:"anticipated"`
	Anticipable                         bool           `json:"anticipable"`
	CreditDate                          string         `json:"creditDate"`
	EstimatedCreditDate                 string         `json:"estimatedCreditDate"`
	Refunds                             []Refund       `json:"refunds"`
	Split                               []Split        `json:"split"`
	Chargeback                          Chargeback     `json:"chargeback"`
}

type Discount struct {
	Value             float64 `json:"value"`
	DueDateLimitDays int     `json:"dueDateLimitDays"`
	Type              string  `json:"type"`
}

type Fine struct {
	Value float64 `json:"value"`
}

type Interest struct {
	Value float64 `json:"value"`
}

type Refund struct {
	DateCreated           string         `json:"dateCreated"`
	Status                string         `json:"status"`
	Value                 float64        `json:"value"`
	EndToEndIdentifier    *string        `json:"endToEndIdentifier,omitempty"` // Pointer for nullability
	Description           *string        `json:"description,omitempty"` // Pointer for nullability
	EffectiveDate         string         `json:"effectiveDate"`
	TransactionReceiptUrl *string        `json:"transactionReceiptUrl,omitempty"` // Pointer for nullability
	RefundedSplits        []RefundedSplit `json:"refundedSplits"`
}

type RefundedSplit struct {
	ID    string  `json:"id"`
	Value float64 `json:"value"`
	Done  bool    `json:"done"`
}

type Split struct {
	ID                 string   `json:"id"`
	WalletId           string   `json:"walletId"`
	FixedValue         float64  `json:"fixedValue"`
	PercentualValue    *float64 `json:"percentualValue,omitempty"`
	TotalValue         float64  `json:"totalValue"`
	CancellationReason string   `json:"cancellationReason"`
	Status             string   `json:"status"`
	ExternalReference  *string  `json:"externalReference,omitempty"`
	Description        *string  `json:"description,omitempty"`
}

type Chargeback struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type Webhook struct {
    ID           string   `json:"id"`
    Name         string   `json:"name"`
    URL          string   `json:"url"`
    Email        string   `json:"email"`
    Enabled      bool     `json:"enabled"`
    Interrupted  bool     `json:"interrupted"`
    APIVersion   int      `json:"apiVersion"`
    AuthToken    string   `json:"authToken"`
    SendType     string   `json:"sendType"`
    Events       []string `json:"events"`
}

type WebhookRequest struct {
    Id   string      `json:"id"`
    Event   string      `json:"event"`
    DateCreated   string      `json:"dateCreated"`
    Payment    Payment
}
