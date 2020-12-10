package property

// Int represents a typed and observable property.
type Int struct {
	Property
}

// Set updates the value and notifies each registered observer.
func (s *Int) Set(v int) {
	s.Property.Set(v)
}

// Get returns the current value.
func (s *Int) Get() int {
	if s.Property.Get() == nil {
		return 0
	}

	return s.Property.Get().(int)
}

// Bind reads the current value from dst into this value. However, every subsequent observed change of the
// property is written into dst.
func (s *Int) Bind(dst *int) Handle {
	h := s.Property.Observe(func(old, new interface{}) {
		*dst = new.(int)
	})

	s.Set(*dst)

	return h
}

// Observe registered a typed observer.
func (s *Int) Observe(onDidSet func(old, new int)) Handle {
	return s.Property.Observe(func(old, new interface{}) {
		if old == nil {
			old = 0
		}

		if old != new {
			onDidSet(old.(int), new.(int))
		}
	})
}
