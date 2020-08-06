package redis

import (
	"utils/utils"
)

type CompressCoder struct {
}

var _ ValueCoder = (*CompressCoder)(nil)

func (c *CompressCoder) Encoder(v interface{}) (bytes []byte, err error) {
	if bytes, err = utils.StructToBytes(v); err != nil {
		return nil, err
	}
	return utils.ZipBytes(bytes)
}

func (c *CompressCoder) DeCoder(o []byte, v interface{}) error {
	if bytes, err := utils.UZipBytes(o); err == nil {
		return utils.BytesToStruct(bytes, v)
	} else {
		return err
	}
}
