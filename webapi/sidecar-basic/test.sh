go run cmd/sidecar/main.go &

go run cmd/main/main.go &

sleep 1

curl -v localhost:8081