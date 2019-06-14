package main

import (
    "fmt"
    "log"

    "periph.io/x/periph/conn/gpio"
    "periph.io/x/periph/conn/gpio/gpioreg"
    "periph.io/x/periph/host"
)

func main() {
    // Load all the drivers:
    if _, err := host.Init(); err != nil {
        log.Fatal(err)
    }

    // Lookup a pin by its number:
    p := gpioreg.ByName("4")
    // p, err := gpioreg.ByName("16")
    // if err != nil {
    //     log.Fatal(err)
    // }
    fmt.Printf("%s: %s\n", p, p.Function())

    // Set it as input.
    // if err = p.In(gpio.PullNoChange, gpio.RisingEdge); err != nil {
    //     log.Fatal(err)
    // }

    // Wait for edges as detected by the hardware.
    for {
        p.WaitForEdge(-1)
        if p.Read() == gpio.High {
          fmt.Printf("You moved!\n")
        }
    }
}