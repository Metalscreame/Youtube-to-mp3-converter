package youtuber

import (
	"fmt"
	"os"
	"time"

	"github.com/rylio/ytdl"
)

const tempFile = "temp_%v.mp4"

type YouTuberUC struct{}

// Download downloads video from youtube and saves it
func (yt *YouTuberUC) Download(url string) (fileName string, err error) {
	vid, err := ytdl.GetVideoInfo(url)
	if err != nil {
		return "", fmt.Errorf("failed to get video info %v", err)
	}

	file, err := os.Create(fmt.Sprintf(tempFile, time.Now().Format("Mon Jan 3 04 00 PM")))
	if err != nil {
		return
	}

	defer file.Close()

	if err = vid.Download(vid.Formats[0], file); err != nil {
		return
	}

	return tempFile, nil
}
