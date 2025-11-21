package search

import "strings"

func SuggestedProducts(products []string, searchWord string) [][]string {

	 sort.Strings(products)
    result := [][]string{}
    prefix := ""

    for _, ch := range searchWord {
        prefix += string(ch)

        // Binary search for the first index where `prefix` can be inserted
        i := sort.Search(len(products), func(j int) bool {
            return products[j] >= prefix
        })

        suggestions := []string{}
        // Check next 3 items for matching prefix
        for k := i; k < len(products) && len(suggestions) < 3; k++ {
            if len(products[k]) >= len(prefix) && products[k][:len(prefix)] == prefix {
                suggestions = append(suggestions, products[k])
            } else {
                break
            }
        }

        result = append(result, suggestions)
    }

    return result
}