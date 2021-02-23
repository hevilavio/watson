package actions

import (
	"log"
	"watson/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
	// "net/url"
	// "strconv"
)

type UserResource struct{}

// lista os usuário do banco
func (ur UserResource) List(c buffalo.Context) error {

	user := models.User{}
	id := c.Param("id")

	err := models.DB.Find(&user, id)

	if err != nil {
		log.Println(err)
		return c.Render(200, r.JSON("error while fetching user"))
	}

	return c.Render(200, r.JSON(user))
}

// cria um novo usuário
func (ur UserResource) Create(c buffalo.Context) error {

	id, _ := uuid.NewV4()

	// new User
	user := &models.User{
		ID: id,
	}

	// add in database
	err := models.DB.Create(user)

	if err != nil {
		app.Logger.Fatal(err)
	}

	return c.Render(201, r.JSON(user))
}

/*
binders

func MyAction(c buffalo.Context) error {
  u := &User{}
  if err := c.Bind(u); err != nil {
    return err
  }
  u.Name // "Ringo"
  u.Email // "ringo@beatles.com"
  u.Password // ""
  // do more work
}
*/

func (ur UserResource) Hello(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}
