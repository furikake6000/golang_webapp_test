package credentials

import (
  "os"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "crypto/cipher"
  "crypto/aes"
)

var (
  credentials map[string]string
  master_key []byte
  block cipher.Block
  commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
)

func init() {
  // get master_key
  master_key := get_master_key()
  // get crypto block
  var err error
  block, err = aes.NewCipher(master_key)
  if err != nil {
    panic(err)
  }
  // refresh encrypted credentials
  encrypt()
  // load credentials
  credentials = make(map[string]string)
  load()
}

// master_keyの取得
func get_master_key() []byte {
  // 環境変数からの読み取り
  master_key_env := os.Getenv("MASTER_KEY")
  if (master_key_env != "") {
    return []byte(master_key_env)
  }

  // ファイルからの読み取り
  file, err := os.Open("secret/master.key")
  if err != nil {
    panic(err)
  }
  buf, err := ioutil.ReadAll(file)
  if err != nil {
    panic(err)
  }
  return buf
}

// 暗号化されたファイルからメモリに読み込む
func load() {
  // ファイル読み取り
  file, err := os.Open("secret/credentials.json.enc")
  if err != nil {
    panic(err)
  }
  buf, err := ioutil.ReadAll(file)
  if err != nil {
    panic(err)
  }

  cfbdec := cipher.NewCFBDecrypter(block, commonIV)
  decrypted := make([]byte, len(buf))
  cfbdec.XORKeyStream(decrypted, buf)

  err = json.Unmarshal(decrypted, &credentials)
  if err != nil {
    panic(err)
  }
}

// 平文ファイルから暗号化ファイルに変換する
func encrypt() {
  // ファイル読み取り
  file, err := os.Open("secret/credentials.json")
  if err != nil {
    fmt.Print("No raw credentials file found. Refresh aborted.")
    return
  }
  buf, err := ioutil.ReadAll(file)
  if err != nil {
    panic(err)
  }

  cfb := cipher.NewCFBEncrypter(block, commonIV)
  encrypted := make([]byte, len(buf))
  cfb.XORKeyStream(encrypted, buf)

  file, err = os.Create("secret/credentials.json.enc")
  if err != nil {
    panic(err)
  }
  file.Write(encrypted)
  file.Close()
}

// 暗号化ファイルから平文ファイルに変換する
func decrypt() {
  // ファイル読み取り
  file, err := os.Open("secret/credentials.json.enc")
  if err != nil {
    panic(err)
  }
  buf, err := ioutil.ReadAll(file)
  if err != nil {
    panic(err)
  }

  cfbdec := cipher.NewCFBDecrypter(block, commonIV)
  decrypted := make([]byte, len(buf))
  cfbdec.XORKeyStream(decrypted, buf)

  file, err = os.Create("secret/credentials.json.dec")
  if err != nil {
    panic(err)
  }
  file.Write(decrypted)
  file.Close()
}