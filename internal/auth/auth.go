package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string, error) {
	apiKey := header.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("authorization header is required")
	}

	splitAuth := strings.Split(apiKey, " ")
	if len(splitAuth) != 2 {
		return "", errors.New("invalid authorization header")
	}

	return splitAuth[1], nil
}
