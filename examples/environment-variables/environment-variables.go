// [Зминніні середовища](https://uk.wikipedia.org/wiki/Змінні_середовища)
// це універсальний спосіб [передачі конфігурацій программам Unix](https://www.12factor.net/uk/config).
// Давайте розгялнемо як встановлювати, отримаувати та показувати зміннні середовища.

package main

import "os"
import "strings"
import "fmt"

func main() {

    // Щоб встановити пару ключ/значення, скористайтесь `os.Setenv`,
    // а щоб отримати, спробуйте `os.Getenv`. Якщо ключ не знайдено в
    // сереовищі, буде повернена пустий рядок.
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    // Використовуйте `os.Environ` для переліку усіх пар ключ/значення
    // середовища. Метод поверне зріз рядків у вигляді `KEY=value`.
    // Ми можете розділіти рядок за допомогою `strings.Split`
    // щоб отримати значення. В цьому прикладі ми розрукуємо ключі.
    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        fmt.Println(pair[0])
    }
}
