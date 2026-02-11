# Changelog

## 0.1.0-alpha.22 (2026-02-11)

Full Changelog: [v0.1.0-alpha.21...v0.1.0-alpha.22](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.21...v0.1.0-alpha.22)

### Bug Fixes

* **encoder:** correctly serialize NullStruct ([14c45d9](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/14c45d9d88796444f09848a4cddbee00beea3727))


### Chores

* **api:** minor updates ([727cf62](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/727cf6227f99d6773e527219f05b3d8a029cd600))

## 0.1.0-alpha.21 (2026-01-29)

Full Changelog: [v0.1.0-alpha.20...v0.1.0-alpha.21](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.20...v0.1.0-alpha.21)

### Bug Fixes

* **docs:** fix mcp installation instructions for remote servers ([30a6180](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/30a61803e1991fc3dc7a9aaace0de1e0cf54fa8a))

## 0.1.0-alpha.20 (2026-01-24)

Full Changelog: [v0.1.0-alpha.19...v0.1.0-alpha.20](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.19...v0.1.0-alpha.20)

### Features

* **client:** add a convenient param.SetJSON helper ([2d2e994](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/2d2e994392c219adda0b0327e8dc8e2a2380ba71))

## 0.1.0-alpha.19 (2026-01-17)

Full Changelog: [v0.1.0-alpha.18...v0.1.0-alpha.19](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.18...v0.1.0-alpha.19)

### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([afed92d](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/afed92d2baded64d2d840caa0f5fea2281ec75de))


### Chores

* add float64 to valid types for RegisterFieldValidator ([26e2a05](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/26e2a05f83b00561d71ec71281ad2bf70d87fc2c))
* **internal:** codegen related update ([d950a8c](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/d950a8c039194338c26f06117583e7105caf3d7d))
* **internal:** update `actions/checkout` version ([013a0f1](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/013a0f19badeea875c4900b93a051e0a623c4d3d))


### Documentation

* prominently feature MCP server setup in root SDK readmes ([dab4837](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/dab48373c2b0c86b72e94bff83fba6145d851613))

## 0.1.0-alpha.18 (2025-12-18)

Full Changelog: [v0.1.0-alpha.17...v0.1.0-alpha.18](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.17...v0.1.0-alpha.18)

### Bug Fixes

* skip usage tests that don't work with Prism ([0db880d](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/0db880dfa8f5ede576467e8f4de63a3d87a4c97e))

## 0.1.0-alpha.17 (2025-12-12)

Full Changelog: [v0.1.0-alpha.16...v0.1.0-alpha.17](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.16...v0.1.0-alpha.17)

### Features

* **encoder:** support bracket encoding form-data object members ([0e045d4](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/0e045d4ba5b8c5cb63eb15a355b23e5d37ada434))


### Bug Fixes

* **mcp:** correct code tool API endpoint ([a8ca0a1](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/a8ca0a1cf3d47bfce0a9add693424c0bfe22ed66))
* rename param to avoid collision ([2e69f68](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/2e69f68e56f00df74c31567ee7dab64b48b58926))


### Chores

* elide duplicate aliases ([3125b39](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/3125b395e8d3443030f33152557e8e219aeaf905))
* fix empty interfaces ([0effbf7](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/0effbf74fdcc286ed84e0a2d29a91f0236bceb44))
* **internal:** codegen related update ([7765969](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/7765969501e45f8292c31d5ec11ed758e050816b))

## 0.1.0-alpha.16 (2025-11-19)

Full Changelog: [v0.1.0-alpha.15...v0.1.0-alpha.16](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.15...v0.1.0-alpha.16)

### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([5fc3924](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/5fc3924160eebd3bb9744c4c2e1b944b69972cb9))


### Chores

