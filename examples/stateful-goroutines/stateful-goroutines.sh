# Запуск нашої программи покаже що базований на горутинах
# приклад управління станом проводить біля 95,000 операцій.
$ go run stateful-goroutines.go
readOps: 87142
writeOps: 8740

# У конкретно цій справі базований на горутинах спосіб
# був трошка більш задіяний ніж [mutex](mutexes)-овий.
# Це може бути корисним у випадках коли і інші канали
# задіяні або коли управління кількома такими `mutex`ами
# буде схильне до помилок. Слудує викоирстовувати той спосіб
# який вам більш до впоодби, з урахуванням розуміння того
# наскільки коректно працює ваша программа.
