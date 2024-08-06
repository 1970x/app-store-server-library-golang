# Apple App Store Server Golang Library
The Golang server library for the [App Store Server API](https://developer.apple.com/documentation/appstoreserverapi) and [App Store Server Notifications](https://developer.apple.com/documentation/appstoreservernotifications).

## Table of Contents
1. [Installation](#installation)
2. [Documentation](#documentation)
3. [Usage](#usage)
4. [Support](#support)

## Installation

#### Requirements

- Golang


## Documentation

[WWDC Video](https://developer.apple.com/videos/play/wwdc2023/10143/)

### Obtaining an In-App Purchase key from App Store Connect

To use the App Store Server API or create promotional offer signatures, a signing key downloaded from App Store Connect is required. To obtain this key, you must have the Admin role. Go to Users and Access > Integrations > In-App Purchase. Here you can create and manage keys, as well as find your issuer ID. When using a key, you'll need the key ID and issuer ID as well.

### Obtaining Apple Root Certificates

Download and store the root certificates found in the Apple Root Certificates section of the [Apple PKI](https://www.apple.com/certificateauthority/) site. Provide these certificates as an array to a SignedDataVerifier to allow verifying the signed data comes from Apple.

## Usage

### API Usage

```golang

```

### Verification Usage

```golang

```

### Receipt Usage

```golang

```

### Promotional Offer Signature Creation

```java

```

## Support

Only the latest major version of the library will receive updates, including security updates. Therefore, it is recommended to update to new major versions.
