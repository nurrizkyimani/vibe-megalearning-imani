# Learn LeetCode DSA

This folder is a LeetCode DSA learning track inspired by the format in
`learn-scholastic-test`.

## Files

| File | Purpose |
| --- | --- |
| `leetcode-curriculum.md` | Topic-by-topic DSA curriculum with concepts, patterns, traps, review rules, and practice questions. |
| `LeetcodeDsaSolutions.java` | Runnable Java examples. Each demo problem has two solution approaches. |

## How to Run the Java File

From the repo root:

```bash
java learn-leetcode-dsa/LeetcodeDsaSolutions.java
```

Expected output includes demo results for:

| Problem | Solution 1 | Solution 2 |
| --- | --- | --- |
| Two Sum | Brute force | Hash map |
| Valid Palindrome | Clean and reverse | Two pointers |
| Best Time to Buy and Sell Stock | Brute force | One pass |
| Valid Anagram | Sorting | Frequency count |
| Valid Parentheses | Open stack | Expected closer stack |
| Binary Search | Iterative | Recursive |
| Reverse Linked List | Iterative | Recursive |
| Maximum Depth of Binary Tree | DFS recursion | BFS levels |
| Number of Islands | DFS | BFS |
| Kth Largest Element | Sorting | Min heap |
| Climbing Stairs | Memoization | Tabulation |

## Study Flow

1. Read one `NumXX` section in `leetcode-curriculum.md`.
2. Solve the practice question without looking at the Java file.
3. Open `LeetcodeDsaSolutions.java` and compare the two approaches.
4. Write the time and space complexity for both approaches.
5. Re-code the optimal approach from memory.

## Notes

- The Java file uses only the standard library.
- Class name and file name must stay the same: `LeetcodeDsaSolutions`.
- The single-command run uses Java source-file mode, available in modern JDKs.
