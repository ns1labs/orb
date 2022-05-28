package grpc

type agentRes struct {
	id      string
	name    string
	channel string
}

type agentGroupRes struct {
	id      string
	name    string
	channel string
}

type agentInfoRes struct {
	ownerID   string
	agentName string
	agentTags map[string]string
}

type emptyRes struct {
	err error
}
