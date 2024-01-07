package errors

import "fmt"

type UnknownStorageTypeError struct {
	StorageType string
}

func NewUnknownStorageTypeError(storageType string) *UnknownStorageTypeError {
	return &UnknownStorageTypeError{StorageType: storageType}
}

func (e *UnknownStorageTypeError) Error() string {
	return fmt.Sprintf("Unknown storage type %s", e.StorageType)
}

type DatabaseConnectionError struct {
	Err string
}

func NewDatabaseConnectionError(err string) *DatabaseConnectionError {
	return &DatabaseConnectionError{Err: err}
}

func (e *DatabaseConnectionError) Error() string {
	return fmt.Sprintf("Database connection error: %s", e.Err)
}

type DatabaseMigrationError struct {
	Err string
}

func NewDatabaseMigrationError(err string) *DatabaseMigrationError {
	return &DatabaseMigrationError{Err: err}
}

func (e *DatabaseMigrationError) Error() string {
	return fmt.Sprintf("Database migration error: %s", e.Err)
}

type ObjectNotFoundError struct {
	Object string
}

func NewObjectNotFoundError(object string) *ObjectNotFoundError {
	return &ObjectNotFoundError{Object: object}
}

func (e *ObjectNotFoundError) Error() string {
	return fmt.Sprintf("Object not found error: %s", e.Object)
}

type ObjectAlreadyExistsError struct {
	Object string
}

func NewObjectAlreadyExistsError(object string) *ObjectAlreadyExistsError {
	return &ObjectAlreadyExistsError{Object: object}
}

func (e *ObjectAlreadyExistsError) Error() string {
	return fmt.Sprintf("Object already exists error: %s", e.Object)
}
