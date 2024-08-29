package mail

import (
	"fmt"
	"n9/config"

	"gopkg.in/mail.v2"
)

type Mailer struct{
	Cfg config.Config
}


func NewMailer(cfg config.Config)*Mailer{
	return &Mailer{Cfg: cfg}
}

func (ml *Mailer)Mail(msg string, to string) {
	m := mail.NewMessage()

	m.SetHeader("From", ml.Cfg.Email_Sender)
	m.SetHeader("To", to)    
	m.SetHeader("Subject", "Hello from Hotel!")
	m.SetBody("text/plain", msg)

	d := mail.NewDialer("smtp.gmail.com", 587, ml.Cfg.Email_Sender, ml.Cfg.Email_Password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully!")
}


