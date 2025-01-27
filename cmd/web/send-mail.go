package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMSG(msg)
		}
	}()
}

func sendMSG(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	if m.Template == "" {
		email.SetBody(mail.TextHTML, m.Content)
	} else {
		data, err := os.ReadFile(fmt.Sprintf("./static/email/templates/%s"))
		if err != nil {
			app.ErrorLog.Println(err)
		}
		mailTemplate := string(data)
		msgToSend := strings.Replace(mailTemplate, "[%E-MAIL-CONTENT%]", m.Content, 1)
		email.SetBody(mail.TextHTML, msgToSend)

	}

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("E-mail sent out!")
	}
}
