package protocol

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContentTopic(t *testing.T) {
	ct := NewContentTopic("waku", 2, "test", "proto")
	require.Equal(t, ct.String(), "/waku/2/test/proto")

	_, err := StringToContentTopic("/waku/-1/a/b")
	require.Error(t, ErrInvalidFormat, err)

	_, err = StringToContentTopic("waku/1/a/b")
	require.Error(t, ErrInvalidFormat, err)

	_, err = StringToContentTopic("////")
	require.Error(t, ErrInvalidFormat, err)

	_, err = StringToContentTopic("/waku/1/a")
	require.Error(t, ErrInvalidFormat, err)

	ct2, err := StringToContentTopic("/waku/2/test/proto")
	require.NoError(t, err)
	require.Equal(t, ct.String(), ct2.String())
	require.True(t, ct.Equal(ct2))

	ct3 := NewContentTopic("waku", 2, "test2", "proto")
	require.False(t, ct.Equal(ct3))
}

func TestNsPubsubTopic(t *testing.T) {
	ns1 := NewNamedShardingPubsubTopic("waku-dev")
	require.Equal(t, "/waku/2/waku-dev", ns1.String())

	ns2 := NewStaticShardingPubsubTopic(0, 2)
	require.Equal(t, "/waku/2/rs/0/2", ns2.String())

	require.True(t, ns1.Equal(ns1))
	require.False(t, ns1.Equal(ns2))

	topic := "/waku/2/waku-dev"
	ns, err := ToShardedPubsubTopic(topic)
	require.NoError(t, err)
	require.Equal(t, NamedSharding, ns.Kind())
	require.Equal(t, "waku-dev", ns.(NamedShardingPubsubTopic).Name())

	topic = "/waku/2/rs/16/42"
	ns, err = ToShardedPubsubTopic(topic)
	require.NoError(t, err)
	require.Equal(t, StaticSharding, ns.Kind())
	require.Equal(t, uint16(16), ns.(StaticShardingPubsubTopic).Cluster())
	require.Equal(t, uint16(42), ns.(StaticShardingPubsubTopic).Shard())

	topic = "/waku/1/rs/16/42"
	_, err = ToShardedPubsubTopic(topic)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrInvalidTopicPrefix)

	topic = "/waku/2/rs//02"
	_, err = ToShardedPubsubTopic(topic)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrMissingClusterIndex)

	topic = "/waku/2/rs/xx/77"
	_, err = ToShardedPubsubTopic(topic)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrInvalidNumberFormat)
}
