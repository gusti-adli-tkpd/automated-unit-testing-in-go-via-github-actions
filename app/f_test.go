package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestF0(t *testing.T) {
	require.Equal(t, struct0{"ccc", 999}, f0(1))
}
