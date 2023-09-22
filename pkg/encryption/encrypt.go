package encryption

// import (
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"errors"
// 	"io"
// 	"savannah/cmd/config"
// )

// type EncryptionService struct {
// 	key *config.Conf
// }

// func NewEncryptionService(key []byte) *EncryptionService {
// 	return &EncryptionService{key: key}
// }

// func (e *EncryptionService) Encrypt(plaintext []byte) ([]byte, error) {
// 	block, err := aes.NewCipher([]byte(e.key.Security.Key))
// 	if err != nil {
// 		return nil, err
// 	}

// 	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
// 	iv := ciphertext[:aes.BlockSize]
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
// 		return nil, err
// 	}

// 	mode := cipher.NewCBCEncrypter(block, iv)
// 	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

// 	return ciphertext, nil
// }

// func (e *EncryptionService) Decrypt(ciphertext []byte) ([]byte, error) {
// 	block, err := aes.NewCipher(e.key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(ciphertext) < aes.BlockSize {
// 		return nil, errors.New("Ciphertext is too short")
// 	}

// 	iv := ciphertext[:aes.BlockSize]
// 	ciphertext = ciphertext[aes.BlockSize:]

// 	if len(ciphertext)%aes.BlockSize != 0 {
// 		return nil, errors.New("Ciphertext is not a multiple of the block size")
// 	}

// 	mode := cipher.NewCBCDecrypter(block, iv)
// 	mode.CryptBlocks(ciphertext, ciphertext)

// 	return ciphertext, nil
// }
