//basic hello go
package main

import "fmt"
import "time"
import "math"
func main(){

fmt.Println("Hello My Time", time.Now())

//strings

fmt.Println("go" + "lang")

fmt.Println("7.0/3.0=", 7.0/3.0)

fmt.Println("1+1=", 1+1)

fmt.Println(true && false)

fmt.Println(true || false)

fmt.Println(!true)

//variables

var a string = "initial"
fmt.Println(a)

var b,c int = 1,2
fmt.Println(b,c)

var d = true
fmt.Println(d)

var e int
fmt.Println (e)

f := "string"
fmt.Println(f)

//constraints

const s string = "constant"
fmt.Println(s)

const n = 500000000

const m = 3e20 / n
fmt.Println(int64(m))
fmt.Println(math.Sin(n)) 

// for example

i := 1
for i <= 3 {
fmt.Println(i)
i = i + 1
}

}
