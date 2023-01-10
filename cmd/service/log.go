package service

type LogService struct{}

func NewLogService() *LogService {
	return &LogService{}
}

func (*LogService) Send(payload interface{}) {

}
