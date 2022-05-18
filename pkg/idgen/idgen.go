package idgen

import (
	"github.com/zmicro-team/zim/pkg/idgen/snowflake"
)

var sf *snowflake.Snowflake

func init() {
	st := snowflake.Settings{
		MachineID: getMachineId,
	}
	sf = snowflake.NewSnowflake(st)
}

func getMachineId() (uint16, error) {
	// TODO
	return 1, nil
}

func Next() int64 {
	return sf.NextId()
}

func GetOne() int64 {
	return Next()
}

func GetMulti(n int) (ids []int64) {
	for i := 0; i < n; i++ {
		ids = append(ids, Next())
	}
	return
}
