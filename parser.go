package main
import (
    "fmt"
    "math"
)

func get_grammar() map[string] (map[string] float64) {
    grammar := map[string] map[string] float64 {}
    temp := make(map[string] float64)
    temp["S"] = 0.8
    grammar["NP VP"] = make(map[string] float64)
    grammar["NP VP"] = temp

    temp = make(map[string] float64)
    temp["S"] = 0.1
    grammar["X1 VP"] = make(map[string] float64)
    grammar["X1 VP"] = temp

    temp = make(map[string] float64)
    temp["X1"] = 1.0
    grammar["Aux NP"] = make(map[string] float64)
    grammar["Aux NP"] = temp

    temp = make(map[string] float64)
    temp["S"] = 0.01
    temp["NP"] = 0.075
    temp["Nom"] = 0.25
    temp["VP"] = 0.1
    temp["V"] = 0.5
    grammar["book"] = make(map[string] float64)
    grammar["book"] = temp

    temp = make(map[string] float64)
    temp["S"] = 0.004
    temp["NP"] = 0.075
    temp["Nom"] = 0.25
    temp["V"] = 0.2
    temp["VP"] = 0.04
    grammar["books"] = make(map[string] float64)
    grammar["books"] = temp

    temp = make(map[string] float64)
    temp["S"] = 0.006
    temp["VP"] = 0.06
    temp["V"] = 0.3
    grammar["prefer"] = make(map[string] float64)
    grammar["prefer"] = temp

    temp = make(map[string] float64)
    temp["S"] = 0.05
    temp["VP"] = 0.5
    grammar["V NP"] = make(map[string] float64)
    grammar["V NP"] = temp

    temp = make(map[string] float64)
    temp["S"] = 0.03
    grammar["VP NP"] = make(map[string] float64)
    grammar["VP NP"] = temp

    temp = make(map[string] float64)
    temp["NP"] = 0.1
    grammar["I"] = make(map[string] float64)
    grammar["I"] = temp

    temp = make(map[string] float64)
    temp["NP"] = 0.02
    grammar["he"] = make(map[string] float64)
    grammar["he"] = temp

    temp = make(map[string] float64)
    temp["NP"] = 0.02
    grammar["she"] = make(map[string] float64)
    grammar["she"] = temp

    temp = make(map[string] float64)
    temp["NP"] = 0.06
    grammar["me"] = make(map[string] float64)
    grammar["me"] = temp

    temp = make(map[string] float64)
    temp["NP"] = 0.16
    grammar["Houston"] = make(map[string] float64)
    grammar["Houston"] = temp

    temp = make(map[string] float64)
    temp["NP"] = 0.04
    grammar["NWA"] = make(map[string] float64)
    grammar["NWA"] = temp

    temp = make(map[string] float64)
    temp["NP"] = 0.3
    grammar["Nom PP"] = make(map[string] float64)
    grammar["Nom PP"] = temp

    temp = make(map[string] float64)
    temp["NP"] = 0.075
    temp["Nom"] = 0.25
    grammar["flight"] = make(map[string] float64)
    grammar["flight"] = temp

    temp = make(map[string] float64)
    temp["NP"] = 0.075
    temp["Nom"] = 0.25
    grammar["money"] = make(map[string] float64)
    grammar["money"] = temp

    // temp = make(map[string] float64)
    // grammar["V NP"] = make(map[string] float64)
    // grammar["V NP"] = temp

    temp = make(map[string] float64)
    temp["VP"] = 0.3
    grammar["VP PP"] = make(map[string] float64)
    grammar["VP PP"] = temp

    temp = make(map[string] float64)
    temp["PP"] = 1.0
    grammar["Prep NP"] = make(map[string] float64)
    grammar["Prep NP"] = temp

    temp = make(map[string] float64)
    temp["Prep"] = 0.5
    grammar["with"] = make(map[string] float64)
    grammar["with"] = temp

    temp = make(map[string] float64)
    temp["Prep"] = 0.5
    grammar["at"] = make(map[string] float64)
    grammar["at"] = temp

    return grammar
}

