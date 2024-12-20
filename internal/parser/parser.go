package parser

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const url = "https://track.4px.com/queryBatch/%s?locale=en_US"

func ParsePage(trackNumbers []string) (string, error) {
	if len(trackNumbers) == 0 {
		return "", errors.New("no track numbers provided")
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(url, strings.Join(trackNumbers, ",")), nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error executing request: %w", err)
	}

	if response.StatusCode != 200 {
		return "", fmt.Errorf("error executing request: %d %s", response.StatusCode, response.Status)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}
