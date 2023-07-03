package main

import (
	"fmt"
	"strconv"
)

var board = [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
var player = "X"
var winner = ""

func main() {
	for winner == "" {
		takeTurn()
	}

	if winner == "draw" {
		printOutput("It's a draw!")
	} else {
		printBoard()
		printOutput("Player " + winner + " wins!")
	}
}

func getUserInputIndex() int {
	var input string
	fmt.Scanln(&input)
	inputInt, err := strconv.Atoi(input)
	if err != nil || inputInt < 1 || inputInt > 9 {
		fmt.Println("Invalid input")
		return getUserInputIndex()
	}
	return inputInt - 1
}

func printOutput(output string) {
	fmt.Println(output)
}

func printBoard() {
	fmt.Println(" " + board[0] + " | " + board[1] + " | " + board[2] + " ")
	fmt.Println("-----------")
	fmt.Println(" " + board[3] + " | " + board[4] + " | " + board[5] + " ")
	fmt.Println("-----------")
	fmt.Println(" " + board[6] + " | " + board[7] + " | " + board[8] + " ")
}

func switchPlayer() {
	if player == "X" {
		player = "O"
	} else {
		player = "X"
	}
}

func checkForWinner() {
	winConditions := [][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, condition := range winConditions {
		if board[condition[0]] == board[condition[1]] && board[condition[1]] == board[condition[2]] && board[condition[0]] != " " {
			winner = player
			return
		}
	}

	isDraw := true
	for _, spot := range board {
		if spot == " " {
			isDraw = false
			break
		}
	}
	if isDraw {
		winner = "draw"
	}
}

func takeTurn() {
	fmt.Println("\n\nPlayer " + player + " turn\n")
	printBoard()
	fmt.Println("\nEnter a number between 1 and 9")
	index := getUserInputIndex()
	if board[index] == " " {
		board[index] = player
	} else {
		fmt.Println("Invalid move")
		takeTurn()
	}

	checkForWinner()

	switchPlayer()
}
