package apachelogformatregex

import (
    "fmt"
    "strings"
)

// ApacheFormatSpecifiers maps Apache log format specifiers to regular expressions.
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
    // Add more directives as needed
}

// ConvertApacheLogFormatToRegex takes an Apache log format string and converts it into a regex.
func ConvertApacheLogFormatToRegex(format string) string {
    // First, replace Apache format specifiers with regex patterns
    for key, value := range ApacheFormatSpecifiers {
        format = strings.Replace(format, key, value, -1)
    }

    // Then, handle literal parts of the format, such as literal spaces and quotes
    // Note: Since we already replaced all specifiers, we assume no literal quotes need additional handling here
    // If the log format includes special regex characters, they should be escaped here as necessary

    return "^" + format + "$"
}

func main() {
    logFormat := `%h %l %u %t "%r" %>s %b "%{Referer}i" "%{User-Agent}i"` // Example Apache log format
    regexPattern := ConvertApacheLogFormatToRegex(logFormat)
    fmt.Println("Regex Pattern:", regexPattern)
}
