// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CachedQuestionMetaDataColumns holds the columns for the "cached_question_meta_data" table.
	CachedQuestionMetaDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "cache_uid", Type: field.TypeString, Unique: true},
		{Name: "is_used", Type: field.TypeBool, Default: false},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "exam_cached_question_metadata", Type: field.TypeInt},
	}
	// CachedQuestionMetaDataTable holds the schema information for the "cached_question_meta_data" table.
	CachedQuestionMetaDataTable = &schema.Table{
		Name:       "cached_question_meta_data",
		Columns:    CachedQuestionMetaDataColumns,
		PrimaryKey: []*schema.Column{CachedQuestionMetaDataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cached_question_meta_data_exams_cached_question_metadata",
				Columns:    []*schema.Column{CachedQuestionMetaDataColumns[6]},
				RefColumns: []*schema.Column{ExamsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ExamsColumns holds the columns for the "exams" table.
	ExamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "exam_category_exams", Type: field.TypeInt, Nullable: true},
	}
	// ExamsTable holds the schema information for the "exams" table.
	ExamsTable = &schema.Table{
		Name:       "exams",
		Columns:    ExamsColumns,
		PrimaryKey: []*schema.Column{ExamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "exams_exam_categories_exams",
				Columns:    []*schema.Column{ExamsColumns[6]},
				RefColumns: []*schema.Column{ExamCategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ExamCategoriesColumns holds the columns for the "exam_categories" table.
	ExamCategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ExamCategoriesTable holds the schema information for the "exam_categories" table.
	ExamCategoriesTable = &schema.Table{
		Name:       "exam_categories",
		Columns:    ExamCategoriesColumns,
		PrimaryKey: []*schema.Column{ExamCategoriesColumns[0]},
	}
	// ExamSettingsColumns holds the columns for the "exam_settings" table.
	ExamSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "number_of_questions", Type: field.TypeInt},
		{Name: "duration_minutes", Type: field.TypeInt},
		{Name: "negative_marking", Type: field.TypeFloat64, Nullable: true},
		{Name: "ai_prompt", Type: field.TypeString, Nullable: true},
		{Name: "other_details", Type: field.TypeJSON, Nullable: true, SchemaType: map[string]string{"postgres": "json"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "exam_setting", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// ExamSettingsTable holds the schema information for the "exam_settings" table.
	ExamSettingsTable = &schema.Table{
		Name:       "exam_settings",
		Columns:    ExamSettingsColumns,
		PrimaryKey: []*schema.Column{ExamSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "exam_settings_exams_setting",
				Columns:    []*schema.Column{ExamSettingsColumns[8]},
				RefColumns: []*schema.Column{ExamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "first_name", Type: field.TypeString, Nullable: true},
		{Name: "last_name", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CachedQuestionMetaDataTable,
		ExamsTable,
		ExamCategoriesTable,
		ExamSettingsTable,
		UsersTable,
	}
)

func init() {
	CachedQuestionMetaDataTable.ForeignKeys[0].RefTable = ExamsTable
	ExamsTable.ForeignKeys[0].RefTable = ExamCategoriesTable
	ExamSettingsTable.ForeignKeys[0].RefTable = ExamsTable
}
