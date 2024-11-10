package lib

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type cssFile struct {
	order int
	content string
}

const DefaultOrder = 1000

func CombineCSS(partialsDir string, outputFile string) {
	cssFiles, err := readFiles(partialsDir)
	if err != nil {
		log.Fatalf("Error readinf files: %v", err)
	}

	sort.Slice(cssFiles, func(i, j int) bool {
		return cssFiles[i].order < cssFiles[j].order
	})

	err = writeOutput(cssFiles, outputFile)
	if err != nil {
		log.Fatalf("Error writing output: %v", err)
	}
	fmt.Println("Combined CSS files into ", outputFile)
}

func readFiles(partialsDir string) ([]cssFile, error) {
	var cssFiles []cssFile
	err := filepath.Walk(partialsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".css" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			cleanedContent := cleanCSS(string(content))
			order := extractOrder(string(content))
			cssFiles = append(cssFiles, cssFile{order: order, content: cleanedContent})
		}
		return nil
	})
	return cssFiles, err
}

func extractOrder(content string) int {
    scanner := bufio.NewScanner(strings.NewReader(content))
    for scanner.Scan() {
        line := scanner.Text()
        // Change to detect CSS comment "/* Order: X */"
        if strings.Contains(line, "/* Order:") && strings.Contains(line, "*/") {
            parts := strings.Split(line, ":")
            if len(parts) == 2 {
                orderStr := strings.TrimSpace(strings.TrimSuffix(parts[1], "*/"))
                order, err := strconv.Atoi(orderStr)
                if err == nil {
                    return order
                }
                log.Printf("Warning: Invalid order value '%s'", orderStr)
                break
            }
        }
    }
    return DefaultOrder
}

func writeOutput(cssFiles []cssFile, outputFile string) error {
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	for _, cssFile := range cssFiles {
		_, err := io.WriteString(output, cssFile.content + "\n")
		if err != nil {
			return err
		}
	}
	return nil 
}

func cleanCSS(content string) string {
	var result strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(content))

	var buffer string 

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, "/* Order:") { 
			continue
		}

		if strings.HasSuffix(line, ",") {
			buffer += line + " " 
		} else if buffer != "" {
			buffer += line
			result.WriteString(buffer + "\n")
			buffer = ""
		} else {
			result.WriteString(line + "\n")
		}
	}

	// Write any remaining buffer content
	if buffer != "" {
		result.WriteString(buffer + "\n")
	}

	return result.String()
}