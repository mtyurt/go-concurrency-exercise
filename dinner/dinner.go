package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Examples taken from http://whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/

type dish struct {
	Name  string
	Count int
}

func startEating(name string, dishController chan string) {
	for {
		sleepDur := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1500) + 300
		time.Sleep(time.Duration(sleepDur) * time.Millisecond)
		dishController <- name
	}
}
func dishController(allDishes []dish, controllerChannel chan string) {
	i := 0
	dishes := allDishes
	fmt.Println("dishes:", dishes)
	for {
		select {
		case name := <-controllerChannel:
			{
				i++
				length := len(dishes)
				randomDish := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(length)
				selectedDish := dishes[randomDish]
				fmt.Println(i, name, "is enjoying some", selectedDish.Name)
				dishes[randomDish].Count--
				if selectedDish.Count == 0 {
					if length == 1 {
						return
					} else {
						dishes = append(dishes[:randomDish], dishes[randomDish+1:]...)
					}
				}
			}
		}
	}
}
func generateRandomMorselCount() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(6) + 5
}

func main() {
	fmt.Println("Bon appetite!")
	dishes := make([]dish, 5)
	dishes[0] = dish{Name: "chorizo", Count: generateRandomMorselCount()}
	dishes[1] = dish{Name: "chopitos", Count: generateRandomMorselCount()}
	dishes[2] = dish{Name: "croquetas", Count: generateRandomMorselCount()}
	dishes[3] = dish{Name: "patatas bravas", Count: generateRandomMorselCount()}
	dishes[4] = dish{Name: "pimientos de padron", Count: generateRandomMorselCount()}

	controllerChannel := make(chan string)
	go startEating("Alice", controllerChannel)
	go startEating("Bob", controllerChannel)
	go startEating("Charlie", controllerChannel)
	go startEating("Dave", controllerChannel)

	dishController(dishes, controllerChannel)
}
