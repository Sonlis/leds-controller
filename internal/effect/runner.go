package effect

import (
    "time"
	"github.com/Sonlis/leds-controller/internal/controller"
)

// filterUnchangedPixels filters out pixels that have not changed since the last frame.
// This is useful to reduce the amount of data sent to the controllers. Network congestion is
// more of an issue than processing power, especially over Wifi or when the amount of controllers
// is high.
func filterUnchangedPixels(pixels [][4]int, previousPixels [][4]int) [][4]int {
    filteredPixels := [][4]int{}
    for index, row := range pixels {
        if pixels[index][1] != previousPixels[index][1] || pixels[index][2] != previousPixels[index][2] || pixels[index][3] != previousPixels[index][3] {
            filteredPixels = append(filteredPixels, row)
        }
    }
    return filteredPixels
}

func splitPixelsIntoPackets(pixels [][4]int) [][][4]int {
    packets := [][][4]int{}
    size := 127
    for i := 0; i < len(pixels); i += size{
        j := i + size
        if j > len(pixels) {
            j = len(pixels)
        }
        packets = append(packets, pixels[i:j])
    }
    return packets
}

func runRainbowEffect(config EffectConfig, controllers []*controller.Controller, ch chan error) {
    packets := [][][4]int{}
    previousPixels := make([][4]int, 240)
    for _, ledController := range controllers {
        err := ledController.Connect()
        if err != nil {
            ch <- err
        }
    }
    for {
        select {
        case <-stopChan:
            return // Stop the function if a signal is received on stopChan
        default:
            for j := 0; j < 1024; j++ {
                pixels := rainbow(config, j)
                filteredPixels := filterUnchangedPixels(pixels, previousPixels)
                previousPixels = pixels
                if len(filteredPixels) > 127 {
                    packets = splitPixelsIntoPackets(filteredPixels)
                } else {
                    packets = [][][4]int{filteredPixels}
                }
                for _, ledController := range controllers {
                    err := ledController.SendPackets(packets)
                    if err != nil {
                        ch <- err
                    }
                }
                time.Sleep(16 * time.Millisecond)
            }
        }
    }
}
