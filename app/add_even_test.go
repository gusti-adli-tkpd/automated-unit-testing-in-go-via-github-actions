package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestAddEvenSuccess(t *testing.T){
	require.Equal(t, 4, add_even(2,2))
}