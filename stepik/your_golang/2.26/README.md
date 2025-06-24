# 2.26 Тип any или interface{}

> https://stepik.org/lesson/1500128/step/2

---

Правильно, оба варианта используют утверждение типа (type assertion), 
чтобы проверить, является ли val строкой. Если да, то ok будет true, иначе false.

- [ ] `if err := val.(error); err != nil { ... }`
- [ ] `if str, ok := val.string(); ok { ... }`
+ [x] `if _, ok := val.(string); ok { ... }`
+ [x] `if str, ok := val.(string); ok { ... }`
- [ ] `if str := val is string { ... }`

---

## overflow

- [search interface](https://pkg.go.dev/search?q=interface&m=)

> https://www.reddit.com/r/golang/comments/15syxg0/type_any/

буквально, type any = interface{}. Это не "новый тип", 
это буквально просто псевдоним типа для типа interface{}, 
и на 100% эквивалентен ему во всех отношениях. 
Это просто более удобный способ представить ровно то же самое.

---

> https://gobyexample.com/generics (VPN!)
