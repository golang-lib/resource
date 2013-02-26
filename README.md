# gosexy/resource

A simple resource downloader.

```go
import "github.com/gosexy/resource"
```

Example:

```go
uri := `http://upload.wikimedia.org/wikipedia/commons/b/be/Kukenan_Roraima_GS.jpg`

file, err := resource.Download(uri, "downloads")

if err != nil {
	t.Errorf("Could not download: %s", err.Error())
}

log.Printf("Downloaded to: %s\n", file)
```

Output:

```
2013/02/25 22:41:18 Downloaded to: downloads/f73/14aea91da675ba7f3846d477edb24fc796e13/Kukenan_Roraima_GS.jpg
```
