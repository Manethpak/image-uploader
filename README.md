<h1 align="center">Image Uploader</h1>

<div align="center">
  A lightweight, and minial image hosting and sharing service built with Go and Svelte. <a href="https://uploader.manethpak.com">Demo here</a>
</div>

<!-- TABLE OF CONTENTS -->

## Table of Contents

- [Overview](#overview)
  - [Built With](#built-with)
- [Features](#features)
- [How to use](#how-to-use)
- [Contact](#contact)
- [Acknowledgements](#acknowledgements)

<!-- OVERVIEW -->

## Overview

![screenshot](https://uploader.manethpak.com/public/9EBu22irK6jQZCB0ZzkYH.png)

A free lightweight, minimal, image hosting and sharing service. Built with Go and Svelte. 

## Built With

### Language

- [Go](https://golang.org/)
- [TypeScript](https://www.typescriptlang.org/)

### Framework & Library

- [Svelte](https://svelte.dev/)
- [Tailwind](https://tailwindcss.com/)
- [gin-gonic](https://gin-gonic.com/)

## Features

- [x] Quick upload and drag-and-drop with no user authentication
- [x] Image sharing
- [x] Supports PNG, JPEG, WebP formats
- [ ] HTTPS-only image serving
- [ ] Automatic image optimization
- [ ] Secure, CDN-backed image hosting

## Run it yourself

To clone and run this application, you'll need [Git](https://git-scm.com) [Node.js](https://nodejs.org/en/download/) (which comes with [npm](http://npmjs.com)) and [Go](https://golang.org) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/Manethpak/image-uploader.git
```

### Running the backend api

```bash
# cd into backend directory
$ cd backend

# install dependencies
$ go mod download

# run the backend
$ go run main.go
```

### Running the frontend

```bash
# cd into frontend directory
$ cd frontend

# install dependencies
$ yarn install

# run the frontend
$ yarn dev
```

### Running everything with Docker

```
docker build -t uploader:latest .

docker run -p 8080:8080 uploader:latest
```

## Acknowledgements

-

## Contact

<!-- - Website [your-website.com](https://{your-web-site-link}) -->

- GitHub [@Manethpak](https://github.com/Manethpak)
<!-- - Twitter [@your-twitter](https://{twitter.com/your-username}) -->
