export GOOS=linux
# for mac os use: export GOOS=darwin
export GOARCH=amd64
# for arm use: export GOARCH=arm
export CGO_ENABLED=0
go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH ../../cmd/server/main.go