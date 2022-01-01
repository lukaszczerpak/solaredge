package solaredgeapi

import (
	"github.com/go-resty/resty/v2"
	"solaredge-sync/util"
	"time"
)

const (
	DATETIME_FORMAT = "2006-01-02 15:04:05"
)

type Client struct {
	ApiClient *resty.Client
}

func createRestyClient(apiKey string) *resty.Client {
	client := resty.New()
	//client.SetDebug(true)
	client.SetTimeout(1 * time.Minute)
	client.SetBaseURL("https://monitoringapi.solaredge.com")
	client.SetHeader("Accept", "application/json")
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   util.USER_AGENT,
	})
	client.SetQueryParam("api_key", apiKey)

	return client
}

func New(apiKey string) *Client {
	client := createRestyClient(apiKey)
	return &Client{client}
}
