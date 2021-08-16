package tinytest

import (
	"testing"
)

func BenchmarkStd_Unmarshal_Env(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s Env
		err := s.UnmarshalJSON(sampleEnvText)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkStd_Unmarshal_Info(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s MessageInfo
		err := s.UnmarshalJSON(sampleMsgInfoText)
		if err != nil {
			b.Error(err)
		}
	}
}
