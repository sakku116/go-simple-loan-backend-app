package seeder_util

import (
	"backend/config"
	"backend/domain/model"
	"backend/repository"
	bcrypt_util "backend/utils/bcrypt"
	"fmt"

	"github.com/google/uuid"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("main")

func SeedUser(userRepo repository.IUserRepo) error {
	users := []model.User{}

	if config.Envs.INITIAL_ADMIN_USERNAME != "" && config.Envs.INITIAL_ADMIN_PASSWORD != "" {
		hashedPassword, _ := bcrypt_util.Hash(config.Envs.INITIAL_ADMIN_PASSWORD)
		users = append(users, model.User{
			UUID:     uuid.New(),
			Username: config.Envs.INITIAL_ADMIN_USERNAME,
			Password: hashedPassword,
			Email:    fmt.Sprint(config.Envs.INITIAL_ADMIN_USERNAME, "@gmail.com"),
			Role:     "admin",
		})
	} else {
		logger.Warningf("initial admin username and password not set")
	}

	if config.Envs.INITIAL_USER_USERNAME != "" && config.Envs.INITIAL_USER_PASSWORD != "" {
		hashedPassword, _ := bcrypt_util.Hash(config.Envs.INITIAL_USER_PASSWORD)
		users = append(users, model.User{
			UUID:     uuid.New(),
			Username: config.Envs.INITIAL_USER_USERNAME,
			Password: hashedPassword,
			Email:    fmt.Sprint(config.Envs.INITIAL_USER_USERNAME, "@gmail.com"),
			Role:     "user",
		})
	} else {
		logger.Warningf("initial user username and password not set")
	}

	for _, user := range users {
		logger.Infof("seeding user: %s", user.Username)

		// check if user already exists
		existing, _ := userRepo.GetByUsername(user.Username)
		if existing != nil {
			logger.Warningf("user already exists: %s", user.Username)
			continue
		}

		// create user
		err := userRepo.Create(&user)
		if err != nil {
			logger.Warningf("failed to seed user: %s", user.Username)
			continue
		}

		logger.Infof("user seeded: %s", user.Username)
	}

	return nil
}
