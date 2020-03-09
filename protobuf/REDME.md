1. export PATH=$PATH:$HOME/go/bin
2. protoc --go_out=. *.proto
3. go run *.go