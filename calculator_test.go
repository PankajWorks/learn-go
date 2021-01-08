package hello_test

import (
	"testing"

	"example.com/hello"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := hello.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}

}

func TestSubstract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := hello.Substract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
