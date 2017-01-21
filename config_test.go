package mlmmjconfig

import (
    "testing"
)

func TestOpen(t *testing.T) {
    var c *MlmmjConfig
    var err error

    c, err = Open("zażółć gęślą jaźń")
    if err == nil {
        t.Fatal("Func Open should return error (file not found)")
    }

    c, err = Open("test/test1")
    if err == nil {
        t.Fatal("Func Open should return error (file is not a directory)")
    }

    c, err = Open("test/test2")
    if err != nil {
        t.Fatal(err)
    }

    if c == nil {
        t.Fatal("Func Open should return non-nil value")
    }
}

func TestLists(t *testing.T) {
    var c *MlmmjConfig
    var err error
    var list []MlmmjList

    c, err = Open("test/test2")
    if err != nil {
        t.Fatal(err)
    }
    if c == nil {
        t.Fatal("Func Open should return non-nil value")
    }

    list, err = c.Lists()
    if err != nil {
        t.Fatal(err)
    }
    if len(list) != 0 {
        t.Fatal("Func Lists should return empty slice")
    }
    c, err = Open("test/test3")
    if err != nil {
        t.Fatal(err)
    }
    if c == nil {
        t.Fatal("Func Open should return non-nil value")
    }

    list, err = c.Lists()
    if err != nil {
        t.Fatal(err)
    }
    if len(list) == 0 {
        t.Fatal("Func Lists should return non-empty slice")
    }
}
