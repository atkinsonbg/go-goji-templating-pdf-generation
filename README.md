[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/O5O63ENS7)

# Go Goji Templating/PDF Generation
This repo is used to experiment with Go's templating engine and use it to generate PDFs. One of the goals is to get a fully accessible PDF, in the least amount of time and memory/CPU footprint. The API uses HTML templates that are combined with data, then converted to PDFs and optimized.

The tools used are:
* Goji: https://github.com/goji/goji - Goji was selected for its minimal footprint.
* Go text/template package: https://golang.org/pkg/text/template/ - text/template was used over html/template since the final output is a PDF. text/template is faster than html/template. We are not concerned with generating "safe" HTML, since we do not serve this to a browser.
* WKHTMLTOPDF: https://wkhtmltopdf.org/ - This is used to perform the HTML to PDF conversion.
* Ghostscript: https://www.ghostscript.com/ - This is used to optimize and apply PDF metadata. It is also used to apply pdfmarks to support accessibility.

## Accessibility
I tried hard to create a fully accessible PDF, however I'm sure there is a lot of work left to do here. The following online checkers were utilized during the development of this code:
* PAVE: https://pave-pdf.org/?lang=en
* Tingtun Checker: http://checkers.eiii.eu/en/pdfcheck/#page2

## Resources
The following blogs, Stack Overflow posts, etc were used to help solve a lot of problems.

* https://github.com/wkhtmltopdf/wkhtmltopdf/issues/2000
* https://stackoverflow.com/questions/10450120/optimize-pdf-files-with-ghostscript-or-other
* https://www.ghostscript.com/doc/9.23/VectorDevices.htm#PDFWRITE
* https://www.ghostscript.com/doc/current/Use.htm
* https://gist.github.com/shreve/0b73d9dcb7ff11336188
* https://www.lexjansen.com/phuse/2018/ad/AD07_ppt.pdf
* http://silas.net.br/tech/apps/ghostscript.html
* https://stackoverflow.com/questions/52532230/ghostscript-ignores-part-of-pdfmarks
* https://stackoverflow.com/questions/21864382/is-there-any-wkhtmltopdf-option-to-convert-html-text-rather-than-file
* https://stackoverflow.com/questions/63757618/use-ghostscript-to-set-pdf-natural-language-via-pdfmarks/63757751
* http://milan.kupcevic.net/ghostscript-ps-pdf/
* https://stackoverflow.com/questions/6394905/wkhtmltopdf-what-paper-sizes-are-valid

## License
 
The MIT License (MIT)

Copyright (c) 2021 Brandon Atkinson

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
