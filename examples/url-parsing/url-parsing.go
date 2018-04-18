// URLи надаютть [однаковий шлях для локації ресурсів](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// І ось як ми розбираємо URL на складові.

package main

import "fmt"
import "net"
import "net/url"

func main() {

    // Ми озбиратимемо цей приклад URL, він включає схему,
    // інфо аутентифікації, хост, порт, шлях,
    // параметри запиту і фрагмент запиту (якір).
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    // Розібр URL дозволяє перевірити наявність помилки.
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    // Отримати доступ до схеми (і інших складових) надзвичайно просто.
    fmt.Println(u.Scheme)

    // `User` збрерігає усю інфомарцію щодо аутентифікації;
    // Зверніться до `Username` та `Password` для отримання
    // індивідуальних значень.
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    // `Host` складається як з імени хоста так і порту,
    // якщо він присутній. СКористайтесь `SplitHostPort`
    // щоб розбити ці значення на складові.
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    // А таким чином ми звернемось до `шляху` та фрагмента після `#`.
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    // Щоб отримати параметри запиту в рядку у форматі
    // `k=v` скористайтесь `RawQuery`. Також ви маєте змогу
    // розпарсити парметри запиту у мапу. Розібрані парарметри
    // запиту в мапі зберігаються як ключ-рядок до значення
    // що зберігається в зрізі рядків, отож звертайтесь за індексом
    // `[0]` якщо вам потрібно буде тільки перше значення.
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}
