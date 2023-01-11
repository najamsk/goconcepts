package mailsender

type MailSender interface {
	Send(to, subject, body string) error
	SendFrom(from, to, subject, body string) error
}

type MockSender struct {
	SendFunc     func(to, subject, body string) error
	SendFromFunc func(from, to, subject, body string) error
}

func (m MockSender) Send(to, subject, body string) error {
	return m.SendFunc(to, subject, body)
}
func (m MockSender) SendFrom(from, to, subject, body string) error {
	return m.SendFromFunc(from, to, subject, body)
}
