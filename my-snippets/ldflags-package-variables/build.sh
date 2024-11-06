go build -ldflags "-X 'main.outputGlobal=apples' -X 'main.outputLocal=oranges'" -o lpv main.go
./lpv
rm lpv