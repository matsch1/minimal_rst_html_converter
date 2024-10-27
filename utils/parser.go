package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
)

// https://www.sphinx-doc.org/en/master/usage/restructuredtext/roles.html#role-math

func latexToHTML(latex string) string {
	// Replace fractions and powers with HTML
	latex = regexp.MustCompile(`\\frac{(.*?)}{(.*?)}`).ReplaceAllString(
		latex,
		`<span class="frac"><span class="numerator">$1</span><span class="denominator">$2</span></span>`,
	)
	// Convert superscripts
	latex = regexp.MustCompile(`([a-zA-Z0-9]+)\^([a-zA-Z0-9]+)`).ReplaceAllString(latex, `$1<sup>$2</sup>`)

	// Convert subscripts
	latex = regexp.MustCompile(`([a-zA-Z]+)_{(.*?)}`).ReplaceAllString(latex, `$1<span class="subscript">$2</span>`)
	latex = regexp.MustCompile(`([a-zA-Z]+)_([a-zA-Z0-9])`).ReplaceAllString(latex, `$1<span class="subscript">$2</span>`)

	// Convert multiplication
	latex = regexp.MustCompile(`\\cdot`).ReplaceAllString(latex, `&middot;`)

	// Additional replacements for +, =, etc. can go here
	return latex
}
func translateMathToHTML(input string) string {
	// Define regex to match :math:`...`
	reMath := regexp.MustCompile(":math:`(.*?)`")

	// Replace LaTeX with HTML
	output := reMath.ReplaceAllStringFunc(input, func(m string) string {
		latex := m[7 : len(m)-1] // Extract the part between :math: and `
		html := latexToHTML(latex)
		return fmt.Sprintf("<span class='math'>%s</span>", html)
	})
	return output
}
func addstyletoTHML(htmlinput string) string {
	stylebyte, err := os.ReadFile("style.html")
	styleContent := string(stylebyte)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	return styleContent + "\n" + htmlinput
}
func ParseRSTtoHTML(rst string) string {
	log.Debug("parsing RST to HTML")
	// Convert headings (.. title::)
	reHeading := regexp.MustCompile(`^(\.\. title::)\s*(.*)`)
	rst = reHeading.ReplaceAllString(rst, "<h1>$2</h1>")

	// Convert headers of the form "header\n^^^^^^^" to <h1>
	reh2 := regexp.MustCompile(`(?m)^(.*?)\n\^+$`)
	rst = reh2.ReplaceAllString(rst, "<h2>$1</h2>")

	// Convert empty lines
	reempty := regexp.MustCompile(`(?m)^\s*\n`)
	rst = reempty.ReplaceAllString(rst, "<br>\n")

	// Convert linebreaks (|)
	relinebreak := regexp.MustCompile(`(?m)^\|`)
	rst = relinebreak.ReplaceAllString(rst, "<br>\n")

	// Convert bold text
	reBold := regexp.MustCompile(`\*\*(.*?)\*\*`)
	rst = reBold.ReplaceAllString(rst, "<b>$1</b>")

	// Convert italic text
	reItalic := regexp.MustCompile(`\*(.*?)\*`)
	rst = reItalic.ReplaceAllString(rst, "<i>$1</i>")

	// Convert bullet lists
	reList := regexp.MustCompile(`(?m)^\s*[-*+]\s+(.*)`)
	rst = reList.ReplaceAllString(rst, "<li>$1</li>")

	// Define regex to match :math:`...`
	rst = translateMathToHTML(rst)

	// Return final HTML
	return strings.TrimSpace(rst)
}
