package embedded

type Foo struct {
	Bar //埋め込みフィールド
}

type Bar struct {
	Baz int
}
