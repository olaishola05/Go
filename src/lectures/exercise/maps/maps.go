package main

import "fmt"

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

func serverStatusInfo(title string, servers []string, serverMap map[string]int) {
	// totalServers := 0
	// var serverStats map[int]int

	// for _, server := range servers {
	// 	totalServers++
	// 	serverMap[server] = Online
	// }
	serverStats := serverStatusStats(serverMap)

	fmt.Println()
	fmt.Println("---", title, "---")
	fmt.Println("Total number of server is:", len(serverMap))
	fmt.Println("Number of servers for Online", ":", serverStats[Online])
	fmt.Println("Number of servers for Offline", ":", serverStats[Offline])
	fmt.Println("Number of servers for Maintenance", ":", serverStats[Maintenance])
	fmt.Println("Number of servers for retired", ":", serverStats[Retired])

}

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatusMap := make(map[string]int)

	for _, server := range servers {
		// totalServers++
		serverStatusMap[server] = Online
		// serverStats = serverStatusStats(serverMap)
	}

	serverStatusInfo("Server Info", servers, serverStatusMap)

	serverStatusMap["darkstar"] = Retired
	serverStatusMap["aiur"] = Offline
	serverStatusInfo("Server status updates", servers, serverStatusMap)

	for server, _ := range serverStatusMap {
		serverStatusMap[server] = Maintenance
	}

	serverStatusInfo("New server update info", servers, serverStatusMap)

}
