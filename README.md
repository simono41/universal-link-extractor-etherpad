# Universal Link Extractor for Etherpad Lite

## Beschreibung

Der Universal Link Extractor for Etherpad Lite ist ein leistungsfähiges Go-Programm, das entwickelt wurde, um Links von Etherpad Lite Instanzen zu extrahieren. Es navigiert durch verschachtelte iFrames und sammelt alle URLs, die einem bestimmten Muster entsprechen. Dieses Tool ist besonders nützlich für die Analyse und Kartierung von Etherpad-basierten Kollaborationsplattformen.

## Hauptfunktionen

- Extrahiert Links von einer oder mehreren initialen URLs
- Navigiert rekursiv durch gefundene Links bis zu einer konfigurierbaren maximalen Tiefe
- Durchsucht verschachtelte iFrames, um versteckte Links zu finden
- Vermeidet doppelte Besuche von URLs
- Bietet eine Fortschrittsanzeige für den Extraktionsprozess
- Sammelt und zeigt Statistiken über gefundene und besuchte Links

## Voraussetzungen

- Go 1.16 oder höher
- go-rod Bibliothek
- progressbar/v3 Bibliothek

## Installation

1. Stellen Sie sicher, dass Go auf Ihrem System installiert ist.

2. Klonen Sie das Repository:
   ```
   git clone https://github.com/yourusername/universal-link-extractor-etherpad.git
   cd universal-link-extractor-etherpad
   ```

3. Installieren Sie die erforderlichen Abhängigkeiten:
   ```
   go mod tidy
   ```

## Konfiguration

Passen Sie die folgenden Variablen am Anfang der `main.go` Datei an Ihre Bedürfnisse an:

```go
maxDepth          = 3 // Maximale Rekursionstiefe
initialURLs       = []string{"https://pad.stratum0.org/p/dc"} // Startseiten
linkRegexPattern  = `https://pad\.stratum0\.org/p/dc[^\s"']+` // Regex für zu extrahierende Links
```

## Verwendung

1. Führen Sie das Programm aus:
   ```
   go run main.go
   ```

2. Das Programm wird mit der Extraktion von Links beginnen und den Fortschritt in der Konsole anzeigen.

3. Nach Abschluss wird eine Liste aller gefundenen Links sowie Statistiken ausgegeben.

## Ausgabe

Das Programm liefert folgende Informationen:

- Fortschrittsmeldungen für jede verarbeitete URL
- Informationen über gefundene und geladene iFrames
- Eine Liste aller gefundenen einzigartigen Links
- Statistiken über die Gesamtanzahl der gefundenen Links, besuchten URLs und die Gesamtlaufzeit

## Anpassung

- Um andere Etherpad Lite Instanzen zu durchsuchen, passen Sie die `initialURLs` und `linkRegexPattern` an.
- Für komplexere Extraktionslogik können Sie die `extractLinks` Funktion modifizieren.

## Fehlerbehebung

- Bei Problemen mit dem Laden von iFrames überprüfen Sie die CSS-Selektoren in der `processNestedIframes` Funktion.
- Wenn Sie Timeout-Probleme haben, passen Sie die `MustWaitLoad` Aufrufe an oder fügen Sie zusätzliche Wartezeiten ein.

## Beitrag

Beiträge zum Projekt sind willkommen. Bitte öffnen Sie ein Issue oder einen Pull Request für Vorschläge oder Verbesserungen.
