/*
1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200
*/

package main

import (
	"flag"
	"io"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		glog.Fatal(err)
	}
}

func healthz(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	client_ip := getClientIP(req)
	glog.V(2).Info("clientIP ,X-Real-Ip : %s ", client_ip)
	io.WriteString(res, "200\n")
}

// 设置header 在io之前做，否则无效
func rootHandler(w http.ResponseWriter, r *http.Request) {
	glog.Info("root handler")
	env_str := os.Getenv("VERSION")
	glog.V(2).Info(env_str)
	w.Header().Add("sys_version", env_str)
	for k, v := range r.Header {
		w.Header().Add(k, v[0])
	}
	clientIP := getClientIP(r)
	glog.V(2).Info("client IP %s", clientIP)
	w.WriteHeader(200)
	io.WriteString(w, "server root")
}

func getClientIP(r *http.Request) string {
	// 先看header
	ip := strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		glog.V(2).Info("clientIP ,X-Real-Ip : %s ", ip)
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		glog.V(2).Info("clientIP : %s ", ip)
		return ip
	}
	return ""
}
