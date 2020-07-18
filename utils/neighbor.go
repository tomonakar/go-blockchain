package utils

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"time"
)

func IsFoundHost(host string, port uint16) bool {
	target := fmt.Sprintf("%s:%d", host, port)
	// 指定したhostにコネクションできたか確認する
	_, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err != nil {

		fmt.Printf("%s %v\n", target, err)
		return false
	}
	return true
}

var PATTERN = regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\.){3})(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

func FindNeighors(myHost string, myPort uint16, startIp uint8, endIp uint8, startPort uint16, endPort uint16) []string {
	address := fmt.Sprintf("%s:%d", myHost, myPort)

	// ex) 127.0.0.1 を FindStringSubmatchした結果
	//[127.0.0.1, 127.0.0., 0., 1]
	m := PATTERN.FindStringSubmatch(myHost)
	if m == nil {
		return nil
	}
	prefixHost := m[1]
	fmt.Printf("%v", m)
	lastIp, _ := strconv.Atoi(m[len(m)-1])

	neighbors := make([]string, 0)

	for port := startPort; port <= endPort; port += 1 {
		for ip := startIp; ip <= endIp; ip += 1 {
			guessHost := fmt.Sprintf("%s%d", prefixHost, lastIp+int(ip))
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port)
			if guessTarget != address && IsFoundHost(guessHost, port) {
				neighbors = append(neighbors, guessTarget)
			}
		}
	}
	return neighbors
}
