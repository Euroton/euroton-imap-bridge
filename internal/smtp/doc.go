// Package smtp is the SMTP submission gateway (Phase 4 of the bridge).
//
// The wire protocol is RFC 5321 + the SMTP-AUTH (RFC 4954) extension
// for client login. Phase 4 commands:
//
//   - EHLO with STARTTLS + AUTH PLAIN advertised
//   - AUTH PLAIN <base64 user\0user\0app-password> → maps to the same
//     app-password flow as the IMAP gateway
//   - MAIL FROM, RCPT TO, DATA, QUIT
//
// On DATA: the bridge unwraps the inbound RFC 5322 message, computes
// a Packages[] envelope on behalf of the user (plaintext-relay only —
// E2EE is dropped per README), and POSTs to euroton-server's
// /mail/v4/messages send pipeline.
//
// VERP envelopes + bounce ingestion stay on the euroton-server side;
// the bridge submits as a normal SMTP client to the configured
// upstream relay (Postfix on loopback, by default).
package smtp
