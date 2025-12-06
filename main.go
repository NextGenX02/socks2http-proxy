package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/elazarl/goproxy"
	"golang.org/x/net/proxy"
)

type serverConfig struct {
	Socks5ProxyAddr string `json:"socks_proxy_addr"`
	HttpProxyPort   string `json:"http_port"`
	HttpProxyHost   string `json:"http_host"`
}

func main() {
	configData := readConfig()
	socksDialer, dialError := proxy.SOCKS5("tcp", configData.Socks5ProxyAddr, nil, proxy.Direct)
	if dialError != nil {
		panic(dialError)
	}

	httpProxySrv := goproxy.NewProxyHttpServer()
	httpProxySrv.Verbose = true

	httpProxySrv.Tr = &http.Transport{
		Dial:                  socksDialer.Dial, //TODO:Deprecated Fix later
		ResponseHeaderTimeout: 30 * time.Second,
	}

	httpProxySrv.ConnectDial = func(network string, addr string) (net.Conn, error) {
		return socksDialer.Dial(network, addr)
	}

	fmt.Printf("Proxy Server is running on port %s | Route to %s", configData.HttpProxyPort, configData.Socks5ProxyAddr)

	listenError := http.ListenAndServe(fmt.Sprintf("%s:%s", configData.HttpProxyHost, configData.HttpProxyPort), httpProxySrv)
	if listenError != nil {
		panic(listenError)
	}
}

func readConfig() *serverConfig {
	wkdir, _ := os.Getwd()
	var configs serverConfig

	fileConfig, readError := os.ReadFile(path.Join(wkdir, "server_config.json"))
	if readError != nil {
		if os.IsNotExist(readError) {
			panic("Can't find server_config.json")
		}
		panic(readError)
	}
	unmarshalError := json.Unmarshal(fileConfig, &configs)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	return &configs
}
