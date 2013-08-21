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

func TestString(t *testing.T) {
  s := "o.OOO"
  output := o(s)
  fmt.Println(output)
  assertOutputContains(t, output, "o.OOO")
}

func TestInt(t * testing.T) {
  i := 12345
  output := o(i)
  fmt.Println(output)
  assertOutputContains(t, output, "12345")
}

func TestBoolTrue(t *testing.T) {
  i := true
  output := o(i)
  assertOutputContains(t, output, "true")
}

func TestBoolFalse(t *testing.T) {
  i := false
  output := o(i)
  assertOutputContains(t, output, "false")
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
  NameStruct StructWithName
  AgeStruct StructWithAge
}

func TestStructWithDepth(t *testing.T) {
  s := StructWithDepth { NameStruct: StructWithName { Name: "Mika" }, AgeStruct: StructWithAge { Age: 10 } }
  output := o(s)
  fmt.Println(output)
  assertOutputContains(t, output, "NameStruct")
  assertOutputContains(t, output, "Name")
  assertOutputContains(t, output, "Mika")

  assertOutputContains(t, output, "AgeStruct")
  assertOutputContains(t, output, "Age")
  assertOutputContains(t, output, "10")
}

type Thing struct {
  ThingValue string
}

type StructWithArrays struct {
  Things []Thing
}

func TestArray(t *testing.T) {
  s := []Thing { { ThingValue: "ererrrmmmm" }, }
  output := o(s)
  fmt.Println(output)
  assertOutputContains(t, output, "slice")
  assertOutputContains(t, output, "ThingValue")
  assertOutputContains(t, output, "ererrrmmmm")
}

func TestStructWithArray(t *testing.T) {
  s := StructWithArrays{ Things: []Thing{ { ThingValue: "ermmm" }, } }
  output := o(s)
  fmt.Println(output)
  assertOutputContains(t, output, "StructWithArrays")
  assertOutputContains(t, output, "Things")
  assertOutputContains(t, output, "slice")
  assertOutputContains(t, output, "ThingValue")
  assertOutputContains(t, output, "ermmm")
}

func TestMap(t *testing.T) {
  s := map[string] string {
    "I like": "cake",
    "And also": "ice cream",
  }
  output := o(s)
  fmt.Println(output)
  assertOutputContains(t, output, "map")
  assertOutputContains(t, output, "I like")
  assertOutputContains(t, output, "cake")
  assertOutputContains(t, output, "And also")
  assertOutputContains(t, output, "ice cream")
}
