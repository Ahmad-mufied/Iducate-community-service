package utils

import (
	"fmt"
	"io"
	"net/http"
)

func RequestGET(url string, headers map[string]string) ([]byte, error) {

	client := new(http.Client)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get data, status code: %d", response.StatusCode)
	}

	return body, nil
}
