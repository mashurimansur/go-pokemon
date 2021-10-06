package controller

import (
	"go-pokemon/entity"
	"go-pokemon/service"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	PokemonService service.PokemonService
}

func NewPokemonController(pokemonService *service.PokemonService) PokemonController {
	return PokemonController{PokemonService: *pokemonService}
}

func (controller *PokemonController) Route(e *gin.Engine) {
	// WEB
	e.GET("/", controller.listPokemon)
	e.GET("/detail/:name", controller.detailPokemon)
	e.POST("/pokemon", controller.catchPokemon)
	e.GET("/my-pokemon", controller.myPokemon)
	e.GET("/my-pokemon/remove/:id", controller.removePokemon)
	//API
	e.GET("/api/random-catch", controller.randomCatch)
	e.GET("/api/prime-number/:number", controller.primeNumber)
	e.GET("/api/fibonacci/:name", controller.fibonacciName)
	e.GET("/api/change-name/:id", controller.updateNamePokemon)

}

func (controller *PokemonController) listPokemon(c *gin.Context) {
	list := controller.PokemonService.ListPokemon()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"host":    c.Request.URL.Host,
		"pokemon": list,
	})
}

func (controller *PokemonController) detailPokemon(c *gin.Context) {
	name := c.Param("name")
	result := controller.PokemonService.DetailPokemon(name)

	c.HTML(http.StatusOK, "detail.html", gin.H{
		"host":   c.Request.URL.Host,
		"result": result,
	})
}

func (controller *PokemonController) catchPokemon(c *gin.Context) {
	myPokemon := entity.MyPokemon{
		Name:     c.PostForm("name"),
		Nickname: c.PostForm("nickname"),
		Image:    c.PostForm("image"),
	}

	err := controller.PokemonService.CatchPokemon(&myPokemon)
	if err != nil {
		log.Fatal(err)
	}

	c.Redirect(http.StatusMovedPermanently, "/my-pokemon")
}

func (controller *PokemonController) myPokemon(c *gin.Context) {
	result, err := controller.PokemonService.MyPokemon()
	if err != nil {
		log.Fatal(err)
	}

	c.HTML(http.StatusOK, "mypokemon.html", gin.H{
		"host":   c.Request.URL.Host,
		"result": result,
	})
}

func (controller *PokemonController) removePokemon(c *gin.Context) {
	id := c.Param("id")

	err := controller.PokemonService.ReleasePokemon(id)
	if err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/my-pokemon")
}

func (controller *PokemonController) randomCatch(c *gin.Context) {
	now := time.Now()

	second := now.Second()
	var status bool
	if second%2 == 0 {
		status = true
	} else {
		status = false
	}

	c.JSON(200, gin.H{"status": status})
}

func (controller *PokemonController) primeNumber(c *gin.Context) {
	number := c.Param("number")
	num, _ := strconv.Atoi(number)
	prime := true

	if num < 2 {
		prime = false
	}

	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			prime = false
		}
	}

	c.JSON(200, gin.H{"status": prime})
}

func (controller *PokemonController) fibonacciName(c *gin.Context) {
	name := c.Param("name")
	res := controller.PokemonService.ChangeNamePokemon(name)
	c.JSON(200, gin.H{"nickname": res})
}

func (controller *PokemonController) updateNamePokemon(c *gin.Context) {
	id := c.Param("id")

	myPokemon, _ := controller.PokemonService.FindByID(id)
	myPokemon.Nickname = controller.PokemonService.ChangeNamePokemon(myPokemon.Nickname)

	errUpdate := controller.PokemonService.UpdateName(&myPokemon)
	if errUpdate != nil {
		c.JSON(400, errUpdate.Error())
		return
	}
	c.JSON(200, myPokemon)
}
