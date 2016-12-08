# Go compilation flags to make the executable small (-s -w) and to make it 
# work with alpine linux
CGO_ENABLED=0 go build -ldflags="-s -w" -a -installsuffix cgo -o bin/capserv main.go

