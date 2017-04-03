package main

import (
	"fmt"
	"math/rand"
	"io/ioutil"
	"encoding/json"
)



type blocLoc struct {
	iRow int
	iCol int
	iColor int
}


const iNumberOfRows int = 12
const iNumberOfCols int = 5


type jsonobject struct {
	Object ObjectType
}

type ObjectType struct {
	ScenesNumber int
	Databases   []DatabasesType
}

type DatabasesType struct {
	IRows   int
	ICols   int
}

func main() {

	fmt.Println("reading configuration...")
	readConfig()
	blocks := prepareBlockArray(1)
	buildScene(blocks)
}

func buildScene(blocks [][]blocLoc){

	for rowsCounter:=0; rowsCounter<iNumberOfRows; rowsCounter++ {
		fmt.Println("\nRow:",rowsCounter)
		for colsCounter:=0; colsCounter<iNumberOfCols; colsCounter++ {
			blocks[rowsCounter][colsCounter].iCol=colsCounter
			blocks[rowsCounter][colsCounter].iRow=rowsCounter
			blocks[rowsCounter][colsCounter].iColor = rand.Intn(3)  // 0 - empty block, 1-blue, 2-red
			fmt.Printf("%d",blocks[rowsCounter][colsCounter].iColor)
		}
	}

}


// Prepare a dynamic two dimensional array to hold the blocLoc struct
func prepareBlockArray(iSceneNumber int)[][]blocLoc {
	blocks := make([][]blocLoc,iNumberOfRows)
	for f:=0;f<iNumberOfRows;f++ {
		blocks[f] = make([]blocLoc,iNumberOfCols)
	}
	return blocks
}





func readConfig(){
	dat, err := ioutil.ReadFile("config/general.json")
	check(err)

	var jsontype jsonobject
	json.Unmarshal(dat, &jsontype)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}