package main

import (
    "fmt"
    "regexp"
    "sync"
    "time"
    "github.com/go-rod/rod"
)

var visitedURLs = sync.Map{}
var allLinks = sync.Map{}
var mutex = &sync.Mutex{}
var wg sync.WaitGroup

func main() {
    fmt.Println("Starte das Programm...")
    startTime := time.Now()

    browser := rod.New().MustConnect()
    defer browser.MustClose()

    initialURL := "https://pad.stratum0.org/p/dc"
    fmt.Printf("Beginne mit der initialen URL: %s\n", initialURL)

    wg.Add(1)
    go extractLinksFromPage(browser, initialURL)

    wg.Wait()

    fmt.Println("\nAlle gefundenen Links:")
    linkCount := 0
    allLinks.Range(func(key, value interface{}) bool {
        fmt.Println(key)
        linkCount++
        return true
    })

    fmt.Printf("\nStatistik:\n")
    fmt.Printf("Gesamtanzahl der gefundenen Links: %d\n", linkCount)
    visitedCount := 0
    visitedURLs.Range(func(key, value interface{}) bool {
        visitedCount++
        return true
    })
    fmt.Printf("Anzahl der besuchten URLs: %d\n", visitedCount)
    fmt.Printf("Gesamtzeit: %v\n", time.Since(startTime))
}

func extractLinksFromPage(browser *rod.Browser, url string) {
    defer wg.Done()

    mutex.Lock()
    fmt.Printf("\nVerarbeite URL: %s\n", url)
    mutex.Unlock()

    if _, visited := visitedURLs.LoadOrStore(url, true); visited {
        mutex.Lock()
        fmt.Printf("URL bereits besucht: %s\n", url)
        mutex.Unlock()
        return
    }

    page := browser.MustPage(url)
    defer page.MustClose()
    page.MustWaitLoad()

    mutex.Lock()
    fmt.Printf("Seite geladen: %s\n", url)
    mutex.Unlock()

    // Verarbeite die Hauptseite
    processPage(page, url)

    // Verarbeite die verschachtelten iFrames
    processNestedIframes(page, url)
}

func processNestedIframes(page *rod.Page, sourceURL string) {
    mutex.Lock()
    fmt.Printf("Suche nach äußerem iFrame auf %s\n", sourceURL)
    mutex.Unlock()

    // Finden Sie das erste iFrame-Element
    outerIframeElement := page.MustElement("#editorcontainer > iframe:nth-child(1)")

    // Wechseln Sie zum Kontext des ersten iFrames
    outerFrame := outerIframeElement.MustFrame()

    // Warten Sie, bis der Inhalt des ersten iFrames geladen ist
    outerFrame.MustWaitLoad()

    mutex.Lock()
    fmt.Printf("Äußeres iFrame geladen auf %s\n", sourceURL)
    mutex.Unlock()

    // Verarbeite das äußere iFrame
    processPage(outerFrame, sourceURL+" (äußeres iFrame)")

    mutex.Lock()
    fmt.Printf("Suche nach innerem iFrame auf %s\n", sourceURL)
    mutex.Unlock()

    // Finden Sie das zweite iFrame-Element innerhalb des ersten iFrames
    innerIframeElement := outerFrame.MustElement("#outerdocbody > iframe:nth-child(1)")

    // Wechseln Sie zum Kontext des zweiten iFrames
    innerFrame := innerIframeElement.MustFrame()
    innerFrame.MustWaitLoad()

    mutex.Lock()
    fmt.Printf("Inneres iFrame geladen auf %s\n", sourceURL)
    mutex.Unlock()

    // Verarbeite das innere iFrame
    processPage(innerFrame, sourceURL+" (inneres iFrame)")
}

func processPage(page *rod.Page, sourceURL string) {
    text := page.MustElement("body").MustText()

    re := regexp.MustCompile(`https://pad\.stratum0\.org/p/dc[^\s"']+`)
    links := re.FindAllString(text, -1)

    mutex.Lock()
    fmt.Printf("Gefundene Links auf %s: %d\n", sourceURL, len(links))
    mutex.Unlock()

    for _, link := range links {
        if _, exists := allLinks.LoadOrStore(link, true); !exists {
            mutex.Lock()
            fmt.Printf("Neuer Link gefunden: %s\n", link)
            mutex.Unlock()
            wg.Add(1)
            go extractLinksFromPage(page.Browser(), link)
        }
    }
}
