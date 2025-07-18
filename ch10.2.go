package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	profaneWords:=map[string]string{"fubb":"****","shiz":"****","witch":"*****"}
	for key,word := range profaneWords{
		*message=strings.ReplaceAll(*message,key,word)
	}
}

func main(){
testCase:=[]string{
			"English, motherfubber, do you speak it?",
			"Oh man I've seen some crazy ass shiz in my time...",
	}

removeProfanity(&testCase[0])
fmt.Println(testCase[0])
}
