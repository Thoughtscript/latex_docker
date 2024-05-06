# latex

- [x] Setup the necessary Latex dependencies
- [x] Mount files
- [x] Generate a PDF
- [x] Host the PDF through a static file server
- [ ] Support for uploading text through the browser
- [x] Support for regenerating the pdf within the container

```bash
docker-compose up
```

View the PDF:

1. https://localhost/public/ - a view, into the container (can also get the PDF from **Docker Desktop** > **Container Files**).
2. https://localhost/api/makepdf - POST endpoint that'll regenerate the PDF file without having to restart the container. Is called from the view above.

## Resources and Links

1. https://wch.github.io/latexsheet/latexsheet-0.png
1. https://linuxconfig.org/how-to-install-latex-on-ubuntu-20-04-focal-fossa-linux
1. https://www.latex-project.org/help/documentation/usrguide-historic.pdf
1. https://askubuntu.com/questions/1034019/18-04-cannot-install-texlive-full
2. https://tutorialedge.net/golang/go-file-upload-tutorial/