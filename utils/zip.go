package utils

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
)

type Closer interface {
	Close() error
}

func ioClose(obj Closer) {
	if obj != nil {
		if err := obj.Close(); err != nil {
			fmt.Println(err)
		}
	}
}

//zip压缩
func ZipBytes(data []byte) ([]byte, error) {

	var in bytes.Buffer
	//z:=zlib.NewWriter(&in)
	z, err := zlib.NewWriterLevel(&in, zlib.DefaultCompression)
	if err != nil {
		return []byte{}, err
	}
	_, err = z.Write(data)
	if err != nil {
		return []byte{}, err
	}
	ioClose(z)
	return in.Bytes(), nil
}

//zip解压
func UZipBytes(data []byte) ([]byte, error) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write(data)
	r, err := zlib.NewReader(&in)
	if err != nil {
		return []byte{}, err
	}

	_, err = io.Copy(&out, r)
	if err != nil {
		return []byte{}, err
	}
	ioClose(r)
	return out.Bytes(), nil
}

//压缩
func GZipBytes(data []byte) ([]byte, error) {
	var input bytes.Buffer
	//g := gzip.NewWriter(&input)
	g, err := gzip.NewWriterLevel(&input, gzip.DefaultCompression)
	if err != nil {
		return []byte{}, err
	}
	_, err = g.Write(data)
	if err != nil {
		return []byte{}, err
	}
	ioClose(g)
	return input.Bytes(), nil
}

//解压
func UGZipBytes(data []byte) ([]byte, error) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write(data)
	r, err := gzip.NewReader(&in)
	if err != nil {
		return []byte{}, err
	}
	_, err = io.Copy(&out, r)
	if err != nil {
		return []byte{}, err
	}
	ioClose(r)
	return out.Bytes(), nil
}
