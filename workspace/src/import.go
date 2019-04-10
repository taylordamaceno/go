package main 

import (

     "fmt"
     "os"
)

func main () {

       //	fmt.Println(len(os.Args))
       //	fmt.Println(os.Args[0])
        if len(os.Args) <= 1 {
	fmt.Println("sem argumentos")	
	os.Exit(1)
	} 
        fmt.Println("recebi argumento")
        fmt.Println(os.Args[1])

	
}
