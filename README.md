gorss is a library for easily parse RSS feeds and convert them to Go structs

```go
 package main

 import "github.com/mvader/gorss"
 import "fmt"

 func main() {
   rss, _ := gorss.LoadFeed("url to my feed")

   fmt.Println(*rss)
 }
 ``