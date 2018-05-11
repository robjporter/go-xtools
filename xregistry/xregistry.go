package xregistry

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"sync"
)

type AnyValue interface{}

// Registry is a thread-safe map with strings as keys and anything can be
// stored as a  value.
// Locking is required because underlying map will be read and written to
// by many go routines.
type Registry struct {
	lock  sync.RWMutex
	store map[string]AnyValue
}

// Returns value read from given key and true or false if the key exists
func (r *Registry) Get(key string) (AnyValue, bool) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	val, status := r.store[key]
	return val, status
}

func (r *Registry) Exists(key string) bool {
	r.lock.RLock()
	defer r.lock.RUnlock()
	_, status := r.store[key]
	return status
}

func (r *Registry) Clear() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.store = make(map[string]AnyValue)
}

func (r *Registry) Delete(key string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	delete(r.store, key)
}

func (r *Registry) Len() int {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return len(r.store)
}

func (r *Registry) Copy(oldkey string, newkey string) bool {
	status := false
	if r.Exists(oldkey) {
		tmp, _ := r.Get(oldkey)
		r.Set(newkey, tmp)
		status = true
	}
	return status
}

// Set (Safely) value under given key
func (r *Registry) Set(key string, val AnyValue) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.store[key] = val
}

func New() *Registry {
	return &Registry{store: make(map[string]AnyValue)}
}

func (r *Registry) SetWithMap(m map[string]interface{}) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for k, v := range m {
		r.store[k] = v
	}
}

func (r *Registry) Keys() []string {
	var keys []string
	r.lock.RLock()
	defer r.lock.RUnlock()

	for k := range r.store {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func (r *Registry) String() string {
	var buf []byte
	var keys []string
	r.lock.RLock()
	defer r.lock.RUnlock()

	for k := range r.store {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		value, err := encodeValue(r.store[k])

		if err != nil {
			continue
		}

		key := encodeKey(k)
		buf = append(buf, fmt.Sprintf(" %s=%s", key, value)...)
	}

	// Remove leading space
	buf = buf[1:len(buf)]

	return string(buf)
}

// Sanitize the key for logging purposes
func encodeKey(k string) (key string) {
	// Keys may not have any spaces
	key = strings.Replace(k, " ", "_", -1)

	return key
}

// Encode the value of the map for certain supported types.
func encodeValue(i interface{}) (buf []byte, err error) {
	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.String:
		buf = append(buf, fmt.Sprintf("%q", i)...)
	case reflect.Array, reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Map, reflect.Ptr, reflect.Struct:
		err = fmt.Errorf("Unable to encode %s value: type=%s value=%v",
			v.Kind(), reflect.TypeOf(i), i)
	default:
		buf = append(buf, fmt.Sprintf("%v", i)...)
	}

	return buf, err
}
