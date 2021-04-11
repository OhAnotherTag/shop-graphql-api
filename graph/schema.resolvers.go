package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/OhAnotherTag/shop-gql-api/graph/generated"
	"github.com/OhAnotherTag/shop-gql-api/graph/model"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	db := r.Database()

	product := &model.Product{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		CategoryID:  input.CategoryID,
	}

	if err := db.Create(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	db := r.Database()

	category := &model.Category{
		Name:     input.Name,
		Products: []*model.Product{},
	}

	if err := db.Create(&category).Error; err != nil {
		return nil, err
	}

	fmt.Println(category)

	return category, nil
}

func (r *queryResolver) Category(ctx context.Context, id int) (*model.Category, error) {
	var category *model.Category
	var products []*model.Product

	db := r.Database()

	preloads := GetPreloads(ctx)
	includeProducts := false

	for _, field := range preloads {
		if field == "products" {
			includeProducts = true
		}
	}

	fmt.Println(includeProducts)

	if includeProducts {
		db.First(&category, id)
		db.Model(&category).Association("Products").Find(&products)

		category.Products = products
	} else {
		db.First(&category, id)
	}

	fmt.Println(category)

	return category, nil
}

func (r *queryResolver) Product(ctx context.Context, id int) (*model.Product, error) {
	db := r.Database()
	var product *model.Product

	db.First(&product, id)
	return product, nil
}

func (r *queryResolver) Inventory(ctx context.Context, filter *string, skip *int, take *int, orderBy *model.ProductOrderByInput) (*model.Inventory, error) {
	db := r.Database()

	var products []*model.Product
	var inventory *model.Inventory
	var title string
	var createdAt string

	if orderBy != nil {
		title = fmt.Sprintf("title %s", orderBy.Title.String())
		createdAt = fmt.Sprintf("created_at %s", orderBy.Title.String())
	} else {
		title = "title desc"
		createdAt = "created_at desc"
	}

	var count int64

	db.Find(&model.Product{}).Count(&count)

	if *filter != "" || *take == -1 {
		db.Offset(*skip).Limit(*take).Where("title LIKE ?", *filter+"%").Order(createdAt).Order(title).Find(&products)

		inventory = &model.Inventory{
			Products: products,
			Count:    int(count),
		}
	} else {
		inventory = &model.Inventory{
			Products: []*model.Product{},
			Count:    0,
		}
	}

	return inventory, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	db := r.Database()

	var categories []*model.Category

	db.Preload("Products").Find(&categories)

	return categories, nil
}

func (r *queryResolver) Cart(ctx context.Context, ids []int) ([]*model.Product, error) {
	db := r.Database()
	var cart []*model.Product

	if len(ids) == 0 {
		return []*model.Product{}, nil
	}

	db.Find(&cart, ids)

	return cart, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}
func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}
func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
