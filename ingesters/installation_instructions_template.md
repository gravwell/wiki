To install the Debian package, make sure the Gravwell Debian repository is configured [as described in the quickstart](debian_repo). Then run the following command as root:

::: {parsed-literal}
apt update && apt install {{package}}
:::

To install the Redhat package, make sure the Gravwell Redhat repository is configured [as described in the quickstart](redhat_repo). Then run the following command as root:

::: {parsed-literal}
yum install {{package}}
:::

To install via the standalone shell installer, download the installer from the [downloads page](/quickstart/downloads), then run the following command as root, replacing X.X.X with the appropriate version:

::: {parsed-literal}
bash {{standalone}}_installer_X.X.X.sh
:::

You may be prompted for additional configuration during the installation.

{{ 'The Docker image is available on [Dockerhub](https://hub.docker.com/r/gravwell/{}).'.format(dockername) if dockername else "There is currently no Docker image for this ingester" }}
