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
	firstBasementHit(code)
}
func getFinalFloor(code []byte){
	floor := 0
	for i := 0; i < len(code); i++ {
		if code[i] == '('{
			floor++
		} else{
			floor--
		}
	}
	fmt.Println(floor)
}

func firstBasementHit(code []byte){
	floor := 0
	for i := 0; i < len(code); i++ {
		if code[i] == '('{
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}