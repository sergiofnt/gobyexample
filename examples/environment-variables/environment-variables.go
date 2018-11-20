// [Змінніні середовища](https://uk.wikipedia.org/wiki/Змінні_середовища)
// це універсальний спосіб [передачі конфігурацій программам, що працююь у середовищі Unix](https://www.12factor.net/uk/config).
// Давайте розглянемо як створювати, отримувати та показувати ці змінні.

package main

import "os"
import "strings"
import "fmt"

func main() {

    // Ми можемо скористатись `os.Setenv` для встановлення пар
    // ключ/значення і `os.Getenv` для отримання значення по ключу
    // (в разі якщо ключ не буде знайдено - повернеться пустий рядок).
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    // `os.Environ` поверне зріз усіх змінних середовища, у вигляді
    // `KEY=value` (який ви надалі можете розділити на частини
    // скориставшись `strings.Split`). В даному прикладі, ми лише
    // виведемо назви ключів змінних.
    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        fmt.Println(pair[0])
    }
}
