package effect

func rainbow(config EffectConfig, rainbowGenerator int) [][4]int {
    pixels := make([][4]int, 240)
    for i := 0; i < 240; i++ {
        r, g, b := wheel(((i * 256 % 240) +rainbowGenerator) & 255)
        pixels[i] = [4]int{i, r, g, b}
    }
    return pixels
}
