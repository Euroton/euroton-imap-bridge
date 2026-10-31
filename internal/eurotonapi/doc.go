// Package eurotonapi is the HTTP client for euroton-server (Phase 1
// of the bridge). The package is intentionally separate from the
// IMAP and SMTP gateways so each can swap mocks for testing.
//
// Implementation lands in a follow-up slice. The wire shapes
// upstream uses (Proton API) are the contract; see the Euroton
// OpenAPI spec at ../openapi-spec/spec/euroton.openapi.yaml for
// the operations the bridge needs:
//
//   - POST /core/v4/auth/info     → SRP info for login
//   - POST /core/v4/auth          → SRP completion
//   - GET  /core/v4/users         → session bootstrap
//   - GET  /mail/v4/conversations → IMAP mailbox listing source
//   - POST /mail/v4/messages      → IMAP APPEND target
//   - GET  /core/v6/events/:eventID → IMAP IDLE diff source
package eurotonapi
