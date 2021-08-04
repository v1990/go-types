package regexp

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

type data struct {
	Regexp *Regexp `json:"regexp,omitempty" yaml:"regexp,omitempty"`
}

func (d *data) String() string {
	if d == nil || d.Regexp == nil || d.Regexp.Regexp == nil {
		return ""
	}
	return fmt.Sprintf("regexp: %s", d.Regexp.String())
}

func TestRegexp(t *testing.T) {
	tests := []struct {
		name      string
		data      data
		txt       string
		unmarshal func(data []byte, v interface{}) error
		marshal   func(v interface{}) ([]byte, error)
		match     map[string]bool // test MatchString after unmarshal
	}{
		{
			name: "json",
			data: data{
				Regexp: MustCompile("1[3-9]\\d{9}"),
			},
			txt:       `{"regexp":"1[3-9]\\d{9}"}`,
			unmarshal: json.Unmarshal,
			marshal:   json.Marshal,
			match: map[string]bool{
				"13123456789": true,
				"134":         false,
				"12345678901": false,
			},
		},
		{
			name: "yaml",
			data: data{
				Regexp: MustCompile("1[3-9]\\d{9}"),
			},
			txt:       "regexp: 1[3-9]\\d{9}\n",
			unmarshal: yaml.Unmarshal,
			marshal:   yaml.Marshal,
			match: map[string]bool{
				"13123456789": true,
				"134":         false,
				"12345678901": false,
			},
		},

		{
			name:      "json_nil",
			data:      data{},
			txt:       "{}",
			unmarshal: json.Unmarshal,
			marshal:   json.Marshal,
			match:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.marshal != nil {
				res, err := tt.marshal(tt.data)
				assert.NoError(t, err)
				assert.Equal(t, tt.txt, string(res))
			}
			if tt.unmarshal != nil {
				var val data
				err := tt.unmarshal([]byte(tt.txt), &val)
				assert.NoError(t, err)
				assert.Equal(t, tt.data.String(), val.String())

				for str, want := range tt.match {
					got := val.Regexp.MatchString(str)
					assert.Equal(t, want, got, fmt.Sprintf("MatchString(%s) want: %v,but got: %v", str, want, got))
				}
			}
		})
	}
}
