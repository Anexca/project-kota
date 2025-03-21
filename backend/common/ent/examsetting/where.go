// Code generated by ent, DO NOT EDIT.

package examsetting

import (
	"common/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldID, id))
}

// NumberOfQuestions applies equality check predicate on the "number_of_questions" field. It's identical to NumberOfQuestionsEQ.
func NumberOfQuestions(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldNumberOfQuestions, v))
}

// DurationSeconds applies equality check predicate on the "duration_seconds" field. It's identical to DurationSecondsEQ.
func DurationSeconds(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldDurationSeconds, v))
}

// NegativeMarking applies equality check predicate on the "negative_marking" field. It's identical to NegativeMarkingEQ.
func NegativeMarking(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldNegativeMarking, v))
}

// AiPrompt applies equality check predicate on the "ai_prompt" field. It's identical to AiPromptEQ.
func AiPrompt(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldAiPrompt, v))
}

// MaxAttempts applies equality check predicate on the "max_attempts" field. It's identical to MaxAttemptsEQ.
func MaxAttempts(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldMaxAttempts, v))
}

// TotalMarks applies equality check predicate on the "total_marks" field. It's identical to TotalMarksEQ.
func TotalMarks(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldTotalMarks, v))
}

// CutoffMarks applies equality check predicate on the "cutoff_marks" field. It's identical to CutoffMarksEQ.
func CutoffMarks(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldCutoffMarks, v))
}

// EvaluationAiPrompt applies equality check predicate on the "evaluation_ai_prompt" field. It's identical to EvaluationAiPromptEQ.
func EvaluationAiPrompt(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldEvaluationAiPrompt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldUpdatedAt, v))
}

// NumberOfQuestionsEQ applies the EQ predicate on the "number_of_questions" field.
func NumberOfQuestionsEQ(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldNumberOfQuestions, v))
}

// NumberOfQuestionsNEQ applies the NEQ predicate on the "number_of_questions" field.
func NumberOfQuestionsNEQ(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldNumberOfQuestions, v))
}

// NumberOfQuestionsIn applies the In predicate on the "number_of_questions" field.
func NumberOfQuestionsIn(vs ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldNumberOfQuestions, vs...))
}

// NumberOfQuestionsNotIn applies the NotIn predicate on the "number_of_questions" field.
func NumberOfQuestionsNotIn(vs ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldNumberOfQuestions, vs...))
}

// NumberOfQuestionsGT applies the GT predicate on the "number_of_questions" field.
func NumberOfQuestionsGT(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldNumberOfQuestions, v))
}

// NumberOfQuestionsGTE applies the GTE predicate on the "number_of_questions" field.
func NumberOfQuestionsGTE(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldNumberOfQuestions, v))
}

// NumberOfQuestionsLT applies the LT predicate on the "number_of_questions" field.
func NumberOfQuestionsLT(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldNumberOfQuestions, v))
}

// NumberOfQuestionsLTE applies the LTE predicate on the "number_of_questions" field.
func NumberOfQuestionsLTE(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldNumberOfQuestions, v))
}

// DurationSecondsEQ applies the EQ predicate on the "duration_seconds" field.
func DurationSecondsEQ(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldDurationSeconds, v))
}

// DurationSecondsNEQ applies the NEQ predicate on the "duration_seconds" field.
func DurationSecondsNEQ(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldDurationSeconds, v))
}

// DurationSecondsIn applies the In predicate on the "duration_seconds" field.
func DurationSecondsIn(vs ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldDurationSeconds, vs...))
}

// DurationSecondsNotIn applies the NotIn predicate on the "duration_seconds" field.
func DurationSecondsNotIn(vs ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldDurationSeconds, vs...))
}

// DurationSecondsGT applies the GT predicate on the "duration_seconds" field.
func DurationSecondsGT(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldDurationSeconds, v))
}

