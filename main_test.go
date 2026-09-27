package main

import "testing"

// Each license's own header must be recognised by its distinctive phrase.
func TestDetectsKnownLicenses(t *testing.T) {
	cases := []struct {
		name string
		text string
		want string
	}{
		{"MIT", "MIT License\n\nPermission is hereby granted, free of charge", "MIT"},
		{"Apache-2.0", "Apache License\nVersion 2.0, January 2004", "Apache-2.0"},
		{"GPL-3.0", "GNU GENERAL PUBLIC LICENSE\nVersion 3, 29 June 2007", "GPL-3.0"},
		{"GPL-2.0", "GNU GENERAL PUBLIC LICENSE\nVersion 2, June 1991", "GPL-2.0"},
		{"LGPL-2.1", "GNU LESSER GENERAL PUBLIC LICENSE\nVersion 2.1", "LGPL-2.1"},
		{"AGPL-3.0", "GNU AFFERO GENERAL PUBLIC LICENSE\nVersion 3", "AGPL-3.0"},
		{"MPL-2.0", "Mozilla Public License Version 2.0", "MPL-2.0"},
		{"Unlicense", "This is free and unencumbered software released into the public domain", "Unlicense"},
		{"0BSD", "Permission to use, copy, modify, and/or distribute this software", "0BSD"},
		{"BSD-3-Clause", "Redistribution and use in source and binary forms", "BSD-3-Clause"},
	}
	for _, c := range cases {
		found := detect(c.text)
		if len(found) == 0 {
			t.Errorf("%s: nothing detected in %q", c.name, c.text)
			continue
		}
		matched := false
		for _, name := range found {
			if name == c.want {
				matched = true
			}
		}
		if !matched {
			t.Errorf("%s: got %v, want it to include %q", c.name, found, c.want)
		}
	}
}

func TestDetectionIsCaseInsensitive(t *testing.T) {
	for _, text := range []string{
		"permission is hereby granted, free of charge",
		"PERMISSION IS HEREBY GRANTED, FREE OF CHARGE",
		"Permission Is Hereby Granted, Free Of Charge",
	} {
		found := detect(text)
		if len(found) == 0 || found[0] != "MIT" {
			t.Errorf("%q -> %v, want MIT", text, found)
		}
	}
}

func TestUnknownLicenseReturnsNothing(t *testing.T) {
	for _, text := range []string{"", "do whatever you want", "all rights reserved"} {
		if found := detect(text); len(found) != 0 {
			t.Errorf("%q -> %v, want no match", text, found)
		}
	}
}

func TestEveryListedLicenseHasAPhrase(t *testing.T) {
	seen := map[string]bool{}
	for _, lic := range LICENSES {
		if lic.name == "" || lic.pattern == "" {
			t.Errorf("incomplete entry: %+v", lic)
		}
		if seen[lic.name] {
			t.Errorf("duplicate license name %q", lic.name)
		}
		seen[lic.name] = true
	}
	// the list must actually detect its own patterns
	for _, lic := range LICENSES {
		found := detect(lic.pattern)
		if len(found) == 0 {
			t.Errorf("%s: its own phrase is not detected", lic.name)
		}
	}
}

func TestRealMitsLicenseHeader(t *testing.T) {
	text := `MIT License

Copyright (c) 2026 The dnsprobe Authors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction.`
	found := detect(text)
	if len(found) != 1 || found[0] != "MIT" {
		t.Errorf("got %v, want [MIT]", found)
	}
}
