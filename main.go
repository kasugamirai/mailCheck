package mian

import (
	"fmt"
	"github.com/kasugamirai/mailCheck/handlers"
	"log"
)

func main() {
	c, err := handlers.ConnectToEmail("your_email@gmail.com", "your_password")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	foundEmails, err := handlers.CheckForContents(c)
	if err != nil {
		log.Fatal(err)
	}

	if len(foundEmails) == 0 {
		fmt.Println("没有包含 '123' 的邮件")
	} else {
		for _, subject := range foundEmails {
			fmt.Println("找到包含 '123' 的邮件:", subject)
		}
	}
}
