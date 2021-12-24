package main

import (
	"fmt"
	"week09-decode/pkg/decoding"
)

const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_heartSize     = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	//_maxPackSize   = MaxBodySize + int32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_heartOffset  = _seqOffset + _seqSize
)

type ProtoReader struct{}

type Proto struct {
	ver  int32
	op   int32
	seq  int32
	body []byte
}

func main() {

}

func (ProtoReader) read(src []byte) Proto {
	packSize := src[:_headerOffset]
	fmt.Printf("read packSize:%v,from proto", packSize)
	ver := src[_verOffset:_opOffset]
	op := src[_opOffset:_seqOffset]
	seq := src[_seqOffset : _seqOffset+4]
	body := src[_seqOffset+4:]
	p := Proto{ver: decoding.BigEndian.Get32(ver),
		op:   decoding.BigEndian.Get32(op),
		seq:  decoding.BigEndian.Get32(seq),
		body: body}

	return p
}
