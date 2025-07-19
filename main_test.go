package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	newSize1 = 99_999_999
)

func getTestData(size int) []int {
	var data []int
	for i := 0; i <= size; i++ {
		data = append(data, i)
	}
	return data
}

func TestGenerateRandomElementsWhenOk(t *testing.T) {
	data := generateRandomElements(newSize1)
	assert.Equal(t, newSize1, len(data))

}

func TestGenerateRandomElementsNegative(t *testing.T) {
	if newSize1 == 0 {
		data := generateRandomElements(newSize1)
		assert.Equal(t, newSize1, len(data))
	}
	if newSize1 == -100 {
		data := generateRandomElements(newSize1)
		assert.Equal(t, 0, len(data))
	}
}

func TestMaximumWhenOk(t *testing.T) {
	data := getTestData(newSize1)
	max := maximum(data)
	assert.Equal(t, newSize1, max)
}

func TestMaximumNegative(t *testing.T) {
	data := make([]int, newSize1)
	expectMax := maximum(data)
	assert.Equal(t, 0, expectMax)

	data = []int{100}
	expectMax = maximum(data)
	assert.Equal(t, data[0], expectMax)
}

// в техническом задании для func maxChunks() не просят писать тесты
