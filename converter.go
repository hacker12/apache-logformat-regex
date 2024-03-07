package apachelogformatregex

import (
    "fmt"
    "strings"
)

// ApacheFormatSpecifiers maps Apache log format specifiers to regular expressions.
// This map covers a wide range of specifiers but might not include every custom specifier.
var ApacheFormatSpecifiers = map[string]string{
    "%h": `(?P<host>[^ ]+)`,                              // Remote hostname
    "%l": `(?P<logname>[^ ]*)`,                           // Remote logname
    "%u": `(?P<user>[^ ]*)`,                              // Remote user
    "%t": `(?P<time>\[[^]]+\])`,                          // Time the request was received
    "%r": `(?P<request>[^"]*)`,                           // First line of request
    "%>s": `(?P<status>[0-9]+)`,                          // Status
    "%b": `(?P<size>[0-9]+|-)`,                           // Size of response in bytes
    "%{Referer}i": `(?P<referer>[^"]*)`,                  // Referer HTTP request header
    "%{User-Agent}i": `(?P<useragent>[^"]*)`,             // User-Agent HTTP request header
    "%{Host}i": `(?P<hostheader>[^"]*)`,                  // Host HTTP request header
    // Add more directives as needed
}

// ConvertApacheLogFormatToRegex takes an Apache log format string and converts it into a regex.
func ConvertApacheLogFormatToRegex(format string) string {
    // Escape string literals between quotes that are not format specifiers
    literalStart := false
    var literalBuilder strings.Builder
    for _, runeValue := range format {
        if runeValue == '"' && !literalStart {
            literalStart = true
            literalBuilder.WriteRune(runeValue)
        } else if runeValue == '"' && literalStart {
            literalStart = false
            literalBuilder.WriteString(`[^"]*`) // Replace literals between quotes with regex
            literalBuilder.WriteRune(runeValue)
        } else if literalStart {
            continue // Skip characters inside literal quotes
        } else {
            literalBuilder.WriteRune(runeValue)
        }
    }
    escapedFormat := literalBuilder.String()

    // Replace Apache format specifiers with regex patterns
    for key, value := range ApacheFormatSpecifiers {
        escapedFormat = strings.Replace(escapedFormat, key, value, -1)
    }

    return "^" + escapedFormat + "$"
}

func main() {
    logFormat := `%h %l %u %t "%r" %>s %b "%{Referer}i" "%{User-Agent}i"` // Example Apache log format
    regexPattern := ConvertApacheLogFormatToRegex(logFormat)
    fmt.Println("Regex Pattern:", regexPattern)
}
