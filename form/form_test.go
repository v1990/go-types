package form

import (
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

func TestTransform(t *testing.T) {
	type T1 struct {
		ID   uint64 `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Addr string `json:"addr,omitempty"`
	}

	tests := []struct {
		name   string
		obj    interface{}
		values url.Values
	}{
		{
			name: "base",
			obj: T1{
				ID:   1234,
				Name: "hello",
				Addr: "hk",
			},
			values: map[string][]string{
				"id":   {"1234"},
				"name": {"hello"},
				"addr": {"hk"},
			},
		},
		{
			name: "omitempty",
			obj: T1{
				ID:   1234,
				Name: "hello",
			},
			values: map[string][]string{
				"id":   {"1234"},
				"name": {"hello"},
				//"addr": {"hk"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form, err := Transform(tt.obj)
			require.NoError(t, err)
			require.Equal(t, tt.values, form)
		})
	}

}
