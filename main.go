package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	var data []int
	if size == 0 {
		log.Println("size must not be zero")
		return data
	}
	if size < 0 {
		log.Println("the size must not be less than zero")
		return data
	}

	src := rand.NewSource(time.Now().Unix())
	for i := 0; i < size; i++ {
		randomNumber := src.Int63()
		data = append(data, int(randomNumber))
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	var max int
	if len(data) == 0 {
		log.Println("slice must not be zero")
		return max
	}
	if len(data) == 1 {
		log.Println("to find the maximum number you need at least two numbers")
		return max
	}

	for _, num := range data {
		if num >= max {
			max = num
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		log.Println("slice must not be zero")
		return 0
	}
	if len(data) == 1 {
		log.Println("to find the maximum number you need at least two numbers")
		return 0
	}
	if len(data)%CHUNKS != 0 {
		log.Printf("the slice size must be a multiple of %d\n", CHUNKS)
		return 0
	}
	var mu sync.Mutex
	var wg sync.WaitGroup
	var maxData []int
	sizeChunk := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		newData := data[i*sizeChunk : (i+1)*sizeChunk]
		wg.Add(1)
		go func() {

			defer wg.Done()
			maxNumber := maximum(newData)
			mu.Lock()
			maxData = append(maxData, maxNumber)
			mu.Unlock()
		}()
	}
	wg.Wait()
	max := maximum(maxData)
	return max
}

func main() {

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
