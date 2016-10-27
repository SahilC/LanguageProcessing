package main
import (
    "fmt"
)
// func getGrammar() map[string] (map[string] float64) {
//     grammar := map[string] map[string] float64 {}
//     temp := make(map[string] float64)
//     temp["S"] = 1.0
//     grammar["NP VP"] = make(map[string] float64)
//     grammar["NP VP"] = temp
//
//     temp = make(map[string] float64,0)
//     temp["PP"] = 1.0
//     grammar["P NP"] = temp
//
//     temp = make(map[string] float64,0)
//     temp["VP"] = 0.7
//     grammar["V NP"] = temp
//
//     temp = make(map[string] float64,0)
//     temp["VP"] = 0.3
//     grammar["VP PP"] = temp
//
//     return grammar
// }

func get_grammar() map[string] (map[string] float64) {
    grammar := map[string] map[string] float64 {}
    temp := make(map[string] float64)
    temp["S"] = 1.0
    grammar["B C"] = make(map[string] float64)
    // grammar["A B"] = make(map[string] float64)
    // grammar["A B"] = temp
    grammar["B C"] = temp

    temp = make(map[string] float64,0)
    temp["A"] = 0.5
    grammar["B A"] = temp

    temp = make(map[string] float64,0)
    temp["A"] = 0.5
    temp["C"] = 0.5
    grammar["a"] = temp

    temp = make(map[string] float64,0)
    temp["B"] = 0.7
    grammar["C C"] = temp
    grammar["b"] = temp

    temp = make(map[string] float64,0)
    temp["C"] = 0.3
    temp["S"] = 1.0
    grammar["A B"] = temp

    return grammar
}

func get_combinations(grammar map[string] map[string] float64,left []string,right []string) []string{
    temp := make([]string,0)
    for _,i := range left {
        for _,j := range right {
            //fmt.Printf("%#v=============\n",i+ " "+j)
            //fmt.Printf("%#v=============\n",i +" "+j)
            for k := range grammar[i +" "+j] {
                temp = append(temp,k)
            }
        }
    }
    //fmt.Printf("%#v\n",temp)
    return temp
}

func parser(tokens []string) {
    grammar := get_grammar()
    cyk_grid := make([][][]string,len(tokens))
    for i,_ := range tokens {
        temp := make([]string,0)
        for j:= range grammar[tokens[i]] {
            temp = append(temp,j)
        }
        //fmt.Printf("%v\n",temp)
        cyk_grid[0] = append(cyk_grid[0],temp)
    }

    //fmt.Printf("%#v\n%#v\n",cyk_grid,grammar)
    for i:= 1; i<len(tokens);i++ {
        cyk_grid[i] = make([][]string,0)
        for j:=0;j< len(tokens) - i;j++ {
            //cyk_grid[i] = append(cyk_grid[i],tokens[j])
            //fmt.Printf("%#v %d %d\n",cyk_grid,i-1,j)
            val := get_combinations(grammar,cyk_grid[i-1][j],cyk_grid[0][i+j])
            cyk_grid[i] = append(cyk_grid[i],val)

            if(i > 2) {
                val = get_combinations(grammar,cyk_grid[0][j],cyk_grid[i-1][j+1])
                cyk_grid[i] = append(cyk_grid[i],val)
            }
            // fmt.Printf("(((%d %d),",i-1,j)
            // fmt.Printf("(%d %d)),",0, i + j)
            // fmt.Printf("((%d %d),",0,j)
            // fmt.Printf("(%d %d))),",i-1,j+1)
        }
        //fmt.Printf("\n")
    }
    fmt.Printf("%#v\n",cyk_grid)
}
