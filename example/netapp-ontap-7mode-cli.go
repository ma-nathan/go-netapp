package main

import (
	"flag"
	"fmt"
	"os"

	na "github.com/ma-nathan/go-netapp/netapp"
)

var nc *na.Client

func main() {

	server := flag.String("server", "", "NetApp Server.")
	version := flag.String("version", "1.21", "NetApp Version No.")
	user := flag.String("user", "", "NetApp user.")
	pass := flag.String("password", "", "NetApp password.")
	volume := flag.String("volume", "", "NetApp volume name.")
	snap := flag.String("snap", "", "NetApp snap name.")

	// action flags
	listFlag := flag.Bool("list", false, "List the snaps of the specified volume")
	createFlag := flag.Bool("create", false, "Create a snap of the specified volume")
	deleteFlag := flag.Bool("delete", false, "Delete a specified snap on specified volume")

	flag.Parse()

	if !(*listFlag || *createFlag || *deleteFlag) {
		fmt.Println("ERROR: You must specify a function: --list or --create or --delete.")
		os.Exit(1)
	}

	if *server == "" || *version == "" || *user == "" || *pass == "" || *volume == "" {
		fmt.Println("ERROR: You must specify server, user, password and volume.")
		os.Exit(1)
	}

	nc = na.NewClient("https://"+*server, *version,
		&na.ClientOptions{SSLVerify: false, BasicAuthUser: *user, BasicAuthPassword: *pass})

	if *listFlag {
		List(*volume)
	} else if *createFlag {
		if *snap == "" {
			fmt.Println("ERROR: You must specify a snap.")
			os.Exit(1)
		}
		Create(*volume, *snap)
	} else if *deleteFlag {
		if *snap == "" {
			fmt.Println("ERROR: You must specify a snap.")
			os.Exit(1)
		}
		Delete(*volume, *snap)
	}
	os.Exit(0)
}

func List(volume string) {
	listResponse, _, err := nc.SnapshotO7M.List(&na.SnapshotO7MOptions{Volume: volume})

  // haven't seen err return a value...
	if err != nil {
		fmt.Println("Error listing snaps on volume (%s): %v.", volume, err)
		os.Exit(1)
	}
	// fmt.Printf("DEBUG: listResponse: %+v\n", listResponse)
	if listResponse.Results.ResultBase.Status == "failed" {
		fmt.Printf("ERROR: Listing snaps on volume (%s): %s\n", volume, listResponse.Results.ResultBase.Reason)
		os.Exit(1)
	}

	if len(listResponse.Results.AttributesList.SnapshotO7MAttributes) != 0 {
		fmt.Printf("Volume: %s\n", volume)
		for _, attrib := range listResponse.Results.AttributesList.SnapshotO7MAttributes {
			// fmt.Printf("[DEBUG] SnapshotAttributes: %+v\n", attrib)
			fmt.Printf("\tsnap name: %s\n", attrib.Name)
		}
	} else {
		fmt.Printf("No snaps for %s volume: %s\n", nc.BaseURL, volume)
	}
	fmt.Printf("%s\n", listResponse.Results.ResultBase.Status)
}

func Create(volume, snapshot string) {
	fmt.Printf("Creating snapshot named \"%s\" on volume \"%s\"...\n", snapshot, volume)
	listResponse, _, err := nc.SnapshotO7M.Create(&na.SnapshotO7MOptions{Volume: volume, Snapshot: snapshot})

	if err != nil {
		fmt.Println("Error creating snap (%s) on volume (%s): %v.\n", snapshot, volume, err)
		os.Exit(1)
	}
	// Results:{ResultBase:{Status:failed Reason:Unable to find API: Params NumRecords:}}
	if listResponse.Results.ResultBase.Status == "failed" {
		fmt.Printf("ERROR: Create snap (%s) on volume (%s) failed: %s\n", snapshot, volume, listResponse.Results.ResultBase.Reason)
		os.Exit(1)
	}
	fmt.Printf("%s\n", listResponse.Results.ResultBase.Status)
}

func Delete(volume, snapshot string) {
	fmt.Printf("Deleting snapshot named \"%s\" on volume \"%s\"...\n", snapshot, volume)
	// listResponse, res, err := nc.SnapshotO7M.Delete(&na.SnapshotO7MOptions{Volume: volume, Snapshot: "test_snap", SnapshotInstanceUuid: snapshot})
	listResponse, _, err := nc.SnapshotO7M.Delete(&na.SnapshotO7MOptions{Volume: volume, Snapshot: snapshot})

	if err != nil {
		fmt.Println("Error deleting snap (%s) on volume (%s): %v.", snapshot, volume, err)
		os.Exit(1)
	}
	if listResponse.Results.ResultBase.Status == "failed" {
		fmt.Printf("ERROR: Deleting snap (%s) on volume (%s) failed: %s\n", snapshot, volume, listResponse.Results.ResultBase.Reason)
		os.Exit(1)
	}
	fmt.Printf("%s\n", listResponse.Results.ResultBase.Status)
}

/*

See /NetApp/netapp-manageability-sdk-9.3/doc/WebHelp/doc/perldoc/Ontap7ModeAPI.html

*/
