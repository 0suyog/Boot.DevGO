package main

import (
	"strings"
)

func _removeProfanity(message *string) {
	profaneWords:=map[string]string{"fubb":"****","shiz":"****","witch":"*****"}
	for key,word := range profaneWords{
		*message=strings.ReplaceAll(*message,key,word)
	}
}