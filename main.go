package main

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)

	go func() {

		loadRecipient("./emails.csv", recipientChannel)
	}()

	go emailWorker(1, recipientChannel)

}
