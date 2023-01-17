package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/nikoksr/notify"
)

type Logger struct {
	services []string
}

type LocalLog struct {
	ServiceName string
	Environment string
	Date        string
	Message     string
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) SetLogger(services ...string) {
	if len(services) == 0 {
		log.Fatal("A service must be set in the logger")
	}
	l.services = services
}

func (l *Logger) Log(message string) {
	services := map[string]notify.Notifier{
		"email": mailService(),
		"slack": slackService(),
	}
	var wg sync.WaitGroup
	for _, v := range l.services {
		if _, ok := services[v]; !ok {
			continue
		}
		wg.Add(1)
		v := v
		go (func() {
			defer wg.Done()
			notify.UseServices(services[v])
			err := notify.Send(
				context.Background(),
				"Demo App:: Attention required",
				message,
			)
			if err != nil {
				log.Fatal(err.Error())
			}
			now := time.Now()
			logToFile(&LocalLog{
				ServiceName: v,
				Date:        now.Format(time.RFC3339),
				Environment: "production",
				Message:     message,
			})
		})()
	}
	wg.Wait()
}

func logToFile(payload *LocalLog) {
	f, err := os.OpenFile("local.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	res := fmt.Sprintf("[%v]::[%s]::[%s]:: %s", payload.ServiceName, payload.Environment, payload.Date, payload.Message)
	_, err = f.WriteString(res + "\n")
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
}

func slackService() *Slack {
	slackService := NewSlack()
	return slackService
}

func mailService() *Mail {
	email := NewMailService("info@gmail.ccom", "smtp.mailtrap.io:2525")
	email.AddReceivers("iamhabbeboy@gmail.com")
	email.AuthenticateSMTP("", "3e077c0853c98d", "617119660145dc", "smtp.mailtrap.io")
	return email
}
