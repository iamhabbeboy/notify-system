package service

type SlackService struct{}

func NewSlackService() *SlackService {
	return &SlackService{}
}

func (*SlackService) Send(payload interface{}) {

}
