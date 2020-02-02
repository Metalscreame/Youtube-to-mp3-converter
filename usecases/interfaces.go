package usecases

type Youtube interface {
	Download(url string) (fileName string, err error)
}

type AudioConverter interface {
	ConvertToMp3(videoFileName string) (audioFileName string, err error)
}