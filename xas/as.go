package xas

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	KindTime reflect.Kind = iota + 1000000000
	KindDuration
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

var timeformats = []string{
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
	"Mon, 2 Jan 2006 15:04:05 -0700",
	"02.01.06",
	"01/02/06",
	"2006-01-02",
	"2006/01/02",
	"01/02/2006",
	"02.01.2006",
	"01/02/06 15:04",
	"2006-01-02 15:04",
	"2006-01-02T15:04",
	"01/02/2006 15:04",
	"02.01.06 15:04:05",
	"01/02/06 15:04:05",
	"01/02/2006 15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02T15:04:05",
	"02.01.2006 15:04:05",
}

func FormatIntToByte(b int64) string {
	multiple := ""
	value := float64(b)

	switch {
	case b < KB:
		return strconv.FormatInt(b, 10) + "B"
	case b < MB:
		value /= KB
		multiple = "KB"
	case b < MB:
		value /= KB
		multiple = "KB"
	case b < GB:
		value /= MB
		multiple = "MB"
	case b < TB:
		value /= GB
		multiple = "GB"
	case b < PB:
		value /= TB
		multiple = "TB"
	case b < EB:
		value /= PB
		multiple = "PB"
	default:
		value = 0
		multiple = ""
	}

	return fmt.Sprintf("%.02f%s", value, multiple)
}

func ToSliceInterface(iface ...interface{}) (ret []interface{}) {
	ret = []interface{}{}
	if len(iface) == 0 || iface[0] == nil {
		return
	}
	t := iface[0]
	val := reflect.ValueOf(t)
	if IsPtr(t) {
		val = val.Elem()
	}
	if val.Kind() != reflect.Slice || val.Len() == 0 {
		return
	}

	ret = make([]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		ret[i] = val.Index(i).Interface()
	}
	return
}

func IsPtr(obj interface{}) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Ptr
}

func Convert(value interface{}, t reflect.Kind) (interface{}, error) {

	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		switch t {
		case reflect.String:
			if reflect.TypeOf(value).Elem().Kind() == reflect.Uint8 {
				return ToString(value.([]byte)), nil
			} else {
				return ToString(value), nil
			}
		case reflect.Slice:
		default:
			return nil, fmt.Errorf("Could not convert slice into non-slice.")
		}
	case reflect.String:
		switch t {
		case reflect.Slice:
			return ToBytes(value), nil
		}
	}

	switch t {

	case reflect.String:
		return ToString(value), nil

	case reflect.Uint64:
		return ToUint(value), nil

	case reflect.Uint32:
		return uint32(ToUint(value)), nil

	case reflect.Uint16:
		return uint16(ToUint(value)), nil

	case reflect.Uint8:
		return uint8(ToUint(value)), nil

	case reflect.Uint:
		return uint(ToUint(value)), nil

	case reflect.Int64:
		return int64(ToInt64(value)), nil

	case reflect.Int32:
		return int32(ToInt64(value)), nil

	case reflect.Int16:
		return int16(ToInt64(value)), nil

	case reflect.Int8:
		return int8(ToInt64(value)), nil

	case reflect.Int:
		return int(ToInt64(value)), nil

	case reflect.Float64:
		return ToFloat(value), nil

	case reflect.Float32:
		return float32(ToFloat(value)), nil

	case reflect.Bool:
		return ToBool(value), nil

	case reflect.Interface:
		return value, nil

	case KindTime:
		return ToTime(false, value), nil

	case KindDuration:
		return ToDuration(value), nil

	}

	return nil, fmt.Errorf("Could not convert %s into %s.", reflect.TypeOf(value).Kind(), t)
}

func ToStringSlice(valuea ...interface{}) []string {
	value := valuea[0]
	var response []string

	switch v := value.(type) {
	case []interface{}:
		for _, u := range v {
			response = append(response, ToString(u))
		}
		return response
	case []string:
		return v
	case string:
		return strings.Fields(v)
	case interface{}:
		str := ToString(v)
		return []string{str}
	default:
		return response
	}
}

