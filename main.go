package main

import (
	"fmt"
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
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}
	rndElem := rand.New(rand.NewSource(time.Now().UnixNano()))
	slc := make([]int, size)
	for i := range slc {
		slc[i] = rndElem.Int()
	}
	return slc
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	var wg sync.WaitGroup
	chSize := len(data) / CHUNKS
	maxVal := make([]int, CHUNKS)
	for i := range CHUNKS {
		start := i * chSize
		fin := start + chSize
		if i == CHUNKS-1 {
			fin = len(data)
		}
		wg.Add(1)
		go func(i, start, fin int) {
			defer wg.Done()
			chunk := data[start:fin]
			maxVal[i] = maximum(chunk)
		}(i, start, fin)
	}
	wg.Wait()
	return maximum(maxVal)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
