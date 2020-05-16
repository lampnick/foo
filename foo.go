package foo

import "fmt"

func Greet(name string) string {
    return fmt.Sprintf("%s, 你好！Foo Version:0.2.0", name)
}
