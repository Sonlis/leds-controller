package effect

func rainbow(rainbowGenerator, pixel_count int) [][4]int {
    pixels := make([][4]int, pixel_count)
    for i := 0; i < pixel_count; i++ {
        r, g, b := wheel(((i * 256 % pixel_count) +rainbowGenerator) & 255)
        pixels[i] = [4]int{i, r, g, b}
    }
    return pixels
}
