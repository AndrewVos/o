package o

import (
  "fmt"
  "reflect"
  "strconv"
  "github.com/AndrewVos/colour"
)

func O(i interface{}) {
  fmt.Println(o(i))
}

func o(i interface{}) string {
  return write("", 0, i)
}

func write(name string, depth int, i interface{}) string {
  result := name

  t := reflect.TypeOf(i)
  if t.Kind() == reflect.Ptr{
    t = t.Elem()
  }
  if t.Kind() == reflect.Int {
    result += writeInt(i)
  } else if t.Kind() == reflect.String {
    result += writeString(i)
  } else if t.Kind() == reflect.Bool {
    result += writeBool(i)
  } else if t.Kind() == reflect.Struct {
    result += writeStruct(depth, i)
  } else if t.Kind() == reflect.Slice {
    result += writeSlice(depth, i)
  } else if t.Kind() == reflect.Map {
    result += writeMap(depth, i)
  }
  return result
}

func writeBool(i interface{}) string {
  if i.(bool) {
    return colourValue("true")
  } else {
    return colourValue("false")
  }
}

func writeInt(i interface{}) string {
  value := reflect.ValueOf(i)
  if value.Kind() == reflect.Ptr {
    value = value.Elem()
  }
  return colourValue(strconv.Itoa(int(value.Int())))
}

func writeSlice(depth int, thing interface{}) string {
  result := colourTitle("slice") + " [" + "\n"
  s := reflect.ValueOf(thing)

  for i := 0; i < s.Len(); i++ {
    result += margin(depth + 1) + write("", depth + 1, s.Index(i).Interface()) + ",\n"
  }

  result += margin(depth) + "]"
  return result
}

func writeString(thing interface{}) string {
  quote := colourQuotes(`"`)
  return quote + colourValue(thing.(string)) + quote
}

func writeStruct(depth int, thing interface{}) string {
  t := reflect.TypeOf(thing)
  if t.Kind() == reflect.Ptr{
    t = t.Elem()
  }
  value := reflect.ValueOf(thing)

  result := colourTitle(t.Name()) + " {\n"

  for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)

    if !field.Anonymous {
      it := value.Field(i).Interface()
      displayName := colourField(field.Name) + ": "
      result += margin(depth + 1) + write(displayName, depth + 1, it) + "\n"
    }
  }

  return result + margin(depth) + "}"
}

func writeMap(depth int, thing interface{}) string {
  result := colourTitle("map") + " {\n"
  t := reflect.TypeOf(thing)
  if t.Kind() == reflect.Ptr{ t = t.Elem() }

  value := reflect.ValueOf(thing)
  for _, key := range value.MapKeys() {
    mapValue := value.MapIndex(key)
    result += margin(depth + 1) + write("", depth + 1, key.Interface()) + ": " + write("", depth + 1, mapValue.Interface()) + ",\n"
  }
  return result + margin(depth) + "}"
}

func margin(depth int) string {
  m := ""
  if depth == 0 { return m }
  for i:= 1; i <= depth; i++ {
    m += "  "
  }
  return m
}

func colourQuotes(quote string) string { return colour.Red(quote) }
func colourTitle(title string) string { return colour.Blue(title) }
func colourField(field string) string { return colour.Green(field) }
func colourValue(value string) string { return colour.Yellow(value) }
