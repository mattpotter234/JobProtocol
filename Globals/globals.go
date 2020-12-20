package globals

const (
	// ConnAddr is a string constant specifying where the job creator connects to.
	ConnAddr = "localhost"
	// ConnPort is a string constant specifying the port the job creator runs on
	ConnPort = 7777
	// ConnType specifies the connection type for the job creator (usually just TCP)
	ConnType = "tcp"
)

// GetJobNames is an array of job names for the C&C to use
func GetJobNames() []string {
	return []string{
		"EQN",
		"HOSTUP",
		"TCPFLOOD",
		"UDPFLOOD",
		"NEIGHBOURS",
	}
}

// GetTCPPorts port is an array of TCP ports to be used in the TCP flood
func GetTCPPorts() []int {
	return []int{
		25,
		80,
		443,
		20,
		21,
		23,
		143,
		3389,
		22,
		53,
		110,
	}
}

// GetUDPPorts is an array of UDP ports to be used in the UDP flood
func GetUDPPorts() []int {
	return []int{
		53,
		67,
		68,
		69,
		123,
		137,
		138,
		139,
		161,
		162,
		389,
		636,
	}
}
