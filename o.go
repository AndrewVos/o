package o

import (
  "fmt"
  "reflect"
  "strings"
  "strconv"
  "github.com/AndrewVos/colour"
)

func O(i interface{}) {
  fmt.Println(o(i))
}

func o(i interface{}) string {
  t := reflect.TypeOf(i)
  if t.Kind() == reflect.Ptr{
    t = t.Elem()
  }

  if t.Kind() == reflect.String {
    return writeString(i)
  } else if t.Kind() == reflect.Int {
    return writeInt(i)
  } else if t.Kind() == reflect.Struct {
    return writeStruct(i, t, 0)
  }

  return t.Name()
}

func writeString(interfaceValue interface{}) string {
  return colouriseStructTitle("string") + " (\n" + margin(1) + colouriseValue(interfaceValue.(string)) + "\n)"
}

func writeInt(interfaceValue interface{}) string {
  return colouriseStructTitle("int") + " (\n" + margin(1) + colouriseValue(strconv.Itoa(interfaceValue.(int))) + "\n)"
}

func writeStruct(interfaceValue interface{}, structType reflect.Type, depth int) string {
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
      } else if value.Kind() == reflect.Struct {
        attributes[field.Name] = writeStruct(value.Interface(), value.Type(), depth + 1)
      } else {
        attributes[field.Name] = value.String()
      }
    }
  }

  widestAttributeName := widestAttributeName(attributes)
  allFields := []string{}
  for name, value := range attributes {
    allFields = append(allFields, margin(depth + 1) + colouriseField(rjust(name, widestAttributeName)) + ": " + colouriseValue(value))
  }

  return colouriseStructTitle(structType.Name()) + " {\n" + strings.Join(allFields, "\n") + "\n" + margin(depth) + "}"
}

func margin(depth int) string {
  m := ""
  if depth == 0 { return m }
  for i:= 1; i <= depth; i++ {
    m += "  "
  }
  return m
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
