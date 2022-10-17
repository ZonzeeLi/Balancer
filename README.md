## Balancer_Selector

Go语言实现了一个负载均衡的管理器，支持服务器节点信息、balancer配置、实例信息、实例方法及选择器的自定义。目前支持的负载均衡方法有随机、轮询、加权轮询、iphash、一致性hash及最小连接数，同时也包含nginx做代理层以grpc通信的负载均衡，