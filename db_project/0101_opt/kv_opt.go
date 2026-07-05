package db0101_opt

import "sync"

// Set only inserts if the key is not already present.
// Del only removes if the key is present.
// Both report whether the call actually changed the store's state.
type KV struct {
	mu  sync.RWMutex
	mem map[string][]byte
}

func (kv *KV) Open() error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.mem = make(map[string][]byte)
	return nil
}

func (kv *KV) Close() error { return nil }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	v, found := kv.mem[string(key)]
	if !found {
		return nil, false, nil
	}
	val = make([]byte, len(v))
	copy(val, v)
	return val, true, nil
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if _, exists := kv.mem[string(key)]; exists {
		return false, nil
	}

	v := make([]byte, len(val))
	copy(v, val)
	kv.mem[string(key)] = v
	return true, nil
}

func (kv *KV) Del(key []byte) (deleted bool, err error) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if _, exists := kv.mem[string(key)]; !exists {
		return false, nil
	}
	delete(kv.mem, string(key))
	return true, nil
}
