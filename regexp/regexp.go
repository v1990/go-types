package regexp

import "regexp"

type Regexp struct {
	*regexp.Regexp
}

func MustCompile(str string) *Regexp {
	reg := regexp.MustCompile(str)
	return &Regexp{
		Regexp: reg,
	}
}

func (r *Regexp) UnmarshalText(text []byte) error {
	reg, err := regexp.Compile(string(text))
	if err != nil {
		return err
	}
	r.Regexp = reg
	return nil
}

func (r *Regexp) MarshalText() (text []byte, err error) {
	if r == nil || r.Regexp == nil {
		return nil, nil
	}
	return []byte(r.String()), nil
}
