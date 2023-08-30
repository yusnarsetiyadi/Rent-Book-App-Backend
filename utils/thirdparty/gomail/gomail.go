package gomail

import (
	"fmt"
	"rentbook/internal/config"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func Gomail(TittleFunc, Subject, Receiver, Message string) (string, error) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.GetConfig().SENDER_NAME)
	mailer.SetHeader("To", Receiver)
	mailer.SetHeader("Subject", Subject)
	mailer.SetBody("text/html", Message)

	dialer := gomail.NewDialer(
		config.GetConfig().SMTP_HOST,
		int(config.GetConfig().SMTP_PORT),
		config.GetConfig().AUTH_EMAIL,
		config.GetConfig().AUTH_PASSWORD,
	)

	returnMsgError := fmt.Sprintf("Error %s", TittleFunc)
	returnMsgSuccess := fmt.Sprintf("Success %s", TittleFunc)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		logrus.Errorf("Error %s", TittleFunc)
		return returnMsgError, err
	}

	logrus.Infof("Success %s", TittleFunc)
	return returnMsgSuccess, nil
}

func SendEmailLoginInfo(Receiver string, Subject string, email string, password string, username string) (string, error) {
	message := fmt.Sprintf("Hello %s, you can login to the app with Email: %s and Password: %s", username, email, password)
	msg, err := Gomail("Send Email Login Info", Subject, Receiver, message)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func SendEmailDeactiveUser(Receiver string, Subject string, username string) (string, error) {
	message := fmt.Sprintf("Hello %s, your account has been deactive.", username)
	msg, err := Gomail("Send Email Deactive User", Subject, Receiver, message)
	if err != nil {
		return msg, err
	}
	return msg, nil
}
