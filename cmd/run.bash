# service端微服务
nohup go run ./cmd/service/user/main.go > /dev/null 2>&1 &
nohup go run ./cmd/service/feed/main.go > /dev/null 2>&1 &
# web端微服务
nohup go run ./cmd/web/user/main.go > /dev/null 2>&1 &
nohup go run ./cmd/web/feed/main.go > /dev/null 2>&1 &