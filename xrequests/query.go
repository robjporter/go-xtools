package xrequests

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func (a *Agent) Query(content interface{}) *Agent {
	switch v := reflect.ValueOf(content); v.Kind() {
	case reflect.String:
		a.queryString(v.String())
	case reflect.Struct:
		a.queryStruct(v.Interface())
	case reflect.Map:
		a.queryMap(v.Interface())
	default:
	}
	return a
}

func (a *Agent) queryMap(content interface{}) *Agent {
	return a.queryStruct(content)
}

func (a *Agent) queryString(content string) *Agent {
	var val map[string]string
	if err := json.Unmarshal([]byte(content), &val); err == nil {
		for k, v := range val {
			a.QueryData.Add(k, v)
		}
	} else {
		if QueryData, err := url.ParseQuery(content); err == nil {
			for k, v := range QueryData {
				for _, v2 := range v {
					a.QueryData.Add(k, string(v2))
				}
			}
		} else {
			a.Errors = append(a.Errors, err)
		}
	}
	return a
}

func (a *Agent) queryStruct(content interface{}) *Agent {
	if mc, err := json.Marshal(content); err != nil {
		a.Errors = append(a.Errors, err)
	} else {
		var val map[string]interface{}
		if err := json.Unmarshal(mc, &val); err != nil {
			a.Errors = append(a.Errors, err)
		} else {
			for k, v := range val {
				k = strings.ToLower(k)
				var queryVal string
				switch t := v.(type) {
				case string:
					queryVal = t
				case float64:
					queryVal = strconv.FormatFloat(t, 'f', -1, 64)
				case time.Time:
					queryVal = t.Format(time.RFC3339)
				default:
					j, err := json.Marshal(v)
					if err != nil {
						continue
					}
					queryVal = string(j)
				}
				a.QueryData.Add(k, queryVal)
			}
		}
	}
	return a
}
