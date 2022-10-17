/**
    @Author:     ZonzeeLi
    @Project:    lb
    @CreateDate: 10/13/2022
    @UpdateDate: xxx
    @Note:       xxxx
**/

package lb

type Balancer interface {
	NewBalancer(...interface{}) error
	DoBalance(...string) (*Instance, error)
	Add(...*Instance) error
}

type Instance struct {
	Host        string
	Port        string
	Weight      Weight
	Connections int
}

func (p *Instance) GetHost() string {
	return p.Host
}

func (p *Instance) GetPort() string {
	return p.Port
}

func (p *Instance) GetWeight() int64 {
	return p.Weight.Weight
}

func (p *Instance) GetAddr() string {
	return p.Host + ":" + p.Port
}
