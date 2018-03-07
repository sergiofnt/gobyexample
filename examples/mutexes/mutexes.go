// У попередньому пиркладі ми бачили як керувати простим
// лічильником використовуючи [атомарні лічильники](atomic-counters).
// В разі більш складних випадків ми можемо використовувати  [_м’ютекси_](https://uk.wikipedia.org/wiki/%D0%9C%27%D1%8E%D1%82%D0%B5%D0%BA%D1%81)
// для безпечного доступу до данних в горутинах.

package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

func main() {

    // В нашому прикладі `state` буде мапою.
    var state = make(map[int]int)

    // А цей `mutex` сихнронізуватиме доступ до `state`.
    var mutex = &sync.Mutex{}

    // Ми будемо слідкувати скільки операцій запису та
    // читання ми робитимемо.
    var readOps uint64
    var writeOps uint64

    // Ось ми запускаємо 100 горутин для запуску повторних
    // зчитувань данних стану нашої мапи раз в мілісекунду
    // в кожній горутині.
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {

                // Для кожного читання ми: генеруємо ключ для
                // доступа, замикаємо (`Lock()`) м’ютекс
                // щоб впевнитись в унікальному доступі до `state`,
                // чичтаємо значення обраного ключа,
                // відімкнути (`Unlock()`) м’ютекс, і зібльшуємо
                // на одиницю лічильник `readOps`.
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddUint64(&readOps, 1)

                // Чекаємо мілісекунду до наступного читання.
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // Ми також стартуємо 10 горутин, для записів,
    // використовуючи цей же шаблон
    // (що ми використали для читанння).
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // Дамо 10 горутинам працювати з `state` та
    // `mutex` рівно 1 секунду.
    time.Sleep(time.Second)

    // Прочитаємо звіти щодо наших автоманих операцій читання
    // та запису.
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

    // І покажемо як же `state` став виглядати після того всього.
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}