func ToStringMapString(valuea ...interface{}) map[string]string {
	value := valuea[0]

	var m = map[string]string{}

	switch v := value.(type) {
	case map[string]string:
		return v
	case map[string]interface{}:
		for k, val := range v {
			m[ToString(k)] = ToString(val)
		}
		return m
	case map[interface{}]string:
		for k, val := range v {
			m[ToString(k)] = ToString(val)
		}
		return m
	case map[interface{}]interface{}:
		for k, val := range v {
			m[ToString(k)] = ToString(val)
		}
		return m
	default:
		return m
	}
}

func ToSlice(valuea ...interface{}) []interface{} {
	var s []interface{}
	value := valuea[0]

	switch v := value.(type) {
	case []interface{}:
		for _, u := range v {
			s = append(s, u)
		}
		return s
	case []map[string]interface{}:
		for _, u := range v {
			s = append(s, u)
		}
		return s
	default:
		return s
	}
}

func ToIP(valuea ...interface{}) net.IP {
	addr := ToString(valuea[0])
	ip := net.ParseIP(addr)
	return ip
}

func OfType(valuea ...interface{}) string {
	value := valuea[0]
	return fmt.Sprintf("%T", value)
}

func IsKind(valuea ...interface{}) bool {
	value := valuea[0]
	value2 := valuea[1]
	return value == OfKind(value2)
}

func OfKind(valuea ...interface{}) string {
	value := valuea[0]
	return reflect.ValueOf(value).Kind().String()
}

func ToBase64(valuea ...interface{}) string {
	value := valuea[0]
	return base64.StdEncoding.EncodeToString([]byte(ToString(value)))
}

