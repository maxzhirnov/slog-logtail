package logtail

import (
	"fmt"
	"strings"
	"time"
)

func formatMessage(data interface{}, indent int) string {
	indentStr := strings.Repeat("  ", indent)
	switch v := data.(type) {
	case map[string]interface{}:
		var lines []string
		// Handle specific keys first
		for _, key := range []string{"level", "msg", "time", "app_info", "env"} {
			if value, ok := v[key]; ok {
				switch key {
				case "level":
					lines = append(lines, fmt.Sprintf("%s[%s]", indentStr, formatMessage(value, indent+1)))
				case "time":
					lines = append(lines, fmt.Sprintf("%s%s: %s", indentStr, key, formatMessage(parseTime(value), indent+1)))
				default:
					lines = append(lines, fmt.Sprintf("%s%s: %s", indentStr, key, formatMessage(value, indent+1)))
				}
				delete(v, key)
			}
		}
		for key, value := range v {
			lines = append(lines, fmt.Sprintf("%s%s: %s", indentStr, key, formatMessage(value, indent+1)))
		}
		return "\n" + strings.Join(lines, "\n")
	case []interface{}:
		var lines []string
		for i, value := range v {
			lines = append(lines, fmt.Sprintf("%s[%d]: %s", indentStr, i, formatMessage(value, indent+1)))
		}
		return "\n" + strings.Join(lines, "\n")
	default:
		return fmt.Sprintf("%v", v)
	}
}

func parseTime(i interface{}) string {
	switch v := i.(type) {
	case string:
		t, err := time.Parse(time.RFC3339Nano, v)
		if err != nil {
			fmt.Println("error parsinf time")
			t = time.Now()
		}
		return t.Format("02-01-2006 15:04:05")
	default:
		return time.Now().Format("02-01-2006 15:04:05")
	}
}
