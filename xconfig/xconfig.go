package xconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/robjporter/go-xtools/xas"
	"gopkg.in/yaml.v2"
	"time"
)

type expire struct {
	expiry  time.Time
	access  time.Time
	expires time.Duration
}

type Config struct {
	root    interface{}
	Expiry  time.Duration
	expires map[interface{}]*expire
}

// ///////////////////////////////////////////////////
// CONSTRUCTOR
// ///////////////////////////////////////////////////
var cfg = New()

func New(expiry ...time.Duration) *Config {
	c := &Config{root: make(map[interface{}]interface{}), expires: make(map[interface{}]*expire)}
	if expiry != nil {
		c.Expiry = expiry[0]
	} else {
		c.Expiry = 1<<63 - 1
	}
	return c
}

// ///////////////////////////////////////////////////
// PUBLIC KEYS
// ///////////////////////////////////////////////////
func GetKeys() ([]string, error) { return cfg.GetKeys() }
func (cfg *Config) GetKeys() ([]string, error) {
	var keys []string
	tmp := cfg.AllSettings()
	for k := range tmp {
		keys = append(keys, k)
	}
	return keys, nil
}

// ///////////////////////////////////////////////////
// PUBLIC SIZE
// ///////////////////////////////////////////////////
func Size() (int, error) { return cfg.Size() }
func (cfg *Config) Size() (int, error) {
	var keys []string
	tmp := cfg.AllSettings()
	for k := range tmp {
		keys = append(keys, k)
	}
	return len(keys), nil
}

// ///////////////////////////////////////////////////
// PUBLIC READ FILES
// ///////////////////////////////////////////////////
func ReadFiles(files ...string) { cfg.ReadFiles(files...) }
func (cfg *Config) ReadFiles(files ...string) {
	for _, file := range files {
		tmp := new(Config)
		err := tmp.readFile(file)
		if err != nil {
			fmt.Printf("Cannot read config file [%s]: %s\n",
				file, err.Error())
		} else {
			merge(&cfg.root, &tmp.root)
		}
	}
}

// ///////////////////////////////////////////////////
// PUBLIC READ STRING
// ///////////////////////////////////////////////////
func (cfg *Config) ReadString(config string) error {
	tmp := new(Config)
	err := tmp.readBuffer(xas.ToBytes(config))
	if err == nil {
		merge(&cfg.root, &tmp.root)
		return nil
	} else {
		return err
	}
}

// ///////////////////////////////////////////////////
// INTERNAL SUPPORTING READ
// ///////////////////////////////////////////////////
func (cfg *Config) readBuffer(buff []byte) error {
	return yaml.Unmarshal(buff, &cfg.root)
}

func (cfg *Config) readFile(file string) error {
	buff, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return cfg.readBuffer(buff)
}

func (cfg *Config) readEnv(prefix string) error {
	return nil
}

// ///////////////////////////////////////////////////
// PUBLIC SUB
// ///////////////////////////////////////////////////
func Sub(path string) *Config { return cfg.Sub(path) }
func (cfg *Config) Sub(path string) *Config {
	data, err := cfg.Get(path)
	if err != nil {
		fmt.Printf("Failed to get %s\n", path)
		return nil
	}

	ncfg := Config{}
	ncfg.root = data
	return &ncfg
}

func Remove(path string) (bool, error) { return cfg.Remove(path) }
func (cfg *Config) Remove(path string) (bool, error) {
	cfg.expires[path] = nil
	return true, nil
}

func remove(path string, data *interface{}) {
	m, ok := (*data).(map[interface{}]interface{})
	if ok {
		for k, _ := range m {
			if path == k {
				m[path] = nil
			}
		}
	}
}

// ///////////////////////////////////////////////////
// PUBLIC GET
// ///////////////////////////////////////////////////
func Get(path string) (interface{}, error) { return cfg.Get(path) }
func (cfg *Config) Get(path string) (interface{}, error) {
	val, err := get(path, cfg.root)
	if err != nil {
		fmt.Printf("Failed to get path: %s\n", path)
		fmt.Printf("%s\n", err.Error())
	} else {
		n := time.Now()
		cfg.expires[path].access = n
		if cfg.expires[path].expiry.After(n) {
			remove(path, &cfg.root)
		}
		return nil, errors.New("Item has expired.")
	}

	return val, err
}

func GetStringSlice(path string) ([]string, error) { return cfg.GetStringSlice(path) }
func (cfg *Config) GetStringSlice(path string) ([]string, error) {
	tmp, err := cfg.Get(path)
	if err == nil {
		n := time.Now()
		if cfg.expires[path].expiry.After(n) {
			return tmp.([]string), nil
		}
		return nil, errors.New("Key missing or expired.")
	} else {
		return nil, err
	}
}

