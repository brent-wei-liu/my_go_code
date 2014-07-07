package main

import (
    "fmt"
    "math/rand"
)

func main() {
    n := 100000
    interval := 100
    var desiredStdDev, desiredMean float64 = 10 , 80 
    dis := make([]int, interval)
    i:=0
    for i < n {
        sample := rand.NormFloat64() * desiredStdDev + desiredMean
        index := int(sample)
        if index < 0 {
            index = 0
        }
        if index >= interval{
            index = interval - 1
            //fmt.Println(index," continue")
        }
        //fmt.Println(sample, index)
        dis[index] ++
        i++
    }
    for i, v := range dis{
        fmt.Print(i,"\t",v,"\r\n")
    }
}

