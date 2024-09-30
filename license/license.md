# Gravwell License System

Gravwell provides several license tiers that enable additional features and capabilities.  The most basic license tier is the Community Edition and the most advanced tier is the Unlimited license.  The Community Edition license allows limited ingest on a single indexer; it is the only license tier that does *not* have unlimited ingest.

Gravwell installations are licensed using a license file that is located at the path /opt/gravwell/etc/license by default.  The license location can be controlled in the gravwell.conf file by setting the License-Location variable in the [Global] configuration block.  Gravwell will validate the license at each startup and periodically check it.

## Installing a License

A license can be installed via the GUI, the CLI, or by simply copying the license file to the appropriate location.  Only administrative users can update the license via the GUI or CLI.  At startup, if no license is installed or the current license is invalid the system will wait for a valid license to be installed.  The GUI will prompt for a license upload and wait for a valid license before completing startup.  The license installation API validates the provided license and will not allow you to install an invalid license.

![Upload a License](upload.png)

![License Install](install.png)

If you use the CLI or GUI to install a license Gravwell will distribute the license to all connected indexers automatically.  Gravwell will also validate the installed license for each indexer and update if necessary at each connection.  That means that if you bring a new indexer online it will automatically receive the appropriate license when the webserver connects.

## Updating a License

Updating a Gravwell license can be performed using the CLI or GUI without restarting Gravwell, but if you just copy the license file to the appropriate directory you must restart Gravwell in order for the license to be reloaded.  Only administrators can update the license file through the GUI or CLI.  To update the license, log into Gravwell as an administrator and navigate to the license section under the administrator panel.  The License page will provide details about your license and its expiration date, select the new license file and upload it to update the webserver and all connected Gravwell indexers.

![License Page](navpanel.png)

![License Info](licinfo.png)

![License Upload](licupload.png)

## License Expiration

All Gravwell licenses except the Free edition have an expiration date and once a license has expired Gravwell will not start.  A license expires in four discrete steps:

1. Warning about upcoming expiration
2. Expiration grace period
3. Ingest is disabled
4. Search is disabled

Prior to expiring Gravwell will post a notification in the GUI warning that the license is about to expire, once the expiration date is reached there is a 14 day grace period where Gravwell is fully functional.  The grace period allows you to continue using Gravwell, but if you restart the Gravwell indexer or webserver it will stop and wait for a valid license.  Once the Grace period has expired Gravwell will disable ingest, in this state you can still query your data but Gravwell will not ingest any new data.  After the query grace period expires Gravwell will disable search functionality.

Gravwell will never delete data due to license expiration, all stored data, resources, dashboards, and query history is preserved.  If a new license is installed everything will be right where you left it.

Here is a handy table that explains the events leading up to and after license expiration.

| Event | Description | Time to License Expiration |
|-------|-------------|:--------------------------:|
| Warning 1 | A notification indicating that the license will expire in less than 30 days | T - 30 days |
| Warning 2 | A notification indicating that the license will expire in less than 15 days | T - 15 days |
| Expiration | A notification indicating that the license is expired, 14 day grace period begins | T - 0 |
| Ingest Disabled | Ingest is disabled and a notification indicating that the license is expired | T + 15 days |
| Query Disabled | Searching is disabled and a notification indicating that the license is expired | T + 30 days |

## Gravwell License Types

| Type                | Identifier  | Basic Features | Unlimited Ingest | Cluster | Replication | CBAC | HA Webservers | SSO | Notes                                               |
|---------------------|-------------|:--------------:|:----------------:|:-------:|:-----------:|:----:|:-------------:|:---:|:----------------------------------------------------|
| Free                | free        | ✅             |                  |         |             |      |               |     | Free with 2GB/day ingest, no sign-up required, non-commercial use only, never expires. |
| Community Edition   | community   | ✅             |                  |         |             |      |               |     | Free signup with 13.9 GB/day ingest, authorized for commercial use, [free licenses with instant delivery](https://www.gravwell.io/community-edition). |
| CE Advanced         | community   | ✅             |                  |         |             |      |               |     | Free signup with 50 GB/day ingest, authorized for commercial use, [free license](https://www.gravwell.io/community-edition-advanced) after validation.  Business email required. |
| Pro                 | single      | ✅             |   ✅             |         |             |      |               |     | Single indexer, unlimited ingest, limited features. |
| Enterprise          | single      | ✅             |   ✅             |         |  ✅         | ✅   |               | ✅  | Single indexer, full feature set, offline replication supported. |
| Cluster             | cluster     | ✅             |   ✅             |  ✅     |  ✅         | ✅   |    ✅         | ✅  | Cluster deployment with online replication, distributed webservers, and full feature set. | 
| Unlimited           | unlimited   | ✅             |   ✅             |  ✅     |  ✅         | ✅   |    ✅         | ✅  | Cluster deployment no limit on indexer count; the *go nuts* license tier. |
| Cloud               | cloud       | ✅             |                  | <img src="/_static/favicon.ico" alt="gravwell managed" width="20"/> | <img src="/_static/favicon.ico" alt="gravwell managed" width="20"/>  | ✅   | <img src="/_static/favicon.ico" alt="gravwell managed" width="20"/>  | ✅  | ✅  | Gravwell managed cloud deployment, opaque infrastructure with contract defined ingest. |


🗸 - rate limited

✅ - full support

<img src="/_static/favicon.ico" alt="gravwell managed" width="20"/> - Gravwell managed

## Free Tier Feature Availability Warnings

As illustrated above, not all features are available when using Free Tier. Gravwell will display warning messages if it detects state that is incompatible with Free Tier

![CBAC is not compatible with Free Tier](free-tier-error.png)

Please consult the following table for advice on resolving Free Tier errors, should you encounter them.

| Message                          | Fix                                                                                                 |
| -------------------------------- | --------------------------------------------------------------------------------------------------- |
| `Overwatch enabled`              | Ensure you're not running an [Overwatch webserver](#gravwell-overwatch)                             |
| `<N> Remote-Indexers configured` | Ensure your `gravwell.conf` has no more than one [remote indexer configured](remote-indexers-conig) |
| `CBAC enabled`                   | Ensure `Enable-CBAC` is not set in your `gravwell.conf` ([CBAC docs](#enabling-cbac))               |
| `Remote Datastore enabled`       | Ensure `Datastore` is not set in your `gravwell.conf` ([Datastore Server docs](#datastore_server))  |
| `Replication configured`         | Ensure your `gravwell.conf` does not have a [replication section](#data-replication)                |