// DurationSecondsGTE applies the GTE predicate on the "duration_seconds" field.
func DurationSecondsGTE(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldDurationSeconds, v))
}

// DurationSecondsLT applies the LT predicate on the "duration_seconds" field.
func DurationSecondsLT(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldDurationSeconds, v))
}

// DurationSecondsLTE applies the LTE predicate on the "duration_seconds" field.
func DurationSecondsLTE(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldDurationSeconds, v))
}

// NegativeMarkingEQ applies the EQ predicate on the "negative_marking" field.
func NegativeMarkingEQ(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldNegativeMarking, v))
}

// NegativeMarkingNEQ applies the NEQ predicate on the "negative_marking" field.
func NegativeMarkingNEQ(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldNegativeMarking, v))
}

// NegativeMarkingIn applies the In predicate on the "negative_marking" field.
func NegativeMarkingIn(vs ...float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldNegativeMarking, vs...))
}

// NegativeMarkingNotIn applies the NotIn predicate on the "negative_marking" field.
func NegativeMarkingNotIn(vs ...float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldNegativeMarking, vs...))
}

// NegativeMarkingGT applies the GT predicate on the "negative_marking" field.
func NegativeMarkingGT(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldNegativeMarking, v))
}

// NegativeMarkingGTE applies the GTE predicate on the "negative_marking" field.
func NegativeMarkingGTE(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldNegativeMarking, v))
}

// NegativeMarkingLT applies the LT predicate on the "negative_marking" field.
func NegativeMarkingLT(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldNegativeMarking, v))
}

// NegativeMarkingLTE applies the LTE predicate on the "negative_marking" field.
func NegativeMarkingLTE(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldNegativeMarking, v))
}

// NegativeMarkingIsNil applies the IsNil predicate on the "negative_marking" field.
func NegativeMarkingIsNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIsNull(FieldNegativeMarking))
}

// NegativeMarkingNotNil applies the NotNil predicate on the "negative_marking" field.
func NegativeMarkingNotNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotNull(FieldNegativeMarking))
}

// AiPromptEQ applies the EQ predicate on the "ai_prompt" field.
func AiPromptEQ(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldAiPrompt, v))
}

// AiPromptNEQ applies the NEQ predicate on the "ai_prompt" field.
func AiPromptNEQ(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldAiPrompt, v))
}

// AiPromptIn applies the In predicate on the "ai_prompt" field.
func AiPromptIn(vs ...string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldAiPrompt, vs...))
}

// AiPromptNotIn applies the NotIn predicate on the "ai_prompt" field.
func AiPromptNotIn(vs ...string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldAiPrompt, vs...))
}

// AiPromptGT applies the GT predicate on the "ai_prompt" field.
func AiPromptGT(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldAiPrompt, v))
}

// AiPromptGTE applies the GTE predicate on the "ai_prompt" field.
func AiPromptGTE(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldAiPrompt, v))
}

// AiPromptLT applies the LT predicate on the "ai_prompt" field.
func AiPromptLT(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldAiPrompt, v))
}

// AiPromptLTE applies the LTE predicate on the "ai_prompt" field.
func AiPromptLTE(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldAiPrompt, v))
}

// AiPromptContains applies the Contains predicate on the "ai_prompt" field.
func AiPromptContains(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldContains(FieldAiPrompt, v))
}

// AiPromptHasPrefix applies the HasPrefix predicate on the "ai_prompt" field.
func AiPromptHasPrefix(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldHasPrefix(FieldAiPrompt, v))
}

// AiPromptHasSuffix applies the HasSuffix predicate on the "ai_prompt" field.
func AiPromptHasSuffix(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldHasSuffix(FieldAiPrompt, v))
}

// AiPromptIsNil applies the IsNil predicate on the "ai_prompt" field.
func AiPromptIsNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIsNull(FieldAiPrompt))
}

// AiPromptNotNil applies the NotNil predicate on the "ai_prompt" field.
func AiPromptNotNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotNull(FieldAiPrompt))
}

