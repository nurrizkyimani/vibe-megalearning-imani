# M4/W4/D33 - Sun, 03 May 2026 (WIB)

## LeetCode DSA - Curriculum

This LeetCode DSA curriculum builds the ability to recognize problem patterns,
choose the right data structure, write clean Java solutions, and explain time
and space trade-offs. This file is designed as one place to study: each numbered
section contains the concept, question pattern, answer pattern, common traps,
review rule, and direct practice.

Target audience: learners who want to start from the basics and gradually move
into DSA patterns that commonly appear in interviews and LeetCode practice.

---

## Curriculum Table

| #   | Topic                   | Core Concept                                           | LeetCode Usage                                                 |
| --- | ----------------------- | ------------------------------------------------------ | -------------------------------------------------------------- |
| 01  | Array and Hash Map      | Fast lookup, frequency, index, and value pairs         | Two Sum, Contains Duplicate, Group Anagrams, Max Palindromes   |
| 02  | Two Pointers            | Two indices moving in the same or opposite directions  | Valid Palindrome, Two Sum II, Container With Most Water        |
| 03  | Sliding Window          | Dynamic windows for subarrays or substrings            | Best Time to Buy and Sell Stock, Longest Substring, Min Window |
| 04  | Stack                   | Last-in-first-out behavior for pairs, order, and undo  | Valid Parentheses, Min Stack, Daily Temperatures               |
| 05  | Binary Search           | Ordered search space and left-right invariants         | Binary Search, Search Insert Position, Search in Rotated Array |
| 06  | Linked List             | Pointers, nodes, dummy heads, and link reversal        | Reverse Linked List, Merge Two Lists, Linked List Cycle        |
| 07  | Tree DFS and BFS        | Recursive traversal or level-order traversal           | Maximum Depth, Invert Tree, Level Order Traversal              |
| 08  | Graph DFS and BFS       | Visited set, components, queue, and recursion boundary | Number of Islands, Clone Graph, Course Schedule                |
| 09  | Heap and Priority Queue | Get the best element without fully sorting everything  | Kth Largest, Top K Frequent, Merge K Lists                     |
| 10  | Dynamic Programming     | State, transition, base case, and subproblem reuse     | Climbing Stairs, House Robber, Coin Change                     |

---

## Grind Plan

| Order | Focus                         | Target      | Output                                                         |
| ----- | ----------------------------- | ----------- | -------------------------------------------------------------- |
| 01-03 | Array, pointer, and window    | Num01-Num03 | Notes on lookup patterns, pointer movement, and window update  |
| 04-06 | Linear structures             | Num04-Num06 | Notes on stack state, binary search invariants, and node flow  |
| 07-08 | Recursive and graph traversal | Num07-Num08 | Notes on DFS/BFS, visited rules, boundaries, and base cases    |
| 09-10 | Optimization and state        | Num09-Num10 | Notes on heap comparators, DP state, transition, and base case |

Timing rules:

- Num01-Num03: 20-35 minutes per problem until you can explain complexity.
- Num04-Num06: 25-45 minutes per problem, with a required pointer or stack-state
  sketch.
- Num07-Num08: 35-60 minutes per problem, with a required base case and visited
  rule.
- Num09-Num10: 45-75 minutes per problem, with a required comparison between
  brute force and the optimal approach.
- Review must record the exact failure point: wrong data structure, wrong
  invariant, wrong boundary, wrong pointer update, wrong visited rule, wrong
  base case, wrong DP state, or wrong complexity.

## Prompt

```
Classify these LeetCode questions by topic and insert each one into the right
NumXX domain section in learn-leetcode-dsa/leetcode-curriculum.md as practice
questions. Keep each answer, Java idea, complexity, and explanation under its
question. Do not make a mixed simulation section. Only create a new topic
section if no existing domain fits.

```

```
visualize image of that solusion iwth the quesiton; focus on the visiual cue ; not make the iamge full code; show visual of the movement; so that incan imagine it, focus on making the visual easy to understand easy to bayangin; the goals of the user to eundersnts easily

use chatgpt image, generate image; generate image

```

---

## Num01 - Array and Hash Map

### The Problem

Array problems often look like ordinary calculation tasks, but the real test is
whether you can avoid repeated scans. If a value must be searched again and
again, there is usually an opportunity to use a `HashMap` or `HashSet`.

Basic pattern example:

```text
nums = [2, 7, 11, 15], target = 9
```

We need two numbers whose sum is `9`. Brute force tries every pair, while the
hash map stores values that have already been seen.

### The Concept

A hash map gives average `O(1)` lookup. The key is deciding what information
must be stored for later: a pair value, frequency, index, or group.

