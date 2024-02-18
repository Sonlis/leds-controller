package effect

import (
    "errors"
    "github.com/Sonlis/leds-controller/internal/controller"
    "sync"
)

var (
    mu        sync.Mutex
    stopChan  = make(chan struct{})
    running   = false
)

type EffectConfig struct {
    Red int `json:"red"`
    Green int `json:"green"`
    Blue int `json:"blue"`
    Delay int `json:"delay"`
    Controllers []string `json:"controllers"`
}

func RunEffect(effect string, config EffectConfig, controllers map[string]string) error{
    controllersList := []*controller.Controller{}
    // Check if the controller configured in the request exists.
    for _, controllerName := range config.Controllers {
        if _, ok := controllers[controllerName]; !ok {
            return errors.New("Controller "+controllerName+" not found")
        }
        controllersList = append(controllersList, &controller.Controller{Name: controllerName, IPAdress: controllers[controllerName]})
    }

    // Check if the controllers are up and running.
    for _, ledController := range controllersList {
        err := ledController.CheckStatus()
        if err != nil {
            return err
        }
        ledController.InitPixelsArrays(240)
    }
    mu.Lock()
    defer mu.Unlock()

    if running {
        stopChan <- struct{}{}
    }

    errCh := make(chan error)
    switch effect {
    case "rainbow":
        go runRainbowEffect(config, controllersList, errCh)
    }
    running = true
    err := <-errCh
    if err != nil {
        return errors.New("Error running effect: "+err.Error())
    }
    return nil
}

func StopEffect() error {
    mu.Lock()
    defer mu.Unlock()
    if running {
        stopChan <- struct{}{}
        running = false
    } else {
        return errors.New("No effect is currently running")
    }
    return nil
}
