package logtail

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Logtail struct {
	url   *url.URL
	token string
}

var (
	client = &http.Client{}
)

func NewLogtail(token string) *Logtail {
	return &Logtail{
		url: &url.URL{
			Scheme: "https",
			Host:   "in.logtail.com",
		},
		token: "Bearer " + token,
	}
}

func (l *Logtail) Send(body string) (int, error) {
	request, _ := http.NewRequest(http.MethodPost, "", strings.NewReader(body))
	request.URL = l.url
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", l.token)

	response, err := client.Do(request)
	if err != nil {
		if response != nil {
			response.Body.Close()
		}
		return 0, fmt.Errorf("log send: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 202 {
		return response.StatusCode, fmt.Errorf("Error sending log: %w", err)
	}
	return response.StatusCode, nil
}
