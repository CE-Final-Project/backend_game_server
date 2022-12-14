package commands

import "github.com/ce-final-project/backend_game_server/gateway/internal/dto"

type AuthCommands struct {
	RegisterAccount RegisterAccountCmdHandler
	LoginAccount    LoginAccountCmdHandler
	FriendInvite    FriendInviteCmdHandler
}

func NewAuthCommands(registerAccount RegisterAccountCmdHandler, loginAccount LoginAccountCmdHandler, friendInvite FriendInviteCmdHandler) *AuthCommands {
	return &AuthCommands{
		RegisterAccount: registerAccount,
		LoginAccount:    loginAccount,
		FriendInvite:    friendInvite,
	}
}

type RegisterAccountCommand struct {
	RegisterDto *dto.RegisterAccount
}

func NewRegisterAccountCommand(registerDto *dto.RegisterAccount) *RegisterAccountCommand {
	return &RegisterAccountCommand{RegisterDto: registerDto}
}

type LoginAccountCommand struct {
	LoginDto *dto.LoginAccount
}

func NewLoginAccountCommand(loginDto *dto.LoginAccount) *LoginAccountCommand {
	return &LoginAccountCommand{LoginDto: loginDto}
}

type FriendInviteCommand struct {
	FriendInviteDto *dto.FriendInvite
}

func NewFriendInviteCommand(friendInviteDto *dto.FriendInvite) *FriendInviteCommand {
	return &FriendInviteCommand{FriendInviteDto: friendInviteDto}
}
