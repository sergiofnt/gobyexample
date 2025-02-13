// [_Прапори командного рядку_](https://uk.wikipedia.org/wiki/Інтерфейс_командного_рядка#Формат_команди) -
// це загально прийнятий спосіб вказувати налаштування програмам
// призначеним для використання з командного рядка. Наприклад,
// у `wc -l` прапорецем є `-l`

package main

// Пакет Go `flag` підтримує основні операції
// розбору прапорців командного рядку.
// Ми скористаємось ним, для створення власної консольної програми.
import (
	"flag"
	"fmt"
)

func main() {

	// Декларування прапорців доступно для насутпних типів даних -
	// рядки, цілі чисела та логічний тип даних. Ми декларуємо
	// прапорець `word` з значенням за замовчуванням `"foo"`
	// та короткою довідкою. Функція `flag.String`
	// поверне вказівник на рядок, а не рядкову змінну, ми ще
	// побачимо як користуватись таким вказівниками в майбутньому.
	wordPtr := flag.String("word", "foo", "a string")

	// Тут ми декларуємо прапорці `numb` та `fork`
	// скориставшись вже знайомим підходом.
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// Також, можливо декларувати прапорець - таким чином
	// щоб вже існуюча змінна прийняла на себе його значення.
	// Зауважте, що потрібно передати вказівник на цю змінну
	// аргументом до функції яка декларуватиме прапорець.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Як тільки усі прапорці задекларовані, ми викличемо `flag.Parse()`
	// для виконання аналізу прапорців командного рядку.
	flag.Parse()

	// Наразі, ми просто виведемо на екран усі розпізнані налаштування
	// та позиційні аргументи. Зауважте, що нам потрібно розіменувати
	// вказівники, наприклад `*wordPtr`, щоб отримати справжнє його
	// значення.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
