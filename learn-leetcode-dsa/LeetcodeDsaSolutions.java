import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.PriorityQueue;
import java.util.Queue;

public class LeetcodeDsaSolutions {
    public static void main(String[] args) {
        runTwoSumDemo();
        runMaximumPalindromesDemo();
        runPalindromeDemo();
        runStockProfitDemo();
        runAnagramDemo();
        runValidParenthesesDemo();
        runBinarySearchDemo();
        runLinkedListDemo();
        runTreeDepthDemo();
        runIslandDemo();
        runKthLargestDemo();
        runClimbingStairsDemo();
    }

    private static void runTwoSumDemo() {
        int[] nums = {2, 7, 11, 15};
        int target = 9;

        System.out.println("Two Sum");
        System.out.println("Solution 1 - Brute force: " + Arrays.toString(twoSumBruteForce(nums, target)));
        System.out.println("Solution 2 - Hash map:    " + Arrays.toString(twoSumHashMap(nums, target)));
        System.out.println();
    }

    private static int[] twoSumBruteForce(int[] nums, int target) {
        for (int i = 0; i < nums.length; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                if (nums[i] + nums[j] == target) {
                    return new int[] {i, j};
                }
            }
        }
        return new int[] {-1, -1};
    }

    private static int[] twoSumHashMap(int[] nums, int target) {
        Map<Integer, Integer> seenIndex = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            if (seenIndex.containsKey(complement)) {
                return new int[] {seenIndex.get(complement), i};
            }
            seenIndex.put(nums[i], i);
        }

        return new int[] {-1, -1};
    }

    private static void runMaximumPalindromesDemo() {
        String[] words = {"abbb", "ba", "aa"};

        System.out.println("3035. Maximum Palindromes After Operations");
        System.out.println("Solution 1 - Pair budget: " + maximumPalindromesByPairBudget(words));
        System.out.println("Solution 2 - Min slots:   " + maximumPalindromesByMinSlots(words));
        System.out.println();
    }

    private static int maximumPalindromesByPairBudget(String[] words) {
        int[] counts = new int[26];
        int[] lengths = new int[words.length];

        for (int i = 0; i < words.length; i++) {
            lengths[i] = words[i].length();
            for (int j = 0; j < words[i].length(); j++) {
                counts[words[i].charAt(j) - 'a']++;
            }
        }

        int pairs = 0;
        for (int count : counts) {
            pairs += count / 2;
        }

        Arrays.sort(lengths);

        int palindromes = 0;
        for (int length : lengths) {
            int neededPairs = length / 2;
            if (pairs < neededPairs) {
                break;
            }
            pairs -= neededPairs;
            palindromes++;
        }

        return palindromes;
    }

    private static int maximumPalindromesByMinSlots(String[] words) {
        int[] counts = new int[26];
        PriorityQueue<Integer> pairSlots = new PriorityQueue<>();

        for (String word : words) {
            pairSlots.add(word.length() / 2);
            for (int i = 0; i < word.length(); i++) {
                counts[word.charAt(i) - 'a']++;
            }
        }

        int pairs = 0;
        for (int count : counts) {
            pairs += count / 2;
        }

        int palindromes = 0;
        while (!pairSlots.isEmpty() && pairs >= pairSlots.peek()) {
            pairs -= pairSlots.remove();
            palindromes++;
        }

        return palindromes;
    }

    private static void runPalindromeDemo() {
        String input = "A man, a plan, a canal: Panama";

        System.out.println("Valid Palindrome");
        System.out.println("Solution 1 - Clean and reverse: " + isPalindromeByCleaning(input));
        System.out.println("Solution 2 - Two pointers:      " + isPalindromeTwoPointers(input));
        System.out.println();
    }

    private static boolean isPalindromeByCleaning(String input) {
        StringBuilder cleaned = new StringBuilder();

        for (int i = 0; i < input.length(); i++) {
            char current = input.charAt(i);
            if (Character.isLetterOrDigit(current)) {
                cleaned.append(Character.toLowerCase(current));
            }
        }

        return cleaned.toString().contentEquals(cleaned.reverse());
    }

    private static boolean isPalindromeTwoPointers(String input) {
        int left = 0;
        int right = input.length() - 1;

        while (left < right) {
            while (left < right && !Character.isLetterOrDigit(input.charAt(left))) {
                left++;
            }
            while (left < right && !Character.isLetterOrDigit(input.charAt(right))) {
                right--;
            }

            char leftChar = Character.toLowerCase(input.charAt(left));
            char rightChar = Character.toLowerCase(input.charAt(right));
            if (leftChar != rightChar) {
                return false;
            }

            left++;
            right--;
        }

        return true;
    }

    private static void runStockProfitDemo() {
        int[] prices = {7, 1, 5, 3, 6, 4};

        System.out.println("Best Time to Buy and Sell Stock");
        System.out.println("Solution 1 - Brute force: " + maxProfitBruteForce(prices));
        System.out.println("Solution 2 - One pass:    " + maxProfitOnePass(prices));
        System.out.println();
    }

    private static int maxProfitBruteForce(int[] prices) {
        int best = 0;

        for (int buy = 0; buy < prices.length; buy++) {
            for (int sell = buy + 1; sell < prices.length; sell++) {
                best = Math.max(best, prices[sell] - prices[buy]);
            }
        }

        return best;
    }

    private static int maxProfitOnePass(int[] prices) {
        int minPrice = Integer.MAX_VALUE;
        int best = 0;

        for (int price : prices) {
            minPrice = Math.min(minPrice, price);
            best = Math.max(best, price - minPrice);
        }

        return best;
    }

    private static void runAnagramDemo() {
        String s = "anagram";
        String t = "nagaram";

        System.out.println("Valid Anagram");
        System.out.println("Solution 1 - Sort:       " + isAnagramBySorting(s, t));
        System.out.println("Solution 2 - Frequency:  " + isAnagramByFrequency(s, t));
        System.out.println();
    }

    private static boolean isAnagramBySorting(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }

        char[] first = s.toCharArray();
        char[] second = t.toCharArray();
        Arrays.sort(first);
        Arrays.sort(second);

        return Arrays.equals(first, second);
    }

    private static boolean isAnagramByFrequency(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }

        int[] counts = new int[26];
        for (int i = 0; i < s.length(); i++) {
            counts[s.charAt(i) - 'a']++;
            counts[t.charAt(i) - 'a']--;
        }

        for (int count : counts) {
            if (count != 0) {
                return false;
            }
        }

        return true;
    }

    private static void runValidParenthesesDemo() {
        String input = "({[]})";

        System.out.println("Valid Parentheses");
        System.out.println("Solution 1 - Open stack:     " + isValidParenthesesByOpenStack(input));
        System.out.println("Solution 2 - Expected stack: " + isValidParenthesesByExpectedStack(input));
        System.out.println();
    }

    private static boolean isValidParenthesesByOpenStack(String input) {
        ArrayDeque<Character> stack = new ArrayDeque<>();

        for (int i = 0; i < input.length(); i++) {
            char current = input.charAt(i);
            if (current == '(' || current == '[' || current == '{') {
                stack.push(current);
                continue;
            }

            if (stack.isEmpty()) {
                return false;
            }

            char open = stack.pop();
            if ((current == ')' && open != '(')
                    || (current == ']' && open != '[')
                    || (current == '}' && open != '{')) {
                return false;
            }
        }

        return stack.isEmpty();
    }

    private static boolean isValidParenthesesByExpectedStack(String input) {
        ArrayDeque<Character> expectedClosers = new ArrayDeque<>();

        for (int i = 0; i < input.length(); i++) {
            char current = input.charAt(i);
            if (current == '(') {
                expectedClosers.push(')');
            } else if (current == '[') {
                expectedClosers.push(']');
            } else if (current == '{') {
                expectedClosers.push('}');
            } else if (expectedClosers.isEmpty() || expectedClosers.pop() != current) {
                return false;
            }
        }

        return expectedClosers.isEmpty();
    }

    private static void runBinarySearchDemo() {
        int[] nums = {-1, 0, 3, 5, 9, 12};
        int target = 9;

        System.out.println("Binary Search");
        System.out.println("Solution 1 - Iterative: " + binarySearchIterative(nums, target));
        System.out.println("Solution 2 - Recursive: " + binarySearchRecursive(nums, target));
        System.out.println();
    }

    private static int binarySearchIterative(int[] nums, int target) {
        int left = 0;
        int right = nums.length - 1;

        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] == target) {
                return mid;
            }
            if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return -1;
    }

    private static int binarySearchRecursive(int[] nums, int target) {
        return binarySearchRecursive(nums, target, 0, nums.length - 1);
    }

    private static int binarySearchRecursive(int[] nums, int target, int left, int right) {
        if (left > right) {
            return -1;
        }

        int mid = left + (right - left) / 2;
        if (nums[mid] == target) {
            return mid;
        }
        if (nums[mid] < target) {
            return binarySearchRecursive(nums, target, mid + 1, right);
        }
        return binarySearchRecursive(nums, target, left, mid - 1);
    }

    private static void runLinkedListDemo() {
        ListNode iterativeHead = buildList(1, 2, 3, 4, 5);
        ListNode recursiveHead = buildList(1, 2, 3, 4, 5);

        System.out.println("Reverse Linked List");
        System.out.println("Solution 1 - Iterative: " + listToString(reverseListIterative(iterativeHead)));
        System.out.println("Solution 2 - Recursive: " + listToString(reverseListRecursive(recursiveHead)));
        System.out.println();
    }

    private static ListNode reverseListIterative(ListNode head) {
        ListNode prev = null;
        ListNode curr = head;

        while (curr != null) {
            ListNode next = curr.next;
            curr.next = prev;
            prev = curr;
            curr = next;
        }

        return prev;
    }

    private static ListNode reverseListRecursive(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }

        ListNode newHead = reverseListRecursive(head.next);
        head.next.next = head;
        head.next = null;
        return newHead;
    }

    private static ListNode buildList(int... values) {
        ListNode dummy = new ListNode(0);
        ListNode tail = dummy;

        for (int value : values) {
            tail.next = new ListNode(value);
            tail = tail.next;
        }

        return dummy.next;
    }

    private static String listToString(ListNode head) {
        StringBuilder result = new StringBuilder();
        ListNode curr = head;

        while (curr != null) {
            if (result.length() > 0) {
                result.append(" -> ");
            }
            result.append(curr.value);
            curr = curr.next;
        }

        return result.toString();
    }

    private static void runTreeDepthDemo() {
        TreeNode root = new TreeNode(
                3,
                new TreeNode(9),
                new TreeNode(20, new TreeNode(15), new TreeNode(7))
        );

        System.out.println("Maximum Depth of Binary Tree");
        System.out.println("Solution 1 - DFS recursion: " + maxDepthDfs(root));
        System.out.println("Solution 2 - BFS levels:    " + maxDepthBfs(root));
        System.out.println();
    }

    private static int maxDepthDfs(TreeNode root) {
        if (root == null) {
            return 0;
        }
        return 1 + Math.max(maxDepthDfs(root.left), maxDepthDfs(root.right));
    }

    private static int maxDepthBfs(TreeNode root) {
        if (root == null) {
            return 0;
        }

        Queue<TreeNode> queue = new ArrayDeque<>();
        queue.add(root);
        int depth = 0;

        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            depth++;

            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queue.remove();
                if (node.left != null) {
                    queue.add(node.left);
                }
                if (node.right != null) {
                    queue.add(node.right);
                }
            }
        }

        return depth;
    }

    private static void runIslandDemo() {
        char[][] grid = {
                {'1', '1', '1', '1', '0'},
                {'1', '1', '0', '1', '0'},
                {'1', '1', '0', '0', '0'},
                {'0', '0', '0', '0', '0'}
        };

        System.out.println("Number of Islands");
        System.out.println("Solution 1 - DFS: " + numIslandsDfs(copyGrid(grid)));
        System.out.println("Solution 2 - BFS: " + numIslandsBfs(copyGrid(grid)));
        System.out.println();
    }

    private static int numIslandsDfs(char[][] grid) {
        int islands = 0;

        for (int row = 0; row < grid.length; row++) {
            for (int col = 0; col < grid[row].length; col++) {
                if (grid[row][col] == '1') {
                    islands++;
                    sinkIslandDfs(grid, row, col);
                }
            }
        }

        return islands;
    }

    private static void sinkIslandDfs(char[][] grid, int row, int col) {
        if (row < 0 || row >= grid.length || col < 0 || col >= grid[row].length) {
            return;
        }
        if (grid[row][col] != '1') {
            return;
        }

        grid[row][col] = '0';
        sinkIslandDfs(grid, row + 1, col);
        sinkIslandDfs(grid, row - 1, col);
        sinkIslandDfs(grid, row, col + 1);
        sinkIslandDfs(grid, row, col - 1);
    }

    private static int numIslandsBfs(char[][] grid) {
        int islands = 0;
        int[][] directions = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};

        for (int row = 0; row < grid.length; row++) {
            for (int col = 0; col < grid[row].length; col++) {
                if (grid[row][col] != '1') {
                    continue;
                }

                islands++;
                Queue<int[]> queue = new ArrayDeque<>();
                queue.add(new int[] {row, col});
                grid[row][col] = '0';

                while (!queue.isEmpty()) {
                    int[] cell = queue.remove();
                    for (int[] direction : directions) {
                        int nextRow = cell[0] + direction[0];
                        int nextCol = cell[1] + direction[1];

                        if (nextRow < 0 || nextRow >= grid.length || nextCol < 0 || nextCol >= grid[nextRow].length) {
                            continue;
                        }
                        if (grid[nextRow][nextCol] != '1') {
                            continue;
                        }

                        grid[nextRow][nextCol] = '0';
                        queue.add(new int[] {nextRow, nextCol});
                    }
                }
            }
        }

        return islands;
    }

    private static char[][] copyGrid(char[][] grid) {
        char[][] copy = new char[grid.length][];
        for (int row = 0; row < grid.length; row++) {
            copy[row] = Arrays.copyOf(grid[row], grid[row].length);
        }
        return copy;
    }

    private static void runKthLargestDemo() {
        int[] nums = {3, 2, 1, 5, 6, 4};
        int k = 2;

        System.out.println("Kth Largest Element");
        System.out.println("Solution 1 - Sorting:  " + findKthLargestBySorting(nums, k));
        System.out.println("Solution 2 - Min heap: " + findKthLargestByMinHeap(nums, k));
        System.out.println();
    }

    private static int findKthLargestBySorting(int[] nums, int k) {
        int[] copy = Arrays.copyOf(nums, nums.length);
        Arrays.sort(copy);
        return copy[copy.length - k];
    }

    private static int findKthLargestByMinHeap(int[] nums, int k) {
        PriorityQueue<Integer> minHeap = new PriorityQueue<>();

        for (int num : nums) {
            minHeap.add(num);
            if (minHeap.size() > k) {
                minHeap.remove();
            }
        }

        return minHeap.peek();
    }

    private static void runClimbingStairsDemo() {
        int n = 5;

        System.out.println("Climbing Stairs");
        System.out.println("Solution 1 - Memoization: " + climbStairsMemo(n));
        System.out.println("Solution 2 - Tabulation:  " + climbStairsTabulation(n));
        System.out.println();
    }

    private static int climbStairsMemo(int n) {
        int[] memo = new int[n + 1];
        return climbStairsMemo(n, memo);
    }

    private static int climbStairsMemo(int n, int[] memo) {
        if (n <= 1) {
            return 1;
        }
        if (memo[n] != 0) {
            return memo[n];
        }

        memo[n] = climbStairsMemo(n - 1, memo) + climbStairsMemo(n - 2, memo);
        return memo[n];
    }

    private static int climbStairsTabulation(int n) {
        if (n <= 1) {
            return 1;
        }

        int[] dp = new int[n + 1];
        dp[0] = 1;
        dp[1] = 1;

        for (int step = 2; step <= n; step++) {
            dp[step] = dp[step - 1] + dp[step - 2];
        }

        return dp[n];
    }

    private static final class ListNode {
        private final int value;
        private ListNode next;

        private ListNode(int value) {
            this.value = value;
        }
    }

    private static final class TreeNode {
        private final int value;
        private final TreeNode left;
        private final TreeNode right;

        private TreeNode(int value) {
            this(value, null, null);
        }

        private TreeNode(int value, TreeNode left, TreeNode right) {
            this.value = value;
            this.left = left;
            this.right = right;
        }
    }
}
