# Gravwell Beta Program

Greetings and salutations!

You are reading this message because you are in the early access group for Gravwell Beta!  We sincerely thank you for your participation and look forward to feedback and bug reports.

Please submit any bugs or feedback to [beta@gravwell.io](mailto:beta@gravwell.io)

Thank you!

## Current "State of the Beta"

We're preparing for the release of Gravwell 4.2.0. The primary new feature of this release is the Data Explorer, which makes it easier to play with your data in a point-and-click manner.

### Desired Testing

Testing desires for this sprint (in order of priority)

* Data Explorer - Try it with as many different data sources as you can. Add filters, remove filters, pivot into other query modes.
* Query Studio - Make sure the query editing box is usable and robust, move between many queries and tags, check the new formatting with other renderers.
* Per Tag Accelerators - Add some [tag specific accelerator definitions](#!configuration/accelerators.md#Accelerating_Specific_Tags).

## Installation and Upgrade

We're very excited to say this build is now available for your use and testing. We have created a new Ubuntu repository and Docker images. Switching from Stable to Beta is done by modifying your apt source repository (or our quick start instructions if installing from scratch).

### Upgrading:
Edit your `/etc/apt/sources.list.d/gravwell.list` file and replace `https://update.gravwell.io/debian/` with `https://update.gravwell.io/debianbeta/`. Then `apt update` and `apt upgrade` and you should be on the new release.

### Installing from scratch:

```
curl https://update.gravwell.io/debian/update.gravwell.io.gpg.key | sudo apt-key add -
echo 'deb [ arch=amd64 ] https://update.gravwell.io/debianbeta/ community main' | sudo tee /etc/apt/sources.list.d/gravwell.list
sudo apt-get install apt-transport-https
sudo apt-get update
sudo apt-get install gravwell
```

### Docker

The Docker image is available at [gravwell/beta](https://hub.docker.com/r/gravwell/beta). You can change `gravwell/gravwell` to `gravwell/beta` in any of our Docker commands in the documentation and it should "just work."


## Thank You

We are very excited about the new capabilities that Gravwell brings. Thank you for your interest and participation in the beta program. We couldn't do it without you!

Please send us feedback, bug reports, and especially show us cool stuff that you build with the new tools!
