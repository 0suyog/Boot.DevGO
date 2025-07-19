package main

func addEmailsToQueue(emails []string) chan string {
	emailChan:=make(chan string,len(emails) )
	for _,email := range emails{
		emailChan<-email
	} 
	return emailChan
}
