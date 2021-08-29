package cloud

type CloudClient interface {
	UploadFile(bucket string, fileName string) error
}
