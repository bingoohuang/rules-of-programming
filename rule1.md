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

想象一下，你试图计算有多少种方法可以爬梯子，你可以一次爬上 1 级、2 级或者最多 3 级台阶。如果梯子有 2 级，有 2 种方法可以爬上去——
11, 2。

同样地，有 4 种方法可以爬上 3 级梯子——111、21、3。4 级梯子可以有 7 方式攀登，5 级梯子有 13 种方式，以此类推。

你可以编写简单的代码来递归计算这个问题：

```go
package rule1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// CountStepWays 计算爬 stepCount 级梯子有多少种方法
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

基本思路是，攀登梯子的任何一种方式都必须在它下面的三个台阶之一上到达顶层。将攀登到每个台阶的方式相加就可以得到攀登到顶层的方式。
然后，只需找出基本情况即可。前面的代码允许负数步数作为基本情况，以使递归变得简单。

不幸的是，这种解决方案并不可行。只有对于较小的 stepCount 值是可行的，`CountStepWays(20) `所需时间将是 `CountStepWays(19)`
所需时间的两倍。
计算机速度很快，但这样的指数增长速度最终会超过计算机速度。在我的测试中，一旦 `stepCount` 超过二十，示例代码的速度开始变慢。

如果你需要计算更长的梯子的攀登方式数量，那么这个代码就太简单了。核心问题是 `CountStepWays` 的所有中间结果都是一遍又一遍地重新计算，
这导致运行时间呈指数级增长。解决这个问题的标准方法是备忘录技术——保存计算出的中间值并复用它们，就像在以下示例中一样：

```go
package rule1

// CountStepWaysV2 计算爬 stepCount 级梯子有多少种方法
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

// CountStepWaysV3 使用动态规划计算爬 rungCount 级梯子有多少种方法
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

原始的递归版本的 `CountStepWays` 存在长梯子时出现问题。简单的代码可以很好地处理短梯子，但对于长梯子却会遇到指数级的性能壁垒。后续版本在稍微增加一些复杂性的代价下避免了指数级壁垒，但不久后又遇到了不同的问题。

如果我运行之前的代码以计算 `CountStepWaysV3(36)`，我会得到正确的答案 `2082876103`。不过，调用 `CountStepWaysV3(37)`
却返回了-463960867。这显然不对！

这是因为我使用的 go 版本使用了 int32 (带符号32位 ) 类型的值，而计算 `CountStepWaysV3(37)`
导致可用位数溢出。有 `3831006429`种方法可以爬上一个
37 阶的梯子，这个数字太大了，无法容纳在 int32 中。

所以也许代码还是太简单了。我们可以合理地期望 `CountStepWays` 可以适用于所有梯子的长度，对吧？

然后我们就构造了一个可以容纳超级大数字的对象，`Ordinal`，然后用它来作为备忘录 mem 的类型，有效防止溢出，

```go
// 为了实现 Orinal 对象 很多行代码，这里省略
// blabla ...
```

所以，看起来问题已经通过引入 Ordinal 得到了解决。它允许精确计算更长梯子的答案。虽然需要添加几百行代码来实现 `Ordinal`
，这对于一个只有 14 行代码的函数来说似乎有点过多，但这是必要的，以确保问题得到正确解决。

**很可能不是**。如果没有一个简单的解决方案，那么在接受复杂的解决方案之前，应该审查问题。**您需要解决的问题是否真正需要解决**
？或者，您是否对问题进行了不必要的假设，这些假设使您的解决方案变得复杂？

在这种情况下，如果您确实在计算实际梯子的台阶模式，那么您可以假设有一个最大梯子长度。如果最大的梯子长度是15级，那么本节中的任何一种解决方案都是完全足够的，即使是最初展示的纯递归示例也是如此。在其中的一个解决方案中添加一个assert（声明），注明函数内置限制并宣布胜利：

```go
package rule1

import (
	"errors"
)

var ErrOverFlowRungCount = errors.New("max rungs is 36")

// CountStepWaysV4 使用动态规划计算爬 rungCount 级梯子有多少种方法
// 如果 rungCount 超过 36， 返回 ErrOverFlowRungCount 错误
func CountStepWaysV4(rungCount int32) (int32, error) {
	// NOTE (chris) can't represent the pattern count in an int
	// once we get past 36 rungs...”
	if rungCount > 36 {
		return 0, ErrOverFlowRungCount
	}

	stepWaysCounts := []int32{0, 0, 1}

	for rungIndex := int32(0); rungIndex < rungCount; rungIndex++ {
		stepWaysCounts = append(stepWaysCounts,
			stepWaysCounts[rungIndex+0]+
				stepWaysCounts[rungIndex+1]+
				stepWaysCounts[rungIndex+2])
	}

	return stepWaysCounts[len(stepWaysCounts)-1], nil
}

```

或者，如果需要支持非常长的梯子，比如处理风力涡轮机的检测梯子，那么是否只计算步数的近似值就足够了呢？很可能是，如果是这样的话，将整数替换为浮点数就很容易，我甚至不打算展示代码。

