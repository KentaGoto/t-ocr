package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func imgCheck(format string, img string, path string, lang string) {
	if strings.Contains(format, img) {
		cmd := exec.Command("tesseract", path, path, "-l", lang)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	} else {

	}
}

// Supported image types: gif, jpeg, png, bmp
func main() {
	var arg string
	if len(os.Args) != 3 {
		fmt.Println("The number of arguments specified is incorrect.")
		os.Exit(1)
	} else {
		arg = os.Args[1]
	}

	lang := os.Args[2]

	paths := dirwalk(arg)
	fmt.Println("Processing...")
	imgs := [...]string{"jpeg", "jpg", "bmp", "png", "gif"}

	for _, path := range paths {
		fmt.Println(path)
		f, _ := os.Open(path)
		defer f.Close()

		_, format, err := image.DecodeConfig(f) // Get the image file format
		if err != nil {
			fmt.Println(err)
		}

		for _, img := range imgs {
			imgCheck(format, img, path, lang)
		}
	}
}
