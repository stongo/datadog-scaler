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

func (s *Scaler) ScaleUpPrecise() (string, error) {
	out, err := exec.Command("echo", "scale up precise").Output()
	if err != nil {
		log.Print(err)
		return "", fmt.Errorf("Command response: %s\n", err)
	}
	return fmt.Sprintf("Command response: %s\n", out), nil
}

func (s *Scaler) ScaleUpTrusty() (string, error) {
	out, err := exec.Command("echo", "scale up trusty").Output()
	if err != nil {
		log.Print(err)
		return "", fmt.Errorf("Command response: %s\n", err)
	}
	return fmt.Sprintf("Command response: %s\n", out), nil
}

func (s *Scaler) ScaleDownPrecise() (string, error) {
	out, err := exec.Command("echo", "scale down precise").Output()
	if err != nil {
		log.Print(err)
		return "", fmt.Errorf("Command response: %s\n", err)
	}
	return fmt.Sprintf("Command response: %s\n", out), nil
}

func (s *Scaler) ScaleDownTrusty() (string, error) {
	out, err := exec.Command("echo", "scale down trusty").Output()
	if err != nil {
		log.Print(err)
		return "", fmt.Errorf("Command response: %s\n", err)
	}
	return fmt.Sprintf("Command response: %s\n", out), nil
}

func (s *Scaler) Scale() (string, error) {
	if strings.Contains(s.Data.EventTitle, "[Scale Up Precise]") {
		return s.ScaleUpPrecise()
	} else if strings.Contains(s.Data.EventTitle, "[Scale Up Trusty]") {
		return s.ScaleUpTrusty()
	} else if strings.Contains(s.Data.EventTitle, "[Scale Down Precise]") {
		return s.ScaleDownPrecise()
	} else if strings.Contains(s.Data.EventTitle, "[Scale Down Trusty]") {
		return s.ScaleDownTrusty()
	}
	return "no command executed", nil
}
