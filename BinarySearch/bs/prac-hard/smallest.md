[译]使用“试错法”解决问题

原文：[Approach the problem using the "trial and error" algorithm](https://leetcode.com/explore/learn/card/binary-search/146/more-practices-ii/1041/discuss/109082/Approach-the-problem-using-the-%22trial-and-error%22-algorithm)


**更新**：如果我们把输入的数组按升序排列，这个问题实际上可以改写成找到排序矩阵中的第 K 小的元素，在 `(i, j)` 位置的矩阵元素由 `matrix[i][j] = nums[j] - nums[i]` 得出，还有 `k = K + n*(n+1)/2` 和 `n = nums.length`。这个问题就可以用我[之前发帖](https://leetcode.com/problems/k-th-smallest-prime-fraction/discuss/115819/Summary-of-solutions-for-problems-%22reducible%22-to-LeetCode-378)中的多种算法（优先队列 PriorityQueue， 二分搜索 BinarySearch， 之字形搜索 ZigzagSearch）来解决。

---

通常我们会避免使用天真的试错算法（[trial and error](https://en.wikipedia.org/wiki/Trial_and_error)）来解决问题，因为它通常会导致糟糕的时间复杂度。但是，有些情况下，这种天真的算法可能胜过其他更复杂的解决方案，而 LeetCode 确实有一些这样的问题（在本文的末尾列出 —— 具有讽刺意味的是，他们大多数都是“难”题）。因此，我觉得应该在此描述应用此算法的一般过程。

试错算法的基本思想实际上非常简单，总结如下：

* 第一步：构建候选解决方案。

* 第二步：验证它是否符合我们的要求。

* 第三步：如果是，则接受解决方案；否则丢弃并重复第一步。

但是，为了使此算法有效工作，需要满足以下两个条件：

* 条件一：我们在`第二步`中一个有效的验证算法；

* 条件二：由所有候选解决方案形成的搜索空间很小，或者如果它很大，我们有有效的方法来遍历（或搜索）这个空间。

第一个条件确保每个验证操作都可以快速完成，而第二个条件限制了需要完成的此类操作的总数。这两个组合将保证我们有一个有效的试错算法（这意味着如果不能满足其中一个，你甚至可能不会考虑这个算法）。

---

现在我们来看看这个问题： `719. Find The K-th Smallest Pair Distance`，看看我们怎么应用试错算法。

**一、 构建候选解决方案**

为了构建候选解决方案，我们首先需要了解所需的解决方案。这个问题描述要求我们输出第 K 个最小的对距离（pair distance），这不过是一个非负整数（因为输入数组 nums 是一个整数数组，对距离的是绝对值）。因此，我们的候选解决方案应该也是非负整数。

**二、 搜索空间由所有的候选解决方案组成**

设 `min` 和 `max` 为输入数组 nums 中的最小和最大值，并且 `d = max - min`，然后任何 nums 中的对距离必须位于 `[0, d]` 的范围。因此，我们所需的解决方案一再此范围内，这意味着搜索空间将为`[0, d]` （超出此范围的任何数字都可以立即排除，无需进一步验证）。

**三、 验证给定的候选解决方案**

这是试错算法的关键部分。因此，给定一个候选整数，我们如何确定它是否是第 K 个最小的对距离呢？

首先，第 K 个最小的对距离究竟意味着什么？根据定义，如果我们计算所有对距离并按照升序排序，那么第 K 个最小的对距离将是索引 `K-1` 处的距离。这实际上是解决这个问题的天真方式（但是不出所料会因为 `MLE` 而被拒绝）。

显然，上述定义不能用于验证，因为它需要显示计算对距离数组。幸运的是，还有另一种方法来定义第 K 个最小的对距离：给定一个整数 `num`，用 `count(num)` 表示不大于 `num` 的对距离数，则第 K 个最小对距离将是满足 `count(num) >= K` 的最小整数。

以下是替代定义。令 `num_k` 为具有索引 `K-1` 的有序对距离数组的第 K 个对距离，如第一定义中所指定的。由于索引 `K-1` 及其之前的所有对距离都不大于 `num_k` ，因此 `count(num_k) >= K` 。现在假设 `num` 是满足 `count(num) >= K` 最小的整数，我们认定 `num` 必须等于 `num_k` ，如下所示：

1. 如果 `num_k < num` ，因为 `count(num) >= K` ，那么 `num`  将不是满足 `count(num) >= K` 的最小整数，这与我们假设相矛盾。
2. 如果 `num_k > num` ，因为 `count(num) >= K` ，根据 `count` 的函数定义，至少有 K 个对距离不大于 `num` ，这意味着至少有 K 个对距离小于 `num_k` 。这意味着 `num_k` 不能是第 K 个对距离，这再次与我们的假设矛盾。

利用第 K 个最小对距离的这种替代定义，我们可以将验证过程转换为一个计数过程，那么我们究竟如何计算呢？

**四、 计算不超过给定整数的对距离数**

正如我提到的，我们不能使用对距离数组，这意味着唯一的选择是输入数组本身。如果其元素之间没有顺序，除了逐个计算和测试每个对距离之外，我们没有更好的方法。这导致 `O(n^2)` 的验证算法，其与上述朴素解决方法一样差。所以我们需要对 nums 强加一些命令，默认是排序。

现在假设 `nums` 按升序排序，我们如何继续计算给定的 `num` ？注意，每个对距离 `d_ij` 的特征在于一对索引 `(i, j)` ，其中 `i < j` ，即 `d_ij <= nums[j] - nums[i]` 。如果我们保持第一个索引 `i` 固定，那么 `d_ij <= num` 等价于 `nums[j] <= nums[i] + num` 。这表明至少我们可以进行二分搜索以找到最小的索引 `j` ，使得每个索引 `i` 的 `nums[j] > nums[i] + num` ，然后索引 `i` 的 count 将是 `j-i-1` ，并且总的来说，我们有一个 O(nlogn) 的验证算法。

事实证明，如果我们使用以下属性，则可以使用经典的双指针技术在线性时间内完成计数：假设有两个起始索引 `i1` 和 `i2` ， `i1 < i2` ， 让 `j1` 和 `j2` 成为最小的索引，使得 `nums[j1] > nums[i1] + num` he `nums[j2] > nums[i2] + num` ，那么 `j2 >= j1` 必定是真的。证明是直截了当的：假设 `j2 < j1` ，因为 `j1` 是满足 `nums[j1] > nums[i1] + num` 的最小的索引， 我们应该有 `nums[j2] <= nums[i1] + num` 。另一方面， `nums[j2] > nums[i2] + num >= nums[i1] + num` 。这两个不等式相互矛盾，从而证实了我们上面的结论。

**五、 如何有效地遍历（或搜索）搜索空间**

到目前为止，我们知道搜索空间，知道如何构建候选解决方案以及如何通过计数来验证它，我们仍然需要最后一块拼图：如何遍历搜索空间。

当然，我们可以通过尝试从 `0` 到 `d` 的每个整数来进行天真的线性遍历，并选择第一个满足 `count(num) >= K`  的整数 `num`  。时间复杂度将是 `O(nd)` 。然而，假设 `d` 可以比 `n` 大得多，则改算法可能比之前提出的朴素 `O(n^2)` 解更差。

这里关键是观察到候选解决方案是按升序自然排序的，因此可能使用二分搜索。另一个事实是 `count` 函数的非递减属性：给定两个整数 `num1` 和 `num2` 满足 `num1 < num2` ，我们有 `count(num1) <= count(num2)` （我将证明过程留给你）。所以二分搜索步骤如下：

1. 设 `[l,r]` 为当前搜索空间，初始化 `l = 0, r = d` 。
2. 如果 `l < r` ，计算中间点 `m = (l + r) / 2` 并评估 `count(m)` 。
3. 如果 `count(m) < K` ，我们丢弃当前搜索空间的左半边，令 `l = m + 1`；否则 `count(m) >= K` 丢弃右半边，令 `r=m` 。

你可能想知道为什么即使 `count(m) == K` ，我们也丢弃搜索空间的右半部分。注意第 K 小的对距离 `num_k` 是满足 `count(num_k) >= K` 的最小整数。当 `count(m) == K` 时，我们知道 `num_k <= K` （但不一定是 `num_k == m`，想想吧！）所以保持右半边没有意义。

**六、 把所有东西放到一起，又称解决方案**

不要被上述分析吓到。一旦理解，最终的解决方案就更容易编写。下面是 Java 版本的试错算法，时间复杂度 `O(nlogd + nlogn)` （别忘了排序），空间复杂度 `O(1)` 。

```Java
public int smallestDistancePair(int[] nums, int k) {
    Arrays.sort(nums);
    
    int n = nums.length;
    int l = 0;
    int r = nums[n - 1] - nums[0];
    
    for (int cnt = 0; l < r; cnt = 0) {
        int m = l + (r - l) / 2;
        
        for (int i = 0, j = 0; i < n; i++) {
            while (j < n && nums[j] <= nums[i] + m) j++;
            cnt += j - i - 1;
        }
        
        if (cnt < k) {
            l = m + 1;
        } else {
            r = m;
        }
    }
    
    return l;
}
```

---

最后是 LeetCode 上可以用试错算法解决的问题列表（欢迎添加新的例子）：

* [786. K-th Smallest Prime Fraction](https://leetcode.com/problems/k-th-smallest-prime-fraction/description/)

* [774	Minimize Max Distance to Gas Station](https://leetcode.com/problems/minimize-max-distance-to-gas-station/description/)

* [719. Find K-th Smallest Pair Distance](https://leetcode.com/problems/find-k-th-smallest-pair-distance/description/)

* [668. Kth Smallest Number in Multiplication Table](https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/description/)

* [644. Maximum Average Subarray II](https://leetcode.com/problems/maximum-average-subarray-ii/description/)

* [378. Kth Smallest Element in a Sorted Matrix](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/description/)

无论如何，这篇文章只是提醒你，如果你所有其他常见解决方案严重受到不良时间或空间性能的影响，则试错算法值得尝试。此外，我们始终建议你在完全致力于此算法前，快速评估搜索空间大小和潜在的验证算法，以及估计算法复杂度。

希望有帮助，编码愉快！