### Pattern of Question

```text
array -> need pair/frequency/duplicate -> fast lookup
```

### Pattern to Answer

1. Decide what gets searched repeatedly.
2. Choose the hash map key and value.
3. Scan once while checking the answer before or after insertion.
4. Explain why a nested loop is unnecessary.

```text
current value -> complement/frequency -> map -> answer
```

### Common Traps

| Trap                | Example                                              |
| ------------------- | ---------------------------------------------------- |
| Inserting too early | Reusing the same element twice in Two Sum            |
| Wrong key           | Storing index as key when the lookup needs the value |
| Missing duplicates  | Frequency is needed, not just existence              |

### Review Rule

For every hash map problem, write:

```text
Key = ...
Value = ...
```

### Practice Questions

#### 1. Two Sum

**Question**

Given `nums = [2, 7, 11, 15]` and `target = 9`, return the indices of two
numbers whose sum equals the target.

**Answer: `[0, 1]`**

**Explanation**  
The number `2` needs complement `7`. When the scan reaches `7`, the map already
stores `2` at index `0`, so the answer is index `0` and index `1`.

Optimal complexity: `O(n)` time and `O(n)` space.

#### 2. 3035. Maximum Palindromes After Operations

**Question**

Given `words = ["abbb", "ba", "aa"]`, you may swap any character from any word
with any character from any other word. Return the maximum number of words that
can become palindromes.

**Answer: `3`**

**Pattern to Answer**

Because any character can be swapped globally, the original word order does not
matter. Only two things matter:

```text
character frequencies -> total available pairs
word lengths -> pair slots needed by each palindrome
```

Steps:

1. Count all characters across all words.
2. Convert counts into available pairs with `count / 2`.
3. Sort word lengths from shortest to longest.
4. For each length, spend `length / 2` pairs.
5. Count the word if there are enough pairs.

**Explanation**

The words have character counts `a:4` and `b:4`, so there are `2 + 2 = 4`
available pairs. Sorted lengths are `[2, 2, 4]`, which need `1`, `1`, and `2`
pairs. All pair slots can be filled after swaps arrange the letters as
palindromes. One possible result is `["bbbb", "aa", "aa"]`.

Optimal complexity: `O(totalChars + n log n)` time and `O(n)` extra space for
the lengths.

![](images/Pasted%20image%2020260503221415.png)

---

## Num02 - Two Pointers

### The Problem

Two pointers are useful when the order of the data gives a reason to move in a
specific direction. For a palindrome string, the left and right pointers move
toward each other. For a sorted array, a sum that is too small moves the left
pointer, while a sum that is too large moves the right pointer.

### The Concept

Two pointers reduce unnecessary combinations through a clear movement rule. Each
step must safely discard part of the search space.

### Pattern of Question

```text
array/string -> compare two positions -> move pointer with a reason
```

### Pattern to Answer

1. Choose the starting pointers.
2. Write the stopping condition.
3. Write the rule for moving the left or right pointer.
4. Make sure no important character or number is skipped.

```text
left/right -> compare -> move one pointer -> repeat
```

### Common Traps

| Trap                  | Example                                           |
| --------------------- | ------------------------------------------------- |
| Pointer does not move | Infinite loop when values are equal               |
| Wrong character skip  | Valid Palindrome forgets to skip non-alphanumeric |
| Data is not sorted    | Using sorted-array pointer logic on unsorted data |

### Review Rule

After solving, write the invariant:

```text
All positions outside the pointers are already valid or safely discarded.
```

### Practice Question

**Question**  
Is the string `"A man, a plan, a canal: Panama"` a palindrome if uppercase
letters, spaces, and punctuation are ignored?

**Answer: `true`**

**Explanation**  
Compare alphanumeric characters from the left and right. Every pair matches
after normalization to lowercase.

---

## Num03 - Sliding Window

### The Problem

Sliding window is used when the answer lives inside a contiguous subarray or
substring. Instead of trying every range, you maintain one active window and
move its boundaries.

### The Concept

There are two common forms:

| Form           | When to Use                                       |
| -------------- | ------------------------------------------------- |
| Fixed window   | The window length is already known                |
| Dynamic window | The window grows and shrinks based on a condition |

### Pattern of Question

```text
contiguous subarray/substring -> find max/min/count -> update window
```

### Pattern to Answer

1. Decide what state the window must track.
2. Move right to add an element.
3. Move left when the window violates the rule.
4. Update the answer only when the window is valid.

### Common Traps

