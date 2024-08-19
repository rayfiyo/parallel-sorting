package ui

import (
	"context"
	"fmt"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/barchart"
)

func Draw(values []int, title string) error {
	if len(values) > max {
		return fmt.Errorf("upper limit exceeded@cmd.Ui")
	}

	t, err := tcell.New()
	if err != nil {
		return fmt.Errorf("1@cmd.Ui: %w", err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())
	bc, err := barchart.New(
		barchart.ShowValues(),
		barchart.BarWidth(1+50/len(values)), // len(values) > 50 ? 1 : 2
	)
	if err != nil {
		cancel()
		return fmt.Errorf("2@cmd.Ui: %w", err)
	}

	go PlaySort(ctx, bc, values)

	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle(title),
		container.PlaceWidget(bc),
	)
	if err != nil {
		cancel()
		return fmt.Errorf("3@cmd.Ui: %w", err)
	}

	keyboardHandler := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c,
		termdash.KeyboardSubscriber(keyboardHandler)); err != nil {
		return fmt.Errorf("4@cmd.Ui: %w", err)
	}

	return nil
}
