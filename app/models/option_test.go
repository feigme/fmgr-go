package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewOption(t *testing.T) {
	option, err := NewOption("NIO211126P40000")
	require.NoError(t, err)
	require.Equal(t, "NIO211126P40000", option.Code)
	require.Equal(t, "NIO", option.Stock)
	require.Equal(t, "211126", option.ExerciseDate)
	require.Equal(t, "P", option.Type)
	require.Equal(t, "40000", option.StrikePrice)

	option, err = NewOption("nio211126p40000")
	require.NoError(t, err)
	require.Equal(t, "NIO211126P40000", option.Code)
	require.Equal(t, "NIO", option.Stock)
	require.Equal(t, "211126", option.ExerciseDate)
	require.Equal(t, "P", option.Type)
	require.Equal(t, "40000", option.StrikePrice)

}
