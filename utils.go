package main

import (
    "io/ioutil"
    "crypto/sha256"
    "log"
    "encoding/hex"
    "os"
    "io"
    "archive/zip"
    "path/filepath"
    "fmt"
)

/* Returns SHA256 checksum of the file asked. */ 
func get_sha256(hashpath string) string {
  hasher := sha256.New()
  s, err := ioutil.ReadFile(hashpath)
  hasher.Write(s)
  if err != nil {
    log.Fatal(err)
  }
  return hex.EncodeToString(hasher.Sum(nil))
}

/* Given a Koodous zipfile or just a random zipfile
it checks whether it's an Android Application or not */ 

func is_apk(zippath string) bool {
  r, err := zip.OpenReader(zippath)
  if err != nil {
    log.Fatal(err)
  }
  defer r.Close()

  for _, f := range r.File {
    if (f.Name == "AndroidManifest.xml") {
      return true
    }
  }
  return false
}

/* Extracts compressed apks downloaded from Koodous */ 

func checkError(e error){
  if e != nil {
    panic(e)
  }
}
func cloneZipItem(f *zip.File, dest string){

    path := filepath.Join(dest, f.Name)
    fmt.Println("Creating", path)
    err := os.MkdirAll(filepath.Dir(path), os.ModeDir|os.ModePerm)
    checkError(err)

    rc, err := f.Open()
    checkError(err)
    if !f.FileInfo().IsDir() {
        fileCopy, err := os.Create(path)
        checkError(err)
        _, err = io.Copy(fileCopy, rc)
        fileCopy.Close()
        checkError(err)
    }
    rc.Close()
}

func Extract(zip_path, dest string) {
    r, err := zip.OpenReader(zip_path)
    checkError(err)
    defer r.Close()
    for _, f := range r.File {
        cloneZipItem(f, dest)
    }
}
