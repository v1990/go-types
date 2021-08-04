package bytesize

import (
	"github.com/alecthomas/units"
)

type Base2Bytes units.Base2Bytes

const (
	Kibibyte = Base2Bytes(units.Kibibyte)
	KiB      = Base2Bytes(units.KiB)
	Mebibyte = Base2Bytes(units.Mebibyte)
	MiB      = Base2Bytes(units.MiB)
	Gibibyte = Base2Bytes(units.Gibibyte)
	GiB      = Base2Bytes(units.GiB)
	Tebibyte = Base2Bytes(units.Tebibyte)
	TiB      = Base2Bytes(units.TiB)
	Pebibyte = Base2Bytes(units.Pebibyte)
	PiB      = Base2Bytes(units.PiB)
	Exbibyte = Base2Bytes(units.Exbibyte)
	EiB      = Base2Bytes(units.EiB)
)

func (b Base2Bytes) Base2Bytes() units.Base2Bytes {
	return units.Base2Bytes(b)
}

func (b Base2Bytes) String() string {
	return b.Base2Bytes().String()
}

func (b *Base2Bytes) UnmarshalText(text []byte) error {
	bb, err := units.ParseBase2Bytes(string(text))
	if err != nil {
		return err
	}

	*b = Base2Bytes(bb)
	return nil
}

func (b Base2Bytes) MarshalText() (text []byte, err error) {
	return []byte(b.String()), nil
}
