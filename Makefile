compile-windows:
	GOOS=windows GOARCH=386 \
  CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc \
  go build cmd/GOt-trivia/main.go

compile-linux:
	GOOS=linux GOARCH=arm \
  go build cmd/GOt-trivia/main.go

compile-darwin:
	GOOS=darwin GOARCH=amd6 \
	go build cmd/GOt-trivia/main.go