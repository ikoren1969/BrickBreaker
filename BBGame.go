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


var jsontype jsonobject

func main() {

	readConfig()
	for f:=0;f<jsontype.Object.ScenesNumber;f++ {
		fmt.Println("Scene number:", f+1)
		blocks := prepareBlockArray(f)
		buildScene(blocks, f)
	}
}

func buildScene(blocks [][]blocLoc, iSceneNumber int){

	for rowsCounter:=0; rowsCounter<jsontype.Object.Databases[iSceneNumber].IRows; rowsCounter++ {
		fmt.Println("\nRow:",rowsCounter)
		for colsCounter:=0; colsCounter<jsontype.Object.Databases[iSceneNumber].ICols; colsCounter++ {
			blocks[rowsCounter][colsCounter].iCol=colsCounter
			blocks[rowsCounter][colsCounter].iRow=rowsCounter
			blocks[rowsCounter][colsCounter].iColor = rand.Intn(3)  // 0 - empty block, 1-blue, 2-red
			fmt.Printf("%d",blocks[rowsCounter][colsCounter].iColor)
		}
	}

}


// Prepare a dynamic two dimensional array to hold the blocLoc struct
func prepareBlockArray(iSceneNumber int)[][]blocLoc {
	blocks := make([][]blocLoc, jsontype.Object.Databases[iSceneNumber].IRows)
	for f:=0;f<jsontype.Object.Databases[iSceneNumber].IRows;f++ {
		blocks[f] = make([]blocLoc,jsontype.Object.Databases[iSceneNumber].ICols)
	}
	return blocks
}





func readConfig(){
	dat, err := ioutil.ReadFile("config/general.json")
	check(err)


	json.Unmarshal(dat, &jsontype)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}