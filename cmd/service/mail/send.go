package service

import "context"

type MailSender struct {
}

func Send() {
	email := New("info@gmail.ccom", "smtp.mailtrap.io:2525")
	email.AddReceivers("iamhabbeboy@gmail.com")
	email.AuthenticateSMTP("", "3e077c0853c98d", "617119660145dc", "smtp.mailtrap.io")
	email.Send(context.Background(), "Hello notifier", "Notifier body message")
}
