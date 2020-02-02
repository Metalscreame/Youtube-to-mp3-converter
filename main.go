package main

import (
	"log"

	"github.com/Metalscreame/Youtube-to-mp3-converter/usecases/audio"
	"github.com/Metalscreame/Youtube-to-mp3-converter/usecases/youtuber"
)

func main() {
	u := youtuber.YouTuberUC{}

	fileName, err := u.Download("https://www.youtube.com/watch?v=IY_4M7crOp4")
	if err != nil {
		log.Fatal(err)
	}

	mp3 := audio.FFmpegConverter{}
	audioFile, err := mp3.ConvertToMp3(fileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Finished %v", audioFile)
}
