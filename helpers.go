package contentsynthesizer

import "github.com/mitranim/jsonfmt"

func fixJSON(data string) string {
	return jsonfmt.FormatString(jsonfmt.Default, data)
}
