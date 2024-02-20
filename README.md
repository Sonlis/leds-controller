# Leds-controller

WIP

Inspired by [naztronaut dancy leds](https://github.com/naztronaut/dancyPi-audio-reactive-led).
Sets up a webserver which listens for JSON packet to display an effect on ws2812b leds. The webserver
computes what to display, and then sends it over wifi to an ESP32 or ESP8266 physically connected to
the leds, which displays the desired pattern.

## Why this repository

The work of naztronaut and scottlawson is awsome. But 2 main things pushed the creation of this repository:
1. It does not include a webserver, so triggering the patterns must be done manually on the machine
computing the pattern. Hence it is impossible to trigger patterns from a remote, low-computing device.
2. Being written in python, the original work is more complicated to setup, require dependencies in python
and perhaps and package manager such as conda. This repository only contains a binary, less hassle to setup.
