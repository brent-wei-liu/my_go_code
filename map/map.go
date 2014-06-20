/*kvs := make(map[string]*list.List)


Go 43 Exercise: Maps

Implement WordCount. It should return a map of the counts of each “word” in the string s. 
The wc.Test function runs a test suite against the provided function and prints success or failure.
You might find strings.Fields helpful.

*/

package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    strs := strings.Fields(s)
    m := make(map[string]int)
    for _,v := range strs{
        _,ok := m[v]
        if !ok{
           m[v] = 0 
        }
        m[v] ++
    }
    return m
}

func main() {
    wc.Test(WordCount)
}
