/**
 * Problem: {@link https://leetcode.com/problems/two-sum/description/}
 * Submission: {@link https://leetcode.com/submissions/detail/157494454/}
 *
 * @summary Runtime: ~64ms, Time: O(n), Space: O(n)
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
const twoSum = (numbers: number[], target: number): number[] => {
  const pairs = new Map();
  for (let [index, number] of numbers.entries()) {
    // Compute the matching number for the current item
    const mate = target - number;

    // If we can find the current item in our map
    // Then return the saved value and the current index
    // Else save the mate to look for in the future
    if (pairs.has(number)) {
      return [pairs.get(number), index];
    }

    pairs.set(mate, index);
  }
  return [];
};

// Test cases:
// [2,7,11,15]
// 9
// [0,1,2,3]
// 1
// [0,1,-1,2]
// 0
// [0,1,-1,2]
// 1
// [0,0,0,0]
// 0
// [3,3]
// 6
// [-1,2,-1,2]
// -2
// [-1,2,-1,2]
// 4
// [0, 4, 3, 2, 5, 1, 6, 7, 9, 10, 25, 1000]
// 1000