func GetSlice(path string) ([]interface{}, error) { return cfg.GetSlice(path) }
func (cfg *Config) GetSlice(path string) ([]interface{}, error) {
	tmp, err := cfg.Get(path)
	if err == nil {
		n := time.Now()
		if cfg.expires[path].expiry.After(n) {
			return tmp.([]interface{}), nil
		}
		return nil, errors.New("Key missing or expired.")
	} else {
		return nil, err
	}
}

func GetSliceSize(path string) (int, error) { return cfg.GetSliceSize(path) }
func (cfg *Config) GetSliceSize(path string) (int, error) {
	tmp, err := cfg.Get(path)
	if err == nil {
		n := time.Now()
		if cfg.expires[path].expiry.After(n) {
			return len(tmp.([]interface{})), nil
		}
		return -1, errors.New("Key missing or expired.")
	} else {
		return -1, err
	}
}

func GetStringMapSliceElement(element interface{}) (map[string]string, error) {
	return cfg.GetStringMapSliceElement(element)
}
func (cfg *Config) GetStringMapSliceElement(element interface{}) (map[string]string, error) {
	return element.(map[string]string), nil
}

func GetInterfaceMapSliceElement(element interface{}) (map[interface{}]interface{}, error) {
	return cfg.GetInterfaceMapSliceElement(element)
}
func (cfg *Config) GetInterfaceMapSliceElement(element interface{}) (map[interface{}]interface{}, error) {
	return element.(map[interface{}]interface{}), nil
}

func GetString(path string) string { return cfg.GetString(path) }
func (cfg *Config) GetString(path string) string {
	val, err := cfg.Get(path)
	if err != nil {
		return ""
	} else {
		str, ok := val.(string)
		if !ok {
			return ""
		} else {
			n := time.Now()
			fmt.Println("NOW: ", n)
			fmt.Println("EXPIRY: ", cfg.expires[path].expiry)
			fmt.Println("EXPIRES: ", cfg.expires[path].expires)
			fmt.Println("ACCESS: ", cfg.expires[path].access)
			if cfg.expires[path].expiry.After(n) {
				return str
			}
			return ""
		}
	}
}

func GetInt(path string) int { return cfg.GetInt(path) }
func (cfg *Config) GetInt(path string) int {
	val, err := cfg.Get(path)
	if err != nil {
		return 0
	} else {
		num, ok := val.(int)
		if !ok {
			return 0
		} else {
			n := time.Now()
			if cfg.expires[path].expiry.After(n) {
				return num
			}
			return 0
		}
	}
}

func GetBool(path string) bool { return cfg.GetBool(path) }
func (cfg *Config) GetBool(path string) bool {
	val, err := cfg.Get(path)
	if err != nil {
		fmt.Printf("No such value")
		return false
	} else {
		b, ok := val.(bool)
		if !ok {
			fmt.Printf("Mismatched type")
			return false
		} else {
			n := time.Now()
			if cfg.expires[path].expiry.After(n) {
				return b
			}
			return false
		}
	}
}

func GetFloat(path string) float64 { return cfg.GetFloat(path) }
func (cfg *Config) GetFloat(path string) float64 {
	val, err := cfg.Get(path)
	if err != nil {
		fmt.Printf("No such value")
		return float64(-1)
	} else {
		b, ok := val.(float64)
		if !ok {
			fmt.Printf("Mismatched type")
			return float64(-1)
		} else {
			n := time.Now()
			if cfg.expires[path].expiry.After(n) {
				return b
			}
			return 0.0
		}
	}
}

// ///////////////////////////////////////////////////
// PUBLIC SET
// ///////////////////////////////////////////////////
func Set(path string, value interface{}) { cfg.Set(path, value) }
func (cfg *Config) Set(path string, value interface{}) {
	data := build(path, value)
	merge(&cfg.root, &data)
	n := time.Now()
	exp := n.Add(cfg.Expiry)
	tmp := &expire{expiry: exp, access: n}
	cfg.expires[path] = tmp
}

func SetCustom(path string, value interface{}, expirys time.Duration) { cfg.SetCustom(path, value, expirys) }
func (cfg *Config) SetCustom(path string, value interface{}, expirys time.Duration) {
	data := build(path, value)
	merge(&cfg.root, &data)
	n := time.Now()
	exp := n.Add(expirys)
	tmp := &expire{expiry: exp, access: n}
	cfg.expires[path] = tmp
}

