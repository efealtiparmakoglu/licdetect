# licdetect

**Open-source license detector — reads a LICENSE file (or a source file
with a license header) and names the license by its distinctive phrases.
Go standard library only.**

```
$ licdetect .
licdetect: MIT
```

Dependency scanners need a license name, and the name lives in prose: the
opening "Permission is hereby granted, free of charge" *is* MIT, the
"Apache License Version 2.0" header *is* Apache-2.0. licdetect matches
those phrases instead of trusting a filename.

## What it does

- Recognizes MIT, Apache-2.0, BSD-3-Clause, BSD-2-Clause, GPL-3.0,
  GPL-2.0, LGPL-2.1, MPL-2.0, ISC and Unlicense.
- Scans a file, a directory tree, or stdin.
- Reports the first distinctive phrase that matches, so a COPYING file
  full of references is still classified by its own header.
- Exits non-zero when no license is recognized — useful in CI.

## Usage

```sh
licdetect <path>      # file or directory
cat LICENSE | licdetect -
```

Exit codes: 0 a license was recognized, 1 nothing recognized, 2 usage
error.

## A note

This matches phrases, it does not read the whole license. It is a
convenience, not legal advice: a file that mentions several licenses in
its body is identified by its header, and SPDX identifiers should
still be confirmed by a human before shipping.

## Development

```sh
go vet ./...
go test ./...
```

## License

MIT — see [LICENSE](LICENSE).
