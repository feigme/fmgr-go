package repository

import (
	"testing"

	"github.com/feigme/fmgr-go/bootstrap"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	bootstrap.InitializeDB()

	option, err := optionRepo.Save("NIO211126P40000", "short", "1.52")
	require.NoError(t, err)
	require.NotNil(t, option.Id)
}
