# Pad Stratum0 Link Extractor

## Beschreibung

Dieses Go-Programm ist ein spezialisierter Web Scraper, der entwickelt wurde, um Links von der Pad Stratum0 Plattform zu extrahieren. Es navigiert durch verschachtelte iFrames und sammelt alle URLs, die mit "https://pad.stratum0.org/p/dc" beginnen.

## Funktionen

- Extrahiert Links von einer initialen URL und folgt diesen rekursiv bis zu einer konfigurierbaren maximalen Tiefe.
- Navigiert durch verschachtelte iFrames, um versteckte Links zu finden.
- Vermeidet doppelte Besuche von URLs.
- Bietet detaillierte Konsolenausgaben über den Fortschritt des Scrapings.
- Sammelt Statistiken über gefundene und besuchte Links.

## Voraussetzungen

- Go 1.16 oder höher
- go-rod Bibliothek

## Installation

1. Stellen Sie sicher, dass Go auf Ihrem System installiert ist.

2. Klonen Sie das Repository:
   ```
   git clone https://github.com/yourusername/pad-stratum0-link-extractor.git
   cd pad-stratum0-link-extractor
   ```

3. Installieren Sie die erforderlichen Abhängigkeiten:
   ```
   go mod tidy
   ```

## Konfiguration

Sie können die maximale Suchtiefe anpassen, indem Sie den Wert der `maxDepth` Variable am Anfang der `main.go` Datei ändern:

```go
maxDepth = 3 // Ändern Sie diesen Wert nach Bedarf
```

## Verwendung

1. Führen Sie das Programm aus:
   ```
   go run main.go
   ```

2. Das Programm wird mit der Extraktion von Links beginnen und den Fortschritt in der Konsole ausgeben.

3. Nach Abschluss wird eine Liste aller gefundenen Links sowie Statistiken angezeigt.

## Ausgabe

Das Programm gibt folgende Informationen aus:

- Fortschrittsmeldungen für jede verarbeitete URL
- Informationen über gefundene und geladene iFrames
- Eine Liste aller gefundenen einzigartigen Links
- Statistiken über die Gesamtanzahl der gefundenen Links, besuchten URLs und die Gesamtlaufzeit

## Anpassung

- Um andere Websites zu scrapen, passen Sie die initiale URL und den regulären Ausdruck für die Link-Erkennung an.
- Für komplexere Scraping-Logik können Sie die `extractLinks` Funktion modifizieren.

## Fehlerbehebung

- Wenn Sie Probleme mit dem Laden von iFrames haben, überprüfen Sie die CSS-Selektoren in der `processNestedIframes` Funktion.
- Bei Timeout-Problemen können Sie die `MustWaitLoad` Aufrufe anpassen oder zusätzliche Wartezeiten einbauen.

## Beitrag

Beiträge zum Projekt sind willkommen. Bitte öffnen Sie ein Issue oder einen Pull Request für Vorschläge oder Verbesserungen.

## Lizenz

[Fügen Sie hier Ihre gewählte Lizenz ein, z.B. MIT, GPL, etc.]

```

Diese README.md bietet eine umfassende Übersicht über Ihr Projekt und enthält Abschnitte für:

1. Eine Beschreibung des Projekts und seiner Hauptfunktionen
2. Installationsanweisungen
3. Konfigurationsmöglichkeiten
4. Verwendungshinweise
5. Erklärung der Ausgabe
6. Anpassungsmöglichkeiten
7. Tipps zur Fehlerbehebung
8. Informationen zum Beitragen zum Projekt
9. Einen Platzhalter für die Lizenz
