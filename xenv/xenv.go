package xenv

import (
	"os"
	"strings"
	"strconv"
)

type Env struct {
	vars map[string]string
}

func New() *Env {
	return &Env{vars: make(map[string]string)}
}

func (e *Env) Size() int {
	return len(e.vars)
}

func (e *Env) Add(key, value string) {
	if os.Getenv(key) == "" {
		e.vars[key] = value
		return
	}
	e.vars[key] = os.Getenv(key)
}

func (e *Env) GetAll() map[string]string {
	return e.vars
}

func (e *Env) Get(key string) string {
	if e.vars[key] == "" {
		v := os.Getenv(key)
		if v != "" {
			e.vars[key] = v
			return e.vars[key]
		}
		return ""
	}
	return e.vars[key]
}

func (e *Env) Set(key, value string) error {
	e.vars[key] = value
	return os.Setenv(key, value)
}

func (e *Env) GetAsBool(key string) bool {
	val := strings.ToLower(strings.TrimSpace(e.Get(key)))
	switch val {
	case "0", "no", "false":
		return false
	default:
		if val != "" {
			return true
		}
		return false
	}
}

func (e *Env) GetAsInt(key string) int {
	i, _ := strconv.Atoi(strings.ToLower(strings.TrimSpace(e.Get(key))))
	return i
}
