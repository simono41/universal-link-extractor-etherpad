package main

import (
    "fmt"
    "regexp"
    "github.com/go-rod/rod"
)

func main() {
    browser := rod.New().MustConnect()
    defer browser.MustClose()

    page := browser.MustPage("https://pad.stratum0.org/p/dc") // URL der Hauptseite
    page.MustWaitLoad()

    // Finden Sie das erste iFrame-Element
    outerIframeElement := page.MustElement("#editorcontainer > iframe:nth-child(1)") // Passen Sie den Selektor an

    // Wechseln Sie zum Kontext des ersten iFrames
    outerFrame := outerIframeElement.MustFrame()

    // Warten Sie, bis der Inhalt des ersten iFrames geladen ist
    outerFrame.MustWaitLoad()

    // Finden Sie das zweite iFrame-Element innerhalb des ersten iFrames
    innerIframeElement := outerFrame.MustElement("#outerdocbody > iframe:nth-child(1)") // Passen Sie den Selektor an

    // Wechseln Sie zum Kontext des zweiten iFrames
    innerFrame := innerIframeElement.MustFrame()
    innerFrame.MustWaitLoad()

    // Extrahieren Sie den Text aus dem zweiten iFrame (Benutze statt #innderdocbody auch body für jedes Element)
    text := innerFrame.MustElement("#innerdocbody").MustText()

    // Regex zum Finden von Links, die mit https beginnen
    re := regexp.MustCompile(`https://[\w\.-]+(?:/[\w\.-]*)*`)
    httpsLinks := re.FindAllString(text, -1)

    fmt.Println("Gefundene https-Links:")
    for _, link := range httpsLinks {
        fmt.Println(link)
    }
}
