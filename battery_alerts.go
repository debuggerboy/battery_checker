package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

    // Battery Alert Cutoff Percentage
    var batt_cutoff_percent = 20

	// Open the battery information files
	statusFile, err := os.Open("/sys/class/power_supply/BAT0/status")
	if err != nil {
		fmt.Println("Error opening battery status file:", err)
		return
	}
	defer statusFile.Close()

	capacityFile, err := os.Open("/sys/class/power_supply/BAT0/capacity")
	if err != nil {
		fmt.Println("Error opening battery capacity file:", err)
		return
	}
	defer capacityFile.Close()

	// Read the battery status and capacity
	statusScanner := bufio.NewScanner(statusFile)
	statusScanner.Scan()
	batteryStatus := strings.TrimSpace(statusScanner.Text())

	capacityScanner := bufio.NewScanner(capacityFile)
	capacityScanner.Scan()
	batteryCapacity, err := strconv.Atoi(strings.TrimSpace(capacityScanner.Text()))
	if err != nil {
		fmt.Println("Error parsing battery capacity:", err)
		return
	}

	// Check if the battery is discharging and below 20%
	if batteryStatus == "Discharging" && batteryCapacity < batt_cutoff_percent {
		fmt.Println("The battery is discharging and below", batt_cutoff_percent, "%.")
    } else if batteryStatus == "Discharging" && batteryCapacity > batt_cutoff_percent {
        fmt.Println("The battery is discharging but above", batt_cutoff_percent, "%.")
	} else {
		fmt.Println("The battery is charging.")
	}
}
