package middleware

import (
	"backend/common"
	"backend/database"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	_ "log"
	"strings"

	_ "github.com/ItsMeSamey/go_utils"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func JWTProtected() fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenString := c.Cookies(common.Cfg.CookieName)
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		token, err := common.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			// log.Printf("Error validating token: %v", utils.WithStack(err))
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		c.Locals("user", token)
		return c.Next()
	}
}

func VerifyAPIKey() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Step 1: Extract the API key from the request headers
		secret := common.Cfg.API_Secret
		apiKeyHeader := c.Get("Authorization")
		if apiKeyHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "API key is missing",
			})
		}

		// Step 2: Decode the API key from base64
		decodedKey, err := base64.URLEncoding.DecodeString(apiKeyHeader)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API key format",
			})
		}

		// Step 3: Split the decoded key into its components
		parts := strings.Split(string(decodedKey), ":")
		if len(parts) != 3 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API key structure",
			})
		}

		serialNumber := parts[0]
		randomBits := parts[1]
		signature := parts[2]

		// Step 4: Recompute the signature using the secret key
		data := fmt.Sprintf("%s:%s", serialNumber, randomBits)
		h := hmac.New(sha256.New, []byte(secret))
		h.Write([]byte(data))
		expectedSignature := hex.EncodeToString(h.Sum(nil))

		// Step 5: Compare the computed signature with the provided signature
		if !hmac.Equal([]byte(signature), []byte(expectedSignature)) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API key signature",
			})
		}

		// Step 6: Check if the API key is revoked (e.g., query a database)
		if isRevoked(serialNumber) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "API key is revoked",
			})
		}

		// Step 7: If the key is valid, proceed to the next handler
		return c.Next()
	}
}

func isRevoked(serialNumber string) bool {
	// Replace this with actual logic to check if the key is revoked
	// For example, query a database to see if the serial number is in a revoked list
	client, exists, _ := database.ClientDB.GetExists(bson.M{"serialNumber": serialNumber})
	log.Println(exists, serialNumber)
	if exists{
		return client.IsRevoked
	}
	return true
}