# Gravwell Releases

## Release Stages

### Alpha

Alpha releases are generally for internal use only, but they may occasionally be made available on a Gravwell-hosted instance in order to solicit early feedback from our customers. Alpha features may be incomplete and may not be documented. Alphas have no release support obligations. 

### Beta

Beta releases are publicly available. Beta features may require additional testing and improvements before they are ready for a production release. You can read more on our [Gravwell Beta Program](beta/index) page. Betas have no release support obligations.

### Production

Production releases are general availability releases that are ready for production use. For details on release support for production builds, see [Release Support](release-support) below.

For detailed information about licensing, see the [Gravwell License System](license/license) page. You can find a table summarizing the features supported by each license type on that page.

## Release Versioning

Gravwell's versioning follows [Semantic Versioning 2.0.0](https://semver.org/). When a new version of Gravwell is released, it has the format MAJOR.MINOR.PATCH-HOTFIX where…
* MAJOR is an integer. 
  * It is incremented when the release contains a high-risk change or a backward-incompatible API change.
* MINOR is an integer. 
  * It is incremented when the release contains a medium-risk change or a backward-compatible API change.
  * It is reset to zero if MAJOR is incremented.
* PATCH is an integer.
  * It is incremented when the release contains only low-risk changes.
  * It is reset to zero if MAJOR is incremented or if MINOR is incremented
* HOTFIX is an integer.
  * It is incremented when there is a failure in the build pipeline.
  * It is reset to one if MAJOR, MINOR, or PATCH are incremented.

(release-support)=
## Release Support

Maintenance and security updates are applied by upgrading to the latest patch release and are not backported to previous versions. 
