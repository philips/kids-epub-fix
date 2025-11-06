package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "regexp"
    "strings"
)

const layoutCSS = `body, .page-background-image, html {
    height: 100% !important; 
    width: 100% !important; 
    margin: 0 !important;
    padding: 0 !important;
    position: static !important;
    z-index: auto !important;
    overflow: hidden !important; 
}

img {
    width: auto !important;
    height: 100% !important;
    max-width: 100% !important;
    max-height: 100% !important;
    object-fit: contain !important;
    display: block !important;
    margin: 0 !important;
    position: static !important;
}`

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: kid-epub-fix <epub-file>")
        os.Exit(1)
    }
    
    inputFile := os.Args[1]
    outputFile := strings.TrimSuffix(inputFile, filepath.Ext(inputFile)) + "-fix.epub"

    // Open input EPUB
    inZip, err := zip.OpenReader(inputFile)
    if err != nil {
        fmt.Printf("Error opening EPUB: %v\n", err)
        os.Exit(1)
    }
    defer inZip.Close()

    // Create output EPUB
    outFile, err := os.Create(outputFile)
    if err != nil {
        fmt.Printf("Error creating output: %v\n", err)
        os.Exit(1)
    }
    defer outFile.Close()
    outZip := zip.NewWriter(outFile)
    defer outZip.Close()

    // Add our CSS to standard location
    cssWriter, _ := outZip.Create("OEBPS/Styles/layout.css")
    cssWriter.Write([]byte(layoutCSS))

    imgRegex := regexp.MustCompile(`<img[^>]*src="([^"]+)"`)

    // Process each file in EPUB
    for _, file := range inZip.File {
        switch {
        case strings.HasSuffix(file.Name, ".css"):
            // Skip existing CSS files
            continue
            
        case strings.HasSuffix(file.Name, ".xhtml"):
            // Process XHTML content
            f, _ := file.Open()
            content, _ := io.ReadAll(f)
            f.Close()

            // Extract first image source
            matches := imgRegex.FindStringSubmatch(string(content))
            imgSrc := ""
            if len(matches) > 1 {
                imgSrc = matches[1]
            }

            // Create simplified XHTML
            newContent := fmt.Sprintf(`<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <link rel="stylesheet" type="text/css" href="../Styles/layout.css"/>
</head>
<body>
    <img src="%s"/>
</body>
</html>`, imgSrc)

            writer, _ := outZip.Create(file.Name)
            writer.Write([]byte(newContent))
            
        default:
            // Copy other files directly
            f, _ := file.Open()
            defer f.Close()
            
            writer, _ := outZip.Create(file.Name)
            io.Copy(writer, f)
        }
    }

    fmt.Printf("Created fixed EPUB: %s\n", outputFile)
}
