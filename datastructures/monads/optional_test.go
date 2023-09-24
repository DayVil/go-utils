package monads

import "testing"

func TestOptional(t *testing.T) {
	// Test NewOptional
	optional := NewOptional(5)
	if !optional.IsPresent() {
		t.Error("NewOptional failed")
	}

	// Test NewEmptyOptional
	emptyOptional := NewEmptyOptional[int]()
	if emptyOptional.IsPresent() {
		t.Error("NewEmptyOptional failed")
	}

	// Test SetValue
	optional.SetValue(6)
	val, err := optional.Get()
	if err != nil {
		t.Error("SetValue failed")
	}
	if val != 6 {
		t.Error("SetValue failed")
	}

	// Test Get
	_, err = emptyOptional.Get()
	if err == nil {
		t.Error("Get failed")
	}

	// Test MustGet
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustGet failed")
		}
	}()
	emptyOptional.MustGet()
}
