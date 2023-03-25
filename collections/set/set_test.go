package set

import "testing"

func TestNew(t *testing.T) {
	expectedSet := Set[int]{
		elements: []int{1, 2, 3, 4},
		size:     4,
	}

	actualSet := New[int](1, 1, 2, 3, 4, 4)

	if expectedSet.size != actualSet.size {
		t.Errorf("Expected %v but got %v", expectedSet.size, actualSet.size)
	}
}
