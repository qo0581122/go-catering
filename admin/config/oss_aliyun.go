package config

type AliyunOSS struct {
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `json:"accessKeyId" yaml:"access-key-id"`
	AccessKeySecret string `json:"accessKeySecret" yaml:"access-key-secret"`
	BucketName      string `json:"bucketName" yaml:"bucket-name"`
	BucketUrl       string `json:"bucketUrl" yaml:"bucket-url"`
	BasePath        string `json:"basePath" yaml:"base-path"`
}
