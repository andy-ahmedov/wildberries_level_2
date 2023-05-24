Что выведет программа? Объяснить вывод программы.

```go
package main
import (
	"fmt"
	"math/rand"
	"time"
)
func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}
func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Программа выведет отправленные в канал 'с' значения из каналов 'a' и 'b', после будет выводить zero values, так как каналы 'a' и 'b' после отправления нужных в них значений были закрыты, а получение из них в функции merge не останавливается, соотвественно цикл получения значений в main будет продолжаться вечно
```