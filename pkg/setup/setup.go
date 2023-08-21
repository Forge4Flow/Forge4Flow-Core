package setup

import (
	"context"
	feature "github.com/forge4flow/forge4flow-core/pkg/authz/feature"
	user "github.com/forge4flow/forge4flow-core/pkg/authz/user"
	warrant "github.com/forge4flow/forge4flow-core/pkg/authz/warrant"
	"github.com/forge4flow/forge4flow-core/pkg/config"
	"log"
)

func InitialSetup(cfg *config.Forge4FlowConfig, featureSvc *feature.FeatureService, userSvc *user.UserService, warrantSvc *warrant.WarrantService) {
	ctx := context.Background()
	featureId := "forge4flow-admin"
	featureName := "Forge4Flow Admin"
	featureDesc := "Enables access to the Forge4Flow Admin Dashboard"

	// Check if the feature exists and create if needed
	_, err := featureSvc.GetByFeatureId(ctx, featureId)
	if err != nil {
		newFeatureSpec := feature.FeatureSpec{
			FeatureId:   featureId,
			Name:        &featureName,
			Description: &featureDesc,
		}

		_, err := featureSvc.Create(ctx, newFeatureSpec)
		if err != nil {
			log.Fatalln("Unable to create admin feature")
		}
	}

	// Verify the admin user exists and create if needed
	_, err = userSvc.GetByUserId(ctx, cfg.AdminAccount)
	if err != nil {
		newUserSpec := user.UserSpec{
			UserId: cfg.AdminAccount,
		}

		_, err = userSvc.Create(ctx, newUserSpec)
		if err != nil {
			log.Fatalln("Unable to create admin role")
		}

		// Add feature to the user
		newWarrantSpecUser := warrant.WarrantSpec{
			ObjectType: "feature",
			ObjectId:   featureId,
			Relation:   "member",
			Subject: &warrant.SubjectSpec{
				ObjectType: "user",
				ObjectId:   cfg.AdminAccount,
			},
		}

		_, err = warrantSvc.Create(ctx, newWarrantSpecUser)
		if err != nil {
			log.Fatalln("Unable to assign admin feature to user")
		}
	}
}
