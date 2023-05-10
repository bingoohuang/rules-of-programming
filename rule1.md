# 规则1： 简单原则 —— 越简单越好，但不可简陋

编程不易。

我猜你已经意识到了这一点。任何拿起并阅读标题为《编程规则》的书的人都很可能：

- 至少能写一点点代码。
- 对编程不易感到沮丧。

编程不易是有很多原因的，但也也有很多策略可以尝试。本书选取了常见的方式来摆脱麻烦和规则以避免这些错误，这些方式都来自于我的多年经验，其中包括我自己犯过的错误和他人的错误处理经验。

这些规则有一个总体的模式，大部分规则都有一个共同的主题。它最好由阿尔伯特·爱因斯坦的一句话来概括：

> 越简单越好，但不要简陋。

老爱的意思是说，最好的物理理论是最简单的那一个，完全描述了所有可观察现象。

这个想法可以应用到编程上，解决任何问题的最简单实现方法是符合该问题所有要求的方法。最好的代码也是最简单的代码。

假设你正在编写代码来计算整数中二进制非零位数的数量。有很多方法可以实现这个功能。您可能会使用位操作技巧逐位将位清零，计算清零的位数：

```go
package rule1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func CountSetBits(value int) int {
	count := 0

	for value != 0 {
		count++
		value &= value - 1
	}

	return count
}

func TestCountSetBits(t *testing.T) {
	assert.Equal(t, 2, CountSetBits(0b0001_0100))
	assert.Equal(t, 2, CountSetBits(0b0000_1010))
	assert.Equal(t, 3, CountSetBits(0b0000_1110))
}
```

或者你可以选择无循环实现，使用位移和屏蔽来并行计数位：

```go
package rule1

func CountSetBitsV2(value int) int {
	value = ((value & 0xaaaaaaaa) >> 1) + (value & 0x55555555)
	value = ((value & 0xcccccccc) >> 2) + (value & 0x33333333)
	value = ((value & 0xf0f0f0f0) >> 4) + (value & 0x0f0f0f0f)
	value = ((value & 0xff00ff00) >> 8) + (value & 0x00ff00ff)
	value = ((value & 0xffff0000) >> 16) + (value & 0x000ffff)

	return value
}

```

或者你可能只是编写最明显的代码：

```go
package rule1

func CountSetBitsV3(value int) int {
	count := 0

	for bit := 0; bit < 32; bit++ {
		if (value & (1 << bit)) > 0 {
			count++
		}
	}

	return count
}
```

前两个答案很聪明…我不是在夸奖他们。你没法秒懂前两个示例中实际发生了什么——它们每个都有一小部分“呃，这是个什么鬼？”的代码。需要经过一些思考，你才可以理解情，当然理解了这个技巧也蛮有趣的，但总得费点脑子。

在展示代码之前，我告诉你函数的作用，函数命名的重要性让人如数家珍。如果你不知道这些代码计算了设置的位数，读懂前两个代码示例会更费力气。

而对于最后一个答案，很显然，它正在计算设置的位数。它是尽可能简单的，而没有更简单的了，这使得它比前两个答案更好。

## 测量简单性

有很多方法来思考什么是简单的代码。

您可以决定根据团队中的其他人理解代码的难度来衡量简单性。如果随意挑选一个小伙伴可以阅读一些代码并不费力地理解了代码，那么这段代码就可能符合了简单性原则了。

或者您可以根据创建代码的容易程度来衡量简单性，而不仅仅是敲代码所需的时间，还包括使代码完全可用和无漏洞所需的时间。复杂的代码需要一段时间才能正确运行；简单的代码更容易顺利完成。

当然，这两个衡量标准有很大的重叠。易于编写的代码通常也易于阅读。还有其他有效的复杂度衡量标准：

1. 编写了多少代码

   代码往往较短，尽管在一行代码中塞入许多复杂性是可能的。

2. 引入了多少新概念

   简单的代码往往建立在团队中每个人都知道的概念上，不会引入解决问题的新思维方式或新术语。

3. 解释需要多长时间

   代码易于解释 —— 在代码审核中，它足够明显，评审人员可以轻松通过。复杂的代码需要解释。

一个在某个标准下看起来简单的代码，在其他标准下也会看起来简单。你只需要选择哪个标准对你的工作提供了最清晰的关注点，但我建议从易于创建和易于理解开始。如果你专注于快速创建易于阅读的代码，那么你正在创建简单的代码。

## ……但不要简陋

代码越简单越好，但它仍然需要解决它打算解决的问题。

想象一下，你试图计算有多少种方法可以爬楼梯，你可以一次爬上 1 级、2 级或者最多 3 级台阶。如果楼梯有 2 级，有 2 种方法可以爬上去——
11, 2。

