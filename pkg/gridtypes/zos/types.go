package zos

import (
	"fmt"

	"github.com/threefoldtech/zos/pkg/gridtypes"
)

const (
	// ContainerType type
	ContainerType gridtypes.WorkloadType = "container"
	// VolumeType type
	VolumeType gridtypes.WorkloadType = "volume"
	// NetworkType type
	NetworkType gridtypes.WorkloadType = "network"
	// ZDBType type
	ZDBType gridtypes.WorkloadType = "zdb"
	// KubernetesType type
	KubernetesType gridtypes.WorkloadType = "kubernetes"
	//PublicIPType reservation
	PublicIPType gridtypes.WorkloadType = "ipv4"
)

func init() {
	gridtypes.RegisterType(ContainerType, Container{})
	gridtypes.RegisterType(VolumeType, Volume{})
	gridtypes.RegisterType(NetworkType, Network{})
	gridtypes.RegisterType(ZDBType, ZDB{})
	gridtypes.RegisterType(KubernetesType, Kubernetes{})
	gridtypes.RegisterType(PublicIPType, PublicIP{})
}

// DeviceType is the actual type of hardware that the storage device runs on,
// i.e. SSD or HDD
type DeviceType string

// Known device types
const (
	SSDDevice DeviceType = "ssd"
	HDDDevice DeviceType = "hdd"
)

func (d DeviceType) String() string {
	return string(d)
}

// Valid validates device type
func (d DeviceType) Valid() error {
	if d != SSDDevice && d != HDDDevice {
		return fmt.Errorf("invalid device type")
	}
	return nil
}