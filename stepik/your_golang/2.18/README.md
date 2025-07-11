# 2.18 Приведение в строку

> https://stepik.org/lesson/1500120/step/2

Мы можем использовать `strconv.FormatInt(int64(num), 10)`, где сначала преобразуем `num` в `int64`, 
а затем используем `FormatInt` для преобразования в строку с основанием 10. 
Либо, есть вариант использовать `strconv.Itoa(num)`. 
`Itoa` — это функция, которая принимает `int` и возвращает строку. Это более простой способ преобразования `int` в строку, если основание не нужно (если оно по умолчанию 10).

```go
fmt.Println(strconv.FormatFloat(price, 'f', 3, 64))
```

---

## Альтернативы:

- `fmt.Printf("%.3f\n", price)`
- `fmt.Println(fmt.Sprintf("%.3f", price))`
