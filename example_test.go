package types

import (
	"encoding/json"
	"fmt"
	"github.com/v1990/go-types/bytesize"
	"github.com/v1990/go-types/duration"
	"github.com/v1990/go-types/regexp"
	"github.com/v1990/go-types/url"
)

func ExampleUnmarshal() {
	type Config struct {
		Duration duration.Duration   `json:"duration"`
		Size     bytesize.Base2Bytes `json:"size"`
		Regexp   *regexp.Regexp      `json:"regexp"`
		URL      *url.URL            `json:"url"`
	}
	jsonStr := `{
		"duration": "1h",
		"size": "1GB",
		"Regexp": "1[3-9]\\d{9}",
		"url": "https://www.google.com/search?q=hello"
	}`

	var c Config
	if err := json.Unmarshal([]byte(jsonStr), &c); err != nil {
		panic(err)
	}

	fmt.Println("duration:", c.Duration.String())
	fmt.Println("size:", c.Size.String())
	fmt.Println("regexp:", c.Regexp.String(), c.Regexp.MatchString("13123456789"))
	fmt.Println("url:", c.URL.String(), c.URL.Host, c.URL.Path, c.URL.RawQuery)

	// Output:
	// duration: 1h0m0s
	// size: 1GiB
	// regexp: 1[3-9]\d{9} true
	// url: https://www.google.com/search?q=hello www.google.com /search q=hello

}
