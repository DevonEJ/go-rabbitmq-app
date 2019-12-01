package main

import (
	"flag"
	"math/rand"
	"time"
	"strconv"
)

// Capture details desired sensor from cli flags
var name = flag.String("name", "sensor", "name of sensor")
var freq = flag.Uint("freq", 5, "update frequency in cycles/sec")
var max = flag.Float64("max", 5., "maximum val for reading")
var min = flag.Float64("min", 1., "min val for reading")
var stepSize = flag.Float64("step", 0.1, "max allowable diff between readings")


func main() {

	// Parse command line flags for sensors
	flag.Parse()

	// Create input data - use nanosecond seed to ensure unique
	var ran = rand.New(rand.NewSource(time.Now().UnixNano()))

	var reading = ran.Float64()*(*max-*min) + *min


	// Create duration and signal from desired frequency of reading recording
	duration , _ := time.ParseDuration(strconv.Itoa(1000/int(*freq)) + "ms")
	signal := time.Tick(duration)

	for range signal {
		calcReading()

	}


}


func calcReading() {
	
}