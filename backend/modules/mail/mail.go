package mail

import (
	"net/smtp"
	"strings"
)



type Mail struct {
    From     string
    Username string
    Password string
    To       string `json:"email"`
    Sub      string
    Message	 string
}

const (
	PORT = "587"
	PROTOCOL = "smtp"
)

func SendEmail(m Mail) error {
	index := strings.Index(m.From,"@")
	host := m.From[index+1:]
    smtpSvr := PROTOCOL + "." + host + ":" + PORT
    auth := smtp.PlainAuth("", m.Username, m.Password, PROTOCOL + "." + host)
    if err := smtp.SendMail(smtpSvr, auth, m.From, []string{m.To}, []byte(m.body())); err != nil {
        return err
    }
    return nil
}

func (m Mail) body() string {
    return "To: " + m.To + "\r\n" +
        "Subject: " + m.Sub + "\r\n\r\n" +
        m.Message + "\r\n"
}
