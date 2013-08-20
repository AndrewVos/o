package o

import (
  "reflect"
  "strings"
  "strconv"
  "github.com/AndrewVos/colour"
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
  s := colouriseStructTitle(structType.Name()) + " {\n"

  attrs := map[string]string {}
  for i := 0; i < structType.NumField(); i++ {
    value := reflect.ValueOf(interfaceValue).Field(i)
    field := structType.Field(i)

    if !field.Anonymous {
      if value.Kind() == reflect.Int {
        attrs[field.Name] = strconv.Itoa(int(value.Int()))
      } else {
        attrs[field.Name] = value.String()
      }
    }
  }

  widest := 0
  for name,_ := range attrs {
    if length := len(name); length > widest { widest = length }
  }

  allFields := []string{}
  for name, value := range attrs {
    allFields = append(allFields, "  " + colouriseField(rjust(name, widest)) + ": " + colouriseValue(value))
  }

  s = s + strings.Join(allFields, "\n") + "\n}\n"

  return s
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
