package xrequests

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
)

func contains(resp int, status []int) bool {
	for _, stat := range status {
		if stat == resp {
			return true
		}
	}
	return false
}

func changeMapToURLValues(data map[string]interface{}) url.Values {
	var newValues = url.Values{}

	for k, v := range data {
		switch val := v.(type) {
		case string:
			newValues.Add(k, val)
		case bool:
			newValues.Add(k, strconv.FormatBool(val))
		case json.Number:
			newValues.Add(k, string(val))
		case int:
			newValues.Add(k, strconv.FormatInt(int64(val), 10))
		case int8:
			newValues.Add(k, strconv.FormatInt(int64(val), 10))
		case int16:
			newValues.Add(k, strconv.FormatInt(int64(val), 10))
		case int32:
			newValues.Add(k, strconv.FormatInt(int64(val), 10))
		case int64:
			newValues.Add(k, strconv.FormatInt(int64(val), 10))
		case float32:
			newValues.Add(k, strconv.FormatFloat(float64(val), 'f', -1, 64))
		case float64:
			newValues.Add(k, strconv.FormatFloat(float64(val), 'f', -1, 64))
		case []string:
			for _, element := range val {
				newValues.Add(k, element)
			}
		case []int:
			for _, element := range val {
				newValues.Add(k, strconv.FormatInt(int64(element), 10))
			}
		case []bool:
			for _, element := range val {
				newValues.Add(k, strconv.FormatBool(element))
			}
		case []float32:
			for _, element := range val {
				newValues.Add(k, strconv.FormatFloat(float64(element), 'f', -1, 64))
			}
		case []float64:
			for _, element := range val {
				newValues.Add(k, strconv.FormatFloat(float64(element), 'f', -1, 64))
			}
		case []interface{}:
			if len(val) < 1 {
				continue
			}

			switch val[0].(type) {
			case string:
				for _, element := range val {
					newValues.Add(k, element.(string))
				}
			case bool:
				for _, element := range val {
					newValues.Add(k, strconv.FormatBool(element.(bool)))
				}
			case json.Number:
				for _, element := range val {
					newValues.Add(k, string(element.(json.Number)))
				}
			}
		default:
		}
	}

	return newValues
}

func makeSliceOfReflectValue(v reflect.Value) (slice []interface{}) {
	kind := v.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return slice
	}

	slice = make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		slice[i] = v.Index(i).Interface()
	}

	return slice
}
