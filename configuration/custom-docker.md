# Custom Docker Deployments

Most Gravwell components are deployed as statically compiled executables which are suitable for execution on most modern Linux hosts, which also enables easy Docker deployment.  Gravwell engineers use docker extensively for developement, testing, and internal deployment.  Docker also enables customers to rapidly standup, teardown, and otherwise manage a large Gravwell deployment.  To enable customers to get running quickly in a Docker deployment, we have provided sample Dockerfiles for both cluster and single edition SKUs.

The complete set of Dockerfiles can be found [here](https://update.gravwell.io/files/docker_buildfiles_ad05723a547d31ee57ed8880bb5ef4e9.tar.bz2) with an MD5 checksum of ad05723a547d31ee57ed8880bb5ef4e9.

## Building Docker Containers

Building Docker containers using the provided Dockerfiles is extremely easy.  Gravwell docker deployments utilize the extremely small busybox base container, enabling very small continers.

### Dockerfile

Modifying the docker file to suite specific deployment requirements requires very little work.  The standard gravwell docker file uses a small startup script to check whether or not we need to regenerate X509 certificates at startup.  If your deployment has valid certificates, you can start Gravwell binaries directly and remove some of the utilities in /opt/gravwell/bin that are deployed by the installer (namely gencert and crashreport).

The base Dockerfile for a single instance:
```
FROM busybox
MAINTAINER support@gravwell.io
ARG INSTALLER=gravwell_installer.sh
COPY $INSTALLER /tmp/installer.sh
COPY start.sh /tmp/start.sh
RUN /bin/sh /tmp/installer.sh --no-questions
RUN rm -f /tmp/installer.sh
RUN mv /tmp/start.sh /opt/gravwell/bin/
CMD ["/bin/sh", "/opt/gravwell/bin/start.sh"]
```

The basic startup script:
```
#!/bin/sh

# unless environment variable says no, generate new SSL certs
if [ "$NO_SSL_GEN" == "true" ]; then
	echo "Skipping SSL certificate generation"
else
	/opt/gravwell/bin/gencert -key-file /opt/gravwell/etc/key.pem -cert-file /opt/gravwell/etc/cert.pem -host 127.0.0.1
	if [ "$?" != "0" ]; then
		echo "Failed to generate certificates"
		exit -1
	fi
fi

#fire up the indexer and webserver processes and wait
/opt/gravwell/bin/gravwell_indexer -config-override /opt/gravwell/etc/ &
/opt/gravwell/bin/gravwell_webserver -config-override /opt/gravwell/etc/ &
wait
```

## Using Environment Variables

The standard docker startup script looks for the environment variable `NO_SSL_GEN` and will skip X509 certificate generation if set to "true".  If your deployment is injecting valid certificates be sure to include the argument `-e NO_SSL_GEN=true` when starting the container.

The indexer, webserver, and ingesters can also have some configuration parameters set via environment variables rather than config file if desired. See the [Environment Variables](environment-variables.md) documentation for more details.

## Sample Dockerfile For Ingesters

Gravwell is continually releasing new ingesters and components, but may not always have a Dockerfile for every ingester.  However, Dockerfiles are pretty straight forward and easy to modify.  Below is an example Dockerfile which builds a Docker container via the SimpleRelay ingester.

```
FROM busybox
MAINTAINER support@gravwell.io
ARG INSTALLER=gravwell_installer.sh
COPY $INSTALLER /tmp/installer.sh
RUN /bin/sh /tmp/installer.sh --no-questions
RUN rm -f /tmp/installer.sh
CMD ["/opt/gravwell/bin/gravwell_simple_relay"]
```

To build the container, copy the simple relay installer to the same working directory as the Docker file and execute the following command:
```
docker build --ulimit nofile=32000:32000 --compress --build-arg INSTALLER=gravwell_simple_relay_installer_2.0.sh --no-cache --tag gravwell:simple_relay_2.0.0 .
```
