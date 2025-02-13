// _Функції_ - це базова складова мови Go. Ми будемо знайомитись з
// ними за допомогою кількох різних прикладів.

package main

import "fmt"

// Ось функція, що приймає два цілих числа (`int`)
// та повертає їх суму (також, як ціле число `int`).
// Go потребує чітко вказувати тип що повертається.
func plus(a int, b int) int {
	return a + b
}

// За виключенням, коли ми вказуємо можемо додатково вказати зміну
// для повернення, (зверніть увагу на приклад з змінною summ).
func plusNames(a int, b int) (summa int) {
	summa = a + b
	return
}

// Коли у вас кілька послідовних параметрів одного типу,
// дозволяється пропускати тип - вказуючи його лише
// для останнього аргумента.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
