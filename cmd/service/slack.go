package service

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Slack struct {
}

func NewSlack() *Slack {
	return &Slack{}
}

func (s Slack) Send(ctx context.Context, subject, message string) error {
	data := map[string]string{
		"text": "```\n" + message + "\n```",
	}
	j, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	url := "https://hooks.slack.com/services/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"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return nil
}
