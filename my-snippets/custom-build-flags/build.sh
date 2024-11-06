go build -o apples -tags "apples"
./apples
rm apples

go build -o oranges -tags "oranges"
./oranges
rm oranges

go build -o onions
./onions
rm onions