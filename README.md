# euroton-imap-bridge

Go skeleton for an HTTP server that would translate the Proton API into IMAP & SMTP, so the Mail client could work as webmail on top of Dovecot, Cyrus or any other standard mail server. **It doesn't do that yet.** Phase 0 of 7 is done: the module & package layout boot, and no gateway code exists.

[ADR 211](https://github.com/Euroton/euroton-server/blob/main/docs/decisions/211-imap-bridge-unpark.md) in `euroton-server` documents the integration points. [Proton Bridge](https://github.com/ProtonMail/proton-bridge) is the reference for which API routes have IMAP equivalents.

## Planned phases

| Phase | Scope | Status |
|---|---|---|
| 0 | Module, package skeleton, boot | Done |
| 1 | API client for `/auth/info`, `/auth` & `/core/v4/users` | Not started |
| 2 | IMAP `LOGIN`, `LIST`, `SELECT`, `FETCH` against an Inbox mapped from `/mail/v4/conversations` | Not started |
| 3 | IMAP `APPEND`, `MOVE`, `STORE` (flags only), `EXPUNGE` | Not started |
| 4 | SMTP submission into `/mail/v4/messages` & the send pipeline | Not started |
| 5 | `IDLE`, with event polling pushed to the client | Not started |
| 6 | TLS, app passwords, systemd unit | Not started |

Calendar & Contacts are out of scope; they'd need separate CalDAV & CardDAV bridges.

## End-to-end encryption

This design drops it. The client's encryption depends on per-recipient `Packages[]` blobs, and IMAP & SMTP have no field for them. The bridge would unwrap the plaintext package on send, and receive mail with transport TLS at best.

Any working version has to label those sessions as not end-to-end encrypted, because people associate this interface with that guarantee.

## Building

Needs Go 1.22 or newer.

```sh
go build -o bin/euroton-bridge ./cmd/euroton-bridge
./bin/euroton-bridge --help
```

## Licence

AGPL-3.0-or-later. See `LICENSE`.
