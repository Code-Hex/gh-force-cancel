package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
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
	runID, repo, err := parseURL(url)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	endpoint := fmt.Sprintf("repos/%s/actions/runs/%s/force-cancel", repo, runID)
	args := []string{
		"api",
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

func parseURL(url string) (runID, repo string, err error) {
	// リポジトリ名を抽出（例: owner/repo）
	repoRegex := regexp.MustCompile(`github\.com/([^/]+/[^/]+)`)
	matches := repoRegex.FindStringSubmatch(url)
	if len(matches) != 2 {
		return "", "", fmt.Errorf("invalid GitHub URL format: %q", url)
	}
	repo = matches[1]

	// ランIDを抽出
	runRegex := regexp.MustCompile(`/runs/(\d+)`)
	matches = runRegex.FindStringSubmatch(url)
	if len(matches) != 2 {
		return "", "", fmt.Errorf("could not find run ID in URL: %q", url)
	}
	runID = matches[1]

	return runID, repo, nil
}
