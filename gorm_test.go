package gormcrypto

import (
	"github.com/pkosilo/gorm-crypto/algorithms"
	"github.com/pkosilo/gorm-crypto/helpers"
	"github.com/pkosilo/gorm-crypto/serialization"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type TestModel struct {
	gorm.Model
	Name    string
	Address EncryptedValue
}

func TestRSA(t *testing.T) {
	privateKey, publicKey, err := helpers.RSAGenerateKeyPair(4096)
	if err != nil {
		t.Fatalf("Failed to generate key pair: %s\n", err.Error())
	}

	Init(algorithms.NewRSA(privateKey, publicKey), serialization.NewJSON())

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Cannot open test DB: %s\n", err.Error())
	}

	err = db.AutoMigrate(&TestModel{})
	if err != nil {
		t.Fatalf("Cannot migrate: %s\n", err.Error())
	}

	test := TestModel{
		Name:    "Anonymous",
		Address: EncryptedValue{Raw: "221b Baker St, Marylebone, London NW1 6XE"},
	}

	if err = db.Create(&test).Error; err != nil {
		t.Fatalf("Cannot create TestModel entity: %s\n", err.Error())
	}

	var testExtracted TestModel
	if err = db.Find(&testExtracted, test.ID).Error; err != nil {
		t.Fatalf("Cannot find TestModel entity: %s\n", err.Error())
	}

	if test.Address.Raw != testExtracted.Address.Raw {
		t.Fatalf("Fields aren't equal: %s != %s\n", test.Address.Raw.(string), testExtracted.Address.Raw.(string))
	}
}

func TestAES(t *testing.T) {
	aes, err := algorithms.NewAES([]byte("passphrasewhichneedstobe32bytes!"))
	// algorithms.NewAES can fall with an error, so you should handle it
	if err != nil {
		panic(err)
	}
	Init(aes, serialization.NewJSON())

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Cannot open test DB: %s\n", err.Error())
	}

	err = db.AutoMigrate(&TestModel{})
	if err != nil {
		t.Fatalf("Cannot migrate: %s\n", err.Error())
	}

	test := TestModel{
		Name:    "Anonymous",
		Address: EncryptedValue{Raw: "221b Baker St, Marylebone, London NW1 6XE"},
	}

	if err = db.Create(&test).Error; err != nil {
		t.Fatalf("Cannot create TestModel entity: %s\n", err.Error())
	}

	var testExtracted TestModel
	if err = db.Find(&testExtracted, test.ID).Error; err != nil {
		t.Fatalf("Cannot find TestModel entity: %s\n", err.Error())
	}

	if test.Address.Raw != testExtracted.Address.Raw {
		t.Fatalf("Fields aren't equal: %s != %s\n", test.Address.Raw.(string), testExtracted.Address.Raw.(string))
	}
}
