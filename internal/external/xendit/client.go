package xendit

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"dinz-rentbike/internal/config"
	"dinz-rentbike/internal/domain/contract"
)

type Client struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func NewClient(cfg *config.XenditConfig) contract.XenditService {
	return &Client{
		baseURL: cfg.BaseURL,
		apiKey:  cfg.SecretKey,
		client:  &http.Client{Timeout: 15 * time.Second},
	}
}

func (c *Client) CreateInvoice(ctx context.Context, req *contract.InvoiceRequest) (*contract.InvoiceResponse, error) {
	body, err := json.Marshal(map[string]any{
		"reference_id": req.ReferenceID,
		"session_type": "PAY",
		"currency":     req.Currency,
		"amount":       req.Amount,
		"description":  req.Description,
		"mode":         "PAYMENT_LINK",
		"country":      "ID",
		"expires_at":   time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		"allowed_payment_channels": []string{
			"QRIS",
			"DANA",
			"SHOPEEPAY",
			"OVO",
			"GOPAY",
		},
		"metadata": map[string]string{
			"email": req.PayerEmail,
			"name":  req.PayerName,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/sessions", c.baseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Basic "+basicAuth(c.apiKey, ""))

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("xendit error [%d]: %s", resp.StatusCode, string(respBody))
	}

	var xenditResp struct {
		ID  string `json:"payment_session_id"`
		URL string `json:"payment_link_url"`
	}
	if err := json.Unmarshal(respBody, &xenditResp); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &contract.InvoiceResponse{
		InvoiceID:  xenditResp.ID,
		InvoiceURL: xenditResp.URL,
	}, nil
}

func (c *Client) CancelInvoice(ctx context.Context, invoiceID string) error {
	url := fmt.Sprintf("%s/sessions/%s/cancel", c.baseURL, invoiceID)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	httpReq.Header.Set("Authorization", "Basic "+basicAuth(c.apiKey, ""))

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("xendit error [%d]: %s", resp.StatusCode, string(respBody))
	}

	return nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
