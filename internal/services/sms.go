package services

import (
	"bytes"
	"net/http"
	"net/url"
)

type sms struct {
	Username    string
	PhoneNumber string
	ShortCode   string
	Message     string
	APIKey      string
}

func sendSMS(sm *sms) error {
	urlx := "https://api.sandbox.africastalking.com/version1/messaging"

	data := url.Values{}
	data.Set("username", sm.Username)
	data.Set("to", sm.PhoneNumber)
	data.Set("message", sm.ShortCode)
	data.Set("from", sm.ShortCode)

	req, err := http.NewRequest("POST", urlx, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apiKey", sm.APIKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err

	}
	defer resp.Body.Close()
	return nil
}
