package resource

import (
	"log"
	"testing"
)

func TestDownload(t *testing.T) {
	uri := `http://upload.wikimedia.org/wikipedia/commons/b/be/Kukenan_Roraima_GS.jpg`

	file, err := Download(uri, "downloads")

	if err != nil {
		t.Errorf("Could not download: %s", err.Error())
	}

	log.Printf("Downloaded to: %s\n", file)
}
