package audio

import "os/exec"

type FFmpegConverter struct{}

func (F FFmpegConverter) ConvertToMp3(videoFileName string) (audioFileName string, err error) {
	exec.Command("ffmpeg", "-i", videoFileName, output-audio.mp3)
}
