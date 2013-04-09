# gosexy/resource

A simple file downloader.

```go
import "menteslibres.net/gosexy/resource"
```

Example:

```go
uri := `http://upload.wikimedia.org/wikipedia/commons/b/be/Kukenan_Roraima_GS.jpg`

file, err := resource.Download(uri, "downloads")

if err != nil {
	log.Fatalf("Could not download: %s", err.Error())
}

log.Printf("Downloaded to: %s\n", file)
```

Output:

```
2013/04/09 21:02:46 Downloaded to: downloads/f731/4aea/91da/675ba7f3846d477edb24fc796e13/Kukenan_Roraima_GS.jpg
```

## Documentation

See the [online docs](http://godoc.org/menteslibres.net/gosexy/resource).
