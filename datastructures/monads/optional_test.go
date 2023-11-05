package monads

import "testing"

func TestOptional(t *testing.T) {
	// Test NewOptional
	optional := NewOptional(5)
	if !optional.IsPresent() {
		t.Error("NewOptional failed")
	}

	sevenOptional := NewOptional(7)

	// Test Get
	_, err := sevenOptional.Get()
	if err != nil {
		t.Error("Get failed")
	}

	// Test IfPresent
	sevenOptional.IfPresent(func(i int) {
		if i != 7 {
			t.Error("IfPresent failed")
		}
	})

	emptyOptional := NewEmptyOptional[int]()
	// Test OrElse
	if emptyOptional.OrElse(7) != 7 {
		t.Error("OrElse failed")
	}

	// Test MustGet
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustGet failed")
		}
	}()
	emptyOptional.MustGet()
}
