package gochallenge3
import (
    "net/http"
    "fmt"
    "io"
    "os"
)

func Download(urls []ImageURL, downloadDir string) ([]string, error) {
    var filePaths = make([]string, len(urls))

    for i, url := range urls {
        CommonLog.Printf("download [%d] => \"%s\"", i, url)

        response, err := http.Get(string(url))
        if err != nil {
            return nil, fmt.Errorf("error while downloading %s: %s", url, err)
        }
        defer response.Body.Close()

        filePaths[i] = fmt.Sprintf("%s/thumb.%d", downloadDir, i)

        output, err := os.Create(filePaths[i])
        if err != nil {
            return nil, fmt.Errorf("error while creating %s: %s", filePaths[i], err)
        }
        defer output.Close()

        _, err = io.Copy(output, response.Body)
        if err != nil {
            return nil, fmt.Errorf("error while downloading", url, "-", err)
        }

        CommonLog.Printf("downloaded %s to %s\n", url, filePaths[i])
    }


    return filePaths, nil
}
