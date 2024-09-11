// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CachedExamsColumns holds the columns for the "cached_exams" table.
	CachedExamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "cache_uid", Type: field.TypeString, Unique: true},
		{Name: "is_used", Type: field.TypeBool, Default: false},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "exam_cached_exam", Type: field.TypeInt},
	}
	// CachedExamsTable holds the schema information for the "cached_exams" table.
	CachedExamsTable = &schema.Table{
		Name:       "cached_exams",
		Columns:    CachedExamsColumns,
		PrimaryKey: []*schema.Column{CachedExamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cached_exams_exams_cached_exam",
				Columns:    []*schema.Column{CachedExamsColumns[6]},
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
	// ExamAssesmentsColumns holds the columns for the "exam_assesments" table.
	ExamAssesmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "completed_seconds", Type: field.TypeInt},
		{Name: "raw_assesment_data", Type: field.TypeJSON, Nullable: true, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "raw_user_submission", Type: field.TypeJSON, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"COMPLETED", "REJECTED", "PENDING"}, SchemaType: map[string]string{"postgres": "status"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "exam_attempt_assesment", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// ExamAssesmentsTable holds the schema information for the "exam_assesments" table.
	ExamAssesmentsTable = &schema.Table{
		Name:       "exam_assesments",
		Columns:    ExamAssesmentsColumns,
		PrimaryKey: []*schema.Column{ExamAssesmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "exam_assesments_exam_attempts_assesment",
				Columns:    []*schema.Column{ExamAssesmentsColumns[7]},
				RefColumns: []*schema.Column{ExamAttemptsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ExamAttemptsColumns holds the columns for the "exam_attempts" table.
	ExamAttemptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "attempt_number", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "generated_exam_attempts", Type: field.TypeInt, Nullable: true},
		{Name: "user_attempts", Type: field.TypeUUID, Nullable: true},
	}
	// ExamAttemptsTable holds the schema information for the "exam_attempts" table.
	ExamAttemptsTable = &schema.Table{
		Name:       "exam_attempts",
		Columns:    ExamAttemptsColumns,
		PrimaryKey: []*schema.Column{ExamAttemptsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "exam_attempts_generated_exams_attempts",
				Columns:    []*schema.Column{ExamAttemptsColumns[4]},
				RefColumns: []*schema.Column{GeneratedExamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "exam_attempts_users_attempts",
				Columns:    []*schema.Column{ExamAttemptsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
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
		{Name: "duration_seconds", Type: field.TypeInt},
		{Name: "negative_marking", Type: field.TypeFloat64, Nullable: true},
		{Name: "ai_prompt", Type: field.TypeString, Nullable: true},
		{Name: "other_details", Type: field.TypeJSON, Nullable: true, SchemaType: map[string]string{"postgres": "json"}},
		{Name: "max_attempts", Type: field.TypeInt, Default: 2},
		{Name: "evaluation_ai_prompt", Type: field.TypeString, Nullable: true},
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
				Columns:    []*schema.Column{ExamSettingsColumns[10]},
				RefColumns: []*schema.Column{ExamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GeneratedExamsColumns holds the columns for the "generated_exams" table.
	GeneratedExamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "raw_exam_data", Type: field.TypeJSON, Nullable: true, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "exam_generatedexams", Type: field.TypeInt, Nullable: true},
	}
	// GeneratedExamsTable holds the schema information for the "generated_exams" table.
	GeneratedExamsTable = &schema.Table{
		Name:       "generated_exams",
		Columns:    GeneratedExamsColumns,
		PrimaryKey: []*schema.Column{GeneratedExamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "generated_exams_exams_generatedexams",
				Columns:    []*schema.Column{GeneratedExamsColumns[5]},
				RefColumns: []*schema.Column{ExamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amount", Type: field.TypeInt},
		{Name: "payment_date", Type: field.TypeTime},
		{Name: "payment_status", Type: field.TypeEnum, Enums: []string{"CREATED", "AUTHORIZED", "CAPTURED", "FAILED", "REFUNDED", "PARTIALLY_REFUNDED", "PENDING", "PROCESSING", "CANCELLED", "DISPUTED"}},
		{Name: "payment_method", Type: field.TypeString},
		{Name: "payment_payment_id", Type: field.TypeString, Unique: true},
		{Name: "receipt_id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_payments", Type: field.TypeUUID, Nullable: true},
		{Name: "user_subscription_payments", Type: field.TypeInt, Nullable: true},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:       "payments",
		Columns:    PaymentsColumns,
		PrimaryKey: []*schema.Column{PaymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payments_users_payments",
				Columns:    []*schema.Column{PaymentsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "payments_user_subscriptions_payments",
				Columns:    []*schema.Column{PaymentsColumns[10]},
				RefColumns: []*schema.Column{UserSubscriptionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "provider_plan_id", Type: field.TypeString},
		{Name: "price", Type: field.TypeInt},
		{Name: "duration_in_months", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool},
		{Name: "name", Type: field.TypeString},
		{Name: "raw_subscription_data", Type: field.TypeJSON, Nullable: true, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
	}
	// SubscriptionExamsColumns holds the columns for the "subscription_exams" table.
	SubscriptionExamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "exam_subscriptions", Type: field.TypeInt, Nullable: true},
		{Name: "subscription_exams", Type: field.TypeInt, Nullable: true},
	}
	// SubscriptionExamsTable holds the schema information for the "subscription_exams" table.
	SubscriptionExamsTable = &schema.Table{
		Name:       "subscription_exams",
		Columns:    SubscriptionExamsColumns,
		PrimaryKey: []*schema.Column{SubscriptionExamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subscription_exams_exams_subscriptions",
				Columns:    []*schema.Column{SubscriptionExamsColumns[3]},
				RefColumns: []*schema.Column{ExamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "subscription_exams_subscriptions_exams",
				Columns:    []*schema.Column{SubscriptionExamsColumns[4]},
				RefColumns: []*schema.Column{SubscriptionsColumns[0]},
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
		{Name: "phone_number", Type: field.TypeString, Nullable: true},
		{Name: "payment_provider_customer_id", Type: field.TypeString, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserSubscriptionsColumns holds the columns for the "user_subscriptions" table.
	UserSubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "is_active", Type: field.TypeBool},
		{Name: "start_date", Type: field.TypeTime, Nullable: true},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "provider_subscription_id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "subscription_user_subscriptions", Type: field.TypeInt, Nullable: true},
		{Name: "user_subscriptions", Type: field.TypeUUID, Nullable: true},
	}
	// UserSubscriptionsTable holds the schema information for the "user_subscriptions" table.
	UserSubscriptionsTable = &schema.Table{
		Name:       "user_subscriptions",
		Columns:    UserSubscriptionsColumns,
		PrimaryKey: []*schema.Column{UserSubscriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_subscriptions_subscriptions_user_subscriptions",
				Columns:    []*schema.Column{UserSubscriptionsColumns[7]},
				RefColumns: []*schema.Column{SubscriptionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "user_subscriptions_users_subscriptions",
				Columns:    []*schema.Column{UserSubscriptionsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CachedExamsTable,
		ExamsTable,
		ExamAssesmentsTable,
		ExamAttemptsTable,
		ExamCategoriesTable,
		ExamSettingsTable,
		GeneratedExamsTable,
		PaymentsTable,
		SubscriptionsTable,
		SubscriptionExamsTable,
		UsersTable,
		UserSubscriptionsTable,
	}
)

func init() {
	CachedExamsTable.ForeignKeys[0].RefTable = ExamsTable
	ExamsTable.ForeignKeys[0].RefTable = ExamCategoriesTable
	ExamAssesmentsTable.ForeignKeys[0].RefTable = ExamAttemptsTable
	ExamAttemptsTable.ForeignKeys[0].RefTable = GeneratedExamsTable
	ExamAttemptsTable.ForeignKeys[1].RefTable = UsersTable
	ExamSettingsTable.ForeignKeys[0].RefTable = ExamsTable
	GeneratedExamsTable.ForeignKeys[0].RefTable = ExamsTable
	PaymentsTable.ForeignKeys[0].RefTable = UsersTable
	PaymentsTable.ForeignKeys[1].RefTable = UserSubscriptionsTable
	SubscriptionExamsTable.ForeignKeys[0].RefTable = ExamsTable
	SubscriptionExamsTable.ForeignKeys[1].RefTable = SubscriptionsTable
	UserSubscriptionsTable.ForeignKeys[0].RefTable = SubscriptionsTable
	UserSubscriptionsTable.ForeignKeys[1].RefTable = UsersTable
}
