package service

import "log"

type LoggerService struct {
	service string
}

type externalServiceManager interface {
	Send(payload interface{})
}

func NewLoggerService() *LoggerService {
	return &LoggerService{}
}

func (*LoggerService) SetLogger(services ...string) string {
	if len(services) == 0 {
		log.Fatal("A service must be set in the logger")
	}

	svcs := map[string]externalServiceManager{
		"log":   NewLogService(),
		"slack": NewSlackService(),
	}

	for _, service := range services {
		if svc, ok := svcs[service]; ok {
			svc.Send("sdf")
		}
	}
	return ""
}

func (*LoggerService) Log(payload string) string {
	return "Hello world"
	// telegramService, err := telegram.New("5351562309:AAG2G2YTA1JoL-nBxRWjR2ckaQU7myLDaYk")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// telegramService.AddReceivers(-1234567890)

	// // // Send a test message.
	// notify.UseServices(email, telegramService)

	// err = notify.Send(
	// 	context.Background(),
	// 	"Subject/Title",
	// 	"The actual message - Hello, you awesome gophers! :)",
	// )
}
