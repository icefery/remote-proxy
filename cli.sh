function build() {
  export GOPROXY=https://goproxy.cn,direct
  go mod tidy
  go build -o bin/remote-proxy src/main.go
}

case $1 in
build) build ;;
esac
