package http

func (s *Server) SetUpRoute() {
	v1 := s.App.Group("/local/v1")

	student := v1.Group("/student")

	student.POST("/sign-in", s.handler.SingIn)
	student.POST("/sign-up", s.handler.SignUp)
	student.GET("/profile/{name}", s.handler.ShowProfile)
	student.PUT("/profile", s.handler.UpdateProfile)
}
