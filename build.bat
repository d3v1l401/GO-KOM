@echo off

echo "Building go application..."

go build main.go RBuffer.go RHeader.go Encryption.go
ren "main.exe" "GoKOM.exe"

echo "Done"