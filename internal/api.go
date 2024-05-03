package internal

import (
	"github.com/sleeyax/smspool-go/internal/utils"
	"io"
	"net/http"
)

type Api struct {
	// The required API key to access the API.
	apiKey string

	// The HTTP client to use for requests.
	client *http.Client
}

// Create creates a new API instance.
func Create(apiKey string) Api {
	return Api{
		apiKey: apiKey,
		client: &http.Client{},
	}
}

func (a Api) setHeaders(req *http.Request, contentType string) {
	req.Header.Set("Authorization", "Bearer "+a.apiKey)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", "smspool-go v0")
}

// Send sends an HTTP POST request to the API.
func (a Api) Send(request interface{}, path string) ([]byte, error) {
	formData, contentType := utils.ToFormData(request, a.apiKey)
	req, err := http.NewRequest(http.MethodPost, BaseUrl+path, formData)
	if err != nil {
		return nil, err
	}
	a.setHeaders(req, contentType)

	res, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
