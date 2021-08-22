package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"os"
	"time"
	)

func main() {
    
	magicNo := generateRandRange()
	//fmt.Printf("Magic No is... %d", magicNo)
	
	for i := 10; i > 0; i-- {
		
		fmt.Printf("\nYou have %d attempts left", i)
		
		var guessStr string
		fmt.Println("\nGuess the number from 1 to 100: ")
		fmt.Scanln(&guessStr)
		
		guessNum, err := strconv.Atoi(guessStr)
		if err != nil {
			fmt.Println(err)
		}
		
		isGuessCorrect(guessNum, magicNo)				
	}
	
	fmt.Printf("\nSorry, You dint guess my number. It was %d", magicNo)
}

func isGuessCorrect(guessNum int,magicNo int) {
	if guessNum > magicNo {
		fmt.Println("Oops, Your guess was HIGH.")
	} else if guessNum < magicNo {
		fmt.Println("Oops, Your guess was LOW.")
	} else {
		fmt.Println("Good Job! You guessed it!")
		os.Exit(2)
	}
}

func generateRandRange() int {
	rand.Seed(time.Now().UnixNano())
    min := 1
    max := 100
	return rand.Intn(max - min + 1) + min
}