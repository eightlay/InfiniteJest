package tests

import (
	"reflect"
	"testing"
)

type A interface{ ID() string }
type B struct{}

func (*B) ID() string {
	return "B"
}

type C struct {
	q1 string
	q2 string
	q3 string
	q4 string
	q5 string
	q6 string
	q7 string
	q8 string
}

func (C) ID() string {
	return "C"
}

func BenchmarkReflect1(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	c := B{}
	b.StartTimer()

	v := reflect.ValueOf(&c).MethodByName("ID").Call([]reflect.Value{})[0].Interface().(string)
	b.StopTimer()
	v += ""
}

func BenchmarkReflect2(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	c := C{"1", "1", "1", "1", "1", "1", "1", "1"}
	b.StartTimer()

	v := c.ID()
	b.StopTimer()
	v += ""
}
