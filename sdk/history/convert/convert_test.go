package convert

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hashicorp/waypoint/sdk/component"
	"github.com/hashicorp/waypoint/sdk/component/mocks"
	"github.com/hashicorp/waypoint/sdk/internal-shared/mapper"
)

func TestComponent(t *testing.T) {
	require := require.New(t)

	// Build our set
	funcA := mapper.TestFunc(t, func(int) *mocks.Deployment { return &mocks.Deployment{} })
	set := mapper.Set([]*mapper.Func{funcA})

	raw, err := Component(set,
		[]int{1, 2, 3},
		(**mocks.Deployment)(nil),
		(*component.Deployment)(nil),
	)
	require.NoError(err)
	require.NotNil(raw)

	result := raw.([]component.Deployment)
	require.Len(result, 3)
}
