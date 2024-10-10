package service

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (s *EmailService) Send() {

}
