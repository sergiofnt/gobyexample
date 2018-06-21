# `zeroval` не змінює `i` в `main`,
# а `zeroptr` робить це бо має вказівник
# на комірку пам’яті виділену для `i`.
$ go run pointers.go
initial: 1
zeroval: 1
0xc4200160b8
zeroptr: 0
pointer: 0xc4200160b8
