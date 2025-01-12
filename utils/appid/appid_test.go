package appid

import "testing"

func TestSimple(t *testing.T) {
	if CounterStrike != 10 {
		t.Fatal("Constant for Counter Strike 1.6 broken")
	}
}

func TestString(t *testing.T) {
	if CounterStrike.String() != "Counter Strike 1.6" {
		t.Fatal("Constant for Counter Strike 1.6 broken")
	}
}

func TestStringFalse(t *testing.T) {
	if AppID(1).String() != "1" {
		t.Fatal("Constant broken")
	}
}
