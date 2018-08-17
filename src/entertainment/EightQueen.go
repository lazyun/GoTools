package entertainment

import (
	"fmt"
	//"os"
	"strconv"
)


// 位计算 8 皇后问题
var (

	ChessBoardBit = [8]int{}
	MaxQueue = 8
    TotalFunctions = 0
)


func EightQueenBit(index int) () {
	if MaxQueue == index {
		return
	}

	for i := 0; i < MaxQueue; i++  {
		temporaryPosition := 1 << uint(i)

		isBreak := false
		for j := 0; j < index; j++ {
			right := temporaryPosition << uint( index - j )
			left := temporaryPosition >> uint( index - j )

			//fmt.Printf("index: %d, temporaryPosition: %d, j: %d, right: %d, left: %d\n", index, temporaryPosition, j, right, left)

			if (temporaryPosition & ChessBoardBit[j]) != 0 ||
				(ChessBoardBit[j] & right) != 0 ||
				(ChessBoardBit[j] & left) != 0 {

					isBreak	= true
					break
			}
		}

		if !isBreak {
			ChessBoardBit[index] = temporaryPosition

			if MaxQueue - 1 == index {
				//ChessBoard = [8]int{}
				//fmt.Printf("ChessBoardBit ret is : %d \n", ChessBoardBit)
				TotalFunctions++
				return
			}

			//fmt.Printf("ChessBoardBit ret is : %d \n", ChessBoardBit)

			//if index == 1 {
			//	os.Exit(1)
			//}


			EightQueenBit( index + 1 )
		}


	}

	fmt.Printf("total ret is %d\n", TotalFunctions)
	//fmt.Printf("now index: %d, %d", index, ChessBoardBit)
	//os.Exit(1)
}



var (
	ArryTotal = 0
	MaxQueueArry = 8
	//chessBoardArry = [8][8]string{}
)

func EightQueenArry(step int, chessBoardArry [8][8]string) {
	//fmt.Printf("this step is %d\t", step)
	if MaxQueueArry == step {
		ArryTotal++
		fmt.Printf("%d final ret is %\n", ArryTotal, chessBoardArry)
		return
	}

	tempotaty := chessBoardArry
	for i :=0; i < MaxQueueArry; i++ {
		if 0 == step {
			tempotaty = [8][8]string{}
			FillArryRow(&tempotaty, step, step)
			FillArryColumn(&tempotaty, step, i)
			FillArrySlash(&tempotaty, step, step, i)
			EightQueenArry( step + 1, tempotaty)
			continue
		}

		if "" != tempotaty[step][i] {
			if MaxQueueArry - 1 == i {
				return
			}
			continue
		}

		if 6 == step {
			fmt.Printf("index %d total %d final ret is %s\n", step, ArryTotal, chessBoardArry)
		}

		if MaxQueueArry - 1 == step {
			fmt.Printf("index %d total %d final ret is %s\n", step, ArryTotal, chessBoardArry)
		}

		FillArryRow(&tempotaty, step, step)
		FillArryColumn(&tempotaty, step, i)
		FillArrySlash(&tempotaty, step, step, i)
		EightQueenArry( step + 1, tempotaty )
	}
	fmt.Printf("index %d total %d final ret is %s\n", step, ArryTotal, chessBoardArry)
}


func FillArryRow(arry *[8][8]string, step, row int) {
	var index int
	for index, _ = range arry[row] {

		arry[row][index] += strconv.Itoa(step)
	}
}

func FillArryColumn(arry *[8][8]string, step, column int) {
	var index int
	for index, _ = range arry {
		arry[index][column] += strconv.Itoa(step)
	}
}

func FillArrySlash(arry *[8][8]string, step, row , column int) {
	for i :=0; i < MaxQueueArry; i++ {
		index := (row + i) % MaxQueueArry
		indey := (column + i) % MaxQueueArry
		arry[index][indey] += strconv.Itoa(step)

		indexx := (row + i) % MaxQueueArry
		indeyy := ( ( (column - i) ^ (column - i) >> 31 ) - ( column - i ) ) >> 31 % MaxQueueArry
		//fmt.Printf("xx is %d, yy is %d row %d, column %d, i %d\n", indexx, indeyy, row, column, i)
		arry[indexx][indeyy] += strconv.Itoa(step)
	}

}