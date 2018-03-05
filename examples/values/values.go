// Go має різні типи значень, включаючи: рядки (strings),
// ціли числа (int), числа з рухомою комою (float),
// логічні (bool) абощо. Ось кілька базових прикладів.

package main

import "fmt"

func main() {

    // Рядки можуть бути об’єднані поміж собою за допомогою
    // оператора `+`.
    fmt.Println("go" + "lang")

    // Цілі числа та числа з рухомою комою
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Логічний тип з логічними операторами,
    // <nobr>як ви могли б очікувати<nobr>.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