func FromBase64(valuea ...interface{}) string {
	value := valuea[0]
	data, err := base64.StdEncoding.DecodeString(ToString(value))
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func IsEmpty(valuea ...interface{}) bool {
	value := valuea[0]
	g := reflect.ValueOf(value)
	if !g.IsValid() {
		return true
	}
	switch g.Kind() {
	default:
		return g.IsNil()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return g.Len() == 0
	case reflect.Bool:
		return g.Bool() == false
	case reflect.Complex64, reflect.Complex128:
		return g.Complex() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return g.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return g.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return g.Float() == 0
	case reflect.Struct:
		return false
	}
	return true
}

func ToRuneLength(valuea ...interface{}) int {
	value := ToString(valuea[0])
	return utf8.RuneCountInString(value)
}

func ToBool(valuea ...interface{}) bool {
	value := valuea[0]
	if ToInt(value) > 0 {
		return true
	}
	b, _ := strconv.ParseBool(ToString(value))
	return b
}

func ToBytes(valuea ...interface{}) []byte {
	value := valuea[0]
	if value == nil {
		return []byte{}
	}

	switch val := value.(type) {
	case bool:
		if val == true {
			return []byte("true")
		}
		return []byte("false")
	case string:
		return []byte(val)
	case []byte:
		return val
	default:
		return []byte(fmt.Sprintf("%v", value))
	}
}

func ToDuration(valuea ...interface{}) time.Duration {
	value := valuea[0]
	switch value.(type) {
	case int, int8, int16, int32, int64:
		return time.Duration(ToInt(value))
	case uint, uint8, uint16, uint32, uint64:
		return time.Duration(ToInt(value))
	case float32, float64:
		return time.Duration(ToInt(value))
	default:
		dur, _ := time.ParseDuration(ToString(value))
		return dur
	}
}

func ToFixedLengthAfter(str string, spacer string, length int) string {
	spacer = spacer[:1]
	l := length - len(str)
	if l > 0 {
		return str + strings.Repeat(spacer, l)
	}
	if l == 0 {
		return str
	}
	return str[:length]
}

func ToFixedLengthBefore(str string, spacer string, length int) string {
	spacer = spacer[:1]
	l := length - len(str)
	if l > 0 {
		return strings.Repeat(spacer, l) + str
	}
	if l == 0 {
		return str
	}
	return str[:length]
}

func ToFixedLengthCenter(str string, spacer string, length int) string {
	spacer = spacer[:1]
	l := length - len(str)
	if l > 0 {
		if l%2 == 0 {
			l = l / 2
			return strings.Repeat(spacer, l) + str + strings.Repeat(spacer, l)
		}
		l = (l + 1) / 2
		return strings.Repeat(spacer, l) + str + strings.Repeat(spacer, l-1)
	}
	if l == 0 {
		return str
	}
	return str[:length]
}

func ToFloat(valuea ...interface{}) float64 {
	value := valuea[0]
	switch val := value.(type) {
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return float64(val)
	case time.Time:
		return float64(val.Unix())
	case bool:
		if val == true {
			return float64(1)
		}
		return float64(0)
	default:
		f, _ := strconv.ParseFloat(ToString(value), 64)
		return float64(f)
	}
}

func ToFloat32(valuea ...interface{}) float32 {
	value := valuea[0]
	switch val := value.(type) {
	case int:
		return float32(val)
	case int8:
		return float32(val)
	case int16:
		return float32(val)
	case int32:
		return float32(val)
	case int64:
		return float32(val)
	case uint:
		return float32(val)
	case uint8:
		return float32(val)
	case uint16:
		return float32(val)
	case uint32:
		return float32(val)
	case uint64:
		return float32(val)
	case float32:
		return float32(val)
	case float64:
		return float32(val)
	case time.Time:
		return float32(val.Unix())
	case bool:
		if val == true {
			return float32(1)
		}
		return float32(0)
	default:
		f, _ := strconv.ParseFloat(ToString(value), 32)
		return float32(f)
	}
}

func ToFloatFromXString(valuea ...string) float64 {
	value := valuea[0]
	value = strings.Trim(value, "\t\n\r¢§$€ ")
	var float float64
	c := strings.Count(value, ",")
	p := strings.Count(value, ".")
	fc := strings.Index(value, ",")
	fp := strings.Index(value, ".")
	if c == 0 && p == 1 {
		float, _ = strconv.ParseFloat(value, 64)
	} else if c == 1 && p == 0 {
		value = strings.Replace(value, ",", ".", 1)
		float, _ = strconv.ParseFloat(value, 64)
	} else if c == 0 && p == 0 {
		intx, _ := strconv.ParseInt(value, 0, 64)
		float = float64(intx)
	} else if c > 1 && p < 2 {
		value = strings.Replace(value, ",", "", -1)
		float, _ = strconv.ParseFloat(value, 64)
	} else if c < 2 && p > 1 {
		value = strings.Replace(value, ".", "", -1)
		value = strings.Replace(value, ",", ".", 1)
		float, _ = strconv.ParseFloat(value, 64)
	} else if c == 1 && p == 1 {
		if fp < fc {
			value = strings.Replace(value, ".", "", -1)
			value = strings.Replace(value, ",", ".", 1)
		} else {
			value = strings.Replace(value, ",", "", -1)
		}
		float, _ = strconv.ParseFloat(value, 64)
	} else {
		value = "0"
		float, _ = strconv.ParseFloat(value, 64)
	}
	return float64(float)
}

func ToInt(valuea ...interface{}) int {
	value := valuea[0]
	switch val := value.(type) {
	case int:
		return int(val)
	case int8:
		return int(val)
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int64:
		return int(val)
	case uint:
		return int(val)
	case uint8:
		return int(val)
	case uint16:
		return int(val)
	case uint32:
		return int(val)
	case uint64:
		return int(val)
	case float32:
		return int(val + 0.5)
	case float64:
		return int(val + 0.5)
	case time.Time:
		return int(val.Unix())
	case bool:
		if val == true {
			return int(1)
		}
		return int(0)
	default:
		i, _ := strconv.ParseFloat(ToString(value), 64)
		return int(i + 0.5)
	}
}

func ToInt64(valuea ...interface{}) int64 {
	value := valuea[0]
	switch val := value.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return int64(val)
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case float32:
		return int64(val + 0.5)
	case float64:
		return int64(val + 0.5)
	case time.Time:
		return int64(val.Unix())
	case bool:
		if val == true {
			return int64(1)
		}
		return int64(0)
	default:
		i, _ := strconv.ParseFloat(ToString(value), 64)
		return int64(i + 0.5)
	}
}

