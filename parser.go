package main
import (
    "fmt"
)
func getGrammar() map[string] (map[string] float64) {
    grammar := map[string] map[string] float64 {}
    temp := make(map[string] float64)
    temp["NP VP"] = 1.0
    grammar["S"] = make(map[string] float64)
    grammar["S"] = temp

    temp = make(map[string] float64,0)
    temp["P NP"] = 1.0
    grammar["PP"] = temp

    temp = make(map[string] float64,0)
    temp["V NP"] = 0.7
    grammar["VP"] = temp

    temp = make(map[string] float64,0)
    temp["VP PP"] = 0.3
    grammar["VP"] = temp

    return grammar
}

func parser(tokens []string) {
    cyk_grid := make([][]string,len(tokens))
    for i,_ := range tokens {
        cyk_grid[0] = append(cyk_grid[0],tokens[i])
    }
    grammar := getGrammar()
    fmt.Printf("%#v\n%#v\n",cyk_grid,grammar)
}
