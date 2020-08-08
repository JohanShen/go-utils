package kredis

type ValueCoder interface {
	// 将对象转成字节组
	Encoder(interface{}) ([]byte, error)
	// 将字节组转成对象
	DeCoder([]byte, interface{}) error
}
