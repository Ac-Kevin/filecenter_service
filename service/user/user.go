package user

type User struct {
	_UserName         string
	_Password         string
	_UploadRateByte   int64
	_DownloadRateByte int64
}

func (u *User) GetUserName() string {
	return u._UserName
}

func (u *User) GetUserPassword() string {
	return u._Password
}
func (u *User) GetFileUploadRateByte() int64 {
	return u._UploadRateByte
}
func (u *User) GetFileDownloadRateByte() int64 {
	return u._DownloadRateByte
}