* bump gjson version ([8c208a4](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/8c208a486b23b102c81091a725031cacb9c03fb6))
* **internal:** grammar fix (it's -&gt; its) ([d2dfd27](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/d2dfd27bc716213fd7384182032a80fbbe0cb2a6))

## 0.1.0-alpha.15 (2025-09-26)

Full Changelog: [v0.1.0-alpha.14...v0.1.0-alpha.15](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.14...v0.1.0-alpha.15)

### Bug Fixes

* bugfix for setting JSON keys with special characters ([5c8da88](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/5c8da88feea95177dd4b163c83a305db3a08f303))

## 0.1.0-alpha.14 (2025-09-20)

Full Changelog: [v0.1.0-alpha.13...v0.1.0-alpha.14](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.13...v0.1.0-alpha.14)

### Bug Fixes

* use slices.Concat instead of sometimes modifying r.Options ([9b15082](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/9b150820514684fc9a123b919fd16a4aad321a50))


### Chores

* bump minimum go version to 1.22 ([d21763e](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/d21763e142558298519938c5caaffdcd8c2b67ce))
* do not install brew dependencies in ./scripts/bootstrap by default ([a1dcbcc](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/a1dcbcc70b2d4f58e5ffe82aed8ea305f752bac4))
* update more docs for 1.22 ([0d687fa](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/0d687fa2c19c682e930e0e8acffc7ca930655c9a))

## 0.1.0-alpha.13 (2025-09-06)

Full Changelog: [v0.1.0-alpha.12...v0.1.0-alpha.13](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.12...v0.1.0-alpha.13)

### Bug Fixes

* **internal:** unmarshal correctly when there are multiple discriminators ([60e8285](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/60e8285b9b90523742e4244fbfbd5a0effc8e798))


### Chores

* **internal:** codegen related update ([7d325c1](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/7d325c1a01b939da023ffc178537df96d65b1a0a))

## 0.1.0-alpha.12 (2025-08-29)

Full Changelog: [v0.1.0-alpha.11...v0.1.0-alpha.12](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.11...v0.1.0-alpha.12)

### Features

* **client:** support optional json html escaping ([b8f9a1b](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/b8f9a1bd201c592af80fcba6dd6357dc1dceef72))


### Bug Fixes

* close body before retrying ([a7796f1](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/a7796f1cca720753d4947e4a7cb969b5b454b828))


### Chores

* **internal:** codegen related update ([5223c31](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/5223c31cb08ec246cd3500da85d57fcc4b9ba343))
* **internal:** update comment in script ([7527f10](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/7527f10aac2e8b08093ab28d7c74091fbfb4e656))
* update @stainless-api/prism-cli to v5.15.0 ([0bdd3a7](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/0bdd3a716f2f8b8a06d4743ad4f5e29ca30a854b))

## 0.1.0-alpha.11 (2025-08-06)

Full Changelog: [v0.1.0-alpha.10...v0.1.0-alpha.11](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.10...v0.1.0-alpha.11)

### Chores

* **api:** dedupe location types ([b7405d8](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/b7405d886e1862a5d6ca5f7410d70282e70dd1ed))

## 0.1.0-alpha.10 (2025-07-31)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-alpha.10](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.9...v0.1.0-alpha.10)

### Features

* **api:** manual updates ([b0c9769](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/b0c97697de6370bc07e42b85a5258668c487fd6e))

## 0.1.0-alpha.9 (2025-07-31)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Bug Fixes

* List&lt;List&gt; problem in java by naming dto.Event ([02d3780](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/02d378046072156abd642e94517c6610cdc4089e))

## 0.1.0-alpha.8 (2025-07-30)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Features

* publish to ruby gem and align namings ([de165c0](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/de165c026daa3c9e643e87a5783e380798fe3301))

## 0.1.0-alpha.7 (2025-07-30)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Features

* **api:** manual updates ([31c8902](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/31c8902b00727f2ab828267b922af2dadcc29792))
* Fix for ruby ([75a394f](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/75a394f16a13640305214af9e05de00afc5d9a60))
* re-order endpoints ([55d8616](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/55d861625befdc95054556099eaf68ab734709ff))

## 0.1.0-alpha.6 (2025-07-28)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** update via SDK Studio ([6314dff](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/6314dffd58878971bf845f8c813a8cf9b630a899))

## 0.1.0-alpha.5 (2025-07-28)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** update via SDK Studio ([ab64ee3](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/ab64ee3afdaf8fcab5cefac5cdbe945e2b3e6c47))
* **api:** update via SDK Studio ([ea84d18](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/ea84d181ffe7389e36c3c3b21a169ba75e7af51f))

## 0.1.0-alpha.4 (2025-07-28)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** update via SDK Studio ([a31bb40](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/a31bb408339b334e3f0f7bd3a7264bc601d28ab8))
* **api:** update via SDK Studio ([048888b](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/048888b95c5b94cc4ac1df48d7855f9f649c798e))

## 0.1.0-alpha.3 (2025-07-28)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** update via SDK Studio ([18d5c99](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/18d5c993da0818a4dab76fa2ff219a1da06612b5))
* **api:** update via SDK Studio ([d4eb36c](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/d4eb36c22fb244e5f26de564b41fc21f0e127edd))


### Chores

* update SDK settings ([4344236](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/434423654a167711947dcf5f0da0dfe35fb02fe9))

## 0.1.0-alpha.2 (2025-07-28)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** update via SDK Studio ([e1f2cc7](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/e1f2cc7354132ff685ef18c52a8bf5f0a5ca23ab))

## 0.1.0-alpha.1 (2025-07-28)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/nextbillion-ai/nextbillion-sdk-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** update via SDK Studio ([9c424d9](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/9c424d9050e74202a8700cbe5718ba23f3b7d5d6))
* **api:** update via SDK Studio ([0a17888](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/0a1788874010f39660221c420c6a2837ea2227d0))
* **api:** update via SDK Studio ([b8929c6](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/b8929c64fac2c9bf215df1fe9c1f7cdcb4ad1032))
* **api:** update via SDK Studio ([e2e4b82](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/e2e4b8288f5d7cecb930f8597ef6bd5c418d547d))
* **api:** update via SDK Studio ([6a41e55](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/6a41e551e6416273df505700ad7b3ebe00044482))
* **api:** update via SDK Studio ([3c4d6bc](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/3c4d6bc7abf4feb12fdb82dd9fd7999cce6dbd39))
* **api:** update via SDK Studio ([7faacc5](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/7faacc5f256fbab349451bec90d637b4e023e209))
* **api:** update via SDK Studio ([a00c15a](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/a00c15a015eeff11ffc336f2231df03fb74ad51f))
* **api:** update via SDK Studio ([4213eda](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/4213edaf02dfc55ea8a3b397209781f4fe9b9a29))
* **api:** update via SDK Studio ([7468e0e](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/7468e0efe996997d2883dc031b2eee016714b024))
* **api:** update via SDK Studio ([e6fed97](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/e6fed978548a6927679c493a69637fce4917edaa))
* **api:** update via SDK Studio ([3230980](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/323098001ee27515fc6e4bbbb4a3494590c90e7a))
* **api:** update via SDK Studio ([c17e5a9](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/c17e5a9dd4ee8c8cf4e50c8dc8611bedf2213bff))
* **api:** update via SDK Studio ([ec4bbca](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/ec4bbcaa328c600ca3d9b7121ab34ed6fa01f058))
* **api:** update via SDK Studio ([0dbcd8c](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/0dbcd8c1da76ff1403cd06d0cf0e0d44194e6da7))
* **api:** update via SDK Studio ([24df534](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/24df534597faddc09be2fd2e1aa07e2798f08985))
* **api:** update via SDK Studio ([a61f5ce](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/a61f5ce1ca4faecbb2f88749c9685c9aac50e1a8))
* **api:** update via SDK Studio ([3ea69b6](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/3ea69b6345695432f5eb46ebf3a0fa53e6158ce3))
* **api:** update via SDK Studio ([de6b8bb](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/de6b8bbcb9ecee39f2488a472a31bd53e44b4863))
* **api:** update via SDK Studio ([dbe1a0c](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/dbe1a0c49d9abd0a7bf4e30c266ce5bac588dc6d))
* **api:** update via SDK Studio ([e642474](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/e6424743f11ac7451fe75b645fe816d7255eaf53))
* **api:** update via SDK Studio ([4e8eeec](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/4e8eeec32517ea1181aef1611c1c346af15df2c6))
* **api:** update via SDK Studio ([3b2e32f](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/3b2e32f868117ebb4684b294e339132930368564))
* **api:** update via SDK Studio ([93043dc](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/93043dcebbe22657cf1848303fc6781429dc4879))
* **api:** update via SDK Studio ([43aca1a](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/43aca1a571aaebd4c5973c4f424cc639bf1173a0))
* **api:** update via SDK Studio ([9a3a201](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/9a3a2013d776486341ac9608e975f0b45b4f608e))
* **api:** update via SDK Studio ([9a12b2a](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/9a12b2af691e39e5311fcdcbbc3943b504ec0075))
* **api:** update via SDK Studio ([9e94f23](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/9e94f2325124a3b6fae28288a5edbc93085525ca))
* **api:** update via SDK Studio ([5066c9e](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/5066c9eb0c11dd5e3c3390b0e75cb6d29cf742a3))
* **api:** update via SDK Studio ([8ab99d3](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/8ab99d370efc6a4914d81e49b60b6e7819fc5002))
* **api:** update via SDK Studio ([3dae7fd](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/3dae7fd9b5964b11522a379f604a42675358ecac))
* **api:** update via SDK Studio ([a207084](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/a207084d6050c0a76a42564a613ef65ed609a1bf))


### Chores

* update SDK settings ([ab90e89](https://github.com/nextbillion-ai/nextbillion-sdk-go/commit/ab90e893145527b8692656b963f30b0019030655))
