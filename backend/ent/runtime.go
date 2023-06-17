// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/collection"
	"backend/ent/members"
	"backend/ent/schema"
	"backend/ent/subject"
	"backend/ent/subjectfield"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	collectionFields := schema.Collection{}.Fields()
	_ = collectionFields
	// collectionDescType is the schema descriptor for type field.
	collectionDescType := collectionFields[1].Descriptor()
	// collection.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	collection.TypeValidator = func() func(uint8) error {
		validators := collectionDescType.Validators
		fns := [...]func(uint8) error{
			validators[0].(func(uint8) error),
			validators[1].(func(uint8) error),
		}
		return func(_type uint8) error {
			for _, fn := range fns {
				if err := fn(_type); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// collectionDescHasComment is the schema descriptor for has_comment field.
	collectionDescHasComment := collectionFields[2].Descriptor()
	// collection.DefaultHasComment holds the default value on creation for the has_comment field.
	collection.DefaultHasComment = collectionDescHasComment.Default.(bool)
	// collectionDescComment is the schema descriptor for comment field.
	collectionDescComment := collectionFields[3].Descriptor()
	// collection.DefaultComment holds the default value on creation for the comment field.
	collection.DefaultComment = collectionDescComment.Default.(string)
	// collection.CommentValidator is a validator for the "comment" field. It is called by the builders before save.
	collection.CommentValidator = collectionDescComment.Validators[0].(func(string) error)
	// collectionDescScore is the schema descriptor for score field.
	collectionDescScore := collectionFields[4].Descriptor()
	// collection.DefaultScore holds the default value on creation for the score field.
	collection.DefaultScore = collectionDescScore.Default.(uint8)
	// collectionDescAddTime is the schema descriptor for add_time field.
	collectionDescAddTime := collectionFields[5].Descriptor()
	// collection.DefaultAddTime holds the default value on creation for the add_time field.
	collection.DefaultAddTime = collectionDescAddTime.Default.(string)
	// collectionDescEpStatus is the schema descriptor for ep_status field.
	collectionDescEpStatus := collectionFields[6].Descriptor()
	// collection.DefaultEpStatus holds the default value on creation for the ep_status field.
	collection.DefaultEpStatus = collectionDescEpStatus.Default.(uint8)
	membersFields := schema.Members{}.Fields()
	_ = membersFields
	// membersDescUsername is the schema descriptor for username field.
	membersDescUsername := membersFields[1].Descriptor()
	// members.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	members.UsernameValidator = func() func(string) error {
		validators := membersDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// membersDescEmail is the schema descriptor for email field.
	membersDescEmail := membersFields[2].Descriptor()
	// members.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	members.EmailValidator = membersDescEmail.Validators[0].(func(string) error)
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
	// subject.ImageValidator is a validator for the "image" field. It is called by the builders before save.
	subject.ImageValidator = subjectDescImage.Validators[0].(func(string) error)
	// subjectDescSummary is the schema descriptor for summary field.
	subjectDescSummary := subjectFields[2].Descriptor()
	// subject.DefaultSummary holds the default value on creation for the summary field.
	subject.DefaultSummary = subjectDescSummary.Default.(string)
	// subject.SummaryValidator is a validator for the "summary" field. It is called by the builders before save.
	subject.SummaryValidator = subjectDescSummary.Validators[0].(func(string) error)
	// subjectDescWish is the schema descriptor for wish field.
	subjectDescWish := subjectFields[6].Descriptor()
	// subject.DefaultWish holds the default value on creation for the wish field.
	subject.DefaultWish = subjectDescWish.Default.(uint32)
	// subjectDescDoing is the schema descriptor for doing field.
	subjectDescDoing := subjectFields[7].Descriptor()
	// subject.DefaultDoing holds the default value on creation for the doing field.
	subject.DefaultDoing = subjectDescDoing.Default.(uint32)
	// subjectDescWatched is the schema descriptor for watched field.
	subjectDescWatched := subjectFields[8].Descriptor()
	// subject.DefaultWatched holds the default value on creation for the watched field.
	subject.DefaultWatched = subjectDescWatched.Default.(uint32)
	// subjectDescOnHold is the schema descriptor for on_hold field.
	subjectDescOnHold := subjectFields[9].Descriptor()
	// subject.DefaultOnHold holds the default value on creation for the on_hold field.
	subject.DefaultOnHold = subjectDescOnHold.Default.(uint32)
	// subjectDescDropped is the schema descriptor for dropped field.
	subjectDescDropped := subjectFields[10].Descriptor()
	// subject.DefaultDropped holds the default value on creation for the dropped field.
	subject.DefaultDropped = subjectDescDropped.Default.(uint32)
	subjectfieldFields := schema.SubjectField{}.Fields()
	_ = subjectfieldFields
	// subjectfieldDescRate1 is the schema descriptor for rate_1 field.
	subjectfieldDescRate1 := subjectfieldFields[0].Descriptor()
	// subjectfield.DefaultRate1 holds the default value on creation for the rate_1 field.
	subjectfield.DefaultRate1 = subjectfieldDescRate1.Default.(uint32)
	// subjectfieldDescRate2 is the schema descriptor for rate_2 field.
	subjectfieldDescRate2 := subjectfieldFields[1].Descriptor()
	// subjectfield.DefaultRate2 holds the default value on creation for the rate_2 field.
	subjectfield.DefaultRate2 = subjectfieldDescRate2.Default.(uint32)
	// subjectfieldDescRate3 is the schema descriptor for rate_3 field.
	subjectfieldDescRate3 := subjectfieldFields[2].Descriptor()
	// subjectfield.DefaultRate3 holds the default value on creation for the rate_3 field.
	subjectfield.DefaultRate3 = subjectfieldDescRate3.Default.(uint32)
	// subjectfieldDescRate4 is the schema descriptor for rate_4 field.
	subjectfieldDescRate4 := subjectfieldFields[3].Descriptor()
	// subjectfield.DefaultRate4 holds the default value on creation for the rate_4 field.
	subjectfield.DefaultRate4 = subjectfieldDescRate4.Default.(uint32)
	// subjectfieldDescRate5 is the schema descriptor for rate_5 field.
	subjectfieldDescRate5 := subjectfieldFields[4].Descriptor()
	// subjectfield.DefaultRate5 holds the default value on creation for the rate_5 field.
	subjectfield.DefaultRate5 = subjectfieldDescRate5.Default.(uint32)
	// subjectfieldDescRate6 is the schema descriptor for rate_6 field.
	subjectfieldDescRate6 := subjectfieldFields[5].Descriptor()
	// subjectfield.DefaultRate6 holds the default value on creation for the rate_6 field.
	subjectfield.DefaultRate6 = subjectfieldDescRate6.Default.(uint32)
	// subjectfieldDescRate7 is the schema descriptor for rate_7 field.
	subjectfieldDescRate7 := subjectfieldFields[6].Descriptor()
	// subjectfield.DefaultRate7 holds the default value on creation for the rate_7 field.
	subjectfield.DefaultRate7 = subjectfieldDescRate7.Default.(uint32)
	// subjectfieldDescRate8 is the schema descriptor for rate_8 field.
	subjectfieldDescRate8 := subjectfieldFields[7].Descriptor()
	// subjectfield.DefaultRate8 holds the default value on creation for the rate_8 field.
	subjectfield.DefaultRate8 = subjectfieldDescRate8.Default.(uint32)
	// subjectfieldDescRate9 is the schema descriptor for rate_9 field.
	subjectfieldDescRate9 := subjectfieldFields[8].Descriptor()
	// subjectfield.DefaultRate9 holds the default value on creation for the rate_9 field.
	subjectfield.DefaultRate9 = subjectfieldDescRate9.Default.(uint32)
	// subjectfieldDescRate10 is the schema descriptor for rate_10 field.
	subjectfieldDescRate10 := subjectfieldFields[9].Descriptor()
	// subjectfield.DefaultRate10 holds the default value on creation for the rate_10 field.
	subjectfield.DefaultRate10 = subjectfieldDescRate10.Default.(uint32)
	// subjectfieldDescAverageScore is the schema descriptor for average_score field.
	subjectfieldDescAverageScore := subjectfieldFields[10].Descriptor()
	// subjectfield.DefaultAverageScore holds the default value on creation for the average_score field.
	subjectfield.DefaultAverageScore = subjectfieldDescAverageScore.Default.(float64)
	// subjectfield.AverageScoreValidator is a validator for the "average_score" field. It is called by the builders before save.
	subjectfield.AverageScoreValidator = subjectfieldDescAverageScore.Validators[0].(func(float64) error)
	// subjectfieldDescRank is the schema descriptor for rank field.
	subjectfieldDescRank := subjectfieldFields[11].Descriptor()
	// subjectfield.DefaultRank holds the default value on creation for the rank field.
	subjectfield.DefaultRank = subjectfieldDescRank.Default.(uint32)
	// subjectfieldDescMonth is the schema descriptor for month field.
	subjectfieldDescMonth := subjectfieldFields[13].Descriptor()
	// subjectfield.MonthValidator is a validator for the "month" field. It is called by the builders before save.
	subjectfield.MonthValidator = func() func(uint8) error {
		validators := subjectfieldDescMonth.Validators
		fns := [...]func(uint8) error{
			validators[0].(func(uint8) error),
			validators[1].(func(uint8) error),
		}
		return func(month uint8) error {
			for _, fn := range fns {
				if err := fn(month); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// subjectfieldDescDate is the schema descriptor for date field.
	subjectfieldDescDate := subjectfieldFields[14].Descriptor()
	// subjectfield.DateValidator is a validator for the "date" field. It is called by the builders before save.
	subjectfield.DateValidator = func() func(uint8) error {
		validators := subjectfieldDescDate.Validators
		fns := [...]func(uint8) error{
			validators[0].(func(uint8) error),
			validators[1].(func(uint8) error),
		}
		return func(date uint8) error {
			for _, fn := range fns {
				if err := fn(date); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// subjectfieldDescWeekday is the schema descriptor for weekday field.
	subjectfieldDescWeekday := subjectfieldFields[15].Descriptor()
	// subjectfield.WeekdayValidator is a validator for the "weekday" field. It is called by the builders before save.
	subjectfield.WeekdayValidator = func() func(uint8) error {
		validators := subjectfieldDescWeekday.Validators
		fns := [...]func(uint8) error{
			validators[0].(func(uint8) error),
			validators[1].(func(uint8) error),
		}
		return func(weekday uint8) error {
			for _, fn := range fns {
				if err := fn(weekday); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
