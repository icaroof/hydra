package hydradblayer

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkMySqlDbReads(b *testing.B) {
	dblayer, err := ConnectDatabase("mysql", "gouser:gouser@/Hydra")

	if err != nil {
		b.Fatal("Could not connect to database", err)
	}

	findMembersBM(b, dblayer)
}

func BenchmarkMongoDbReads(b *testing.B) {
	dblayer, err := ConnectDatabase("mongodb", "mongodb://127.0.0.1")

	if err != nil {
		b.Error("Could not connect to database", err)
		return
	}

	findMembersBM(b, dblayer)
}

func findMembersBM(b *testing.B, dblayer DBLayer) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		_, err := dblayer.FindMember(rand.Intn(16) + 1)

		if err != nil {
			b.Error("Query failed", err)
			return
		}
	}
}
