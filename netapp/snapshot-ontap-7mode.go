package netapp

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type SnapshotO7M struct {
	Base
	Params struct {
		XMLName xml.Name
		*SnapshotO7MOptions
	}
}

type SnapshotO7MQuery struct {
	SnapshotO7MInfo *SnapshotO7MInfo `xml:"snapshot-info,omitempty"`
}

type SnapshotO7MOptions struct {
	Is7ModeSnapshot   bool   `xml:"is-7-mode-snapshot,omitempty"`
	ContainsLunClones bool   `xml:"lun-clone-snapshot,omitempty"`
	SnapOwners        bool   `xml:"snapowners,omitempty"`
	TargetName        string `xml:"target-name,omitempty"`
	TargetType        string `xml:"target-type,omitempty"`
	Terse             bool   `xml:"terse,omitempty"`
	Volume            string `xml:"volume,omitempty"`
	IgnoreOwners      bool   `xml:"ignore-owners,omitempty"`
	// delete only
	SnapshotInstanceUuid string `xml:"snapshot-instance-uuid,omitempty"`
	// delete & create
	Snapshot string `xml:"snapshot,omitempty"`
	// create only
	Async                   bool `xml:"async,omitempty"`
	IsValidLunCloneSnapshot bool `xml:"is-valid-lun-clone-snapshot,omitempty"`
}

type SnapshotO7MInfo struct {
	AccessTime                        int    `xml:"access-time,omitempty"`
	Busy                              bool   `xml:"busy,omitempty"`
	CumulativePercentageOfTotalBlocks int    `xml:"cumulative-percentage-of-total-blocks,omitempty"`
	CumulativePercentageOfUsedBlocks  int    `xml:"cumulative-percentage-of-used-blocks,omitempty"`
	CumulativeTotal                   int    `xml:"cumulative-total,omitempty"`
	Dependency                        string `xml:"dependency,omitempty"`
	FsBlockFormat                     string `xml:"fs-block-format,omitempty"`
	Name                              string `xml:"name,omitempty"`
	PercentageOfTotalBlocks           int    `xml:"percentage-of-total-blocks,omitempty"`
	PercentageOfUsedBlocks            int    `xml:"percentage-of-used-blocks,omitempty"`
	SnapshotInstanceUuid              string `xml:"snapshot-instance-uuid,omitempty"`
	SnapshotVersionUuid               string `xml:"snapshot-version-uuid,omitempty"`
	Total                             int    `xml:"total,omitempty"`
}

type SnapshotO7MListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			SnapshotO7MAttributes []SnapshotO7MInfo `xml:"snapshot-info"`
		} `xml:"snapshots"`
	} `xml:"results"`
}

type SnapshotInfoO7MResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		SnapshotO7MInfo `xml:"snapshot-info"`
	} `xml:"results"`
}

