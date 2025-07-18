package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	newSize1 = 99_999_999
)

func getTestData() []int {
	var data []int
	for i := 0; i <= newSize1; i++ {
		data = append(data, i)
	}
	return data
}

func TestGenerateRandomElementsWhenOk(t *testing.T) {
	data := generateRandomElements(newSize1)
	assert.Equal(t, newSize1, len(data))
	if newSize1 == 0 {
		data := generateRandomElements(newSize1)
		assert.Equal(t, newSize1, len(data))
	}
}

func TestMaximumWhenOk(t *testing.T) {
	data := getTestData()
	max := maximum(data)
	assert.Equal(t, newSize1, max)
}
