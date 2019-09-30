// Go має декілька корисних функцій для роботи з
// *директоріями* в файловій системі.

package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // Створити нову директорію, під назвою *subdir* в поточній.
    // Другим параметром являється налаштування [доступу](https://en.wikipedia.org/wiki/File_system_permissions#Numeric_notation) до директорії.
    err := os.Mkdir("subdir", 0755)
    check(err)

    // При створенні тимчасової директорії, непоганою ідеєю являється
    // _відкласти_ її видалення за допомогою [`defer`](./defer).
    // `os.RemoveAll` видалить директорію *subdir* повністю (разом з вкладеними директоріями) -
    // аналогічно з `rm -rf subdir` в UNIX-подібних системах.
    defer os.RemoveAll("subdir")

    // Допоміжна функція для створення тимчасового файлу.
    createEmptyFile := func(name string) {
        d := []byte("")
        check(ioutil.WriteFile(name, d, 0644))
    }

    createEmptyFile("subdir/file1")

    // Також можливо створити повну ієрархію директорій за допомогою `MkdirAll`.
    // Дана команда являється аналогічною до команди `mkdir -p subdir/parent/child`.
    err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)

    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

    // `ReadDir` зчитує зміст директорії *parent*,
    // і повертає [зріз](./slice) об'єктів `os.FileInfo`.
    c, err := ioutil.ReadDir("subdir/parent")
    check(err)

    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // `Chdir` дозволяє змінити поточну робочу директорію,
    // подібно до `cd`.
    err = os.Chdir("subdir/parent/child")
    check(err)

    // Дана команда поверне [зріз](./slice) об'єктів `os.FileInfo` для поточної директорії,
    // якою на даний момент являється `subdir/parent/child`
    c, err = ioutil.ReadDir(".")
    check(err)

    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // `cd` Повернутись на стартову позицію.
    err = os.Chdir("../../..")
    check(err)

    // Також ми можемо обійти директорію *рекурсивно*,
    // включаючи всі вкладені директорії.
    // `Walk` другим параметром приймає функцію зворотного виклику,
    // яка буде викликана для кожного *відвіданого* файлу та директорії.
    fmt.Println("Visiting subdir")
    err = filepath.Walk("subdir", visit)
}

// `visit` - функція яка буде викликана для кожного файлу чи директорії
// знайденого в процесі обходу *subdir*.
func visit(p string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", p, info.IsDir())
    return nil
}
