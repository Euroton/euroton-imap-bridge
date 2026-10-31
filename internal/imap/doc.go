// Package imap is the IMAP gateway (Phases 2-3 of the bridge).
//
// IMAP4rev1 (RFC 3501) is the wire protocol. The minimal command set
// Phase 2 implements:
//
//   - CAPABILITY: returns IMAP4rev1 + IDLE.
//   - LOGIN <user> <app-password>: authenticates against euroton-server's
//     SRP flow under the hood. The bridge maps "app passwords" (long
//     random strings the user generates in account settings) onto
//     standard plain LOGIN since IMAP clients don't speak SRP.
//   - LIST "" "*": enumerate folders from euroton-server's labels API.
//   - SELECT <mailbox>: pin a mailbox for subsequent FETCH/STORE.
//   - FETCH 1:* (UID FLAGS RFC822): pull messages.
//   - LOGOUT: cleanly close.
//
// Phase 3 layers in APPEND, MOVE, STORE (flags), EXPUNGE.
// Phase 5 layers in IDLE driven by euroton-server's events poll.
//
// The package is intentionally separate from the SMTP gateway so each
// can swap mocks for testing.
package imap
