# LaTeX Docker

[![](https://img.shields.io/badge/Go-1.22.2-blue.svg)](https://golang.org/pkg/)
[![](https://img.shields.io/badge/LaTeX-TexLive-turquoise.svg)](https://www.latex-project.org/get/)
[![](https://img.shields.io/badge/LATEX.js-0.12.6-red.svg)](https://latex.js.org/)

- [x] Setup the necessary LaTeX dependencies
- [x] Mount files
- [x] Generate a PDF
- [x] Host the PDF through a static file server
- [x] Support for uploading text through the browser
- [x] Support for regenerating the pdf within the container
- [x] Support for writing LaTeX in the browser (and rendering it)
- [x] Support for saving LaTeX in the browser to `paper.tex`
- [ ] Better styling
- [ ] Better editing
- [ ] Better navigation
- [x] Tweak some JS and HTML functionality around redirects, rerendering, and prevent default

```bash
docker-compose up
```

View the PDF and write LaTeX:

1. https://localhost/public/ - a view, into the container (can also get the PDF from **Docker Desktop** > **Container Files**). (**Experimental**)
2. https://localhost/public/editor.html - a compressed view into the container to edit and modify LaTeX. (**Experimental**)
3. https://localhost/public/viewer.html - a compressed view into the container to compile and view pdf's.
4. https://localhost/public/pdfeditor.html - slightly more reliable way to edit and render text live.

## API Endpoints

1. https://localhost/api/pdf/make - POST endpoint that'll regenerate the PDF file without having to restart the container. Is called from the view above.
1. https://localhost/api/pdf/upload - POST endpoint that'll allow a file to be uploaded to overwrite `paper.tex` so it can be recompile (allows raw `.tex` files to uploaded to the container to be compiled into `.pdf`).
1. https://localhost/api/latex/save - POST endpoint to save inputted text to `paper.tex` (better UI and Nav features to be added to provide basic but fuller text editor support).

> Note: all of the commands assume `paper.tex` is the relevant filename and filetype.

## Resources and Links

1. https://wch.github.io/latexsheet/latexsheet-0.png
1. https://linuxconfig.org/how-to-install-latex-on-ubuntu-20-04-focal-fossa-linux
1. https://www.latex-project.org/help/documentation/usrguide-historic.pdf
1. https://askubuntu.com/questions/1034019/18-04-cannot-install-texlive-full
2. https://tutorialedge.net/golang/go-file-upload-tutorial/
3. https://opensource.com/article/18/6/copying-files-go