| Trap                  | Example                                     |
| --------------------- | ------------------------------------------- |
| Wrong answer update   | Updating while the window is still invalid  |
| Missing shrink step   | Window keeps growing and violates the rule  |
| Non-contiguous target | Using sliding window for a free subsequence |

### Review Rule

Write the valid-window condition:

```text
The window is valid when ...
```

### Practice Question

**Question**  
Given prices `[7, 1, 5, 3, 6, 4]`, find the maximum profit from one buy and one
sell.

**Answer: `5`**

**Explanation**  
Buy at price `1`, sell at price `6`, and get profit `5`. The optimal solution
stores the minimum price seen so far and the best profit.

---

## Num04 - Stack

### The Problem

A stack fits problems where the most recent unresolved element must be solved
first, such as a closing bracket that must match the latest opening bracket.

### The Concept

A stack stores unfinished state. When an element arrives that resolves that
state, you `pop` and validate it.

### Pattern of Question

```text
symbol/number sequence -> need latest pair -> push/pop
```

### Pattern to Answer

1. Push opening elements or candidates.
2. When a closing element appears, check the stack top.
3. If it does not match, return false.
4. At the end, the stack must be empty for pairing problems.

### Common Traps

| Trap                   | Example                                        |
| ---------------------- | ---------------------------------------------- |
| Popping an empty stack | The string starts with a closing bracket       |
| Counting only quantity | `([)]` has the right count but the wrong order |
| Forgetting leftovers   | An opening bracket remains unclosed            |

### Review Rule

For stack problems, draw at least three states:

```text
input char -> stack before -> stack after
```

### Practice Question

**Question**  
Is the string `"({[]})"` valid?

**Answer: `true`**

**Explanation**  
Every closing bracket matches the latest opening bracket on the stack, and the
stack is empty at the end.

---

## Num05 - Binary Search

### The Problem

Binary search is not only for finding a number in an array. The pattern is to
search inside a space that can be split in half because it is ordered or has a
monotonic property.

### The Concept

The key to binary search is the invariant: which part can still contain the
answer. Do not only memorize `mid`.

### Pattern of Question

```text
sorted/monotonic -> check middle -> discard half of the search space
```

### Pattern to Answer

1. Choose `left` and `right`.
2. Compute safe `mid`: `left + (right - left) / 2`.
3. If the target is smaller, move the right boundary.
4. If the target is larger, move the left boundary.
5. Explain the return value when the target is not found.

### Common Traps

| Trap                | Example                                         |
| ------------------- | ----------------------------------------------- |
| Off-by-one          | Loop does not stop or skips the target          |
| Mid overflow        | `(left + right) / 2` on very large indices      |
| Wrong insert return | Search Insert Position needs the final position |

### Review Rule

Before coding, write the interval:

```text
Search space: inclusive [left, right] or half-open [left, right)
```

### Practice Question

**Question**  
Given `nums = [-1, 0, 3, 5, 9, 12]` and `target = 9`, return the target index.

**Answer: `4`**

**Explanation**  
Binary search discards half of the array until `9` is found at index `4`.

---

## Num06 - Linked List

### The Problem

Linked list problems test whether you can move pointers without losing the rest
of the list. A common mistake is changing `next` too early before saving the next
node.

### The Concept

A node only knows `next`. Because of that, operations such as reversal usually
need three pointers: `prev`, `curr`, and `next`.

### Pattern of Question

```text
node -> pointer changes -> old connection must be saved
```

### Pattern to Answer

1. Save `next` before changing the link.
2. Point `curr.next` to the new node.
3. Move `prev` and `curr`.
4. Return the new head.

### Common Traps

| Trap              | Example                                  |
| ----------------- | ---------------------------------------- |
| Losing nodes      | `curr.next` changes before next is saved |
| Wrong return      | Returning the old head after reversing   |
| Missing null case | Empty list or single-node list           |

### Review Rule

Draw the pointer state:

```text
prev <- curr    next
```

### Practice Question

**Question**  
Reverse linked list `1 -> 2 -> 3 -> 4 -> 5`.

**Answer: `5 -> 4 -> 3 -> 2 -> 1`**

**Explanation**  
Iterate from left to right while reversing the direction of `next`. The new head
is the last node of the original list.

---

## Num07 - Tree DFS and BFS

### The Problem

Tree problems can often be solved with recursive DFS or queue-based BFS. The
choice depends on whether you need a result per branch or per level.

### The Concept

DFS goes as deep as possible before returning. BFS reads the tree level by
level.

### Pattern of Question

```text
root -> left/right subtree or level -> combine result
```

### Pattern to Answer

1. Write the base case for `null`.
2. Choose DFS or BFS traversal.
3. For DFS, combine the left and right answers.
4. For BFS, process the queue size per level.

