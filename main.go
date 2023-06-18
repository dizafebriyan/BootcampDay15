package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"profile/connection"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// nama dari structnya adalah Blog
type Blog struct {
	ID                int
	Title             string
	Content           string
	StartDate         time.Time
	StartDateFormator string
	EndDate           time.Time
	EndDateFormator   string
	Technologies      string
	// box1      string
	// box2      string
	// box3      string
	// box4      string
	PostDate time.Time
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

// type statuslogin struct {
// 	isLogin bool
// 	Name    string
// }

// var sessionLogin = statuslogin{}

// var dataProject = []Blog{}

func main() {
	connection.DatabaseConnect()
	// e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World!")
	// })

	// e.Logger.Fatal(e.Start("localhost:5000"))
	e := echo.New()

	e.Static("/public", "public")

	//to use sessions using echo
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))
	//Routing

	//get
	e.GET("/", home)
	e.GET("/myproject", myproject)
	e.GET("/contact", contact)
	e.GET("/testimonial", testimonial)
	e.GET("/myproject/:id", blogDetail)
	e.GET("project-edit/:id", editProject)
	e.GET("/form-login", formLogin)
	e.GET("/form-register", formRegister)
	e.GET("/logout-button", logoutButton)

	//post
	e.POST("project-edit-post/:id", postEditProject)
	e.POST("/addmyproject", addmyproject)
	e.POST("/project-delete/:id", deleteproject)

	//login & register
	e.POST("/login", login)
	// e.POST("/logout", logout)

	e.Logger.Fatal(e.Start("localhost:5500"))

}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id_serial,title,description,start_date,end_date,technologies,postdate FROM tb_projects")

	var result []Blog
	for data.Next() {
		var each = Blog{}
		err := data.Scan(&each.ID, &each.Title, &each.Content, &each.StartDate, &each.EndDate, &each.Technologies, &each.PostDate)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		each.EndDateFormator = each.EndDate.Format("02 January 2006")
		each.StartDateFormator = each.StartDate.Format("02 January 2006")

		result = append(result, each)
	}

	sess, _ := session.Get("session", c)

	var tmpl, err = template.ParseFiles("views/index.html")

	// data := map[string]interface{}{
	// 	"login": true,
	// }

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{}{
		"project":         result,
		"FlashStatus":     sess.Values["status"],
		"FlashMessage":    sess.Values["message"],
		"FlashName":       sess.Values["name"],
		"buttonIndicator": sess.Values["buttonIndicator"],
		// "FlashName":   sessionLogin.Name,
		// "FlashStatus": sessionLogin.isLogin,
	}
	delete(sess.Values, "buttonIndicator")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), projects)
}

func myproject(c echo.Context) error {

	// data := map[string]interface{}{
	// 	"login": true,
	// }
	sess, _ := session.Get("session", c)
	var tmpl, err = template.ParseFiles("views/form-project.html")

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{}{
		"FlashStatus":     sess.Values["status"],
		"FlashMessage":    sess.Values["message"],
		"FlashName":       sess.Values["name"],
		"buttonIndicator": sess.Values["buttonIndicator"],
		// "FlashName":   sessionLogin.Name,
		// "FlashStatus": sessionLogin.isLogin,
	}

	return tmpl.Execute(c.Response(), projects)
}

func contact(c echo.Context) error {
	// data := map[string]interface{}{
	// 	"login": true,
	// }
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	sess, _ := session.Get("session", c)

	projects := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"FlashName":    sess.Values["name"],
		// "FlashName":   sessionLogin.Name,
		// "FlashStatus": sessionLogin.isLogin,
	}

	return tmpl.Execute(c.Response(), projects)
}

