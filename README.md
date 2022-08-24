## Build

```bash
export GOPROXY=https://goproxy.cn,direct
go mod tidy
go build -o bin/remote-proxy src/main.go
```

## Use

```bash
bin/remote-proxy
```

Open `http://<REMOTE_IP>:8080/download?url=<DOWNLOAD_URL>`
