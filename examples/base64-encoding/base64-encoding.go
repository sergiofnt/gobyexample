// Go підтримує кодування/декодування
// [base64](https://uk.wikipedia.org/wiki/Base64).

package main

// для початку давйте включемо для імпорту пакет `encoding/base64`,
// до якого будемо звертатись, як до `b64` замість стандартного
// звернення `base64`.
import b64 "encoding/base64"
import "fmt"

func main() {

    // Це рядок - що ми будемо кодувати та декодувати.
    data := "abc123!?$*&()'-=@~"

    // Go підтримує як стандартний так і URL-сумісний base64.
    // Спочатку наведемо приклад використання стандартного base64.
    // Механізм кодування потребує зріз байтів, отож ми
    // сконвернуємо наш рядок до потрібного типу.
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    // Розкодування може повертати помилку, яку ви можете перевірити.
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))

    // А це ми кодуємо, за допомогою URL-сумісного base64.
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)

    // Розкодовуємо, аналогічно тому, як ми це зробили для
    // стандартного методу кодування base64, кількома рядками вище.
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
