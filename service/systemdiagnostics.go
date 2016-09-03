package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type SystemDiagnostics struct {
}

/**
 * Gets the diagnostics for the system NiFi is running on
 *
 *
 * Tags: ["system-diagnostics"]
 *
 * @param nodewise Whether or not to include the breakdown per node. Optional, defaults to false
 * @param clusterNodeId The id of the node where to get the status.
 * @return model.SystemDiagnosticsEntity
 */
func (self SystemDiagnostics) GetSystemDiagnostics(nodewise bool, clusterNodeId string) model.SystemDiagnosticsEntity {
	s := api.Setup()
	res := model.SystemDiagnosticsEntity{}
	url := "http://localhost:8080/nifi-api/system-diagnostics"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
