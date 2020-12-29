package gorm_crypto

import (
	"github.com/pkosilo/gorm-crypto/encryption"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
	"testing"
)

type TestModel struct {
	gorm.Model
	Name	string
	Address EncryptedValue
}

func TestGormIntegration(t *testing.T) {
	tmpDir := t.TempDir()
	dbFile := filepath.Join(tmpDir, "test.db")

	privateKey, publicKey, err := encryption.GenerateKeyPair(4096)
	if err != nil {
		t.Fatalf("Failed to generate key pair: %s\n", err.Error())
	}
	InitFromKeyPair(privateKey, publicKey)

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
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
