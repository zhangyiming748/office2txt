package stand

import (
	"bytes"
	"encoding/xml"
	"io"
)

/*
我输入的字节流包括xml的控制语句，我需要删除这些语句，只保留xml中的正文内容。
*/
func Clean(text []byte) []byte {
	decoder := xml.NewDecoder(bytes.NewReader(text))
	var result bytes.Buffer

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return text
		}

		switch t := token.(type) {
		case xml.CharData:
			result.Write(t)
		}
	}

	return result.Bytes()
}
