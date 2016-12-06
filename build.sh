CGO_ENABLED=0 go build -ldflags="-s -w" -a -installsuffix cgo -o bin/capserv main.go

