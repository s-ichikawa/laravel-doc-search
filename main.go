package main

import (
    "io/ioutil"
    "regexp"
    "fmt"
    "os"
    "bufio"
    "os/exec"
)

var titles = []string{}

func init() {
    dir, _ := ioutil.ReadDir("./docs")
    md_pattern := regexp.MustCompile(`.*.md`)
    for _, fi := range dir {
        if !md_pattern.MatchString(fi.Name()) {
            continue
        }
        titles = append(titles, regexp.MustCompile(`.md$`).ReplaceAllString(fi.Name(), ""))

    }
}

func main() {
    for _, url := range titles {
        fmt.Println(url)
    }
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        keyword := s.Text()

        if searchTitle(keyword) {
            open(keyword)
        }

        if result, title := searchText(keyword); result {
            open(title)
        }
    }
}

func open(dir string) {
    s := bufio.NewScanner(os.Stdin)
    url := "https://laravel.com/docs/5.3/" + dir
    fmt.Println("found: " + url)
    fmt.Print("open? [y/n]:")
    s.Scan()
    if (s.Text() == "y") {
        exec.Command("open", url).Start()
    }
}

func searchTitle(keyword string) bool {
    if inArray(keyword, titles) {
        return true
    }
    return false
}

func searchText(keyword string) (bool, string) {
    pattern := regexp.MustCompile(keyword)
    for _, title := range titles {
        file, _ := ioutil.ReadFile("./docs/" + title + ".md")
        if pattern.MatchString(string(file)) {
            return true, title
        }
    }
    return false, ""
}

func inArray(needle string, list []string) bool {
    for _, str := range list {
        if str == needle {
            return true
        }
    }
    return false
}
