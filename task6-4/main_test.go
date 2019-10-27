package main

import "testing"

func TestGetDiscriminant(t *testing.T) {
	d := GetDiscriminant(2, 5, 0)
	if d != 25 {
		t.Error("Expected 25, got ", d)
	}
}

func TestGetX1X2(t *testing.T) {
	x1, x2, err := GetX1X2(2, 5, -3)

	if err != nil {
		t.Error("Expected no error, got error", err)
	}

	if x1 != 0.5 {
		t.Error("Expected 0.5, got ", x1)
	}

	if x2 != -3 {
		t.Error("Expected -3, got ", x2)
	}

	x1, x2, err = GetX1X2(2, 4, 2)

	if err != nil {
		t.Error("Expected no error, got error", err)
	}

	if x1 != -1 {
		t.Error("Expected -1, got ", x1)
	}

	if x2 != -1 {
		t.Error("Expected -1, got ", x2)
	}

	_, _, err = GetX1X2(4, 1, 1)

	if err == nil {
		t.Error("Expected error, got no error")
	}
}
