# Contributing

General guidelines for contributing to this project.

## Issues

If there is no issue that suits the change you want to make, feel free to open an issue with a descirption and reason for the change.
Milestones and labels will be added by the owners of the project.

## Branches

Banch names are up to the person creating the branch, there is no specific standard to follow.

## Pull requests

To contribute, fork the repository and make the updates in your forked version.
Once complete, create a pull request to merge your changes from your forked repository into the main repository.
Pull requests require an approval by the owners of, or contributors to the project.
The merging strategy is to rebase and merge without a merge commit to keep the main branch clean.
There is no SLA on how long it will take your pull request to get merged.
Once the pull request is approved, either the owner or contributors will merge the PR when appropriate.
Approved pull requests will be rebased and managed once approval has been obtained, so at the point of approval, you can consider the change complete.

## Commits

The scope of each commit should be small and specific.
This will result in commits that are easy to review and will effectively show the changes to the project over time.
Commit messages should use imperitive mood and explain the change being made in the commit.
The reason for the commit strategy is that there is no CHANGELOG.md for this project.
It is extra admin that the project's maintainers do not want to deal with.
The commit tree thus becomes the effective changelog and the details of those commit messages become important in tracking the changes over time.

1. Capitalize the first word and do not end in punctuation.
2. Use imperative mood in the subject line.
   Imperative mood gives the tone you are giving an order or request.
   Example: `Add fix for dark mode toggle state`
3. Specify the type of commit.
   It is recommended and can be even more beneficial to have a consistent set of words to describe your changes.
   Example: Add, change, fix, remove etc.
4. There should only be on sentence and it should ideally be no longer than 50 characters, but if length is needed for context, it can be motivated.
5. Be direct, try to state the change as a fact based on the content of the commit.

## Releases

The project follows [semantic versioning][1].
Releases are created when a new tag is pushed with a new version.

[1]: https://semver.org/
