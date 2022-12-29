package enum

import (
	"bytes"
	"strings"
	"unicode"
)

func Convert(name string) string {
	return Camel2Space(name)
}

func Camel2Space(name string) string {
	// 全大写保留原格式
	if strings.ToUpper(name) == name {
		return name
	}

	buffer := bytes.Buffer{}
	for i, c := range name {
		if i == 0 {
			buffer.WriteRune(unicode.ToLower(c))
			continue
		}
		if unicode.IsUpper(c) {
			buffer.WriteRune(0x20)
		}
		buffer.WriteRune(unicode.ToLower(c))
	}

	return buffer.String()
}
