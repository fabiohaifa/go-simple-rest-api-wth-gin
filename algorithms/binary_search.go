package main

import "fmt"

func main() {

	fmt.Println("Binary Search sample in GoLang!")

	var arrSearch []int
	arrSearch = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	indexSearch := 70

	var lowIndex int = 0
	var highIndex int = len(arrSearch) - 1
	var loop int = 0

	for lowIndex <= highIndex {
		loop++
		median := (lowIndex + highIndex) / 2

		if arrSearch[median] == indexSearch {
			fmt.Println(arrSearch[median])
			fmt.Println("Found!")
			break
		} else if arrSearch[median] < indexSearch {
			lowIndex = median + 1
		} else {
			highIndex = median - 1
		}
	}

	if lowIndex == len(arrSearch) || arrSearch[lowIndex] != indexSearch {
		fmt.Println("Not Found")
	} else {
		fmt.Printf("Found in %d interactions", loop)
	}

	// for i := 0; i < len(arrSearch); i++ {
	// 	checked, arrLook := checkArray(arrSearch, indexSearch)
	// 	if checked {
	// 		fmt.Printf("found in %d interactions", i+1)
	// 		break
	// 	} else {
	// 		arrSearch = arrLook
	// 	}
	// 	fmt.Println(arrSearch[i])
	// }

}

func checkArray(array []int, valueCheck int) (bool, []int) {
	// check recusively the middle of the array
	middle := (len(array) - 0) / 2
	if valueCheck == array[middle] {
		fmt.Println("found!")
		return true, nil
	} else if valueCheck > middle {
		return false, remountArray(array, middle, len(array))
	} else {
		return false, remountArray(array, 0, middle)
	}
}

func remountArray(array []int, startPosition int, endPosition int) []int {
	var retArray []int
	for i := startPosition; i < endPosition; i++ {
		retArray = append(retArray, array[i])
	}
	return retArray
}
