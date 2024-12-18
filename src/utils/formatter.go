package utils

import (
	"fmt"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func FormatHeader(columns []string) string {
	var headerParts []string
	for _, col := range columns {
		headerParts = append(headerParts, fmt.Sprintf("%s%-15s%s", Cyan, strings.ToUpper(col), Reset))
	}
	return strings.Join(headerParts, " | ")
}

func FormatRow(row string) string {
	parts := strings.Split(row, "|")
	for i, part := range parts {
		parts[i] = fmt.Sprintf("%s%-15s%s", Green, strings.TrimSpace(part), Reset)
	}
	return strings.Join(parts, " | ")
}

func PrintTable(columns []string, rows []string) {
	header := FormatHeader(columns)
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	for _, row := range rows {
		if row == "" {
			continue
		}
		fmt.Println(FormatRow(row))
	}
}
