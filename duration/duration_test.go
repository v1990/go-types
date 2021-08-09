package duration

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"testing"
	"time"
)

type data struct {
	Timeout   Duration   `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	Intervals []Duration `json:"intervals,omitempty" yaml:"intervals,omitempty,flow"`
}

func TestDuration(t *testing.T) {
	t.Run("json_unmarshal_number", func(t *testing.T) {
		var d data
		var s = fmt.Sprintf(`{"timeout":%d}`, time.Second)
		err := json.Unmarshal([]byte(s), &d)
		require.NoError(t, err)
		require.Equal(t, time.Second, d.Timeout.Duration())
	})

	tests := []struct {
		name             string
		data             data
		txt              string
		unmarshal        func(data []byte, v interface{}) error
		wantUnmarshalErr error
		marshal          func(v interface{}) ([]byte, error)
	}{
		{
			name: "json",
			data: data{
				Timeout: Second,
			},
			txt:       `{"timeout":"1s"}`,
			unmarshal: json.Unmarshal,
			marshal:   json.Marshal,
		},
		{
			name: "json_array",
			data: data{
				Timeout:   Second,
				Intervals: []Duration{Nanosecond, Microsecond, Millisecond, Second, Minute, Hour},
			},
			txt:       `{"timeout":"1s","intervals":["1ns","1µs","1ms","1s","1m0s","1h0m0s"]}`,
			unmarshal: json.Unmarshal,
			marshal:   json.Marshal,
		},
		{
			name: "yaml",
			data: data{
				Timeout: Second,
			},
			txt:       "timeout: 1s\n",
			unmarshal: yaml.Unmarshal,
			marshal:   yaml.Marshal,
		},
		{
			name: "yaml_array",
			data: data{
				Timeout:   Second,
				Intervals: []Duration{Nanosecond, Microsecond, Millisecond, Second, Minute, Hour},
			},
			txt:       "timeout: 1s\nintervals: [1ns, 1µs, 1ms, 1s, 1m0s, 1h0m0s]\n",
			unmarshal: yaml.Unmarshal,
			marshal:   yaml.Marshal,
		},
		{
			name:             "json_unmarshal_error",
			txt:              `{"timeout":"BadDuration"}`,
			unmarshal:        json.Unmarshal,
			wantUnmarshalErr: fmt.Errorf(`time: invalid duration "BadDuration"`),
		},
		{
			name:             "yaml_unmarshal_error",
			txt:              `timeout: BadDuration`,
			unmarshal:        yaml.Unmarshal,
			wantUnmarshalErr: fmt.Errorf(`time: invalid duration "BadDuration"`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.marshal != nil {
				body, err := tt.marshal(tt.data)
				require.NoError(t, err)
				require.Equal(t, tt.txt, string(body))
			}

			if tt.unmarshal != nil {
				var d2 data
				err := tt.unmarshal([]byte(tt.txt), &d2)

				if tt.wantUnmarshalErr == nil {
					require.NoError(t, err)
					require.Equal(t, tt.data, d2)
				} else {
					require.Error(t, err)
					require.Equal(t, tt.wantUnmarshalErr.Error(), err.Error())
				}
			}
		})
	}

}
