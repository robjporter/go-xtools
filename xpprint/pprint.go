package xpprint

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

var (
	defaultPrinter = &Printer{level: -1}
)

type Printer struct {
	level int
}

func (p *Printer) key(v string) {
	indent := ""
	for i := 0; i < p.level; i++ {
		indent += "    "
	}

	fmt.Printf("%s%s: ", indent, v)
}

func (p *Printer) val(v string) {
	fmt.Println(v)
}

func (p *Printer) indent() {
	p.level += 1
}

func (p *Printer) deindent() {
	p.level -= 1
}

func (p *Printer) newline() {
	fmt.Println()
}

func (p *Printer) enter() {
	if p.level > -1 {
		p.newline()
	}
	p.indent()
}

func Pprint(o interface{}) {
	if o == nil {
		defaultPrinter.newline()
		return
	}

	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Slice:
		if s, ok := v.Interface().([]byte); ok {
			defaultPrinter.val(string(s))
			break
		}
		defaultPrinter.enter()
		for i := 0; i < v.Len(); i++ {
			defaultPrinter.key(strconv.Itoa(i))
			Pprint(v.Index(i).Interface())
		}
		defaultPrinter.deindent()

	case reflect.Map:
		defaultPrinter.enter()
		var keys []string
		var kmap = make(map[string]reflect.Value)
		for _, k := range v.MapKeys() {
			keys = append(keys, k.String())
			kmap[k.String()] = k
		}
		sort.Strings(keys)
		for _, k := range keys {
			defaultPrinter.key(k)
			Pprint(v.MapIndex(kmap[k]).Interface())
		}
		defaultPrinter.deindent()

	case reflect.Struct:
		defaultPrinter.enter()
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			t := v.Type().Field(i)
			if s, ok := f.Interface().(fmt.Stringer); ok {
				defaultPrinter.key(t.Name)
				defaultPrinter.val(s.String())
				continue
			}
			defaultPrinter.key(t.Name)
			Pprint(f.Interface())
		}
		defaultPrinter.deindent()

	case reflect.String:
		defaultPrinter.val(v.String())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		defaultPrinter.val(strconv.FormatInt(v.Int(), 10))
	case reflect.Float64, reflect.Float32:

		defaultPrinter.val(strconv.FormatFloat(v.Float(), 'f', -1, 64))
	case reflect.Bool:
		defaultPrinter.val(strconv.FormatBool(v.Bool()))

	default:
		fmt.Printf("Unknown Type: %v", v.Kind())
	}
}
