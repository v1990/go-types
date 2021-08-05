package url

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

type data struct {
	URL *URL `json:"url" yaml:"url"`
}

func TestURL(t *testing.T) {
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
				URL: MustParse("https://www.baidu.com"),
			},
			txt:       `{"url":"https://www.baidu.com"}`,
			unmarshal: json.Unmarshal,
			marshal:   json.Marshal,
		},
		{
			name: "yaml",
			data: data{
				URL: MustParse("https://www.baidu.com"),
			},
			txt:       "url: https://www.baidu.com\n",
			unmarshal: yaml.Unmarshal,
			marshal:   yaml.Marshal,
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
				assert.Equal(t, tt.data, val)
				if tt.data.URL != nil {
					assert.NotNil(t, val.URL)
					assert.Equal(t, tt.data.URL.Scheme, val.URL.Scheme)
					assert.Equal(t, tt.data.URL.Host, val.URL.Host)
					assert.Equal(t, tt.data.URL.Path, val.URL.Path)
					assert.Equal(t, tt.data.URL.RawQuery, val.URL.RawQuery)
				}
			}
		})
	}
}
