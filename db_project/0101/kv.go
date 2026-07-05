package db0101

import (
	"fmt"
)

type KV struct {
	mem map[string][]byte
}

func (kv *KV) Open() error {
	kv.mem = map[string][]byte{} // empty
	return nil
}

func (kv *KV) Close() error { return nil }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error) {
	if val = kv.mem[string(key)]; val != nil {
		ok = true
	} else {
		ok = false
		// err = errors.New("DB: Could not find key " + string(key) + " in memory")
	}
	fmt.Println(string(key), val, ok, err)
	return val, ok, err
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error) {
	if kv.mem[string(key)] != nil {
		updated = false
		// err = errors.New("DB: Key " + string(key) + " already exists!")
	} else {
		kv.mem[string(key)] = val
		updated = true
	}
	fmt.Println(string(key), updated, err)
	return updated, err
}

func (kv *KV) Del(key []byte) (deleted bool, err error) {
	if kv.mem[string(key)] != nil {
		kv.mem[string(key)] = nil
		deleted = true
	} else {
		deleted = false
		// err = errors.New("DB: Key " + string(key) + " does not exist!")
	}
	fmt.Println(string(key), deleted, err)
	return deleted, err
}

// QzBQWVJJOUhU https://trialofcode.org/
