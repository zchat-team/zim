package zcontext

import (
	"context"
	"github.com/spf13/cast"
	"github.com/zmicro-team/zmicro/core/transport"

	"github.com/zchat-team/zim/pkg/auth"
)

type accountKey struct{}

func ContextWithAccount(ctx context.Context, acc *auth.Account) context.Context {
	return context.WithValue(ctx, accountKey{}, acc)
}

func AccountFromContext(ctx context.Context) (*auth.Account, bool) {
	acc, ok := ctx.Value(accountKey{}).(*auth.Account)
	return acc, ok
}

func GetUid(ctx context.Context) int64 {
	if acc, ok := AccountFromContext(ctx); ok {
		return cast.ToInt64(acc.ID)
	}

	return 0
}

func GetLoginType(ctx context.Context) int {
	if acc, ok := AccountFromContext(ctx); ok {
		t := acc.Metadata["login_type"]
		return cast.ToInt(t)
	}

	return 0
}

func GetUidStr(ctx context.Context) string {
	if acc, ok := AccountFromContext(ctx); ok {
		return acc.ID
	}

	return ""
}

func GetClientIp(ctx context.Context) string {
	if tr, ok := transport.FromTransporter(ctx); !ok {
		return ""
	} else {
		return tr.ClientIp()
	}
}

func GetToken(ctx context.Context) string {
	if tr, ok := transport.FromTransporter(ctx); !ok {
		return ""
	} else {
		return tr.RequestHeader().Get("Token")
	}
}
