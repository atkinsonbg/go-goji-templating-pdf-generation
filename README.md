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

