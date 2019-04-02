# Rails n' Buffalo

This repository is primarily to investigate the benefits, development practices, and difficulties between [Rails](https://rubyonrails.org) and [Buffalo.](https://gobuffalo.io/) In order to do so a simple "feature flag" application will be built in both frameworks to attempt (largely in vain) to compare them both.

## Project Setup

Each directory will maintain its own `README.md` to facilitate setup instructions and how to run each application.

## Problem Statement

Build a web application that can create, manage, and delete Feature Flags. Feature Flags are boolean values that determine if a requested feature is "On" or "Off."

Each Feature Flag will belong to a Project, which in turn belongs to a User.

Both applications must use Postgres for its Database, and to be considered complete should also be deployable in a Docker container.
