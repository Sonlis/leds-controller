// Receives POST request with json data,  and run visualization program accordingly
package main

import (
	"github.com/Sonlis/leds-controller/internal/effect"
    "github.com/Sonlis/leds-controller/internal/controller"
	"github.com/gin-gonic/gin"
    "net/http"
    "os"
    "gopkg.in/yaml.v2"
    "strconv"
)

func setupRouter(controllers map[string]*controller.Controller, pixel_count int) *gin.Engine {
	r := gin.Default()

	// Health endpoint.
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "we living")
	})

    // Returns a list of configured controllers. Useful for the client to know what controllers are available.
    r.GET("/controllers", func(c *gin.Context) {
        c.JSON(http.StatusOK, controllers)
    })

	// Start an effect on a controller.
    // The effect type is specified in the URL, and the configuration is sent in the body of the request.
	r.POST("/effect/:effect", func(c *gin.Context) {
		effectName := c.Params.ByName("effect")
        var config effect.EffectConfig
        if err := c.BindJSON(&config); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"Error parsing request JSON body": err.Error()})
            return
        }

        err := effect.RunEffect(effectName, config, controllers, pixel_count)
		if err == nil {
            c.JSON(http.StatusOK, gin.H{"executed": true, "message": "Effect "+effectName+" started correctly"})
		} else {
            c.JSON(http.StatusOK, gin.H{"executed": false, "message": "Error starting effect "+effectName+": "+err.Error()})
		}
	})

    r.GET("/effect/stop", func(c *gin.Context) {
      err := effect.StopEffect(controllers)
        if err != nil {
            c.JSON(http.StatusOK, gin.H{"executed": false, "message": "Error stopping effect: "+err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"executed": true, "message": "Effect stopped"})
    })

	return r
}

func parseControllers(filePath string) (map[string]*controller.Controller, error) {
    temporary := struct {
        Controllers []controller.Controller `yaml:"controllers"`
    }{}

    yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &temporary)
	if err != nil {
		return nil, err
	}

    controllers := make(map[string]*controller.Controller)

    for _, controller := range temporary.Controllers {
        controllers[controller.Name] = &controller
    }

    return controllers, nil
}

func main() {
    controller_config := "controllers.yaml"
    if value, ok := os.LookupEnv("CONTROLLER_CONFIG_FILE_PATH"); ok {
        controller_config = value
    }
    controllers, err := parseControllers(controller_config)
    if err != nil {
        panic(err)
    }
    pixel_count := 60
    if value, ok := os.LookupEnv("PIXEL_COUNT"); ok {
        pixel_count, err = strconv.Atoi(value)
        if err != nil {
            panic(err)
        }
    }
    router := setupRouter(controllers, pixel_count)
	router.Run(":8082")
}

