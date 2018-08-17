package main

import (
	FUN "../entertainment"
	"fmt"
)

func main() {
	//FUN.EightQueenBit(0)
	//print( FUN.TotalFunctions )

	//FUN.EightQueenArry()

	//qwe := [8][8]string{}
	//
	//FUN.FillArryRow(&qwe, 0, 0)
	//fmt.Printf("fill arry row ret is %s\n", qwe)
	//
	//FUN.FillArryColumn(&qwe, 1, 1)
	//fmt.Printf("fill arry column ret is %s\n", qwe)
	//
	//FUN.FillArrySlash(&qwe, 2, 2, 2)
	//fmt.Printf("fill arry slash ret is %s\n", qwe)

	chessBoardArry := [8][8]string{}
	FUN.EightQueenArry(0, chessBoardArry)
	fmt.Printf("total is %d\n", FUN.ArryTotal)
}