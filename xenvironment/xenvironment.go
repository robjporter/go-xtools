package xenvironment

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	Prefix = ""
	debug  = os.Getenv("BUILDDEBUG") != ""
)

func GetUsername() string {
	u, e := user.Current()
	if e != nil {
		return ""
	}
	return u.Username
}

func GOPATHBIN() string {
	return os.Getenv("GOPATH") + PathSeparator() + "bin"
}

func GetAllEnvironment() map[string]string {
	getenvironment := func(data []string, getkeyval func(item string) (key, val string)) map[string]string {
		items := make(map[string]string)
		for _, item := range data {
			key, val := getkeyval(item)
			items[key] = val
		}
		return items
	}
	return getenvironment(os.Environ(), func(item string) (key, val string) {
		splits := strings.Split(item, "=")
		key = splits[0]
		val = strings.Join(splits[1:], "=")
		return
	})
}

func PathSeparator() string {
	return fmt.Sprintf("%c", os.PathSeparator)
}

func ListSeparator() string {
	return fmt.Sprintf("%c", os.PathListSeparator)
}

func IsCompiled() bool {
	if strings.HasPrefix(os.Args[0], "/var/folders/") ||
		strings.HasPrefix(os.Args[0], "/tmp/go-build") ||
		strings.Contains(os.Args[0], "\\AppData\\Local\\Temp\\") {
		return false
	}
	return true
}

func BuildDebug() bool {
	return debug
}

func CheckArchitecture() bool {
	switch runtime.GOARCH {
	case "386", "amd64":
		return true
	case "arm64", "arm", "ppc64", "ppc64le", "mips", "mipsle":
		fmt.Println("This is an untested architecture: %q; proceed with caution!", runtime.GOARCH)
		return false
	default:
		fmt.Printf("Unknown goarch %q; proceed with caution!", runtime.GOARCH)
		return false
	}
}

func BuildStamp() int64 {
	if s, _ := strconv.ParseInt(os.Getenv("SOURCE_DATE_EPOCH"), 10, 64); s > 0 {
		return s
	}
	bs, err := runError("git", "show", "-s", "--format=%ct")
	if err != nil {
		return time.Now().Unix()
	}
	s, _ := strconv.ParseInt(string(bs), 10, 64)
	return s
}

func runError(cmd string, args ...string) ([]byte, error) {
	if debug {
		t0 := time.Now()
		fmt.Println("runError:", cmd, strings.Join(args, " "))
		defer func() {
			fmt.Println("... in", time.Since(t0))
		}()
	}
	ecmd := exec.Command(cmd, args...)
	bs, err := ecmd.CombinedOutput()
	return bytes.TrimSpace(bs), err
}

func BuildHost() string {
	if v := os.Getenv("BUILD_HOST"); v != "" {
		return v
	}

	h, err := os.Hostname()
	if err != nil {
		return "unknown-host"
	}
	return h
}

func Compiler() string {
	return runtime.Compiler
}

func GOARCH() string {
	return runtime.GOARCH
}

func GOOS() string {
	return runtime.GOOS
}

func GOROOT() string {
	return runtime.GOROOT()
}

func GOVER() string {
	return runtime.Version()
}

func NumCPU() int {
	return runtime.NumCPU()
}

func GOPATH() string {
	return os.Getenv("GOPATH")
}

func GetFormattedTime() string {
	return Now("Monday, 2 Jan 2006")
}

func Now(layout string) string {
	return time.Now().Format(layout)
}

func GetEnv(n string, def interface{}) interface{} {
	value := os.Getenv(prefixedName(n))
	if value == "" {
		return def
	}
	return value
}

//GetString returns a environment variable as string
func GetEnvString(n string, def string) string {
	return GetEnv(n, def).(string)
}

//GetBool returns a environment variable as bool
func GetEnvBool(n string, def bool) bool {
	return GetEnv(n, def).(bool)
}

//GetInt returns a environment variable as int
func GetEnvInt(n string, def int) int {
	return GetEnv(n, def).(int)
}

//GetFloat returns a environment variable as float
func GetEnvFloat(n string, def float64) float64 {
	return GetEnv(n, def).(float64)
}

//prefixedName returns the prefixed name
func prefixedName(s string) string {
	if Prefix == "" {
		return s
	}
	return fmt.Sprintf("%s_%s", Prefix, s)
}
