package main

// import "github.com/gin-gonic/gin"

// func SetUpRouter() *gin.Engine{
//     router := gin.Default()
//     return router
// }

// func TestHomepageHandler(t *testing.T) {
//     mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
//     r := SetUpRouter()
//     r.GET("/", HomepageHandler)
//     req, _ := http.NewRequest("GET", "/", nil)
//     w := httptest.NewRecorder()
//     r.ServeHTTP(w, req)

//     responseData, _ := ioutil.ReadAll(w.Body)
//     assert.Equal(t, mockResponse, string(responseData))
//     assert.Equal(t, http.StatusOK, w.Code)
// }
