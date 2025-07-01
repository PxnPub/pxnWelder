
clear
go mod tidy  || exit 1
go run .     || exit 1
