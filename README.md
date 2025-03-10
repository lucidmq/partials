# Partials

This repository contains reuseable components that are used across all the landing pages. It contains data structures, css and go templates to render websites.

## How to install

1. Add submodule to existing project
```bash
git submodule add git@github.com:lucidmq/partials.git
```


2. Updat go mod file to reference this submodule.
```go.mod
require partials v1.0.0

replace partials v1.0.0 => ./partials
```

## Updating
When new commits are made to the submodule, the submodule hash needs to be updated. To pull down the lastest commits use the following command:
```bash
git submodule update --remote --merge
```