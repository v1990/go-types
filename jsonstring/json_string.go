package jsonstring

import (
	"bytes"
	"encoding/json"
)

// JSONString 标识这是一个json格式的string，无需转义
//  否则会被当成普通的string，会被加上双引号并转义里面的内容
type JSONString string

func (s *JSONString) UnmarshalJSON(data []byte) error {
	data = bytes.Trim(data, `"`) // 可能带有双引号
	*s = JSONString(data)
	return nil
}
func (s JSONString) MarshalJSON() ([]byte, error) {
	if len(s) == 0 {
		return nil, nil
	}
	// 是一个 json 格式的 string
	if s[0] == '{' {
		return []byte(s), nil
	}
	// 普通string，需要转义
	return json.Marshal(string(s))
}
