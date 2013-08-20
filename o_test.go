package o

import (
  "testing"
  "strings"
)

type Struct1 struct {
  Name string
}

func assertOutputContains(t *testing.T, output string, expected string) {
  if strings.Contains(output, expected) == false {
    t.Errorf("Expected output to contain %q, but was:\n%v", expected, output)
  }
}

func TestStruct(t *testing.T) {
  s := Struct1{Name: "Arthur"}
  output := OO(s)
  assertOutputContains(t, output, "Struct1")
  assertOutputContains(t, output, "Name")
  assertOutputContains(t, output, "Arthur")
}
