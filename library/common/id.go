package common

import "github.com/bwmarrin/snowflake"

// GenerateUUID 采用雪花算法生成id
func GenerateUUID() (snowflake.ID, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}
	return node.Generate(), nil
}
