package uuid

import (
	"github.com/google/uuid"
	"github.com/gophero/gotools/valuex"
	"strings"
)

func UUID() string {
	return valuex.Must(uuid.NewUUID()).String()
}

func UUID32() string {
	uid := UUID()
	return strings.ReplaceAll(uid, "-", "")
}
