// [_Аргументи командного рядка_](https://uk.wikipedia.org/wiki/%D0%86%D0%BD%D1%82%D0%B5%D1%80%D1%84%D0%B5%D0%B9%D1%81_%D0%BA%D0%BE%D0%BC%D0%B0%D0%BD%D0%B4%D0%BD%D0%BE%D0%B3%D0%BE_%D1%80%D1%8F%D0%B4%D0%BA%D0%B0#%D0%A4%D0%BE%D1%80%D0%BC%D0%B0%D1%82_%D0%BA%D0%BE%D0%BC%D0%B0%D0%BD%D0%B4%D0%B8) це загальноприйнятий спосіб
// запуску програм з командний інтерфейсом. Наприклад,
// `go run hello.go` використовує `run` і  `hello.go`
// як аргументи для программи `go`.

package main

import "os"
import "fmt"

func main() {

    // `os.Args` надає доступ для незмінених аргументів
    // командного рядка. Зауважте що першим значенням у
    // цьому зрізі буде шлях до самої программи, а `os.Args[1:]`
    // буде утримувати вже аргументи до программи
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    // Ви можете отримати значення індивідуальних аргументів
    // використовуючи звичайний синтаксис доступа по індексу.
    arg := os.Args[1]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
