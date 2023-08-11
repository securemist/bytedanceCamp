/********************************************************************************
* @author: Yakult
* @date: 2023/8/3 15:39
* @description:
********************************************************************************/

package util

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"net"
)

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

// EncodePassword 将明文的密码进行加密
func EncodePassword(plaintext string) (ciphertext string) {
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(plaintext, options)
	ciphertext = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	return ciphertext
}
