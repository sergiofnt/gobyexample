// _Інтерфейси_ - це іменовані коллекції описів методів.

package main

import "fmt"
import "math"

// Базовий приклад інтерфейсу для геометричних фігур.
type geometry interface {
    area() float64
    perim() float64
}

// Задля прикладу ми реалізуємо цей інтефейс для типів
// `rect` (квадрат) та `circle` (круг).
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

// Щоб релізувати інтерфейс в Go, нам лише потрібно
// створити усі методи в цьому інтерфейсі. Як у
// прикладі де ми імплементуємо `geometry` для
// `rect` (квадрата).
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// І реалізація для `circle` (круга).
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// У разі, якщо змінна має тип інтерфейс, нам дозволяється
// викликати оглошені в цьому інтерфейсі методи.
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    // Обидві структури `circle` (круг) і `rect` (квадрат)
    // реалізуються інтерфейс `geometry` отож ми можемо викостати
    // їх зразки як аргумент до фукнції `measure`.
    measure(r)
    measure(c)
}
