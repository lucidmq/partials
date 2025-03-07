# Todo
[ ] Create a way to go from markdown to markdown format -> code lines see `partials/components/codeShowcase/codeShowcase.html`
[ ] Create structure to feed in data to `partials/components/codeShowcase`
[ ] Create title and description to feed in data to `partials/components/codeShowcase`
[ ] Clean up labels for the code containers `partials/components/codeShowcase`

## How to Generate a new Templated Component

```go
    var content types.BaseTemplatePageData
	partials.GenerateHtmlPage(content, "test.html", "test.css")
	return
```