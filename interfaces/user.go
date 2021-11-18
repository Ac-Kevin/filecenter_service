package interfaces

// UserHandler Def .
type UserHandler interface {
	Save(user User) error
	Del(user User) error
	Get(username string) (User, error)
	List(index, limit int) ([]User, error)
	All() ([]User, error)
}

// UserInterface Def .
type User interface {
	GetUserName() string
	GetUserPassword() string
	GetFileUploadRateByte() int64
	GetFileDownloadRateByte() int64
	// SetUserName(username string) error
	// SetUserPassword(password string) error
	// SetFileUploadRateByte(rate int64) error
	// SetFileDownloadRateByte(rate int64) error
}
