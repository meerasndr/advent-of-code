package main
import (
	"fmt"
)

func main(){
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