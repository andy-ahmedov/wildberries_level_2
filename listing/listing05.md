Что выведет программа? Объяснить вывод программы.

```go
package main
type customError struct {
	msg string
}
func (e *customError) Error() string {
	return e.msg
}
func test() *customError {
	{
		// do something
	}
	return nil
}
func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод: "error"

Интерфейс хранит в себе тип интерфейса и тип самого значение.

Значение любого интерфейса, не только error, является nil в случае когда И значение И тип являются nil.

Функция test возвращает nil типа *customError, результат мы сравниваем с nil типа nil, откуда и следует их неравенство.
```