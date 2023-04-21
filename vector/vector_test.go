package vector_test

import (
	"testing"

	"github.com/eatmoreapple/vi/vector"
)

func TestVector_Push(t *testing.T) {
	v := []int{1, 2, 3}
	vec := vector.New(&v)
	vec.Push(4)
	if vec.Len() != 4 {
		t.Errorf("expected %d, got %d", 4, vec.Len())
	}
	if v[3] != 4 {
		t.Errorf("expected %d, got %d", 4, v[3])
	}
}

func TestVector_Pop(t *testing.T) {
	v := []int{1, 2, 3}
	vec := vector.New(&v)
	t1, ok := vec.Pop()
	if !ok {
		t.Errorf("expected %t, got %t", true, ok)
	}
	if t1 != 3 {
		t.Errorf("expected %d, got %d", 3, t1)
	}
	if vec.Len() != 2 {
		t.Errorf("expected %d, got %d", 2, vec.Len())
	}
}

func TestVector_Delete(t *testing.T) {
	v := []int{1, 2, 3}
	vec := vector.New(&v)
	vec.Delete(1)
	if vec.Len() != 2 {
		t.Errorf("expected %d, got %d", 2, vec.Len())
	}
	if v[1] != 3 {
		t.Errorf("expected %d, got %d", 3, v[1])
	}
}

func TestVector_IsEmpty(t *testing.T) {
	v := []int{1, 2, 3}
	vec := vector.New(&v)
	if vec.IsEmpty() {
		t.Errorf("expected %t, got %t", false, vec.IsEmpty())
	}
	vec.Clear()
	if !vec.IsEmpty() {
		t.Errorf("expected %t, got %t", true, vec.IsEmpty())
	}
}

func TestVector_Len(t *testing.T) {
	v := []int{1, 2, 3}
	vec := vector.New(&v)
	if vec.Len() != 3 {
		t.Errorf("expected %d, got %d", 3, vec.Len())
	}
}

func TestVector_Remove(t *testing.T) {
	v := []int{1, 2, 3}
	vec := vector.New(&v)
	vec.Remove(2)
	if vec.Len() != 2 {
		t.Errorf("expected %d, got %d", 2, vec.Len())
	}
	if v[1] != 3 {
		t.Errorf("expected %d, got %d", 3, v[1])
	}
}

func TestVector_Foreach(t *testing.T) {
	v := []int{1, 2, 3}
	vec := vector.New(&v)
	var index int
	vec.Foreach(func(i int) {
		if v[index] != i {
			t.Errorf("expected %d, got %d", i+1, v[index])
		}
		index++
	})
}

func TestVector_Reduce(t *testing.T) {
	v := []int{1, 2, 3}
	vec := vector.New(&v)
	result := vec.Reduce(func(i int, i2 int) int {
		return i + i2
	})
	if result != 6 {
		t.Errorf("expected %d, got %d", 6, result)
	}
}

func TestVector_Map(t *testing.T) {
	v := []int{1, 2, 3}
	vec := vector.New(&v)
	result := vec.Map(func(i int) int {
		return i + 1
	})
	if i, _ := result.At(0); i != 2 {
		t.Errorf("expected %d, got %d", 2, i)
	}
}

func TestVector_Filter(t *testing.T) {
	v := []int{1, 2, 3}
	vec := vector.New(&v)
	result := vec.Filter(func(i int) bool {
		return i > 1
	})
	if i, _ := result.At(0); i != 2 {
		t.Errorf("expected %d, got %d", 2, i)
	}
}

func TestNew(t *testing.T) {
	vec := vector.New[int](nil)
	vec.Push(1)
	if vec.Len() != 1 {
		t.Errorf("expected %d, got %d", 1, vec.Len())
	}
	if a, _ := vec.At(0); a != 1 {
		t.Errorf("expected %d, got %d", 1, a)
	}
}
