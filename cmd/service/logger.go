package service

import (
	"context"
	"log"
	service "main/cmd/service/mail"
	"sync"

	"github.com/nikoksr/notify"
)

type LoggerService struct {
	services []string
}

func NewLoggerService() *LoggerService {
	return &LoggerService{}
}

func (l *LoggerService) SetLogger(services ...string) {
	if len(services) == 0 {
		log.Fatal("A service must be set in the logger")
	}
	l.services = services
}

func (l *LoggerService) Log(message string) {
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
		})()
	}
	wg.Wait()
}

func slackService() *Slack {
	slackService := NewSlack()
	return slackService
}

func mailService() *service.Mail {
	email := service.NewMailService("info@gmail.ccom", "smtp.mailtrap.io:2525")
	email.AddReceivers("iamhabbeboy@gmail.com")
	email.AuthenticateSMTP("", "3e077c0853c98d", "617119660145dc", "smtp.mailtrap.io")
	return email
}
