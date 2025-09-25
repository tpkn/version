package version

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Update(t *testing.T) {
	_, err := Update(&Args{IncrementMajor: true, IncrementOptions: "M", FilePath: ""})
	require.ErrorContains(t, err, "you can't use a targeted version incrementation")

	_, err = Update(&Args{IncrementMajor: true, IncrementPatch: true, FilePath: ""})
	require.ErrorContains(t, err, "targeted version incrementation")

	_, err = Update(&Args{IncrementOptions: "XYU", FilePath: ""})
	require.ErrorContains(t, err, "unknown incrementation flag")

	// ------------------------------

	temp_version_path := filepath.Join(t.TempDir(), "version.test")

	args := &Args{IncrementPatch: true, FilePath: temp_version_path}
	got, err := Update(args)
	require.NoError(t, err)
	require.Equal(t, "0.0.1", got)

	args = &Args{IncrementMinor: true, FilePath: temp_version_path}
	got, err = Update(args)
	require.NoError(t, err)
	require.Equal(t, "0.1.0", got)

	args = &Args{IncrementMajor: true, FilePath: temp_version_path}
	got, err = Update(args)
	require.NoError(t, err)
	require.Equal(t, "1.0.0", got)

	args = &Args{IncrementPatch: true, FilePath: temp_version_path}
	got, err = Update(args)
	require.NoError(t, err)
	require.Equal(t, "1.0.1", got)

	args = &Args{IncrementOptions: "M", FilePath: temp_version_path}
	got, err = Update(args)
	require.NoError(t, err)
	require.Equal(t, "2.0.0", got)

	temp_version_path = filepath.Join(t.TempDir(), "version.test")

	args = &Args{IncrementOptions: "MMM", FilePath: temp_version_path}
	got, err = Update(args)
	require.NoError(t, err)
	require.Equal(t, "3.0.0", got)

	args = &Args{IncrementOptions: "MMMppp", FilePath: temp_version_path}
	got, err = Update(args)
	require.NoError(t, err)
	require.Equal(t, "6.0.3", got)

	args = &Args{IncrementOptions: "mmm", FilePath: temp_version_path}
	got, err = Update(args)
	require.NoError(t, err)
	require.Equal(t, "6.3.0", got)

	args = &Args{IncrementOptions: "mmmM", FilePath: temp_version_path}
	got, err = Update(args)
	require.NoError(t, err)
	require.Equal(t, "7.0.0", got)
}

func Test_changeVersion(t *testing.T) {
	_, err := changeVersion("a.b.c", "M")
	require.ErrorContains(t, err, "version should match")

	_, err = changeVersion("0.0.0", "XYU")
	require.ErrorContains(t, err, "unknown incrementation flag")

	// ------------------------------

	tests := []struct {
		input string
		want  string
	}{
		{"", "0.0.0"},
		{"M", "1.0.0"},
		{"m", "0.1.0"},
		{"p", "0.0.1"},
		{"mm", "0.2.0"},
		{"MMMM", "4.0.0"},
		{"Mmp", "1.1.1"},
		{"mMp", "1.0.1"},
		{"mmmmp", "0.4.1"},
		{"Mpppppppppppppp", "1.0.14"},
		{"ppppppppppppppM", "1.0.0"},
		{"MMMMMMmmmmmMmmmmmppppppppppppp", "7.5.13"},
	}

	for i, tt := range tests {
		got, err := changeVersion("0.0.0", tt.input)
		require.NoError(t, err)
		require.Equal(t, tt.want, got, i)
	}
}

func Test_isValidVersion(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"a.b.cc", false},
		{"0.0.0", true},
		{"123.767.990", true},
		{"1.2.3.4.5", false},
		{"1-.0.0", false},
		{"1...0.0", false},
	}

	for _, tt := range tests {
		got := isValidVersion(tt.input)
		require.Equal(t, tt.want, got)
	}
}