// AiPromptEqualFold applies the EqualFold predicate on the "ai_prompt" field.
func AiPromptEqualFold(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEqualFold(FieldAiPrompt, v))
}

// AiPromptContainsFold applies the ContainsFold predicate on the "ai_prompt" field.
func AiPromptContainsFold(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldContainsFold(FieldAiPrompt, v))
}

// OtherDetailsIsNil applies the IsNil predicate on the "other_details" field.
func OtherDetailsIsNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIsNull(FieldOtherDetails))
}

// OtherDetailsNotNil applies the NotNil predicate on the "other_details" field.
func OtherDetailsNotNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotNull(FieldOtherDetails))
}

// MaxAttemptsEQ applies the EQ predicate on the "max_attempts" field.
func MaxAttemptsEQ(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldMaxAttempts, v))
}

// MaxAttemptsNEQ applies the NEQ predicate on the "max_attempts" field.
func MaxAttemptsNEQ(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldMaxAttempts, v))
}

// MaxAttemptsIn applies the In predicate on the "max_attempts" field.
func MaxAttemptsIn(vs ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldMaxAttempts, vs...))
}

// MaxAttemptsNotIn applies the NotIn predicate on the "max_attempts" field.
func MaxAttemptsNotIn(vs ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldMaxAttempts, vs...))
}

// MaxAttemptsGT applies the GT predicate on the "max_attempts" field.
func MaxAttemptsGT(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldMaxAttempts, v))
}

// MaxAttemptsGTE applies the GTE predicate on the "max_attempts" field.
func MaxAttemptsGTE(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldMaxAttempts, v))
}

// MaxAttemptsLT applies the LT predicate on the "max_attempts" field.
func MaxAttemptsLT(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldMaxAttempts, v))
}

// MaxAttemptsLTE applies the LTE predicate on the "max_attempts" field.
func MaxAttemptsLTE(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldMaxAttempts, v))
}

// TotalMarksEQ applies the EQ predicate on the "total_marks" field.
func TotalMarksEQ(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldTotalMarks, v))
}

// TotalMarksNEQ applies the NEQ predicate on the "total_marks" field.
func TotalMarksNEQ(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldTotalMarks, v))
}

// TotalMarksIn applies the In predicate on the "total_marks" field.
func TotalMarksIn(vs ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldTotalMarks, vs...))
}

// TotalMarksNotIn applies the NotIn predicate on the "total_marks" field.
func TotalMarksNotIn(vs ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldTotalMarks, vs...))
}

// TotalMarksGT applies the GT predicate on the "total_marks" field.
func TotalMarksGT(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldTotalMarks, v))
}

// TotalMarksGTE applies the GTE predicate on the "total_marks" field.
func TotalMarksGTE(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldTotalMarks, v))
}

// TotalMarksLT applies the LT predicate on the "total_marks" field.
func TotalMarksLT(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldTotalMarks, v))
}

// TotalMarksLTE applies the LTE predicate on the "total_marks" field.
func TotalMarksLTE(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldTotalMarks, v))
}

// TotalMarksIsNil applies the IsNil predicate on the "total_marks" field.
func TotalMarksIsNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIsNull(FieldTotalMarks))
}

// TotalMarksNotNil applies the NotNil predicate on the "total_marks" field.
func TotalMarksNotNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotNull(FieldTotalMarks))
}

// CutoffMarksEQ applies the EQ predicate on the "cutoff_marks" field.
func CutoffMarksEQ(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldCutoffMarks, v))
}

// CutoffMarksNEQ applies the NEQ predicate on the "cutoff_marks" field.
func CutoffMarksNEQ(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldCutoffMarks, v))
}

// CutoffMarksIn applies the In predicate on the "cutoff_marks" field.
func CutoffMarksIn(vs ...float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldCutoffMarks, vs...))
}

