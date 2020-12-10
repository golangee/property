package property

// Float64 represents a typed and observable property.
type Float64 struct {
	Property
}

// NewFloat64 allocates a new property and sets a default value.
func NewFloat64(value float64) *Float64 {
	b := &Float64{}
	b.Set(value)

	return b
}

// Set updates the value and notifies each registered observer.
func (s *Float64) Set(v float64) {
	s.Property.Set(v)
}

// Get returns the current value.
func (s *Float64) Get() float64 {
	if s.Property.Get() == nil {
		return 0
	}

	return s.Property.Get().(float64)
}

// Bind reads the current value from dst into this value. However, every subsequent observed change of the
// property is written into dst.
func (s *Float64) Bind(dst *float64) Handle {
	h := s.Property.Observe(func(old, new interface{}) {
		*dst = new.(float64)
	})

	s.Set(*dst)

	return h
}

// Observe registered a typed observer.
func (s *Float64) Observe(onDidSet func(old, new float64)) Handle {
	return s.Property.Observe(func(old, new interface{}) {
		if old == nil {
			old = 0
		}

		if old != new {
			onDidSet(old.(float64), new.(float64))
		}
	})
}
