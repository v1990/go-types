package bytesize

import (
	"encoding/json"
	"github.com/alecthomas/units"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"testing"
)

type data struct {
	Size     Base2Bytes   `json:"size,omitempty" yaml:"size,omitempty"`
	SizeList []Base2Bytes `json:"size_list,omitempty" yaml:"size_list,omitempty,flow"`
}

func TestBase2Bytes(t *testing.T) {
	tests := []struct {
		name      string
		data      data
		txt       string
		unmarshal func(data []byte, v interface{}) error
		marshal   func(v interface{}) ([]byte, error)
	}{

		{
			name: "json",
			data: data{
				Size:     10 * KiB,
				SizeList: []Base2Bytes{KiB, MiB, GiB, TiB, PiB, EiB},
			},
			txt:       `{"size":"10KiB","size_list":["1KiB","1MiB","1GiB","1TiB","1PiB","1EiB"]}`,
			unmarshal: json.Unmarshal,
			marshal:   json.Marshal,
		},
		{
			name: "json_unmarshal",
			data: data{
				Size:     5 * MiB,
				SizeList: []Base2Bytes{KiB, MiB, GiB, TiB, PiB, EiB},
			},
			txt:       `{"size":"5MB","size_list":["1KB","1MB","1GB","1TB","1PB","1EB"]}`,
			unmarshal: json.Unmarshal,
			//marshal:   json.Marshal,
		},
		{
			name: "yaml",
			data: data{
				Size:     10 * KiB,
				SizeList: []Base2Bytes{KiB, MiB, GiB, TiB, PiB, EiB},
			},
			txt:       "size: 10KiB\nsize_list: [1KiB, 1MiB, 1GiB, 1TiB, 1PiB, 1EiB]\n",
			unmarshal: yaml.Unmarshal,
			marshal:   yaml.Marshal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.marshal != nil {
				s, err := tt.marshal(tt.data)
				require.NoError(t, err)
				require.Equal(t, tt.txt, string(s))
			}
			if tt.unmarshal != nil {
				var d data
				err := tt.unmarshal([]byte(tt.txt), &d)
				require.NoError(t, err)
				require.Equal(t, d, tt.data)
			}
		})
	}
}
func TestParseBase2Bytes(t *testing.T) {
	tests := []struct {
		txt string
		n   Base2Bytes
	}{
		{"1MB", 1 * MiB},
		{"500KB", 500 * KiB},
		{"0.5MB", 512 * KiB}, // NOTE: NOT 500KB
		{"1.5MB", 1*MiB + 512*KiB},
		{"1MB1KB25B", 1*MiB + 1*KiB + 25},
	}
	for _, tt := range tests {
		t.Run(tt.txt, func(t *testing.T) {
			n, err := units.ParseBase2Bytes(tt.txt)
			assert.NoError(t, err)
			assert.Equal(t, tt.n, Base2Bytes(n))
		})
	}
}
