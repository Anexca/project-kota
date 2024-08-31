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
	}
	// CachedQuestionMetaDataTable holds the schema information for the "cached_question_meta_data" table.
	CachedQuestionMetaDataTable = &schema.Table{
		Name:       "cached_question_meta_data",
		Columns:    CachedQuestionMetaDataColumns,
		PrimaryKey: []*schema.Column{CachedQuestionMetaDataColumns[0]},
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
	// ExamCachedQuestionMetadataColumns holds the columns for the "exam_cached_question_metadata" table.
	ExamCachedQuestionMetadataColumns = []*schema.Column{
		{Name: "exam_id", Type: field.TypeInt},
		{Name: "cached_question_meta_data_id", Type: field.TypeInt},
	}
	// ExamCachedQuestionMetadataTable holds the schema information for the "exam_cached_question_metadata" table.
	ExamCachedQuestionMetadataTable = &schema.Table{
		Name:       "exam_cached_question_metadata",
		Columns:    ExamCachedQuestionMetadataColumns,
		PrimaryKey: []*schema.Column{ExamCachedQuestionMetadataColumns[0], ExamCachedQuestionMetadataColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "exam_cached_question_metadata_exam_id",
				Columns:    []*schema.Column{ExamCachedQuestionMetadataColumns[0]},
				RefColumns: []*schema.Column{ExamsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "exam_cached_question_metadata_cached_question_meta_data_id",
				Columns:    []*schema.Column{ExamCachedQuestionMetadataColumns[1]},
				RefColumns: []*schema.Column{CachedQuestionMetaDataColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CachedQuestionMetaDataTable,
		ExamsTable,
		ExamCategoriesTable,
		ExamSettingsTable,
		ExamCachedQuestionMetadataTable,
	}
)

func init() {
	ExamsTable.ForeignKeys[0].RefTable = ExamCategoriesTable
	ExamSettingsTable.ForeignKeys[0].RefTable = ExamsTable
	ExamCachedQuestionMetadataTable.ForeignKeys[0].RefTable = ExamsTable
	ExamCachedQuestionMetadataTable.ForeignKeys[1].RefTable = CachedQuestionMetaDataTable
}
