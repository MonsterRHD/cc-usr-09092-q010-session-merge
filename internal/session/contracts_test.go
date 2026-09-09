package session

import "testing"

func TestBoundaryEventCarriesLogicalVersion(t *testing.T) {
	event := BoundaryEvent{EventID: "edge-1", DeviceID: "tablet-3", LogicalVersion: 5}
	if event.LogicalVersion == 0 {
		t.Fatal("观看边界事件必须携带逻辑版本")
	}
}
