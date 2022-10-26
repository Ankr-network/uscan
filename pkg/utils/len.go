package utils

const (
	FullLen = 0xFF
)

func WrapLen(data []byte) []byte {
	var bl byte
	l := len(data)

	if l > FullLen {
		bl = FullLen
	} else {
		bl = byte(l)
	}

	ds := make([]byte, l+1)
	ds[0] = bl
	copy(ds[1:], data)
	return ds
}
