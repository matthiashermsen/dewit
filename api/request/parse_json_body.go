package request

import (
	"encoding/json"
	"net/http"
)

func ParseJSONBody(r *http.Request, requestBody any) error {
	requestBodyDecoder := json.NewDecoder(r.Body)
	err := requestBodyDecoder.Decode(&requestBody)

	return err
}
