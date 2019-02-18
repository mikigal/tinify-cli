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
		doStuff(args[0], "tiny-"+args[0], key)
		return
	}
	if len(args) == 2 {
		doStuff(args[0], args[1], key)
		return
	}

	showHelp()
}

func showHelp() {
	fmt.Println("Help:")
	fmt.Println("tinify key <api key> - Set API KEY. You can get it on: http://tinypng.con")
	fmt.Println("tinify <target> [new name] - Tinify selected file, if new name is not selected file will be saved with 'tiny-' prefix")
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

func doStuff(target string, name string, key string) {
	if !(strings.HasSuffix(target, ".png") || strings.HasSuffix(target, ".jpg") || strings.HasSuffix(target, ".jpeg")) {
		log.Fatal("Supported file types: png, jpg, jpeg")
	}

	file, err := os.Open(target)
	if err != nil {
		log.Fatal("Selected file does not exists!")
	}

	res, err := tinify.Upload(key, file)
	tinify.Check(err)

	res.Download(name, key)
	fmt.Printf("Panda just saved you %d%% (%d KB)!", res.CalcPercent(), res.CalcSizeKB())
}
