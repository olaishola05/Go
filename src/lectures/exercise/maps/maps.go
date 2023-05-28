package main

import "fmt"

func serversToMap(servers []string, status int) map[string]int {
	serverStatusMap := make(map[string]int)
	for _, server := range servers {
		serverStatusMap[server] = status
	}
	return serverStatusMap
}

func serverStatusStats(serverMap map[string]int) map[int]int {
	serverStats := make(map[int]int)
	for _, eachServerStatus := range serverMap {
		switch eachServerStatus {
		case Online:
			serverStats[Online] += 1
		case Offline:
			serverStats[Offline] += 1
		case Maintenance:
			serverStats[Maintenance] += 1
		case Retired:
			serverStats[Retired] += 1
		default:
			panic("Some error occur")
		}
	}

	return serverStats
}

func serverStatusInfo(title string, serverMap map[string]int) {
	serverStats := serverStatusStats(serverMap)

	fmt.Println()
	fmt.Println("---", title, "---")
	fmt.Println("Total number of servers is:", len(serverMap))
	fmt.Println("Number of servers for Online", ":", serverStats[Online])
	fmt.Println("Number of servers for Offline", ":", serverStats[Offline])
	fmt.Println("Number of servers for retired", ":", serverStats[Retired])
	fmt.Println("Number of servers for Maintenance", ":", serverStats[Maintenance])
}

func updateServer(serverStatusMap map[string]int, status int, server string) map[string]int {
	serverStatusMap[server] = status
	return serverStatusMap
}

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatusMap := serversToMap(servers, Online)
	serverStatusInfo("Server Info", serverStatusMap)

	serverStatusMap = updateServer(serverStatusMap, Retired, "darkstar")
	serverStatusMap = updateServer(serverStatusMap, Offline, "aiur")
	serverStatusInfo("Server status updates", serverStatusMap)

	serverStatusMap = serversToMap(servers, Maintenance)
	serverStatusInfo("All servers on maintainance", serverStatusMap)

}
