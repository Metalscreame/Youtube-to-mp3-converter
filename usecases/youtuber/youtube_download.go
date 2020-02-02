package youtuber

import (
	"fmt"
	"os"
	"strings"

	"github.com/rylio/ytdl"
)

const tempFile = "%v.mp4"

type YouTuberUC struct{}

// Download downloads video from youtube and saves it
func (yt *YouTuberUC) Download(url string) (fileName string, err error) {
	vid, err := ytdl.GetVideoInfo(url)
	if err != nil {
		return "", fmt.Errorf("failed to get video info %v", err)
	}

	title := strings.ReplaceAll(vid.Title,"|", "")
	title = strings.ReplaceAll(title,":", "")
	title = strings.ReplaceAll(title,"-", "")

	fileName = fmt.Sprintf(tempFile, title)
	file, err := os.Create(fileName)
	if err != nil {
		return
	}

	defer file.Close()

	if err = vid.Download(vid.Formats[0], file); err != nil {
		return
	}

	return fileName, nil
}
