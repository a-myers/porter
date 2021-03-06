package input

import (
	"encoding/json"
)

type EKS struct {
	AWSRegion    string `json:"aws_region"`
	AWSAccessKey string `json:"aws_access_key"`
	AWSSecretKey string `json:"aws_secret_key"`
	ClusterName  string `json:"cluster_name"`
}

func (eks *EKS) GetInput() ([]byte, error) {
	return json.Marshal(eks)
}

func GetEKSInput(bytes []byte) (*EKS, error) {
	res := &EKS{}

	err := json.Unmarshal(bytes, res)

	return res, err
}
