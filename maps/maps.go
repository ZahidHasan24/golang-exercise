//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import (
	"fmt"
)

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//* Create a function to print server status, including:
func displayServerStatus(serverList map[string]int) {
	//  - Number of servers
	fmt.Println("\nThere are", len(serverList), "servers")
	serverStats := make(map[int]int)
	for _, status := range serverList {
		if status == Online {
			serverStats[Online] += 1
		} else if status == Offline {
			serverStats[Offline] += 1
		} else if status == Maintenance {
			serverStats[Maintenance] += 1
		} else {
			serverStats[Retired] += 1
		}
	}
	//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
	fmt.Println(serverStats[Online], "servers are online")
	fmt.Println(serverStats[Offline], "servers are offline")
	fmt.Println(serverStats[Maintenance], "servers are undergoing maintenance")
	fmt.Println(serverStats[Retired], "servers are retired")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Store the existing slice of servers in a map
	serverStore := make(map[string]int)

	//* Default all of the servers to `Online`
	for _, server := range servers {
		serverStore[server] = Online
	}

	//* Perform the following status changes and display server info:
	//  - display server info
	displayServerStatus(serverStore)
	//  - change `darkstar` to `Retired`
	serverStore["darkstar"] = Retired
	//  - change `aiur` to `Offline`
	serverStore["aiur"] = Offline
	//  - display server info
	displayServerStatus(serverStore)
	//  - change all servers to `Maintenance`
	for server := range serverStore {
		serverStore[server] = Maintenance
	}
	//  - display server info
	displayServerStatus(serverStore)
}
