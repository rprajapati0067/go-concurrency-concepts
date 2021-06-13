package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var initialString string
var finalString string
var stringLength int

// var loremIpSum string
// var finalIpSum string
// var letterSentChan chan string

// func deliverToFinal(letter string, finalIpSum *string) {
// 	*finalIpSum += letter
// }

// func capitalize(current *int, length int, letters []byte, finalIpSum *string) {

// 	for *current < length {
// 		thisLetter := strings.ToUpper(string(letters[*current]))

// 		deliverToFinal(thisLetter, finalIpSum)
// 		*current++
// 	}

// }

func addToFinalstack(letterChannel chan string, wg *sync.WaitGroup) {
	letter := <-letterChannel
	finalString += letter
	wg.Done()

}
func capitalize(letterChannel chan string, currentLetter string, wg *sync.WaitGroup) {
	thisLetter := strings.ToUpper(currentLetter)
	wg.Done()
	letterChannel <- thisLetter
}

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	//index := new(int)

	// *index = 0
	// loremIpSum = "The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make’s return type is the same as the type of its argument, not a pointer to it."
	// letters := []byte(loremIpSum)
	// length := len(letters)

	// go capitalize(index, length, letters, &finalIpSum)
	// go func() {
	// 	go capitalize(index, length, letters, &finalIpSum)
	// }()

	initialString := "The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make’s return type is the same as the type of its argument, not a pointer to it."
	initialBytes := []byte(initialString)
	var letterChannel chan string = make(chan string)

	stringLength = len(initialBytes)
	for i := 0; i < stringLength; i++ {
		wg.Add(2)

		go capitalize(letterChannel, string(initialBytes[i]), &wg)
		go addToFinalstack(letterChannel, &wg)

		wg.Wait()
	}
	fmt.Println(finalString)
}
