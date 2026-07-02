# HTML Escape

A Go program that escapes HTML special characters, converting raw HTML tags into safe, display-friendly text. Built to demonstrate string manipulation and real-world web security concepts.

## Example

**Input:**
```
<h1>Hello</h1>
```

**Output:**
```
&lt;h1&gt;Hello&lt;/h1&gt;
```

## Why This Matters

Without escaping, a browser reads `<h1>Hello</h1>` as an actual HTML tag and renders it as a heading. With escaping, the browser displays it as plain, literal text instead.

More importantly, this is how web applications defend against **XSS (Cross Site Scripting)** attacks — where a malicious user injects a script tag into a form or comment box hoping the browser executes it. Escaping the input neutralises it before it ever reaches the page.

## Characters Escaped

| Character | Escaped Version |
|-----------|-----------------|
| `&`       | `&amp;`         |
| `<`       | `&lt;`          |
| `>`       | `&gt;`          |
| `"`       | `&quot;`        |
| `'`       | `&#39;`         |

> Note: `&` is always replaced first. If `<` were replaced first, producing `&lt;`, a subsequent `&` replacement would corrupt it into `&amp;lt;`.

## Project Structure

```
html-escape/
├── main.go       # Entry point — defines the input and prints the result
├── escape.go     # escapeHTML — loops through the input byte by byte
└── replacer.go   # replaceChar — maps each special character to its escaped version
```

All files share `package main` and compile together as one program.

## Concepts Demonstrated

- **Byte-level string manipulation** — walking through a string one character at a time without regex.
- **Switch statements** — clean, readable alternative to long if/else chains for multiple specific values.
- **Multi-file structure** — logic split by responsibility across separate files.
- **Real-world security awareness** — understanding why escaping matters, not just how to do it.

## Running the Program

From inside the project folder:

```bash
go run .
```

> Use `go run .` (not `go run main.go`) since the project spans multiple files.

## Edge Cases Handled

- Empty string `""` → returns `""` without crashing.
- No special characters, e.g. `"Hello"` → returned unchanged.
- Only special characters, e.g. `"<>&"` → returns `"&lt;&gt;&amp;"`.
- `&` appearing naturally in text, e.g. `"fish & chips"` → returns `"fish &amp; chips"`.
- Nested tags, e.g. `"<div><p>Hi</p></div>"` → every `<` and `>` escaped independently.

## Author

Ayomide Ajisegiri (Show Manny)