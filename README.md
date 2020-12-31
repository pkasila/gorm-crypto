# gorm-crypto
Encrypts any value of JSON compatible type

# Installation
To install `gorm-crypto` you need to, as usual, run this command:
```shell
go get github.com/pkosilo/gorm-crypto
```

# How to use it?
## Initialize
First of all you need to initialize the library **before migrating models**. To do so, you need to decide which
algorithm you want to use. (By the way, it's possible to select serializer/deserialize, so in the future you will
be able to customize your flow even more).
### RSA
If you want to use RSA, then you need to generate or load key pair and initialize library with `algorithm.RSA`
with your private and public keys. To do so, you can simply run code **like** this (you can also use different
public and private key files) in your app:
```golang
var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

// Different behaviour if there is already a PEM with private key or not
if _, err := os.Stat("private_key.pem"); os.IsNotExist(err) {
  // There is no PEM
  
  // Generate key pair
  privateKey, publicKey, err = helpers.RSAGenerateKeyPair(4096)

  if err != nil {
    panic(err)
  }

  // Store it
  privateKeyBytes := helpers.RSAPrivateKeyToBytes(privateKey)
  err := ioutil.WriteFile("private_key.pem", privateKeyBytes, 0644)

  if err != nil {
    panic(err)
  }
} else {
  // There is a PEM
  
  // Read PEM file with private key
  bytes, err := ioutil.ReadFile("private_key.pem")
  if err != nil {
    panic(err)
  }

  // Bytes to private key
  privateKey, err = helpers.RSABytesToPrivateKey(bytes)
  if err != nil {
    panic(err)
  }
  publicKey = &privateKey.PublicKey
}

// Use privateKey and publicKey to initialize gormcrypto
gormcrypto.Init(algorithms.NewRSA(privateKey, publicKey), serialization.NewJSON())
```
### AES
To use this library with AES, you need to initialize it with `algorithm.AES` with your key passed.
There is an example how to initialize library with AES:
```golang
aes, err := algorithms.NewAES([]byte("passphrasewhichneedstobe32bytes!"))
// algorithms.NewAES can fall with an error, so you should handle it
if err != nil {
panic(err)
}
gorm.Init(aes, serialization.NewJSON())
```

## Use it in a model
There is an example of a model with `EncryptedValue`:
```golang
type Application struct {
  gorm.Model
  Name               string                      `json:"name"`
  Phone              string                      `json:"phone"`
  Email              string                      `json:"email"`
  Address            gormcrypto.EncryptedValue  `json:"address"`
}
```

## Create a record with it
There is an example of `DB.Create` with a model with `EncryptedValue`:
```golang
application := Application {
  Name: "Anonymous",
  Phone: "+375290000000",
  Email: "example@example.com",
  Address: gormcrypto.EncryptedValue{Raw: "Oktyabr'skaya Ploshchad' 1, Minsk 220030"},
}

if err = models.DB.Create(&model).Error; err != nil {
  panic(err)
}
```

## Read a record with it
There is an example of finding a value using `DB.Find` and using `EncryptedValue.Raw` to access the decrypted value:
```golang
var application Application

if err := db.Find(&application, 1).Error; err != nil {
  panic(err)
}

fmt.Printf("Decrypted (raw) address: %s", application.Address.Raw.(string))
```
