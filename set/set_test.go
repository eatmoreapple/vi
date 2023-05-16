package set

import (
	"testing"
)

func TestSet_Add(t *testing.T) {
	s := New[int]()
	s.Add(1)
	if !s.Contains(1) {
		t.Error("set should contain 1")
	}
}

func TestSet_Clear(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Clear()
	if s.Contains(1) {
		t.Error("set should not contain 1")
	}
}

func TestSet_Contains(t *testing.T) {
	s := New[int]()
	s.Add(1)
	if !s.Contains(1) {
		t.Error("set should contain 1")
	}
	if s.Contains(2) {
		t.Error("set should not contain 2")
	}
}

func TestSet_Difference(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)
	s3 := s.Difference(s2)
	if s3.Len() != 1 {
		t.Error("difference should have length 1")
	}
	if !s3.Contains(1) {
		t.Error("difference should contain 1")
	}
}

func TestSet_Intersection(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)
	s3 := s.Intersection(s2)
	if s3.Len() != 1 {
		t.Error("intersection should have length 1")
	}
	if !s3.Contains(2) {
		t.Error("intersection should contain 2")
	}
}

func TestSet_Len(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	if s.Len() != 2 {
		t.Error("set should have length 2")
	}
}

func TestSet_Remove(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Remove(1)
	if s.Contains(1) {
		t.Error("set should not contain 1")
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)
	s3 := s.SymmetricDifference(s2)
	if s3.Len() != 2 {
		t.Error("symmetric difference should have length 2")
	}
	if !s3.Contains(1) {
		t.Error("symmetric difference should contain 1")
	}
	if !s3.Contains(3) {
		t.Error("symmetric difference should contain 3")
	}
}

func TestSet_ToSlice(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	ts := s.ToSlice()
	if len(ts) != 2 {
		t.Error("slice should have length 2")
	}
}

func TestSet_Union(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)
	s3 := s.Union(s2)
	if s3.Len() != 3 {
		t.Error("union should have length 3")
	}
	if !s3.Contains(1) {
		t.Error("union should contain 1")
	}
	if !s3.Contains(2) {
		t.Error("union should contain 2")
	}
	if !s3.Contains(3) {
		t.Error("union should contain 3")
	}
}

func TestSet_Equal(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(1)
	if !s.Equal(s2) {
		t.Error("sets should be equal")
	}
}

func TestSet_Equal2(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)
	if s.Equal(s2) {
		t.Error("sets should not be equal")
	}
}

func TestSet_Equal3(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s2 := New[int]()
	s2.Add(2)
	if s.Equal(s2) {
		t.Error("sets should not be equal")
	}
}
