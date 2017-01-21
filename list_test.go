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
