package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestAddEvenSuccess(t *testing.T){
	require.Equal(t, 4, add_even(2,2))
	require.Equal(t, 100, add_even(100,1))
	require.Equal(t, 2, add_even(1,2))
}