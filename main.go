package main

import "fmt"

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

// ilk parantez hangi liste ikinci parantez o listenin hangi elemani

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

func correctionCheck(num string, rowNumber, columnNumber int) bool {
	for i := 0; i < 9; i++ {
		if allRows[rowNumber][i] == num {
			return false
		}
	}

	for a := 0; a < 9; a++ {
		if allColumns[columnNumber][a] == num {
			return false
		}
	}

	// 3x3 kutuları kontrol et
	startRow := (rowNumber / 3) * 3    // 7 ÷ 3 = 2 , 2 * 3 = 6 // iki ornekte bu matematiksel ifadenin isleyisi mevcut 7/3 gibi bir islemden
	startCol := (columnNumber / 3) * 3 //4 ÷ 3 = 1 , 1 * 3 = 3  // variable larimiz int oldugu icin tam sayi cikmaktadir
	for i := 0; i < 3; i++ {           //yukaridan asagi 3 kare
		for b := 0; b < 3; b++ { //soldan saga 3 kare
			if allRows[startRow+i][startCol+b] == num {
				return false
			}
		}
	}

	return true
}

func fixGrids() bool {
	for rowNumber := 0; rowNumber < 9; rowNumber++ { //satirlari dondurur
		for colIndex := 0; colIndex < 9; colIndex++ { // sutunlari dondurur
			if allRows[rowNumber][colIndex] == "." { // eger "." var ise
				for num := 1; num <= 9; num++ { // 1 ile 9 arasi rakamlar
					numStr := fmt.Sprintf("%d", num)
					if correctionCheck(numStr, rowNumber, colIndex) { // "." yerine neler gelebilir
						placeNumber(numStr, rowNumber, colIndex) //gecici olarak tabloya yerlestir
						if fixGrids() {                          //recursive call geri kalanini yerlestir
							return true
						}
						removeNumber(rowNumber, colIndex) // eger kod buraya geldi yani false ise koyulan gecici numarayi degistirip sirasiyla diger rakama geciyor
					}
				}
				return false //not ok
			}
		}
	}
	return true //hepsi ok
}

func placeNumber(numStr string, row, col int) {
	allRows[row][col] = numStr
	allColumns[col][row] = numStr
}

func removeNumber(row, col int) {
	allRows[row][col] = "."
	allColumns[col][row] = "."
}

func main() {
	if fixGrids() {
		fmt.Println("islem tamamdir!")
		// cozumu yazdir
		for i := 0; i < 9; i++ {
			fmt.Println(allRows[i])
		}
	} else {
		fmt.Println("islem tamamlanmadi")
	}
}

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
