# MongoDB Server

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [File Follower](/ingesters/file_follow)
:::

This guide collects two logs from MongoDB Server, the self-managed MongoDB database (MongoDB Community Server and MongoDB Enterprise Server):

* The **server log** (`mongod.log`): MongoDB's diagnostic log of connections, authentication, slow operations, replication, storage and startup events. MongoDB Community and Enterprise (4.4 and later) write it in the same JSON format. Gravwell stores it under the `mongodb-diag` tag.
* The **audit log**: a record of who did what (authentication, authorization, user and role changes, schema changes). Only MongoDB Enterprise writes it. Gravwell stores it under the `mongodb-audit` tag.

Both are files on the MongoDB server, collected by File Follower on that server.

The steps in this guide were tested on Ubuntu 24.04 LTS with MongoDB 8.0. Other Linux distributions can use different package commands, user names and file locations. Check each step against your system before you run it.

## MongoDB Configuration

### Server Log

MongoDB writes its server log to the file set in the `systemLog` section of `/etc/mongod.conf`. Confirm that section writes to a file:

```yaml
systemLog:
  destination: file
  logAppend: true
  path: /var/log/mongodb/mongod.log
```

Gravwell reads the file as written; no MongoDB log format settings are needed. See [MongoDB: Log Messages](https://www.mongodb.com/docs/manual/reference/log-messages/).

### Audit Log (MongoDB Enterprise)

Skip this section if you run MongoDB Community or don't need the audit log.

Auditing is off until `auditLog` is set. Add the following to `/etc/mongod.conf`, then restart MongoDB (`sudo systemctl restart mongod`; on a replica set, one member at a time). Set it on every `mongod` and `mongos` you want audited:

```yaml
auditLog:
  destination: file
  format: JSON
  path: /var/log/mongodb/auditLog.json
```

* Use `format: JSON`. File Follower reads text log files, and `BSON` is a binary format.
* Keep the file in `/var/log/mongodb`, so the read access and rotation below cover it.

To check, log in with a user that doesn't exist, then look for the failed attempt:

```bash
mongosh "mongodb://localhost:27017/?authSource=admin" -u nosuchuser -p wrong --eval 'db.runCommand({ping: 1})'
sudo grep authenticate /var/log/mongodb/auditLog.json | tail -n 1
```

You should see a JSON line with `atype` set to `authenticate`.

### File Access for File Follower

File Follower runs as the `gravwell` user in the `adm` group. MongoDB creates its log files readable only by its own user (`0600` in a `0750` directory), so File Follower can't read them until you grant access. MongoDB ignores the system umask for its log files, so adding `gravwell` to the `mongodb` group isn't enough. Grant read access with an access control list (ACL) on the log directory, which also applies to log files MongoDB creates later:

1. On Debian and Ubuntu, install the ACL tools if they're missing (`getfacl: command not found`):

   ```bash
   sudo apt-get install -y acl
   ```

2. Show the current permissions:

   ```bash
   getfacl /var/log/mongodb /var/log/mongodb/mongod.log
   ```

3. Grant File Follower read access, and give the `mongodb` group no access to new files:

   ```bash
   sudo setfacl -m u:gravwell:rx /var/log/mongodb
   sudo setfacl -d -m u:gravwell:r,g::--- /var/log/mongodb
   sudo setfacl -m u:gravwell:r /var/log/mongodb/*.log
   ```

   If the audit log already exists, grant it too: `sudo setfacl -m u:gravwell:r /var/log/mongodb/auditLog.json`. An audit log created later is covered by the second command.

4. Check that File Follower can read the server log. The message confirms it:

   ```bash
   sudo -u gravwell head -c 0 /var/log/mongodb/mongod.log && echo "File Follower can read the log"
   ```

   If you enabled the audit log, repeat the check with `/var/log/mongodb/auditLog.json`. If it prints `Permission denied`, repeat step 3 and check that the paths match `systemLog.path` and `auditLog.path`.

To undo, remove the ACLs with `sudo setfacl -b /var/log/mongodb /var/log/mongodb/*.log` (add `/var/log/mongodb/auditLog.json` if you enabled the audit log). File Follower then can't read the logs.

### Log Rotation

MongoDB rotates its log files only when something tells it to, so most deployments set up rotation. How the logs are rotated affects File Follower.

```{note}
MongoDB's default rotation mode (`systemLog.logRotate: rename`) renames the log to `mongod.log.<timestamp>` and starts a new file. File Follower doesn't follow those timestamped names reliably: if File Follower is stopped or behind when the rename happens, the last lines of the old file are never read; if `File-Filter` also matches the timestamped names, a restarting File Follower reads the renamed file again from the start and stores every line twice. Rotating with the Linux `logrotate` utility and MongoDB's `reopen` mode avoids both, and the `File-Filter` in this guide is written for it. The audit log follows the same setting.
```

To rotate with `logrotate`:

1. Back up the MongoDB configuration. Use a new backup name each time:

   ```bash
   sudo cp /etc/mongod.conf /etc/mongod.conf.bak-$(date +%Y%m%d-%H%M%S)
   ```

2. In `/etc/mongod.conf`, set `systemLog` to the following, then restart MongoDB (`sudo systemctl restart mongod`; on a replica set, one member at a time):

   ```yaml
   systemLog:
     destination: file
     path: /var/log/mongodb/mongod.log
     logAppend: true
     logRotate: reopen
   ```

3. Create `/etc/logrotate.d/mongod`. Schedule and retention are up to you; keep the `create` line as shown:

   ```text
   /var/log/mongodb/mongod.log {
     su mongodb mongodb
     daily
     rotate 4
     missingok
     nodateext
     create 0640 mongodb mongodb
     postrotate
       systemctl kill -s USR1 mongod
     endscript
   }
   ```

   `create 0640` matters: the group permission bits limit the read access granted to File Follower above, so `0600` cuts File Follower off at the first rotation. MongoDB's RHEL packages run as `mongod` instead of `mongodb`; on RHEL, use `mongod mongod` in the `su` and `create` lines.

   Keep `nodateext`: it overrides a `dateext` set in `/etc/logrotate.conf`, which would rename rotated files to `mongod.log-<date>`, and `File-Filter` doesn't match those names. If you add `compress`, also add `delaycompress`, so `mongod.log.1` stays uncompressed until File Follower has finished reading it.

   If you enabled the audit log, list both files in this one entry and add `sharedscripts`. One signal makes MongoDB reopen both files, so both must be renamed before it's sent:

   ```text
   /var/log/mongodb/mongod.log /var/log/mongodb/auditLog.json {
     su mongodb mongodb
     daily
     rotate 4
     missingok
     nodateext
     sharedscripts
     create 0640 mongodb mongodb
     postrotate
       systemctl kill -s USR1 mongod
     endscript
   }
   ```

4. Rotate once to check:

   ```bash
   sudo logrotate -f /etc/logrotate.d/mongod
   ls -l /var/log/mongodb/
   sudo -u gravwell head -c 0 /var/log/mongodb/mongod.log && echo "File Follower can read the new log"
   ```

   You should see `mongod.log` and `mongod.log.1` (and `auditLog.json` and `auditLog.json.1` if you enabled the audit log), followed by the "can read" message.

```{warning}
* Keep File Follower running when rotation runs. If File Follower is stopped during a rotation, it reads the rotated file again from the start when it restarts, and those entries are stored twice.
* With `logRotate: reopen`, let `logrotate` trigger every rotation. A `SIGUSR1` or `db.adminCommand({logRotate: 1})` sent without `logrotate` renaming the file first makes MongoDB empty the current log file, despite `logAppend: true`. For the same reason, list each MongoDB log in only one `logrotate` entry, and keep the server and audit logs in the same entry.
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/mongodb-well.conf`
```ini
[Storage-Well "mongodb-diag"]
    Location=/opt/gravwell/storage/mongodb-diag
    Tags=mongodb-diag

[Storage-Well "mongodb-audit"]
    Location=/opt/gravwell/storage/mongodb-audit
    Tags=mongodb-audit
```

Leave out the `mongodb-audit` well if you don't collect the audit log.

### Gravwell Ingester Configuration: File Follower

Install File Follower on your MongoDB host by following the instructions in [File Follower](/ingesters/file_follow). Then add the following configuration.

**Sample MongoDB config:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf.d/mongodb.conf`
```ini
[Follower "mongodb"]
    Base-Directory="/var/log/mongodb"
    File-Filter="mongod.log,mongod.log.[0-9]"
    Tag-Name=mongodb-diag
    Assume-Local-Timezone=true #Default for assume localtime is false

[Follower "mongodb-audit"]
    Base-Directory="/var/log/mongodb"
    File-Filter="auditLog.json,auditLog.json.[0-9]"
    Tag-Name=mongodb-audit
    Assume-Local-Timezone=true #Default for assume localtime is false
```

Leave out the `mongodb-audit` follower if you don't collect the audit log. The `.[0-9]` names let File Follower finish the most recent rotated file after `logrotate` renames it. Follow only these names: MongoDB's own timestamped rotation names aren't supported (see Log Rotation).

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_file_follow.service`
```

To confirm data is arriving, run `tag=mongodb-diag count` over the last few minutes. MongoDB logs connections and periodic storage checkpoints, so even a quiet server produces entries. For the audit log, repeat the failed login from Audit Log, then run `tag=mongodb-audit count`.

## Additional Resources
* [MongoDB: Log Messages](https://www.mongodb.com/docs/manual/reference/log-messages/)
* [MongoDB: Configure Auditing](https://www.mongodb.com/docs/manual/tutorial/configure-auditing/)
* [MongoDB: Configure Audit Filters](https://www.mongodb.com/docs/manual/tutorial/configure-audit-filters/)
* [MongoDB: mongo Schema Audit Messages](https://www.mongodb.com/docs/manual/reference/audit-message/mongo/)
* [MongoDB: Rotate Log Files](https://www.mongodb.com/docs/manual/tutorial/rotate-log-files/)
* [MongoDB University: MongoDB Server Log Rotation and Retention](https://learn.mongodb.com/learn/course/mongodb-logging-basics/lesson-5-mongodb-server-log-rotation-and-retention/learn?page=2)
* [File Follower](/ingesters/file_follow)
