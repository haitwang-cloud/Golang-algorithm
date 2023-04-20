# Golang-algorithm
## reference

- [https://labuladong.gitee.io/algo/di-ling-zh-bfe1b/](https://labuladong.gitee.io/algo/di-ling-zh-bfe1b/)

## Tree view of the project
```bash
.
├── README.md
└── codes
    ├── offer-30-days-training
    │   ├── Day-01-栈和队列
    │   │   ├── 包含min函数的栈
    │   │   │   └── MinStack.go
    │   │   └── 两个栈实现队列
    │   │       └── queueWithTwoStack.go
    │   ├── Day-02-链表
    │   │   └── 链表
    │   │       ├── 反转链表
    │   │       │   └── reverseList.go
    │   │       ├── 从尾到头打印链表
    │   │       │   └── reversePrint.go
    │   │       └── 复制带随机指针的链表
    │   │           └── copyRandomList.go
    │   ├── Day-03-字符串
    │   │   ├── 替换空格
    │   │   │   └── replaceSpace.go
    │   │   └── 左旋转字符串
    │   │       └── reverseLeftWords.go
    │   ├── Day-04-查找
    │   │   ├── 数组中缺失的数字
    │   │   │   └── missingNumber.go
    │   │   ├── 数组中重复的数字
    │   │   │   └── findRepeatNumber.go
    │   │   └── 在排序数组中查找数字 I
    │   │       └── search.go
    │   ├── Day-05-查找算法
    │   │   ├── 二维数组中的查找
    │   │   │   └── findNumberIn2DArray.go
    │   │   ├── 旋转数组的最小数字
    │   │   │   └── minArray.go
    │   │   └── 第一个只出现一次的字符
    │   │       └── firstUniqChar.go
    │   ├── Day-06-搜索与回溯算法
    │   │   ├── 从上到下打印二叉树
    │   │   │   └── levelOrder.go
    │   │   ├── 从上到下打印二叉树_||
    │   │   │   └── levelOrder.go
    │   │   └── 从上到下打印二叉树_|||
    │   │       └── levelOrder.go
    │   ├── Day-07-搜索与回溯算法（简单）
    │   │   ├──  二叉树的镜像
    │   │   │   └── mirrorTree.go
    │   │   ├── 树的子结构
    │   │   │   └── isSubStructure.go
    │   │   └── 对称的二叉树
    │   │       └── isSymmetric.go
    │   ├── Day-08-动态规划
    │   │   ├── 斐波那契数列
    │   │   │   └── fib.go
    │   │   ├── 股票的最大利润
    │   │   │   └── maxProfit.go
    │   │   └── 青蛙跳台阶问题
    │   │       └── numWays.go
    │   ├── Day-09-动态规划（中等）
    │   │   ├── 礼物的最大价值
    │   │   │   └── maxValue.go
    │   │   └── 连续数组的最大和
    │   │       └── maxSubArray.go
    │   ├── Day-10-动态规划（中等）
    │   │   ├── 把数字翻译成字符串
    │   │   │   └── translateNum.go
    │   │   └── 最长不含重复字符的子字符串
    │   │       └── lengthOfLongestSubstring.go
    │   ├── Day-11-双指针
    │   │   ├── 链表中倒数第k个节点
    │   │   │   └── getKthFromEnd.go
    │   │   └── 删除链表的节点
    │   │       └── deleteNode.go
    │   ├── Day-12-双指针
    │   │   ├──  两个链表的第一个公共节点
    │   │   │   └── getIntersectionNode.go
    │   │   └── 合并两个排序的链表
    │   │       └── mergeTwoLists.go
    │   ├── Day-13-双指针
    │   │   ├──  和为s的两个数字
    │   │   │   └── twoSum.go
    │   │   ├── 翻转单词顺序
    │   │   │   └── reverseWords.go
    │   │   └── 调整数组顺序使奇数位于偶数前面
    │   │       └── exchange.go
    │   ├── Day-14-搜索与回溯算法（中等)
    │   │   ├── 矩阵中的路径
    │   │   │   └── exist.go
    │   │   └── 机器人的运动范围
    │   │       └── movingCount.go
    │   ├── Day-15-搜索与回溯算法（中等）
    │   │   ├── 二叉搜索树的第k大节点
    │   │   │   └── kthLargest.go
    │   │   ├── 二叉搜索树与双向链表
    │   │   │   └── treeToDoublyList.go
    │   │   └── 二叉树中和为某一值的路径
    │   │       └── pathSum.go
    │   ├── Day-16-排序（简单）
    │   │   ├── 扑克牌中的顺子
    │   │   │   └── isStraight.go
    │   │   └── 把数组排成最小的数
    │   │       └── minNumber.go
    │   ├── Day-17-排序（中等）
    │   │   ├── 最小的k个数
    │   │   │   └── getLeastNumbers.go
    │   │   └── 数据流中的中位数
    │   │       └── FindMedian.go
    │   ├── Day-18-搜索与回溯算法（中等）
    │   │   ├── 平衡二叉树
    │   │   │   └── isBalanced.go
    │   │   └── 二叉树的深度
    │   │       └── maxDepth.go
    │   ├── Day-19-搜索与回溯算法（中等）
    │   │   ├── 求1+2+…+n
    │   │   │   └── sumNums.go
    │   │   ├── 二叉树的最近公共祖先
    │   │   │   └── lowestCommonAncestor.go
    │   │   └── 二叉搜索树的最近公共祖先
    │   │       └── lowestCommonAncestor.py
    │   ├── Day-20-分治算法（中等）
    │   │   ├── 重建二叉树
    │   │   │   └── buildTree.go
    │   │   ├── 数值的整数次方
    │   │   │   └── myPow.go
    │   │   └── 二叉搜索树的后序遍历序列
    │   │       └── verifyPostorder.go
    │   ├── Day-21- 位运算（简单）
    │   │   ├──  二进制中1的个数
    │   │   │   └── hammingWeight.go
    │   │   └── 用加减乘除做加法
    │   │       └── add.go
    │   ├── Day-22-位运算（中等）
    │   │   ├── 数组中数字出现的次数
    │   │   │   └── singleNumbers.go
    │   │   └── 数组中数字出现的次数 II
    │   │       └── singleNumber.go
    │   ├── Day-23- 数学（简单）
    │   │   ├── 构建乘积数组
    │   │   │   └── constructArr.go
    │   │   └── 数组中出现次数超过一半的数字
    │   │       └── majorityElement.go
    │   ├── Day-24-数学（中等）
    │   │   ├──  圆圈中最后剩下的数字
    │   │   │   └── lastRemaining.go
    │   │   ├── 和为s的连续正数序列
    │   │   │   └── findContinuousSequence.go
    │   │   └── 剪绳子
    │   │       └── cuttingRope.go
    │   ├── Day-25-模拟（中等）
    │   │   ├── 顺时针打印矩阵
    │   │   │   └── spiralOrder.go
    │   │   └── 栈的压入、弹出序列
    │   │       └── validateStackSequences.go
    │   ├── Day-26-字符串（中等）
    │   │   ├── 表示数值的字符串
    │   │   │   └── isNumber.go
    │   │   └── 把字符串转换成整数
    │   │       └── strToInt.go
    │   ├── Day-27-栈与队列（困难）
    │   │   ├── 队列的最大值
    │   │   │   └── MaxQueue.go
    │   │   └── 滑动窗口的最大值
    │   │       └── maxSlidingWindow.go
    │   ├── Day-28-搜索与回溯算法（困难）
    │   │   ├── 字符串的排列
    │   │   │   └── permutation.go
    │   │   └── 序列化二叉树
    │   │       └── serializ.py
    │   ├── Day-29-动态规划（困难）
    │   │   ├── n个骰子的点数
    │   │   │   └── dicesProbability.go
    │   │   ├── 丑数
    │   │   │   └── nthUglyNumber.go
    │   │   └── 正则表达式匹配
    │   │       └── isMatch.go
    │   ├── Day-30-分治算法（困难）
    │   │   ├── 打印从1到最大的n位数
    │   │   │   └── printNumbers.go
    │   │   └── 数组中的逆序对
    │   │       └── reversePairs.go
    │   ├── Day-31-数学（困难）
    │   │   ├── 剪绳子 II
    │   │   │   └── cuttingRope.go
    │   │   └── 数字序列中某一位的数字
    │   │       └── findNthDigit.go
    │   └── README.MD
    └── routine-training
        ├──  二叉树
        │   ├── 层序遍历
        │   │   ├── levelOrder.go
        │   │   └── levelOrder_test.go
        │   ├── 二叉树的直径
        │   │   ├── diameterOfBinaryTree.go
        │   │   └── diameterOfBinaryTree_test.go
        │   ├── 二叉树的前序遍历
        │   │   ├── preorderTraversal.go
        │   │   └── preorderTraversal_test.go
        │   └── 二叉树的最大深度
        │       ├── maxDepth.go
        │       └── maxDepth_test.go
        ├── README.md
        ├── 排序
        │   ├── 前 K 个高频元素
        │   │   └── topKFrequent.go
        │   ├── 最大数
        │   │   └── largestNumber.go
        │   ├── 荷兰国旗
        │   │   └── sortColors.go
        │   ├── 数组中的第K个最大元素
        │   │   └── Kth_Element.go
        │   ├── 把数组排成最小的数
        │   │   └── minNumber.go
        │   └── 根据字符出现频率排序
        │       └── frequencySort.go
        ├── 数组
        │   ├── 双指针
        │   │   ├── 移动零
        │   │   │   ├── moveZeroes.go
        │   │   │   └── moveZeroes_test.go
        │   │   ├── 移除元素
        │   │   │   ├── removeElement.go
        │   │   │   └── removeElement_test.go
        │   │   ├── 两数之和 II - 输入有序数组
        │   │   │   ├── twoSum.go
        │   │   │   └── twoSum_test.go
        │   │   ├── 反转字符串
        │   │   │   ├── reverseString.go
        │   │   │   └── reverseString_test.go
        │   │   ├── 回文字符串
        │   │   │   ├── validPalindrome.go
        │   │   │   └── validPalindrome_test.go
        │   │   ├── 最长回文子串
        │   │   │   ├── longestPalindrome.go
        │   │   │   └── longestPalindrome_test.go
        │   │   └── 删除有序数组中的重复项
        │   │       ├── removeDuplicates.go
        │   │       └── removeDuplicates_test.go
        │   ├── 旋转数组
        │   │   └── rotate.go
        │   ├── 螺旋矩阵
        │   │   └── generateMatrix..go
        │   ├── 最长公共前缀
        │   │   └── longestCommonPrefix.go
        │   ├── 两个数组的交集
        │   │   └── intersect.go
        │   ├── 二维数组中查找
        │   │   └── findNumberIn2DArray.go
        │   └── 顺时针打印矩阵
        │       └── spiralOrder.go
        ├── 贪心
        │   ├── 分发饼干
        │   │   └── findContentChildren.go
        │   ├── 种植花朵
        │   │   └── canPlaceFlowers.go
        │   ├── 划分字母区间
        │   │   └── partitionLabels.go
        │   ├── 最大子数组和
        │   │   └── maxSubArray.go
        │   ├── 不重叠的区间个数
        │   │   └── eraseOverlapIntervals.go
        │   ├── 根据身高重建队列
        │   │   └── reconstructQueue.go
        │   ├── 买卖股票的最佳时机
        │   │   └── maxProfit.go
        │   ├── 用最少数量的箭引爆气球
        │   │   └── findMinArrowShots.go
        │   └── 修改一个数成为非递减数组
        │       └── checkPossibility.go
        ├── 链表
        │   ├── 双指针
        │   │   ├── 分割链表
        │   │   │   ├── partition.go
        │   │   │   └── partition_test.go
        │   │   ├── 环形链表
        │   │   │   ├── hasCycle.go
        │   │   │   └── hasCycle_test.go
        │   │   ├── 相交链表
        │   │   │   ├── getIntersectionNode.go
        │   │   │   └── getIntersectionNode_test.go
        │   │   ├── 移除链表元素
        │   │   │   ├── removeElements.go
        │   │   │   └── removeElements_test.go
        │   │   ├── 链表中倒数第k个节点
        │   │   │   ├── getKthFromEnd.go
        │   │   │   └── getKthFromEnd_test.go
        │   │   ├── 链表的中间结点
        │   │   │   ├── middleNode.go
        │   │   │   └── middleNode_test.go
        │   │   ├── 合并两个有序链表
        │   │   │   ├── mergeTwoLists.go
        │   │   │   └── mergeTwoLists_test.go
        │   │   ├── 删除链表的倒数第 N 个结点
        │   │   │   ├── removeNthFromEnd.go
        │   │   │   └── removeNthFromEnd_test.go
        │   │   ├── 链表中环的入口节点
        │   │   │   ├── detectCycle.go
        │   │   │   └── detectCycle_test.go
        │   │   └── 删除排序链表中的重复元素
        │   │       ├── deleteDuplicates.go
        │   │       └── deleteDuplicates_test.go
        │   └── 反转链表
        │       ├── reverseList.go
        │       └── reverseList_test.go
        ├── 双指针
        │   ├── 两数之和
        │   │   └── 平方数之和
        │   │       └── judgeSquareSum.go
        │   ├── 原地删除
        │   │   └── 移除元素
        │   │       └── removeElement.go
        │   ├── 判断子序列
        │   │   └── isSubsequence.go
        │   ├── 最长子序列
        │   │   └── findLongestWord.go
        │   ├── 荷兰国旗问题
        │   │   └── sortColors.go
        │   ├── 有序数组的平方
        │   │   └── sortedSquares.go
        │   ├── 合并两个有序数组
        │   │   └── merge.go
        │   └── 反转字符串中的元音字母
        │       └── reverseVowels.go
        ├── 二分搜索
        │   ├──  x 的平方根 
        │   │   └── mySqrt.go
        │   ├── 供暖器
        │   │   └── findRadius.go
        │   ├── 二分查找
        │   │   └── search.go
        │   ├── 狒狒吃香蕉
        │   │   └── minEatingSpeed.go
        │   ├── 搜索插入位置
        │   │   └── searchInsert.go
        │   ├── 旋转数组中的最小值
        │   │   └── minArray.go
        │   └── 在排序数组中查找元素的第一个和最后一个位置
        │       └── searchRange.go
        ├── 求和问题
        │   ├── 三数之和
        │   │   └── threeSum.go
        │   ├── 两数之和
        │   │   └── twoSum.go
        │   └── 两数之和(有序数组)
        │       └── twoSum.go
        ├── 滑动窗口
        │   ├── template.go
        │   ├── 水果成篮
        │   │   └── totalFruit.go
        │   └── 长度最小的子数组
        │       └── minSubArrayLen.go
        └── 排序算法汇总
            ├── readme.md
            ├── sort.webp
            ├── 桶排序
            │   ├── bucketSort.go
            │   └── bucketSort_test.go
            ├── 冒泡排序
            │   ├── bubbleSort.go
            │   ├── bullleSort_test.go
            │   └── readme.md
            ├── 基数排序
            │   ├── radixSort.go
            │   └── radixSort_test.go
            ├── 希尔排序
            │   ├── shellSort.go
            │   └── shellSort_test.go
            ├── 归并排序
            │   ├── mergeSort.go
            │   └── mergeSort_test.go
            ├── 快速排序
            │   ├── quickSort.go
            │   └── quickSort_test.go
            ├── 插入排序
            │   ├── straightInsertionSort.go
            │   └── straightInsertionSort_test.go
            ├── 计数排序
            │   ├── countingSort.go
            │   └── countingSort_test.go
            └── 简单选择排序
                ├── simpleSelect.go
                └── simpleSelect_test.go
```
- [algorithms-by-category](./codes/routine-training/README.md)
- [剑指offer-30-days-training](./codes/offer-30-days-training/README.MD)
- [leetcode-hot-100](./codes/leetcode-hot-100/README.MD)

