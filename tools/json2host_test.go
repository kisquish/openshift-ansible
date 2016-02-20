package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	foo := `{
    "msg": {
        "lenaic": {
            "master": {
                "default": [
                    {
                        "name": "lenaic-master-6f732",
                        "private IP": "192.168.55.159",
                        "public IP": "192.168.55.159"
                    }
                ]
            },
            "node": {
                "compute": [
                    {
                        "name": "lenaic-node-compute-e05f2",
                        "private IP": "192.168.55.55",
                        "public IP": "192.168.55.55"
                    },
                    {
                        "name": "lenaic-node-compute-82447",
                        "private IP": "192.168.55.214",
                        "public IP": "192.168.55.214"
                    }
                ],
                "infra": [
                    {
                        "name": "lenaic-node-infra-24671",
                        "private IP": "192.168.55.34",
                        "public IP": "192.168.55.34"
                    }
                ]
            }
        }
    }
}`
	New(foo)
}

func TestCreateJSon(t *testing.T) {
	nodeMaster := node{"lenaic-master", "127.0.0.1", "127.0.0.1"}
	node1 := node{"name1", "127.0.0.1", "127.0.0.1"}
	node2 := node{"name2", "127.0.0.2", "127.0.0.2"}
	var nodeList1 nodeList
	nodeList1 = nodeList{node1, node2}
	var nodeListMaster nodeList
	nodeListMaster = nodeList{nodeMaster}

	nodeGroup1 := make(nodeGroup)
	nodeGroup1["foo"] = nodeList1

	nodeGroupMaster := make(nodeGroup)
	nodeGroupMaster["default"] = nodeListMaster

	cluster1 := make(cluster)
	cluster1["master"] = nodeGroupMaster
	cluster1["node"] = nodeGroup1

	msg := make(map[string]cluster)
	msg["lenaic"] = cluster1
	rsp := AnsibleResponse{msg}

	bolB, _ := json.Marshal(rsp)
	fmt.Println(string(bolB))

	var nodeTest nodeList
	json.Unmarshal(bolB, &nodeTest)
}
