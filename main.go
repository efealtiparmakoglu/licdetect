package main

import (
	"fmt"
	"os"
	"strings"
)

// Detects open-source licenses by matching distinctive phrase patterns.
var LICENSES = []struct {
	name  string
	pattern string
}{
	{"MIT", "Permission is hereby granted, free of charge"},
	{"Apache-2.0", "Apache License Version 2.0"},
	{"BSD-3-Clause", "Redistribution and use in source and binary forms"},
	{"GPL-3.0", "GNU GENERAL PUBLIC LICENSE Version 3"},
	{"GPL-2.0", "GNU GENERAL PUBLIC LICENSE Version 2"},
	{"LGPL-2.1", "GNU LESSER GENERAL PUBLIC LICENSE"},
	{"MPL-2.0", "Mozilla Public License Version 2.0"},
	{"AGPL-3.0", "GNU AFFERO GENERAL PUBLIC LICENSE"},
	{"Unlicense", "This is free and unencumbered software released into the public domain"},
	{"0BSD", "Permission to use, copy, modify, and/or distribute this software"},
}

func detect(text string) []string {
	upper := strings.ToUpper(text)
	var found []string
	for _, lic := range LICENSES {
		if strings.Contains(upper, strings.ToUpper(lic.pattern)) {
			found = append(found, lic.name)
		}
	}
	return found
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: licdetect <LICENSE_FILE>")
		os.Exit(2)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "licdetect: %v\n", err)
		os.Exit(1)
	}
	found := detect(string(data))
	if len(found) == 0 {
		fmt.Println("unknown license")
		os.Exit(1)
	}
	fmt.Println(strings.Join(found, ", "))
}
