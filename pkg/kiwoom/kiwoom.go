package kiwoom

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/net/websocket"
	"resty.dev/v3"
)

type Client struct {
	AppKey        string
	AppSecret     string
	RestBaseUrl   string
	WsBaseUrl     string
	AccessToken   string
	TokenTTL      time.Time
	RestClient    *resty.Client
	WebSocketConn *websocket.Conn
}

func ensureAbsolute(path string) (string, error) {
	if filepath.IsAbs(path) {
		return filepath.Clean(path), nil
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	return absPath, nil
}

func NewClient(envPath string) (*Client, error) {
	absPath, err := ensureAbsolute(envPath)
	if err != nil {
		return nil, err
	}
	if err := godotenv.Load(absPath); err != nil {
		return nil, err
	}

	appKey := os.Getenv("APP_KEY")
	appSecret := os.Getenv("APP_SECRET")

	return &Client{
		AppKey:      appKey,
		AppSecret:   appSecret,
		RestBaseUrl: DefaultRestBaseUrl,
		WsBaseUrl:   DefaultWsBaseUrl,
		RestClient:  resty.New().SetBaseURL(DefaultRestBaseUrl),
	}, nil
}
