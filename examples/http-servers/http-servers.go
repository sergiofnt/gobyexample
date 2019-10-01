// Написання *HTTP*-серверу не являється складною задачою,
// якщо використовувати пакет `net/http`.
package main

import (
    "fmt"
    "net/http"
)

// Фундаментальним концептом в серверах `net/http` являються обробники(*handlers*)
// Обробник - об'єкт, що імплементує `http.Handler` інтерфейс.
// Зазвичай для написання обробника використовується адаптер `http.HandlerFunc`
// для функції з відповідною сигнатурою.
func hello(w http.ResponseWriter, req *http.Request) {

    // Функції що використовуються як обробники приймають
    // `http.ResponseWriter` і `http.Request` як аргументи.
    // `http.ResponseWriter` використовується для запису *HTTP-відповіді*
    // В даному прикладі відповіддю являється стрічка "hello\n".
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    // В даному обробнику проводиться вичитування всіх HTTP-заголовків,
    // і передавання їх в якості відповіді запитуючій стороні.
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    // Дані стрічки реєструють маршрути, які необхідно обслуговувати серверу,
    // за допомогою функції `http.HandleFunc`. Ця функція
    // встановлює маршрутизатор по замовчуванню і обробник який буде виконано,
    // якщо буде виявлено запит по вказаних маршрутах.
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    // Запускає *HTTP*-сервер. В якості порту на якому очікуються з'єднання
    // встановлено ":8090". Другий параметр вказує на те, що потрібно використовувати
    // маршрутизатор по замовчуванню, який було заздалегідь встановлено.
    http.ListenAndServe(":8090", nil)
}
