/********************************************************************************
* @author: Yakult
* @date: 2023/8/3 15:39
* @description:
********************************************************************************/

package register

import "net"

// GetFreePort 获取可用的端口号
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
