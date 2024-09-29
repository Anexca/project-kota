package repositories_test

import (
	"context"
	"testing"

	"common/ent"
	"common/repositories"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *ent.Client {
	// Create an in-memory SQLite database for testing
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)

	// Run the migration to create the schema
	err = client.Schema.Create(context.Background())
	require.NoError(t, err)

	return client
}

func TestUserRepository_Get(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	ctx := context.Background()

	// Manually generate a UUID for the user
	userId := uuid.New()

	// Create a user in the test database
	user, err := client.User.Create().
		SetID(userId). // Set the UUID manually
		SetEmail("test@example.com").
		SetFirstName("John").
		SetLastName("Doe").
		Save(ctx)
	require.NoError(t, err)

	// Initialize the repository
	userRepo := repositories.NewUserRepository(client)

	// Call the Get method with the created user's ID
	retrievedUser, err := userRepo.Get(ctx, user.ID.String())

	// Assert no errors and verify the retrieved user
	require.NoError(t, err)
	require.NotNil(t, retrievedUser)
	require.Equal(t, "John", retrievedUser.FirstName)
	require.Equal(t, "Doe", retrievedUser.LastName)
	require.Equal(t, "test@example.com", retrievedUser.Email)
}

func TestUserRepository_Get_NonExistentUser(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	ctx := context.Background()

	// Initialize the repository
	userRepo := repositories.NewUserRepository(client)

	// Call the Get method with a non-existent user ID
	nonExistentID := uuid.New().String()
	retrievedUser, err := userRepo.Get(ctx, nonExistentID)

	// Assert that an error is returned and user is nil
	require.Error(t, err)
	require.Nil(t, retrievedUser)
}

func TestUserRepository_Get_EmptyUserID(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	ctx := context.Background()

	// Initialize the repository
	userRepo := repositories.NewUserRepository(client)

	// Call the Get method with an empty user ID
	retrievedUser, err := userRepo.Get(ctx, "")

	// Assert that an error is returned and user is nil
	require.Error(t, err)
	require.Nil(t, retrievedUser)
}

func TestUserRepository_CreateMultipleUsers(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	ctx := context.Background()

	// Manually generate UUIDs for the users
	user1ID := uuid.New()
	user2ID := uuid.New()

	// Create multiple users with manually generated UUIDs
	user1, err := client.User.Create().
		SetID(user1ID). // Set UUID for user1
		SetEmail("user1@example.com").
		SetFirstName("Alice").
		SetLastName("Smith").
		Save(ctx)
	require.NoError(t, err)

	user2, err := client.User.Create().
		SetID(user2ID). // Set UUID for user2
		SetEmail("user2@example.com").
		SetFirstName("Bob").
		SetLastName("Jones").
		Save(ctx)
	require.NoError(t, err)

	// Initialize the repository
	userRepo := repositories.NewUserRepository(client)

	// Verify that both users are correctly retrieved
	retrievedUser1, err := userRepo.Get(ctx, user1.ID.String())
	require.NoError(t, err)
	require.NotNil(t, retrievedUser1)
	require.Equal(t, "Alice", retrievedUser1.FirstName)
	require.Equal(t, "Smith", retrievedUser1.LastName)
	require.Equal(t, "user1@example.com", retrievedUser1.Email)

	retrievedUser2, err := userRepo.Get(ctx, user2.ID.String())
	require.NoError(t, err)
	require.NotNil(t, retrievedUser2)
	require.Equal(t, "Bob", retrievedUser2.FirstName)
	require.Equal(t, "Jones", retrievedUser2.LastName)
	require.Equal(t, "user2@example.com", retrievedUser2.Email)
}

func TestUserRepository_UpdateUser(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	ctx := context.Background()

	// Manually generate a UUID for the user
	userId := uuid.New()

	// Create a user in the test database
	user, err := client.User.Create().
		SetID(userId). // Set the UUID manually
		SetEmail("test@example.com").
		SetFirstName("John").
		SetLastName("Doe").
		Save(ctx)
	require.NoError(t, err)

	// Update the user's information
	updatedUser, err := client.User.UpdateOneID(user.ID).
		SetFirstName("Jane").
		SetLastName("Doe").
		SetEmail("jane.doe@example.com").
		Save(ctx)
	require.NoError(t, err)

	// Initialize the repository
	userRepo := repositories.NewUserRepository(client)

	// Verify that the updated user is retrieved
	retrievedUser, err := userRepo.Get(ctx, updatedUser.ID.String())
	require.NoError(t, err)
	require.NotNil(t, retrievedUser)
	require.Equal(t, "Jane", retrievedUser.FirstName)
	require.Equal(t, "Doe", retrievedUser.LastName)
	require.Equal(t, "jane.doe@example.com", retrievedUser.Email)
}
