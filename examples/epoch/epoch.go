// Нерідко вимагається взнати скільки пройшло в
// секундах, мілісекундах або наносекундах з моменту
// [початку часу Unix](https://uk.wikipedia.org/wiki/Час_Unix).
// І ось як ми це робимо в Go.

package main

import "fmt"
import "time"

func main() {

	// Використайте `time.Now` з `Unix` або `UnixNano` для
	// отрмання часу з моменту початку Unix часу для отримання
	// секунд або наносекунд відповідно.
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// Зауважте що не існує `UnixMillis`, отож щоб отримати мілісекунди
	// вам необзідно буде ділити наносекунди.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// Є можливим конвернувати цілочисельні секунди або наносекунди
	// з моменту початку епохі Unix у відповідний `time`.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
