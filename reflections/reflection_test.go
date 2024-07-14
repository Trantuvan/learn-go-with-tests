package reflections

import (
	"slices"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{{
		"struct with one string field",
		struct {
			Name string
		}{"Chris"},
		[]string{"Chris"},
	}, {
		"struct with two string fields",
		struct {
			Name string
			City string
		}{"Chris", "London"},
		[]string{"Chris", "London"},
	}, {
		"struct with non string field",
		struct {
			Name string
			Age  int
		}{"Chris", 33},
		[]string{"Chris"},
	}}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(s string) {
				got = append(got, s)
			})

			if !slices.Equal(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
