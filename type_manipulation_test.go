package main

import "testing"

func TestAtoi(t *testing.T) {
	testStr := "123"
	expected := 123

	got, err := Atoi(testStr)
	if err != "" || got != expected {
		t.Errorf("Got: %d", got)
		t.Errorf("Expected: %d", expected)
		t.Errorf("TestAtoi Failed!")
	}
}

func TestTransfm(t *testing.T) {
	testSlc := []string{"1", "2", "3"}
	expected := []int{1, 2, 3}
	got := sliceTransfm(testSlc)

	for i := 0; i < len(testSlc); i++ {
		if got[i] != expected[i] {
			t.Errorf("Got %d", got)
			t.Errorf("Expected: %d", expected)
			t.Errorf("TestTransfm Failed!")
		}
	}
}