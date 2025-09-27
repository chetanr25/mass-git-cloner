package ui

import (
	"fmt"
	"strings"
	"time"
)

type ProgressTracker struct {
	total     int
	completed int
	failed    int
	current   string
	startTime time.Time
}

func InitProgressTracker(total int) *ProgressTracker {
	return &ProgressTracker{
		total:     total,
		completed: 0,
		failed:    0,
		current:   "",
		startTime: time.Now(),
	}
}

func (p *ProgressTracker) Update(current string) {
	p.current = current
	p.display()
}

func (p *ProgressTracker) Success(repoName string) {
	p.completed++
	p.current = fmt.Sprintf("✅ %s", repoName)
	p.display()
	time.Sleep(100 * time.Millisecond)
}

func (p *ProgressTracker) Failure(repoName string, err error) {
	p.failed++
	p.current = fmt.Sprintf("❌ %s: %v", repoName, err)
	p.display()
	time.Sleep(500 * time.Millisecond)
}

func (p *ProgressTracker) display() {

	fmt.Print("\r" + strings.Repeat(" ", 100) + "\r")

	percentage := float64(p.completed+p.failed) / float64(p.total) * 100
	elapsed := time.Since(p.startTime)

	barWidth := 30
	filled := int(float64(barWidth) * percentage / 100)
	bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

	status := fmt.Sprintf("[%s] %.1f%% (%d/%d) %s",
		bar,
		percentage,
		p.completed+p.failed,
		p.total,
		elapsed.Truncate(time.Second),
	)

	fmt.Print(status)

	if p.current != "" {
		fmt.Printf("\n%s", p.current)
		if p.completed+p.failed < p.total {
			fmt.Print("\n")
		}
	}
}

func (p *ProgressTracker) Finish() {
	fmt.Print("\r" + strings.Repeat(" ", 100) + "\r")
	elapsed := time.Since(p.startTime)
	// Color codes
	reset := "\033[0m"
	green := "\033[32m"
	red := "\033[31m"
	blue := "\033[34m"
	yellow := "\033[33m"
	bold := "\033[1m"

	fmt.Printf("%s%s🎉 Cloning completed!%s\n", bold, green, reset)
	fmt.Printf("   %sTotal:%s %d repositories\n", blue, reset, p.total)
	fmt.Printf("   %sSuccessful:%s %d\n", green, reset, p.completed)
	fmt.Printf("   %sFailed:%s %d\n", red, reset, p.failed)
	fmt.Printf("   %sDuration:%s %s\n", yellow, reset, elapsed.Truncate(time.Second))

	if p.failed > 0 {
		fmt.Printf("   Success rate: %.1f%%\n", float64(p.completed)/float64(p.total)*100)
	}
}

func (p *ProgressTracker) Total() int {
	return p.total
}
