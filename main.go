package main

import (
	"fmt"
	"minimal_rst_html_converter/utils"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
)

type Converter struct {
	Rstcontentfile   string `help:"RST file to convert"`
	Rstcontentstring string `help:"RST content in string"`
	Htmloutput       string `help:"outpufile for html content"`
}

func (c *Converter) Run() error {
	log.Debug("Converter Run")
	rstContent, errgetrst := c.getRstContent()
	if errgetrst != nil {
		log.Error("error get rst content: ", "err", errgetrst)
	}
	htmlContent := utils.ParseRSTtoHTML(rstContent)
	if c.Htmloutput != "" {
		utils.WriteHTMLToFile(htmlContent, c.Htmloutput)
	}
	log.Info(htmlContent)

	return nil
}
func (c *Converter) getRstContent() (string, error) {
	log.Debug("getRstContent")
	if c.Rstcontentstring != "" {
		return c.Rstcontentstring, nil
	} else if c.Rstcontentfile != "" {
		if strings.Contains(c.Rstcontentfile, ".rst") {
			rstbyte, errreadfile := os.ReadFile(c.Rstcontentfile)
			rstContent := string(rstbyte)
			if errreadfile != nil {
				return "", fmt.Errorf("error: %s", errreadfile)
			}
			return rstContent, nil
		} else {
			return "", fmt.Errorf("error: input is no RST file")
		}
	} else {
		return "", fmt.Errorf("nothing to convert")
	}
}

var cli struct {
	Debug     bool      `help:"activate debug loggin" `
	Converter Converter `cmd:"" help:"converter"`
}

func main() {

	ctx := kong.Parse(&cli)
	if cli.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("log debuglevel active")
	}

	errrun := ctx.Run()
	if errrun != nil {
		log.Error("error: ", "err", errrun)
	}

}
