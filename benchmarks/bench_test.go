package benchmarks

import (
	"testing"

	"github.com/lalamove/konfig"
	"github.com/spf13/viper"
	"go-micro.dev/v5/config"
	"go-micro.dev/v5/config/source/memory"
)

func BenchmarkGetKonfig(b *testing.B) {
	var k = konfig.New(konfig.DefaultConfig())
	k.Set("foo", "bar")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.Get("foo")
	}
}

func BenchmarkStringKonfig(b *testing.B) {
	var k = konfig.New(konfig.DefaultConfig())
	k.Set("foo", "bar")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.String("foo")
	}
}

func BenchmarkGetViper(b *testing.B) {
	var v = viper.New()
	v.Set("foo", "bar")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Get("foo")
	}

}

func BenchmarkStringViper(b *testing.B) {
	var v = viper.New()
	v.Set("foo", "bar")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.GetString("foo")
	}
}

var data = []byte(`{
    "foo": "bar"
}`)

func newGoConfig() config.Config {
	memorySource := memory.NewSource(
		memory.WithJSON(data),
	)
	// Create new config
	conf, err := config.NewConfig()
	if err != nil {
		panic("Error when creating config: " + err.Error())
	}
	// Load file source
	conf.Load(memorySource)

	return conf
}

func BenchmarkGetGoConfig(b *testing.B) {
	conf := newGoConfig()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conf.Get("foo")
	}
}

func BenchmarkStringGoConfig(b *testing.B) {
	conf := newGoConfig()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conf.Get("foo").String("bar")
	}
}
