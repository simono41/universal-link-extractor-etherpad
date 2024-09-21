package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "regexp"
    "time"
    "io"

    "github.com/go-rod/rod"
    "github.com/schollz/progressbar/v3"
)

var (
    visitedURLs = make(map[string]bool)
    allLinks    = make(map[string]bool)
    maxDepth    = 3
    initialURLs = []string{"https://pad.stratum0.org/p/dc"}
    linkRegexPattern = `https://pad\.stratum0\.org/p/dc[^\s"']+`
    debugMode   bool
    logger      *log.Logger
)

func init() {
    flag.BoolVar(&debugMode, "debug", false, "Enable debug mode")
    flag.Parse()

    if debugMode {
        logger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
    } else {
        logger = log.New(io.Discard, "", 0)
    }
}

func main() {
    startTime := time.Now()

    logger.Println("Starte das Programm...")
    browser := rod.New().MustConnect()
    defer browser.MustClose()

    logger.Printf("Initiale URLs: %v\n", initialURLs)
    logger.Printf("Link Regex Pattern: %s\n", linkRegexPattern)
    logger.Printf("Maximale Tiefe: %d\n", maxDepth)

    toVisit := make([]struct {
        url   string
        depth int
    }, len(initialURLs))

    for i, url := range initialURLs {
        toVisit[i] = struct {
            url   string
            depth int
        }{url, 0}
    }

    totalURLs := len(toVisit)
    bar := progressbar.Default(int64(totalURLs))

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
                totalURLs++
                bar.ChangeMax(totalURLs)
            }
        }

        bar.Add(1)
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
    logger.Printf("\nVerarbeite URL: %s (Tiefe: %d)\n", url, depth)

    if depth > maxDepth {
        logger.Printf("Maximale Tiefe erreicht für URL: %s\n", url)
        return nil
    }

    if visitedURLs[url] {
        logger.Printf("URL bereits besucht: %s\n", url)
        return nil
    }

    visitedURLs[url] = true

    page := browser.MustPage(url)
    defer page.MustClose()
    page.MustWaitLoad()

    logger.Printf("Seite geladen: %s\n", url)

    var newLinks []string

    mainLinks := extractLinks(page, url)
    newLinks = append(newLinks, mainLinks...)

    iframeLinks := processNestedIframes(page, url)
    newLinks = append(newLinks, iframeLinks...)

    return newLinks
}

func processNestedIframes(page *rod.Page, sourceURL string) []string {
    logger.Printf("Suche nach äußerem iFrame auf %s\n", sourceURL)

    outerIframeElement := page.MustElement("#editorcontainer > iframe:nth-child(1)")
    outerFrame := outerIframeElement.MustFrame()
    outerFrame.MustWaitLoad()

    logger.Printf("Äußeres iFrame geladen auf %s\n", sourceURL)

    outerLinks := extractLinks(outerFrame, sourceURL+" (äußeres iFrame)")

    logger.Printf("Suche nach innerem iFrame auf %s\n", sourceURL)

    innerIframeElement := outerFrame.MustElement("#outerdocbody > iframe:nth-child(1)")
    innerFrame := innerIframeElement.MustFrame()
    innerFrame.MustWaitLoad()

    logger.Printf("Inneres iFrame geladen auf %s\n", sourceURL)

    innerLinks := extractLinks(innerFrame, sourceURL+" (inneres iFrame)")

    return append(outerLinks, innerLinks...)
}

func extractLinks(page *rod.Page, sourceURL string) []string {
    text := page.MustElement("body").MustText()

    re := regexp.MustCompile(linkRegexPattern)
    links := re.FindAllString(text, -1)

    logger.Printf("Gefundene Links auf %s: %d\n", sourceURL, len(links))

    var newLinks []string
    for _, link := range links {
        if !allLinks[link] {
            allLinks[link] = true
            logger.Printf("Neuer Link gefunden: %s\n", link)
            newLinks = append(newLinks, link)
        }
    }

    return newLinks
}
