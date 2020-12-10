// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package property

// Bool represents a typed and observable property.
type Bool struct {
	Property
}

// NewBool allocates a new property and sets a default value.
func NewBool(value bool) *Bool {
	b := &Bool{}
	b.Set(value)

	return b
}

// Set updates the value and notifies each registered observer.
func (s *Bool) Set(v bool) {
	s.Property.Set(v)
}

// Get returns the current value.
func (s *Bool) Get() bool {
	if s.Property.Get() == nil {
		return false
	}

	return s.Property.Get().(bool)
}

// Bind reads the current value from dst into this value. However, every subsequent observed change of the property
// is written into dst.
func (s *Bool) Bind(dst *bool) Handle {
	h := s.Property.Observe(func(old, new interface{}) {
		*dst = new.(bool)
	})

	s.Set(*dst)

	return h
}

// Toggle inverts the current state.
func (s *Bool) Toggle() {
	newV := !s.Get()
	s.Set(newV)
}

// Observe registered a typed observer.
func (s *Bool) Observe(onDidSet func(old, new bool)) Handle {
	return s.Property.Observe(func(old, new interface{}) {
		if old == nil {
			old = false
		}

		if old != new {
			onDidSet(old.(bool), new.(bool))
		}
	})
}
