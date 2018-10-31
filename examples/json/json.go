// Go підтримує кодування та розкодування JSON на рівні
// як базових типів даних так і ваших власних структур даних.

package main

import "encoding/json"
import "fmt"
import "os"

// Ми використаємо ці структури щоб продемонструвати
// кодування та розкодування не типових структур данних.
type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// Спочатку розглянемо кодування базових типів в JSON рядки.
	// Це в нас приклади атомарних значень.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// А це приклад з зрізами та мапам, що кодуються в JSON
	// масиви та об’єкти саме так, як ми від них цього чекаємо.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Пакет JSON може автоматично кодувати ваші типи
	// данних. Це дозволяє автоматично включати експортовані
	// поля до вашого виводу. Ці поля по-замовчуванню
	// будуть носити тіж назви що і поля структури яка
	// кодувалась.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Ви можете скористатись тегами (матеданими) під час
	// декларації полів структури для налаштування на власний розсуд того
	// як кодуватимуться імен ключів для JSON. Задля прикладу, переірте, подане вище, визначення структури `response2`.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Тепер, розглянемо як декодувати JSON дані в
	// значення Go, на прикладі узагальненої структури.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Ми маємо надати змінну, де б пакет JSON був здатний
	// зберігати розкодовані данні. Змінна dat` з типом
	// `map[string]interface{}` використовуватиме як ми бачимо
	// мапу з ключами типу рядок та довільним значенням.
	var dat map[string]interface{}

	// Власне розкодування та перевірка на супутні помилки
	// відбувається тут.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Щоб використати данні з розкодованої мапи, нам потрібно буде
	// скастувати (від англійської _to cast_ - надавати форму, що
	// використовується в сенсі "перетворювати один тип
	// данних у інший шляхом, що підтримується обраною мовою") їх
	// до відповідного типу. Для прикладу - ми кастуємо значення
	// значення `num` до очікуваного типу `float64`.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Доступ до вкладених данних потребує від нас серій
	// кастувань.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Дозволяється розкодувати JSON у певну структуру
	// даних і слід надавати перевагу цьому методу оскільки
	// він гарантує нашій программі  додаткову типобезпеку
	// і усуває потребу в перевірках типу при використанні розкодованих
	// даних.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// В прикладах вище, ми використовували байти та рядки
	// як посередники між даними і JSON представленням при
	// виведенні до стандартного потоку виведення. Ми також
	// можемо напрвляти поток кодувань на пряму в `os.Writer`
	// так само як `os.Stdout` або навіть у HTTP Response.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
