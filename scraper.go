package main

import (
    "fmt"
    "regexp"
    "time"
    "github.com/go-rod/rod"
)

var (
    visitedURLs = make(map[string]bool)
    allLinks    = make(map[string]bool)
    maxDepth    = 3 // Hier können Sie die maximale Tiefe festlegen
)

func main() {
    fmt.Println("Starte das Programm...")
    startTime := time.Now()

    browser := rod.New().MustConnect()
    defer browser.MustClose()

    initialURL := "https://pad.stratum0.org/p/dc"
    fmt.Printf("Beginne mit der initialen URL: %s\n", initialURL)
    fmt.Printf("Maximale Tiefe: %d\n", maxDepth)

    toVisit := []struct {
        url   string
        depth int
    }{{initialURL, 0}}

    for len(toVisit) > 0 {
        current := toVisit[0]
        toVisit = toVisit[1:]

        newLinks := extractLinksFromPage(browser, current.url, current.depth)

        for _, link := range newLinks {
            if !visitedURLs[link] && current.depth < maxDepth {
                toVisit = append(toVisit, struct {
                    url   string
                    depth int
                }{link, current.depth + 1})
            }
        }
    }

    fmt.Println("\nAlle gefundenen Links:")
    for link := range allLinks {
        fmt.Println(link)
    }

    fmt.Printf("\nStatistik:\n")
    fmt.Printf("Gesamtanzahl der gefundenen Links: %d\n", len(allLinks))
    fmt.Printf("Anzahl der besuchten URLs: %d\n", len(visitedURLs))
    fmt.Printf("Gesamtzeit: %v\n", time.Since(startTime))
}

func extractLinksFromPage(browser *rod.Browser, url string, depth int) []string {
    fmt.Printf("\nVerarbeite URL: %s (Tiefe: %d)\n", url, depth)

    if depth > maxDepth {
        fmt.Printf("Maximale Tiefe erreicht für URL: %s\n", url)
        return nil
    }

    if visitedURLs[url] {
        fmt.Printf("URL bereits besucht: %s\n", url)
        return nil
    }

    visitedURLs[url] = true

    page := browser.MustPage(url)
    defer page.MustClose()
    page.MustWaitLoad()

    fmt.Printf("Seite geladen: %s\n", url)

    var newLinks []string

    // Verarbeite die Hauptseite
    newLinks = append(newLinks, extractLinks(page, url)...)

    // Verarbeite die verschachtelten iFrames
    newLinks = append(newLinks, processNestedIframes(page, url)...)

    return newLinks
}

func processNestedIframes(page *rod.Page, sourceURL string) []string {
    fmt.Printf("Suche nach äußerem iFrame auf %s\n", sourceURL)

    // Finden Sie das erste iFrame-Element
    outerIframeElement := page.MustElement("#editorcontainer > iframe:nth-child(1)")

    // Wechseln Sie zum Kontext des ersten iFrames
    outerFrame := outerIframeElement.MustFrame()

    // Warten Sie, bis der Inhalt des ersten iFrames geladen ist
    outerFrame.MustWaitLoad()

    fmt.Printf("Äußeres iFrame geladen auf %s\n", sourceURL)

    // Extrahiere Links aus dem äußeren iFrame
    outerLinks := extractLinks(outerFrame, sourceURL+" (äußeres iFrame)")

    fmt.Printf("Suche nach innerem iFrame auf %s\n", sourceURL)

    // Finden Sie das zweite iFrame-Element innerhalb des ersten iFrames
    innerIframeElement := outerFrame.MustElement("#outerdocbody > iframe:nth-child(1)")

    // Wechseln Sie zum Kontext des zweiten iFrames
    innerFrame := innerIframeElement.MustFrame()
    innerFrame.MustWaitLoad()

    fmt.Printf("Inneres iFrame geladen auf %s\n", sourceURL)

    // Extrahiere Links aus dem inneren iFrame
    innerLinks := extractLinks(innerFrame, sourceURL+" (inneres iFrame)")

    return append(outerLinks, innerLinks...)
}

func extractLinks(page *rod.Page, sourceURL string) []string {
    text := page.MustElement("body").MustText()

    re := regexp.MustCompile(`https://pad\.stratum0\.org/p/dc[^\s"']+`)
    links := re.FindAllString(text, -1)

    fmt.Printf("Gefundene Links auf %s: %d\n", sourceURL, len(links))

    var newLinks []string
    for _, link := range links {
        if !allLinks[link] {
            allLinks[link] = true
            fmt.Printf("Neuer Link gefunden: %s\n", link)
            newLinks = append(newLinks, link)
        }
    }

    return newLinks
}
