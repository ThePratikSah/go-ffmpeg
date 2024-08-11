# Go FFMPEG

So, I was wondering how HLS are made and found FFMPEG is the tool for this requirement.

I tried installing **FFMPEG** first using `brew`

```bash
brew install ffmpeg
```

And if you are on a Linux machine,

```bash
sudo apt update
sudo apt install ffmpeg
```

Once downloaded, you can start using it from the CLI.

```bash
ffmpeg -version
```

Start with file format conversion
```bash
ffmpeg -i ./video/video.mp4 ./output/output.mov
```

To run my code, make sure you have Go installed on your machine. If now, visit [Go](https://go.dev/doc/install) official website.

Once you have go installed, you can proceed with the code.

```bash
go version
```

To run this code
```bash
go run main.go
```

Follow [**FFMPEG docs**](https://www.ffmpeg.org/documentation.html) for some more advance concepts.