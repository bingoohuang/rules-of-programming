package rule2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func SumVector(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func Reduce(initialValue int, reduceFunction func(int, int) int, values []int) int {
	reducedValue := initialValue
	for _, value := range values {
		reducedValue = reduceFunction(reducedValue, value)
	}
	return reducedValue
}

func sum(value, otherValue int) int {
	return value + otherValue
}

func TestReduce(t *testing.T) {
	assert.Equal(t, 15, Reduce(0, sum, []int{1, 2, 3, 4, 5}))
}

type Character struct {
	priority int
	index    int
}

var allCharacters []*Character

func NewCharacter(priority int) *Character {
	c := &Character{
		priority: priority,
		index:    0,
	}

	index := 0
	for ; index < len(allCharacters); index++ {
		if priority <= allCharacters[index].priority {
			break
		}
	}

	allCharacters = append(allCharacters, c)

	for ; index < len(allCharacters); index++ {
		allCharacters[index].index = index
	}

	return c
}

// Close 当字符被销毁时清除其索引
func (c *Character) Close() error {
	return nil
}

// SetPriority 如果字符的优先级发生更改，则最小化向前和向后移动该字符
func (c *Character) SetPriority(priority int) {

}
