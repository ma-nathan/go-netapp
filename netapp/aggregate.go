package netapp

import (
	"encoding/xml"
	"net/http"
)

type Aggregate struct {
	Base
	Params struct {
		XMLName xml.Name
		*AggrOptions
	}
}

type AggrOptions struct {
	DesiredAttributes *AggrInfo `xml:"desired-attributes,omitempty"`
	MaxRecords        int       `xml:"max-records,omitempty"`
	Query             *AggrInfo `xml:"query,omitempty"`
	Tag               string    `xml:"tag,omitempty"`
}

type AggrListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			AggrAttributes []AggrInfo `xml:"aggr-attributes"`
		} `xml:"attributes-list"`
	} `xml:"results"`
}

func (a *Aggregate) List(options *AggrOptions) (*AggrListResponse, *http.Response, error) {
	a.Params.XMLName = xml.Name{Local: "aggr-get-iter"}
	a.Params.AggrOptions = options
	r := AggrListResponse{}
	res, err := a.get(a, &r)
	return &r, res, err
}

type AggrInfo struct {
	AggregateName       string              `xml:"aggregate-name"`
	AggrInodeAttributes AggrInodeAttributes `xml:"aggr-inode-attributes"`
	AggrSpaceAttributes AggrSpaceAttributes `xml:"aggr-space-attributes"`
}

type AggrInodeAttributes struct {
	FilesPrivateUsed         string `xml:"files-private-used"`
	FilesTotal               string `xml:"files-total"`
	FilesUsed                string `xml:"files-used"`
	InodefilePrivateCapacity string `xml:"inodefile-private-capacity"`
	InodefilePublicCapacity  string `xml:"inodefile-public-capacity"`
	MaxfilesAvailable        string `xml:"maxfiles-available"`
	MaxfilesPossible         string `xml:"maxfiles-possible"`
	MaxfilesUsed             string `xml:"maxfiles-used"`
	PercentInodeUsedCapacity string `xml:"percent-inode-used-capacity"`
}

type AggrSpaceAttributes struct {
	AggregateMetadata            string `xml:"aggregate-metadata"`
	HybridCacheSizeTotal         string `xml:"hybrid-cache-size-total"`
	PercentUsedCapacity          string `xml:"percent-used-capacity"`
	PhysicalUsed                 string `xml:"physical-used"`
	PhysicalUsedPercent          string `xml:"physical-used-percent"`
	SizeAvailable                string `xml:"size-available"`
	SizeTotal                    string `xml:"size-total"`
	SizeUsed                     string `xml:"size-used"`
	TotalReservedSpace           string `xml:"total-reserved-space"`
	UsedIncludingSnapshotReserve string `xml:"used-including-snapshot-reserve"`
	VolumeFootprints             string `xml:"volume-footprints"`
}

type AggregateSpace struct {
	Base
	Params struct {
		XMLName xml.Name
		*AggrSpaceOptions
	}
}

type AggrSpaceOptions struct {
	DesiredAttributes *AggrSpaceInfo `xml:"desired-attributes,omitempty"`
	MaxRecords        int            `xml:"max-records,omitempty"`
	Query             *AggrSpaceInfo `xml:"query,omitempty"`
	Tag               string         `xml:"tag,omitempty"`
}

type AggrSpaceInfo struct {
	Aggregate                           string `xml:"aggregate"`
	AggregateMetadata                   string `xml:"aggregate-metadata"`
	AggregateMetadataPercent            string `xml:"aggregate-metadata-percent"`
	AggregateSize                       string `xml:"aggregate-size"`
	PercentSnapshotSpace                string `xml:"percent-snapshot-space"`
	PhysicalUsed                        string `xml:"physical-used"`
	PhysicalUsedPercent                 string `xml:"physical-used-percent"`
	SnapSizeTotal                       string `xml:"snap-size-total"`
	SnapshotReserveUnusable             string `xml:"snapshot-reserve-unusable"`
	SnapshotReserveUnusablePercent      string `xml:"snapshot-reserve-unusable-percent"`
	UsedIncludingSnapshotReserve        string `xml:"used-including-snapshot-reserve"`
	UsedIncludingSnapshotReservePercent string `xml:"used-including-snapshot-reserve-percent"`
	VolumeFootprints                    string `xml:"volume-footprints"`
	VolumeFootprintsPercent             string `xml:"volume-footprints-percent"`
}

type AggrSpaceListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			AggrAttributes []AggrSpaceInfo `xml:"space-information"`
		} `xml:"attributes-list"`
	} `xml:"results"`
}

func (a *AggregateSpace) List(options *AggrSpaceOptions) (*AggrSpaceListResponse, *http.Response, error) {
	a.Params.XMLName = xml.Name{Local: "aggr-space-get-iter"}
	a.Params.AggrSpaceOptions = options
	r := AggrSpaceListResponse{}
	res, err := a.get(a, &r)
	return &r, res, err
}
