package audio

import (
	"fmt"
	"os"
	"os/exec"
)

type FFmpegConverter struct{}

func (F FFmpegConverter) ConvertToMp3(videoFileName string) (audioFileName string, err error) {
	audioFileName = fmt.Sprintf("%v.mp3",videoFileName)
	cmd := exec.Command("ffmpeg", "-i", videoFileName, audioFileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	return audioFileName, cmd.Run()
}
