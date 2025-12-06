# socks2http-proxy
Connect to SOCKS5 proxy via http proxy, usefull for program that doesn't have socks proxy support.
## Prerequisites
* GO Version 1.25.5 or newer
* Socks proxy server used as the upstream server

## Installation
1. Clone this repo.
   ```bash
   $ git clone https://github.com/NextGenX02/socks2http-proxy.git
   $ cd socks2http
    ```
2. Open **server_config.json** and fill the following config.
    ```json
    {
    "socks_proxy_addr": "Your SOCKS5 server address",
    "http_port": "HTTP Proxy Port",
    "http_host": "HTTP Proxy Host"
    }
    ```
3. Download Required Library.
    ```bash
    $ go mod download
    ```
4. Build the code into executable format.
    ```bash
    $ go build
    ```
   Or you can just run it without build it.
    ```bash
    $ go run main.go
    ```

5. You should see this message popup in the console which tell the server is running properly.
    ```bash
    Proxy Server is running on port <Your port> | Route to <Your SOCKS5 proxy address>
   ```
## Testing
* For testing, you can use simple curl command by typing
  ```bash
    $ culr -x http://<your http proxy host>:<your http proxy port> "https://ipinfo.io/json"
    ```
  And you should get the result like this
    ```json
    {
  "ip": "Your SOCKS5 proxy server",
  "city": "[Remove]",
  "region": "[Remove]",
  "country": "[Remove]",
  "loc": "1.xxxx,1xx.xxxx",
  "org": "[Remove]",
  "postal": "[Remove]",
  "timezone": "[Remove]"
    }
  ```
  
## Library used in this project
* [GoProxy](https://github.com/elazarl/goproxy) by elazarl.
* [proxy](https://pkg.go.dev/golang.org/x/net/proxy) by golang.org