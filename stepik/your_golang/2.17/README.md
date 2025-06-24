# 2.17 Strings

> https://stepik.org/lesson/1500119/step/3

`2025-06-24 18:58`

```go
fmt.Println(strings.TrimSpace(message))
fmt.Println(strings.TrimSpace(strings.ToLower(message)))
fmt.Println(strings.Contains(message, "Go"))
```

```go
fmt.Println(strings.TrimSpace(message))
fmt.Println(strings.ToLower(strings.TrimSpace(message)))
fmt.Println(strings.HasPrefix(strings.ToLower(strings.TrimSpace(message)), "go"))
```
