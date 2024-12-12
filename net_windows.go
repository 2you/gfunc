package gfunc

import "golang.org/x/sys/windows/registry"

func GetHttpProxyAddr() (string, error) {
	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`,
		registry.QUERY_VALUE)
	if err != nil {
		return ``, err
	}
	proxyAddr, _, err := key.GetStringValue("ProxyServer")
	return proxyAddr, err
}