同样地，有 4 种方法可以爬上 3 级楼梯——111、21、3。4 级楼梯可以有 7 方式攀登，5 级楼梯有 13 种方式，以此类推。

你可以编写简单的代码来递归计算这个问题：

```go
package rule1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// CountStepWays 计算爬 stepCount 级楼梯有多少种方法
func CountStepWays(stepCount int) int {
	switch {
	case stepCount < 0:
		return 0
	case stepCount == 0:
		return 1
	}

	return CountStepWays(stepCount-3) +
		CountStepWays(stepCount-2) +
		CountStepWays(stepCount-1)
}

func TestCountStepWays(t *testing.T) {
	assert.Equal(t, 2, CountStepWays(2))
	assert.Equal(t, 4, CountStepWays(3))
	assert.Equal(t, 7, CountStepWays(4))
	assert.Equal(t, 13, CountStepWays(5))
	assert.Equal(t, 66012, CountStepWays(19))
}

```

基本思路是，攀登楼梯的任何一种方式都必须在它下面的三个台阶之一上到达顶层。将攀登到每个台阶的方式相加就可以得到攀登到顶层的方式。
然后，只需找出基本情况即可。前面的代码允许负数步数作为基本情况，以使递归变得简单。

不幸的是，这种解决方案并不可行。只有对于较小的 stepCount 值是可行的，`CountStepWays(20) `所需时间将是 `CountStepWays(19)`
所需时间的两倍。
计算机速度很快，但这样的指数增长速度最终会超过计算机速度。在我的测试中，一旦 `stepCount` 超过二十，示例代码的速度开始变慢。

如果你需要计算更长的楼梯的攀登方式数量，那么这个代码就太简单了。核心问题是 `CountStepWays` 的所有中间结果都是一遍又一遍地重新计算，
这导致运行时间呈指数级增长。解决这个问题的标准方法是备忘录技术——保存计算出的中间值并复用它们，就像在以下示例中一样：

```go
package rule1

// CountStepWaysV2 计算爬 stepCount 级楼梯有多少种方法
func CountStepWaysV2(stepCount int) int {
	// 备忘录技术——保存计算出的中间值并复用它们
	memo := map[int]int{}
	return countStepWaysV2(memo, stepCount)

}

func countStepWaysV2(memo map[int]int, stepCount int) int {
	switch {
	case stepCount < 0:
		return 0
	case stepCount == 0:
		return 1
	}

	if ways, ok := memo[stepCount]; ok {
		return ways
	}

	ways := countStepWaysV2(memo, stepCount-3) +
		countStepWaysV2(memo, stepCount-2) +
		countStepWaysV2(memo, stepCount-1)
	memo[stepCount] = ways
	return ways
}

```

使用备忘录技术后，每个值只计算一次并插入哈希表中。随后的调用在哈希表中以大致恒定的时间找到已经计算出的值，指数增长现象消失了。
备忘录技术使代码稍微复杂了一些，但它不会遇到性能瓶颈。

你还可以使用动态规划，将一些概念上的复杂性换成更好的代码简洁性：

```go
package rule1

// CountStepWaysV3 使用动态规划计算爬 rungCount 级楼梯有多少种方法
func CountStepWaysV3(rungCount int32) int32 {
	stepWaysCounts := []int32{0, 0, 1}

	for rungIndex := int32(0); rungIndex < rungCount; rungIndex++ {
		stepWaysCounts = append(stepWaysCounts,
			stepWaysCounts[rungIndex+0]+
				stepWaysCounts[rungIndex+1]+
				stepWaysCounts[rungIndex+2])
	}

	return stepWaysCounts[len(stepWaysCounts)-1]
}

```

这种方法也运行得足够快，而且比备忘录递归版本还要简单。

## 有时候，简化问题而不是解决问题会更好。

原始的递归版本的 `CountStepWays` 存在长楼梯时出现问题。简单的代码可以很好地处理短楼梯，但对于长楼梯却会遇到指数级的性能壁垒。后续版本在稍微增加一些复杂性的代价下避免了指数级壁垒，但不久后又遇到了不同的问题。

如果我运行之前的代码以计算 `CountStepWaysV3(36)`，我会得到正确的答案 `2082876103`。不过，调用 `CountStepWaysV3(37)`
却返回了-463960867。这显然不对！

这是因为我使用的 go 版本将 int32 (带符号32位 ) 值，而计算 `CountStepWaysV3(37)` 导致可用位数溢出。有 `3831006429`种方法可以爬上一个
37 阶的楼梯，这个数字太大了，无法容纳在 int32 中。

所以也许代码还是太简单了。我们可以合理地期望 `CountStepWays` 可以适用于所有楼梯的长度，对吧？

（未完待续）
