// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
)

type Querier interface {
	CreateAuthor(ctx context.Context, name string) (int32, error)
	CreateBanner(ctx context.Context, arg CreateBannerParams) (int32, error)
	CreateBook(ctx context.Context, arg CreateBookParams) (int32, error)
	CreateBookCategory(ctx context.Context, arg CreateBookCategoryParams) (int32, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (int32, error)
	CreateCategoryImage(ctx context.Context, arg CreateCategoryImageParams) (int32, error)
	CreateCountry(ctx context.Context, arg CreateCountryParams) (int32, error)
	CreateOffer(ctx context.Context, arg CreateOfferParams) (int32, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (int32, error)
	CreatePayment(ctx context.Context, arg CreatePaymentParams) (int32, error)
	CreatePermission(ctx context.Context, arg CreatePermissionParams) (int32, error)
	CreateRole(ctx context.Context, arg CreateRoleParams) (int32, error)
	CreateRolePermission(ctx context.Context, arg CreateRolePermissionParams) (int32, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (int32, error)
	DeleteAuthor(ctx context.Context, id int32) (int32, error)
	DeleteBanner(ctx context.Context, id int32) (int32, error)
	DeleteBook(ctx context.Context, id int32) (int32, error)
	DeleteBookCategory(ctx context.Context, id int32) (int32, error)
	DeleteCategory(ctx context.Context, id int32) (int32, error)
	DeleteCategoryImage(ctx context.Context, id int32) (int32, error)
	DeleteCountry(ctx context.Context, id int32) (int32, error)
	DeleteOffer(ctx context.Context, id int32) (int32, error)
	DeleteOrder(ctx context.Context, id int32) (int32, error)
	DeletePayment(ctx context.Context, id int32) (int32, error)
	DeletePermission(ctx context.Context, id int32) (int32, error)
	DeleteRole(ctx context.Context, id int32) (int32, error)
	DeleteRolePermission(ctx context.Context, id int32) (int32, error)
	DeleteUser(ctx context.Context, id int32) (int32, error)
	GetAllAuthors(ctx context.Context) ([]Author, error)
	GetAllBanners(ctx context.Context) ([]Banner, error)
	GetAllBookCategories(ctx context.Context) ([]BooksCategory, error)
	GetAllBooks(ctx context.Context) ([]Book, error)
	GetAllCategories(ctx context.Context) ([]Category, error)
	GetAllCategoryImages(ctx context.Context) ([]CategoriesImage, error)
	GetAllCountries(ctx context.Context) ([]Country, error)
	GetAllOffers(ctx context.Context) ([]Offer, error)
	GetAllOrders(ctx context.Context) ([]Order, error)
	GetAllPayments(ctx context.Context) ([]Payment, error)
	GetAllPermissions(ctx context.Context) ([]Permission, error)
	GetAllRolePermissions(ctx context.Context) ([]RolesPermission, error)
	GetAllRoles(ctx context.Context) ([]Role, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	GetAuthor(ctx context.Context, id int32) (Author, error)
	GetBanner(ctx context.Context, id int32) (Banner, error)
	GetBook(ctx context.Context, id int32) (Book, error)
	GetBookCategory(ctx context.Context, id int32) (BooksCategory, error)
	GetCategory(ctx context.Context, id int32) (Category, error)
	GetCategoryImage(ctx context.Context, id int32) (CategoriesImage, error)
	GetCountry(ctx context.Context, id int32) (Country, error)
	GetOffer(ctx context.Context, id int32) (Offer, error)
	GetOrder(ctx context.Context, id int32) (Order, error)
	GetPayment(ctx context.Context, id int32) (Payment, error)
	GetPermission(ctx context.Context, id int32) (Permission, error)
	GetRole(ctx context.Context, id int32) (Role, error)
	GetRolePermission(ctx context.Context, id int32) (RolesPermission, error)
	GetUser(ctx context.Context, id int32) (User, error)
	UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (int32, error)
	UpdateBanner(ctx context.Context, arg UpdateBannerParams) (int32, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams) (int32, error)
	UpdateBookCategory(ctx context.Context, arg UpdateBookCategoryParams) (int32, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (int32, error)
	UpdateCategoryImage(ctx context.Context, arg UpdateCategoryImageParams) (int32, error)
	UpdateCountry(ctx context.Context, arg UpdateCountryParams) (int32, error)
	UpdateOffer(ctx context.Context, arg UpdateOfferParams) (int32, error)
	UpdateOrder(ctx context.Context, arg UpdateOrderParams) (int32, error)
	UpdatePayment(ctx context.Context, arg UpdatePaymentParams) (int32, error)
	UpdatePermission(ctx context.Context, arg UpdatePermissionParams) (int32, error)
	UpdateRole(ctx context.Context, arg UpdateRoleParams) (int32, error)
	UpdateRolePermission(ctx context.Context, arg UpdateRolePermissionParams) (int32, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (int32, error)
}

//*Queries implements the interface Querier at compile time.
var _ Querier = (*Queries)(nil)
