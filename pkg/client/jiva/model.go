package client

import (
	"net/http"
	"strings"
)

const (
	metadataSuffix     = ".meta"
	imgSuffix          = ".img"
	volumeMetaData     = "volume.meta"
	defaultSectorSize  = 4096
	headPrefix         = "volume-head-"
	headSuffix         = ".img"
	headName           = headPrefix + "%03d" + headSuffix
	diskPrefix         = "volume-snap-"
	diskSuffix         = ".img"
	diskName           = diskPrefix + "%s" + diskSuffix
	maximumChainLength = 250
	snapPrefix         = "volume-snap-"
	snapSuffix         = ".img"
)

// Volumes 
type Volumes struct {
	Resource
	Name         string `json:"name"`
	ReplicaCount int    `json:"replicaCount"`
	Endpoint     string `json:"endpoint"`
}
// VolumeCollection contains Collection, Volumes
type VolumeCollection struct {
	Collection
	Data []Volumes `json:"data"`
}
// Replica contains a copy of Address, Mode
type Replica struct {
	Resource
	Address string `json:"address"`
	Mode    string `json:"mode"`
}
// InfoReplica contains a copy of info
type InfoReplica struct {
	Resource
	Dirty           bool                `json:"dirty"`
	Rebuilding      bool                `json:"rebuilding"`
	Head            string              `json:"head"`
	Parent          string              `json:"parent"`
	Size            string              `json:"size"`
	SectorSize      int64               `json:"sectorSize"`
	State           string              `json:"state"`
	Chain           []string            `json:"chain"`
	Disks           map[string]DiskInfo `json:"disks"`
	RemainSnapshots int                 `json:"remainsnapshots"`
	RevisionCounter string              `json:"revisioncounter"`
}
// DiskInfo contains info related to the disk
type DiskInfo struct {
	Name        string   `json:"name"`
	Parent      string   `json:"parent"`
	Children    []string `json:"children"`
	Removed     bool     `json:"removed"`
	UserCreated bool     `json:"usercreated"`
	Created     string   `json:"created"`
	Size        string   `json:"size"`
}
// ReplicaCollection copies Collection into Data
type ReplicaCollection struct {
	Collection
	Data []Replica `json:"data"`
}
// MarkDiskAsRemovedInput contains Resource 'Name'
type MarkDiskAsRemovedInput struct {
	Resource
	Name string `json:"name"`
}

// ReplicaClient is Client structure
type ReplicaClient struct {
	Address    string
	SyncAgent  string
	Host       string
	httpClient *http.Client
}
// ControllerClient contains Address, Host, httpClient
type ControllerClient struct {
	Address    string
	Host       string
	httpClient *http.Client
}
// RevertInput returns back Name
type RevertInput struct {
	Resource
	Name string `json:"name"`
}

// SnapshotInput is Input struct to create
// snapshot
type SnapshotInput struct {
	Resource
	Name        string            `json:"name"`
	UserCreated bool              `json:"usercreated"`
	Created     string            `json:"created"`
	Labels      map[string]string `json:"labels"`
}

type SnapshotOutput struct {
	Resource
}

type Resource struct {
	Id      string            `json:"id,omitempty"`
	Type    string            `json:"type,omitempty"`
	Links   map[string]string `json:"links"`
	Actions map[string]string `json:"actions"`
}
// Collection structure, used to contain info about the container 
type Collection struct {
	Type         string                 `json:"type,omitempty"`
	ResourceType string                 `json:"resourceType,omitempty"`
	Links        map[string]string      `json:"links,omitempty"`
	CreateTypes  map[string]string      `json:"createTypes,omitempty"`
	Actions      map[string]string      `json:"actions,omitempty"`
	SortLinks    map[string]string      `json:"sortLinks,omitempty"`
	Pagination   *Pagination            `json:"pagination,omitempty"`
	Sort         *Sort                  `json:"sort,omitempty"`
	Filters      map[string][]Condition `json:"filters,omitempty"`
}

// Sort used to sort based on Name, Order 
type Sort struct {
	Name    string `json:"name,omitempty"`
	Order   string `json:"order,omitempty"`
	Reverse string `json:"reverse,omitempty"`
}

// Condition structure
type Condition struct {
	Modifier string      `json:"modifier,omitempty"`
	Value    interface{} `json:"value,omitempty"`
}
// Pagination contains paths to neighbour nodes
type Pagination struct {
	Marker   string `json:"marker,omitempty"`
	First    string `json:"first,omitempty"`
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	Total    *int64 `json:"total,omitempty"`
	Partial  bool   `json:"partial,omitempty"`
}

func Filter(list []string, check func(string) bool) []string {
	result := make([]string, 0, len(list))
	for _, i := range list {
		if check(i) {
			result = append(result, i)
		}
	}
	return result
}

// Contains returns True if length of string array == value
func Contains(arr []string, val string) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

// IsHeadDisk checks if Disk is a Head
func IsHeadDisk(diskName string) bool {
	if strings.HasPrefix(diskName, headPrefix) && strings.HasSuffix(diskName, headSuffix) {
		return true
	}
	return false
}

// TrimSnapshotName trims the snapshot name from diskName string
func TrimSnapshotName(diskName string) string {
	if diskName == "" {
		return "NA"
	} else if strings.HasPrefix(diskName, headPrefix) && strings.HasSuffix(diskName, headSuffix) {
		return strings.TrimSuffix(strings.TrimPrefix(diskName, "volume-"), ".img")
	} else if strings.HasPrefix(diskName, snapPrefix) && strings.HasSuffix(diskName, snapSuffix) {
		return strings.TrimSuffix(strings.TrimPrefix(diskName, "volume-snap-"), ".img")
	}

	return "NA"
}

// TrimSnapshotNamesOfSlice trims the each snapshot name in diskNames Slice
func TrimSnapshotNamesOfSlice(diskNames []string) []string {
	if len(diskNames) == 0 {
		return []string{"NA"}
	}

	snapChildren := make([]string, len(diskNames))
	for index, diskName := range diskNames {
		snapChildren[index] = TrimSnapshotName(diskName)
	}
	return snapChildren
}
