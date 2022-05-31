func buildSql(email string) string {
	return fmt.Sprintf("SELECT * FROM users WHERE email='%s';", email)
  }
  
  buildSql("oyetoketoby80@gmail.com")
  
  // SELECT * FROM users WHERE email='oyetoketoby80@gmail.com';