func ToInt32(valuea ...interface{}) int32 {
	value := valuea[0]
	switch val := value.(type) {
	case int:
		return int32(val)
	case int8:
		return int32(val)
	case int16:
		return int32(val)
	case int32:
		return int32(val)
	case int64:
		return int32(val)
	case uint:
		return int32(val)
	case uint8:
		return int32(val)
	case uint16:
		return int32(val)
	case uint32:
		return int32(val)
	case uint64:
		return int32(val)
	case float32:
		return int32(val + 0.5)
	case float64:
		return int32(val + 0.5)
	case time.Time:
		return int32(val.Unix())
	case bool:
		if val == true {
			return int32(1)
		}
		return int32(0)
	default:
		i, _ := strconv.ParseFloat(ToString(value), 32)
		return int32(i + 0.5)
	}
}

func ToInt16(valuea ...interface{}) int16 {
	value := valuea[0]
	switch val := value.(type) {
	case int:
		return int16(val)
	case int8:
		return int16(val)
	case int16:
		return int16(val)
	case int32:
		return int16(val)
	case int64:
		return int16(val)
	case uint:
		return int16(val)
	case uint8:
		return int16(val)
	case uint16:
		return int16(val)
	case uint32:
		return int16(val)
	case uint64:
		return int16(val)
	case float32:
		return int16(val + 0.5)
	case float64:
		return int16(val + 0.5)
	case time.Time:
		return int16(val.Unix())
	case bool:
		if val == true {
			return int16(1)
		}
		return int16(0)
	default:
		i, _ := strconv.ParseFloat(ToString(value), 16)
		return int16(i + 0.5)
	}
}

func ToInt8(valuea ...interface{}) int8 {
	value := valuea[0]
	switch val := value.(type) {
	case int:
		return int8(val)
	case int8:
		return int8(val)
	case int16:
		return int8(val)
	case int32:
		return int8(val)
	case int64:
		return int8(val)
	case uint:
		return int8(val)
	case uint8:
		return int8(val)
	case uint16:
		return int8(val)
	case uint32:
		return int8(val)
	case uint64:
		return int8(val)
	case float32:
		return int8(val + 0.5)
	case float64:
		return int8(val + 0.5)
	case time.Time:
		return int8(val.Unix())
	case bool:
		if val == true {
			return int8(1)
		}
		return int8(0)
	default:
		i, _ := strconv.ParseFloat(ToString(value), 8)
		return int8(i + 0.5)
	}
}

func ToStringMap(valuea ...interface{}) map[string]interface{} {
	value := valuea[0]

	var m = map[string]interface{}{}

	switch v := value.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[ToString(k)] = val
		}
		return m
	case map[string]interface{}:
		return v
	default:
		return m
	}
}

