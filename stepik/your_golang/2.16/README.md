# README.md

> https://stepik.org/lesson/1500118/step/4?unit=1520227
> `2025-06-21 22:48`

---

```go
// Вы уже находитесь в функции, создавать функцию, прописывать package и import не нужно.
// Пакеты "unicode/utf8" и "fmt" импортированы.
// Переменная str уже существует и содержит строку, с которой нужно работать.

fmt.Println(str, len(str), utf8.RuneCountInString(str))

```

```terminal
Input:
Ты ведь любишь Golang?
Output:
Ты ведь любишь Golang? 34 22
```

or...

`fmt.Printf("%s %d %d\n", str, len(str), utf8.RuneCountInString(str))`
