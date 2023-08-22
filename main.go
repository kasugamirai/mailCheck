package mian

import (
	"fmt"
	"github.com/kasugamirai/mailCheck/handlers"
	"log"
	"os"
)

func main() {
	email := os.Getenv("your_email")
	password := os.Getenv("your_password")
	c, err := handlers.ConnectToEmail(email, password)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()
	content := "123"
	foundEmails, err := handlers.CheckForContents(c, content)
	if err != nil {
		log.Fatal(err)
	}

	if len(foundEmails) == 0 {
		fmt.Println("没有包含" + content + "的邮件")
	} else {
		for _, subject := range foundEmails {
			fmt.Println("找到包含"+content+"的邮件:", subject)
		}
	}
}
