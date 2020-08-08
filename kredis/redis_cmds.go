package kredis

import (
	"github.com/go-redis/redis/v7"
	"time"
)

func (r *Redis) SetByCoder(key string, value interface{}, expire time.Duration) error {
	encodeVal := ValueEncode(r.ValueCoder, value)
	err := r.Cmdable.Set(key, encodeVal, expire).Err()
	return err
}

// Get(key string) *StringCmd
func (r *Redis) GetByCoder(key string, val interface{}) error {
	strObj, err := r.Cmdable.Get(key).Bytes()
	if err != nil && err != redis.Nil {
		return err
	}
	return ValueDecode(r.ValueCoder, strObj, val)
}
func (r *Redis) GetRaw(key string) ([]byte, error) {
	c, err := r.Cmdable.Get(key).Bytes()
	if err != nil && err != redis.Nil {
		return []byte{}, err
	}
	return c, nil
}

func (r *Redis) GetSetByCoder(key string, val interface{}, newVal interface{}) error {
	strObj, err := r.Cmdable.GetSet(key, newVal).Bytes()
	if err != nil && err != redis.Nil {
		return err
	}
	return ValueDecode(r.ValueCoder, strObj, val)
}
func (r *Redis) HSetByCoder(key, field string, value interface{}) error {
	encodeVal := ValueEncode(r.ValueCoder, value)
	err := r.Cmdable.HSet(key, field, encodeVal).Err()
	return err
}

func (r *Redis) HGetByCoder(key, field string, val interface{}) error {
	strObj, err := r.Cmdable.HGet(key, field).Bytes()
	if err != nil && err != redis.Nil {
		return err
	}
	return ValueDecode(r.ValueCoder, strObj, val)
}

func (r *Redis) HGetRaw(key string, fields string) ([]byte, error) {
	strObj, err := r.Cmdable.HGet(key, fields).Bytes()
	if err != nil && err != redis.Nil {
		return []byte{}, err
	}
	if err == redis.Nil {
		return []byte{}, nil
	}
	return strObj, nil
}
