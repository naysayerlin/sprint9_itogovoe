package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenOfRandEl(t *testing.T) {
	tests := []struct {
		name  string
		size  int
		expct int
	}{
		{"Positive size", 100, 100},
		{"Zero size", 0, 0},
		{"Large size", 10000, 10000},
	}
	for _, x := range tests {
		t.Run(x.name, func(t *testing.T) {
			reslt := generateRandomElements(x.size)
			assert.Equal(t, x.expct, len(reslt), "Длинна должна быть равна ожидаемой")
			if x.size > 0 {
				assert.NotEmpty(t, reslt, "Слайс не должен быть пустым")
				assert.IsType(t, []int{}, reslt, "Слайс не должен быть пустым")
			} else {
				assert.Empty(t, reslt, "Если размер слайса неверный, должен быть возвращен пустой слайс")
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name  string
		data  []int
		expct int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{35}, 35},
		{"Multiple elements", []int{2, 6, 3, 7, 2}, 7},
		{"All same elements", []int{5, 5, 5, 5}, 5},
	}
	for _, x := range tests {
		t.Run(x.name, func(t *testing.T) {
			reslt := maximum(x.data)
			assert.Equal(t, x.expct, reslt, "Максимальное значение должно быть равно ожидаемому")
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{35}, 35},
		{"Multiple elements", []int{4, 8, 1, 7, 2, 6, 9, 8, 5}, 9},
		{"All same elements", []int{6, 6, 6, 6, 6, 6, 6, 6}, 6},
		{"Less than chunks", []int{1, 2, 3}, 3},
	}
	for _, x := range tests {
		t.Run(x.name, func(t *testing.T) {
			reslt := maxChunks(x.data)
			assert.Equal(t, x.expected, reslt, "Поиск максимума должен выдавать корректный результат")
		})
	}
}