请注意，任何东西都会溢出。对于一个问题解决极限边界情况会导致过于复杂的解决方案。不要被限制于解决问题的最严格定义。相对于针对问题更广泛的定义而言，将精力集中于实际需要解决的部分，用简单的解决方案更为明智。如果无法简化解决方案，请试着简化问题。

## 简单算法

有时候，选择错误的算法会增加代码的复杂性。毕竟，有很多方法可以解决任何特定的问题，有些方法比其他方法更复杂。简单的算法会导致简单的代码。问题是简单的算法并不总是显而易见！

假设你要编写代码来对一副牌进行排序。一种显而易见的方法是模拟日常的洗牌手法——将牌堆分成两组，然后将它们交错叠在一起，让每侧的卡牌有近似相等的机会进入重新组合的牌堆下一张牌。重复此过程直至完成洗牌。

代码可能如下所示：

```go
package rule1

import (
	"math/rand"
	"time"
)

// Card 表示扑克牌
type Card int

// Shuffle 对若干扑克牌洗牌
func Shuffle(cards []Card) []Card {
	shuffledCards := cards
	for i := 0; i < 7; i++ {
		shuffledCards = shuffleOnce(shuffledCards)
	}

	return shuffledCards
}

func shuffleOnce(cards []Card) (shuffledCards []Card) {
	ran := rand.New(rand.NewSource(time.Now().UnixMilli()))

	splitIndex := len(cards) / 2
	leftIndex := 0
	rightIndex := splitIndex

	for {
		if leftIndex >= splitIndex {
			shuffledCards = append(shuffledCards, cards[rightIndex:]...)
			break
		} else if rightIndex >= len(cards) {
			shuffledCards = append(shuffledCards, cards[leftIndex:splitIndex]...)
			break
		} else if ran.Intn(2) == 1 {
			shuffledCards = append(shuffledCards, cards[rightIndex])
			rightIndex++
		} else {
			shuffledCards = append(shuffledCards, cards[leftIndex])
			leftIndex++
		}
	}

	return shuffledCards
}

```

这个模拟洗牌的算法是有效的，我在这里编写的代码是该算法的相当简单的实现。你需要花一点能量确保所有的索引检查都是正确的，但并不太难。

但是，还有更简单的算法来洗一副牌。例如，您可以逐张构建一副洗好的牌。在每次迭代中，取一张新牌并将其与该迭代中的随机一张牌交换。实际上，您可以在原地进行操作：

```go
package rule1

import (
	"math/rand"
	"time"
)

// Card 表示扑克牌
type Card int

// ShuffleV2 对若干扑克牌洗牌
func ShuffleV2(cards []Card) {
	ran := rand.New(rand.NewSource(time.Now().UnixMilli()))
	for i := len(cards) - 1; i > 0; i-- {
		j := ran.Intn(i)
		cards[i], cards[j] = cards[j], cards[i]
	}
}

```

根据先前介绍的简洁度指标，这个版本更加卓越。编写时间更短。阅读更容易。代码更少。解释更容易。它更简单、更好——不是因为代码，而是因为更好的算法选择。

## 不要丢失重点

简单的代码易于阅读——最简单的代码可以像读书一样从上到下直接阅读。不过程序并不是书籍。如果代码的流程不简单，很容易出现难以理解的代码。当代码变得拐弯抹角，让你从一个地方跳到另一个地方跟随执行流程，那么阅读起来就会更加困难。

过于强调在一个地方精确地表达每一个想法可能会导致代码变得拐弯抹角。

> 译者注：由于使用 Go 语言简洁性的原因，原文的 C++ 的进一步重构代码，在 Go 语言中已经显得毫无必要了。

减少代码中的重复量是有用的！但是必须认识到，消除重复的代价很高，对于小量的代码和简单的想法，最好只保留重复的副本。这样代码将更容易编写和阅读。

## 大道至简

本书中剩余的规则中，很多都会回到简单这个主题，保持代码尽可能简单，但又不要过于简单。

从本质上讲，编程是处理复杂性的斗争。添加新功能通常意味着让代码变得更加复杂——随着代码变得越来越复杂，工作变得越来越困难，进展也越来越缓慢。最终，你会到达一个事件地平线，在那里，任何试图前进——修复漏洞或添加功能——都会带来和解决问题同等数量的问题。进一步的进展实际上是不可能的。

最终，会是复杂性杀死你的项目。

这意味着有效的编程就是拖延不可避免的结果。在添加功能和修复漏洞时尽可能少地添加复杂性。寻找消除复杂性的机会，或者架构设计，使得新功能不会给系统整体复杂性增加太多。在团队合作方面，尽可能创造简单性。

如果你认真对待，你可以无限期地推迟不可避免的结果。我在25年前写下了《激昂亢奋》的第一行代码，代码库从那以后不断发展。没有尽头——我们的代码比25年前复杂得多，但我们已经能够控制这种复杂性，我们仍然能够取得有效的进展。

我们能够管理复杂性，你也可以。保持警觉，记住复杂性是终极敌人，你将做得很好。

