package xregex

import (
	"regexp"
)

var (
	patterns map[string]string
)

func init() {
	patterns = make(map[string]string)
	addDefaults()
}

func addDefaults() {
	patterns["email_Pattern"] = `([a-zA-Z0-9_\.-]+)@([\da-zA-Z\.-]+)\.([a-zA-Z\.]{2,4})`
	patterns["email_2_Pattern"] = `(\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3})`
	patterns["email_free_Pattern"] = `([a-zA-Z0-9_\.-]+)@(gmail.com|yahoo.com|hotmail.com)`
	patterns["ipv4_Pattern"] = `(?:(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}(?:25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]|[0-9])`
	patterns["ipv4_cidr_Pattern"] = `(([01]?\d?\d|2[0-4]\d|25[0-5])\.){3}([01]?\d?\d|2[0-4]\d|25[0-5])\/(\d{1,2})`
	patterns["ipv6_Pattern"] = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	patterns["https_Pattern"] = `(https:\/\/)([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-])`
	patterns["http_Pattern"] = `(http:\/\/)([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-])`
	patterns["url_Pattern"] = `((http|https|ftp):\/\/)([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-])`
	patterns["html_Pattern"] = `<([a-z1-6]+)([^<]+)*(?:>(.*)<\/\1>| *\/>)`
	patterns["mac_Pattern"] = `(([a-fA-F0-9]{2}[:-]){5}([a-fA-F0-9]{2}))`
	patterns["dollar_Pattern"] = `\$[0-9]*.?[0-9][0-9]`
	patterns["pound_Pattern"] = `(£[0-9]*.?[0-9][0-9])`
	patterns["date_Pattern"] = `(([1-9]|[12][0-9]|3[01])[- /.](0?[1-9]|1[012])[- /.](19|20)?[0-9]{2})`
	patterns["postcode_uk_Pattern"] = `([A-Z]{1,2}[0-9][A-Z0-9]? [0-9][ABD-HJLNP-UW-Z]{2})`
	patterns["password_low_Pattern"] = `([a-zA-Z])(.*\d)[a-zA-Z\d]{4,}`
	patterns["password_med_Pattern"] = ``
	patterns["password_hig_Pattern"] = ``
	patterns["time24_Pattern"] = `4([01]?[0-9]|2[0-3]):[0-5][0-9]`
	patterns["ucs_version_Pattern"] = `(\d){1}.(\d){1}.(\d){1}([a-z]{1})`
	patterns["clean_string"] = `[^a-zA-Z0-9]+`
}

func AddPattern(name string, pattern string) bool {
	if patterns[name] == "" {
		patterns[name] = pattern
		return true
	}
	return false
}

func Pattern(name, text string) []string {
	if patterns[name] != "" {
		return match(text, name)
	}
	return []string{}
}

func match(text string, pattern string) []string {
	if patterns[pattern] == "" {
		return []string{}
	} else {
		compiled := regexp.MustCompile(patterns[pattern])
		parsed := compiled.FindAllString(text, -1)
		return parsed
	}
}

func CleanString(text string) string {
	return match(text, "clean_string")[0]
}

func UCSVersion(text string) []string {
	return match(text, "ucs_version_Pattern")
}

func IP(text string) []string {
	return match(text, "ipv4_Pattern")
}

func IPCidr(text string) []string {
	return match(text, "ipv4_cidr_Pattern")
}

func IPv6(text string) []string {
	return match(text, "ipv6_Pattern")
}

func Email(text string) []string {
	return match(text, "email_2_Pattern")
}

func EmailFree(text string) []string {
	return match(text, "email_free_Pattern")
}

func Https(text string) []string {
	return match(text, "https_Pattern")
}

func Http(text string) []string {
	return match(text, "http_Pattern")
}

func URL(text string) []string {
	return match(text, "url_Pattern")
}

func Time24(text string) []string {
	return match(text, "time24_Pattern")
}

func Html(text string) []string {
	return match(text, "html_Pattern")
}

func Mac(text string) []string {
	return match(text, "mac_Pattern")
}

func Dollar(text string) []string {
	return match(text, "dollar_Pattern")
}

func Pound(text string) []string {
	return match(text, "pound_Pattern")
}

func Date(text string) []string {
	return match(text, "date_Pattern")
}

func PostcodeUK(text string) []string {
	return match(text, "postcode_uk_Pattern")
}

func PasswordLevel(text string) int {
	if PasswordHigh(text) {
		return 3
	} else if PasswordMedium(text) {
		return 2
	} else if PasswordLow(text) {
		return 1
	}
	return 0
}

func PasswordLow(text string) bool {
	result := match(text, "password_low_Pattern")
	if len(result) > 0 {
		return true
	}
	return false
}

func PasswordMedium(text string) bool {
	result := match(text, "password_med_Pattern")
	if len(result) > 0 {
		return true
	}
	return false
}

func PasswordHigh(text string) bool {
	result := match(text, "password_hig_Pattern")
	if len(result) > 0 {
		return true
	}
	return false
}
