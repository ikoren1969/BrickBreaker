package main

import (
	"fmt"
	"math/rand"
)



type blocLoc struct {
	iRow int
	iCol int
	iColor int
}


const iNumberOfRows int = 10
const iNumberOfCols int = 5


func main() {

	fmt.Printf("hi")
	blocks := prepareBlockArray()
	buildScene(blocks)
}

func buildScene(blocks [][]blocLoc){

	for rowsCounter:=0; rowsCounter<iNumberOfRows; rowsCounter++ {
		fmt.Println("\nRow:",rowsCounter)
		for colsCounter:=0; colsCounter<iNumberOfCols; colsCounter++ {
			blocks[rowsCounter][colsCounter].iCol=colsCounter
			blocks[rowsCounter][colsCounter].iRow=rowsCounter
			blocks[rowsCounter][colsCounter].iColor = rand.Intn(3)
			fmt.Printf("%d",blocks[rowsCounter][colsCounter].iColor)
		}
	}

}


// Prepare a dynamic two dimensional array to hold the blocLoc struct
func prepareBlockArray()[][]blocLoc {
	blocks := make([][]blocLoc,iNumberOfRows)
	for f:=0;f<iNumberOfRows;f++ {
		blocks[f] = make([]blocLoc,iNumberOfCols)
	}

	return blocks

}