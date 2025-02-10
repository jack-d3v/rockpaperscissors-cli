package main 

import (
	"fmt"
	"strings"
	"math/rand"
)

var userScore int = 0
var compScore int = 0
var gameOptions = []string{
	"Rock",
	"Paper",
	"Scissors",
}
var userGameScore int = 0
var compGameScore int = 0

func getCompVote () string {
	var compVote = gameOptions[rand.Intn(len(gameOptions))]
	return compVote
	
}

func getUserVote () string {
	var userGuess string
	userVote := "No value"
	fmt.Println(`
================================

[R]ock, [P]aper or [S]cissors? 

================================
`)
	fmt.Scan(&userGuess)
	if strings.ToUpper(userGuess) == "R"{
		userVote = "Rock"
	} else if strings.ToUpper(userGuess) == "P"{
		userVote = "Paper"
	} else if strings.ToUpper(userGuess) == "S"{
		userVote = "Scissors"
	}

return userVote
}

func playRound () (string, int, int) {
	userVote := getUserVote()
	compVote := getCompVote()
	fmt.Println("--------------------------------")
	fmt.Println(" ")
	fmt.Println("User Vote: " + userVote)	
	fmt.Println("Computer Vote: " + compVote)

	userScore := 2
	compScore := 0
	if userVote == "Rock" {
		if compVote == "Rock"{
			compScore = 2
		} else if compVote == "Paper"{
			compScore = 3
		} else {
			compScore = 1
		}
	}
	if userVote == "Paper" {
		if compVote == "Paper"{
			compScore = 2
		} else if compVote == "Scissors"{
			compScore = 3
		} else {
			compScore = 1
		}
	}
	
	if userVote == "Scissors" {
		if compVote == "Scissors"{
			compScore = 2
		} else if compVote == "Rock"{
			compScore = 3
		} else {
			compScore = 1
		}
	}
	var winnerText string
	//var compGameScore int
	//var userGameScore int

	if compScore > userScore {
		winnerText = "Computer Wins! "
		compGameScore += 1
		//userGameScore = userGameScore
	} else if compScore == userScore {
		winnerText = "It's a draw!"
		//compGameScore = compGameScore
		//userGameScore = userGameScore
	} else {
		winnerText = "User wins!"
		userGameScore += 1
		//compGameScore = compGameScore
	}
 
	return winnerText, compGameScore, userGameScore
}

func main () {
	//compVote := getCompVote()
	//userVote := getUserVote()
	//fmt.Println("User Vote: " + userVote)	
	//fmt.Println("Computer Vote: " + compVote)
	playGame := "Y"
	playReq := "default"
	for playGame == "Y"{
	winnerText, compGameScore, userGameScore := playRound()
	fmt.Println(winnerText)
	fmt.Printf("Computer Score: %d\n", compGameScore)
	fmt.Printf("User Score: %d\n", userGameScore)
	fmt.Println(" ")
	fmt.Println("--------------------------------")
	fmt.Println("Play again? [Y]es or [N]o? ")
	fmt.Scan(&playReq)
	if strings.ToUpper(playReq) == "Y"{
		playGame = "Y"
	} else if strings.ToUpper(playReq) == "N" {
		playGame = "N"
		fmt.Println("Hope you had fun!")
	}
	
}


}
