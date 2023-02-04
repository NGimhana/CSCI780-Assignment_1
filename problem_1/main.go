package main

import "fmt"

var ArrayOfNumbers = []uint{2, 7, 15, 11}
var Sum uint = 20

/*
Time complexity: O(n)
Explanation: Time complexity for each line/block of lines is mentioned below

	line 17 -> Create an empty dictionary. O(1)
	line 21-23 -> Insert all numbers into the dictionary. O(n). Here, n = len(ArrayOfNumbers)
	line 27 -> Iterate over the each element in the array.
	Under the Worst case this loop will run n times. Hence O(n).
	line 31 -> Check if the (Sum - number) is present in the dictionary. Since this is a dictionary lookup, it takes O(1) time.
	line 34 -> If the pair is found, print the index and return. This will take O(1) time.
	line 40 -> If no pair found, print the message. This will take O(1) time.

	So, the total time complexity is O(1) + O(n) + O(n) + O(1) + O(1) + O(1) = O(n)

	Space complexity will be O(n) since we are using a dictionary to store the numbers.
*/
func main() {

	// dict = [number, index]
	dict := make(map[uint]uint)

	// insert all numbers into the dictionary
	for index, value := range ArrayOfNumbers {
		dict[value] = uint(index)
	}

	// check if the (Sum - number) is present in the dictionary
	for index, value := range ArrayOfNumbers {
		if _, isKeyPresent := dict[Sum-value]; isKeyPresent {
			fmt.Printf("Index1 : %v, Index2: %v\n", index, dict[Sum-value])
			// return only the 1st pair
			return
		}
	}

	// if no pair found
	fmt.Println("No such pair found")
}
