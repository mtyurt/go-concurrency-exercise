package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Examples taken from http://whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/

func doAction(name string, action string, durBase int, durOffset int, c chan int) {
	fmt.Println(name + " started " + action)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dur := durBase + r.Intn(durOffset)
	time.Sleep(time.Duration(dur*10) * time.Millisecond)
	fmt.Println(name + " spent " + strconv.Itoa(dur) + " seconds " + action)
	c <- dur
}

func getReady(name string, c chan int) {
	doAction(name, "getting ready", 60, 30, c)
}

func putOnShoes(name string, c chan int) {
	doAction(name, "putting on shoes", 35, 10, c)
}

func armAlarm(c chan int) {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Alarm is counting down.")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Alarm is armed.")
	c <- 1
}

func main() {
	fmt.Println("Let's go for a walk!")
	c := make(chan int)
	go getReady("Bob", c)
	go getReady("Alice", c)
	_, _ = <-c, <-c
	alarmChannel := make(chan int)
	fmt.Println("Arming alarm.")
	go armAlarm(alarmChannel)
	go putOnShoes("Bob", c)
	go putOnShoes("Alice", c)
	_, _ = <-c, <-c
	fmt.Println("Exiting and locking the door.")
	_ = <-alarmChannel
}
