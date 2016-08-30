package main

import (
	"fmt"
	"github.com/abbot/go-http-auth"
	"log"
	"net/http"
	"os/exec"
)

func ScalerPostHandler(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	switch r.Method {
	case http.MethodPost:
		out, err := exec.Command("echo", "hello").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Command response: %s\n", out)
	}
}
