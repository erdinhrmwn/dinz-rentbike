package contract

import (
	"context"
)

type InvoiceRequest struct {
	ReferenceID string
	Amount      float64
	Currency    string
	PayerEmail  string
	Description string
}

type InvoiceResponse struct {
	InvoiceID  string
	InvoiceURL string
}

type XenditService interface {
	CreateInvoice(ctx context.Context, req *InvoiceRequest) (*InvoiceResponse, error)
	CancelInvoice(ctx context.Context, invoiceID string) error
}
