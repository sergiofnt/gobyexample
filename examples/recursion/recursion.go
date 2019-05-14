// Як і інші мови останніх декад у Go підтримується [_рекурсія_](https://uk.wikipedia.org/wiki/Рекурсія_(програмування)).

package main

import "fmt"

// Функція `fact` викликає сама себе допоки не досягне
// базового випадку `fact(0)`, коли її виконання припиняється.
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))
}
