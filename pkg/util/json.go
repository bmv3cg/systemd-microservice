package util

import (
	"encoding/json"
	"net/http"
)

// JSONResponse form a JSON respose for given struct
func JSONResponse(response interface{}, w http.ResponseWriter) error {
	//json, err := json.Marshal(response)
	json, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	return nil
}
