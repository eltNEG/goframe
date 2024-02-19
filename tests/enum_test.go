package tests

import (
	"goframe/enum"
	"testing"
)

func TestEnum(t *testing.T) {
	type sampleEnum string

	sampleEnums := enum.MakeEnum[sampleEnum](struct {
		VALUE1 sampleEnum
		VALUE2 sampleEnum
	}{
		VALUE1: "value1",
		VALUE2: "value2",
	})
	tests := []struct {
		name  string
		value sampleEnum
		want  bool
	}{
		{
			name:  "test exist value1",
			value: sampleEnums.V.VALUE1,
			want:  true,
		},
		{
			name:  "test exit value2",
			value: sampleEnums.V.VALUE1,
			want:  true,
		},
		{
			name:  "test non exist value3",
			value: "value3",
			want:  false,
		},
		{
			name:  "test non exist empty value",
			value: "",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleEnums.IsValid(tt.value); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
