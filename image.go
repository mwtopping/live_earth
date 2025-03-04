package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func read_image(fname string) image.Image {
	fmt.Printf("Loading image from file: %v\n", fname)
	imageFile, err := os.Open(fname)
	if err != nil {
		log.Println("Error opening image")
		return nil
	}
	defer imageFile.Close()

	imageData, imageType, err := image.Decode(imageFile)
	if err != nil {
		log.Println("Error decoding image", err)
	}
	fmt.Printf("Loaded image of type %v\n", imageType)

	//imageSize := imageData.Bounds().Size()

	return imageData
}
