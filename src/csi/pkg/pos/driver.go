/*
Copyright (c) Arm Limited and Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pos

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"k8s.io/klog"

	csicommon "github.com/poseidonos/pos-csi/pkg/csi-common"
	"github.com/poseidonos/pos-csi/pkg/util"
)

func Run(conf *util.Config) {
	var (
		cd  *csicommon.CSIDriver
		ids *identityServer
		cs  *controllerServer
		ns  *nodeServer

		controllerCaps = []csi.ControllerServiceCapability_RPC_Type{
                        csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME, csi.ControllerServiceCapability_RPC_LIST_VOLUMES, csi.ControllerServiceCapability_RPC_GET_VOLUME, csi.ControllerServiceCapability_RPC_VOLUME_CONDITION,
		}
		volumeModes = []csi.VolumeCapability_AccessMode_Mode{
			csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
		}
	)

	cd = csicommon.NewCSIDriver(conf.DriverName, conf.DriverVersion, conf.NodeID)
	if cd == nil {
		klog.Fatalln("Failed to initialize CSI Driver.")
	}
	if conf.IsControllerServer {
		cd.AddControllerServiceCapabilities(controllerCaps)
		cd.AddVolumeCapabilityAccessModes(volumeModes)
	}

	ids = newIdentityServer(cd)

	if conf.IsNodeServer {
		ns = newNodeServer(cd)
	}

	if conf.IsControllerServer {
		var err error
		cs, err = newControllerServer(cd)
		if err != nil {
			klog.Fatalf("failed to create controller server: %s", err)
		}
	}

	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(conf.Endpoint, ids, cs, ns)
	s.Wait()
}
