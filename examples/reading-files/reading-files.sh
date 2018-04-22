$ echo "hello" > /tmp/dat
$ echo "go" >>   /tmp/dat
$ go run reading-files.go
hello
go
5 байт: hello
2 байт @ 6: go
2 байт @ 6: go
5 байт: hello

# Далі, ми побачимо як записувати файли.
