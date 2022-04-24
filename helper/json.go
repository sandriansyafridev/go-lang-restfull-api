package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(r *http.Request, request interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(request)
	if err != nil {
		return err
	}

	return nil
}

func WriteRequestBody(w http.ResponseWriter, response interface{}) {

	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)

}
