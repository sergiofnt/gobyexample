// Використовуючи канали як параметри функції, дозволяється
// вказувати призначення каналу - чи він буде використаний
// для відправки чи тільки для отримання значень.
// Використання цієї особливості підвищує безпеку программи.

package main

import "fmt"

// Функція `ping` прийме канал лише для надсилання значень.
// Ви отримаєте помилку компіляції, якщо спробуєете отримати
// значення з цього каналу.
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// Функція `pong` приймає перший канал для доставки
// (`pings`), a другий для відправки (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "передане повідомлення")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
