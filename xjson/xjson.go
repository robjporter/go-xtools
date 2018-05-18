package xjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type result struct {
	data interface{}
	err  error
}

type Xjson struct {
	raw  string
	data interface{}
}

func New() *Xjson {
	return &Xjson{}
}

func (x *Xjson) ParseBytes(json []byte) *Xjson {
	return x.ParseString(string(json))
}

func (x *Xjson) ParseString(data string) *Xjson {
	var f interface{}
	err := json.Unmarshal([]byte(data), &f)
	if err != nil {
		panic(err)
	}

	switch f.(type) {
	case []interface{}:
		x.data = f.([]interface{})
	case map[string]interface{}:
		x.data = f.(map[string]interface{})
	}

	x.raw = data
	return x
}

// PRINT ///////////////////////////////////

func (x *Xjson) Print() string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(x.raw), "", "\t")
	if err != nil {
		return x.raw
	}
	return out.String()
}

func (x *Xjson) PrintTrim() string {
	return strings.Replace(x.raw, " ", "", -1)
}

// FOR ///////////////////////////////////
func (r *result) ForEach(iterator func(key string, value interface{}) bool) {
	if !r.Exists() {
		return
	}
	switch r.data.(type) {
	case []interface{}:
		tmp := r.data.([]interface{})
		for i := 0; i < len(tmp); i++ {
			iterator(strconv.Itoa(i), tmp[i])
		}
	case map[string]interface{}:
		tmp := r.data.(map[string]interface{})
		for k, v := range tmp {
			iterator(k, v)
		}
	}
	return
}

// IS ///////////////////////////////////
func (r *result) IsArray() bool {
	switch r.data.(type) {
	case []interface{}:
		return true
	case map[string]interface{}:
		return true
	}
	return false
}

func (r *result) Exists() bool {
	if r.data == nil {
		return false
	}
	return true
}

func (x *Xjson) Exists() bool {
	return len(x.raw) != 0
}

// RESULT ///////////////////////////////////

func (r *result) String() string {
	return r.data.(string)
}

func (r *result) StringDef(def string) string {
	tmp := r.String()
	if tmp == "" {
		tmp = def
	}
	return tmp
}

func (r *result) StringArray() []string {
	var result []string
	switch r.data.(type) {
	case []interface{}:
		tmp := r.data.([]interface{})
		for i := 0; i < len(tmp); i++ {
			result = append(result, tmp[i].(string))
		}
	}
	return result
}

func (r *result) Int() int {
	switch r.data.(type) {
	case int:
		return r.data.(int)
	case float64:
		return int(r.data.(float64))
	}
	return -1234567
}

func (r *result) IntDef(def int) int {
	tmp := r.Int()
	if tmp == -1234567 {
		tmp = def
	}
	return tmp
}

func (r *result) IntArray() []int {
	var result []int
	switch r.data.(type) {
	case []interface{}:
		tmp := r.data.([]interface{})
		for i := 0; i < len(tmp); i++ {
			data := 0
			switch tmp[i].(type) {
			case int:
				data = tmp[i].(int)
			case float64:
				data = int(tmp[i].(float64))
			}
			result = append(result, data)
		}
	}
	return result
}

func (r *result) Bool() bool {
	return r.data.(bool)
}

func (r *result) Time() time.Time {
	return r.data.(time.Time)
}

func (r *result) Float() float64 {
	return r.data.(float64)
}

func (r *result) Interface() interface{} {
	return r.data
}

func (r *result) Map() map[string]interface{} {
	return r.data.(map[string]interface{})
}

// GET ///////////////////////////////////

func (x *Xjson) Get(path string) *result {
	var results result
	if x.data != "" {
		r, e := getValue(x.data, strings.Split(path, "."))
		results = result{data: r, err: e}
		return &results
	}
	return &results
}

func getValue(value interface{}, props []string) (val interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	if len(props) == 0 {
		return value, nil
	}
	first := props[0]
	rest := props[1:]
	switch vv := value.(type) {
	case map[string]interface{}:
		v := value.(map[string]interface{})[first]
		if first == "#" {
			return len(value.(map[string]interface{})), nil
		} else if first == "*" {
			fmt.Println("HERE")
		}
		return getValue(v, rest)
	case []interface{}:
		values := value.([]interface{})
		if first == "*" {
			vs := make([]interface{}, len(values))
			for i, v := range values {
				vs[i], _ = getValue(v, rest)
			}
			return vs, nil
		} else if first == "#" {
			return len(values), nil
		} else if strings.HasPrefix(first, "[") && strings.HasSuffix(first, "]") {
			first = strings.TrimLeft(first, "[")
			first = strings.TrimRight(first, "]")
			if strings.Contains(first, ":") {
				splits := strings.Split(first, ":")
				index := strings.Index(first, ":")
				start := 0
				stop := 0
				if index > 0 {
					a, e := strconv.Atoi(splits[0])
					if e == nil {
						start = a
					}
				}
				if len(splits) == 2 {
					a, e := strconv.Atoi(splits[1])
					if e == nil {
						stop = a
					}
				}
				if splits[1] == "" {
					stop = len(values)
				}
				pos := 0
				count := stop - start
				vs := make([]interface{}, count)
				for i, v := range values {
					if i >= start && i < stop {
						vs[pos], _ = getValue(v, rest)
						pos++
					}
				}
				return vs, nil
			} else if strings.Contains(first, "=") {
				splits := strings.Split(first, "=")
				var vs []interface{}
				if len(splits) == 2 {
					key := strings.TrimSpace(splits[0])
					value := strings.TrimSpace(splits[1])
					for _, v := range values {
						tmp := v.(map[string]interface{})[key]
						switch tmp.(type) {
						case int, int8, int16, int32, int64:
							a, _ := strconv.Atoi(value)
							if tmp.(int) == a {
								data, _ := getValue(v, rest)
								vs = append(vs, data)
							}
						case float64:
							a, _ := strconv.Atoi(value)
							if int(tmp.(float64)) == a {
								data, _ := getValue(v, rest)
								vs = append(vs, data)
							}
						case string:
							if tmp.(string) == value {
								data, _ := getValue(v, rest)
								vs = append(vs, data)
							}
						}
					}
					fmt.Println(key, value)
					fmt.Println(len(values))
				}
				return vs, nil
			}
		}
		i, err := strconv.ParseInt(first, 10, 0)
		if err != nil {
			return nil, err
		}
		return getValue(values[i], rest)
	default:
		err := fmt.Errorf("Unsupported type: %v, for value: %#v", vv, value)
		return value, err
	}
}
