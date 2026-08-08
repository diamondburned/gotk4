package intern

import (
	"reflect"
	"testing"
)

func TestBoxDummyContainsScannedPointer(t *testing.T) {
	typ := reflect.TypeFor[boxDummy]()
	found := false
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Type.Kind() == reflect.Pointer {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("boxDummy has no pointer field and is eligible for tiny allocation")
	}
}

func TestFinalizerCleanupQueuesWeakEntries(t *testing.T) {
	if !shouldQueueFinalizerCleanup(entryActive, false) {
		t.Fatal("a finalizer must queue cleanup for an entry without a strong box")
	}
}

func TestFinalizerCleanupRejectsStrongOrInactiveEntries(t *testing.T) {
	if shouldQueueFinalizerCleanup(entryActive, true) {
		t.Fatal("a strongly held box must not be finalized")
	}
	if shouldQueueFinalizerCleanup(entryCleanupQueued, false) {
		t.Fatal("queued cleanup must not be scheduled twice")
	}
}
