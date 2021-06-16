package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	code, err := ioutil.ReadFile("code.txt")
	if err != nil {
		panic(err)
	}
	getFinalFloor(code)
}
func getFinalFloor(code []byte){
	floor := 0
	basement := 0
	for i := 0; i < len(code); i++ {
		if code[i] == '('{
			floor++
		} else {
			floor--
		}
		if floor == -1 && basement == 0{
			basement = 1
			fmt.Println("First basement hit:", i + 1)
		}
	}
	fmt.Println("Final floor:", floor)
}