type Ontap7ModeVolumeListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			VolumeAttributes []Ontap7ModeVolumeInfo `xml:"volume-info"`
		} `xml:"volumes"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

type PlexInfo struct {
	Name        string `xml:"name"`
	IsOnline    bool   `xml:"is-online"`
	ResyncLevel int8   `xml:"resync-level"`
	Pool        int8   `xml:"pool"`
	PlexStatus  string `xml:"plex-status"`
}

type Ontap7ModeVolumeInfo struct {
	Name                     string `xml:"name"`
	Uuid                     string `xml:"uuid"`
	Type                     string `xml:"type"`
	BlockType                string `xml:"block-type"`
	State                    string `xml:"state"`
	FilesystemSize           int64  `xml:"filesystem-size"`
	SizeTotal                int64  `xml:"size-total"`
	SizeUsed                 int64  `xml:"size-used"`
	SizeAvailable            int64  `xml:"size-available"`
	PercentageUsed           int8   `xml:"percentage-used"`
	SnapshotPercentReserved  int8   `xml:"snapshot-percent-reserved"`
	SnapshotBlocksReserved   int64  `xml:"snapshot-blocks-reserved"`
	ReserveRequired          int    `xml:"reserve-required"`
	Reserve                  int    `xml:"reserve"`
	ReserveUsed              int    `xml:"reserve-used"`
	ReserveUsedActual        int    `xml:"reserve-used-actual"`
	FilesTotal               int64  `xml:"files-total"`
	FilesUsed                int64  `xml:"files-used"`
	FilesPrivateUsed         int64  `xml:"files-private-used"`
	InodefilePublicCapacity  int64  `xml:"inodefile-public-capacity"`
	InodefilePrivateCapacity int64  `xml:"inodefile-private-capacity"`
	QuotaInit                int64  `xml:"quota-init"`
	IsSnaplock               bool   `xml:"is-snaplock"`
	ContainingAggregate      string `xml:"containing-aggregate"`
	CompressionInfo          struct {
		IsCompressionEnabled bool `xml:"is-compression-enabled"`
	} `xml:compression-info"`
	InstanceUuid            string `xml:"instance-uuid"`
	SpaceReserve            string `xml:"space-reserve"`
	SpaceReserveEnabled     bool   `xml:"space-reserve-enabled"`
	ProvenanceUuid          string `xml:"provenance-uuid"`
	RaidSize                int64  `xml:"raid-size"`
	RaidStatus              string `xml:"raid-status"`
	ChecksumStyle           string `xml:"checksum-style"`
	IsChecksumEnabled       bool   `xml:"is-checksum-enabled"`
	IsInconsistent          bool   `xml:"is-inconsistent"`
	IsUnrecoverable         bool   `xml:"is-unrecoverable"`
	IsInvalid               bool   `xml:"is-invalid"`
	IsInSnapmirrorJumpahead bool   `xml:"is-in-snapmirror-jumpahead"`
	MirrorStatus            string `xml:"mirror-status"`
	DiskCount               int64  `xml:"disk-count"`
	PlexCount               int64  `xml:"plex-count"`
	Plexes                  struct {
		PlexInfo []PlexInfo `xml:"plex-info"`
	} `xml:"plexes"`
	Volume64bitUpgrade struct {
		Volume64bitUpgradeInfo string `xml:"volume-64bit-upgrade-info"`
	} `xml:"volume-64bit-upgrade"`
}

type CreateSnapshotO7MListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
	} `xml:"results"`
}

type DeleteSnapshotO7MListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
	} `xml:"results"`
}

func (v *SnapshotO7M) Create(options *SnapshotO7MOptions) (*CreateSnapshotO7MListResponse, *http.Response, error) {
	v.Params.XMLName = xml.Name{Local: "snapshot-create"}
	v.Params.SnapshotO7MOptions = options
	r := CreateSnapshotO7MListResponse{}
	res, err := v.get(v, &r)
	return &r, res, err
}

func (v *SnapshotO7M) VolumeInfo(options *SnapshotO7MOptions) (*Ontap7ModeVolumeListResponse, *http.Response, error) {
	v.Params.XMLName = xml.Name{Local: "volume-list-info"}
	v.Params.SnapshotO7MOptions = options
	r := Ontap7ModeVolumeListResponse{}
	res, err := v.get(v, &r)
	return &r, res, err
}

func (v *SnapshotO7M) Delete(options *SnapshotO7MOptions) (*DeleteSnapshotO7MListResponse, *http.Response, error) {
	v.Params.XMLName = xml.Name{Local: "snapshot-delete"}
	v.Params.SnapshotO7MOptions = options
	r := DeleteSnapshotO7MListResponse{}
	res, err := v.get(v, &r)
	return &r, res, err
}

// snapshot-get-iter
// AKA snapshot info
// * is-7-mode-snapshot
// * lun-clone-snapshot
// * snapowners
// * target-name
// * target-type
// * terse
// * volume
func (v *SnapshotO7M) SnapshotInfo(options *SnapshotO7MOptions) (*SnapshotInfoO7MResponse, *http.Response, error) {
	v.Params.XMLName = xml.Name{Local: "snapshot-get-iter"}
	v.Params.SnapshotO7MOptions = options
	r := SnapshotInfoO7MResponse{}
	res, err := v.get(v, &r)
	if err != nil {
		fmt.Printf("get returned error: %v\n", err)
	}
	return &r, res, err
}

func (v *SnapshotO7M) List(options *SnapshotO7MOptions) (*SnapshotO7MListResponse, *http.Response, error) {
	v.Params.XMLName = xml.Name{Local: "snapshot-list-info"}
	v.Params.SnapshotO7MOptions = options
	r := SnapshotO7MListResponse{}
	res, err := v.get(v, &r)
	if err != nil {
		fmt.Printf("get returned error: %v\n", err)
	}
	return &r, res, err
}
