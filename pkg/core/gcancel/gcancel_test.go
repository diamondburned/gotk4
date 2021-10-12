package gcancel

import (
	"context"
	"testing"
	"time"
)

// TestCancelThisBackground tests when this context is canceled with a
// background parent.
func TestCancelThisBackground(t *testing.T) {
	ctx, cancel := WithCancel(context.Background())
	defer cancel()

	done := cancellableDoneCh(ctx)

	cancel()
	waitCh(t, done)
}

// TestCancelThisParent tests when this context is canceled but not the parent.
func TestCancelThisParent(t *testing.T) {
	pctx, pcancel := context.WithCancel(context.Background())
	defer pcancel()

	ctx, cancel := WithCancel(pctx)
	defer cancel()

	done := cancellableDoneCh(ctx)

	cancel()
	waitCh(t, done)
}

// TestCancelParent tests when the parent context is canceled.
func TestCancelParent(t *testing.T) {
	pctx, pcancel := context.WithCancel(context.Background())
	defer pcancel()

	ctx, cancel := WithCancel(pctx)
	defer cancel()

	done := cancellableDoneCh(ctx)

	pcancel()
	waitCh(t, done)
}

func cancellableDoneCh(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})

	gc := GCancellableFromContext(ctx)
	if gc == nil {
		panic("given ctx is not *Cancellable")
	}

	gc.Connect("cancelled", func() {
		go func() {
			done <- struct{}{}
		}()
	})

	return done
}

func waitCh(t *testing.T, ch <-chan struct{}) {
	t.Helper()

	select {
	case <-ch:
		return
	case <-time.After(time.Second):
		t.Error("timeout waiting for ch")
	}
}
