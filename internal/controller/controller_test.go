package controller

import (
    "testing"
)

func TestConnect(t *testing.T) {
    c := Controller{IPAdress: "192.168.0.10"}
    err := c.Connect()
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if c.Conn == nil {
        t.Errorf("Expected to have a connection, got %v", c.Conn)
    }
    c = Controller{IPAdress: "oui"}
    err = c.Connect()
    if err == nil {
        t.Errorf("Expected an error, got nil")
    }
}

func TestInitPixelsArrays(t *testing.T) {
    c := Controller{}
    c.InitPixelsArrays(10)
    if len(c.pixels) != 10 {
        t.Errorf("Expected to have 10 pixels, got %d", len(c.pixels))
    }
}

func TestSendPackets(t *testing.T) {
    c := Controller{IPAdress: "127.0.0.1"}
    c.Connect()
    packets := [][][4]int{{{1, 2, 3, 4}, {2, 2, 3, 4}, {3, 2, 3, 4}, {4, 2, 3, 4}}}
    err := c.SendPackets(packets)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
}
