package util

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// MD5String 将指定字符串加密为 md5，
// 返回加密后的 md5 字符串。
//
// str 字符串参数，传入要加密的字符串
func MD5String(str string) string {
	// 将字符串转换为字节数组
	bs := []byte(str)
	// 加密字符串
	h := md5.Sum(bs)

	return fmt.Sprintf("%x", h)
}

// Base64 将指定文件编码为 Base64 字符串，
// 返回编码信息及错误信息。
//
// file 字符串参数，传入文件路径
func Base64(file string) (string, error) {
	// 检查错误
	f, err := os.Open(file)
	// 如果出错
	if err != nil {
		return Base64ForURI(file)
	}
	// 关闭
	defer f.Close()

	// 初始化byte
	buff := make([]byte, 500000)
	// 读取文件
	n, err := f.Read(buff)
	// 检查错误
	if err != nil {
		return "", err
	}

	// Base64编码
	source := base64.StdEncoding.EncodeToString(buff[:n])

	return source, nil
}

// Base64ForURI 从远程文件获取 Base64
// 返回编码信息及错误信息。
//
// uri 字符串参数，传入远程链接
func Base64ForURI(uri string) (string, error) {
	// 解析网址
	u, err := url.Parse(uri)
	if err != nil || u == nil {
		return "", err
	}

	// 打开远程链接
	res, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	// 关闭
	defer res.Body.Close()

	// 初始化byte
	buff := make([]byte, 500000)
	// 读取文件
	n, err := res.Body.Read(buff)
	// 检查错误
	if err != nil {
		return "", err
	}

	// Base64编码
	source := base64.StdEncoding.EncodeToString(buff[:n])

	return source, nil
}
