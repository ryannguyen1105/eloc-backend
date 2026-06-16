package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ryannguyen1105/eloc-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	arg := CreateCategoryParams{
		Name: util.RandomNameCategory(),
		Slug: util.RandomSlug(),
	}
	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.Slug, category.Slug)

	require.NotZero(t, category.ID)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategoryByID(t *testing.T) {
	category1 := createRandomCategory(t)
	arg := GetCategoryByIDParams {
		ID: category1.ID,
	}
	category2, err := testQueries.GetCategoryByID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.Equal(t, category1.Slug, category2.Slug)

}

func TestGetCategoryBySlug(t *testing.T) {
	category1 := createRandomCategory(t)
	arg := GetCategoryBySlugParams {
		Slug: category1.Slug,
	}
	category2, err := testQueries.GetCategoryBySlug(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.Equal(t, category1.Slug, category2.Slug)

}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	updateArg := UpdateCategoryParams {
		ID: category1.ID,
		Name: util.RandomNameCategory(),
		Slug: util.RandomSlug(),
	}
	err := testQueries.UpdateCategory(context.Background(), updateArg)
	require.NoError(t, err)

	arg := GetCategoryByIDParams {
		ID: updateArg.ID,
	}
	category2, err := testQueries.GetCategoryByID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, updateArg.ID, category2.ID)
	require.Equal(t, updateArg.Name, category2.Name)
	require.Equal(t, updateArg.Slug, category2.Slug)
}

func TestDeleteCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	deleteArg := DeleteCategoryParams {
		ID: category1.ID,
	}
	err := testQueries.DeleteCategory(context.Background(), deleteArg)
	require.NoError(t, err)

	arg := GetCategoryByIDParams {
		ID: category1.ID,
	}
	category2, err := testQueries.GetCategoryByID(context.Background(), arg)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)

	require.Empty(t, category2)
}

func TestListCategories(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomCategory(t)
	}
	arg := ListCategoriesParams {
		Limit: 5,
		Offset: 0,
	}
	categories, err := testQueries.ListCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, categories, 5)

	for _, category := range categories {
		require.NotEmpty(t, category)
		require.NotEmpty(t, category.ID)
		require.NotEmpty(t, category.Name)
	}
}