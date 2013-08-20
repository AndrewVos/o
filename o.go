package o

import (
  "fmt"
  "reflect"
  "strings"
  "strconv"
  "github.com/AndrewVos/colour"
)

func O(i interface{}) {
  fmt.Println(OO(i))
}

func OO(i interface{}) string {
  t := reflect.TypeOf(i)
  if t.Kind() == reflect.Ptr{
    t = t.Elem()
  }

  if t.Kind() == reflect.Struct {
    return writeStruct(i, t)
  }

  return t.Name()
}

func writeStruct(interfaceValue interface{}, structType reflect.Type) string {
  attributes := map[string]string {}
  for i := 0; i < structType.NumField(); i++ {
    field := structType.Field(i)

    if !field.Anonymous {
      value := reflect.ValueOf(interfaceValue)
      if value.Kind() == reflect.Ptr {
        value = value.Elem()
      }
      value = value.Field(i)
      if value.Kind() == reflect.Int {
        attributes[field.Name] = strconv.Itoa(int(value.Int()))
      } else if value.Kind() == reflect.Bool {
        if value.Bool() {
          attributes[field.Name] = "true"
        } else {
          attributes[field.Name] = "false"
        }
      } else {
        attributes[field.Name] = value.String()
      }
    }
  }

  widestAttributeName := widestAttributeName(attributes)
  allFields := []string{}
  for name, value := range attributes {
    allFields = append(allFields, "  " + colouriseField(rjust(name, widestAttributeName)) + ": " + colouriseValue(value))
  }

  return colouriseStructTitle(structType.Name()) + " {\n" + strings.Join(allFields, "\n") + "\n}\n"
}

func widestAttributeName(attributes map[string]string) int {
  widest := 0
  for name,_ := range attributes {
    if length := len(name); length > widest { widest = length }
  }
  return widest
}
func colouriseStructTitle(title string) string { return colour.Blue(title) }
func colouriseField(field string) string { return colour.Green(field) }
func colouriseValue(value string) string { return colour.Yellow(value) }

func rjust(text string, width int) string {
  if len(text) < width {
    for len(text) < width {
      text = " " + text
    }
  }
  return text
}
