# Pi-hole

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Kit, [Pi-hole Kit](https://github.com/gravwell/kits/tree/main/pihole)
:::

## Pi-hole Configuration

An API key is required. Use the following steps to get the key:

*Pi-hole v6.0+:*
1. Log into your Admin Dashboard
2. Go to `Settings > API/Web Interface` tab
3. Switch the view from Basic to Expert using the toggle at the top right
4. Click **Configure app password**
5. Copy the generated password

![](images/PIHOLE_API.png)

*Pi-hole v5.x and earlier:*
1. Log into your Admin Dashboard
2. Go to `Settings > API/Web Interface` tab
3. Click the **Show API token** button
4. A confirmation box appears; click **Yes, show API token** and copy the raw text string

## Gravwell Configuration

Gravwell uses its scripting interface (in the Pi-hole Kit) to request data from the Pi-hole API.

1. Set the `$PIHOLE_IP` macro to the IP address of your Pi-hole instance
2. Set the `$PIHOLE_PORT` macro to the port of your Pi-hole instance (usually 80)
3. Get the API key (or app password) as described in the Pi-hole Configuration section above
4. Set the `$PIHOLE_APIKEY` macro to the token from the previous step
5. Set the `$PIHOLE_TAG` macro to the desired tag name. The default is `pihole-queries`
6. If you used a tag name other than the default, update the extractor to your new tag name
7. Go to **Scripts** and enable the `PiHole Script` to run every 5 minutes with a cron schedule of `*/5 * * * *`

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/pihole-well.conf`
```ini
[Storage-Well "pihole"]
    Location=/opt/gravwell/storage/pihole
    Tags=pihole*
```
