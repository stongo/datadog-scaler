package main

import (
	"fmt"
	"github.com/abbot/go-http-auth"
	"github.com/containous/flaeg"
	"github.com/containous/staert"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	LogLevel string               `description:"Log level - ERROR, WARN, INFO, DEBUG"`
	Listen   *ListenConfiguration `description:"A pointer field"`
}

type ListenConfiguration struct {
	Port int    `description:"Port to listen on"`
	IP   string `description:"IP to bind to"`
}

func HttpBasicAuthSecret(user, realm string) string {
	if user == "circleci" {
		// password is "test"
		return "$apr1$QvVS4zmg$ZN0ZJVC1.a0/MkwZcmnwF0"
	}
	return ""
}

func main() {
	listenDefault := &ListenConfiguration{
		Port: 8000,
		IP:   "127.0.0.1",
	}

	config := &Configuration{
		LogLevel: "WARN",
		Listen:   listenDefault,
	}
	//Set default pointers value
	defaultPointersConfig := &Configuration{
		Listen: listenDefault,
	}

	command := &flaeg.Command{
		Name:                  "CircleCI Auto-scaler",
		Description:           "A webhook that runs circle-admin-cli reactively from datadog alerts",
		Config:                config,
		DefaultPointersConfig: defaultPointersConfig,
		Run: func() error {

			mux := http.NewServeMux()
			authenticator := auth.NewBasicAuthenticator("127.0.0.1", HttpBasicAuthSecret)
			mux.HandleFunc("/", authenticator.Wrap(ScalerPostHandler))
			listen := fmt.Sprintf("%s:%d", config.Listen.IP, config.Listen.Port)
			fmt.Printf("listen: %v", listen)
			http.ListenAndServe(listen, mux)
			return nil
		},
	}

	s := staert.NewStaert(command)
	f := flaeg.New(command, os.Args[1:])
	s.AddSource(f)

	_, err := s.LoadConfig()
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	if err := command.Run(); err != nil {
		log.Fatalf("Error %s", err.Error())
	}
}
