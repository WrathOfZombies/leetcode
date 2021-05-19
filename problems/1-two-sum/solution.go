package problems

func main(nums []int, target int) []int {
	numberMap := make(map[int]int)

	// For each number, see if its compliment exists
	for index, number := range nums {
		compliment := target - number

		// (TIP: Second value in a map assignment is a boolean)
		// Find the compliment in the map and if so, return the result
		if complimentIndex, found := numberMap[compliment]; found {
			return []int{index, complimentIndex}
		} else {
			// Add the current number to the map to be found later
			numberMap[number] = index
		}
	}

	// Return default
	return []int{0, 0}
}