func removeDuplicates(elements []string) []string {
    encountered := map[string]bool{}
    result := []string{}

    for v := range elements {
    	if encountered[elements[v]] == true {
    	    // Do not add duplicate.
    	} else {
    	    encountered[elements[v]] = true
    	    result = append(result, elements[v])
    	}
    }
    return result
}

// func get_grammar() map[string] (map[string] float64) {
//     grammar := map[string] map[string] float64 {}
//     temp := make(map[string] float64)
//     temp["S"] = 1.0
//     grammar["B C"] = make(map[string] float64)
//     grammar["B C"] = temp
//
//     temp = make(map[string] float64,0)
//     temp["A"] = 0.5
//     grammar["B A"] = temp
//
//     temp = make(map[string] float64,0)
//     temp["A"] = 0.5
//     temp["C"] = 0.5
//     grammar["a"] = temp
//
//     temp = make(map[string] float64,0)
//     temp["B"] = 0.7
//     grammar["C C"] = temp
//     grammar["b"] = temp
//
//     temp = make(map[string] float64,0)
//     temp["C"] = 0.3
//     temp["S"] = 1.0
//     grammar["A B"] = temp
//
//     return grammar
// }

func get_combinations(grammar map[string] map[string] float64,cyk_grid [][][]string,cyk_prob_grid [][]map[string] float64,a int,b int) []string {
    temp := make([]string,0)
    temp_prob := make(map[string] float64,0)
    for _,i := range cyk_grid[a-1][b] {
        for _,j := range cyk_grid[0][a+b] {
            for k := range grammar[i +" "+j] {
                temp = append(temp,k)
                if(temp_prob[k] > 0.0) {
                    temp_prob[k] = math.Max(temp_prob[k],grammar[i +" "+j][k]*cyk_prob_grid[a-1][b][i]*cyk_prob_grid[0][a+b][j])
                } else {
                    temp_prob[k] = grammar[i +" "+j][k]*cyk_prob_grid[a-1][b][i]*cyk_prob_grid[0][a+b][j]
                }
            }
        }
    }

    if(a > 1) {
        for _,i := range cyk_grid[0][b] {
            for _,j := range cyk_grid[a-1][b+1] {
                for k := range grammar[i +" "+j] {
                    temp = append(temp,k)
                    if(temp_prob[k] > 0.0) {
                        temp_prob[k] = math.Max(temp_prob[k],grammar[i +" "+j][k]*cyk_prob_grid[0][b][i]*cyk_prob_grid[a-1][b+1][j])
                    } else {
                        temp_prob[k] = grammar[i +" "+j][k]*cyk_prob_grid[0][b][i]*cyk_prob_grid[a-1][b+1][j]
                    }
                }
            }
        }
        temp = removeDuplicates(temp)
    }
    cyk_prob_grid[a] = append(cyk_prob_grid[a],temp_prob)
    return temp
}

func cyk_parser(tokens []string) {
    grammar := get_grammar()
    cyk_grid := make([][][]string,len(tokens))
    cyk_prob_grid := make([][]map[string] float64,len(tokens))
    for i,_ := range tokens {
        temp := make([]string,0)
        temp_prob := make(map[string] float64,0)
        for j:= range grammar[tokens[i]] {
            temp = append(temp,j)
            temp_prob[j] = grammar[tokens[i]][j]
        }
        cyk_grid[0] = append(cyk_grid[0],temp)
        cyk_prob_grid[0] = append(cyk_prob_grid[0],temp_prob)
    }

    for i:= 1; i<len(tokens);i++ {
        cyk_grid[i] = make([][]string,0)
        for j:=0;j< len(tokens) - i;j++ {
            val := get_combinations(grammar,cyk_grid,cyk_prob_grid,i,j)
            cyk_grid[i] = append(cyk_grid[i],val)
        }
    }
    for i,_ := range cyk_prob_grid {
            fmt.Printf("%v\n",cyk_prob_grid[i])
    }
}
