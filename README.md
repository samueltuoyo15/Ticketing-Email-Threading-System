# Ticketing-Email-Threading-System
SQLite → stores tickets + messages  Redis Streams → background job queue  Worker services → process and send emails  Email SMTP library → send emails  Email threading → manage “reply to this ticket ID”  Webhook listener → receive replies and attach to ticket
