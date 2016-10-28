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

func get_grammar() map[string] (map[string] float64) {
    grammar := map[string] map[string] float64 {}
    temp := make(map[string] float64)
    temp["S"] = 1.0
    grammar["B C"] = make(map[string] float64)
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

func get_combinations(grammar map[string] map[string] float64,cyk_grid [][][]string,cyk_prob_grid [][]map[string] float64,a int,b int) []string {
    temp := make([]string,0)
    temp_prob := make(map[string] float64,0)
    for _,i := range cyk_grid[a-1][b] {
        for _,j := range cyk_grid[0][a+b] {
            for k := range grammar[i +" "+j] {
                temp = append(temp,k)
                temp_prob[k] = grammar[i +" "+j][k]*cyk_prob_grid[a-1][b][i]*cyk_prob_grid[0][a+b][j]
            }
        }
    }

    if(a > 1) {
        for _,i := range cyk_grid[0][b] {
            for _,j := range cyk_grid[a-1][b+1] {
                for k := range grammar[i +" "+j] {
                    temp = append(temp,k)
                    temp_prob[k] = grammar[i +" "+j][k]*cyk_prob_grid[0][b][i]*cyk_prob_grid[a-1][b+1][j]
                }
            }
        }
        temp = removeDuplicates(temp)
    }
    cyk_prob_grid[a] = append(cyk_prob_grid[a],temp_prob)
    return temp
}

func parser(tokens []string) {
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
        //cyk_prob_grid[i] = make([]map[string] float64,0)
        for j:=0;j< len(tokens) - i;j++ {
            //val := make([]string,0)
            val := get_combinations(grammar,cyk_grid,cyk_prob_grid,i,j)
            // for _,k := range temp {
            //     val = append(val,k)
            // }
            // if(i > 1) {
            //     temp2 := get_combinations(grammar,cyk_grid[0][j],cyk_grid[i-1][j+1])
            //     for _,k := range temp2 {
            //         val = append(val,k)
            //     }
            //     val = removeDuplicates(val)
            // }
            cyk_grid[i] = append(cyk_grid[i],val)
        }
    }
    for i,_ := range cyk_prob_grid {
            fmt.Printf("%v\n",cyk_prob_grid[i])
    }
}
