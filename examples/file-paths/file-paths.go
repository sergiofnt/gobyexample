// Пакет `filepath` забезпечує функції розбору,
// та побудуви *шляхів до файлу* у спосіб, який є універсальним
// між операційними системами. Наприклад: `dir/file` у UNIX-подібних
// системах чи `dir\file` у Windows.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` можна використовувати для побудови шляхів.
	// Функція приймає будь-яку кількість аргументів
	// і будує з них ієрархічний шлях.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Ви завжди повинні використовувати `Join` замість
	// об'єднання `/` або `\` вручну. В додаток функція `Join`
	// нормалізує шляхи, видаливши зайві роздільники та лишні каталоги.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` та `Base` можна використовувати для розділення шляху до
	// каталога чи файла. Як варіант, `Split` буде
	// повернути обидва значення в одному виклику.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Також ми можемо перевірити, чи шлях абсолютний.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Деякі назви файлів мають розширення після крапки.
	// Ми можемо дізнатись розширення файлу за допомогою `Ext`.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Щоб дізнатись ім'я файлу без його розширення,
	// використовуйте `strings.TrimSuffix`.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` знаходить шлях між *base* і *target*.
	// Він повертає помилку, якщо це не можна зробити.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
