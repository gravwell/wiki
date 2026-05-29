# AWS GuardDuty

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [S3 Ingester](amazon_cloudtrail_log_handling)
         Kit, [GuardDuty Kit](https://github.com/gravwell/kits/tree/main/)
:::

## GuardDuty Configuration
It is recommend to export GuardDuty findings to an S3 bucket for ingestion.

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

### Gravwell S3 Ingester Configuration

The simplest workflow is to ingest the GuardDuty logs from an S3 bucket using the [S3 Ingester](amazon_cloudtrail_log_handling). The polling required may introduce a 5-15 minute delay on new logs landing in Gravwell once they hit S3, but only requires an identity that can query the GuardDuty S3 bucket.

**Sample S3 ingester config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/guardduty.conf`
```
[Bucket "default"]
    Region="us-east-1"
    ID="AKI..."
    Secret="SuperSecretKey..."
    Bucket-ARN = "arn:aws:s3:::aws-guardduty-logs-123456-7890"
    Tag-Name="aws-guardduty"
    File-Filters=**/*.json.gz
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```