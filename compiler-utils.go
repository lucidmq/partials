package partials

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"io"
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

	templates_to_fetch := append([]string{"topbar", "footer"}, templates_to_bring...)
	all_templates := []string{}

	val, ok := TemplatesMap[templateName]
	// If the key exists
	if ok {
		log.Printf("Found template %s \n", templateName)
		for _, template := range templates_to_fetch {
			baseCssPath := SYLES_OUTPUT_DIRECTORY + template + CSS_FILE_EXTENSION
			if checkFileExists(baseCssPath) {
				css_web_path := STYLES_WEB_PATH + template + CSS_FILE_EXTENSION
				cssFiles = append(cssFiles, css_web_path)
			}
			baseJsPath := JS_OUTPUT_DIRECOTRY + template + JS_FILE_EXTENSION
			if checkFileExists(baseJsPath) {
				js_web_path := JS_WEB_PATH + template + JS_FILE_EXTENSION
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
		template_locations := PARTIALS_DIRECTORY + template + "/" + template + HTMl_FILE_EXTENSION
		all_templates = append(all_templates, template_locations)
		css_file_path := STYLES_WEB_PATH + template + CSS_FILE_EXTENSION
		js_file_path := JS_WEB_PATH + template + JS_FILE_EXTENSION
		cssFiles = append(cssFiles, css_file_path)
		jsFiles = append(jsFiles, js_file_path)
		if template != "topbar" && template != "footer" {
			formatted_template_string := fmt.Sprintf("{{ template \"%s.html\" . }}\n", template)
			template_builder_string.WriteString(formatted_template_string)
		}
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
	var directories_to_build []string
	entries, err := os.ReadDir(PARTIALS_DIRECTORY)
	if err != nil {
		log.Fatal(io.ErrNoProgress)
	}
	// This iterates through all the entries, let's clean this up so we don't need to do that
	for _, entry := range entries {
		if entry.IsDir() {
			directories_to_build = append(directories_to_build, entry.Name())
		} else {
			log.Println("Entry is not a directory", entry.Name())
		}
	}
	assetCompiler(directories_to_build)
}

// TODO: test this and see if it works...
func CompileGivenAssets(templateNames []string) {
	var directories_to_build []string
	entries, err := os.ReadDir(PARTIALS_DIRECTORY)
	if err != nil {
		log.Fatal(io.ErrNoProgress)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			for _, templateName := range templateNames {
				if entry.Name() == templateName {
					directories_to_build = append(directories_to_build, entry.Name())
				}
			}
		} else {
			log.Println("Entry is not a directory", entry.Name())
		}
	}

	assetCompiler(directories_to_build)
}

func assetCompiler(directories_to_build []string) {
	SYLES_OUTPUT_DIRECTORY := STATIC_BASE_PATH + "/styles/"
	JS_OUTPUT_DIRECOTRY := STATIC_BASE_PATH + "/js/"

	for _, directory := range directories_to_build {
		src_dictory_path := PARTIALS_DIRECTORY + directory
		log.Println("Reading directory " + src_dictory_path)
		files, err := os.ReadDir(src_dictory_path)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			switch extension := path.Ext(file.Name()); extension {
			case CSS_FILE_EXTENSION:
				src_css_file := src_dictory_path + "/" + file.Name()
				output_css_file := SYLES_OUTPUT_DIRECTORY + file.Name()
				err := symLinker(src_css_file, output_css_file)
				if err != nil {
					log.Fatalln(err)
				}
			case JS_FILE_EXTENSION:
				src_js_file := src_dictory_path + "/" + file.Name()
				output_js_file := JS_OUTPUT_DIRECOTRY + file.Name()
				err := symLinker(src_js_file, output_js_file)
				if err != nil {
					log.Fatalln()
				}
			case HTMl_FILE_EXTENSION:
				log.Println(file.Name(), extension)
			default:
				log.Println("Unexpected type ", file.Name())
			}
		}
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

func js_writer(outputFile *os.File, inputFileName string) {
	// open input file
	fi, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}
	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		// write a chunk
		if _, err := outputFile.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	newLineContent := []byte("\n")
	// write a chunk
	if _, err := outputFile.Write(newLineContent); err != nil {
		panic(err)
	}
}
