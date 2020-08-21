# Gravwell Beta Program

Greetings and salutations!

You are reading this message because you are in the early access group for Gravwell Beta!  We sincerely thank you for your participation and look forward to feedback and bug reports.

Please submit any bugs or feedback:

* https://www.gravwell.io/beta
* Email to beta@gravwell.io

Thank you!

## Current "State of the Beta"

### Desired Testing

Testing desires for this sprint (in order of priority)

* Browsing, installing, and managing kits
* Search template creation and use to create "investigative" dashboards
* Playbooks - New features are coming to enrich playbooks a bit more, but testing of creating and using them as markdown READMEs is the desire for this phase
* Macros
* Templates


## Installation and Upgrade

We're very excited to say this build is now available for your use and testing. We have created a new ubuntu repository and Docker images. Switching from Stable to Beta is done by modifying your apt source respository (or our quick start instructions if installing from scratch).

### Upgrading:
Edit your `/etc/apt/sources.list.d/gravwell.list` file and substitute `https://update.gravwell.io/debian/` for `https://update.gravwell.io/debianbeta/`. Then `apt update` and `apt upgrade` and you should be on the new release.

### Installing from scratch:

```
curl https://update.gravwell.io/debian/update.gravwell.io.gpg.key | sudo apt-key add -
echo 'deb [ arch=amd64 ] https://update.gravwell.io/debianbeta/ community main' | sudo tee /etc/apt/sources.list.d/gravwell.list
sudo apt-get install apt-transport-https
sudo apt-get update
sudo apt-get install gravwell
```

### Docker

The Docker images is available at [gravwell/beta](https://hub.docker.com/r/gravwell/beta). You can substitue `gravwell/gravwell` with `gravwell/beta` in any of the docker documentation and it should "just work."




## Thank You

We are very excited about the new capabilities that Gravwell brings. Thank you for your interest and participation in the beta program. We couldn't do it without you!

Please send us feedback, bug reports, and especially show us cool stuff that you build with the new tools!
