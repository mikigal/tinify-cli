package main

import (
	"./tinify"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := make([]string, len(os.Args)-1)
	for i := 0; i < len(args); i++ {
		args[i] = os.Args[i+1]
	}

	if len(args) == 0 {
		showHelp()
		return
	}

	if args[0] == "key" {
		setKey(args)
		return
	}

	key, err := tinify.GetKey()
	if key == "" || err != nil {
		log.Fatal("Please set API KEY with command: tinify key <api key>")
	}

	fmt.Println("Using API KEY: " + key)
	fmt.Println()

	if len(args) == 1 {
		compress(args[0], args[0], key)
		return
	}
	if len(args) == 2 {
		compress(args[0], args[1], key)
		return
	}

	showHelp()
}

func showHelp() {
	fmt.Println("Help:")
	fmt.Println("tinify key <api key> - Set API KEY. You can get it on: http://tinypng.con")
	fmt.Println("tinify <input> [output] - Tinify selected file, if new name is not selected file will be saved with 'tiny-' prefix")
}

func setKey(args []string) {
	if len(args) != 2 {
		log.Fatal("Correct usage: tinify key <api key>")
	}

	key := args[1]
	err := tinify.SaveKey(key)
	if err != nil {
		fmt.Printf("Can't save API KEY to ~/.tinify! %s", err)
	}

	fmt.Println("Your API KEY has been set correctly! You can use Tinify CLI now!")
}

func compress(input string, output string, key string) {
	if strings.HasSuffix(input, "*") {
		input = strings.TrimSuffix(input, "*")
	}

	file, err := os.Open(input)
	if err != nil {
		log.Fatal("Selected file does not exists!")
	}

	if tinify.IsDirectory(file) {
		compressDirectory(input, output, key)
	} else {
		compressFile(*file, output, key)
	}
}

func compressDirectory(path string, output string, key string) {
	dir, err := tinify.GetFilesFromDir(path)
	tinify.Check(err)

	for _, path := range dir {
		file, err := os.Open(path)
		tinify.Check(err)

		compressFile(*file, output, key)
	}
}

func compressFile(input os.File, output string, key string) {
	if !(strings.HasSuffix(input.Name(), ".png") || strings.HasSuffix(input.Name(), ".jpg") || strings.HasSuffix(input.Name(), ".jpeg")) {
		log.Fatal("Supported file types: png, jpg, jpeg")
	}

	o, err := os.Open(output)

	if err == nil && tinify.IsDirectory(o) {
		output = output + "/tiny-" + strings.Split(input.Name(), "/")[len(strings.Split(input.Name(), "/"))-1]
	} else {
		name := strings.Split(input.Name(), "/")[len(strings.Split(input.Name(), "/"))-1]
		output = strings.Replace(output, name, "tiny-"+name, 1)
	}

	res, err := tinify.Upload(key, &input)
	tinify.Check(err)

	res.Download(output, key)
	fmt.Printf("Panda just saved you %d%% (%d KB)! \n", res.CalcPercent(), res.CalcSizeKB())
}
