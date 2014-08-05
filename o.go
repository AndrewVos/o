package o

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/AndrewVos/colour"
)

func O(thing interface{}) {
	fmt.Println(o(thing))
}

func o(thing interface{}) string {
	return write("", 0, thing)
}

func write(name string, depth int, thing interface{}) string {
	result := name

	thingType := reflect.TypeOf(thing)
	if thingType.Kind() == reflect.Ptr {
		thingType = thingType.Elem()
	}

	if thingType.Kind() == reflect.Int {
		result += writeInt(thing)
	} else if thingType.Kind() == reflect.Uint8 {
		result += writeUint(thing)
	} else if thingType.Kind() == reflect.String {
		result += writeString(thing)
	} else if thingType.Kind() == reflect.Bool {
		result += writeBool(thing)
	} else if thingType == reflect.TypeOf(time.Time{}) {
		result += writeTime(thing)
	} else if thingType.Kind() == reflect.Slice {
		result += writeSlice(depth, thing)
	} else if thingType.Kind() == reflect.Map {
		result += writeMap(depth, thing)
	} else if thingType.Kind() == reflect.Struct {
		result += writeStruct(depth, thing)
	}
	return result
}

func writeBool(thing interface{}) string {
	thingValue := reflect.ValueOf(thing)
	if thingValue.Kind() == reflect.Ptr {
		thingValue = thingValue.Elem()
	}
	if thingValue.Bool() {
		return colourValue("true")
	} else {
		return colourValue("false")
	}
}

func writeInt(thing interface{}) string {
	thingValue := reflect.ValueOf(thing)
	if thingValue.Kind() == reflect.Ptr {
		thingValue = thingValue.Elem()
	}
	return colourValue(strconv.Itoa(int(thingValue.Int())))
}

func writeUint(thing interface{}) string {
	thingValue := reflect.ValueOf(thing)
	if thingValue.Kind() == reflect.Ptr {
		thingValue = thingValue.Elem()
	}
	return colourValue(strconv.Itoa(int(thingValue.Uint())))
}

func writeSlice(depth int, thing interface{}) string {
	result := colourTitle("slice") + " ["
	thingValue := reflect.ValueOf(thing)
	if thingValue.Kind() == reflect.Ptr {
		thingValue = thingValue.Elem()
	}

	hasItems := false
	if thingValue.IsValid() {
		if thingValue.Len() > 0 {
			result += "\n"
			hasItems = true
		}
		for elementIndex := 0; elementIndex < thingValue.Len(); elementIndex++ {
			element := thingValue.Index(elementIndex).Interface()
			result += margin(depth+1) + write("", depth+1, element) + ",\n"
		}
	}

	if hasItems {
		result += margin(depth) + "]"
	} else {
		result += "]"
	}
	return result
}

func writeString(thing interface{}) string {
	thingValue := reflect.ValueOf(thing)
	if thingValue.Kind() == reflect.Ptr {
		thingValue = thingValue.Elem()
	}
	quote := colourQuotes(`"`)
	return quote + colourValue(thingValue.String()) + quote
}

func writeTime(thing interface{}) string {
	thingValue := reflect.ValueOf(thing)
	if thingValue.Kind() == reflect.Ptr {
		thingValue = thingValue.Elem()
	}
	t := thingValue.Interface().(time.Time)
	return colourValue(fmt.Sprintf("%v", t))
}

func writeStruct(depth int, thing interface{}) string {
	thingType := reflect.TypeOf(thing)
	if thingType.Kind() == reflect.Ptr {
		thingType = thingType.Elem()
	}
	value := reflect.ValueOf(thing)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	result := colourTitle(thingType.Name()) + " {\n"

	if value.IsValid() {
		for fieldIndex := 0; fieldIndex < thingType.NumField(); fieldIndex++ {
			field := thingType.Field(fieldIndex)

			if !field.Anonymous {
				childThingField := value.Field(fieldIndex)
				if childThingField.CanInterface() {
					childThing := childThingField.Interface()
					displayName := colourField(field.Name) + ": "
					result += margin(depth+1) + write(displayName, depth+1, childThing) + "\n"
				}
			}
		}
	}

	return result + margin(depth) + "}"
}

func writeMap(depth int, thing interface{}) string {
	result := colourTitle("map") + " {\n"

	thingValue := reflect.ValueOf(thing)
	if thingValue.Kind() == reflect.Ptr {
		thingValue = thingValue.Elem()
	}
	for _, key := range thingValue.MapKeys() {
		mapValue := thingValue.MapIndex(key)
		result += margin(depth+1) + write("", depth+1, key.Interface()) + ": " + write("", depth+1, mapValue.Interface()) + ",\n"
	}
	return result + margin(depth) + "}"
}

func margin(depth int) string {
	m := ""
	if depth == 0 {
		return m
	}
	for i := 1; i <= depth; i++ {
		m += "  "
	}
	return m
}

func colourQuotes(quote string) string { return colour.Red(quote) }
func colourTitle(title string) string  { return colour.Blue(title) }
func colourField(field string) string  { return colour.Green(field) }
func colourValue(value string) string  { return colour.Yellow(value) }
