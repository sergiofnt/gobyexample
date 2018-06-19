## Go за прикладом [![Build Status](https://travis-ci.org/butuzov/gobyexample.svg?branch=ukrainian)](https://travis-ci.org/butuzov/gobyexample)

Наповнення та інструментарій для [Go за Прикладом](https://gobyexample.com.ua), сайту що навчає Go за допомогою анотованих прикладів.

### УВАГА

Робота, наразі, не завершена. Ми перевіряємо на банальні орфографічні помилки і стилистику перекладів.

### Загальне

Сайт "Go за прикладом" збудовано шляхом обробки коду та коментарів отриманих з першоджерельних файлів (що знаходяться в директорії `examples`) та форматуванню їх  за допомогою шаблонів (з директорій `templates`) у статичі файли (що лежатимуть у директорії `public`). Інстурменти що забезпечують весь процес створення сайт знаходяться у директорії `tools`, поряд з деякими залежностями (що лежать у `vendor`).

Створена директорія `public` може буде завантажена на будьякий shared/CDN/cloud хостинг.

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

### Publishing

Ця варіація проекту **gobyexample** розрортується автоматично на сервера [linode](https://linode.com) безпосередньо зі сховища артефактів (див гілку "**gobyexample.com.ua**")

Оригінальний сайт розгортувався на AWS S3 наступним чином.

```bash
> gem install aws-sdk
> export AWS_ACCESS_KEY_ID=...
> export AWS_SECRET_ACCESS_KEY=...
> tools/upload-site
```

### Ліцензія

Ця робота є авторським правом Mark McGranaghan та ліцензована за
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Go's Гофер є авторським правом [Renée French](http://reneefrench.blogspot.com/) та ліцензовано за
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### Інші переклади

Переклади "Go by Example" від волонтерів доступні в наступних версіях:

* [Китайська](https://gobyexample.xgwang.me/) за авторства [xg-wang](https://github.com/xg-wang/gobyexample)
* [Французька](http://le-go-par-l-exemple.keiruaprod.fr) за авторства [keirua](https://github.com/keirua/gobyexample)
* [Італійська](http://gobyexample.it) за авторства [Go Italian community](https://github.com/golangit/gobyexample-it)
* [Корейська](https://mingrammer.com/gobyexample/) за авторства [mingrammer](https://github.com/mingrammer)
* [Іспанська](http://goconejemplos.com) за авторства [Спільноти Go у Мексиці](https://github.com/dabit/gobyexample)

### Дякуємо

Дякуюємо [Jeremy Ashkenas](https://github.com/jashkenas)
за [Docco](http://jashkenas.github.com/docco/), що надихнув на цей проект.
