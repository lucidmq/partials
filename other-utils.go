package partials

import (
	"fmt"
	"html/template"
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
