package provider

import (
	"database/sql"
	"log"
	"time"

	"server/pkg/config"
	"server/pkg/database/transactor"
	"server/pkg/middleware"
	"server/pkg/utils/cloudinaryutils"
	"server/pkg/utils/encryptutils"
	"server/pkg/utils/jwtutils"
	"server/pkg/utils/redisutils"
	"server/pkg/utils/smtputils"

	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
)

var (
	cloudinaryUtil    cloudinaryutils.CloudinaryUtil
	jwtUtil           jwtutils.JwtUtil
	smtpUtil          smtputils.SMTPUtils
	redisUtil         redisutils.RedisUtil
	passwordEncryptor encryptutils.PasswordEncryptor
	base64Encryptor   encryptutils.Base64Encryptor
	store             transactor.Transactor
	authMiddleware    *middleware.AuthMiddleware
	cronJob           *cron.Cron
)

func ProvideUtils(cfg *config.Config, db *sql.DB, rdb *redis.Client) {
	cloudinaryUtil = cloudinaryutils.NewCloudinaryUtil()
	jwtUtil = jwtutils.NewJwtUtil(cfg.Jwt)
	smtpUtil = smtputils.NewSMTPUtils(cfg.SMTP)
	passwordEncryptor = encryptutils.NewBcryptPasswordEncryptor(cfg.App.BCryptCost)
	base64Encryptor = encryptutils.NewBase64Encryptor()
	redisUtil = redisutils.NewRedisUtils(cfg.Redis, rdb)
	store = transactor.NewTransactor(db)

	authMiddleware = middleware.NewAuthMiddleware(jwtUtil)

	wib, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatalf("Failed to load WIB timezone: %v", err)
	}
	cronJob = cron.New(cron.WithLocation(wib))
}
