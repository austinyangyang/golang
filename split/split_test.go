package split

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {

	got := Split("a:b:c:d", ":")

	want := []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%v, got:%v", got, want)
	}
	fmt.Printf("got: %v\n want: %v\n", got, want)
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

func TestSplit1(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	// tests := []test{
	// 	 {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
	// 	{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
	// 	{input: "abcd", sep: "bc", want: []string{"a", "d"}},
	// 	{input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	// }

	tests := map[string]test{
		"case1": {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"case2": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"case3": {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"case4": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	for name, tc := range tests {

		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got: %#v", tc.want, got)
			}

		})
		// got := strings.Split(tc.input, tc.sep)
		// if !reflect.DeepEqual(got, tc.want) {
		// 	t.Errorf("excepted:%#v, got: %#v", tc.want, got)
		// }

	}
}

func BenchmarkSplit(b *testing.B) {
	time.Sleep(5* time.Second)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}
