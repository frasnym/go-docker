package http

import (
	"github.com/frasnym/go-docker/errors"
	"github.com/frasnym/go-docker/models"
	"github.com/frasnym/go-docker/user"
	"github.com/kataras/iris/v12"
)

// UserHandlerIris  represent the httphandler for user with iris
type UserHandlerIris struct {
	UUsecase user.Usecase
}

// NewUserHandler will initialize the user/ resources endpoint
func NewUserHandler(r iris.Party, us user.Usecase) {
	handler := &UserHandlerIris{
		UUsecase: us,
	}

	r.Post("/users", handler.SignUp)
	// r.Get("/", handler.Fetchuser)
	// r.Get("/users/{id:int64}", handler.GetByID)
	// r.Delete("/users/{id:int64}", handler.Delete)
}

// Store will store the user by given request body
func (userHandler *UserHandlerIris) SignUp(ctx iris.Context) {
	var user models.User

	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StatusCode(iris.StatusUnprocessableEntity)
		ctx.JSON(err.Error())
		return
	}

	if err := userHandler.UUsecase.Validate(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(errors.ServiceError{Message: err.Error()})
		return
	}

	// c := ctx.Request().Context()
	// if c == nil {
	// 	c = context.Background()
	// }

	// err = a.AUsecase.Store(c, &user)

	// if err != nil {
	// 	handleError(ctx, err)
	// 	return
	// }

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(user)
}
