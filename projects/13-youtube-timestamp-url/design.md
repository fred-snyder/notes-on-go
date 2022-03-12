# YoutubeTime (Go)

- [ ] Start the project

## v1 notes

Version 2 to could add functionality for the clipboard. To read URLs to and from the clipboard.

```golang
import LIB_FOR_READ_STDIN

// store user link
var link string

// time to start the video
var (
	hour int
	minute int
	second int
)

// extract youtube video ID from an URL
// return a base URL
func extractID(string) string {
	// how can I get the ID from an arbitrary YT-link
	// for example: for playlist URLs or URLs with "crap" in it
	// what are the conventions for YT-urls?

	// return two values? error??
}

// 
func createLink(string) string {
	// combine ID with base link and new time
	// return new link

}

func main() {
	// ask user for input

	ID := extractID(link)
	timeURL := createLink(ID)

	// output new URL
	fmt.Println(timedURL)
}
```


```plain
=======
$ youtubetime.exe
=======

# This application generates a youtube link that starts at a specific time.
Enter a Youtube Link: INPUT

Start at video at:
- Hour: INPUT
- Minute: INPUT
- Second: INPUT

New link:
$ OUTPUT URL
```
