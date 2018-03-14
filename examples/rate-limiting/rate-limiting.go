// <em>[Обмеження Частоти Запитів](http://en.wikipedia.org/wiki/Rate_limiting)</em>
// це важливий механізм контролю за ресурсом, використаня
// та підтримки необхідної якості роботи. Go елегентно
// підтримує обмеження частоти за допомогою
// [горутин](goroutines), [каналів](channels)
// та [маятників](tickers).

package main

import "time"
import "fmt"

func main() {

    // Спершу подивимось на звичайний приклад обмеження
    // частоти. Припустимо ми хочемо обмежити обслуговування
    // вхідних запитів, ми будемо обслууговувати ці запити
    // через одноіменний канал (`requests`).
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    // Цей канало `limiter` (обмежувач) буде приймати
    // значення кожні 200 мілісекунд, це регулятор
    // нашої системи обмеження частоти запитів.
    limiter := time.Tick(200 * time.Millisecond)

    // Блокуючи на отримувачі з каналу `limiter` перед
    // обробкою кожного запиту,ми обмежуємо себе до 1
    // запиту кожні  200 мілісекунд.
    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }
    // Можливо ми схочемо дозволити невеликі напливи запитів,
    // вцілому притримуючись нашої стандартної схеми
    // обмеженя. Ми можемо досигнути цього створюючи буфер
    // в нашому каналі обмеження. Цей канал `burstyLimiter`
    // дозволятиме напиливи аж трьох подій.
    burstyLimiter := make(chan time.Time, 3)

    // Заповнюємо канал, для представлення допустимого
    // напливу.
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    // Кожні 200 мілісекунд ми спробуємо додати значення в
    // `burstyLimiter`, аж до його ліміту в 3 події.
    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    // Тупер ми симулюємо 5 вхідних запитів. Перші три,
    // отримають перевагу від реалізації буфера каналу
    // `burstyLimiter`.
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
