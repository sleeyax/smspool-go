package utils

import (
	"bytes"
	"mime/multipart"
	"reflect"
)

// ToFormData converts an arbitrary struct to a multipart form data buffer and its content type.
func ToFormData(s interface{}, apiKey string) (*bytes.Buffer, string) {
	var buffer bytes.Buffer

	writer := multipart.NewWriter(&buffer)
	defer writer.Close()

	values := reflect.ValueOf(s)

	for i := 0; i < values.NumField(); i++ {
		field := values.Type().Field(i)
		value := values.Field(i).String()
		if !values.Field(i).IsZero() {
			_ = writer.WriteField(field.Tag.Get("form"), value)
		}
	}

	_ = writer.WriteField("key", apiKey)

	return &buffer, writer.FormDataContentType()
}
