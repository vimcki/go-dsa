package openaddr

import (
	"crypto/md5"
	"errors"
)

type entry struct {
	key   string
	value string
}

type HashMap struct {
	mod   int
	array []entry
}

func New() *HashMap {
	mod := 100
	return &HashMap{
		mod:   mod,
		array: make([]entry, mod),
	}
}

func (h *HashMap) Insert(key, value string) {
	hash := hash(key, h.mod)
	current := h.array[hash]
	if current.value != "" {
		h.array[hash] = entry{key, value}
		return
	}
	for i := hash; ; i = (i + 1) % h.mod {
		if h.array[i].value == "" {
			h.array[i] = entry{key, value}
			return
		}
	}
}

func (h *HashMap) Get(key string) (string, error) {
	hash := hash(key, h.mod)
	current := h.array[hash]
	if current.value != "" {
		return current.value, nil
	}
	return "", errors.New("Key not found")
}

func (h *HashMap) Delete(key string) {
	hash := hash(key, h.mod)
	current := h.array[hash]
	for current.key != key {
		hash = (hash + 1) % h.mod
		current = h.array[hash]
	}
	h.array[hash] = entry{}
}

func hash(key string, mod int) int {
	md5 := md5.Sum([]byte(key))
	return int(md5[0]) % mod
}
