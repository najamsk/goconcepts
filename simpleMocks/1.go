package main

import (
	"fmt"
	"log"
	mail "mocking/pkgs/mailsender"
)

func main() {
	fmt.Println("hi")

	ms := mail.MockSender{
		SendFunc: func(to, subject, body string) error {
			log.Printf("Mocked send func -> subject: %s \n", subject)
			return nil
		},
		SendFromFunc: func(from, to, subject, body string) error {
			log.Printf("Mocked sendFrom func -> subject: %s \n", subject)
			return nil
		},
	}
	ms.Send("najamsk@gmail.com", "paypal payment", "hey friend I got your money")
	ms.SendFrom("prince@nigeria.com", "najamsk@gmail.com", "100 million Gold bars", "I have money and gold that need your bank account. I take only 2 percent.")
}