### Common Traps

| Trap                | Example                                       |
| ------------------- | --------------------------------------------- |
| Missing base case   | NullPointerException on a null child          |
| Wrong level count   | Depth increases per node instead of per level |
| Mixed DFS/BFS state | Queue and recursion are used without a plan   |

### Review Rule

For recursive tree problems, write:

```text
Base case: ...
Combine: ...
```

### Practice Question

**Question**  
Find the maximum depth of tree `[3, 9, 20, null, null, 15, 7]`.

**Answer: `3`**

**Explanation**  
The deepest path is `3 -> 20 -> 15` or `3 -> 20 -> 7`, so the depth is `3`.

---

## Num08 - Graph DFS and BFS

### The Problem

Graph problems test how to visit nodes without repeating forever. In a grid,
each cell can be treated as a node connected to its up, down, left, and right
neighbors.

### The Concept

Graph DFS/BFS always needs a visited rule. For grid mutation, marking a cell as
visited can be done by changing the cell value.

### Pattern of Question

```text
node/grid -> neighbor -> visited -> component/count/path
```

### Pattern to Answer

1. Define the neighbor representation.
2. Check the boundary.
3. Mark visited before going deeper.
4. Count components or distance based on the problem.

### Common Traps

| Trap                  | Example                                  |
| --------------------- | ---------------------------------------- |
| Marking visited late  | Node enters the queue multiple times     |
| Wrong boundary check  | Accessing an index outside the grid      |
| Invalid diagonal move | Number of Islands only uses 4 directions |

### Review Rule

Before coding, write the movement directions:

```text
Directions = up, down, left, right
```

### Practice Question

**Question**  
Count the number of islands in this grid:

```text
11110
11010
11000
00000
```

**Answer: `1`**

**Explanation**  
All `1` cells connected horizontally or vertically form one island.

---

## Num09 - Heap and Priority Queue

### The Problem

A heap is useful when you need the largest or smallest element repeatedly, but
you do not need to sort all data every time.

### The Concept

A priority queue keeps the best element at the top. For top K problems, a heap
of size `k` is often cheaper than sorting every element.

### Pattern of Question

```text
stream/list -> need top/min/max k -> heap
```

### Pattern to Answer

1. Decide whether you need a min-heap or max-heap.
2. Insert elements with the correct comparator.
3. For top K, keep the heap size at `k`.
4. Explain the trade-off against sorting.

### Common Traps

| Trap                | Example                                       |
| ------------------- | --------------------------------------------- |
| Reversed comparator | Using a min-heap when a max-heap is needed    |
| Heap too large      | Storing every element when size `k` is enough |
| Missing tie rule    | Equal frequencies may have flexible ordering  |

### Review Rule

Write what the heap represents:

```text
The heap stores candidates for ...
```

### Practice Question

**Question**  
Given `nums = [3, 2, 1, 5, 6, 4]` and `k = 2`, find the 2nd largest element.

**Answer: `5`**

**Explanation**  
With a min-heap of size `2`, the heap stores the two largest numbers seen so far.
The heap top after all numbers are processed is the 2nd largest element.

---

## Num10 - Dynamic Programming

### The Problem

DP is useful when a brute-force solution computes the same subproblem many
times. The challenge is not only coding, but defining the state correctly.

### The Concept

Four DP components:

| Component  | Question                                              |
| ---------- | ----------------------------------------------------- |
| State      | What does `dp[i]` represent?                          |
| Transition | How is the current state built from older states?     |
| Base case  | What starting value is definitely correct?            |
| Order      | Are states computed from small to large, or backward? |

### Pattern of Question

```text
repeated choices -> overlapping subproblem -> state -> transition
```

### Pattern to Answer

1. Mentally try the brute-force recursion first.
2. Find the input that gets computed repeatedly.
3. Define the state.
4. Write the base case and transition.
5. Choose memoization or tabulation.

### Common Traps

| Trap               | Example                                  |
| ------------------ | ---------------------------------------- |
| Vague state        | `dp[i]` does not have a clear meaning    |
| Wrong base case    | `n = 0` or `n = 1` is not handled        |
| Invalid transition | Using a state that has not been computed |

### Review Rule

For every DP problem, write:

```text
dp[i] = ...
transition = ...
base case = ...
```

### Practice Question

**Question**  
For `n = 5`, how many ways can you climb stairs if each move can climb `1` or
`2` steps?

**Answer: `8`**

**Explanation**  
`dp[i] = dp[i - 1] + dp[i - 2]`. With base cases `dp[0] = 1` and `dp[1] = 1`,
the result for `n = 5` is `8`.
