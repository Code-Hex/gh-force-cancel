package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: gh force-cancel <workflow-run-url>")
		os.Exit(1)
	}

	if err := run(flag.Arg(0)); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(url string) error {
	hostname, runID, repo, err := parseURL(url)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	endpoint := fmt.Sprintf("repos/%s/actions/runs/%s/force-cancel", repo, runID)
	args := []string{
		"api",
		"--hostname", hostname,
		"--method", "POST",
		"-H", "Accept: application/vnd.github+json",
		"-H", "X-GitHub-Api-Version: 2022-11-28",
		endpoint,
	}

	// gh コマンドを実行
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, "gh", args...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to execute gh command: %w\n%s", err, output)
	}

	fmt.Printf("Successfully initiated force cancellation of workflow run %s in repository %s\n", runID, repo)
	return nil
}

func parseURL(urlstr string) (hostname, runID, repo string, err error) {
	// リポジトリ名を抽出（例: owner/repo）
	u, err := url.Parse(urlstr)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to parse URL: %w", err)
	}
	paths := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")
	if len(paths) != 5 || paths[2] != "actions" || paths[3] != "runs" {
		return "", "", "", fmt.Errorf("invalid GitHub URL format: %q", urlstr)
	}
	repo = fmt.Sprintf("%s/%s", paths[0], paths[1])
	runID = paths[4]
	return u.Hostname(), runID, repo, nil
}
