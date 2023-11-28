package conf

import (
	"fmt"
)

type nacosConfS struct {
	Name string `yaml:"name" json:"name"`
	Url  string `yaml:"url" json:"url"`
	Ns   string `yaml:"ns" json:"ns"`
	User string `yaml:"user" json:"user"`
	Pwd  string `yaml:"pwd" json:"pwd"`
}

func GetConfYaml() string {
	nacosConf := nacosConfS{
		Name: "rpcNacos",
		Url:  "http://nacosconf:8848/nacos",
		Ns:   "62c3bcf9-7948-4c26-a353-cebc0a7c9712",
		User: "dev",
		Pwd:  "123456",
	}
	yaml := `kitex:  
  client:
    rpc_timeout: 1s #超时设置
    conn_timeout: 1s #连接超时
    max_retry_times: 2 #最大重试次数
    eer_threshold: 10  #重试熔断错误率阈值, 方法级别请求错误率超过阈值则停止重试 10 是指百分之10
    conn_type: long #使用长连接还是短连接
    warmup: true  #是否预热
    warmup_conn_nums: 2 #预热连接数
    pool:
      max_idle_per_address: 10 #表示每个后端实例可允许的最大闲置连接数
      max_idle_global: 1000 #表示全局最大闲置连接数
      max_idle_timeout: 180s #表示连接的闲置时长，超过这个时长的连接会被关闭（最小值 3s，默认值 30s ）
      min_idle_per_address: 2 #对每个后端实例维护的最小空闲连接数，这部分连接即使空闲时间超过 MaxIdleTimeout 也不会被清理。
    service_find:
      type: nacos
  nacos:
    name: %s
    url: %s
    ns: %s
    user: %s
    pwd: %s
  `
	y := fmt.Sprintf(yaml, nacosConf.Name, nacosConf.Url, nacosConf.Ns, nacosConf.User, nacosConf.Pwd)
	return y
}
