# gh gitignore

This `gh` extension enables developers to quickly bootstrap new projects with the appropriate `.gitignore` through a CLI interface.

## Examples

Bootstrap a node project with the [nodejs gitignore](https://github.com/github/gitignore/blob/main/Node.gitignore).

```bash
gh gitignore node > .gitignore
```

## What types of projects does this extension support?

```bash
Load gitignore files from GitHub into your project.

Usage:
  gh-gitignore [arg] [flags]

Examples:
gh-gitignore node: returns the .gitignore for node projects

Flags:
  -h, --help     help for gh-gitignore
```
