package service

import (
    "strings"
    "qrzero/internal/v1/repository"
)

func FindUncreatedAccounts(files []repository.FileInfo, customers []repository.Customer) []repository.Customer {
    exist := make(map[string]struct{})
    for _, f := range files {
        name := strings.TrimSuffix(f.Name, ".png")
        exist[name] = struct{}{}
    }
    var result []repository.Customer
    for _, c := range customers {
        if _, ok := exist[c.AccountNum]; !ok {
            result = append(result, c)
        }
    }
    return result
}
