// Go підтримує _методи_ визначені на типі структура.

package main

import "fmt"

type rect struct {
    width, height int
}

// Ця структура `area` має метод з _отримувачем_ типу `*rect`*.
func (r *rect) area() int {
    return r.width * r.height
}

// Методи можуть визначатись дял отримувачів типу
// вказівник або значення. Ось приклад метода з
// отримувачем типу значення.
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    // Тут ми викличемо 2 методи визначені для нашої структури.
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    // Go автоматично визначає чи потрібне конвернування між
    // значеннями та вказівниками для виклику методів. Можливо
    // ви захочете використати отримувач вказівника щоб уникнути
    // копіювання данних при виклиці методу або для зміни данних
    // в структурі.
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
