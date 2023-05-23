package utils

var config = NewConfig("config.toml")

// func GenerateToken(user *models.User) (string, error) {
// 	claims := jwt.MapClaims{
// 		"id":  user.ID,
// 		"exp": time.Now().Add(24 * time.Hour).Unix(),
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	tokenString, err := token.SignedString(config.GetJWTSecret())
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }

// func ParseToken(tokenString string) (jwt.MapClaims, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return config.GetJWTSecret(), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		return nil, errors.New("invalid token")
// 	}
// 	return claims, nil
// }

// func HashPassword(password string) string {
// 	hash := sha256.New()
// 	hash.Write([]byte(password))
// 	hashedPassword := hash.Sum(nil)
// 	return hex.EncodeToString(hashedPassword)
// }

// func VerifyToken(tokenString string) (bool, error) {
// 	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return config.GetJWTSecret(), nil
// 	})
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
// 			c.Abort()
// 			return
// 		}

// 		// 调用VerifyToken函数验证token
// 		valid, err := VerifyToken(tokenString)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 			c.Abort()
// 			return
// 		}

// 		claims, err := ParseToken(tokenString)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("user", claims)

// 		if !valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }
