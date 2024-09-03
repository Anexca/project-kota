// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/cachedexam"
	"common/ent/exam"
	"common/ent/examattempt"
	"common/ent/examcategory"
	"common/ent/examsetting"
	"common/ent/generatedexam"
	"common/ent/schema"
	"common/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cachedexamFields := schema.CachedExam{}.Fields()
	_ = cachedexamFields
	// cachedexamDescIsUsed is the schema descriptor for is_used field.
	cachedexamDescIsUsed := cachedexamFields[1].Descriptor()
	// cachedexam.DefaultIsUsed holds the default value on creation for the is_used field.
	cachedexam.DefaultIsUsed = cachedexamDescIsUsed.Default.(bool)
	// cachedexamDescCreatedAt is the schema descriptor for created_at field.
	cachedexamDescCreatedAt := cachedexamFields[3].Descriptor()
	// cachedexam.DefaultCreatedAt holds the default value on creation for the created_at field.
	cachedexam.DefaultCreatedAt = cachedexamDescCreatedAt.Default.(func() time.Time)
	// cachedexamDescUpdatedAt is the schema descriptor for updated_at field.
	cachedexamDescUpdatedAt := cachedexamFields[4].Descriptor()
	// cachedexam.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cachedexam.DefaultUpdatedAt = cachedexamDescUpdatedAt.Default.(func() time.Time)
	// cachedexam.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	cachedexam.UpdateDefaultUpdatedAt = cachedexamDescUpdatedAt.UpdateDefault.(func() time.Time)
	examFields := schema.Exam{}.Fields()
	_ = examFields
	// examDescIsActive is the schema descriptor for is_active field.
	examDescIsActive := examFields[2].Descriptor()
	// exam.DefaultIsActive holds the default value on creation for the is_active field.
	exam.DefaultIsActive = examDescIsActive.Default.(bool)
	// examDescCreatedAt is the schema descriptor for created_at field.
	examDescCreatedAt := examFields[3].Descriptor()
	// exam.DefaultCreatedAt holds the default value on creation for the created_at field.
	exam.DefaultCreatedAt = examDescCreatedAt.Default.(func() time.Time)
	// examDescUpdatedAt is the schema descriptor for updated_at field.
	examDescUpdatedAt := examFields[4].Descriptor()
	// exam.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	exam.DefaultUpdatedAt = examDescUpdatedAt.Default.(func() time.Time)
	// exam.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	exam.UpdateDefaultUpdatedAt = examDescUpdatedAt.UpdateDefault.(func() time.Time)
	examattemptFields := schema.ExamAttempt{}.Fields()
	_ = examattemptFields
	// examattemptDescCreatedAt is the schema descriptor for created_at field.
	examattemptDescCreatedAt := examattemptFields[1].Descriptor()
	// examattempt.DefaultCreatedAt holds the default value on creation for the created_at field.
	examattempt.DefaultCreatedAt = examattemptDescCreatedAt.Default.(func() time.Time)
	// examattemptDescUpdatedAt is the schema descriptor for updated_at field.
	examattemptDescUpdatedAt := examattemptFields[2].Descriptor()
	// examattempt.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	examattempt.DefaultUpdatedAt = examattemptDescUpdatedAt.Default.(func() time.Time)
	// examattempt.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	examattempt.UpdateDefaultUpdatedAt = examattemptDescUpdatedAt.UpdateDefault.(func() time.Time)
	examcategoryFields := schema.ExamCategory{}.Fields()
	_ = examcategoryFields
	// examcategoryDescIsActive is the schema descriptor for is_active field.
	examcategoryDescIsActive := examcategoryFields[2].Descriptor()
	// examcategory.DefaultIsActive holds the default value on creation for the is_active field.
	examcategory.DefaultIsActive = examcategoryDescIsActive.Default.(bool)
	// examcategoryDescCreatedAt is the schema descriptor for created_at field.
	examcategoryDescCreatedAt := examcategoryFields[3].Descriptor()
	// examcategory.DefaultCreatedAt holds the default value on creation for the created_at field.
	examcategory.DefaultCreatedAt = examcategoryDescCreatedAt.Default.(func() time.Time)
	// examcategoryDescUpdatedAt is the schema descriptor for updated_at field.
	examcategoryDescUpdatedAt := examcategoryFields[4].Descriptor()
	// examcategory.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	examcategory.DefaultUpdatedAt = examcategoryDescUpdatedAt.Default.(func() time.Time)
	// examcategory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	examcategory.UpdateDefaultUpdatedAt = examcategoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	examsettingFields := schema.ExamSetting{}.Fields()
	_ = examsettingFields
	// examsettingDescMaxAttempts is the schema descriptor for max_attempts field.
	examsettingDescMaxAttempts := examsettingFields[5].Descriptor()
	// examsetting.DefaultMaxAttempts holds the default value on creation for the max_attempts field.
	examsetting.DefaultMaxAttempts = examsettingDescMaxAttempts.Default.(int)
	// examsettingDescCreatedAt is the schema descriptor for created_at field.
	examsettingDescCreatedAt := examsettingFields[7].Descriptor()
	// examsetting.DefaultCreatedAt holds the default value on creation for the created_at field.
	examsetting.DefaultCreatedAt = examsettingDescCreatedAt.Default.(func() time.Time)
	// examsettingDescUpdatedAt is the schema descriptor for updated_at field.
	examsettingDescUpdatedAt := examsettingFields[8].Descriptor()
	// examsetting.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	examsetting.DefaultUpdatedAt = examsettingDescUpdatedAt.Default.(func() time.Time)
	// examsetting.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	examsetting.UpdateDefaultUpdatedAt = examsettingDescUpdatedAt.UpdateDefault.(func() time.Time)
	generatedexamFields := schema.GeneratedExam{}.Fields()
	_ = generatedexamFields
	// generatedexamDescIsActive is the schema descriptor for is_active field.
	generatedexamDescIsActive := generatedexamFields[0].Descriptor()
	// generatedexam.DefaultIsActive holds the default value on creation for the is_active field.
	generatedexam.DefaultIsActive = generatedexamDescIsActive.Default.(bool)
	// generatedexamDescCreatedAt is the schema descriptor for created_at field.
	generatedexamDescCreatedAt := generatedexamFields[2].Descriptor()
	// generatedexam.DefaultCreatedAt holds the default value on creation for the created_at field.
	generatedexam.DefaultCreatedAt = generatedexamDescCreatedAt.Default.(func() time.Time)
	// generatedexamDescUpdatedAt is the schema descriptor for updated_at field.
	generatedexamDescUpdatedAt := generatedexamFields[3].Descriptor()
	// generatedexam.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	generatedexam.DefaultUpdatedAt = generatedexamDescUpdatedAt.Default.(func() time.Time)
	// generatedexam.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	generatedexam.UpdateDefaultUpdatedAt = generatedexamDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
}
