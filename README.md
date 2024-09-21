# iFrame Link Extractor

## Beschreibung

Dieses Go-Programm verwendet die go-rod Bibliothek, um den Inhalt von verschachtelten iFrames auf einer Webseite zu extrahieren und alle darin enthaltenen HTTPS-Links zu identifizieren und auszugeben. Es ist besonders nützlich für Webseiten mit komplexen iFrame-Strukturen, bei denen der Zugriff auf den Inhalt tiefer verschachtelter iFrames erforderlich ist.

## Funktionen

- Navigiert zu einer angegebenen URL
- Wartet auf das Laden der Hauptseite
- Findet und wechselt in den Kontext eines äußeren iFrames
- Findet und wechselt in den Kontext eines inneren iFrames innerhalb des äußeren iFrames
- Extrahiert den Textinhalt des inneren iFrames
- Identifiziert alle HTTPS-Links im extrahierten Text
- Gibt die gefundenen Links aus

## Voraussetzungen

- Go 1.16 oder höher
- go-rod Bibliothek

## Installation

1. Stellen Sie sicher, dass Go auf Ihrem System installiert ist.

2. Klonen Sie das Repository:
   ```
   git clone https://github.com/yourusername/iframe-link-extractor.git
   cd iframe-link-extractor
   ```

3. Installieren Sie die erforderlichen Abhängigkeiten:
   ```
   go mod tidy
   ```

## Verwendung

1. Öffnen Sie die `main.go` Datei und passen Sie die URL und iFrame-Selektoren an Ihre Zielwebseite an.

2. Führen Sie das Programm aus:
   ```
   go run main.go
   ```

3. Das Programm wird die gefundenen HTTPS-Links in der Konsole ausgeben.

## Anpassung

- Um auf spezifische Elemente innerhalb der iFrames zu warten, können Sie zusätzliche `MustElement().MustWaitVisible()` Aufrufe hinzufügen.
- Der reguläre Ausdruck für die Link-Erkennung kann angepasst werden, um spezifischere URL-Muster zu erfassen.

## Fehlerbehebung

- Wenn Sie Probleme mit Timeouts haben, erhöhen Sie die Wartezeiten mit `page.Timeout()` oder `MustElement().WaitVisible()`.
- Bei Problemen mit dem Zugriff auf iFrames aufgrund von Sicherheitseinstellungen, erwägen Sie die Verwendung zusätzlicher Browser-Optionen wie im Kommentar im Code beschrieben.

## Beitrag

Beiträge zum Projekt sind willkommen. Bitte öffnen Sie ein Issue oder einen Pull Request für Vorschläge oder Verbesserungen.

## Lizenz

[Fügen Sie hier Ihre gewählte Lizenz ein, z.B. MIT, GPL, etc.]
```

Diese README.md bietet eine umfassende Übersicht über Ihr Projekt und enthält Abschnitte für:

1. Eine kurze Beschreibung des Projekts
2. Die Hauptfunktionen
3. Voraussetzungen für die Nutzung
4. Installationsanweisungen
5. Verwendungshinweise
6. Anpassungsmöglichkeiten
7. Tipps zur Fehlerbehebung
8. Informationen zum Beitragen zum Projekt
9. Einen Platzhalter für die Lizenz
