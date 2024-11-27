package utils

import (
	"fmt"
	"net"
)

// GetInternalIP 获取本机内网ip
func GetInternalIP() (string, error) {
	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	// 遍历所有网络接口
	for _, iface := range interfaces {
		// 过滤掉非活动的接口和回环接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// 获取接口的地址
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}

		// 遍历接口的地址
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 过滤掉回环地址和非 IPv4 地址
			if ip == nil || ip.IsLoopback() || ip.To4() == nil {
				continue
			}

			// 返回第一个找到的 IPv4 地址
			return ip.String(), nil
		}
	}

	return "", fmt.Errorf("no internal IP found")
}