// CutoffMarksNotIn applies the NotIn predicate on the "cutoff_marks" field.
func CutoffMarksNotIn(vs ...float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldCutoffMarks, vs...))
}

// CutoffMarksGT applies the GT predicate on the "cutoff_marks" field.
func CutoffMarksGT(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldCutoffMarks, v))
}

// CutoffMarksGTE applies the GTE predicate on the "cutoff_marks" field.
func CutoffMarksGTE(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldCutoffMarks, v))
}

// CutoffMarksLT applies the LT predicate on the "cutoff_marks" field.
func CutoffMarksLT(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldCutoffMarks, v))
}

// CutoffMarksLTE applies the LTE predicate on the "cutoff_marks" field.
func CutoffMarksLTE(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldCutoffMarks, v))
}

// CutoffMarksIsNil applies the IsNil predicate on the "cutoff_marks" field.
func CutoffMarksIsNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIsNull(FieldCutoffMarks))
}

// CutoffMarksNotNil applies the NotNil predicate on the "cutoff_marks" field.
func CutoffMarksNotNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotNull(FieldCutoffMarks))
}

// EvaluationAiPromptEQ applies the EQ predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptEQ(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldEvaluationAiPrompt, v))
}

// EvaluationAiPromptNEQ applies the NEQ predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptNEQ(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldEvaluationAiPrompt, v))
}

// EvaluationAiPromptIn applies the In predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptIn(vs ...string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldEvaluationAiPrompt, vs...))
}

// EvaluationAiPromptNotIn applies the NotIn predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptNotIn(vs ...string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldEvaluationAiPrompt, vs...))
}

// EvaluationAiPromptGT applies the GT predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptGT(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldEvaluationAiPrompt, v))
}

// EvaluationAiPromptGTE applies the GTE predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptGTE(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldEvaluationAiPrompt, v))
}

// EvaluationAiPromptLT applies the LT predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptLT(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldEvaluationAiPrompt, v))
}

// EvaluationAiPromptLTE applies the LTE predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptLTE(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldEvaluationAiPrompt, v))
}

// EvaluationAiPromptContains applies the Contains predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptContains(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldContains(FieldEvaluationAiPrompt, v))
}

// EvaluationAiPromptHasPrefix applies the HasPrefix predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptHasPrefix(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldHasPrefix(FieldEvaluationAiPrompt, v))
}

// EvaluationAiPromptHasSuffix applies the HasSuffix predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptHasSuffix(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldHasSuffix(FieldEvaluationAiPrompt, v))
}

// EvaluationAiPromptIsNil applies the IsNil predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptIsNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIsNull(FieldEvaluationAiPrompt))
}

// EvaluationAiPromptNotNil applies the NotNil predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptNotNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotNull(FieldEvaluationAiPrompt))
}

// EvaluationAiPromptEqualFold applies the EqualFold predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptEqualFold(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEqualFold(FieldEvaluationAiPrompt, v))
}

// EvaluationAiPromptContainsFold applies the ContainsFold predicate on the "evaluation_ai_prompt" field.
func EvaluationAiPromptContainsFold(v string) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldContainsFold(FieldEvaluationAiPrompt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasExam applies the HasEdge predicate on the "exam" edge.
func HasExam() predicate.ExamSetting {
	return predicate.ExamSetting(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ExamTable, ExamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExamWith applies the HasEdge predicate on the "exam" edge with a given conditions (other predicates).
func HasExamWith(preds ...predicate.Exam) predicate.ExamSetting {
	return predicate.ExamSetting(func(s *sql.Selector) {
		step := newExamStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ExamSetting) predicate.ExamSetting {
	return predicate.ExamSetting(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ExamSetting) predicate.ExamSetting {
	return predicate.ExamSetting(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ExamSetting) predicate.ExamSetting {
	return predicate.ExamSetting(sql.NotPredicates(p))
}
