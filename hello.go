package main

import "fmt"
import "os"

//comment
func main() {
    var x string = "Hello, I am Shadow2"
    
    //declare constant and infer type from RHS without type specified
    const y = "Bringer of darkness"

    //declare, initialize and infer type from RHS with ':=' assignment
    l := len(x)
    x += "\n" + y
    fmt.Println(x)
    fmt.Println(l)
    //fmt.Println("Or am I " + stringutil.Reverse("Shadow") + "?")
    os.Exit(0)
}
