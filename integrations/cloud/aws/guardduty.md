# AWS GuardDuty

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, • [S3 Ingester](amazon_cloudtrail_log_handling) <br /> • [SQS Ingester](s3_sqs_references)
         Kit, [GuardDuty Kit](https://github.com/gravwell/kits/tree/main/)
:::

## GuardDuty Configuration
It is recommend to use Cloudtrail to export GuardDuty Logs.

**Sample KMC Policy**
```{note}
  Replace *ACCOUNT_ID:key/YOUR_KMS_KEY_ID* with your KMS Resource
```

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "GuardDuty",
            "Effect": "Allow",
            "Principal": {
                "Service": "guardduty.AWS_REGION.amazonaws.com"
            },
            "Action": [
                "kms:Encrypt",
                "kms:GenerateDataKey"
            ],
            "Resource": "arn:aws:kms:REGION:ACCOUNT_ID:key/YOUR_KMS_KEY_ID"
        }
    ]
}
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:** 
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/guardduty-well.conf`
```ini
[Storage-Well "guardduty"]
    Location=/opt/gravwell/storage/guardduty
    Tags=aws-guardduty*
```

### [Option 1] Gravwell S3 Ingester Configuration

The simplest workflow is to ingest the GuardDuty logs from CloudTrail's S3 bucket using the [S3 Ingester](amazon_cloudtrail_log_handling). The polling required may introduce a 5-15 minute delay on new logs landing in Gravwell once they hit CloudTrail, but only requires an identity that can query the CloudTrail S3 bucket.

**Sample S3 ingester config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/guardduty.conf`
```
[Bucket "default"]
    Region="us-east-1"
    ID="AKI..."
    Secret="SuperSecretKey..."
    Bucket-ARN = "arn:aws:s3:::aws-guardduty-logs-123456-7890"
    Tag-Name="aws-guardduty"
    Reader=guardduty
    File-Filters=**/*.json.gz
```

#### [Option 2] Gravwell SQS Ingester Configuration

Configure S3 ingester to utilize SQS. A new configuration snippet is required that replaces the [Bucket] declaration.

**Sample SQS config snippet**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/guardduty.conf`
```ini
[SQS-S3-Listener "GuardDuty_Queue-Placeholder"] ## REPLACE PLACEHOLDER WITH YOUR ACTUAL SQS QUEUE NAME
    Tag-Name="aws-guardduty"
    Reader=guardduty
    ID="AKI..."
    Secret="SuperSecretKey..."
    Region="us-east-1"
    Queue-URL="https://sqs.us-east-1.amazonaws.com/..."
    File-Filters=**/*.json.gz
```
This allows the S3 ingester to pull from the SQS queue, and then will pull the S3 object items referenced by it. That log will then be processed and sent to the Gravwell indexer. 
