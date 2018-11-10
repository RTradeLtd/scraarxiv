package magnifier

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/RTradeLtd/scraarxiv/utils"
)

// DownloadFiles is used to download files from the given urls, returning their file paths
func (g *Glass) DownloadFiles(urls []string, maxResults int) ([]string, error) {
	var filePaths []string
	rand := utils.GenerateRandomUtils()
	count := 0
	for _, v := range urls {
		count++
		fmt.Println("count ", count)
		if count > maxResults {
			break
		}
		randPath := rand.GenerateString(15, utils.LetterBytes)
		out, err := os.Create(fmt.Sprintf("/tmp/%s", randPath))
		if err != nil {
			return nil, err
		}
		defer out.Close()
		fmt.Println("downloading pdf from url ", v)
		// read the pdf
		resp, err := http.Get(v)
		if err != nil {
			continue
		}
		// write to file
		if _, err = io.Copy(out, resp.Body); err != nil {
			return nil, err
		}
		filePaths = append(filePaths, fmt.Sprintf("/tmp/%s", randPath))
	}
	return filePaths, nil
}
