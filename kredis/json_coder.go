package kredis

import "encoding/json"

type JsonCoder struct {
}

var _ ValueCoder = (*JsonCoder)(nil)

func (c *JsonCoder) Encoder(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *JsonCoder) DeCoder(o []byte, t interface{}) error {
	return json.Unmarshal(o, t)
}
