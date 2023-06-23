package storage

type Storage interface {}

var CurrentStorage Storage

func InitStorage(storageType string) error {
  switch storageType {
  }

  return nil
}

