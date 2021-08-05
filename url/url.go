package url

import "net/url"

type URL struct {
	*url.URL
}

func Parse(rawurl string) (*URL, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	return &URL{URL: u}, nil
}

func MustParse(rawurl string) *URL {
	u, err := Parse(rawurl)
	if err != nil {
		panic(err)
	}
	return u
}

func (u *URL) UnmarshalText(text []byte) (err error) {
	u.URL, err = url.Parse(string(text))
	return
}

func (u *URL) MarshalText() ([]byte, error) {
	if u == nil || u.URL == nil {
		return nil, nil
	}
	return []byte(u.URL.String()), nil
}
