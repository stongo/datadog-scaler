package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type ScalerData struct {
	EventTitle  string
	AlertStatus string
}

type Scaler struct {
	Data *ScalerData
}

func NewScaler(data *ScalerData) *Scaler {
	return &Scaler{Data: data}
}

func (s *Scaler) Scale() (string, error) {
	if strings.Contains(s.Data.EventTitle, "[Scale Up]") {
		out, commandErr := exec.Command("echo", "hello").Output()
		if commandErr != nil {
			log.Print(commandErr)
			return "", fmt.Errorf("Command response: %s\n", commandErr)
		}
		return fmt.Sprintf("Command response: %s\n", out), nil
	} else if strings.Contains(s.Data.EventTitle, "[Scale Down]") {
		out, commandErr := exec.Command("echo", "bye").Output()
		if commandErr != nil {
			log.Print(commandErr)
			return "", fmt.Errorf("Command response: %s\n", commandErr)
		}
		return fmt.Sprintf("Command response: %s\n", out), nil
	}
	return "no command executed", nil
}
