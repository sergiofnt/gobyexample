// В попередньому уроці ми використали явне закріплення
// за допомогою [mutexes](mutexes) для синхронізації доступу
// до спільного стану кількома горутинами. Іншою можливістю
// якою варто скористатись є особливіть виокристання
// горутин та каналів для синхронізації задля отримання
// такого ж результату. Цей базований на каналах спосіб
// співпадає з ідіоматичним Go в питанні спільної пам’яті
// і наданні кожного шматочка пам’яті у володіння лише 1ій
// горутині.

package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

// У цьому прикладі наш стан буде переданий у володіння
// єдиній горутині. Це гарантуватие що данні ніколи не будуть
// зіпсовані з одночасним доступом. Щоб читати та писати стан
// інші оперуючі гоуртини будуть надсилати повідомлення
// головній горутині і отримувати відповідні відповіді.
// Ці структури `readOp` та `writeOp` приховують ці запити
// і те як головна горутина відповідатиме.
type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {

    // Як і раніше, ми будемо прораховувати операції що ми
    // виконуємо.
    var readOps uint64
    var writeOps uint64

    // Наші канали `reads` та `writes` будуть викорстані
    // іншими горутинами для створення запитів на запис або на
    // читання відповідно.
    writes := make(chan *writeOp)
    reads := make(chan *readOp)

    // Це горутина що є власником стану (який був мапою в
    // попередньому прикладі) який наразі є приватним для
    // горутини зі станом. Ця горутина у циклі зчитує запити
    // на читання та запис з відповідних каналів і відповідає
    // на запити як тільки ті надходять. Відповідь відбувається
    // шляхом виконання операції що була запитана і надсиланням
    // відповіді до каналу `resp` для відмітки про успіх і
    // бажане значення у у випадку  `reads`).
    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    // Тут ми запустимо 100 горутин що будуть просити
    // зчитування у головної горутини через канал `reads`.
    // Кожна горутина створює `readOp`, надсилає її до каналу
    //  `reads` і отримує відповідь з каналу `resp`.
    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := &readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // Використовуючи такий же метод ми запускаємо
    // 10 горутин що писатимуть данні.
    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := &writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // Даємо нашим горутинам попрацювати близько секунди.
    time.Sleep(time.Second)

    // Нашеті, дивимось скільки було зчитувань та записів:
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}
