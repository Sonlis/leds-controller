package controller

import (
	"errors"
	"net"
	"time"
	"github.com/prometheus-community/pro-bing"
)

type Controller struct {
    IPAdress string `json:"ip_adress"`
    Port int `json:"port"`
    Name string `json:"name"`
    pixels [][4]int
    net.Conn
}

func (c *Controller) Connect() error {
    conn, err := net.Dial("udp", c.IPAdress+":7777")
    if err != nil {
        return errors.New("Error connecting to controller "+c.Name+": "+err.Error())
    }
    c.Conn = conn
    return nil
}

func (c *Controller) InitPixelsArrays(pixelCount int) {
    c.pixels = make([][4]int, pixelCount)
}

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
