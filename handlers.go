package main

import (
	"encoding/json"
	"fmt"
	"github.com/abbot/go-http-auth"
	"log"
	"net/http"
)

func ScalerPostHandler(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	switch r.Method {
	case http.MethodPost:
		var sd ScalerData
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
		}
		err := json.NewDecoder(r.Body).Decode(&sd)
		if err != nil {
			log.Print(err)
			http.Error(w, fmt.Sprintf("invalid request: %s", err), 400)
		}
		scaler := NewScaler(&sd)
		output, error := scaler.Scale()
		if error != nil {
			http.Error(w, fmt.Sprintf("internal error: %s", error), 500)
		}
		fmt.Fprint(w, output)
	}
}
