package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "path/filepath"
)



func main() {
    tags := walkFiles(os.Args[1])
    uniqueTags := getUniqueTags(tags)
    for _, tag := range uniqueTags {
        fmt.Println(tag)
    }

}

func Find(slice []string, val string) bool {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}

func getUniqueTags(tags map[string][]string) []string {
    uniqueTags := []string{}
    for _, tags := range tags {
        for _, tag := range tags {
            if Find(uniqueTags, tag) != true {
                uniqueTags = append(uniqueTags, tag)
            }
        }
    }
    return uniqueTags
}

func walkFiles(rootPath string) map[string][]string {
    tags := map[string][]string{}

    root := rootPath
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        fi, err := os.Stat(path)
        if err != nil {
            panic(err)
        }
        mode := fi.Mode()
        if mode.IsRegular() {
            f, err := os.Open(path)
            defer f.Close()
            if err != nil {
                panic(err)
            }
            rd := bufio.NewReader(f)
            line, err := rd.ReadString('\n')
            tags[path] = strings.Split(strings.TrimSuffix(line, "\n"), ",")
        }
        return nil
    })
    if err != nil {
        panic(err)
    }
    return tags
}
