package effect

import (
	"context"
	"errors"
	"log"

	"github.com/Sonlis/leds-controller/internal/controller"
)

var (
    contexts  = map[string]context.CancelFunc{}
)

type EffectConfig struct {
    Red int `json:"red"`
    Green int `json:"green"`
    Blue int `json:"blue"`
    Delay int `json:"delay"`
    Controllers []string `json:"controllers"`
}

func RunEffect(effect string, config EffectConfig, controllers map[string]*controller.Controller, pixel_count int) error{
    ctx, cancel := context.WithCancel(context.Background())
    // Check if the controller configured in the request exists.
    for _, controllerName := range config.Controllers {
        if _, ok := controllers[controllerName]; !ok {
            cancel()
            return errors.New("Controller "+controllerName+" not found")
        }
    }

    // Check if the controllers are up and running.
    for _, ledController := range controllers {
        err := ledController.CheckStatus()
        log.Println("Checking status of controller "+ledController.Name)
        if err != nil {
            cancel()
            return err
        }
        ledController.InitPixelsArrays(pixel_count)
        err = ledController.Connect()
        if err != nil {
            cancel()
            return err
        }
    }
    if len(contexts) > 0 {
        for effectRunning, cancelFunc := range contexts {
            cancelFunc()
            delete(contexts, effectRunning)
        }

    }
    contexts[effect] = cancel

    switch effect {
    case "rainbow":
        go runRainbowEffect(controllers, ctx, pixel_count)
    }
    return nil
}


// StopEffect stops the effect running on the controllers.
// It sends a cancel signal to the context of the effect, and clears the pixels on the controllers.
func StopEffect(ledControllers map[string]*controller.Controller) error {
    for _, cancelFunc := range contexts {
        cancelFunc()
    }
    for _, ledController := range ledControllers {
        err := ledController.Clear()
        if err != nil {
            return err
        }
    }
    return nil
}
