# tinify-cli
CLI tool to compress png/jpg images with tinypng.com, tinyjpg.com, tinify.com

# How to use
Before use copy binary to your folder in PATH. For Unix or macOS copy it to `/usr/local/bin`

For Windows copy it to any directory and add this directory to PATH. After editing PATH restart your Terminal App

Remember to set API KEY before first use! Get it on https://tinypng.com. You can compress 500 images per month for free, you don't need any credit card, etc for free usage. After exceed limit you will get error.

```bash
tinify key <api key> # Only before first use
tinify <input> [output] # Input and output can be directory
```

![img](https://i.imgur.com/8wSYDKW.png)