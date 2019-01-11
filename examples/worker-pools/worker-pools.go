// Розглянемо реалізацію патерна
// _worker pools_ використовуючи горутини та канали.
// Вільний переклад "wooker pools" українською "фонд робочої сили",
// але оригінальний термін вживаєтья частіше

package main

import "fmt"
import "time"

// Ось приклад, функції-робітника, якою ми скористаємось,
// щоб запустити кілька різних екземплярів `робітників`. Вони будуть
// отримувати роботу на каналі `jobs` і надсилати сповіщення
// про завершення роботи до каналу `results`. Роботою нашого
// "працівника" буде "сон довжиною в секунду" (така собі симуляція
// реальної роботи).
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("працівник", id,
            "почав проботу над завданням", j)
        time.Sleep(time.Second)
        fmt.Println("працівник", id,
            "закінчив роботу над завданням", j)
        results <- j * 2
    }
}

func main() {

    // Щоб скористатись "worker pool" нам необхідно
    // створити два канали: в один будемо передавати завдання
    // робітнику, а з іншого отримавути результати роботи.
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Спершу - запускаємо 3 функції-працівника, які блокуватимутся
    // під час виконання (тому що жодної роботи для них поки що існує).
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Далі - Надішлемо 5 завдань до каналу роботи, після чого
    // закриває канал щоб вказати що це вся робота що, наразі, доступна.
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // І, нарешті, отримуємо результати роботи наших працівників.
    for a := 1; a <= 5; a++ {
        <-results
    }
}
