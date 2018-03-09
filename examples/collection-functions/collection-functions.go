// Часто, напотрібно щоб программи робили операції
// по збору данних, наприклад вибір усіх елементів
// що задовільняють заданий предикат або розмітка усіх
// елементів у нову колекцію за допомогою користувацьких
// функцій. У деяких мовах ідіоматично в такому разі використовувати
// [узагальнені (generic)](https://uk.wikipedia.org/wiki/%D0%A3%D0%B7%D0%B0%D0%B3%D0%B0%D0%BB%D1%8C%D0%BD%D0%B5%D0%BD%D0%B5_%D0%BF%D1%80%D0%BE%D0%B3%D1%80%D0%B0%D0%BC%D1%83%D0%B2%D0%B0%D0%BD%D0%BD%D1%8F) структури даних та алгоритми.
// Go не підтримує узагальнення, в Go загальноприйнятою
// практикою є надавати функцій що працюють з колекціями -
// якщо і коли вони будуть потрібні в ваших программах та
// типах данних.

// Ось деякі приклади, що оперують над зрізом рядків. Ви можете
// використовувати їх для створення власного функціоналу.
// Зауважте що в деяких випадких можливо буде набагато ефективніше
// якщо б ви вставили код напряму без створення окремої функції.

package main

import "strings"
import "fmt"

// Index повертає перший ідекс елемента який дорівнює `t`,
// або -1 якщо співпадень не знайдено.
func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

// Include поверне `true` якщо рядок `t` знайдено у зрізі.
func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

// Any поверне `true` якщо хочаб один з рядків задовольняє [предикат](https://uk.wikipedia.org/wiki/%D0%9F%D1%80%D0%B5%D0%B4%D0%B8%D0%BA%D0%B0%D1%82) `f`.
func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

// All поверне `true` якщо усі рядки задовольняють предикат `f`.
func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

// Filter поверне новий зріз, що містить усі рядки
// які задовільняють предикат `f`.
func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

// Map поверне новий зріз що містить результати роботи предиката
// `f` над кожним рядком оригінального зріза.
func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {

    // Приклад застосування функцій роботи з колеціями що
    // ми щойно переглянули.
    var strs = []string{"peach", "apple", "pear", "plum"}

    fmt.Println(Index(strs, "pear"))

    fmt.Println(Include(strs, "grape"))

    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))

    // Више приведені приклади усі використовують анонімні
    // фукнції але ви можете використовувати і звичайні
    // іменовні якщо вони задовольняють потрібний нам тип.
}
