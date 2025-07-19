package main

import (
	"time"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)
go	sendIsOld(isOldChan, emails)
	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

 func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}


func main() {
	// Sample time values
	oldDate := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	newDate := time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC)

	// First call: all old
	result1 := checkEmailAge([3]email{
		{body: "old1", date: oldDate},
		{body: "old2", date: oldDate},
		{body: "old3", date: oldDate},
	})
	println(result1[0], result1[1], result1[2])

	// Second call: all new
	result2 := checkEmailAge([3]email{
		{body: "new1", date: newDate},
		{body: "new2", date: newDate},
		{body: "new3", date: newDate},
	})
	println(result2[0], result2[1], result2[2])

	// Third call: mix
	result3 := checkEmailAge([3]email{
		{body: "mix1", date: oldDate},
		{body: "mix2", date: newDate},
		{body: "mix3", date: oldDate},
	})
	println(result3[0], result3[1], result3[2])
}