func testimonial(c echo.Context) error {

	// data := map[string]interface{}{
	// 	"login": true,
	// }
	sess, _ := session.Get("session", c)
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{}{
		"FlashStatus":     sess.Values["status"],
		"FlashMessage":    sess.Values["message"],
		"FlashName":       sess.Values["name"],
		"buttonIndicator": sess.Values["buttonIndicator"],
		// "FlashName":   sessionLogin.Name,
		// "FlashStatus": sessionLogin.isLogin,
	}
	return tmpl.Execute(c.Response(), projects)
}

func formLogin(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-login.html")
	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, _ := session.Get("session", c)

	projects := map[string]interface{}{
		"FlashStatus":     sess.Values["status"],
		"FlashMessage":    sess.Values["message"],
		"FlashName":       sess.Values["name"],
		"buttonIndicator": sess.Values["buttonIndicator"],
	}
	delete(sess.Values, "buttonIndicator")
	delete(sess.Values, "message")
	sess.Save(c.Request(), c.Response())

	flashStatus := projects["FlashStatus"] //fungsi untuk memaksa jika telah login maka langsung ke generate ke home jika mengakses formlogin
	println(flashStatus)
	if flashStatus == true {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	return tmpl.Execute(c.Response(), projects)
}

func formRegister(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-register.html")

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := c.FormValue("inputEmail")
	// password := c.FormValue("inputPassowrd")

	user := User{}

	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		sess, _ := session.Get("session", c)
		sess.Values["message"] = "Login Failed"
		sess.Values["status"] = false
		sess.Values["buttonIndicator"] = false
		return redirectWithMessage(c, "Email Incorrect", false, "/form-login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800
	sess.Values["message"] = "Login Success"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true
	sess.Values["buttonIndicator"] = true
	sess.Save(c.Request(), c.Response())

	// sessionLogin.Name = sess.Values["name"].(string)
	// sessionLogin.isLogin = sess.Values["status"].(bool)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func logoutButton(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-login.html")

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	println("Logout Successfull")

	return tmpl.Execute(c.Response(), nil)

}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}

func blogDetail(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	// data := map[string]interface{}{
	// 	"id":      id,
	// 	"title":   "Duino-Coin, Sebagai pembelajaran Mining",
	// 	"content": "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Neque, eveniet quia in asperiores ducimus, quam maiores excepturi at voluptate quos ipsum? Labore quis dolore id optio nemo placeat voluptas sunt! Lorem ipsum dolor sit amet consectetur adipisicing elit. Delectus ratione, pariatur a quo quis quas in atque magnam, eveniet dolores, animi excepturi quidem dignissimos voluptates expedita repellendus aliquid quasi sed.Lorem ipsum dolor sit, amet consectetur adipisicing elit. Est harum illo vero! Quibusdam esse quidem mollitia amet necessitatibus voluptates inventore sapiente eius expedita. Qui magni placeat error? Libero, debitis minus. Lorem ipsum dolor sit amet consectetur adipisicing elit. Eveniet nam corporis vitae omnis soluta, deserunt reprehenderit temporibus debitis error, voluptatem aliquam corrupti necessitatibus aperiam voluptas aut! Deserunt alias nobis dolor? Lorem ipsum dolor sit amet consectetur adipisicing elit. Repudiandae maiores consectetur non nesciunt rerum unde blanditiis ea, eum ducimus libero nulla corporis in a, consequuntur dolorem voluptatum alias. Explicabo, veritatis.",
	// }

	var ProjectDetail = Blog{}

	//intinya disini membangun sebuah variabel "Blog" Baru bernama ProjectDetail untuk menampung 1 data Detailnya.
	//perulangan dibawah digunakan untuk mencari index yang sesuai yang terdapat pada inputan dan nantinya jika index
	//sudah sesuai maka akan langsung di salin datanya ke variabel ProjectDetail.
	// for i, data := range dataProject {
	// 	if id == i {
	// 		ProjectDetail = Blog{
	// 			Title:        data.Title,
	// 			Content:      data.Content,
	// 			StartDate:    data.StartDate,
	// 			EndDate:      data.EndDate,
	// 			Technologies: data.Technologies,
	// 		}
	// 	}
	// }

	// fungsi $1 untuk mendefinisikan menggunakan data id, misal $1,$2",id,title , maka dia menggunakan kondisi id dan title
	err := connection.Conn.QueryRow(context.Background(), "SELECT id_serial,title,description,start_date,end_date,technologies,postdate FROM tb_projects WHERE id_serial=$1", id).Scan(
		&ProjectDetail.ID, &ProjectDetail.Title, &ProjectDetail.Content, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Technologies, &ProjectDetail.PostDate,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Blog": ProjectDetail,
	}

	var tmpl, errtempl = template.ParseFiles("views/myProjectDetail.html")

	if errtempl != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), data)
}

