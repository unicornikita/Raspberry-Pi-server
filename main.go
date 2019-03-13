package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"
)

func main() {
	http.HandleFunc("/openedDoor", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("processing")
		now := time.Now()
		hour := now.Hour()
		fmt.Print(hour)
		if hour >= 21 || hour <= 5 { // the time when I don't want the door to open
			sendmail("message")  // change message with the text you want to recieve
			fmt.Println("works") // Terminal prints "works" and lets me know the script is running smoothly
		}
	})
	http.ListenAndServe(":80", nil)
}

func sendmail(body string) {
	from := "name@gmail.com" // replace name@gmail.com with the gmail address you want to recieve messages from
	pass := "pass"           // password of the said mail address
	to := "name@gmail.com"   // reciever's mail address

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Alert From Magnetic Sensor" + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent") // terminal lets me know the message has been sent
}
