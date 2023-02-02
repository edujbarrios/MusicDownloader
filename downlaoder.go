package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rylio/ytdl"
)

func main() {
	fmt.Println("Enter the YouTube video URL:")
	var videoURL string
	fmt.Scanln(&videoURL)

	fmt.Println("Select the format to convert the video to:")
	fmt.Println("1. WAV")
	fmt.Println("2. MP3")
	var format int
	fmt.Scanln(&format)

	video, err := ytdl.GetVideoInfo(videoURL)
	if err != nil {
		fmt.Println("Error getting video information:", err)
		return
	}

	audioFormat := video.Formats.Extremes(ytdl.FormatAudioBitrateKey, true)[0]

	title := video.Title
	title = strings.Replace(title, "(Official Video)", "", -1)
	title = strings.Replace(title, "(Official Audio)", "", -1)
	title = strings.Replace(title, "(Videoclip Oficial)", "", -1)
	title = strings.Replace(title, "(VIDEOCLIP OFICIAL)", "", -1)
	title = strings.Replace(title, "(OFFICIAL AUDIO)", "", -1)
	title = strings.Replace(title, "(OFFICIAL VIDEO)", "", -1)
	
	
	

	fileName := title + "." + audioFormat.Extension
	if format == 2 {
		fileName = title + ".mp3"
	}

	err = audioFormat.Download(fileName)
	if err != nil {
		fmt.Println("Error downloading video:", err)
		return
	}

	// Move the file to the "Music" folder
	homeDir, _ := os.UserHomeDir()
	musicDir := filepath.Join(homeDir, "Music")
	newPath := filepath.Join(musicDir, fileName)
	err = os.Rename(fileName, newPath)
	if err != nil {
		fmt.Println("Error moving file to music folder:", err)
		return
	}

	fmt.Println("Video converted successfully to", fileName, "and added to", musicDir)
}
