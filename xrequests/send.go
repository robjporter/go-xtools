package xrequests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

func (a *Agent) Send(content interface{}) *Agent {
	a.ClearSendDataNow()
	switch v := reflect.ValueOf(content); v.Kind() {
	case reflect.String:
		a.SendString(v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		a.SendString(strconv.FormatInt(v.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		a.SendString(strconv.FormatUint(v.Uint(), 10))
	case reflect.Float32:
		a.SendString(strconv.FormatFloat(v.Float(), 'f', -1, 32))
	case reflect.Float64:
		a.SendString(strconv.FormatFloat(v.Float(), 'f', -1, 64))
	case reflect.Bool:
		a.SendString(strconv.FormatBool(v.Bool()))
	case reflect.Struct:
		a.SendStruct(v.Interface())
	case reflect.Slice:
		a.SendSlice(makeSliceOfReflectValue(v))
	case reflect.Array:
		a.SendSlice(makeSliceOfReflectValue(v))
	case reflect.Ptr:
		a.Send(v.Elem().Interface())
	case reflect.Map:
		a.SendMap(v.Interface())
	default:
		return a
	}
	return a
}

func (a *Agent) SendSlice(content []interface{}) *Agent {
	a.SliceData = append(a.SliceData, content...)
	return a
}

func (a *Agent) SendMap(content interface{}) *Agent {
	return a.SendStruct(content)
}

func (a *Agent) SendString(content interface{}) *Agent {
	content2 := content.(string)
	if !a.BounceToRawString {
		var val interface{}

		b := json.NewDecoder(strings.NewReader(content2))
		b.UseNumber()

		if err := b.Decode(&val); err == nil {
			switch v := reflect.ValueOf(val); v.Kind() {
			case reflect.Map:
				for k, v := range val.(map[string]interface{}) {
					a.Data[k] = v
				}
			case reflect.Slice:
				a.SendSlice(val.([]interface{}))
			default:
				a.BounceToRawString = true
			}
		} else if formData, err := url.ParseQuery(content2); err == nil {
			for k, v := range formData {
				for _, v2 := range v {
					if val, ok := a.Data[k]; ok {
						var strA []string
						strA = append(strA, string(v2))
						switch oldv := val.(type) {
						case []string:
							strA = append(strA, oldv...)
						case string:
							strA = append(strA, oldv)
						}
						a.Data[k] = strA
					} else {
						a.Data[k] = v2
					}
				}
			}
		} else {
			a.BounceToRawString = true
		}
	}

	a.RawString += content2
	return a
}

func (a *Agent) SendStruct(content interface{}) *Agent {
	if mc, err := json.Marshal(content); err != nil {
		a.Errors = append(a.Errors, err)
	} else {
		var val map[string]interface{}
		b := json.NewDecoder(bytes.NewBuffer(mc))
		b.UseNumber()
		if err := b.Decode(&val); err != nil {
			a.Errors = append(a.Errors, err)
		} else {
			for k, v := range val {
				a.Data[k] = v
			}
		}
	}

	return a
}

func (a *Agent) SendFile(file interface{}, args ...string) *Agent {
	filename := ""
	fieldname := "file"

	if len(args) >= 1 && len(args[0]) > 0 {
		filename = strings.TrimSpace(args[0])
	}
	if len(args) >= 2 && len(args[1]) > 0 {
		fieldname = strings.TrimSpace(args[1])
	}
	if fieldname == "file" || fieldname == "" {
		filename = "file" + strconv.Itoa(len(a.FileData)+1)
	}

	switch v := reflect.ValueOf(file); v.Kind() {
	case reflect.String:
		a.processReflectString(v, filename, fieldname)
	case reflect.Slice:
		a.processReflectSlice(v, filename, fieldname)
	case reflect.Ptr:
		return a.ProcessReflectPointer(v, filename, fieldname, args)
	default:
		return a.processReflectDefault(v, filename, fieldname)
	}

	return a
}

func (a *Agent) processReflectString(v reflect.Value, filename string, fieldname string) {
	pathToFile, err := filepath.Abs(v.String())
	if err != nil {
		a.Errors = append(a.Errors, err)
		return
	}

	if filename == "" {
		filename = filepath.Base(pathToFile)
	}

	data, err := ioutil.ReadFile(v.String())
	if err != nil {
		a.Errors = append(a.Errors, err)
		return
	}

	a.FileData = append(a.FileData, File{
		Filename:  filename,
		Fieldname: fieldname,
		Data:      data,
	})
}

func (a *Agent) processReflectSlice(v reflect.Value, filename string, fieldname string) {
	slice := makeSliceOfReflectValue(v)
	if filename == "" {
		filename = "filename"
	}

	f := File{
		Filename:  filename,
		Fieldname: fieldname,
		Data:      make([]byte, len(slice)),
	}

	for i := range slice {
		f.Data[i] = slice[i].(byte)
	}

	a.FileData = append(a.FileData, f)
}

func (a *Agent) ProcessReflectPointer(v reflect.Value, filename string, fieldname string, args []string) *Agent {
	if len(args) == 1 {
		return a.SendFile(v.Elem().Interface(), args[0])
	}
	if len(args) >= 2 {
		return a.SendFile(v.Elem().Interface(), args[0], args[1])
	}
	return a.SendFile(v.Elem().Interface())
}

func (a *Agent) processReflectDefault(v reflect.Value, filename string, fieldname string) *Agent {
	if v.Type() == reflect.TypeOf(os.File{}) {
		osfile := v.Interface().(os.File)

		if filename == "" {
			filename = filepath.Base(osfile.Name())
		}

		data, err := ioutil.ReadFile(osfile.Name())
		if err != nil {
			a.Errors = append(a.Errors, err)
			return a
		}

		a.FileData = append(a.FileData, File{
			Filename:  filename,
			Fieldname: fieldname,
			Data:      data,
		})

		return a
	}
	a.Errors = append(a.Errors, errors.New("Unsupported Send format"))
	return a
}
