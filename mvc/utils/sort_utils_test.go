package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	//initialization
	elements := []int{9, 8, 7, 6, 5}

	//execution
	BubbleSort(elements)

	//validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, []int{5, 6, 7, 8, 9}, elements)
}

func TestBubbleSortBestCase(t *testing.T) {
	//initialization
	elements := []int{5, 6, 7, 8, 9}

	//execution
	BubbleSort(elements)

	//validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, []int{5, 6, 7, 8, 9}, elements)
}

func TestBubbleSortNillSlice(t *testing.T) {
	//initialization

	//execution
	BubbleSort(nil)
}