func addmyproject(c echo.Context) error {
	title := c.FormValue("inputTitle")
	content := c.FormValue("inputContent")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	cbox1 := c.FormValue("iot")
	cbox2 := c.FormValue("ui")
	cbox3 := c.FormValue("full")
	cbox4 := c.FormValue("ml")

	var technologies = "IoT : " + cbox1 + " UI : " + cbox2 + " FULL : " + cbox3 + " ML : " + cbox4

	// datestart, _ := time.Parse("02/01/2006 MST", startDate)
	// dateend, _ := time.Parse("02/01/2006 MST", endDate)
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }

	println("Title : " + title)
	println("Content : " + content)
	fmt.Println("Start Date : ", startDate)
	fmt.Println("End Date : ", endDate)
	println("Box IoT : " + cbox1)
	println("Box UI UX : " + cbox2)
	println("Box FullStack : " + cbox3)
	println("Box Machine Learning : " + cbox4)
	println("technologis : " + technologies)

	// var newProject = Blog{
	// 	Title:        title,
	// 	Content:      content,
	// 	StartDate:    datestart,
	// 	EndDate:      dateend,
	// 	Technologies: cbox1 + cbox2 + cbox3 + cbox4,
	// }

	// dataProject = append(dataProject, newProject)

	// println(dataProject)

	_, err := connection.Conn.Exec(context.Background(), "INSERT into tb_projects(title,description,start_date,end_date,technologies,postdate) VALUES ($1, $2, $3, $4, $5, $6)", title, content, startDate, endDate, technologies, time.Now())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteproject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("index: ", id)

	// dataProject = append(dataProject[:id], dataProject[id+1:]...)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id_serial=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, _ := connection.Conn.Query(context.Background(), "SELECT id_serial,title,description,start_date,end_date,technologies,postdate FROM tb_projects WHERE id_serial=$1", id)

	var result []Blog
	for data.Next() {
		var each = Blog{}
		err := data.Scan(&each.ID, &each.Title, &each.Content, &each.StartDate, &each.EndDate, &each.Technologies, &each.PostDate)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		each.StartDateFormator = each.StartDate.Format("2006-01-02")
		each.EndDateFormator = each.EndDate.Format("2006-01-02")

		result = append(result, each)
	}
	var tmpl, err = template.ParseFiles("views/edit-project.html")

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{}{
		"project": result,
	}

	fmt.Println(projects)

	return tmpl.Execute(c.Response(), projects)
}

func postEditProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	title := c.FormValue("inputTitle")
	content := c.FormValue("inputContent")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	cbox1 := c.FormValue("iot")
	cbox2 := c.FormValue("ui")
	cbox3 := c.FormValue("full")
	cbox4 := c.FormValue("ml")

	var technologies = "IoT : " + cbox1 + " UI : " + cbox2 + " FULL : " + cbox3 + " ML : " + cbox4
	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_projects SET title=$1,description=$2,start_date=$3,end_date=$4,technologies=$5,postdate=$6 WHERE id_serial=$7", title, content, startDate, endDate, technologies, time.Now(), id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.Redirect(http.StatusMovedPermanently, "/")
}
