package mlmmjconfig

import (
	"testing"
)

func TestName(t *testing.T) {
	var l = MlmmjList{"test/list1"}
	if l.Name() != "list1" {
		t.Fatal("List should be named list1")
	}
}

func TestClosedList(t *testing.T) {
	var l = MlmmjList{"test/test3/list1"}
	if l.ClosedList() {
		t.Fatal("List should be opened")
	}
	l = MlmmjList{"test/test3/list2"}
	if !l.ClosedList() {
		t.Fatal("List should be closed")
	}
	l = MlmmjList{"test/test3/list-rw"}
	if l.ClosedList() {
		t.Fatal("List should be opened")
	}
	err := l.SetClosedList(true)
	if err != nil {
		t.Fatal("Problem with setting flag closedlist")
	}
	if !l.ClosedList() {
		t.Fatal("List should be closed")
	}
	err = l.SetClosedList(false)
	if err != nil {
		t.Fatal("Problem with unsetting flag closedlist")
	}
	if l.ClosedList() {
		t.Fatal("List should be opened")
	}
}
