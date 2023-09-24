package monads

import "testing"

func TestOptional(t *testing.T) {
	t.Run("NewOptional", func(t *testing.T) {
		o := NewOptional("foo")
		if !o.IsPresent() {
			t.Errorf("expected IsPresent() to be true")
		}
	})

	t.Run("NewEmptyOptional", func(t *testing.T) {
		o := NewEmptyOptional[string]()
		if o.IsPresent() {
			t.Errorf("expected IsPresent() to be false")
		}
	})

	t.Run("SetValue", func(t *testing.T) {
		o := NewEmptyOptional[string]()
		o.SetValue("foo")
		if !o.IsPresent() {
			t.Errorf("expected IsPresent() to be true")
		}
	})

	t.Run("Get", func(t *testing.T) {
		o := NewOptional("foo")
		val, err := o.Get()
		if err != nil {
			t.Errorf("expected err to be nil")
		}
		if val != "foo" {
			t.Errorf("expected val to be foo")
		}
	})

	t.Run("Get (empty)", func(t *testing.T) {
		o := NewEmptyOptional[string]()
		_, err := o.Get()
		if err == nil {
			t.Errorf("expected err to not be nil")
		}
	})

	t.Run("MustGet", func(t *testing.T) {
		o := NewOptional("foo")
		val := o.MustGet()
		if val != "foo" {
			t.Errorf("expected val to be foo")
		}
	})

	t.Run("MustGet (empty)", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		o := NewEmptyOptional[string]()
		o.MustGet()
	})

	t.Run("IfPresent", func(t *testing.T) {
		o := NewOptional("foo")
		o.IfPresent(func(val string) {
			if val != "foo" {
				t.Errorf("expected val to be foo")
			}
		})
	})

	t.Run("IfPresent (empty)", func(t *testing.T) {
		o := NewEmptyOptional[string]()
		o.IfPresent(func(val string) {
			t.Errorf("expected function to not be called")
		})
	})
}
