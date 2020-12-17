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

// String represents a typed and observable property.
type String struct {
	Property
}

// NewString allocates a new property and sets a default value.
func NewString(value string) *String {
	b := &String{}
	b.Set(value)

	return b
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
