# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- New feature or enhancement.

### Changed

- Changes to existing functionality.

### Fixed

- Bug fixes or issue resolutions.

## [1.2.0] - 2023-08-29

### Added

- Qontak SDK Integration: The Qontalk SDK is now seamlessly integrated into the Qontak ecosystem, allowing you to interact with Qontak's messaging and communication services effortlessly. 📞💬
- Finite State Machine (FSM) Support: The SDK includes a powerful Finite State Machine (FSM) library, making it easier than ever to manage complex state transitions in your applications. 🔄
- WhatsApp Integration: Send WhatsApp messages directly from your Go applications using the Qontalk Go SDK. 📱💬
- Direct WhatsApp Broadcast: Send direct WhatsApp broadcasts with custom parameters, including message templates, language settings, and buttons. 📢
- Improved Documentation: Our documentation has been enhanced to provide more detailed information and examples, making it easier for developers to get started. 📖

### Fixed

- Fixed several issues related to message interactions and interactive messages. 🐛
- Improved error handling and error messages for a smoother development experience. 🛠️
- Enhanced performance and reliability in various scenarios. 💪

## [1.1.0] - 2023-08-15

### Changed

- Introduced PutMultipart magic to MockRequestStrategy
- Revamped unit test for SendInteractiveMessage to embrace PutMultipart wizardry
- Nurtured test coverage and code tidying
- Sprinkled some developer love and bug fixes

## [1.0.0] - 2023-08-14

### Added

- Initial implementation of the project.

[Unreleased]: https://github.com/maskentir/qontalk/compare/v1.2.0...HEAD
[1.2.0]: https://github.com/maskentir/qontalk/releases/tag/v1.2.0
[1.1.0]: https://github.com/maskentir/qontalk/releases/tag/v1.1.0
[1.0.0]: https://github.com/maskentir/qontalk/releases/tag/v1.0.0
