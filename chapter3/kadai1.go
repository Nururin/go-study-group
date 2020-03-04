package chapter3

// 課題1
// 以下のstructにgetterとsetterを実装してください。
// Getterの関数名ID, Name
// Setterの関数名SetID, SetName
type Kadai1 struct {
	id   int
	name string
}

func (k Kadai1) ID() int {
	return k.id
}
func (k *Kadai1) SetID(x int) {
	k.id = x
}

func (k Kadai1) Name() string {
	return k.name
}
func (k *Kadai1) SetName(x string) {
	k.name = x
}
