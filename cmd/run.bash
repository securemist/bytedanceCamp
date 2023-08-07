# service端微服务
nohup go run ./cmd/service/user/main.go &
nohup go run ./cmd/service/feed/main.go &
# web端微服务
nohup go run ./cmd/web/user/main.go &
nohup go run ./cmd/web/feed/main.go &