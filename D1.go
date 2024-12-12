package main

// import (
// 	"fmt"
// 	"math"
// 	"os"
// 	"sort"
// 	"strconv"
// 	"strings"
// )

// func main() {

// 	// PART 1

// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 	}
// 	defer file.Close()

// 	content, err := os.ReadFile("input.txt")
// 	if err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}

// 	stringContent := string(content)
// 	nums := strings.Split(stringContent, "\n")
	
// 	leftList := []int{}
// 	rightList := []int{}
// 	for _, v := range(nums) {
// 		if v != "" {
// 			lr := strings.Split(v, "   ")

// 			numL, err := strconv.Atoi(lr[0])
// 			if err != nil {
// 				fmt.Println("Atoi error: ", err)
// 			}
// 			numR, err := strconv.Atoi(lr[1])
// 			if err != nil {
// 				fmt.Println("Atoi error: ", err)
// 			}

// 			leftList = append(leftList, numL)
// 			rightList = append(rightList, numR)
// 		}
// 	}

// 	sort.Ints(leftList)
// 	sort.Ints(rightList)

// 	difs := []int{}
// 	for item, value := range(leftList) {
// 		dif := math.Abs(float64(value) - float64(rightList[item]))
// 		difs = append(difs, int(dif))
// 	}

// 	sum := 0
// 	for _, v := range(difs) {
// 		sum += v
// 	}

// 	fmt.Println(sum)

// 	// PART 2

// 	score := 0
// 	for _, target := range(leftList) {
// 		counter := 0
// 		for _, v := range(rightList) {
// 			if target == v {
// 				counter++
// 			}
// 		}
// 		score += target * counter
// 	}
// 	fmt.Println(score)
// }