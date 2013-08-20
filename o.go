package o

import (
  "reflect"
  "strings"
)

func OO(i interface{} ) string {
  t := reflect.TypeOf(i)
  if t.Kind() == reflect.Ptr{
    t = t.Elem()
  }

  if t.Kind() == reflect.Struct {
    return writeStruct(i, t)
  }

  return ""
}

func writeStruct(interfaceValue interface{}, structType reflect.Type) string {
  s := structType.Name()

  attrs := map[string]string {}
  for i := 0; i < structType.NumField(); i++ {
    value := reflect.ValueOf(interfaceValue).Field(i)
    field := structType.Field(i)

    if !field.Anonymous {
      attrs[field.Name] = (value.String())
    }
  }

  widest := 0
  for name,_ := range attrs {
    if length := len(name); length > widest { widest = length }
  }

  allFields := []string{}
  for name, value := range attrs {
    allFields = append(allFields, rjust(name, widest) + ":" + value)
  }

  s = s + "\n" + strings.Join(allFields, "\n")

  return s
}

func rjust(text string, width int) string {
  if len(text) < width {
    for len(text) < width {
      text = " " + text
    }
  }
  return text
}
