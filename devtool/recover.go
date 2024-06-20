package devtool

import (
	"os"
	"runtime"
	"strconv"
	"strings"
	
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

type debugStackTrace struct {
	line    int
	path    string
	content string
}

func CreateRecoverPage(assets Node, err error) Node {
	stackTraceList, stackTraceErr := parseStackTrace()
	if stackTraceErr != nil {
		err = stackTraceErr
	}
	return layout(
		assets,
		Div(
			tempest.Class().TextRed(400).FontBold().Mb(4).TextMain(),
			Text("Error: "+err.Error()),
		),
		Div(
			tempest.Class().Grid().Gap(2).Mb(6),
			Range(
				stackTraceList, func(item debugStackTrace, _ int) Node {
					return Div(
						tempest.Class().Border(1).BorderSlate(500).Rounded().P(6).Lh(6),
						Div(
							tempest.Class().TextSky(300).FontSemibold().Mb(2),
							Text(item.path+":"+strconv.Itoa(item.line)),
						),
						Div(
							tempest.Class().TextAmber(200).BgAmber(900).Py(2).Px(4),
							Text(item.content),
						),
					)
				},
			),
		),
		Div(
			tempest.Class().TextCenter(),
			A(
				tempest.Class().TextWhite().Underline().NoUnderline(tempest.Hover()),
				Href(""),
				Text("Refresh"),
			),
		),
	)
}

func parseStackTrace() ([]debugStackTrace, error) {
	var buf [4096]byte
	result := make([]debugStackTrace, 0)
	n := runtime.Stack(buf[:], false)
	stackTrace := string(buf[:n])
	wd, err := os.Getwd()
	if err != nil {
		return result, err
	}
	for _, item := range strings.Split(stackTrace, "\n") {
		if strings.Contains(item, "arcanum") ||
			strings.Contains(item, "/go/") ||
			strings.Contains(item, "panic") ||
			strings.Contains(item, "net/http") ||
			strings.Contains(item, "goroutine") {
			continue
		}
		item = strings.TrimSpace(item)
		if len(item) == 0 || !strings.Contains(item, ".go:") {
			continue
		}
		parts := strings.Split(strings.TrimPrefix(item, wd), " ")
		if len(parts) < 1 {
			continue
		}
		pathLine := parts[0]
		if !strings.Contains(pathLine, ":") {
			continue
		}
		path := pathLine[:strings.LastIndex(pathLine, ":")]
		line, err := strconv.Atoi(pathLine[strings.LastIndex(pathLine, ":")+1:])
		if err != nil {
			return result, err
		}
		fileBytes, err := os.ReadFile(wd + path)
		if err != nil {
			return result, err
		}
		for i, fileLineContent := range strings.Split(string(fileBytes), "\n") {
			if i != line-1 {
				continue
			}
			result = append(
				result, debugStackTrace{
					line:    line,
					path:    path,
					content: strings.TrimSpace(fileLineContent),
				},
			)
		}
	}
	return result, nil
}
