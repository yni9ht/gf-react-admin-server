package common

import "github.com/bwmarrin/snowflake"

type idUtils struct{}

var (
	Id = &idUtils{}
)

// GenerateUUID 采用雪花算法生成id
func (i *idUtils) GenerateUUID() (int64, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}
	return int64(node.Generate()), nil
}
