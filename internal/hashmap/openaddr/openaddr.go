package openaddr

type HashMap struct {
	array []string
}

func New() *HashMap {
	return &HashMap{
		array: make([]string, 10),
	}
}

func (h *HashMap) Insert(key, value string) {
}

func (h *HashMap) Get(key string) (string, error) {
	return "", nil
}

func (h *HashMap) Delete(key string) {
}
