package config

import (
	"bytes"

	"gintpl/pkg/hook"
	"gintpl/pkg/log"
	"github.com/jinzhu/copier"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

// InitNacos 初始化nacos配置中心，cfg是nacos配置，appCfg是应用配置对象，从nacos拉取到的配置会解析到appCfg
func InitNacos[T any](cfg *Nacos, appCfg T) {
	nacos := &nacosConfigType[T]{
		appConfig: appCfg,
		dataID:    cfg.DataID,
		group:     cfg.Group,
	}

	cc := constant.ClientConfig{}
	if err := copier.Copy(&cc, cfg.Client); err != nil {
		panic(err)
	}
	cc.NamespaceId = cfg.Client.NamespaceID

	var sc []constant.ServerConfig
	for _, v := range cfg.Server {
		sc = append(sc, constant.ServerConfig{
			IpAddr: v.IP,
			Port:   uint64(v.Port),
		})
	}

	cfc, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	nacos.nacosConfigClient = cfc
	nacos.get()

	hook.Exit.Register(nacos.cancelListen)
}

type nacosConfigType[T any] struct {
	nacosConfigClient config_client.IConfigClient
	dataID            string
	group             string
	appConfig         T
}

func (nc *nacosConfigType[T]) get() {
	content, err := nc.nacosConfigClient.GetConfig(vo.ConfigParam{
		DataId: nc.dataID,
		Group:  nc.group,
	})
	if err != nil {
		panic(err)
	}
	nc.readConfig(content)

	// 监听配置变化
	err = nc.nacosConfigClient.ListenConfig(vo.ConfigParam{
		DataId:   nc.dataID,
		Group:    nc.group,
		OnChange: nc.onChange,
	})
	if err != nil {
		panic(err)
	}
}

func (nc *nacosConfigType[T]) onChange(namespace, group, dataID, data string) {
	log.Logger.Infof("nacos config changed: %s %s %s", namespace, group, dataID)
	nc.readConfig(data)
}

func (nc *nacosConfigType[T]) readConfig(cfgData string) {
	var err error
	viper.SetConfigType(defaultType)
	if err = viper.ReadConfig(bytes.NewBufferString(cfgData)); err != nil {
		panic(err)
	}
	if err = viper.Unmarshal(&nc.appConfig); err != nil {
		panic(err)
	}
}

// cancelListen 取消配置监听
func (nc *nacosConfigType[T]) cancelListen() {
	hook.Exit.Register(func() {
		err := nc.nacosConfigClient.CancelListenConfig(vo.ConfigParam{
			DataId: nc.dataID,
			Group:  nc.group,
		})
		if err != nil {
			log.Logger.Warnf("cancelListen: %v", err)
		}
	})
}
