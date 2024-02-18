// Receives POST request with json data,  and run visualization program accordingly
package main

import (
	"github.com/Sonlis/leds-controller/internal/effect"
	"github.com/gin-gonic/gin"
    "net/http"
    "os"
    "gopkg.in/yaml.v2"
)

func setupRouter(controllers map[string]string) *gin.Engine {
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

        err := effect.RunEffect(effectName, config, controllers)
		if err == nil {
            c.JSON(http.StatusOK, gin.H{"executed": true, "message": "Effect "+effectName+" started correctly"})
		} else {
            c.JSON(http.StatusOK, gin.H{"executed": false, "message": "Error starting effect "+effectName+": "+err.Error()})
		}
	})

    r.GET("/effect/stop", func(c *gin.Context) {
        err := effect.StopEffect()
        if err != nil {
            c.JSON(http.StatusOK, gin.H{"executed": false, "message": "Error stopping effect: "+err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"executed": true, "message": "Effect stopped"})
    })

	return r
}

func parseControllers(filePath string) (map[string]string, error) {
    var temporary struct {
        Controllers []struct {
            Name string `yaml:"name"`
            IPAdress string `yaml:"ipAdress"`
            Port int `yaml:"port"`
        } `yaml:"controllers"`
    }

    yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &temporary)
	if err != nil {
		return nil, err
	}

    controllers := make(map[string]string)

    for _, controller := range temporary.Controllers {
        controllers[controller.Name] = controller.IPAdress
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
    router := setupRouter(controllers)
	router.Run(":8082")
}

