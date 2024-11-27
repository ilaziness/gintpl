package server

import (
	"fmt"

	"gintpl/pkg/config"
	"gintpl/pkg/hook"
	"gintpl/pkg/log"
	"gintpl/pkg/utils"
	"github.com/jinzhu/copier"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

// 提供服务注册和获取功能
// Register：注册服务
// GetInstance: 获取健康服务的ip和端口

var (
	defaultNameService *nameService
)

type nameService struct {
	nacosClient naming_client.INamingClient
	dataID      string
	group       string
	ip          string
	name        string
	weight      int
	port        uint16
}

// Register 注册服务
func Register(cfg *config.Nacos, appCfg *config.App) {
	weight := 1
	if appCfg.Weight > 0 {
		weight = appCfg.Weight
	}
	nacos := &nameService{
		name:   appCfg.ID,
		weight: weight,
		port:   appCfg.Port,
		dataID: cfg.DataID,
		group:  cfg.Group,
	}

	cc := constant.ClientConfig{}
	if err := copier.Copy(&cc, cfg.Client); err != nil {
		panic(err)
	}
	cc.NamespaceId = cfg.Client.NamespaceID

	var err error
	var sc []constant.ServerConfig
	for _, v := range cfg.Server {
		sc = append(sc, constant.ServerConfig{
			IpAddr: v.IP,
			Port:   uint64(v.Port),
		})
	}

	nacos.nacosClient, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	defaultNameService = nacos
	defaultNameService.register()

	hook.Exit.Register(func() {
		defaultNameService.deregister()
	})
}

// GetInstance 获取服务得地址和端口
func GetInstance(name string) (string, error) {
	instance, err := defaultNameService.nacosClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
	})
	if err != nil {
		return "", err
	}
	log.Logger.Infof("%v", instance)
	return fmt.Sprintf("%s:%d", instance.Ip, instance.Port), nil
}

func (n *nameService) register() {
	ip, err := utils.GetInternalIP()
	if err != nil {
		panic(err)
	}
	ok, err := n.nacosClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        uint64(n.port),
		ServiceName: n.name,
		Weight:      float64(n.weight),
		Enable:      true,
		Healthy:     true,
	})
	if err != nil {
		panic(err)
	}
	if !ok {
		log.Logger.Error("register instance fail")
	}
	n.ip = ip
}

// Deregister 注销服务
func (n *nameService) deregister() {
	_, err := defaultNameService.nacosClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          n.ip,
		Port:        uint64(n.port),
		ServiceName: n.name,
	})
	if err != nil {
		log.Logger.Warnf("deregister instance %s[%s] fail", n.name, n.ip)
	}
}
