package jsonstring

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSONString_MarshalJSON(t *testing.T) {
	type T1 struct {
		S1 JSONString
		S2 string
	}

	tests := []struct {
		name string
		arg  T1
		want string
	}{
		{
			name: "json-string",
			arg: T1{
				S1: `{"name":"go"}`,
				S2: `{"name":"gogo"}`,
			},
			want: `{"S1":{"name":"go"},"S2":"{\"name\":\"gogo\"}"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.arg)
			require.NoError(t, err)
			require.Equal(t, tt.want, string(data))

			var v2 T1
			err = json.Unmarshal(data, &v2)
			require.NoError(t, err)

			require.Equal(t, tt.arg, v2)

			//var v2 APNSMsg
			//err = json.Unmarshal(data, &v2)
			//require.NoError(t, err)
			//require.Equal(t, tt.arg, v2)

		})
	}
}
