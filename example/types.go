package main

type object struct {
	id  int
	val string
}

func (o object) ID() int {
	return o.id
}

func (o object) Value() string {
	return o.val
}

func (o object) String() {
	println(o.val)
}
