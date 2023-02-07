package models

type StorageType string
type ShardNum int

type ZincIndex struct {
	Name        string      `json:"name"`
	StorageType StorageType `json:"storage_type"`
	ShardNum    ShardNum    `json:"shard_num"`
	Mapping     interface{} `json:"mappings"`
}
