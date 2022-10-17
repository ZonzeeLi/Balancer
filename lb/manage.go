//管理所有的负载均衡算法
package lb

import "fmt"

type Selector struct {
	allBalancer map[string]Balancer
}

var selector = Selector{
	allBalancer: make(map[string]Balancer),
}

func (p *Selector) registerBalancer(name string, b Balancer) {
	p.allBalancer[name] = b
}

func RegisterBalancer(name string, b Balancer) {
	selector.registerBalancer(name, b)
}

// Select 选择LB
func Select(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := selector.allBalancer[name]
	if !ok {
		err = fmt.Errorf("not found %s balancer", name)
		return
	}
	err = balancer.Add(insts...)
	if err != nil {
		return nil, err
	}
	inst, err = balancer.DoBalance()
	selector.allBalancer[name] = balancer // 更新
	return
}
