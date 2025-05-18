package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

func SendEmailCode(to string, code int, smtpHost, smtpPort, email, password string) error {
	// Trim any trailing periods from the email address
	to = strings.TrimRight(to, ".")

	// Log the email configuration (excluding password)
	log.Printf("Attempting to send email to: %s", to)
	log.Printf("SMTP Host: %s, Port: %s", smtpHost, smtpPort)
	log.Printf("From email: %s", email)

	body := fmt.Sprintf("To: %s\r\nSubject: Confirmation Code\r\n\r\nYour confirmation code is %d", to, code)
	auth := smtp.PlainAuth("", email, password, smtpHost)

	// Construct the full SMTP address
	smtpAddr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)
	log.Printf("Connecting to SMTP server: %s", smtpAddr)

	err := smtp.SendMail(smtpAddr, auth, email, []string{to}, []byte(body))
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return fmt.Errorf("failed to send email: %v", err)
	}

	log.Printf("Email sent successfully to %s", to)
	return nil
}
