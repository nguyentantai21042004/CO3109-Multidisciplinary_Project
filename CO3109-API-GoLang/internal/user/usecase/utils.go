package usecase

import (
	"context"
	"net/http"

	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/internal/role"
	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/recognize"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

func (uc implUsecase) getAndCheckUserPermission(ctx context.Context, sc scope.Scope, tgtUserID string) ([]models.User, error) {
	// Check if user is updating their own avatar
	isSelf := sc.UserID == tgtUserID

	// If self-update, we only need to fetch the current user
	if isSelf {
		users, err := uc.repo.List(ctx, sc, user.ListOptions{
			IDs: []string{sc.UserID},
		})
		if err != nil {
			uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.repo.List: %v", err)
			return nil, err
		}

		if len(users) != 1 {
			uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.repo.List: %v", "user not found")
			return nil, user.ErrUserNotFound
		}

		return users, nil
	}

	// For admin updates, fetch both users and check permissions
	us, err := uc.repo.List(ctx, sc, user.ListOptions{
		IDs: []string{sc.UserID, tgtUserID},
	})
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.repo.List: %v", err)
		return nil, err
	}

	if len(us) != 2 {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.repo.List: %v", "user not found")
		return nil, user.ErrUserNotFound
	}

	// Find the session user
	var u models.User
	for _, usr := range us {
		if usr.ID == sc.UserID {
			u = usr
			break
		}
	}

	// Get roles for permission check
	roles, err := uc.roleUC.List(ctx, sc, role.ListInput{})
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.roleUC.List: %v", err)
		return nil, err
	}

	// Find the session user's role
	var r models.Role
	found := false
	for _, rl := range roles.Roles {
		if rl.ID == u.RoleID {
			r = rl
			found = true
			break
		}
	}

	if !found {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar: %v", "session user role not found")
		return nil, role.ErrRoleNotFound
	}

	// Check if session user has admin privileges
	if r.Code != "ADMIN" && r.Code != "SUPER_ADMIN" {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar: %v", "permission denied - only admins can update other users")
		return nil, user.ErrPermissionDenied
	}

	return us, nil
}

func (uc implUsecase) syncAvatarUser(ctx context.Context, url string, userID string) error {
	file, err := recognize.DownloadFile(url)
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.curl.DownloadFile: %v", err)
		return err
	}

	req, err := recognize.CreateSaveImageRequest(file, userID)
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.recognize.CreateSaveImageRequest: %v", err)
		return err
	}

	resp, err := recognize.SendRequest(req)
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.recognize.SendRequest: %v", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.recognize.SendRequest: %v", "status code not ok")
		return err
	}

	return nil
}
