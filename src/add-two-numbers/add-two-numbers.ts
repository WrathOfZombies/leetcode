class ListNode {
  public next: ListNode | null = null;
  constructor(public val: number) {}
}

/**
 * Problem: {@link https://leetcode.com/problems/two-sum/description/}
 * Submission: {@link https://leetcode.com/submissions/detail/157567425/}
 * @summary Runtime: ~132ms, Time: O(n), Space: O(n)
 */
const addTwoNumbers = (numA: ListNode, numB: ListNode) => {
  let A: ListNode = numA;
  let B: ListNode = numB;

  let carry = 0;
  let digit = 0;
  let longer: ListNode | null = A || B;
  let head: ListNode | null = null;
  let result: any = null;

  do {
    ({ carry, digit } = addDigts(A, B, carry));
    ({ A, B, longer } = moveNextAndReturnLonger(A, B));
    if (head) {
      head.next = new ListNode(digit);
      head = head.next;
    } else {
      head = new ListNode(digit);
      result = head;
    }
  } while (longer !== null);
  if (carry !== 0) {
    head.next = new ListNode(carry);
  }
  return result;
};

const moveNextAndReturnLonger = (
  numA: ListNode | null,
  numB: ListNode | null
): any => {
  let A: ListNode | null = null,
    B: ListNode | null = null;
  if (numA != null) {
    A = numA.next;
  }
  if (numB != null) {
    B = numB.next;
  }
  return { A, B, longer: A || B };
};

const addDigts = (
  A: ListNode = new ListNode(0),
  B: ListNode = new ListNode(0),
  carry = 0
) => {
  const sum = (A ? A.val : 0) + (B ? B.val : 0) + carry;
  const digit = sum % 10;
  const newCarry = Math.floor(sum / 10);
  return { digit, carry: newCarry };
};

// Test cases:
// [2,4,3]
// [5,6,4]
// [1,1,1,1,1]
// [9,9]
// [5]
// [5]
// [1]
// [9,9]
//
