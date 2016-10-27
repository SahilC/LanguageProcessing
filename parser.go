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
    //cyk_grid := make([][]string,len(tokens))
    // for i,_ := range tokens {
    //     cyk_grid[0] = append(cyk_grid[0],tokens[i])
    // }
    //grammar := getGrammar()

    //fmt.Printf("%#v\n%#v\n",cyk_grid,grammar)
    for i:=0; i<len(tokens);i++ {
        for j:=0;j< len(tokens) - i;j++ {
            //cyk_grid[i] = append(cyk_grid[i],tokens[j])
            if(i >= 1) {
                fmt.Printf("(((%d %d),",i-1,j)
                fmt.Printf("(%d %d)),",0, i + j)
                fmt.Printf("((%d %d),",0,j)
                fmt.Printf("(%d %d))),",i-1,j+1)
            }
        }
        fmt.Printf("\n")
    }
    //fmt.Printf("%#v\n",cyk_grid)
}
