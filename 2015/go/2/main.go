package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Open list.txt for reading
	file, _ := os.Open("list.txt")
	//Line-by-line scanner
	scanner := bufio.NewScanner(file)
	total := 0
	ribbon := 0
	//For each line in the scanner
	for scanner.Scan(){
		//Splitting at x and removing x from the input
		//The split string is stored as an array
		arr := strings.Split(scanner.Text(), "x")
		n := []int{}
		//Convert string input to integer
		for _, strval := range arr {
			intval, _ := strconv.Atoi(strval)
			n = append(n, intval)
		}
		//sort the integer array
		sort.Ints(n)
		// surface area calc
		total += 2 * (n[0] * n[1] + n[1] * n[2] + n[2] * n[0]) + n[0] * n[1]
		ribbon += 2 * (n[0] + n[1]) + n[0] * n[1] * n[2]
	}
	fmt.Println(total)
	fmt.Println(ribbon)
}
