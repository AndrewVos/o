package o

import (
  "testing"
  "strings"
  "fmt"
)

type Struct1 struct {
  Name string
  Age int
  Married bool
}

func assertOutputContains(t *testing.T, output string, expected string) {
  if strings.Contains(output, expected) == false {
    t.Errorf("Expected output to contain %q, but was:\n%v", expected, output)
  }
}

func TestStruct(t *testing.T) {
  s := Struct1{Name: "Arthur", Age: 42}
  output := o(s)
  fmt.Println(output)
  assertOutputContains(t, output, "Struct1")

  assertOutputContains(t, output, "Name")
  assertOutputContains(t, output, "Arthur")

  assertOutputContains(t, output, "Age")
  assertOutputContains(t, output, "42")

  assertOutputContains(t, output, "Married")
  assertOutputContains(t, output, "false")
}

type StructWithName struct {
  Name string
}
type StructWithAge struct {
  Age int
}

type StructWithDepth struct {
  Name StructWithName
  Age StructWithAge
}

func TestStructWithDepth(t *testing.T) {
  s := StructWithDepth { Name: StructWithName { Name: "Mika" }, Age: StructWithAge { Age: 10 } }
  output := o(s)
  fmt.Println(output)
  assertOutputContains(t, output, "Name")
  assertOutputContains(t, output, "Mika")
}
