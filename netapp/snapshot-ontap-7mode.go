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
	IgnoreOwners         bool   `xml:"ignore-owners,omitempty"`
  // delete only
	SnapshotInstanceUuid string `xml:"snapshot-instance-uuid,omitempty"`
  // delete & create
	Snapshot             string `xml:"snapshot,omitempty"`
  // create only
	Async                   bool   `xml:"async,omitempty"`
	IsValidLunCloneSnapshot bool   `xml:"is-valid-lun-clone-snapshot,omitempty"`
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

func (v *SnapshotO7M) Delete(options *SnapshotO7MOptions) (*DeleteSnapshotO7MListResponse, *http.Response, error) {
	v.Params.XMLName = xml.Name{Local: "snapshot-delete"}
	v.Params.SnapshotO7MOptions = options
	r := DeleteSnapshotO7MListResponse{}
	res, err := v.get(v, &r)
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
