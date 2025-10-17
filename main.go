package main

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)

	loadRecipient("./emails.csv", recipientChannel)
}
