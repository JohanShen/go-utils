// Copyright 2020 Douyu
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kredis

import (
	"errors"
	"github.com/go-redis/redis/v7"
)

//Redis client (cmdable and config)
type Redis struct {
	redis.Cmdable
	Config     *Config
	ValueCoder ValueCoder
}

func (self *Redis) UseValueCoder(coder ValueCoder) *Redis {
	_redis := &Redis{
		Config:     self.Config,
		Cmdable:    self.Cmdable,
		ValueCoder: coder,
	}
	return _redis
}

func ValueEncode(coder ValueCoder, v interface{}) interface{} {
	if coder == nil {
		return v
	}
	if bytes, err := coder.Encoder(v); err == nil {
		return bytes
	}
	return v
}

func ValueDecode(coder ValueCoder, o []byte, v interface{}) error {
	if coder == nil {
		return errors.New("ValueDecoder can't make nil values")
	}
	return coder.DeCoder(o, v)
}
