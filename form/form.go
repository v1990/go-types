package form

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

const (
	_formTag = "form"
	_jsonTag = "json"
)

// Transform struct to form( url.Values )
func Transform(obj interface{}) (url.Values, error) {
	v := reflect.Indirect(reflect.ValueOf(obj))
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("only support struct,but pass %T", obj)
	}
	t := v.Type()
	numField := v.NumField()
	form := make(url.Values, numField)
	for i := 0; i < numField; i++ {
		vf := v.Field(i)
		tf := t.Field(i)

		tag, ok := tf.Tag.Lookup(_formTag)
		if !ok {
			tag, ok = tf.Tag.Lookup(_jsonTag)
		}
		if !ok {
			continue
		}
		switch tag {
		case "", "-":
			continue
		}

		items := strings.Split(tag, ",")
		name := items[0]
		if len(items) > 1 {
			if items[1] == "omitempty" && (!vf.IsValid() || vf.IsZero()) {
				continue
			}
		}
		value := toString(vf)
		form.Set(name, value)
	}

	return form, nil

}

func toString(vf reflect.Value) string {
	if !vf.IsValid() {
		return ""
	}
	switch vf.Kind() {
	case reflect.String:
		return vf.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(vf.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(vf.Uint(), 10)
	}
	return fmt.Sprint(vf.Interface())
}
