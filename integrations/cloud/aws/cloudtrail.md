# AWS CloudTrail

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
Ingester, • [S3 Ingester](amazon_cloudtrail_log_handling) <br /> • [SQS Ingester](s3_sqs_references)
     Kit, [CloudTrail](https://github.com/gravwell/kits/tree/main/aws_cloudtrail)
:::

## CloudTrail Configuration

There are two primary methods to get CloudTrail data into Gravwell - **S3 Bucket Ingest**, and **SQS-backed ingest**. 
* **S3 Bucket Ingester (Simplest):** This workflow uses the [S3 Ingester](amazon_cloudtrail_log_handling) to pull logs directly from the CloudTrail S3 bucket. It requires an identity with read access to the S3 bucket. The polling interval may introduce a **5-15 minute delay** before logs appear in Gravwell.
* **SQS-S3-Listener Ingester (Near Real-Time):** For faster delivery [SQS can be configured](s3_sqs_references) to notify the ingester as events are generated. 

```{note}
Regardless of the ingestion method CloudTrail inherently has its own internal delay between an action occurring and the event being logged by AWS.
```

### [Option 1] CloudTrail Configuration for: S3 Bucket Ingest

```{note}
We HIGHLY recommend creating a dedicated S3 IAM user for the Gravwell ingester. It's never a good idea to use privileged credentials for dedicated applications like data ingestion.
```

Otherwise the ingester requires no other configuration. For additional configuration options: [S3 Ingester](amazon_cloudtrail_log_handling)

### [Option 2] CloudTrail Configuration for: SQS-S3-Listener Ingest

* Configure CloudTrail with an SNS topic
    * Set up SQS in AWS
    * Configure IAM policies
    * User that can query SQS, and pull from S3
    * **CRITICAL:** [Configure long polling for SQS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/best-practices-setting-up-long-polling.html). By default, SQS queues don't specify a delay between received messages. The ingester will query SQS endlessly resulting in ~12,000 requests per minute, which inflates SQS costs. 
        * If you specify a limited range (1-20 seconds), that significantly reduces the amount of requests that SQS will respond to.
        * To configure long polling, access the SQS queue in AWS Console. Hit "edit" in the top right, and set a "Receive messages wait time". 20 seconds usually results in an imperceptible delay in the time between polling for new logs.
    * Note: If you have KMS encryption enabled (default setting) for SQS, you will need additional permissions for *kms:Decrypt*. 

**_Sample JSON policy for a user that can query CloudTrail utilizing SQS_**
```{note}
  Replace `ACCOUNT_ID:key/YOUR_KMS_KEY_ID` with your KMS Resource
```

```{code-block} json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "SQSConsume",
            "Effect": "Allow",
            "Action": [
                "sqs:ReceiveMessage",
                "sqs:DeleteMessage",
                "sqs:GetQueueAttributes",
                "sqs:ChangeMessageVisibility"
            ],
            "Resource": "arn:aws:sqs:REGION:ACCOUNT_ID:YOUR_QUEUE_NAME"
        },
        {
            "Sid": "S3Read",
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:ListBucket"
            ],
            "Resource": [
                "arn:aws:s3:::YOUR_CLOUDTRAIL_S3_BUCKET",
                "arn:aws:s3:::YOUR_CLOUDTRAIL_S3_BUCKET/*"
            ]
        },
        {
            "Sid": "KMSDecryptForSQS",
            "Effect": "Allow",
            "Action": [
                "kms:Decrypt"
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
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/cloudtrail-well.conf`
```ini
[Storage-Well "cloudtrail"]
    Location=/opt/gravwell/storage/cloudtrail
    Tags=aws-cloudtrail*
```

### Gravwell Ingester Configuration

#### [Option 1] Gravwell S3 Ingester Configuration

The simplest workflow is to ingest logs from CloudTrail's S3 bucket using the [S3 Ingester](https://docs.gravwell.io/ingesters/s3.html#amazon-cloudtrail-log-handling). The polling required may introduce a 5-15 minute delay on new logs landing in Gravwell once they hit CloudTrail, but only requires an identity that can query the CloudTrail S3 bucket.

**Sample S3 ingester config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/cloudtrail.conf`
```
[Bucket "default"]
    Region="us-east-1"
    ID="AKI..."
    Secret="SuperSecretKey..."
    Bucket-ARN = "arn:aws:s3:::aws-cloudtrail-logs-123456-7890"
    Tag-Name="aws-cloudtrail"
    Reader=cloudtrail
    File-Filters=**/*.json.gz
```

##### [Option 2] Gravwell SQS Ingester Configuration

Configure S3 ingester to utilize SQS. A new configuration snippet is required that replaces the [Bucket] declaration.

**Sample SQS config snippet**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/cloudtrail.conf`
```ini
[SQS-S3-Listener "CloudTrail_Queue-Placeholder"] ## REPLACE PLACEHOLDER WITH YOUR ACTUAL SQS QUEUE NAME
    Tag-Name="aws-cloudtrail"
    Reader=cloudtrail
    ID="AKI..."
    Secret="SuperSecretKey..."
    Region="us-east-1"
    Queue-URL="https://sqs.us-east-1.amazonaws.com/..."
    File-Filters=**/*.json.gz
```
This allows the S3 ingester to pull from the SQS queue, and then will pull the S3 object items referenced by it. That log will then be processed and sent to the Gravwell indexer.


```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```