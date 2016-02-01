package debug

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jsimonetti/ldapserv/ldap"
)

func (d *DebugBackend) Abandon(w ldap.ResponseWriter, m *ldap.Message) {
	r := m.GetAbandonRequest()
	spew.Dump(r)
	// retreive the request to abandon, and send a abort signal to it
	if requestToAbandon, ok := m.Client.GetMessageByID(int(r)); ok {
		requestToAbandon.Abandon()
	}
}
