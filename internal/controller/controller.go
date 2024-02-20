package controller

import (
	"errors"
	"net"
	"time"
	"github.com/prometheus-community/pro-bing"
    "fmt"
)

type Controller struct {
    IPAdress string `yaml:"ipAdress" json:"ip_adress"`
    Port int `yaml:"port" json:"port"`
    Name string `yaml:"name" json:"name"`
    pixels [][4]int
    net.Conn
}

// Connect sets up the client connection to the controller.
// Its name is a bit confusing, as it does not send any packets to the controller.
func (c *Controller) Connect() error {
    conn, err := net.Dial("udp", c.IPAdress+":"+fmt.Sprint(c.Port))
    if err != nil {
        return errors.New("Error connecting to controller "+c.Name+": "+err.Error())
    }
    c.Conn = conn
    return nil
}

func (c *Controller) InitPixelsArrays(pixelCount int) {
    c.pixels = make([][4]int, pixelCount)
}


// SendPackets sends packets to the controller. Each packet is a list of pixels, and each pixel is a list of 4 integers: the index of the pixel, and the RGB values.
// As we are sending only pixels that change, the controller needs to know the index of the pixel to update. It cannot rely on the order of the pixels in the packet.
func (c *Controller) SendPackets(packets [][][4]int) error {
    packetsList := []byte{}
    for _, packet := range packets {
        for _, pixel := range packet {
            packetsList = append(packetsList, byte(pixel[0]), byte(pixel[1]), byte(pixel[2]), byte(pixel[3]))
        }
    }
    _, err := c.Conn.Write(packetsList)
    return err
}


// CheckStatus send ICMP packets to check if the IPv4 host is reachable.
// We only send one packet and wait for 2 seconds for a response. If we don't get a response, we consider the host as unreachable.
// Reliability could be improved, with a higher number of packets sent.
func (c *Controller) CheckStatus() error {
    ip := net.ParseIP(c.IPAdress)
    if ip.To4() == nil {
        return errors.New("IP address "+c.IPAdress+" is not an IPv4 address")
    }

    timeout := 2 * time.Second

    pinger, err := probing.NewPinger(c.IPAdress)
    if err != nil {
        panic(err)
    }

    pinger.Count = 1
    pinger.Timeout = timeout
    err = pinger.Run() // Blocks until finished.
    if err != nil {
        return errors.New("Error checking health of controller "+c.Name+": "+err.Error())
    }

    if pinger.Statistics().PacketsRecv == 0 {
        return errors.New("Error checking health of controller "+c.Name+": host is unreachable")
    }

    return nil
}

// CLear sends packets to the controller to turn off all the pixels.
func (c *Controller) Clear() error {
    pixels := make([][4]int, len(c.pixels))
    for i := range pixels {
        pixels[i] = [4]int{i, 0, 0, 0}
    }
    return c.SendPackets([][][4]int{pixels})
}
