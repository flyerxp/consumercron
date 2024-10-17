package nacos

import (
	"context"
	"fmt"
	"github.com/flyerxp/lib/v2/config"
	"github.com/flyerxp/lib/v2/logger"
	"github.com/flyerxp/lib/middleware/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"net/url"
	"strconv"
	"time"
)

func Listen(commd *cobra.Command, name string) {
	ctx := logger.GetContext(context.Background(), "cron")
	confList := config.GetConf().Nacos
	conf := config.MidNacos{}
	for _, v := range confList {
		if v.Name == name {
			conf = v
		}
	}
	if conf.Url == "" {
		fmt.Println("未找到配置 conf:" + name)
		return
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         conf.Ns, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		Username:            conf.User,
		Password:            conf.Pwd,
		LogDir:              "log",
		CacheDir:            "log/cache",
		LogLevel:            "debug",
	}
	pInfo, _ := url.Parse(conf.Url)
	port, _ := strconv.Atoi(pInfo.Port())
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      pInfo.Hostname(),
			ContextPath: pInfo.Path,
			Port:        uint64(port),
			Scheme:      pInfo.Scheme,
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	ldData := [10]string{
		"redis", "mysql", "pulsar",
		"topic", "elastic", "zookeeper",
		"t3d", "iot", "system", "mqtt",
	}
	for _, v := range ldData {
		err = configClient.ListenConfig(vo.ConfigParam{
			DataId: v,
			Group:  v,
			OnChange: func(namespace, group, dataId, data string) {
				fmt.Println(dataId + "变化值:")
				fmt.Println(data)
				client, e := nacos.GetEngine(ctx, "nacosConf")
				if e != nil {
					logger.AddError(ctx, zap.Error(e))
				}
				key := client.DeleteCache(context.Background(), dataId, group, namespace)
				fmt.Println(key + " 缓存已经删除")
			},
		})
		if err != nil {
			fmt.Println(err)
		}
	}
	time.Sleep(365 * time.Hour)
}
