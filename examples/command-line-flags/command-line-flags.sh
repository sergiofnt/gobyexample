# Щоб проексперементувати з прапорцями командного рядку,
# краще всього спершу скомпілювати нашу программу і вже
# потім запускати двійковий файл напряму.
$ go build command-line-flags.go

# Спробуйте спершу скомпільовану программу давши
# значення усім прапорцям.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Зауважте, що якщо ви пропустите прапорці, вони
# автоматично приймуть свої стандартні налаштування.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Заключні позиційні аргументи можуть бути вказані після
# будь яких прапорців.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Note that the `flag` package requires all flags to
# appear before positional arguments (otherwise the flags
# will be interpreted as positional arguments).
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Скористайтес прапорцями `-h` або `--help` щоб отримати автоматично
# згенеровану довідку по роботі з команднним рядком программи.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Якщо ви вкажете прапорець, якиц не було вказано в нашій программі,
# у відповідь, ви отримаєте повідомлення помилки та
# (повторно) тест довідки.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

# Далі, ми розгялдемо змінні середовища, інший
# загальноприйнятий метод параметризації програм.
