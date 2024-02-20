package effect

import (
	"testing"
)

func TestRainbow(t *testing.T) {
    pixels := rainbow(0, 10)
    expected := [][4]int{{0, 0, 255, 0}, {1, 18, 237, 0}, {2, 6, 249, 0}, {3, 24, 231, 0}, {4, 12, 243, 0},
        {5, 0, 255, 0}, {6, 18, 237, 0}, {7, 6, 249, 0}, {8, 24, 231, 0}, {9, 12, 243, 0}}
    for i, pixel := range pixels {
        if pixel != expected[i] {
            t.Errorf("Expected pixel to be %v, got %v", expected[i], pixel)
        }
    }
}
