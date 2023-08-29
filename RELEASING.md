# Releasing Guidelines

🚀 Thank you for working on a new release for the Qontak Go SDK! This document provides guidelines on how to create and publish a new release.

## Versioning

We use [Semantic Versioning (SemVer)](https://semver.org/) for version numbers. A version number consists of `MAJOR.MINOR.PATCH`:

- `MAJOR` version increment for incompatible changes.
- `MINOR` version increment for backward-compatible new features.
- `PATCH` version increment for backward-compatible bug fixes.

## Release Process

1. **Create a Release Branch:** Create a new branch for the release, typically named `release-x.y.z`, where `x.y.z` is the new version number.

2. **Update CHANGELOG.md:** Update the `CHANGELOG.md` file with details about the changes in this release. Include a section for the new version at the top, and use emoji for added visual appeal.

   ```markdown
   ## Version x.y.z (yyyy-mm-dd)

   ✨ New Features:

   - Describe new features or enhancements.

   🐛 Bug Fixes:

   - Describe bug fixes or issues resolved.

   📚 Documentation:

   - Mention any documentation updates.
   ```

3. **Review and Commit Changes:** Ensure that all changes for the release are committed and pushed to the release branch.

4. **Create a Pull Request:** Open a pull request from the release branch to the `main` branch. This pull request should include the updates to the `CHANGELOG.md` file.

5. **Review and Merge:** Have team members review the pull request. Once approved, merge it into the `main` branch.

6. **Tag the Release:** Create a new Git tag for the release using the version number. For example:

   ```sh
    git tag -a v-x.y.z -m "Version x.y.z"
    git push origin v-x.y.z
   ```

7. **Publish to Go.dev:** The release will be automatically published to Go.dev when the tag is pushed. Make sure your Go module follows Go's best practices.

8. **Announce the Release:** Make an announcement in relevant channels or documentation to let users know about the new release.

9. **Cleanup:** You can delete the release branch after the release is successfully tagged and merged.

## Example

Here's an example of a release branch workflow:

```sh
# Create a release branch
git checkout -b release-x.y.z

# Update CHANGELOG.md

# Commit and push changes
git add CHANGELOG.md
git commit -m "Prepare for release x.y.z"
git push origin release-x.y.z

# Create a pull request and merge

# Create a Git tag
git tag -a v-x.y.z -m "Version x.y.z"
git push origin v-x.y.z

# The release will be automatically published to Go.dev
```

Thank you for your contributions and for helping maintain the Qontak Go SDK!
