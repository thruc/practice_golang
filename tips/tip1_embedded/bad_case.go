package embedded

import "sync"

type BadInMem struct {
	sync.Mutex
	m map[string]int
}

func NewBadInMem() *BadInMem {
	return &BadInMem{m: make(map[string]int)}
}

func (i *BadInMem) Get(key string) (int, bool) {
	i.Lock()
	defer i.Unlock()
	v, contains := i.m[key]
	return v, contains
}

type GoodInMem struct {
	mu sync.Mutex // 埋め込むことでmethodにaccessできないようにしている
	m  map[string]int
}

func NewGoodInMem() *BadInMem {
	return &BadInMem{m: make(map[string]int)}
}

func (i *GoodInMem) Get(key string) (int, bool) {
	i.mu.Lock()
	defer i.mu.Unlock()
	v, contains := i.m[key]
	return v, contains
}

func useInOutside() {
	i := NewBadInMem()
	i.Get("foo")
	i.Lock() // 外部のクライアントからアクセスできてしまう
}
