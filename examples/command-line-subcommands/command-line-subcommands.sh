$ go build command-line-subcommands.go 

<<<<<<< HEAD
# Для початку викликати команду foo.
=======
# First invoke the foo subcommand.
>>>>>>> master
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

<<<<<<< HEAD
# Тепер викликати команду bar.
=======
# Now try bar.
>>>>>>> master
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

<<<<<<< HEAD
# Переконатись, що bar не приймає параметри визначенні в foo.
=======
# But bar won't accept foo's flags.
>>>>>>> master
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

<<<<<<< HEAD
# Далі будуть розглянуті змінні середовища,
# ще один поширений спосіб для заданя необхідних параметрів.
=======
# Next we'll look at environment variables, another common
# way to parameterize programs.
>>>>>>> master
