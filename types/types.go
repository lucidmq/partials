package types

import (
	"fmt"
	"html/template"
	"strings"
	"time"
)

type BaseTemplatePageData struct {
	TemplateName    string
	Partials        []string
	CssFiles        []string
	ProdMode        bool
	EnableHighlight bool
	MetaTags        template.HTML
	TemplateData    map[string]interface{}
}

type BlogPageDetails struct {
	Title      string
	Author     string
	Authorrole string
	Authorurl  string
	Dateposted time.Time
	Snippet    string
	Slug       string
	Mainimgurl string
	Tag        string
}

func (bpd BlogPageDetails) GetDateString() string {
	formatted := fmt.Sprintf("%d-%02d-%02d", bpd.Dateposted.Year(), bpd.Dateposted.Month(), bpd.Dateposted.Day())
	return formatted
}

func (bpd BlogPageDetails) GetMonthDateString() string {
	formatted := fmt.Sprintf("%s %02d, %d", bpd.Dateposted.Month().String()[:3], bpd.Dateposted.Day(), bpd.Dateposted.Year())
	return formatted
}

func (bpd BlogPageDetails) GetOgMetaTag(siteName string) string {
	websiteLink := "https://%s%s"
	pageUrl := fmt.Sprintf(websiteLink, siteName, bpd.Slug)
	imgURL := fmt.Sprintf(websiteLink, siteName, bpd.Mainimgurl)
	ogTagBase := "<meta property=\"%s\" content=\"%s\" />"
	twitterTagBase := "<meta name=\"%s\" content=\"%s\" />"
	var og_tags_sb strings.Builder
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:type", "article"))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:locale", "en_US"))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:title", bpd.Title))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:description", bpd.Snippet))
	og_tags_sb.WriteString("\n")
	// Fix this image size...
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:image", imgURL))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:image:width", "500"))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:image:height", "500"))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:site_name", "Integrand"))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(ogTagBase, "og:url", pageUrl))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(twitterTagBase, "twitter:card", "summary_large_image"))
	og_tags_sb.WriteString("\n")
	og_tags_sb.WriteString(fmt.Sprintf(twitterTagBase, "twitter:image", imgURL))
	// <meta property="article:published_time" content="2019-01-08T06:00:59-05:00" />
	// <meta property="article:modified_time" content="2024-06-18T15:57:17-05:00" />
	// For facebook...
	// <meta property="article:publisher" content="https://www.facebook.com/Backlinko"/>
	return og_tags_sb.String()
}
