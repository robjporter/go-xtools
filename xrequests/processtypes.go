package xrequests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"strings"
)

func (a *Agent) processTypeJSON() (string, io.Reader) {
	var contentJson []byte

	if a.BounceToRawString {
		contentJson = []byte(a.RawString)
	} else if len(a.Data) != 0 {
		contentJson, _ = json.Marshal(a.Data)
	} else if len(a.SliceData) != 0 {
		contentJson, _ = json.Marshal(a.SliceData)
	}

	if contentJson != nil {
		return "application/json", bytes.NewReader(contentJson)
	}
	return "", nil
}

func (a *Agent) processTypeForm() (string, io.Reader) {
	var contentForm []byte
	if a.BounceToRawString || len(a.SliceData) != 0 {
		contentForm = []byte(a.RawString)
	} else {
		formData := changeMapToURLValues(a.Data)
		contentForm = []byte(formData.Encode())
	}
	if len(contentForm) != 0 {
		return "application/x-www-form-urlencoded", bytes.NewReader(contentForm)
	}
	return "", nil
}

func (a *Agent) processTypeText() (string, io.Reader) {
	if len(a.RawString) != 0 {
		return "text/plain", strings.NewReader(a.RawString)
	}
	return "", nil
}

func (a *Agent) processTypeXML() (string, io.Reader) {
	if len(a.RawString) != 0 {
		return "application/xml", strings.NewReader(a.RawString)
	}
	return "", nil
}

func (a *Agent) processTypeMultipart() (string, io.Reader) {
	var buf = &bytes.Buffer{}
	var mw = multipart.NewWriter(buf)
	var contentReader io.Reader
	var contentType string

	if a.BounceToRawString {
		fieldName := a.Header.Get("data_fieldname")
		if fieldName == "" {
			fieldName = "data"
		}
		fw, _ := mw.CreateFormField(fieldName)
		fw.Write([]byte(a.RawString))
		contentReader = buf
	}

	if len(a.Data) != 0 {
		formData := changeMapToURLValues(a.Data)
		for key, values := range formData {
			for _, value := range values {
				fw, _ := mw.CreateFormField(key)
				fw.Write([]byte(value))
			}
		}
		contentReader = buf
	}

	if len(a.SliceData) != 0 {
		fieldName := a.Header.Get("json_fieldname")
		if fieldName == "" {
			fieldName = "data"
		}
		h := make(textproto.MIMEHeader)
		fieldName = strings.Replace(strings.Replace(fieldName, "\\", "\\\\", -1), `"`, "\\\"", -1)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"`, fieldName))
		h.Set("Content-Type", "application/json")
		fw, _ := mw.CreatePart(h)
		contentJSON, err := json.Marshal(a.SliceData)
		if err != nil {
			return "", nil
		}
		fw.Write(contentJSON)
		contentReader = buf
	}

	if len(a.FileData) != 0 {
		for _, file := range a.FileData {
			fw, _ := mw.CreateFormFile(file.Fieldname, file.Filename)
			fw.Write(file.Data)
		}
		contentReader = buf
	}

	mw.Close()

	if contentReader != nil {
		contentType = mw.FormDataContentType()
	}

	return contentType, contentReader
}