func ToString(valuea ...interface{}) string {
	value := valuea[0]
	if value == nil {
		return ""
	}

	switch val := value.(type) {
	case bool:
		if value.(bool) == true {
			return "true"
		}
		return "false"
	case time.Duration:
		return string(val.String())
	case time.Time:
		return string(val.Format(time.RFC3339))
	case string:
		return val
	case []byte:
		return string(val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

func Trimmed(valuea ...interface{}) string {
	value := valuea[0]
	return strings.TrimSpace(ToString(value))
}

func ToUint(valuea ...interface{}) uint64 {
	value := valuea[0]
	switch val := value.(type) {
	case int:
		return uint64(val)
	case int8:
		return uint64(val)
	case int16:
		return uint64(val)
	case int32:
		return uint64(val)
	case int64:
		return uint64(val)
	case uint:
		return uint64(val)
	case uint8:
		return uint64(val)
	case uint16:
		return uint64(val)
	case uint32:
		return uint64(val)
	case uint64:
		return uint64(val)
	case float32:
		return uint64(val + 0.5)
	case float64:
		return uint64(val + 0.5)
	case time.Time:
		return uint64(val.Unix())
	case bool:
		if val == true {
			return uint64(1)
		}
		return uint64(0)
	default:
		i, _ := strconv.ParseFloat(ToString(value), 64)
		return uint64(i + 0.5)
	}
}

func IsInt(input interface{}) bool {
	switch input.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}

func IsFloat(input interface{}) bool {
	switch input.(type) {
	case float32, float64:
		return true
	}
	return false
}

func IsBool(input interface{}) bool {
	switch input.(type) {
	case bool:
		return true
	}
	return false
}

func IsString(input interface{}) bool {
	switch input.(type) {
	case string:
		return true
	}
	return false
}

func IsTime(input interface{}) bool {
	switch input.(type) {
	case time.Time:
		return true
	}
	return false
}

func IsNillable(input interface{}) bool {
	switch reflect.TypeOf(input).Kind() {
	case reflect.Chan:
		return true
	case reflect.Slice:
		return true
	case reflect.UnsafePointer:
		return true
	default:
		return false
	}
}

func ToMapInterfaceFromMapString(src map[string]string) map[string]interface{} {
	ret := map[string]interface{}{}
	for k, v := range src {
		ret[k] = v
	}
	return ret
}

func ToMapStringFromMapInterface(src map[interface{}]interface{}) map[string]interface{} {
	tgt := map[string]interface{}{}
	for k, v := range src {
		tgt[fmt.Sprintf("%v", k)] = v
	}
	return tgt
}

func FromTime(time time.Time) string {
	return strconv.FormatInt(time.Unix(), 10)
}

func ToTime(simple bool, valuea ...interface{}) time.Time {
	value := valuea[0]
	s := ToString(value)
	if simple {
		intValue, _ := strconv.ParseInt(s, 10, 0)
		return time.Unix(intValue, 0)
	}
	for _, format := range timeformats {
		r, err := time.Parse(format, s)
		if err == nil {
			return r
		}
	}
	return time.Time{}
}

func ToCamelCase(str string) string {
	if len(str) == 0 {
		return ""
	}

	buf := &bytes.Buffer{}
	var r0, r1 rune
	var size int

	// leading '_' will appear in output.
	for len(str) > 0 {
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		if r0 != '_' {
			break
		}

		buf.WriteRune(r0)
	}

	if len(str) == 0 {
		return buf.String()
	}

	r0 = unicode.ToUpper(r0)

	for len(str) > 0 {
		r1 = r0
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		if r1 == '_' && r0 == '_' {
			buf.WriteRune(r1)
			continue
		}

		if r1 == '_' {
			r0 = unicode.ToUpper(r0)
		} else {
			r0 = unicode.ToLower(r0)
		}

		if r1 != '_' {
			buf.WriteRune(r1)
		}
	}

	buf.WriteRune(r0)
	return buf.String()
}

func Width(str string) int {
	var r rune
	var size, n int

	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)
		n += RuneWidth(r)
		str = str[size:]
	}

	return n
}

func RuneWidth(r rune) int {
	switch {
	case r == utf8.RuneError || r < '\x20':
		return 0

	case '\x20' <= r && r < '\u2000':
		return 1

	case '\u2000' <= r && r < '\uFF61':
		return 2

	case '\uFF61' <= r && r < '\uFFA0':
		return 1

	case '\uFFA0' <= r:
		return 2
	}

	return 0
}

func SetStructField(structPtr interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(structPtr).Elem()
	fieldValue := structValue.FieldByName(name)

	if !fieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in structPtr", name)
	}

	if !fieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	fieldType := fieldValue.Type()
	val := reflect.ValueOf(value)
	if fieldType != val.Type() && fieldType.Kind() != reflect.Interface {
		return errors.New("Provided value type didn't match structPtr field type")
	}

	fieldValue.Set(val)
	return nil
}

// ChangeStruct applies map of changes to struct
func ChangeStructFields(structPtr interface{}, changesMap map[string]interface{}) {
	for k, v := range changesMap {
		SetField(structPtr, k, v)
	}
}

func SetField(structPtr interface{}, k string, v interface{}) {
	//structPtr[k] = v
}

func TypeName(v interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(v))
}
