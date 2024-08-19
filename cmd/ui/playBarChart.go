package ui

import (
	"context"
	"fmt"
	"time"

	"github.com/mum4k/termdash/widgets/barchart"
)

func playBarChart(ctx context.Context, bc *barchart.BarChart, values *[]int) error {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := bc.Values(*values, max); err != nil {
				return fmt.Errorf(
					"値の表示でエラー@cmd.playBarChart: %w", err)
			}

		case <-ctx.Done():
			return nil
		}
	}
}
