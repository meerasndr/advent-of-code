package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	floor := 0
	code, err := ioutil.ReadFile("code.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(code); i++ {
		if code[i] == '('{
			floor++
		} else{
			floor--
		}
	}
	fmt.Println(floor)
}