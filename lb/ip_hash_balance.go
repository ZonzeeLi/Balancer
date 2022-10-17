package lb

import (
	"errors"
	"fmt"
	"hash/crc32"
	"math/rand"
)

type IpHashBalance struct {
	insts []*Instance
}

func init() {
	RegisterBalancer("hash", &IpHashBalance{})
}

func (p *IpHashBalance) NewBalancer(param ...interface{}) (err error) {
	return
}

func (p *IpHashBalance) Add(insts ...*Instance) (err error) {
	if len(insts) <= 0 {
		return errors.New("no instance")
	}
	for _, v := range insts {
		p.insts = append(p.insts, v)
	}
	return
}

func (p *IpHashBalance) DoBalance(ip ...string) (inst *Instance, err error) {
	insts := p.insts
	var def = fmt.Sprintf("%d", rand.Int()) // 默认，应该取一个ip地址
	if len(ip) > 0 {
		def = ip[0]
	}
	lens := len(insts)
	if lens == 0 {
		err = errors.New("no instances")
		return
	}
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(def), crcTable) // 计算hash值
	index := int(hashVal) % lens
	inst = insts[index]
	return
}
