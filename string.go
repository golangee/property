package property

// String represents a typed and observable property.
type String struct {
	Property
}

// Set updates the value and notifies each registered observer.
func (s *String) Set(v string) *String {
	s.Property.Set(v)

	return s
}

// Get returns the current value.
func (s *String) Get() string {
	if s.Property.Get() == nil {
		return ""
	}

	return s.Property.Get().(string)
}

// Bind reads the current value from dst into this value. However, every subsequent observed change of the
// property is written into dst.
func (s *String) Bind(dst *string) Handle {
	h := s.Property.Observe(func(old, new interface{}) {
		*dst = new.(string)
	})

	s.Set(*dst)

	return h
}

// Observe registered a typed observer.
func (s *String) Observe(onDidSet func(old, new string)) Handle {
	return s.Property.Observe(func(old, new interface{}) {
		if old == nil {
			old = ""
		}

		if old != new {
			onDidSet(old.(string), new.(string))
		}
	})
}
