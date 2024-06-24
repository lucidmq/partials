# Partials

This repository contains reuseable components that are used across all the landing pages. It contains data structures, css and go templates to render websites.

## How to install

1. Add submodule to existing project
```
git submodule add git@github.com:integrandio/partials.git
```


2. Updat go mod file to reference this submodule.
```go.mod
require partials v1.0.0

replace partials v1.0.0 => ./partials
```