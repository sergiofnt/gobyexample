// Юніт-тестування (або модульне тестування) є важливою частиною написання
// принципових Go програм. Пакет `testing` надає
// інструменти, необхідні для написання юніт-тестів,
// а команда `go test` запускає виконання тестів.

// Для демонстрації ми приводимо код з пакету
// `main`, але це міг би бути будь-який інший пакет. Код тестів
// зазвичай розміщують у тому ж самому пакеті, код якого тестують.
package main

import (
	"fmt"
	"testing"
)

// Ми тестуємо цю просту реалізацію функції знаходження 
// найменшого числа. Зазвичай, якщо код, який ми збираємось тестувати
// знаходиться у файлі з іменем, наприклад
// `intutils.go`, то файл з кодом тестів має
// називатись відповідно `intutils_test.go`.
func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Тест створюється шляхом написання функції з іменем,
// яке має починатися з `Test`.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` повідомить про помилку, але продовжить
		// виконання тесту. `t.Fail*` повідомить про помилку
		// та негайно зупинить виконання тесту.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Написання тестів може бути повторюваним, тому це ідіоматично
// використовувати *стиль керований таблицею* (*table-driven style*),
// де тестові вхідні та очікувані вихідні дані наведені в таблиці, та 
// в одному циклі проходити ці дані й виконувати алгоритми тестування.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 1},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// t.Run дозволяє запускати "підтести", по одному для кожного
		// запису таблиці. Вони відображаються окремо
		// при виконанні `go test -v`.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}