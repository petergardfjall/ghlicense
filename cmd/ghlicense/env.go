package main

import (
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	AccessTokenEnv = "GITHUB_ACCESS_TOKEN"
)

var (
	accessToken = ""
	timeout     = 30 * time.Second
	verbose     = false
)

func mustGetAccessToken() string {
	if accessToken != "" {
		return accessToken
	}

	envToken := os.Getenv(AccessTokenEnv)
	if envToken == "" {
		log.Fatal().Msg("No GitHub access token supplied: use --access-token or GITHUB_ACCESS_TOKEN")
	}
	return envToken
}
