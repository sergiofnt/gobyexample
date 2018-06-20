// Оператор розгалуження `switch` використовують
// в основному для представлення у вигляді гілок
// умов.

package main

import "fmt"
import "time"

func main() {

    // Ось базовий приклад `switch`.
    i := 2
    fmt.Print("Напишемо ", i, " як ")
    switch i {
    case 1:
        fmt.Println("один")
    case 2:
        fmt.Println("два")
    case 3:
        fmt.Println("три")
    }

    // Дозволяється використовувати коми - так ми розділяємо
    // кілька умов в одній гілці `case`. Також існує `default` -
    // стандартна умова для усіх випадків (яка виконається -
    // за умови що жодна умова ще не виконалась).
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("Це вихідні")
    default:
        fmt.Println("Це робочий день")
    }

    // `switch` без виразу - являє собою альтернативний спосіб
    // реалізувати `if/else` логіку. Приклад також демонструє,
    // що гілки `case` можуть містити не тільки константи.
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Над полудень")
    default:
        fmt.Println("По полудні")
    }

    // Так само - як і в попередньому прикладі з `if`, у
    // операторі `switch` можливо задекларувати змінну,
    // і перевірити умову в одній з гілок вибору.
    // Наразі, ми перевіряємо це для визначення типу
    // вхідної змінної (тип [`itnerface{}`](interfaces)
    // дозволяє працювати з будь-яким типом даних).
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
