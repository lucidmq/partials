package services

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"time"
)

func SendEmail(website string, formBody map[string]string) {
	var sb strings.Builder
	for k, v := range formBody {
		s := fmt.Sprintf("%s: %s\n", k, v)
		sb.WriteString(s)
	}
	s := fmt.Sprintf("Website: %s\n", website)
	sb.WriteString(s)
	currentTime := time.Now()
	subject := fmt.Sprintf("New Website Lead: %s", currentTime.Format("2006-01-02 03:04 pm"))
	sendEmail(subject, sb.String())
}

func sendEmail(subject string, body string) error {
	// Sender data.
	from := "leads.integrand@gmail.com"
	password := "ztmv lroc ggxp muqs"
	to := "burak@bkconsulting.io"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Receiver email address.
	to_list := []string{
		to,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(msg)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to_list, message)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Email Sent Successfully!")
	return nil
}
