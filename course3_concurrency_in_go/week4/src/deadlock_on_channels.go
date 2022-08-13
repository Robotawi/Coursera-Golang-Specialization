package main

func worker1(ch1 chan int, ch2 chan int) int {
	// wait to get something from ch1
	rcv := <-ch1
	// then send on ch2
	ch2 <- 1

	return rcv
}

func worker2(ch1 chan int, ch2 chan int) int {
	// wait to get something from ch2
	rcv := <-ch2
	// then send on ch2
	ch1 <- 1

	return rcv
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	worker1(ch1, ch2)

}
