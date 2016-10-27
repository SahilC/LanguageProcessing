package main
import (
    "fmt"
)
func parser(tokens []string) {
    cyk_grid := make([][]string,len(tokens))
    for i,_ := range tokens {
        cyk_grid[0] = append(cyk_grid[0],tokens[i])
    }
    fmt.Printf("%#v\n",cyk_grid)
}
