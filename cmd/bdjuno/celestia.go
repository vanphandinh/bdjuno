package main

import (
	"encoding/json"
	blobtypes "github.com/celestiaorg/celestia-app/x/blob/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	junomessages "github.com/forbole/juno/v4/modules/messages"
	"strconv"
)

// celestiaMessageAddressesParser represents a parser able to get the addresses of the involved
// account from a Celestia message
var celestiaMessageAddressesParser = junomessages.JoinMessageParsers(
	blobMessageAddressesParser,
)

// blobMessageAddressesParser represents a MessageAddressesParser for the x/blob module
func blobMessageAddressesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {

	case *blobtypes.MsgPayForBlobs:
		{
			// msg.NamespaceIds
			namespaceIdsArray := make([]string, 0)
			namespaceIds := ""
			for _, e := range msg.NamespaceIds {
				namespaceIdsArray = append(namespaceIdsArray, string(e))
			}
			j, err := json.Marshal(namespaceIdsArray)
			if err == nil {
				namespaceIds = string(j)
			}

			// msg.ShareCommitments
			shareCommitmentsArray := make([]string, 0)
			shareCommitments := ""
			for _, e := range msg.ShareCommitments {
				shareCommitmentsArray = append(shareCommitmentsArray, string(e))
			}
			j, err = json.Marshal(shareCommitmentsArray)
			if err == nil {
				shareCommitments = string(j)
			}

			// msg.BlobSizes
			blobSizesArray := make([]string, 0)
			blobSizes := ""
			for _, e := range msg.BlobSizes {
				blobSizesArray = append(blobSizesArray, strconv.FormatUint(uint64(e), 10))
			}
			j, err = json.Marshal(blobSizesArray)
			if err == nil {
				blobSizes = string(j)
			}

			// msg.ShareVersions
			shareVersionsArray := make([]string, 0)
			shareVersions := ""
			for _, e := range msg.ShareVersions {
				shareVersionsArray = append(shareVersionsArray, strconv.FormatUint(uint64(e), 10))
			}
			j, err = json.Marshal(shareVersionsArray)
			if err == nil {
				shareVersions = string(j)
			}

			return []string{msg.Signer, namespaceIds, shareCommitments, blobSizes, shareVersions}, nil
		}

	}
	return nil, junomessages.MessageNotSupported(cosmosMsg)
}
