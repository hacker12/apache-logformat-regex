package apachelogformatregex

import (
    "regexp"
    "testing"
)

// TestConvertApacheLogFormatToRegex tests the conversion of Apache log format strings to regex patterns.
func TestConvertApacheLogFormatToRegex(t *testing.T) {
    // Define test cases
    testCases := []struct {
        name           string
        format         string
        expectedRegex  string
        testLogEntry   string
        shouldMatch    bool
    }{
        {
            name:           "Standard Common Log Format",
            format:         `%h %l %u %t "%r" %>s %b`,
            expectedRegex:  `^(?P<host>[^ ]+) (?P<logname>[^ ]*) (?P<user>[^ ]*) (?P<time>\[[^]]+\]) "(?P<request>[^"]*)" (?P<status>[0-9]+) (?P<size>[0-9]+|-)$`,
            testLogEntry:   `127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326`,
            shouldMatch:    true,
        },
        {
            name:           "Combined Log Format",
            format:         `%h %l %u %t "%r" %>s %b "%{Referer}i" "%{User-Agent}i"`,
            expectedRegex:  `^(?P<host>[^ ]+) (?P<logname>[^ ]*) (?P<user>[^ ]*) (?P<time>\[[^]]+\]) "(?P<request>[^"]*)" (?P<status>[0-9]+) (?P<size>[0-9]+|-) "(?P<referer>[^"]*)" "(?P<useragent>[^"]*)"$`,
            testLogEntry:   `127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326 "http://www.example.com/start.html" "Mozilla/4.08 [en] (Win98; I ;Nav)"`,
            shouldMatch:    true,
        },
        // Add more test cases as needed for different formats and scenarios
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            regexPattern := ConvertApacheLogFormatToRegex(tc.format)
            if regexPattern != tc.expectedRegex {
                t.Errorf("Expected regex pattern does not match. Got %s, want %s", regexPattern, tc.expectedRegex)
            }

            re := regexp.MustCompile(regexPattern)
            if match := re.MatchString(tc.testLogEntry); match != tc.shouldMatch {
                t.Errorf("Log entry did not match as expected. Got %t, want %t, using pattern: %s", match, tc.shouldMatch, regexPattern)
            }
        })
    }
}
