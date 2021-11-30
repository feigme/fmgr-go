package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewOption(t *testing.T) {
	option, err := NewOption("NIO211126P40000", "short", "1.52")
	require.NoError(t, err)
	require.Equal(t, "NIO211126P40000", option.Code)
	require.Equal(t, "NIO", option.UnderlyingSecurities)
	require.Equal(t, "211126", option.ExerciseDate)
	require.Equal(t, "P", option.Type)
	require.Equal(t, int64(40000), option.StrikePrice)
	require.Equal(t, "short", option.Position)
	require.Equal(t, "1.52", option.SellPrice)

	option, err = NewOption("nio211126p40000", "Short", "1.52")
	require.NoError(t, err)
	require.Equal(t, "NIO211126P40000", option.Code)
	require.Equal(t, "NIO", option.UnderlyingSecurities)
	require.Equal(t, "211126", option.ExerciseDate)
	require.Equal(t, "P", option.Type)
	require.Equal(t, int64(40000), option.StrikePrice)
	require.Equal(t, "short", option.Position)
	require.Equal(t, "1.52", option.SellPrice)
}
