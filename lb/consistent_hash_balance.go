/**
    @Author:     ZonzeeLi
    @Project:    balance
    @CreateDate: 10/17/2022
    @UpdateDate: xxx
    @Note:       xxxx
**/

package lb

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type ConsistentHashBalance struct {
	insts []*Instance
	chash *ConsistentHash
}

func init() {
	RegisterBalancer("ConsistentHash", &ConsistentHashBalance{})
}

type ConsistentHash struct {
	mux      sync.RWMutex
	hash     Hash                 //hash函数
	replicas int                  //复制因子
	keys     Uint32Slice          //已排序的hash节点切片
	hashmap  map[uint32]*Instance //key为hash值val为节点
}

func (p *ConsistentHashBalance) NewBalancer(param ...interface{}) (err error) {
	if len(param) < 2 {
		err = errors.New("not enough params for ConsistentHashBalance")
		return
	}
	hash, ok := param[0].(Hash)
	if !ok {
		err = errors.New("invalid params")
		return
	}
	replicas := param[1].(int)
	if ok {
		err = errors.New("invalid params")
		return
	}
	if p.chash == nil {
		p.chash = &ConsistentHash{
			hash:     hash,
			replicas: replicas,
			hashmap:  make(map[uint32]*Instance),
			keys:     make([]uint32, 0, 100),
		}
		if p.chash.hash == nil {
			p.chash.hash = crc32.ChecksumIEEE
		}
	}
	return
}

func (p *ConsistentHashBalance) Add(insts ...*Instance) (err error) {
	c := p.chash
	if c == nil {
		return errors.New("invalid balancer")
	}
	if len(insts) <= 0 {
		return errors.New("no instance")
	}
	for _, v := range insts {
		p.insts = append(p.insts, v)
		addr := v.GetAddr()
		//计算虚拟节点hash值
		for i := 0; i < c.replicas; i++ {
			hash := c.hash([]byte(strconv.Itoa(i) + addr))
			//实现了排序接口
			c.keys = append(c.keys, hash)
			c.hashmap[hash] = v
		}
	}
	sort.Sort(c.keys)
	return
}

func (p *ConsistentHashBalance) DoBalance(key ...string) (inst *Instance, err error) {
	c := p.chash
	if len(c.keys) == 0 {
		return
	}
	hash := c.hash([]byte(key[0]))
	//通过二分查询到最优节点（第一个hash大于资源hash的服务器）
	idx := sort.Search(len(c.keys), func(i int) bool {
		return c.keys[i] > hash
	})
	if idx == len(c.keys) {
		//没有找到服务器，说明此时处于环的尾部，那么直接用第0台服务器
		idx = 0
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	inst = c.hashmap[c.keys[idx]]
	return
}

type Hash func([]byte) uint32

type Uint32Slice []uint32

func (s Uint32Slice) Len() int {
	return len(s)
}

func (s Uint32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Uint32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
