package gotoolkit

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := new(Stack)
	stack.Push("Test")

	if stack.Size == 0 {
		t.Errorf("Incorrect result\ngot:  %#d\nwant: %#d", stack.Size, 0)
	}
}

func TestPop(t *testing.T) {
	stack := new(Stack)
	stack.Push("Test")

	item, err := stack.Pop()

	if err != nil {
		t.Errorf("Incorrect result\ngot:  %#s\nwant: %#s", err, nil)
	}

	if item != "Test" {
		t.Errorf("Incorrect result\ngot:  %#s\nwant: %#s", item, "Test")
	}

	if stack.Size != 0 {
		t.Errorf("Incorrect result\ngot:  %#s\nwant: %#s", stack.Size, 0)
	}
}

func TestPopWithEmptyStack(t *testing.T) {
	stack := new(Stack)
	item, err := stack.Pop()
	want := "Unable to pop element, stack is empty"

	if err.Error() != want {
		t.Errorf("Incorrect result\ngot:  %#s\nwant: %#s", err, want)
	}

	if item != nil {
		t.Errorf("Incorrect result\ngot:  %#s\nwant: %#s", item, nil)
	}

	if !stack.IsEmpty() {
		t.Errorf("Incorrect result\ngot:  %t\nwant: %t", stack.IsEmpty(), true)
	}
}
