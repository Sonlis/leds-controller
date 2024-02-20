package effect

import (
    "testing"
)

func TestFilterUnchangedPixels(t *testing.T) {
    pixels := [][4]int{{1, 2, 3, 4}, {2, 2, 3, 4}, {3, 2, 3, 4}, {4, 2, 3, 4}}
    previousPixels := [][4]int{{1, 2, 3, 4}, {2, 2, 3, 4}, {3, 2, 3, 4}, {4, 2, 3, 4}}
    filteredPixels := filterUnchangedPixels(pixels, previousPixels)
    if len(filteredPixels) != 0 {
        t.Errorf("Expected filtered pixels to be empty, got %v", filteredPixels)
    }

    pixels = [][4]int{{1, 2, 3, 4}, {2, 2, 3, 4}, {3, 2, 3, 4}, {4, 2, 3, 4}}
    previousPixels = [][4]int{{1, 2, 3, 4}, {2, 2, 3, 4}, {3, 2, 3, 4}, {4, 2, 3, 5}}
    filteredPixels = filterUnchangedPixels(pixels, previousPixels)
    if len(filteredPixels) != 1 {
        t.Errorf("Expected filtered pixels to have 1 element, got %v", filteredPixels)
    }
    for _, pixel := range filteredPixels {
        if pixel != [4]int{4, 2, 3, 4} {
            t.Errorf("Expected filtered pixels to be [4 2 3 4] got %v", pixel)
        }
    }
}

func TestSplitPixelsIntoPackets(t *testing.T) {
    pixels := [][4]int{{1, 2, 3, 4}, {2, 2, 3, 4}, {3, 2, 3, 4}, {4, 2, 3, 4}}
    packets := splitPixelsIntoPackets(pixels)
    if len(packets) != 1 {
        t.Errorf("Expected to have 1 packet, got %d of value %v", len(packets), packets)
    }
    for _, pixels := range packets {
        if len(pixels) != 4 {
            t.Errorf("Expected to have 4 pixels, got %d of value %v", len(pixels), pixels)
        }
    }

    pixels = [][4]int{}
    for i := 0; i < 128; i++ {
        pixels = append(pixels, [4]int{i, 2, 3, 4})
    }
    packets = splitPixelsIntoPackets(pixels)
    if len(packets) != 1 {
        t.Errorf("Expected to have 1 packet, got %d of value %v", len(packets), packets)
    }
    for _, pixels := range packets {
        if len(pixels) != 128 {
            t.Errorf("Expected packet to have 128 pixels, got %d of value %v", len(pixels), pixels)
        }
    }

    pixels = [][4]int{}
    for i := 0; i < 256; i++ {
        pixels = append(pixels, [4]int{i, 2, 3, 4})
    }
    packets = splitPixelsIntoPackets(pixels)
    if len(packets) != 2 {
        t.Errorf("Expected to have 2 packets, got %d of value %v", len(packets), packets)
    }
    for _, pixels := range packets {
        if len(pixels) != 128 {
            t.Errorf("Expected to have 128 pixels, got %d of value %v", len(pixels), pixels)
        }
    }
}

func TestBuildPackets(t *testing.T) {
    pixels := [][4]int{{1, 2, 3, 4}, {2, 2, 3, 4}, {3, 2, 3, 4}, {4, 2, 3, 4}}
    previousPixels := [][4]int{{1, 2, 3, 4}, {2, 2, 3, 4}, {3, 2, 3, 4}, {4, 2, 3, 5}}
    packets := buildPackets(pixels, previousPixels)
    if len(packets) != 1 {
        t.Errorf("Expected to have 1 packets, got %d packet of value %v", len(packets), packets)
    }
    for _, pixels := range packets {
        if len(pixels) != 1 {
            t.Errorf("Expected to have 1 pixel, got %d of value %v", len(pixels), pixels)
        }
        for _, pixel := range pixels {
            if pixel != [4]int{4, 2, 3, 4} {
                t.Errorf("Expected pixel to be [4 2 3 4], got %v", pixel)
            }
    }
}
    pixels = [][4]int{}
    for i := 0; i < 256; i++ {
        pixels = append(pixels, [4]int{i, i, i+1, i+2})
    }
    previousPixels = [][4]int{}
    for i := 0; i < 256; i++ {
        previousPixels = append(previousPixels, [4]int{i, i, i+10, i+20})
    }
    packets = buildPackets(pixels, previousPixels)
    if len(packets) != 2 {
        t.Errorf("Expected to have 2 packets, got %d packets of values %v", len(packets), packets)
    }
}
