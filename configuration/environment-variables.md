# ���ϐ�

�C���f�N�T�A�E�F�u�T�[�o�A�C���W�F�X�^�̊e�R���|�[�l���g�ł́A�ꕔ�̃p�����[�^��ݒ�t�@�C���ł͂Ȃ����ϐ��Őݒ�ł���悤�ɂȂ�܂����B����́A��K�͂ȃf�v���C�����g�̂��߂ɁA����ʓI�Ȑݒ�t�@�C�����쐬����̂ɖ𗧂��܂��B�����̃f�B���N�e�B�u���܂ސݒ�ϐ��́A�J���}��؂�̃��X�g���g���Ċ��ϐ��Őݒ肵�܂��B�Ⴆ�΁AFederator�̋N�����ɃC���W�F�X�g�V�[�N���b�g���w�肷��ɂ́A���̂悤�ɂ��܂��B

```
GRAVWELL_INGEST_SECRET=MyIngestSecret /opt/gravwell/bin/gravwell_federator
```

���ϐ����̍Ō��"_FILE "��t����ƁAGravwell�͂��̊��ϐ��Ƀt�@�C���ւ̃p�X���܂܂�Ă���A���̃t�@�C���ɖړI�̃f�[�^���܂܂�Ă���Ƃ݂Ȃ��܂��B����́A[Docker��"secrets"�@�\](https://docs.docker.com/engine/swarm/secrets/)�Ƒg�ݍ��킹��Ɠ��ɕ֗��ł��B

```
GRAVWELL_INGEST_AUTH_FILE=/run/secrets/ingest_secret /opt/gravwell/bin/gravwell_indexer
```

���F���ϐ��̒l�́A�Ή�����t�B�[���h���K�؂Ȑݒ�t�@�C���igravwell.conf�܂��̓C���W�F�X�^�[�̐ݒ�t�@�C���j��**�����I�ɐݒ肳��Ă��Ȃ��ꍇ�ɂ̂�**�g�p����܂� �B

## �C���f�N�T�[�ƃE�F�u�T�[�o�[

���̕\�́A`gravwell.conf`�̂ǂ̃p�����[�^���A�C���f�N�T�ƃE�F�u�T�[�o�̊��ϐ��Ƃ��Đݒ�ł��邩�������Ă��܂��B�����̕ϐ��́A�p�����[�^�� `gravwell.conf` ��**�ݒ肳��Ă��Ȃ��ꍇ�ɂ̂�**�g�p�ł��邱�Ƃɒ��ӂ��Ă��������B

| gravwell.conf ���̕ϐ� | ���ϐ� | �� |
|:------|:----|:---|----:|
| Ingest-Auth | GRAVWELL_INGEST_AUTH | GRAVWELL_INGEST_AUTH=CE58DD3F22422C2E348FCE56FABA131A |
| Control-Auth | GRAVWELL_CONTROL_AUTH | GRAVWELL_CONTROL_AUTH=C2018569D613932A6BBD62A03A101E84 |
| Indexer-UUID | GRAVWELL_INDEXER_UUID | GRAVWELL_INDEXER_UUID=a6bb4386-3433-11e8-bc0b-b7a5a01a3120 |
| Webserver-UUID | GRAVWELL_WEBSERVER_UUID | GRAVWELL_WEBSERVER_UUID=b3191f54-3433-11e8-a0c2-afbff4695836 |
| Remote-Indexers | GRAVWELL_REMOTE_INDEXERS | GRAVWELL_REMOTE_INDEXERS=172.20.0.1:9404,172.20.0.2:9404|
| Replication-Peers | GRAVWELL_REPLICATION_PEERS | GRAVWELL_REPLICATION_PEERS=172.20.0.1:9406,172.20.0.2:9406 |
| Datastore | GRAVWELL_DATASTORE | GRAVWELL_DATASTORE=172.20.0.10:9405 |

## �C���W�F�X�^�[

�C���W�F�X�^�[�ł����l�ɁA�ꕔ�̃p�����[�^��ݒ�t�@�C���Ŗ����I�ɐݒ肷��̂ł͂Ȃ��A���ϐ��Ƃ��Đݒ肷�邱�Ƃ��ł��܂��B

| Config�t�@�C�����̕ϐ� | ���ϐ� | �� |
|:------|:----|:---|
| Ingest-Secret | GRAVWELL_INGEST_SECRET | GRAVWELL_INGEST_SECRET=CE58DD3F22422C2E348FCE56FABA131A |
| Log-Level | GRAVWELL_LOG_LEVEL | GRAVWELL_LOG_LEVEL=DEBUG |
| Cleartext-Backend-target | GRAVWELL_CLEARTEXT_TARGETS | GRAVWELL_CLEARTEXT_TARGETS=172.20.0.1:4023,172.20.0.2:4023 |
| Encrypted-Backend-target | GRAVWELL_ENCRYPTED_TARGETS | GRAVWELL_ENCRYPTED_TARGETS=172.20.0.1:4024,172.20.0.2:4024 |
| Pipe-Backend-target | GRAVWELL_PIPE_TARGETS | GRAVWELL_PIPE_TARGETS=/opt/gravwell/comms/pipe |


### �t�F�f���[�^�[�ŗL�̕ϐ�

�t�F�f���[�^�ɂ���đ����̃��X�i�[�����s�ł��܂����A���ꂼ��̃��X�i�[�ɂ͂��ꂼ��قȂ�C���W�F�X�g�E�V�[�N���b�g���֘A�t�����Ă��邽�߁A���s���ɂ����̃��X�i�[�̃C���W�F�X�g�E�V�[�N���b�g���ʂɐݒ肷�邽�߂̓��ʂȊ��ϐ��̃Z�b�g���g���܂��B

�e���X�i�[�ɂ͖��O������܂��B�ȉ��̗�ł́A���X�i�[�̖��O��"base"�ł�:

```
[IngestListener "base"]
	Cleartext-Bind = 0.0.0.0:4023
	Tags=syslog
```

���s���ɂ���"base"�Ƃ������O�̃��X�i�[�̃C���W�F�X�g�V�[�N���b�g���w�肷��̂ɂ́A�ϐ�`FEDERATOR_base_INGEST_SECRET`���g���܂�:

```
FEDERATOR_base_INGEST_SECRET=SuperSecret /opt/gravwell/bin/gravwell_federator
```

���邢�́A���̊��ϐ��̏ꍇ�Ɠ��l�A�ݒ�p�t�@�C�����w�肷�邱�Ƃ��ł��܂�:

```
FEDERATOR_base_INGEST_SECRET_FILE=/run/secrets/federator_base_secret /opt/gravwell/bin/gravwell_federator
```

### �f�[�^�X�g�A�ŗL�̕ϐ�

[�f�[�^�X�g�A](#!distributed/frontend.md) ���A���s���Ɋ��ϐ��ɂ���Đݒ�ł��܂��B

| gravwell.conf ���̕ϐ� | ���ϐ� | �� |
|------------------------|----------------------|---------|
| Datastore-Listen-Address | GRAVWELL_DATASTORE_LISTEN_ADDRESS | GRAVWELL_DATASTORE_LISTEN_ADDRESS=192.168.1.100 |
| Datastore-Port | GRAVWELL_DATASTORE_LISTEN_PORT | GRAVWELL_DATASTORE_LISTEN_PORT=9995 |
