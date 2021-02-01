package cmd_release

import (
	"fmt"
	"github.com/kintohub/kinto-kube-core/pkg/types"
)

func UpdateReleaseCommitSha(kintoCoreHost, envId, blockName, releaseId, commitSha string, kintoCoreOverTls bool) error {
	commitShaRequest := &types.UpdateBuildCommitShaRequest{
		BlockName: blockName,
		EnvId:     envId,
		ReleaseId: releaseId,
		CommitSha: commitSha,
	}

	kintoCoreClient, err := newKintoCoreReleaseClient(kintoCoreHost, kintoCoreOverTls)
	if err != nil {
		return err
	}

	fmt.Printf("DEBUG Sending commit sha request info - %v\n", commitShaRequest)
	err = kintoCoreClient.updateBuildCommitSha(commitShaRequest)
	return err
}