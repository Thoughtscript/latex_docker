FROM ubuntu:24.04

RUN echo "Creating working dir and copying files"
RUN mkdir /app
WORKDIR /app
COPY . .

RUN apt update
RUN apt upgrade
RUN apt autoremove

RUN apt purge libkpathsea6
RUN apt install texlive-latex-recommended -y

CMD ["pdflatex", "hello-world.tex"]