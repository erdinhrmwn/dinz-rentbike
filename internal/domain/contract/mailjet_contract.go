package contract

type EmailRequest struct {
	ToEmail  string
	ToName   string
	Subject  string
	HTMLBody string
}

type MailjetService interface {
	Send(req *EmailRequest) error
}
