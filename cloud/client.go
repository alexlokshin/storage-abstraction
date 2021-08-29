package cloud

import (
	"errors"

	"github.com/alexlokshin/storage-abstraction/cloud/awsImpl"
	"github.com/alexlokshin/storage-abstraction/cloud/azureImpl"
)

type CloudClient interface {
	UploadFile(bucket string, fileName string) error
}

func CreateClient(cloud string) (CloudClient, error) {
	switch cloud {
	case "aws":
		return awsImpl.AwsClient{}, nil
	case "azure":
		return azureImpl.AzureClient{}, nil
	default:
		return nil, errors.New("Cloud provider not supported")
	}
}
