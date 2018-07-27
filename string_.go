package main
import "fmt"

func main(){
    s:="Hello World"
    a:=[]rune(s)
    a[0] = 'W'
    s2:=string(a)

    fmt.Println(s2)
}
