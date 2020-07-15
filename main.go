package main

import (
	"fmt"

	"email_planner/email"
)

func main() {
	fmt.Printf("Hello, world!\n")
	content := "Subject: Subject!\r\n" +
	"\r\n" +
	"This is the email body.\r\n"
	email.Connect_smtp(content)
}