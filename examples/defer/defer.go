// _Відкладення_ використовують щоб впевнитись,
// що виклик функції буде виконано пізніше в
// виконанні программи, нерідко для очистки.
// `defer` також  використовують там саме де
// і `ensure` та `finally` використались би в
// інших мовах.

package main

import "fmt"
import "os"

// Допустимо, ми бажаємо створити файл, записати щось до
// нього, і коли ми закінчим з цим - закрити його.
// Ось як би ми це зробили за допомогою `defer`
// (ключове слово для відкладення).
func main() {

    // Відразу після отримання обєкту `File` (що повертається
    // `createFile`), ми відкладуємо його закриття за допомогою
    // `closeFile`. ВІдкладений виклик буде виконано в самому
    // кінці закриваючої функції (`main`), але після того,
    // як `writeFile` завершить своє виконання.
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("створєюмо")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("записуємо")
    fmt.Fprintln(f, "данні")

}

func closeFile(f *os.File) {
    fmt.Println("закриваємо")
    f.Close()
}
