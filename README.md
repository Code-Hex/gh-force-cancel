# gh force-cancel

A GitHub CLI extension to force cancel GitHub Actions workflow runs.

https://stackoverflow.com/a/78608665

## Installation

```bash
gh extension install Code-Hex/gh-force-cancel
```

## Usage

```bash
gh force-cancel <workflow-run-url>
```

Example:
```bash
gh force-cancel https://github.com/owner/repo/actions/runs/123456789
```

## Description

This extension provides a simple way to force cancel GitHub Actions workflow runs. Unlike the normal cancel operation, force cancel immediately terminates the workflow run without waiting for in-progress jobs to complete cleanly.

## Features

- Force cancels GitHub Actions workflow runs using a single command
- Accepts workflow URLs directly from the browser
- Automatically extracts repository and run ID from the URL
- Provides clear success/error messages

## Requirements

- GitHub CLI (`gh`) installed and authenticated
- Write access to the target repository

## Note

Use this command with caution as force cancellation immediately terminates the workflow without proper cleanup.
