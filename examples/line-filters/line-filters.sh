# Щоб спробувати наш рядковий фільтр, створимо файл з кількома
# рядками у нижньому регістрі.
$ echo 'hello'   > "line-filters"
$ echo 'filter' >> "line-filters"

# Після чого скористаємось рядковим фільтром, щоб отримати
# ціж рядки у верхньому регістрі.
$ cat "line-filters" | go run line-filters.go
HELLO
FILTER
