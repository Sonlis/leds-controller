package effect

import (
    "testing"
)

func TestWheel(t *testing.T) {
    tests := []struct {
        pos int
        r   int
        g   int
        b   int
    }{
        {0, 0, 255, 0},
        {85, 255, 0, 0},
        {86, 252, 0, 3},
        {170, 0, 0, 255},
        {171, 0, 3, 252},
        {255, 0, 255, 0},
    }
    for _, test := range tests {
        r, g, b := wheel(test.pos)
        if r != test.r || g != test.g || b != test.b {
            t.Errorf("wheel(%d) = %d, %d, %d; want %d, %d, %d", test.pos, r, g, b, test.r, test.g, test.b)
        }
    }
}
