package equal

import (
	"bytes"
	"testing"
)

func Test(t *testing.T) {
	one, oneAgain, two := 1, 1, 2
	
	type mystring string

	ch1, ch2 := make(chan int), make(chan int)
	var ch1ro <-chan int = ch1

	var iface1, iface1Again, iface2 interface{} = &one, &oneAgain, &two

	for _, test := range []struct{
		x, y interface{}
		want bool
	} {
		{1, 1, true},
		{1, 2, false},  //different values
		{1, 1.0, false},  // different types
		{"foo", "foo", true},
		{"foo", "bar", false},
		{mystring("foo"), "foo", false}, //different types
		{[]string{"foo"}, []string{"foo"}, true},
		{[]string{"foo"}, []string{"bar"}, false},
		{[]string{}, []string(nil), true},
		// maps
		{
			map[string][]int{"foo": {1, 2, 3}},
			map[string][]int{"foo": {1, 2, 3}},
			true,
		},
		{
			map[string][]int{"foo": {1, 2, 3}},
			map[string][]int{"foo": {1, 2, 3, 4}},
			false,
		},
		{
			map[string][]int{},
			map[string][]int(nil),
			true,
		},
		// pointers
		{&one, &one, true},
		{&one, &two, false},
		{&one, &oneAgain, true},
		{new(bytes.Buffer), new(bytes.Buffer), true},
		// functions
		{(func())(nil), (func())(nil), true},
		{(func())(nil), func() {}, false},
		{func() {}, func() {}, false},
		// arrays
		{[...]int{1, 2, 3}, [...]int{1, 2, 3}, true},
		{[...]int{1, 2, 3}, [...]int{1, 2, 4}, false},
		// channels
		{ch1, ch1, true},
		{ch1, ch2, false},
		{ch1ro, ch1, false}, // NOTE: not equal
		// interfaces
		{&iface1, &iface1, true},
		{&iface1, &iface2, false},
		{&iface1Again, &iface1, true},
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) == %t", test.x, test.y, !test.want)
		}
	}
}