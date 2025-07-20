package main

func countReports(numSentCh chan int) int {
	for total:=0;;{
		if n,ok:=<- numSentCh;!ok{
			return  total
		}else{
		total+=n
		}
	}

}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
