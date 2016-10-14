package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/mtyurt/internetcafe/queue"
)

// Examples taken from http://whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/

func tourist(number int, leavingChannel chan bool) {
	fmt.Println("Tourist", strconv.Itoa(number), "is online.")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dur := 15 + r.Intn(105)
	time.Sleep(time.Duration(dur*10) * time.Millisecond)
	fmt.Println("Tourist", number, "is done, having spent", dur, "minutes online.")
	leavingChannel <- true
}

func manager(incomingTourists chan int, completed chan bool) {
	waitingTourists := queue.CreateQueue()
	activeTouristCount := 0
	touristLeaving := make(chan bool)
	for {
		select {
		case incoming := <-incomingTourists:
			{
				if activeTouristCount < 8 {
					go tourist(incoming, touristLeaving)
					activeTouristCount = activeTouristCount + 1
				} else {
					fmt.Println("Tourist", incoming, "waiting for turn.")
					waitingTourists.Push(incoming)
				}
			}
		case _ = <-touristLeaving:
			{
				if waitingTourists.Len() > 0 {
					newTourist, _ := waitingTourists.Pop()
					go tourist(newTourist, touristLeaving)
				} else {
					activeTouristCount = activeTouristCount - 1
					if activeTouristCount == 0 {
						completed <- true
					}
				}
			}
		default:
		}
	}
}

func main() {
	completed := make(chan bool)
	incomingTourists := make(chan int)
	go manager(incomingTourists, completed)
	for i := 1; i < 26; i++ {
		incomingTourists <- i
	}
	_ = <-completed
}
