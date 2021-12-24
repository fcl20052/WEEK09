package decoding

var BigEndian bigEndian

type bigEndian struct{}

func (bigEndian) Get32(b []byte) int32 {
	b1 := int32(b[0]) << 24
	b2 := int32(b[1]) << 16
	b3 := int32(b[2]) << 8
	b4 := int32(b[3])
	return b1+b2+b3+b4

}
