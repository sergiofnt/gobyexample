// Твердежння `switch` виражають умовності на кілька гілок.
package main

import "fmt"
import "time"

func main() {

    // Ось базовий `switch`.
    i := 2
    fmt.Print("Напиши ", i, " як ")
    switch i {
    case 1:
        fmt.Println("один")
    case 2:
        fmt.Println("два")
    case 3:
        fmt.Println("три")
    }

    // Ви можете використовувати коми - задля розподілу
    // кількох умов в одній гілці твердженням `case`.
    // Ми також використовуємо необов’язкове твердження `default`,
    // для розгалуження умови, що відображатиме стандартно-задану дію.
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("Це вихідні")
    default:
        fmt.Println("Це робочий день")
    }

    // `switch` без виразу - являє собою альтернативний спосіб
    // виразити `if/else` логіку. Приклад також демонструє,
    // що гілки `case` можуть містити не тільки константи.
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Над полудень")
    default:
        fmt.Println("По полудні")
    }

    // Тип `switch` що порівнює типи замість значень. Ви можете
    // використати це для визначення типу значення інтерфейса.
    // В цьому прикладі `switch` видасть повідомлення базуючись на
    // значенні типу змінної `t` якщо йому про це буде відомо.
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("Я логічна змінна")
        case int:
            fmt.Println("Я ціле число")
        default:
            fmt.Printf("Незнаю нічого про тип %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("Гей!")
}
