package duration

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	Nanosecond  = Duration(time.Nanosecond)
	Microsecond = Duration(time.Microsecond)
	Millisecond = Duration(time.Millisecond)
	Second      = Duration(time.Second)
	Minute      = Duration(time.Minute)
	Hour        = Duration(time.Hour)
)

type Duration time.Duration

// Duration cast to time.Duration
func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}

// String ...
func (d Duration) String() string {
	return d.Duration().String()
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}
func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	return d.setValue(v)
}

func (d Duration) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

//func (d *Duration) UnmarshalText(b []byte) error {
//	var v interface{}
//	if err := json.Unmarshal(b, &v); err != nil {
//		return err
//	}
//	return d.setValue(v)
//}

func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v interface{}
	if err := unmarshal(&v); err != nil {
		return err
	}
	return d.setValue(v)
}

func (d *Duration) setValue(v interface{}) error {
	var dd time.Duration
	switch v := v.(type) {
	case float64:
		dd = time.Duration(v)
	case string:
		var err error
		dd, err = time.ParseDuration(v)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid duration")
	}

	*d = Duration(dd)

	return nil
}
