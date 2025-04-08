package controllers

import (
	"context"
	"fmt"

	operatorsv1 "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/apis/operators/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func GetLatestCSVFromCatalog(ctx context.Context, c client.Client, operatorName, channelName, catalogNamespace string) (string, error) {
	logger := log.FromContext(ctx)

	var pkg operatorsv1.PackageManifest
	if err := c.Get(ctx, client.ObjectKey{
		Name:      operatorName,
		Namespace: catalogNamespace,
	}, &pkg); err != nil {
		logger.Error(err, "unable to get PackageManifest", "name", operatorName)
		return "", fmt.Errorf("PackageManifest not found: %w", err)
	}

	for _, ch := range pkg.Status.Channels {
		if ch.Name == channelName {
			return ch.CurrentCSV, nil
		}
	}

	return "", fmt.Errorf("channel %s not found in PackageManifest %s", channelName, operatorName)
}
