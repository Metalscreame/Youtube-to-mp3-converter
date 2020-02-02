package youtuber

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rylio/ytdl"
)

const tempFile = "%v_%v.mp4"

type YouTuberUC struct{}

// Download downloads video from youtube and saves it
func (yt *YouTuberUC) Download(url string) (fileName string, err error) {
	vid, err := ytdl.GetVideoInfo(url)
	if err != nil {
		return "", fmt.Errorf("failed to get video info %v", err)
	}

	title := strings.ReplaceAll(vid.Title,"|", "")
	title = strings.ReplaceAll(vid.Title,":", "")

	fileName = fmt.Sprintf(tempFile, title, time.Now().Format("Mon Jan 3_04_00 PM"))
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
