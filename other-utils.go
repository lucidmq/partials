package partials

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"partials/types"
	"strconv"
	"strings"
)

func GetEnvrionmentVariableString(env string, substitueVariable string) string {
	value := os.Getenv(env)
	if value == "" {
		os.Setenv(env, substitueVariable)
		return substitueVariable
	} else {
		return value
	}
}

func GetEnvrionmentVariableInt(env string, substitueValue int) int {
	subStringValue := strconv.Itoa(substitueValue)
	strValue := os.Getenv(env)
	if strValue != "" {
		value, err := strconv.Atoi(strValue)
		if err == nil {
			return value
		}
	}
	os.Setenv(env, subStringValue)
	return substitueValue
}

func GetBaseHeaderTags(title string, description string) template.HTML {
	baseTitleTag := "<title>%s</title>"
	descriptionTag := "<meta name=\"description\" content=\"%s\">"
	var meta_sb strings.Builder
	meta_sb.WriteString(fmt.Sprintf(baseTitleTag, title))
	meta_sb.WriteString("\n")
	meta_sb.WriteString(fmt.Sprintf(descriptionTag, description))
	meta_sb.WriteString("\n")
	return template.HTML(meta_sb.String())
}

func GetOgMetaTag(siteName string, slug string, imgUrl string, title string, description string) template.HTML {
	websiteLink := "https://%s%s"
	pageUrl := fmt.Sprintf(websiteLink, siteName, slug)
	imgURL := fmt.Sprintf(websiteLink, siteName, imgUrl)
	ogTagBase := "<meta property=\"%s\" content=\"%s\" />"
	twitterTagBase := "<meta name=\"%s\" content=\"%s\" />"
	var og_tags_sb strings.Builder
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:locale", "en_US"))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:title", title))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:description", description))
	og_tags_sb.WriteString("\n")
	// Fix this image size...
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:image", imgURL))
	og_tags_sb.WriteString("\n")
	// og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:image:width", "500"))
	// og_tags_sb.WriteString("\n")
	// og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:image:height", "500"))
	// og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:site_name", siteName))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:url", pageUrl))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(twitterTagBase, "twitter:card", "summary_large_image"))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(twitterTagBase, "twitter:image", imgURL))
	return template.HTML(og_tags_sb.String())
}

var testHTMLTemplate = `
    <div class="FVRLGYXIOVVZXPN">
	  {{ range $index, $element := .}}
      	<pre><code class="SASNLXGQKQEAZDE">{{$element}}</code></pre>
	  {{end}}
	</div>
`

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Used to generate code demo that's used in the landing page
func HtmlCodeFormatter(filePath string) string {
	t, err := template.New("foo").Parse(testHTMLTemplate)
	if err != nil {
		log.Fatal(err)
	}
	lines, err := readLines(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	err = t.ExecuteTemplate(&out, "foo", lines)
	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}

func CombineFilesWithSpace(files []string, outputFile string) error {
	log.Println("Combining CSS file")
	// Create or overwrite the output file
	out, err := os.Create(outputFile)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer out.Close()

	for i, file := range files {
		if strings.Contains(file, "topbar") {
			continue
		} else if strings.Contains(file, "footer") {
			continue
		}
		file = "web" + file
		// Open each input file
		in, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %v", file, err)
		}

		// Copy the contents of the file to the output file
		_, err = io.Copy(out, in)
		in.Close()
		if err != nil {
			return fmt.Errorf("failed to copy contents of file %s: %v", file, err)
		}

		// Add a space after each file's contents except the last one
		if i < len(files)-1 {
			_, err = out.WriteString(" ")
			if err != nil {
				return fmt.Errorf("failed to write space to output file: %v", err)
			}
		}
	}

	return nil
}

func GenerateHtmlPage(pageData types.BaseTemplatePageData, htmlOutputFile string, cssOutputFile string) {
	tmpl, err := BaseTemplateGenerator(&pageData, "partials/basic_template.html")
	if err != nil {
		log.Fatal(err)
	}
	//Write to file
	err = CombineFilesWithSpace(pageData.CssFiles, cssOutputFile)
	if err != nil {
		log.Fatal(err)
	}
	file, _ := os.Create(htmlOutputFile)
	defer file.Close()
	tmpl.Execute(file, pageData)
}

// func componentCleaner() {
// 	fmt.Println("starting thing")
// 	componentsDir := "components"
// 	items, err := os.ReadDir(componentsDir)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	directoryIterator(componentsDir, items)
// }

// func handlefile(completePath string, item fs.DirEntry) {
// 	switch extension := path.Ext(item.Name()); extension {
// 	case CSS_FILE_EXTENSION:
// 		// handle file there
// 		fmt.Println(completePath)
// 		_, err := CssCopy(completePath, completePath)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 	}
// }
