package main

import (
	"log"
	"net/smtp"
)

func main() {
	timu := ""
	faaddress := "zgslzgsl@163.com"
	sender := "zgslzgsl@163.com"
	shouaddress := "zgslzgsl@gmail.com"
	server := "mail.163.com:25"
	mima := "123456"
	ipaddress := "mail.163.com"
	to := []string{shouaddress}
	msg := []byte("To:recipient@example.net\r\n" + "Subject:discount Gophers!\r\n" + "This is the email body.\r\n")
	auth := smtp.PlainAuth(timu, faaddress, mima, ipaddress)
	if err := smtp.SendMail(server, auth, sender, to, msg); err != nil {
		log.Fatal(err)
	}
}
