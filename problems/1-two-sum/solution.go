// Problem: https://leetcode.com/problems/two-sum/
// Submission: https://leetcode.com/problems/two-sum/submissions/848547911/

package twoSum

/*
 * Considerations:
 * 1. The problem states that there's exactly one solution
 * 2. There's possiblity of duplicates but the numbers are considered unique
 * 3. You can answer in any order
 *
 * So using the above we are modifying A + B = target as, B = target - A.
 * We will store all numbers in a map, this also flattens duplicates and the map
 * will always contain the position of the last duplicate.
 * This means when we deal with the case: A + A = target, we don't reuse A unless the indexes
 * are different. Since we are iterating from the left of the array, if we see a dupe, it should have
 * a different index than the one in the map. Else either we haven't seen a dupe, or this is one of those cases
 * where you subtract a number from target and it is the same.
 */

// Returns the positions of two numbers in nums that add up to the target
func TwoSum(nums []int, target int) []int {
	// Add all the numbers to a search map
	search := make(map[int]int, len(nums))
	for posA, a := range nums {
		search[a] = posA
	}

	// Foreach number, check if B = Target - A, exists in the search map
	for posA, a := range nums {
		b := target - a

		// If B was found was different from A, then return the result
		if posB, found := search[b]; found && posA != posB {
			return []int{posA, posB}
		}
	}

	// Bail
	return nil
}
