package main
 
//TODO: Change the output of each command to output from uname
//Compile with go build uname.go
 
import (
        "fmt"
        "flag"
)
 
func main(){
    rPtr := flag.Bool("r", false, "kernel release");
    aPtr := flag.Bool("a", false, "all info");
    mPtr := flag.Bool("m", false, "machine");
    sPtr := flag.Bool("s", false, "kernel name");
 
 
    flag.Parse();
 
    if(*mPtr){
        //output of uname -m
        fmt.Println("x86_64");
    }else if(*aPtr){
        //output of uname -a, remember that in x.y.z z has to be lower than 255
        fmt.Println("Linux localhost 4.4.254-19935-g0ab0a71a40fd #1 SMP PREEMPT Mon Apr 19 20:04:52 PDT 2021 x86_64 Intel(R) Celeron(R) CPU N2840 @ 2.16GHz GenuineIntel GNU/Linux");
    }else if(*rPtr){
        //output of uname -r, remember that in x.y.z z has to be lower than 255
            fmt.Println("4.4.254-19935-g0ab0a71a40fd");
    }else if(*sPtr){
        //output of uname -s
            fmt.Println("Linux");
    }else{
            fmt.Println("Linux");
    }
}
//The output has been modified for machine use your own
