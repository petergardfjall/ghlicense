package main

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

func main() {
	Execute()
}

func prettyJSON(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.Fatal().Err(err).Msg("marshal json")
	}
	return string(b)
}
