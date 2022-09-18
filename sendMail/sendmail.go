package sendMail

import (
	"log"
	"net/smtp"

	"github.com/cryptix/smtpAuthLogin"
)

func Sendto(mail string, UUID string) {

	from := "raylin366@hotmail.com"
	password := "05ray08ray"

	to := []string{
		mail,
	}

	smtpHost := "smtp-mail.outlook.com"
	smtpPort := "587"

	message := []byte("From: raylin366@hotmail.com\r\n" +
		"To: " + mail + "\r\n" +
		"Subject: UUID number\r\n\r\n" +
		UUID + "\r\n")

	auth := smtpAuthLogin.LoginAuth(from, password) //smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println("error sending mail -> ", err)
		log.Println("Email couldnt be sent maybe a typo or no internet connection")
		return
	}
	log.Println("Email sent Succesfully")
}
