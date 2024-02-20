package effect

func wheel(pos int) (int, int, int) {
    if pos < 85 {
        return pos * 3, 255 - pos * 3, 0
    } else if pos < 170 {
        pos -= 85
        return 255 - pos * 3, 0, pos * 3
    } else {
        pos -= 170
        return 0, pos * 3, 255 - pos * 3
    }
}
