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

func CountSetBitsV2(value int) int {
	value = ((value & 0xaaaaaaaa) >> 1) + (value & 0x55555555)
	value = ((value & 0xcccccccc) >> 2) + (value & 0x33333333)
	value = ((value & 0xf0f0f0f0) >> 4) + (value & 0x0f0f0f0f)
	value = ((value & 0xff00ff00) >> 8) + (value & 0x00ff00ff)
	value = ((value & 0xffff0000) >> 16) + (value & 0x000ffff)

	return value
}

func TestCountSetBitsV2(t *testing.T) {
	assert.Equal(t, 2, CountSetBitsV2(0b0001_0100))
	assert.Equal(t, 2, CountSetBitsV2(0b0000_1010))
	assert.Equal(t, 3, CountSetBitsV2(0b0000_1110))
}

func CountSetBitsV3(value int) int {
	count := 0

	for bit := 0; bit < 32; bit++ {
		if (value & (1 << bit)) > 0 {
			count++
		}
	}

	return count
}

func TestCountSetBitsV3(t *testing.T) {
	assert.Equal(t, 2, CountSetBitsV3(0b0001_0100))
	assert.Equal(t, 2, CountSetBitsV3(0b0000_1010))
	assert.Equal(t, 3, CountSetBitsV3(0b0000_1110))
}

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

// CountStepWaysV2 计算爬 stepCount 级楼梯有多少种方法
func CountStepWaysV2(stepCount int) int {
	// 备忘录技术————保存计算出的中间值并复用它们
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

func TestCountStepWaysV2(t *testing.T) {
	assert.Equal(t, 2, CountStepWaysV2(2))
	assert.Equal(t, 4, CountStepWaysV2(3))
	assert.Equal(t, 7, CountStepWaysV2(4))
	assert.Equal(t, 13, CountStepWaysV2(5))
	assert.Equal(t, 66012, CountStepWaysV2(19))
	assert.Equal(t, 3831006429, CountStepWaysV2(37))
}

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

func TestCountStepWaysV3(t *testing.T) {
	assert.Equal(t, int32(2), CountStepWaysV3(2))
	assert.Equal(t, int32(4), CountStepWaysV3(3))
	assert.Equal(t, int32(7), CountStepWaysV3(4))
	assert.Equal(t, int32(13), CountStepWaysV3(5))
	assert.Equal(t, int32(66012), CountStepWaysV3(19))
	assert.Equal(t, int32(2082876103), CountStepWaysV3(36))
	assert.Equal(t, int32(-463960867), CountStepWaysV3(37))
}
