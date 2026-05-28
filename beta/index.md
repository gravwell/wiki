# Gravwell Beta




Gravwell Beta is publicly available through our Ubuntu repository and Docker images. It is updated prior to each production release, but beta features may require additional testing and improvements before they are ready for production. We have made the beta public because we are excited to share our new capabilities with you and get early feedback. 

## Installation and Upgrade

Switching from Stable to Beta is done by modifying your apt source repository (or our quick start instructions if installing from scratch).

### Upgrading:
Edit your `/etc/apt/sources.list.d/gravwell.list` file and replace `https://update.gravwell.io/debian/` with `https://update.gravwell.io/debianbeta/`. Then `apt update` and `apt upgrade` and you should be on the new release.

### Installing from scratch:

```
# Get our signing key
apt install apt-transport-https gnupg wget
wget -O /usr/share/keyrings/gravwell.asc https://update.gravwell.io/debian/update.gravwell.io.gpg.key

# Add the beta repository
echo 'deb [ arch=amd64 signed-by=/usr/share/keyrings/gravwell.asc ] https://update.gravwell.io/debianbeta/ community main' | sudo tee /etc/apt/sources.list.d/gravwell.list

# Install the package
apt update
apt install gravwell
```

### Docker

The Docker image is available at [gravwell/beta](https://hub.docker.com/r/gravwell/beta). You can change `gravwell/gravwell` to `gravwell/beta` in any of our Docker commands in the documentation and it should "just work."


## Thank You

Thank you for your interest and participation in the beta program. We couldn't do it without you!

Please send us feedback, bug reports, and especially show us cool stuff that you build with the new tools!

Please submit any bugs or feedback to [beta@gravwell.io](mailto:beta@gravwell.io)
