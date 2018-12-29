package context

import (
	"testing"
)

func TestSyncWithWaitGroup(t *testing.T) {
	SyncWithWaitGroup()
}

func TestSyncWithChannelSelect(t *testing.T) {
	SyncWithChannelSelect()
}

func TestSyncWithContext(t *testing.T) {
	SyncWithContext()
}

func TestSyncWithChannelSelect2(t *testing.T) {
	SyncWithContext2()
}

