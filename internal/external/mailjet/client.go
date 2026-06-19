package mailjet

import (
	"bytes"
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
	secret  string
	client  *http.Client
}

func NewClient(cfg *config.MailjetConfig) contract.MailjetService {
	return &Client{
		baseURL: cfg.BaseURL,
		apiKey:  cfg.ApiKey,
		secret:  cfg.SecretKey,
		client:  &http.Client{Timeout: 15 * time.Second},
	}
}

func (c *Client) Send(req *contract.EmailRequest) error {
	body, err := json.Marshal(map[string]any{
		"Messages": []map[string]any{
			{
				"From": map[string]string{
					"Email": "noreply@erdinhrmwn.studio",
					"Name":  "Dinz RentBike",
				},
				"To": []map[string]string{
					{
						"Email": req.ToEmail,
						"Name":  req.ToName,
					},
				},
				"Subject":  req.Subject,
				"HTMLPart": req.HTMLBody,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, c.baseURL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Basic "+basicAuth(c.apiKey, c.secret))

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("mailjet error [%d]: %s", resp.StatusCode, string(respBody))
	}

	return nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
