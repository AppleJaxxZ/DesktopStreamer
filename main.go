package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Live Desktop</title>
</head>
<body style="margin:0;background:black;">
<img src="/stream" style="width:100vw;height:auto;">
</body>
</html>
`)
	})

	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher, ok := w.(http.Flusher)
		if !ok {
			return
		}

		for {
			img := CaptureScreenJPEG()
			if img == nil {
				time.Sleep(500 * time.Millisecond)
				continue
			}

			fmt.Fprintf(w, "--frame\r\nContent-Type: image/jpeg\r\n\r\n")
			w.Write(img)
			fmt.Fprintf(w, "\r\n")
			flusher.Flush()

			time.Sleep(100 * time.Millisecond)
		}
	})

	fmt.Println("Streaming on :8080")
	http.ListenAndServe(":8080", nil)
}
