package idgen

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zchat-team/zim/pkg/idgen/snowflake"
)

func TestNext(t *testing.T) {
	require.NotEqual(t, 0, Next(), "should generate id")

	// setup
	st := snowflake.Settings{
		StartTime: time.Date(1883, 1, 1, 0, 0, 0, 0, time.UTC),
		MachineID: getMachineId,
	}
	sf = snowflake.NewSnowflake(st)

	id := Next()
	require.Equal(t, 0, id, "should return zero when over the time limit")

	// teardown
	st = snowflake.Settings{
		MachineID: getMachineId,
	}
	sf = snowflake.NewSnowflake(st)
}

func TestGetOne(t *testing.T) {
	require.NotEqual(t, 0, GetOne(), "should generate one ID")
}

func TestGetMulti(t *testing.T) {
	ids := GetMulti(3)
	require.Len(t, ids, 3)
	for _, v := range ids[:] {
		require.Greater(t, v, int64(0))
	}
}
