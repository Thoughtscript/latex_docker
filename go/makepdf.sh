#!/usr/bin/env bash

echo "Generating PDF" &

pdflatex -output-directory=assets -jobname=paper assets/paper.tex &