package usecase

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"dinz-rentbike/internal/domain/constants"
	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
	"dinz-rentbike/pkg/jwt"
	"dinz-rentbike/pkg/utils"
)

type authUsecase struct {
	userRepo       contract.UserRepository
	authManager    *jwt.AuthManager
	mailjetService contract.MailjetService
}

func NewAuthUsecase(userRepo contract.UserRepository, authManager *jwt.AuthManager, mailjetService contract.MailjetService) contract.AuthUsecase {
	return &authUsecase{userRepo: userRepo, authManager: authManager, mailjetService: mailjetService}
}

func (u *authUsecase) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashed,
		Role:     constants.UserRoleCustomer,
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	if u.mailjetService != nil {
		go u.mailjetService.Send(&contract.EmailRequest{
			ToEmail:  req.Email,
			ToName:   req.Name,
			Subject:  "Welcome to Dinz RentBike!",
			HTMLBody: welcomeEmailHTML(req.Name),
		})
	}

	return &dto.RegisterResponse{
		User: dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
			Role:  user.Role,
		},
	}, nil
}

func (u *authUsecase) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !utils.ValidatePassword(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := u.authManager.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}

func welcomeEmailHTML(name string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
</head>
<body style="margin:0;padding:0;background:#f4f4f4;font-family:Arial,sans-serif">
    <table width="100%%" cellpadding="0" cellspacing="0" style="background:#f4f4f4;padding:40px 0">
        <tr>
            <td align="center">
                <table width="600" cellpadding="0" cellspacing="0" style="background:#fff;border-radius:12px;overflow:hidden;box-shadow:0 2px 12px rgba(0,0,0,0.08)">
                    <!-- Header -->
                    <tr>
                        <td style="background:linear-gradient(135deg,#1a1a2e,#16213e);padding:40px 30px;text-align:center">
                            <h1 style="color:#fff;margin:0;font-size:28px">🏍️ Dinz RentBike</h1>
                            <p style="color:#a0aec0;margin:8px 0 0;font-size:14px">Rental Motor &amp; Mobil Terpercaya</p>
                        </td>
                    </tr>
                    <!-- Body -->
                    <tr>
                        <td style="padding:40px 30px">
                            <h2 style="color:#1a1a2e;margin:0 0 12px;font-size:22px">Selamat Datang, %s! 👋</h2>
                            <p style="color:#4a5568;margin:0 0 24px;font-size:16px;line-height:1.6">
                                Terima kasih telah bergabung dengan <strong>Dinz RentBike</strong>. Akun kamu sudah aktif dan siap digunakan untuk menyewa motor atau mobil pilihanmu.
                            </p>
                            <!-- Features -->
                            <table width="100%%" cellpadding="0" cellspacing="0" style="margin-bottom:24px">
                                <tr>
                                    <td style="padding:16px;background:#f0f9ff;border-radius:8px">
                                        <p style="margin:0;font-size:15px">🏍️ <strong>Motor</strong> — Matic, Sport, Trail</p>
                                    </td>
                                </tr>
                                <tr><td style="height:8px"></td></tr>
                                <tr>
                                    <td style="padding:16px;background:#f0fff4;border-radius:8px">
                                        <p style="margin:0;font-size:15px">🚗 <strong>Mobil</strong> — MPV, SUV, Hatchback</p>
                                    </td>
                                </tr>
                                <tr><td style="height:8px"></td></tr>
                                <tr>
                                    <td style="padding:16px;background:#fff5f5;border-radius:8px">
                                        <p style="margin:0;font-size:15px">⚡ <strong>Pembayaran Mudah</strong> — QRIS, GoPay, OVO, DANA</p>
                                    </td>
                                </tr>
                            </table>
                            <p style="color:#4a5568;margin:0 0 24px;font-size:14px;line-height:1.6">
                                Mulai jelajahi katalog kendaraan kami dan temukan kendaraan yang cocok untuk perjalananmu. Proses booking cepat, harga transparan, tanpa biaya tersembunyi.
                            </p>
                            <table width="100%%" cellpadding="0" cellspacing="0">
                                <tr>
                                    <td align="center">
                                        <a href="https://dinz-rentbike.up.railway.app/vehicles" target="_blank" style="display:inline-block;background:#e94560;color:#fff;text-decoration:none;padding:14px 40px;border-radius:8px;font-size:16px;font-weight:bold">Jelajahi Katalog</a>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                    <!-- Footer -->
                    <tr>
                        <td style="background:#1a1a2e;padding:24px 30px;text-align:center">
                            <p style="color:#a0aec0;margin:0;font-size:12px">
                                © 2026 Dinz RentBike. All rights reserved.
                            </p>
                            <p style="color:#718096;margin:4px 0 0;font-size:11px">
                                Butuh bantuan? Hubungi kami di <a href="mailto:support@rentbike.com" style="color:#e94560;text-decoration:none">support@rentbike.com</a>
                            </p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>`, name)
}
