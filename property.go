package property

import "sync"

// DidSet is observed, after a new value has been made available. A call to Property.Get will yield the new value.
// Note, that in concurrency, Property.Get may already return an even newer value.
type DidSet func(old, new interface{})

// Handle represents a manual lifecycle manager to a registered Property observer callback.
type Handle struct {
	parent *Property
	idx    int
}

// Release detaches the handle from its parent, so that the registered function can
// be garbage collected, as long as there are no other references. A call is idempotent.
func (h Handle) Release() {
	if h.parent == nil {
		return
	}

	h.parent.lock.Lock()
	defer h.parent.lock.Unlock()

	h.parent.observers[h.idx] = nil
}

// Property is an observable container which can notify all registered callbacks
// when calling Invalidate or updating the contained value. The usage is concurrency and recursive safe.
type Property struct {
	value     interface{}
	observers []DidSet   // we only append, invalid observers will be set to nil
	lock      sync.Mutex // we ever need this very short, no deadlocks possible
}

// Set updates the contained value atomically and notifies all registered observers.
func (p *Property) Set(v interface{}) {
	p.lock.Lock()
	old := p.value
	p.value = v
	p.lock.Unlock()

	p.notify(old, v)
}

// Get returns the value.
func (p *Property) Get() interface{} {
	p.lock.Lock()
	defer p.lock.Unlock()

	return p.value
}

// Observe registers the given function to be called the next time Invalidate is called.
func (p *Property) Observe(f DidSet) Handle {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.observers = append(p.observers, f)

	return Handle{parent: p, idx: len(p.observers) - 1}
}

// Attach connects the given invalidate functions with this Property. If the value of the property has actually changed
// (or is not comparable), function is called.
func (p *Property) Attach(invalidate func()) Handle {
	return p.Observe(func(old, new interface{}) {
		if !equals(old, new) && invalidate != nil {
			invalidate()
		}
	})
}

// notify will invoke all observers which have not been released yet.
// Observers are free to release their Handle or to register new observers.
func (p *Property) notify(old, new interface{}) {
	p.lock.Lock()
	length := len(p.observers)
	p.lock.Unlock()

	for i := 0; i < length; i++ {
		p.lock.Lock()
		observer := p.observers[i]
		p.lock.Unlock()

		if observer != nil {
			observer(old, new)
		}
	}
}

// equals just compares the two interfaces. If not comparable, also false is returned.
func equals(a, b interface{}) bool {
	defer func() {
		_ = recover()
	}()

	return a == b
}