// ///////////////////////////////////////////////////
// PUBLIC OUTPUT
// ///////////////////////////////////////////////////
func AllSettings() map[string]interface{} { return cfg.AllSettings() }
func (cfg *Config) AllSettings() map[string]interface{} {
	all_settings := map[string]interface{}{}
	list(&all_settings, "", cfg.root)
	return all_settings
}

// ///////////////////////////////////////////////////
// PUBLIC SIZE
// ///////////////////////////////////////////////////
func (cfg *Config) Len() int {
	all_settings := cfg.AllSettings()
	return len(all_settings)
}

// ///////////////////////////////////////////////////
// PUBLIC ENVIRONMENT
// ///////////////////////////////////////////////////
func BindEnvs(prefix string) {
	prefix = prefix + "_"
	for _, line := range os.Environ() {
		name := strings.Split(line, "=")[0]
		value := strings.Join(strings.Split(line, "=")[1:], "=")
		if strings.HasPrefix(name, prefix) {
			key := strings.TrimPrefix(name, prefix)
			key = strings.ToLower(key)
			key = strings.Replace(key, "_", ".", -1)
			Set(key, guess(value))
		}
	}
}

// ///////////////////////////////////////////////////
// PRIVATE
// ///////////////////////////////////////////////////
func merge(dst, src *interface{}) {
	src_kind := reflect.TypeOf(*src).Kind()
	dst_kind := reflect.TypeOf(*dst).Kind()

	if dst_kind != src_kind {
		*dst = *src
	}

	switch src_kind {
	case reflect.Int, reflect.String, reflect.Bool:
		*dst = *src
	case reflect.Map:
		src_map, ok := (*src).(map[interface{}]interface{})
		if !ok {
			fmt.Printf("Failed to convert src data to map: %v\n", src)
		}
		dst_map, ok2 := (*dst).(map[interface{}]interface{})
		if !ok2 {
			fmt.Printf("Failed to convert dst data to map: %v\n", dst)
		}

		for k, src_v := range src_map {
			dst_v, ok3 := dst_map[k]
			if ok3 {
				merge(&dst_v, &src_v)
				dst_map[k] = dst_v
			} else {
				dst_map[k] = src_v
			}
		}
	default:
		fmt.Printf("Unknown type kind: %s\n", src_kind.String())
	}
}

func get(path string, data interface{}) (interface{}, error) {
	if path == "" {
		return data, nil
	}
	segs := strings.Split(path, ".")
	seg := segs[0]
	data_map, ok := data.(map[interface{}]interface{})
	if !ok {
		return nil, errors.New("Mismatched type")
	}

	val, ok2 := data_map[seg]
	if !ok2 {
		return nil, errors.New("No such key")
	}

	return get(strings.Join(segs[1:], "."), val)

}

func build(path string, value interface{}) interface{} {
	if path == "" {
		return value
	}

	segs := strings.Split(path, ".")
	seg := segs[len(segs)-1]

	data := make(map[interface{}]interface{})
	data[seg] = value
	return build(strings.Join(segs[0:len(segs)-1], "."), data)
}

func list(result *map[string]interface{}, prefix string, data interface{}) {
	m, ok := data.(map[interface{}]interface{})
	if ok {
		for key, value := range m {
			nprefix := ""
			if prefix == "" {
				nprefix = fmt.Sprintf("%s", key)
			} else {
				nprefix = fmt.Sprintf("%s.%s", prefix, key)
			}
			list(result, nprefix, value)
		}
	} else {
		(*result)[prefix] = data
	}
}

func guess(str string) interface{} {
	nv, err := strconv.Atoi(str)
	if err == nil {
		return nv
	}

	bstr := strings.ToLower(str)
	if bstr == "true" {
		return true
	} else if bstr == "false" {
		return false
	}

	return str
}

// ///////////////////////////////////////////////////
// WRITE OUT
// ///////////////////////////////////////////////////
func (cfg *Config) WriteYaml(filename string) error {
	out, err := yaml.Marshal(cfg.root)
	if err == nil {
		fp, err := os.Create(filename)
		if err == nil {
			defer fp.Close()
			_, err = fp.Write(out)
		}
		return nil
	} else {
		return err
	}
}

func (cfg *Config) WriteJson(filename string) error {
	out, err := json.Marshal(convert(cfg.root))
	fp, err := os.Create(filename)
	if err == nil {
		defer fp.Close()
		_, err = fp.Write(out)
	}
	return nil
}

// ///////////////////////////////////////////////////
// UNUSED
// ///////////////////////////////////////////////////
func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}
