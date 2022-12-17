# gh gitignore

This `gh` extension enables developers to quickly bootstrap new projects with the appropriate `.gitignore` through a CLI interface.

## Examples

Bootstrap a node project with the [nodejs gitignore](https://github.com/github/gitignore/blob/main/Node.gitignore).

```bash
gh gitignore node > .gitignore
```

## What types of projects does this extension support?

```bash
Load gitignore files from GitHub into your project

Usage:
  gh-gitignore [flags]

Examples:
gh-gitignore [go]

Flags:
  -h, --help     help for gh-gitignore
  -t, --toggle   Help message for toggle
```