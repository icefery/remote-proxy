function build() {
  go build -o out/remote-proxy src/main.go
}

case $1 in
build) build ;;
esac
