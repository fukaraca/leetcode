package main

import "fmt"

/*
Input: board =
{{"5","3",".",".","7",".",".",".","."}
,{"6",".",".","1","9","5",".",".","."}
,{".","9","8",".",".",".",".","6","."}
,{"8",".",".",".","6",".",".",".","3"}
,{"4",".",".","8",".","3",".",".","1"}
,{"7",".",".",".","2",".",".",".","6"}
,{".","6",".",".",".",".","2","8","."}
,{".",".",".","4","1","9",".",".","5"}
,{".",".",".",".","8",".",".","7","9"}}
46 49-57
*/

func main() {
	arr := [][]string{{"5", "3", ".", ".", "7", ".", ".", ".", "."}, {"6", ".", ".", "1", "9", "5", ".", ".", "."}, {".", "9", "8", ".", ".", ".", ".", "6", "."}, {"8", ".", ".", ".", "6", ".", ".", ".", "3"}, {"4", ".", ".", "8", ".", "3", ".", ".", "1"}, {"7", ".", ".", ".", "2", ".", ".", ".", "6"}, {".", "6", ".", ".", ".", ".", "2", "8", "."}, {".", ".", ".", "4", "1", "9", ".", ".", "5"}, {".", ".", ".", ".", "8", ".", ".", "7", "9"}}
	bytArr := make([][]byte, 9)
	for i, strings := range arr {
		for _, s := range strings {
			bytArr[i] = append(bytArr[i], s[0])
		}

	}
	fmt.Println(bytArr)
	fmt.Println(isValidSudoku(bytArr))
}
func isValidSudoku(board [][]byte) bool {
	verArr := make([]map[byte]struct{}, 9)
	//ver := make(map[byte]struct{})
	for i := 0; i < 9; i++ {
		verArr[i] = make(map[byte]struct{})
		hor := make(map[byte]struct{})

		for j := 0; j < 9; j++ {
			//horizontal check
			if _, ok := hor[board[i][j]]; ok && board[i][j] != 46 {
				return false
			}
			hor[board[i][j]] = struct{}{}
			//vertical check
			if _, ok := verArr[i][board[j][i]]; ok && board[j][i] != 46 {

				return false
			}
			verArr[i][board[j][i]] = struct{}{}
			//sub-table check

			if (i+1)%3 == 0 && (j+1)%3 == 0 {
				subSud := make(map[byte]struct{})
				for k := (i - 2); k <= i; k++ {
					for l := (j - 2); l <= j; l++ {
						if _, ok := subSud[board[k][l]]; ok && board[k][l] != 46 {
							return false
						}
						subSud[board[k][l]] = struct{}{}
					}
				}

			}
		}

	}
	return true
}
