# Contributing to Pigo

Thanks for your interest in contributing to Pigo.

Please take a moment to review this document before submitting your first pull request. I also strongly recommend that you check for open issues and pull requests to see if someone else is working on something similar.

If you need any help, feel free to reach out.

## Structure

This repository is structured as follows:

```
pigo
├── cmd
├── internal
└── pkg
    ├── icons
    ├── lib
    ├── ui
    └── themes
```

| Path                  | Description                              |
| --------------------- | ---------------------------------------- |
| `pigo/cmd`            | The main development server.             |
| `pigo/internal`       | The internal unexposed utilities.        |
| `pigo/pkg/icons`      | The svg embedded icon templ components.  |
| `pigo/pkg/lib`        | The exposed package library functions.   |
| `pigo/pkg/ui`         | The templ components and matching css.   |
| `pigo/pkg/themes`     | The css theme configurations.            |

## Development

### Fork this repo

You can fork this repo by clicking the fork button in the top right corner of this page.

### Clone on your local machine

```bash
git clone https://github.com/your-username/pigo.git
```

### Navigate to project directory

```bash
cd pigo
```

### Create a new Branch

```bash
git checkout -b my-new-branch
```

### Install dependencies

```bash
make deps
```

#### Running the dev server

1. To run the dev server:

```bash
make dev
```

2. Go to localhost:8000 in your browser

If you wish to change the default port from 8000 to something else:
- open .air.toml
- add the flag `-p` and your desired port number to the `args_bin` array
  - e.g., `args_bin = ["-p", "6969"]`

## Documentation

The documentation for this project and components are located either in the README.md or as Go Doc comments in the code as per [Go Docs](https://tip.golang.org/doc/comment)

## Components

The spirit of this project is derived from [pico](https://picocss.com/), a minimal css framework that utilises semantic HTML.

Components will ideally be developed with this concept in mind, using as little JavaScript as possible. Best attempts should be made so that JavaScript fails gracefully such that the component should be functional without it.

In the event that a component absolutely requires JavaScript, ensure this is documented prominantly in the Go Doc comment. i.e., `// NOTE: component requires JavaScript to function`.

This project also borrows heavily from the concept of Locality of Behaviour [LoB](https://htmx.org/essays/locality-of-behaviour/). To this end, if inlining some css is more helpful than having it in the component's corresponding css file, just do it. Similarly for JavaScript if it is required, inline it, use a script tag, or a templ script component and keep it with the ui component.

## Requests for new components

If you have a request for a new component, please open a discussion on GitHub. I'll be happy to take a look or help you out.
