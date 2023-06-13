// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/members"
	"backend/ent/schema"
	"backend/ent/subject"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	membersFields := schema.Members{}.Fields()
	_ = membersFields
	// membersDescUsername is the schema descriptor for username field.
	membersDescUsername := membersFields[1].Descriptor()
	// members.DefaultUsername holds the default value on creation for the username field.
	members.DefaultUsername = membersDescUsername.Default.(string)
	// members.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	members.UsernameValidator = membersDescUsername.Validators[0].(func(string) error)
	// membersDescEmail is the schema descriptor for email field.
	membersDescEmail := membersFields[2].Descriptor()
	// members.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	members.EmailValidator = membersDescEmail.Validators[0].(func(string) error)
	// membersDescPassword is the schema descriptor for password field.
	membersDescPassword := membersFields[3].Descriptor()
	// members.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	members.PasswordValidator = func() func(string) error {
		validators := membersDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// membersDescNickname is the schema descriptor for nickname field.
	membersDescNickname := membersFields[4].Descriptor()
	// members.NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	members.NicknameValidator = membersDescNickname.Validators[0].(func(string) error)
	// membersDescAvatar is the schema descriptor for avatar field.
	membersDescAvatar := membersFields[5].Descriptor()
	// members.DefaultAvatar holds the default value on creation for the avatar field.
	members.DefaultAvatar = membersDescAvatar.Default.(string)
	// members.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	members.AvatarValidator = membersDescAvatar.Validators[0].(func(string) error)
	// membersDescGid is the schema descriptor for gid field.
	membersDescGid := membersFields[6].Descriptor()
	// members.DefaultGid holds the default value on creation for the gid field.
	members.DefaultGid = membersDescGid.Default.(uint8)
	subjectFields := schema.Subject{}.Fields()
	_ = subjectFields
	// subjectDescImage is the schema descriptor for image field.
	subjectDescImage := subjectFields[1].Descriptor()
	// subject.DefaultImage holds the default value on creation for the image field.
	subject.DefaultImage = subjectDescImage.Default.(string)
	// subjectDescSummary is the schema descriptor for summary field.
	subjectDescSummary := subjectFields[2].Descriptor()
	// subject.SummaryValidator is a validator for the "summary" field. It is called by the builders before save.
	subject.SummaryValidator = subjectDescSummary.Validators[0].(func(string) error)
	// subjectDescSubjectType is the schema descriptor for subject_type field.
	subjectDescSubjectType := subjectFields[9].Descriptor()
	// subject.DefaultSubjectType holds the default value on creation for the subject_type field.
	subject.DefaultSubjectType = subjectDescSubjectType.Default.(uint8)
	// subjectDescCollect is the schema descriptor for collect field.
	subjectDescCollect := subjectFields[10].Descriptor()
	// subject.DefaultCollect holds the default value on creation for the collect field.
	subject.DefaultCollect = subjectDescCollect.Default.(uint32)
}
