package partials

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
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
