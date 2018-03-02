# To try out our line filter, first make a file with a few
# lowercase lines.
$ echo 'hello'   > "line-filters"
$ echo 'filter' >> "line-filters"

# Then use the line filter to get uppercase lines.
$ cat "line-filters" | go run line-filters.go
HELLO
FILTER
