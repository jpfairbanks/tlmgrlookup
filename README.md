# tlmgrlookup

A cli tool for finding texlive packages from filenames

## Motivating example

Have you ever seen a latex error complaining about some style file missing?

Suppose you compiled this latex document:
```latex
\documentclass[ignorenonframetext]{article}
\usepackage{xintfrac}

\begin{document}

Hello evince!

\end{document}
``` 

When you compile this document you will get this error: `ERROR: LaTeX Error: File `xintfrac.sty' not found.` Suppose you
wanted to install `xintfrac`, you would run `tlmgr install xinitfrac`, but then you would get the following totally
frustrating error message:

```sh
$ sudo tlmgr install xintfrac
tlmgr: package repository http://ftp.math.purdue.edu/mirrors/ctan.org/systems/texlive/tlnet (verified)
tlmgr install: package xintfrac not present in repository.
tlmgr: action install returned an error; continuing.
tlmgr: An error has occurred. See above messages. Exiting.```

Thanks `tlmgr`, but what is the package that provides `xintfrac`?

tlmgrlookup to the rescue!

```sh
 $ tlmgrlookup xintfrac
Looking up latex package: xintfrac
Result: 337	tex/generic/xint/xintfrac.sty	xint
To install this package run:

	 sudo tlmgr install xint
```

Now we know exactly which package to install and the tlmgr command to install it.


## Installation

```sh
go get github.com/jpfairbanks/tlmgrlookup
cd $GOPATH/github.com/jpfairbanks/tlmgrlookup
go install .
```

## Usage
To use this package you just provide
`$ tlmgrlookup filename.sty`

Example:

```bash
$ tlmgrlookup readprov.sty
Looking up latex package: readprov.sty
Result: 1892	tex/latex/fileinfo/readprov.sty	fileinfo
To install this package run:

	 sudo tlmgr install fileinfo$ ./tlmgrlookup readprov.sty

```

# LICENSE

This software is licensed under the AGPL.
