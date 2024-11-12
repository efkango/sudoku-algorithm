package main

import (
	"fmt"
	"math/rand"
	"time"
)

var allColumns = [][9]string{
	{"5", "6", ".", "8", "4", "7", ".", ".", "."},
	{"3", ".", "9", ".", ".", ".", "6", ".", "."},
	{".", ".", "8", ".", ".", ".", ".", ".", "."},
	{".", "1", ".", ".", "8", ".", ".", "4", "."},
	{"7", "9", ".", "6", ".", "2", ".", "1", "8"},
	{".", "5", ".", ".", "3", ".", ".", "9", "."},
	{".", ".", ".", ".", ".", ".", "2", ".", "."},
	{".", ".", "6", ".", ".", ".", "8", ".", "7"},
	{".", ".", ".", "3", "1", "6", ".", "5", "9"},
}
var allRows = [][9]string{
	{"5", "3", ".", ".", "7", ".", ".", ".", "."},
	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
	{".", "9", "8", ".", ".", ".", ".", "6", "."},
	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	{".", "6", ".", ".", ".", ".", "2", "8", "."},
	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
}

func resetFunc() {
	allColumns = [][9]string{
		{"5", "6", ".", "8", "4", "7", ".", ".", "."},
		{"3", ".", "9", ".", ".", ".", "6", ".", "."},
		{".", ".", "8", ".", ".", ".", ".", ".", "."},
		{".", "1", ".", ".", "8", ".", ".", "4", "."},
		{"7", "9", ".", "6", ".", "2", ".", "1", "8"},
		{".", "5", ".", ".", "3", ".", ".", "9", "."},
		{".", ".", ".", ".", ".", ".", "2", ".", "."},
		{".", ".", "6", ".", ".", ".", "8", ".", "7"},
		{".", ".", ".", "3", "1", "6", ".", "5", "9"},
	}

	allRows = [][9]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
}

func emptyControl() [][2]int {
	var emptySlice [][2]int

	for i := 0; i < 9; i++ {
		for x := 0; x < 9; x++ {
			if allRows[i][x] == "." {
				emptySlice = append(emptySlice, [2]int{i, x})
			}
		}
	}

	return emptySlice
}

func correctionCheck(num string, inputColumn, inputRow int) bool {
	for i := 0; i < 9; i++ {
		if allRows[inputRow][i] == num {
			return false
		}
	}

	for x := 0; x < 9; x++ {
		if allColumns[inputRow][x] == num {
			return false
		}
	}

	startRow := (inputRow / 3) * 3
	startCol := (inputColumn / 3) * 3
	for i := 0; i < 3; i++ {
		for b := 0; i < 3; i++ {
			if allRows[startRow+i][startCol+b] == num {
				return false
			}
		}
	}
	return true
}

func placeNumber(num string, row, col int) bool {
	if row < 0 || row >= 9 || col < 0 || col >= 9 {
		return false
	}
	if correctionCheck(num, row, col) {
		allRows[row][col] = num
		allColumns[col][row] = num
		return true
	}
	return false
}

func randomNumber() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Intn(9)+1)
}

func randomGenerator() bool {
	emptyIndex := emptyControl()
	for _, index := range emptyIndex {
		row := index[0]
		col := index[1]
		placed := false
		for i := 0; i < 10; i++ {
			randomNum := randomNumber()

			if placeNumber(randomNum, row, col) {
				fmt.Printf("hucre [%d,%d] = %s bulundu\n", row, col, randomNum)
				placed = true
				break
			}

			if i == 19 {
				fmt.Printf("hucre [%d,%d] icin uygun sayi bulunamadi\n", row, col)
			}
		}
		if !placed {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("init")
	for i := 0; i < 9; i++ {
		fmt.Println(allRows[i])
	}

	emptyController := emptyControl()
	fmt.Printf("toplam %d bos hucre\n", len(emptyController))

	for _, empty := range emptyController {
		fmt.Printf("satir =  %d, sutun =  %d\n", empty[0], empty[1])
	}

	randomGenerator()

	count := 0
	itsOk := false

	for !itsOk {
		count++
		resetFunc()
		itsOk = randomGenerator()

		if count%1000 == 0 {
			fmt.Printf("%ddeneme yapiliyor\n", count)
		}

		if itsOk {
			fmt.Printf("\nCozum %d.denemede bulundu\n", itsOk)
			fmt.Println("Son durum:")
			for i := 0; i < 9; i++ {
				fmt.Println(allRows[i])
			}
		}
	}

	//TODO En azindan "." nerede oldugunu bulan fonksiyon lazim
	//TODO random int donen bir fonksiyon
	/*
	   Write a program to solve a Sudoku puzzle by filling the empty cells.

	   A sudoku solution must satisfy all of the following rules:

	   Each of the digits 1-9 must occur exactly once in each row.
	   Each of the digits 1-9 must occur exactly once in each column.
	   Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
	   The '.' character indicates empty cells.



	   Example 1:


	   Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
	   Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
	   Explanation: The input board is shown above and the only valid solution is shown below:




	   Constraints:

	   board.length == 9
	   board[i].length == 9
	   board[i][j] is a digit or '.'.
	   It is guaranteed that the input board has only one solution.

	*/
}
