package partials

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"partials/types"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	HTMl_FILE_EXTENSION = ".html"
	CSS_FILE_EXTENSION  = ".css"
	JS_FILE_EXTENSION   = ".js"
	PARTIALS_DIRECTORY  = "partials/components/"
	STATIC_BASE_PATH    = "web/static"
	STYLES_WEB_PATH     = "/static/styles/"
	JS_WEB_PATH         = "/static/js/"
)

var TemplatesMap = map[string]*template.Template{}

func BaseTemplateGenerator(templater *types.BaseTemplatePageData, starting_template string) (*template.Template, error) {
	SYLES_OUTPUT_DIRECTORY := STATIC_BASE_PATH + "/styles/"
	JS_OUTPUT_DIRECOTRY := STATIC_BASE_PATH + "/js/"

	templateName := templater.TemplateName
	templates_to_bring := templater.Partials

	cssFiles := []string{}
	jsFiles := []string{}

	templates_to_fetch := append([]string{}, templates_to_bring...)
	all_templates := []string{}

	val, ok := TemplatesMap[templateName]
	// If the key exists
	if ok {
		for _, template := range templates_to_fetch {
			fileName := filepath.Base(template)
			baseCssPath := SYLES_OUTPUT_DIRECTORY + fileName + CSS_FILE_EXTENSION
			if checkFileExists(baseCssPath) {
				css_web_path := STYLES_WEB_PATH + fileName + CSS_FILE_EXTENSION
				cssFiles = append(cssFiles, css_web_path)
			}
			baseJsPath := JS_OUTPUT_DIRECOTRY + fileName + JS_FILE_EXTENSION
			if checkFileExists(baseJsPath) {
				js_web_path := JS_WEB_PATH + fileName + JS_FILE_EXTENSION
				jsFiles = append(jsFiles, js_web_path)
			}

		}
		templater.CssFiles = cssFiles
		templater.JsFiles = jsFiles
		return val, nil
	}
	log.Printf("Generating template %s \n", templateName)
	var template_builder_string strings.Builder
	for _, template := range templates_to_fetch {
		baseFileName := filepath.Base(template)

		template_locations := PARTIALS_DIRECTORY + template + "/" + baseFileName + HTMl_FILE_EXTENSION
		all_templates = append(all_templates, template_locations)

		baseCssPath := SYLES_OUTPUT_DIRECTORY + baseFileName + CSS_FILE_EXTENSION
		if checkFileExists(baseCssPath) {
			css_file_path := STYLES_WEB_PATH + baseFileName + CSS_FILE_EXTENSION
			cssFiles = append(cssFiles, css_file_path)
		}

		baseJsPath := JS_OUTPUT_DIRECOTRY + baseFileName + JS_FILE_EXTENSION
		if checkFileExists(baseJsPath) {
			js_file_path := JS_WEB_PATH + baseFileName + JS_FILE_EXTENSION
			jsFiles = append(jsFiles, js_file_path)
		}

		full_template_name := baseFileName + HTMl_FILE_EXTENSION

		formatted_template_string := fmt.Sprintf("{{ template \"%s\" . }}\n", full_template_name)
		template_builder_string.WriteString(formatted_template_string)
	}

	// Set our css and our js files
	templater.CssFiles = cssFiles
	templater.JsFiles = jsFiles

	// Start Templating our inside template
	fileContents, err := os.ReadFile(starting_template)
	if err != nil {
		return nil, err
	}
	templateString := strings.Replace(string(fileContents), "#replace_me#", template_builder_string.String(), 1)

	funcMap := GetFuncMap()

	tmpl, err := template.New(templateName).Funcs(funcMap).Parse(templateString)
	if err != nil {
		return nil, err
	}
	_, err = tmpl.ParseFiles(all_templates...)
	if err != nil {
		return nil, err
	}

	TemplatesMap[templateName] = tmpl
	return tmpl, nil
}

func CompileAllAssets() {
	entries, err := os.ReadDir(PARTIALS_DIRECTORY)
	if err != nil {
		log.Fatal(io.ErrNoProgress)
	}
	directoryIterator(PARTIALS_DIRECTORY, entries)
}

func directoryIterator(startingPath string, items []fs.DirEntry) {
	for _, item := range items {
		completePath := filepath.Join(startingPath, item.Name())
		if item.IsDir() {
			subitems, err := os.ReadDir(completePath)
			if err != nil {
				log.Fatal(err)
			}
			directoryIterator(completePath, subitems)
		} else {
			assetCompiler(completePath, item)
		}
	}
}

func assetCompiler(completePath string, item fs.DirEntry) {
	SYLES_OUTPUT_DIRECTORY := STATIC_BASE_PATH + "/styles/"
	JS_OUTPUT_DIRECOTRY := STATIC_BASE_PATH + "/js/"

	switch extension := path.Ext(item.Name()); extension {
	case CSS_FILE_EXTENSION:
		output_css_file := SYLES_OUTPUT_DIRECTORY + item.Name()
		err := symLinker(completePath, output_css_file)
		if err != nil {
			log.Fatalln(err)
		}
	case JS_FILE_EXTENSION:
		output_js_file := JS_OUTPUT_DIRECOTRY + item.Name()
		err := symLinker(completePath, output_js_file)
		if err != nil {
			log.Fatalln(err)
		}
	case HTMl_FILE_EXTENSION:
	default:
	}
}

func checkFileExists(src string) bool {
	fullSrcPath, err := filepath.Abs(src)
	if err != nil {
		log.Println(err)
		return false
	}
	// Check if our src file exists, if not then we just want to exit
	if _, err := os.Stat(fullSrcPath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func symLinker(src, dst string) error {
	fullSrcPath, err := filepath.Abs(src)
	if err != nil {
		return err
	}
	fullDstPath, err := filepath.Abs(dst)
	if err != nil {
		return err
	}
	// Remove the symlink if it already exists
	if _, err := os.Lstat(fullDstPath); err == nil {
		err = os.Remove(fullDstPath)
		if err != nil {
			return err
		}
	}
	return os.Symlink(fullSrcPath, fullDstPath)
}

var comment_re = regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")

func CssCopy(src, dst string) (int, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	// Read the file into a byte slice
	cssBytes := make([]byte, sourceFileStat.Size())
	_, err = bufio.NewReader(source).Read(cssBytes)
	if err != nil {
		return 0, err
	}

	strippedCSSBytes := comment_re.ReplaceAll(cssBytes, nil)

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := destination.Write(strippedCSSBytes)
	return nBytes, err
}

// TODO: test this and see if it works...
// func CompileGivenAssets(templateNames []string) {
// 	var directories_to_build []string
// 	entries, err := os.ReadDir(PARTIALS_DIRECTORY)
// 	if err != nil {
// 		log.Fatal(io.ErrNoProgress)
// 	}
// 	for _, entry := range entries {
// 		if entry.IsDir() {
// 			for _, templateName := range templateNames {
// 				if entry.Name() == templateName {
// 					directories_to_build = append(directories_to_build, entry.Name())
// 				}
// 			}
// 		} else {
// 			log.Println("Entry is not a directory", entry.Name())
// 		}
// 	}

// 	assetCompiler(directories_to_build)
// }
