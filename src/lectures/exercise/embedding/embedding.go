package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func (b *BandwidthUsage) bandwidthAverage() int {
	avg := 0
	for _, value := range b.amount {
		avg += int(value) / (len(b.amount))
	}
	return avg
}

func (c *CpuTemp) cpuTempAverage() float32 {
	avg := float32(0)
	for _, value := range c.temp {
		avg += float32(value) / float32(len(c.temp))
	}
	return avg
}

func (m *MemoryUsage) memoryUsageAverage() int {
	avg := 0
	for _, value := range m.amount {
		avg += int(value) / (len(m.amount))
	}
	return avg
}

func printAverageInfo(component Dashboard, title string) {
	fmt.Println()
	fmt.Println("----", title, "-----")
	fmt.Printf("Average for bandwidth usage: %v\n", component.bandwidthAverage())
	fmt.Printf("Average for cpu temperature usage: %v\n", component.cpuTempAverage())
	fmt.Printf("Average for memory usage: %v\n", component.memoryUsageAverage())
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	printAverageInfo(Dashboard{bandwidth, temp, memory}, "Dashboard components Averages")
}
