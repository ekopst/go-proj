package web

import "encoding/gob"

func init() {
	gob.Register(CreatePostForm{})
	gob.Register(CreateThreadForm{})
	gob.Register(CreateCommentForm{})
	gob.Register(RegisterForm{})
	gob.Register(LoginForm{})
	gob.Register(FormErrors{})
}

type FormErrors map[string]string

type CreatePostForm struct {
	Title   string
	Content string

	Errors FormErrors
}

func (f *CreatePostForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Title == "" {
		f.Errors["Title"] = "Masukin judul."
	}
	if f.Content == "" {
		f.Errors["Content"] = "Masukin text."
	}

	return len(f.Errors) == 0
}

type CreateThreadForm struct {
	Title       string
	Description string

	Errors FormErrors
}

func (f *CreateThreadForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Title == "" {
		f.Errors["Title"] = "Masukin Judul."
	}
	if f.Description == "" {
		f.Errors["Description"] = "Masukin description."
	}

	return len(f.Errors) == 0
}

type CreateCommentForm struct {
	Content string

	Errors FormErrors
}

func (f *CreateCommentForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Content == "" {
		f.Errors["Content"] = "Masukin komen."
	}

	return len(f.Errors) == 0
}

type RegisterForm struct {
	Username      string
	Password      string
	UsernameTaken bool

	Errors FormErrors
}

func (f *RegisterForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Username == "" {
		f.Errors["Username"] = "Masukin username."
	} else if f.UsernameTaken {
		f.Errors["Username"] = "Username ini sudah dipakai."
	}

	if f.Password == "" {
		f.Errors["Password"] = "Masukin Password."
	} else if len(f.Password) < 8 {
		f.Errors["Password"] = "Pssword minimal 8 karakter."
	}

	return len(f.Errors) == 0
}

type LoginForm struct {
	Username             string
	Password             string
	IncorrectCredentials bool

	Errors FormErrors
}

func (f *LoginForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Username == "" {
		f.Errors["Username"] = "Masukin username."
	} else if f.IncorrectCredentials {
		f.Errors["Username"] = "Username atau password salah."
	}

	if f.Password == "" {
		f.Errors["Password"] = "Masukin password."
	}

	return len(f.Errors) == 0
}
