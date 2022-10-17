/**
    @Author:     ZonzeeLi
    @Project:    lb
    @CreateDate: 10/13/2022
    @UpdateDate: xxx
    @Note:       xxxx
**/

package lb

import "errors"

type WeightBalance struct {
	insts []*Instance
}

type Weight struct {
	Weight          int64
	curWeight       int64
	effectiveWeight int64
	totalWeight     int64
}

func init() {
	RegisterBalancer("weight", &WeightBalance{})
}

func (p *WeightBalance) NewBalancer(param ...interface{}) (err error) {
	return
}

func (p *WeightBalance) Add(insts ...*Instance) (err error) {
	if len(insts) <= 0 {
		return errors.New("no instance")
	}
	for _, v := range insts {
		p.insts = append(p.insts, v)
	}
	return
}

func (p *WeightBalance) DoBalance(weight ...string) (inst *Instance, err error) {
	insts := p.insts
	var total int64
	for k, v := range insts {
		w := v.Weight
		total += w.curWeight
		insts[k].Weight.curWeight += insts[k].Weight.effectiveWeight
		if w.effectiveWeight < w.curWeight {
			insts[k].Weight.effectiveWeight++
		}
		if inst == nil || inst.Weight.curWeight < w.curWeight {
			inst = insts[k]
		}
	}
	inst.Weight.curWeight -= total
	return
}
