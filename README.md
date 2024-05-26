# beecs-ui

[![Test status](https://img.shields.io/github/actions/workflow/status/mlange-42/beecs-ui/tests.yml?branch=main&label=Tests&logo=github)](https://github.com/mlange-42/beecs-ui/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mlange-42/beecs-ui)](https://goreportcard.com/report/github.com/mlange-42/beecs-ui)
[![Go Reference](https://img.shields.io/badge/reference-%23007D9C?logo=go&logoColor=white&labelColor=gray)](https://pkg.go.dev/github.com/mlange-42/beecs-ui)
[![GitHub](https://img.shields.io/badge/github-repo-blue?logo=github)](https://github.com/mlange-42/beecs-ui)

Graphical user interface for the [beecs](https://github.com/mlange-42/beecs) honeybee model and derivatives.

![beecs-ui screenshot](https://github.com/mlange-42/beecs-ui/assets/44003176/a8897dbd-4608-4d74-88da-08e78ee68c1c)

## Features

* Model parametrization from JSON files.
* Customizable parameter manipulation UI.
* Customizable plots and other live visualizations.

For a command line interface for the beecs model, see [beecs-cli](https://github.com/mlange-42/beecs-cli).

## Web version

A web version of the model UI can be used on https://mlange-42.github.io/beecs-ui/.
However, for best performance and configuration options we recommend local use. See below.

## Installation

Pre-compiled binaries for Linux, Windows and MacOS are available in the [Releases](https://github.com/mlange-42/beecs-ui/releases).

> To install the latest **development version** using [Go](https://go.dev), run:
> 
> ```
> go install github.com/mlange-42/beecs-ui@main
> ```

## Usage

Simply run the executable via double click or with

```
beecs-ui
```

To load a parameter file, use the `-p` option:

```
beecs-ui -p path/to/params.json
```

To run with a custom layout, use the `-l` option:

```
beecs-ui -l path/to/layout.json
```
