package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

	"site-admin-api/config"
)

// CDNClient handles communication with CDN File Server
type CDNClient struct {
	baseURL string
	token   string
	timeout time.Duration
}

// CDNUploadResponse represents the response from CDN upload
type CDNUploadResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		FileID       string    `json:"file_id"`
		OriginalName string    `json:"original_name"`
		URL          string    `json:"url"`
		Tag          string    `json:"tag"`
		Size         int64     `json:"size"`
		ContentType  string    `json:"content_type"`
		Public       bool      `json:"public"`
		UploadedAt   time.Time `json:"uploaded_at"`
		UploadedBy   string    `json:"uploaded_by"`
	} `json:"data"`
	Error string `json:"error,omitempty"`
}

// CDNDeleteResponse represents the response from CDN delete
type CDNDeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// NewCDNClient creates a new CDN client
func NewCDNClient(cfg *config.Config) *CDNClient {
	baseURL := cfg.CDN.BaseURL
	if cfg.IsProduction() {
		baseURL = cfg.CDN.BaseURLProd
	}

	timeout := time.Duration(cfg.CDN.Timeout) * time.Second
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	return &CDNClient{
		baseURL: baseURL,
		token:   cfg.CDN.Token,
		timeout: timeout,
	}
}

// UploadFile uploads a file to CDN server
func (c *CDNClient) UploadFile(fileContent []byte, filename, tag string, isPublic bool) (*CDNUploadResponse, error) {
	// Create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file field
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := io.Copy(part, bytes.NewReader(fileContent)); err != nil {
		return nil, fmt.Errorf("failed to copy file content: %w", err)
	}

	// Add tag field
	if err := writer.WriteField("tag", tag); err != nil {
		return nil, fmt.Errorf("failed to write tag field: %w", err)
	}

	// Add public field
	publicStr := "false"
	if isPublic {
		publicStr = "true"
	}
	if err := writer.WriteField("public", publicStr); err != nil {
		return nil, fmt.Errorf("failed to write public field: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	// Create request
	url := c.baseURL + "/upload"
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+c.token)

	// Send request
	client := &http.Client{Timeout: c.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Parse response
	var uploadResp CDNUploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if !uploadResp.Success {
		return nil, errors.New(uploadResp.Error)
	}

	return &uploadResp, nil
}

// DeleteFile deletes a file from CDN server
func (c *CDNClient) DeleteFile(tag, filename string) error {
	url := fmt.Sprintf("%s/api/files/%s/%s", c.baseURL, tag, filename)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	client := &http.Client{Timeout: c.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var deleteResp CDNDeleteResponse
	if err := json.NewDecoder(resp.Body).Decode(&deleteResp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if !deleteResp.Success {
		return errors.New(deleteResp.Error)
	}

	return nil
}

// GetFileURL returns the full URL for a file
func (c *CDNClient) GetFileURL(tag, filename string) string {
	return fmt.Sprintf("%s/%s/%s", c.baseURL, tag, filename)
}

// ExtractFilenameFromURL extracts filename from CDN URL
func ExtractFilenameFromURL(url string) string {
	return filepath.Base(url)
}

// ExtractTagFromURL extracts tag from CDN URL
func ExtractTagFromURL(url string) string {
	parts := filepath.Dir(url)
	return filepath.Base(parts)
}
