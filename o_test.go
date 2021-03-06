package o

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"
)

type SimpleStruct struct {
	Name    string
	Age     int
	Married bool
}
type StructWithName struct{ Name string }
type StructWithAge struct{ Age int }
type StructWithDepth struct {
	NameStruct StructWithName
	AgeStruct  StructWithAge
}
type Thing struct{ ThingValue string }
type StructWithSlices struct{ Things []Thing }
type StructWithPointerFields struct{ SomeField *int }
type StructWithUnexportedFields struct {
	SomeField  int
	unexported int
}
type StructWithPointerField struct {
	SomeField *SimpleStruct
}

func assertOutput(t *testing.T, value interface{}, expected string) {
	begin := regexp.MustCompile("\\x1b\\[3[1-9];1m")
	end := regexp.MustCompile("\\x1b\\[0m")
	actual := o(value)
	if os.Getenv("OUTPUT") != "" {
		fmt.Println(actual)
		fmt.Println("----------------------------------------------------------------------------------------------------")
	}
	actual = begin.ReplaceAllString(actual, "")
	actual = end.ReplaceAllString(actual, "")
	actual = strings.TrimSpace(actual)
	expected = strings.TrimSpace(expected)
	if actual != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v\n", expected, actual)
	}
}

func TestString(t *testing.T) {
	assertOutput(t, "o.OOO", `"o.OOO"`)
}

func TestPointerString(t *testing.T) {
	s := "sdfwef"
	assertOutput(t, &s, `"sdfwef"`)
}

func TestInt(t *testing.T) {
	assertOutput(t, 12345, "12345")
}

func TestPointerInt(t *testing.T) {
	s := 13113
	assertOutput(t, &s, "13113")
}

func TestBool(t *testing.T) {
	assertOutput(t, true, "true")
	assertOutput(t, false, "false")
}

func TestPointerBool(t *testing.T) {
	s := true
	assertOutput(t, &s, "true")
	s = false
	assertOutput(t, &s, "false")
}

func TestStruct(t *testing.T) {
	s := SimpleStruct{Name: "Arthur", Age: 42}
	expected := `
SimpleStruct {
  Name: "Arthur"
  Age: 42
  Married: false
}
	`
	assertOutput(t, s, expected)
}

func TestPointerStruct(t *testing.T) {
	s := SimpleStruct{Name: "Hello"}
	expected := `
SimpleStruct {
  Name: "Hello"
  Age: 0
  Married: false
}
	`
	assertOutput(t, &s, expected)
}

func TestStructWithPointerFields(t *testing.T) {
	i := 1000
	s := StructWithPointerFields{SomeField: &i}
	expected := `
StructWithPointerFields {
  SomeField: 1000
}
	`
	assertOutput(t, &s, expected)
}

func TestStructWithDepth(t *testing.T) {
	s := StructWithDepth{NameStruct: StructWithName{Name: "Mika"}, AgeStruct: StructWithAge{Age: 10}}
	expected := `
StructWithDepth {
  NameStruct: StructWithName {
    Name: "Mika"
  }
  AgeStruct: StructWithAge {
    Age: 10
  }
}
	`
	assertOutput(t, s, expected)
}

func TestSlice(t *testing.T) {
	s := []Thing{{ThingValue: "ererrrmmmm"}}
	expected := `
slice [
  Thing {
    ThingValue: "ererrrmmmm"
  },
]
	`
	assertOutput(t, s, expected)
}

func TestPointerSlice(t *testing.T) {
	a := "sdfsd"
	s := []*string{&a}
	expected := `
slice [
  "sdfsd",
]
	`
	assertOutput(t, &s, expected)
}

func TestStructWithSlices(t *testing.T) {
	s := StructWithSlices{Things: []Thing{{ThingValue: "ermmm"}}}
	expected := `
StructWithSlices {
  Things: slice [
    Thing {
      ThingValue: "ermmm"
    },
  ]
}
	`
	assertOutput(t, s, expected)
}

func TestMap(t *testing.T) {
	s := map[string]string{
		"I like":   "cake",
		"And also": "ice cream",
	}
	expected := `
map {
  "I like": "cake",
  "And also": "ice cream",
}
	`
	assertOutput(t, s, expected)
}

func TestPointerMap(t *testing.T) {
	s := map[string]string{"I like": "cake"}
	expected := `
map {
  "I like": "cake",
}
	`
	assertOutput(t, &s, expected)
}

func TestMapOfMaps(t *testing.T) {
	s := map[string]map[int]string{
		"meh":      map[int]string{9: "mergh"},
		"tired of": map[int]string{123: "thinking of test cases"},
	}
	expected := `
map {
  "meh": map {
    9: "mergh",
  },
  "tired of": map {
    123: "thinking of test cases",
  },
}
	`
	assertOutput(t, s, expected)
}

func TestMapWithStruct(t *testing.T) {
	s := map[int]Thing{
		982: Thing{"nwle22"},
		892: Thing{"ekel2n2l"},
	}
	expected := `
map {
  982: Thing {
    ThingValue: "nwle22"
  },
  892: Thing {
    ThingValue: "ekel2n2l"
  },
}
	`
	assertOutput(t, s, expected)
}

func TestStructWithUnexportedFields(t *testing.T) {
	s := StructWithUnexportedFields{}
	expected := `
StructWithUnexportedFields {
  SomeField: 0
}`
	assertOutput(t, s, expected)
}

func TestTime(t *testing.T) {
	s := time.Date(1983, time.October, 23, 10, 0, 0, 0, time.UTC)
	expected := "1983-10-23 10:00:00 +0000 UTC"
	assertOutput(t, s, expected)
}

func TestZeroSlice(t *testing.T) {
	var slice []string
	expected := `
slice []
	`
	assertOutput(t, slice, expected)
}

func TestStructWithZeroStructField(t *testing.T) {
	var s StructWithPointerField
	expected := `
StructWithPointerField {
  SomeField: SimpleStruct {
  }
}
	`
	assertOutput(t, s, expected)
}

func TestBytes(t *testing.T) {
	s := byte(123)
	expected := "123"
	assertOutput(t, s, expected)
}

func TestInts(t *testing.T) {
	var a int = 1
	var b int8 = 2
	var c int16 = 3
	var d int32 = 4
	var e int64 = 5

	s := []interface{}{a, b, c, d, e}

	expected := `
slice [
  1,
  2,
  3,
  4,
  5,
]
	`
	assertOutput(t, s, expected)
}

func TestUints(t *testing.T) {
	var a uint = 1
	var b uint8 = 2
	var c uint16 = 3
	var d uint32 = 4
	var e uint64 = 5

	s := []interface{}{a, b, c, d, e}

	expected := `
slice [
  1,
  2,
  3,
  4,
  5,
]
	`
	assertOutput(t, s, expected)
}
