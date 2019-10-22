 # Go за прикладом [![Build Status](https://travis-ci.org/butuzov/gobyexample.svg?branch=ukrainian)](https://travis-ci.org/butuzov/gobyexample)


Наповнення та інструментарій для [Go за Прикладом](https://butuzov.github.io/gobyexample/), сайту що навчає Go за допомогою анотованих прикладів.

### Загальне

Сайт "Go за прикладом" збудовано шляхом обробки коду та коментарів отриманих з першоджерельних файлів (що знаходяться в директорії `examples`) та форматуванню їх за допомогою шаблонів (з директорій `templates`) у статичні файли (що лежатимуть у директорії `public`). Інструменти що забезпечують весь процес створення сайт знаходяться у директорії `tools`, поряд з деякими залежностями (що лежать у `vendor`).


Створена директорія `public` може буде завантажена на будь-який shared/CDN/cloud хостинг.

### Побудова сторінок сайту

Щоб побудувати сайт - вам знадобляться Go та Python. Запустіть:

```bash
> go get github.com/russross/blackfriday
> tools/build
> open public/index.html
```

Щоб будувати безперервно скористайтесь (хоч, насправді, це не так зручно):

```bash
> tools/build-loop
```

### Ліцензія

Ця робота є авторським правом Mark McGranaghan та ліцензована за
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Go's Гофер є авторським правом [Renée French](http://reneefrench.blogspot.com/) та ліцензовано за
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

### Інші переклади

Переклади "Go by Example" від волонтерів доступні в наступних версіях:


* [Китайська](https://gobyexample.xgwang.me/) за авторства [xg-wang](https://github.com/xg-wang/gobyexample)
* [Чеська](http://gobyexamples.sweb.cz/) за авторства [martinkunc](https://github.com/martinkunc/gobyexample-cz)
* [Японська](http://spinute.org/go-by-example) за авторства [spinute](https://github.com/spinute)
* [Французька](http://le-go-par-l-exemple.keiruaprod.fr) за авторства [keirua](https://github.com/keirua/gobyexample)
* [Італійська](http://gobyexample.it) за авторства [Go Italian community](https://github.com/golangit/gobyexample-it)
* [Корейська](https://mingrammer.com/gobyexample/) за авторства [mingrammer](https://github.com/mingrammer)
* [Іспанська](http://goconejemplos.com) за авторства [Спільноти Go у Мексиці](https://github.com/dabit/gobyexample)
* [Українська](https://butuzov.github.io/gobyexample/) за авторства [butuzov](https://github.com/butuzov/gobyexample)

### Дякуємо

Дякуюємо [Jeremy Ashkenas](https://github.com/jashkenas)
за [Docco](http://jashkenas.github.com/docco/), що надихнули на цей проект.
