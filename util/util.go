package main

func strInStrSlice(str string, slice []string) bool {
    for _, s := range slice {
        if s == str {
            return true
        }
    }

    return false
}
