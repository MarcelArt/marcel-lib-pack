package arrays_test

import (
	"testing"

	"github.com/MarcelArt/marcel-lib-pack/pkg/arrays"
	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	t.Run("should sum array correctly", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		total := arrays.Reduce(s, 0, func(total, currentValue int) int {
			return total + currentValue
		})

		assert.Equal(t, 15, total)
	})
}

func TestMap(t *testing.T) {
	t.Run("should multiply element by two and return an array with same size", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		result := arrays.Map(s, func(currentValue int) int {
			return currentValue * 2
		})

		assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
	})
}

func TestFilter(t *testing.T) {
	t.Run("should be able to filter even numbers", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		result := arrays.Filter(s, func(currentValue int) bool {
			return currentValue%2 == 0
		})

		assert.Equal(t, []int{2, 4}, result)
	})

	t.Run("should return an empty array if nothing is found", func(t *testing.T) {
		s := []int{1, 3, 5}
		result := arrays.Filter(s, func(currentValue int) bool {
			return currentValue == 2
		})

		assert.Equal(t, []int{}, result)
	})

	t.Run("should be able to return more than one element", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		result := arrays.Filter(s, func(currentValue int) bool {
			return currentValue%2 == 0
		})

		assert.Len(t, result, 2)
	})
}

func TestFind(t *testing.T) {
	t.Run("should find the first even number", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		result := arrays.Find(s, func(currentValue int) bool {
			return currentValue%2 == 0
		})

		assert.Equal(t, 2, *result)
	})

	t.Run("shoud return nil if nothing found", func(t *testing.T) {
		s := []int{1, 3, 5}
		result := arrays.Find(s, func(currentValue int) bool {
			return currentValue == 2
		})

		assert.Nil(t, result)
	})
}
