# Changelog

## 0.21.7 (2024-07-15)

Full Changelog: [v0.21.6...v0.21.7](https://github.com/Finch-API/finch-api-go/compare/v0.21.6...v0.21.7)

### Bug Fixes

* **api:** remove unused fields ([#207](https://github.com/Finch-API/finch-api-go/issues/207)) ([1506c81](https://github.com/Finch-API/finch-api-go/commit/1506c8100af9c4e1c025b998c0d5273b04f56bb3))

## 0.21.6 (2024-07-11)

Full Changelog: [v0.21.5...v0.21.6](https://github.com/Finch-API/finch-api-go/compare/v0.21.5...v0.21.6)

### Chores

* **ci:** also run workflows for PRs targeting `next` ([#204](https://github.com/Finch-API/finch-api-go/issues/204)) ([f805b5e](https://github.com/Finch-API/finch-api-go/commit/f805b5e8c977534cb43ac84e528bd2c426388b3d))


### Documentation

* **examples:** update example values ([#206](https://github.com/Finch-API/finch-api-go/issues/206)) ([c307461](https://github.com/Finch-API/finch-api-go/commit/c307461221ab0d418388dc063e33e039d65bd1e1))

## 0.21.5 (2024-07-07)

Full Changelog: [v0.21.4...v0.21.5](https://github.com/Finch-API/finch-api-go/compare/v0.21.4...v0.21.5)

### Bug Fixes

* **internal:** fix MarshalJSON logic for interface elemnets ([#202](https://github.com/Finch-API/finch-api-go/issues/202)) ([1ec9f9d](https://github.com/Finch-API/finch-api-go/commit/1ec9f9dd99c5c7185916d243fbd3da918198dfd4))

## 0.21.4 (2024-07-04)

Full Changelog: [v0.21.3...v0.21.4](https://github.com/Finch-API/finch-api-go/compare/v0.21.3...v0.21.4)

### Chores

* **internal:** improve deserialization of embedded structs ([#200](https://github.com/Finch-API/finch-api-go/issues/200)) ([86e5a63](https://github.com/Finch-API/finch-api-go/commit/86e5a638588a40c42c95ee11988f49d2d046855f))

## 0.21.3 (2024-07-02)

Full Changelog: [v0.21.2...v0.21.3](https://github.com/Finch-API/finch-api-go/compare/v0.21.2...v0.21.3)

### Bug Fixes

* use slice instead of appending to r.Options ([#199](https://github.com/Finch-API/finch-api-go/issues/199)) ([57474b7](https://github.com/Finch-API/finch-api-go/commit/57474b70cb33d0140ccc89e64daf0f11f17038df))


### Chores

* sync openapi spec ([#197](https://github.com/Finch-API/finch-api-go/issues/197)) ([f5e5b4d](https://github.com/Finch-API/finch-api-go/commit/f5e5b4dc73ba31ef422dda8c8333516230f459c1))

## 0.21.2 (2024-06-28)

Full Changelog: [v0.21.1...v0.21.2](https://github.com/Finch-API/finch-api-go/compare/v0.21.1...v0.21.2)

### Chores

* gitignore test server logs ([#195](https://github.com/Finch-API/finch-api-go/issues/195)) ([845ee15](https://github.com/Finch-API/finch-api-go/commit/845ee15108e21fb48ca405404ad67f3fc9af86d3))

## 0.21.1 (2024-06-21)

Full Changelog: [v0.21.0...v0.21.1](https://github.com/Finch-API/finch-api-go/compare/v0.21.0...v0.21.1)

### Bug Fixes

* fix ExtraFields serialization / deserialization ([#193](https://github.com/Finch-API/finch-api-go/issues/193)) ([f6c540f](https://github.com/Finch-API/finch-api-go/commit/f6c540fad0f0a44b5ce6281038c35157b5fbfa8f))

## 0.21.0 (2024-06-07)

Full Changelog: [v0.20.1...v0.21.0](https://github.com/Finch-API/finch-api-go/compare/v0.20.1...v0.21.0)

### Features

* **api:** updates ([#191](https://github.com/Finch-API/finch-api-go/issues/191)) ([c3c38df](https://github.com/Finch-API/finch-api-go/commit/c3c38df08c5e00eb0d5d31b0fd3b74c0bedbff31))


### Bug Fixes

* fix port function for interface{} types ([#190](https://github.com/Finch-API/finch-api-go/issues/190)) ([88bd8b3](https://github.com/Finch-API/finch-api-go/commit/88bd8b389081de92add470c7fa969a3c8a676bd1))
* **internal:** fix the way that unions are deserialized in nested arrays ([#188](https://github.com/Finch-API/finch-api-go/issues/188)) ([66df98b](https://github.com/Finch-API/finch-api-go/commit/66df98b992c0716ccd458bf690e73bbcf009d7df))
* rename pay groups endpoint ([#192](https://github.com/Finch-API/finch-api-go/issues/192)) ([b7f911b](https://github.com/Finch-API/finch-api-go/commit/b7f911b06a0b24c2e6e632a05db89dc4493595e0))

## 0.20.1 (2024-05-28)

Full Changelog: [v0.20.0...v0.20.1](https://github.com/Finch-API/finch-api-go/compare/v0.20.0...v0.20.1)

### Chores

* **internal:** fix format script ([#186](https://github.com/Finch-API/finch-api-go/issues/186)) ([9dfbaf1](https://github.com/Finch-API/finch-api-go/commit/9dfbaf1b0a8ac60b6491f447b09b3edeb26e9899))

## 0.20.0 (2024-05-27)

Full Changelog: [v0.19.3...v0.20.0](https://github.com/Finch-API/finch-api-go/compare/v0.19.3...v0.20.0)

### Features

* better validation of path params ([#184](https://github.com/Finch-API/finch-api-go/issues/184)) ([3a0aa0e](https://github.com/Finch-API/finch-api-go/commit/3a0aa0e1349f79bc9272f35d557812d1d6a9fdcf))

## 0.19.3 (2024-05-23)

Full Changelog: [v0.19.2...v0.19.3](https://github.com/Finch-API/finch-api-go/compare/v0.19.2...v0.19.3)

### Chores

* **tests:** example client_id ([#182](https://github.com/Finch-API/finch-api-go/issues/182)) ([76e8855](https://github.com/Finch-API/finch-api-go/commit/76e8855242b49f57b964066e1752b48a7ed0eb7e))

## 0.19.2 (2024-05-17)

Full Changelog: [v0.19.1...v0.19.2](https://github.com/Finch-API/finch-api-go/compare/v0.19.1...v0.19.2)

### Chores

* **internal:** resolve spec formatting issue ([#181](https://github.com/Finch-API/finch-api-go/issues/181)) ([fe9770c](https://github.com/Finch-API/finch-api-go/commit/fe9770c1c1af1197d595100995f0d91bc416b4ac))
* **tests:** update some example values ([#179](https://github.com/Finch-API/finch-api-go/issues/179)) ([f1c7c5e](https://github.com/Finch-API/finch-api-go/commit/f1c7c5ef0118266d9f76e56a03cbe44c7582029c))

## 0.19.1 (2024-05-15)

Full Changelog: [v0.19.0...v0.19.1](https://github.com/Finch-API/finch-api-go/compare/v0.19.0...v0.19.1)

### Chores

* **internal:** support parsing other json content types ([#177](https://github.com/Finch-API/finch-api-go/issues/177)) ([a192156](https://github.com/Finch-API/finch-api-go/commit/a192156c1e2752b1512c2526db311d87f2a01659))

## 0.19.0 (2024-05-15)

Full Changelog: [v0.18.5...v0.19.0](https://github.com/Finch-API/finch-api-go/compare/v0.18.5...v0.19.0)

### Features

* propagate resource description field from stainless config to SDK docs ([#168](https://github.com/Finch-API/finch-api-go/issues/168)) ([bd5200b](https://github.com/Finch-API/finch-api-go/commit/bd5200bb3d54853ca97a6cf587cf6be839e08b61))


### Bug Fixes

* **api:** correct authentication methods type ([#173](https://github.com/Finch-API/finch-api-go/issues/173)) ([a21cfba](https://github.com/Finch-API/finch-api-go/commit/a21cfba96a4a9b78c72fbbbe38a7dc09dd3e65e6))
* fix reading the error body more than once ([#176](https://github.com/Finch-API/finch-api-go/issues/176)) ([e144ce8](https://github.com/Finch-API/finch-api-go/commit/e144ce809239cfdc4b219f2c56d3d0d0f12e7f20))


### Chores

* **docs:** add SECURITY.md ([#170](https://github.com/Finch-API/finch-api-go/issues/170)) ([8fe5bcb](https://github.com/Finch-API/finch-api-go/commit/8fe5bcbaa36df84be10e693a65aa7a79ed075c86))
* **internal:** add slightly better logging to scripts ([#175](https://github.com/Finch-API/finch-api-go/issues/175)) ([62478b8](https://github.com/Finch-API/finch-api-go/commit/62478b849ccb1b7a506c1582876f8dc67d7e06ef))
* **internal:** minor updates ([#172](https://github.com/Finch-API/finch-api-go/issues/172)) ([8e015d7](https://github.com/Finch-API/finch-api-go/commit/8e015d72b9416a3da962834427a2ed6e87ce7a82))
* **tests:** update client ID example value ([#174](https://github.com/Finch-API/finch-api-go/issues/174)) ([5c7b044](https://github.com/Finch-API/finch-api-go/commit/5c7b04462b916362bfb8189afaf421e7ed1d6d9c))

## 0.18.5 (2024-05-01)

Full Changelog: [v0.18.4...v0.18.5](https://github.com/Finch-API/finch-api-go/compare/v0.18.4...v0.18.5)

### Chores

* **internal:** fix bootstrap script ([#165](https://github.com/Finch-API/finch-api-go/issues/165)) ([2132c0b](https://github.com/Finch-API/finch-api-go/commit/2132c0b38d792916f3c84b7fd84162c3a6845eb7))

## 0.18.4 (2024-05-01)

Full Changelog: [v0.18.3...v0.18.4](https://github.com/Finch-API/finch-api-go/compare/v0.18.3...v0.18.4)

### Bug Fixes

* **test:** fix test github actions job ([#163](https://github.com/Finch-API/finch-api-go/issues/163)) ([29a071b](https://github.com/Finch-API/finch-api-go/commit/29a071b9f4e9dafbb7dcb8cec2cccd58fe9b8bcb))

## 0.18.3 (2024-04-30)

Full Changelog: [v0.18.2...v0.18.3](https://github.com/Finch-API/finch-api-go/compare/v0.18.2...v0.18.3)

### Chores

* **internal:** add scripts/test, scripts/mock and add ci job ([#161](https://github.com/Finch-API/finch-api-go/issues/161)) ([889fc86](https://github.com/Finch-API/finch-api-go/commit/889fc86a7d01eeb675be253d46bee2ab08b70fe2))

## 0.18.2 (2024-04-30)

Full Changelog: [v0.18.1...v0.18.2](https://github.com/Finch-API/finch-api-go/compare/v0.18.1...v0.18.2)

### Chores

* **internal:** add link to openapi spec ([#159](https://github.com/Finch-API/finch-api-go/issues/159)) ([974edac](https://github.com/Finch-API/finch-api-go/commit/974edacb68517fedf92813da1db02a463004474e))

## 0.18.1 (2024-04-26)

Full Changelog: [v0.18.0...v0.18.1](https://github.com/Finch-API/finch-api-go/compare/v0.18.0...v0.18.1)

### Bug Fixes

* make shared package public ([#158](https://github.com/Finch-API/finch-api-go/issues/158)) ([865ad75](https://github.com/Finch-API/finch-api-go/commit/865ad75e002cb6b0762284cfb9d39ffc6586b456))


### Chores

* **internal:** fix Port function for number and boolean enums ([#156](https://github.com/Finch-API/finch-api-go/issues/156)) ([fa91abc](https://github.com/Finch-API/finch-api-go/commit/fa91abcea7f5f86cbc7904408cf84d9f1f8e7128))
* **internal:** use actions/checkout@v4 for codeflow ([#154](https://github.com/Finch-API/finch-api-go/issues/154)) ([f454a5f](https://github.com/Finch-API/finch-api-go/commit/f454a5f142b42a262586223d85b9548a48239737))


### Documentation

* clarify behavior around employee_size ([#157](https://github.com/Finch-API/finch-api-go/issues/157)) ([fada270](https://github.com/Finch-API/finch-api-go/commit/fada270a7bd1ccbff6e8d78031bbeb57191fb865))

## 0.18.0 (2024-04-12)

Full Changelog: [v0.17.0...v0.18.0](https://github.com/Finch-API/finch-api-go/compare/v0.17.0...v0.18.0)

### Features

* **api:** updates ([#149](https://github.com/Finch-API/finch-api-go/issues/149)) ([9a16b33](https://github.com/Finch-API/finch-api-go/commit/9a16b3318028c537dc14e9588b6531de8f1508c5))
* **option:** add option to provide a raw request body ([#152](https://github.com/Finch-API/finch-api-go/issues/152)) ([f89654f](https://github.com/Finch-API/finch-api-go/commit/f89654fda20ddf470972806413ec3d7e2c407428))


### Chores

* **internal:** formatting ([#151](https://github.com/Finch-API/finch-api-go/issues/151)) ([9270f80](https://github.com/Finch-API/finch-api-go/commit/9270f80668099e2c4f663d13fc72e76c3a663610))

## 0.17.0 (2024-04-10)

Full Changelog: [v0.16.1...v0.17.0](https://github.com/Finch-API/finch-api-go/compare/v0.16.1...v0.17.0)

### Features

* **api:** updates ([#147](https://github.com/Finch-API/finch-api-go/issues/147)) ([6abcbc3](https://github.com/Finch-API/finch-api-go/commit/6abcbc3fd3fd5a8a1afe797cf4449d7187892fd0))


### Bug Fixes

* remove param.F wrapper from SandboxDirectoryNewParamsBody ([#148](https://github.com/Finch-API/finch-api-go/issues/148)) ([effba6e](https://github.com/Finch-API/finch-api-go/commit/effba6e2bb4201cbf37e6cb818d1a99207fac681))


### Chores

* **api:** improve descriptions ([#144](https://github.com/Finch-API/finch-api-go/issues/144)) ([3e9910d](https://github.com/Finch-API/finch-api-go/commit/3e9910d13c7ce4e329e45f5390d5f42f10539ba7))
* **internal:** implement Port function in apijson ([#146](https://github.com/Finch-API/finch-api-go/issues/146)) ([004dc21](https://github.com/Finch-API/finch-api-go/commit/004dc21b4fde460ad8dd7f9a6353bf976e8c605d))
* **webhooks:** fix type name ([0838adf](https://github.com/Finch-API/finch-api-go/commit/0838adf0754fc0db4654e5f9359568416b31fc5c))

## 0.16.1 (2024-04-01)

Full Changelog: [v0.16.0...v0.16.1](https://github.com/Finch-API/finch-api-go/compare/v0.16.0...v0.16.1)

### Chores

* **internal:** use a time zone less likely to conflict with the local one ([#142](https://github.com/Finch-API/finch-api-go/issues/142)) ([db30d1c](https://github.com/Finch-API/finch-api-go/commit/db30d1ccbafb8d2d7c8176d9d675f90aea3743e8))

## 0.16.0 (2024-03-29)

Full Changelog: [v0.15.0...v0.16.0](https://github.com/Finch-API/finch-api-go/compare/v0.15.0...v0.16.0)

### Features

* add IsKnown method to enums ([#135](https://github.com/Finch-API/finch-api-go/issues/135)) ([a12efdc](https://github.com/Finch-API/finch-api-go/commit/a12efdce5d7ac1c4486cccfee5c2b83ed99fa279))
* **client:** implement raw requests methods on client ([#139](https://github.com/Finch-API/finch-api-go/issues/139)) ([0469fdc](https://github.com/Finch-API/finch-api-go/commit/0469fdcad5e8323a6d61419f5190624160f29318))


### Chores

* **internal:** move pagination types to pagination package ([#140](https://github.com/Finch-API/finch-api-go/issues/140)) ([b4172a5](https://github.com/Finch-API/finch-api-go/commit/b4172a582939c4b111bbc947b8fd228fa6763ef6))
* **internal:** update generated pragma comment ([#134](https://github.com/Finch-API/finch-api-go/issues/134)) ([5acbf83](https://github.com/Finch-API/finch-api-go/commit/5acbf8307554cc006f6a6cd6c6bdf9c2de039eae))
* revert temporary commit ([4acd051](https://github.com/Finch-API/finch-api-go/commit/4acd051a9bbc90843d33bc7f99c5035670f6573a))
* temporary commit ([46f8947](https://github.com/Finch-API/finch-api-go/commit/46f8947f7277ba5c50e91a87c485078ed32235e3))


### Documentation

* fix typo in CONTRIBUTING.md ([#132](https://github.com/Finch-API/finch-api-go/issues/132)) ([dd445c7](https://github.com/Finch-API/finch-api-go/commit/dd445c72799beab386d9e5e5055381ea0bcfa184))
* fix typo in docstring for Null() ([#138](https://github.com/Finch-API/finch-api-go/issues/138)) ([da24dc4](https://github.com/Finch-API/finch-api-go/commit/da24dc425772a4ba905931de5bc231d8e90665fa))
* **readme:** consistent use of sentence case in headings ([#136](https://github.com/Finch-API/finch-api-go/issues/136)) ([2c2438f](https://github.com/Finch-API/finch-api-go/commit/2c2438ffd87570117e560da4dd4410c5c1cccc5b))
* **readme:** document file uploads ([#137](https://github.com/Finch-API/finch-api-go/issues/137)) ([4fea46d](https://github.com/Finch-API/finch-api-go/commit/4fea46dd3dd728c5ef594a3af2faeea248d3cbbf))

## 0.15.0 (2024-03-13)

Full Changelog: [v0.14.3...v0.15.0](https://github.com/Finch-API/finch-api-go/compare/v0.14.3...v0.15.0)

### Features

* set user-agent header by default when making requests ([#130](https://github.com/Finch-API/finch-api-go/issues/130)) ([bb66a49](https://github.com/Finch-API/finch-api-go/commit/bb66a49c0b01296111d749bc71a19aa4505bd4e0))

## 0.14.3 (2024-03-12)

Full Changelog: [v0.14.2...v0.14.3](https://github.com/Finch-API/finch-api-go/compare/v0.14.2...v0.14.3)

### Bug Fixes

* **client:** don't include ? in path unless necessary ([#129](https://github.com/Finch-API/finch-api-go/issues/129)) ([2096a3e](https://github.com/Finch-API/finch-api-go/commit/2096a3e2dedb875524b81203e7f7291eeb1cf672))
* fix String() behavior of param.Field ([#127](https://github.com/Finch-API/finch-api-go/issues/127)) ([1b81f9f](https://github.com/Finch-API/finch-api-go/commit/1b81f9f0b3010c738d02a926a2596fe5e0c1263b))

## 0.14.2 (2024-03-07)

Full Changelog: [v0.14.1...v0.14.2](https://github.com/Finch-API/finch-api-go/compare/v0.14.1...v0.14.2)

### Documentation

* **contributing:** add a CONTRIBUTING.md ([#125](https://github.com/Finch-API/finch-api-go/issues/125)) ([5e99acd](https://github.com/Finch-API/finch-api-go/commit/5e99acd0e78c9a70e93043654850132ae5764206))

## 0.14.1 (2024-03-07)

Full Changelog: [v0.14.0...v0.14.1](https://github.com/Finch-API/finch-api-go/compare/v0.14.0...v0.14.1)

### Chores

* **internal:** improve union deserialization logic ([#123](https://github.com/Finch-API/finch-api-go/issues/123)) ([c810956](https://github.com/Finch-API/finch-api-go/commit/c8109566b573efa31ff88cabe30a3adf872bbe71))

## 0.14.0 (2024-03-06)

Full Changelog: [v0.13.3...v0.14.0](https://github.com/Finch-API/finch-api-go/compare/v0.13.3...v0.14.0)

### Features

* implement public RawJSON of response structs ([#121](https://github.com/Finch-API/finch-api-go/issues/121)) ([d1161a4](https://github.com/Finch-API/finch-api-go/commit/d1161a4d59a1aca535eb7a5e903ce1b4d339aa02))


### Bug Fixes

* fix union deserialization for multiple objects ([#122](https://github.com/Finch-API/finch-api-go/issues/122)) ([b72853d](https://github.com/Finch-API/finch-api-go/commit/b72853d933486f5a5adf54bc566e9480360af5ff))


### Documentation

* deprecate old access token getter ([#119](https://github.com/Finch-API/finch-api-go/issues/119)) ([0818f5b](https://github.com/Finch-API/finch-api-go/commit/0818f5b24ffa4c03a33ae29cc62fda4e984a7774))

## 0.13.3 (2024-03-05)

Full Changelog: [v0.13.2...v0.13.3](https://github.com/Finch-API/finch-api-go/compare/v0.13.2...v0.13.3)

### Chores

* remove internal-only field ([#117](https://github.com/Finch-API/finch-api-go/issues/117)) ([d45a652](https://github.com/Finch-API/finch-api-go/commit/d45a652cb44781d2f02553e1e8d88350284d213e))

## 0.13.2 (2024-03-04)

Full Changelog: [v0.13.1...v0.13.2](https://github.com/Finch-API/finch-api-go/compare/v0.13.1...v0.13.2)

### Documentation

* update some doc strings ([#115](https://github.com/Finch-API/finch-api-go/issues/115)) ([23fc3e4](https://github.com/Finch-API/finch-api-go/commit/23fc3e490b3da3600d6d19d1d6ae268ec5d240aa))

## 0.13.1 (2024-03-04)

Full Changelog: [v0.13.0...v0.13.1](https://github.com/Finch-API/finch-api-go/compare/v0.13.0...v0.13.1)

### Chores

* update some doc strings ([#113](https://github.com/Finch-API/finch-api-go/issues/113)) ([4894bbe](https://github.com/Finch-API/finch-api-go/commit/4894bbef822bee7fd2be6627c4a044c4071342ba))

## 0.13.0 (2024-02-29)

Full Changelog: [v0.12.0...v0.13.0](https://github.com/Finch-API/finch-api-go/compare/v0.12.0...v0.13.0)

### Features

* **api:** make redirect_uri optional ([#111](https://github.com/Finch-API/finch-api-go/issues/111)) ([b2c907d](https://github.com/Finch-API/finch-api-go/commit/b2c907db4151cdac43f62025064f3fa5bfbf57da))


### Chores

* **internal:** update deps ([#109](https://github.com/Finch-API/finch-api-go/issues/109)) ([0836cd7](https://github.com/Finch-API/finch-api-go/commit/0836cd7b2f614e489b97c7c065aad70cca0e5090))


### Documentation

* **readme:** improve wording ([#112](https://github.com/Finch-API/finch-api-go/issues/112)) ([b8c363c](https://github.com/Finch-API/finch-api-go/commit/b8c363cb39889aca9268c1f4936e78971f0672a9))

## 0.12.0 (2024-02-20)

Full Changelog: [v0.11.1...v0.12.0](https://github.com/Finch-API/finch-api-go/compare/v0.11.1...v0.12.0)

### Features

* **api:** create new sandbox job ([#107](https://github.com/Finch-API/finch-api-go/issues/107)) ([0850d5a](https://github.com/Finch-API/finch-api-go/commit/0850d5a15e6967bcf0015e929f0e461becea3dd3))

## 0.11.1 (2024-02-15)

Full Changelog: [v0.11.0...v0.11.1](https://github.com/Finch-API/finch-api-go/compare/v0.11.0...v0.11.1)

### Chores

* **internal:** refactor release environment script ([#104](https://github.com/Finch-API/finch-api-go/issues/104)) ([4b9567a](https://github.com/Finch-API/finch-api-go/commit/4b9567a8565c54214ae2261033c8daaee610772f))

## 0.11.0 (2024-02-09)

Full Changelog: [v0.10.5...v0.11.0](https://github.com/Finch-API/finch-api-go/compare/v0.10.5...v0.11.0)

### Features

* **api:** create access token reads client opts if not provided ([#98](https://github.com/Finch-API/finch-api-go/issues/98)) ([2b02a6c](https://github.com/Finch-API/finch-api-go/commit/2b02a6ce9b2bd84ea3389ce2951a33c3fcb81abc))


### Chores

* **interal:** make link to api.md relative ([#100](https://github.com/Finch-API/finch-api-go/issues/100)) ([041fa16](https://github.com/Finch-API/finch-api-go/commit/041fa16178cb196e158361afc6a13849ecf19352))
* **internal:** adjust formatting ([#102](https://github.com/Finch-API/finch-api-go/issues/102)) ([6a39496](https://github.com/Finch-API/finch-api-go/commit/6a39496840312848a81061776724ac03f6e9fbba))
* **internal:** bump timeout threshold in tests ([#103](https://github.com/Finch-API/finch-api-go/issues/103)) ([9839c83](https://github.com/Finch-API/finch-api-go/commit/9839c834fe830b6e8df1d74b7d7b1e8387e51bdc))
* **internal:** minor pagination restructuring ([#101](https://github.com/Finch-API/finch-api-go/issues/101)) ([dba623c](https://github.com/Finch-API/finch-api-go/commit/dba623c6c9c936aa31516ab0664581b40e1639e9))

## 0.10.5 (2024-01-30)

Full Changelog: [v0.10.4...v0.10.5](https://github.com/Finch-API/finch-api-go/compare/v0.10.4...v0.10.5)

### Chores

* **internal:** support pre-release versioning ([#96](https://github.com/Finch-API/finch-api-go/issues/96)) ([5d95e24](https://github.com/Finch-API/finch-api-go/commit/5d95e24c9b4fb1749c76b86fe89000c6eca559d6))

## 0.10.4 (2024-01-29)

Full Changelog: [v0.10.3...v0.10.4](https://github.com/Finch-API/finch-api-go/compare/v0.10.3...v0.10.4)

### Chores

* **internal:** parse date-time strings more leniently ([#94](https://github.com/Finch-API/finch-api-go/issues/94)) ([8f4dafc](https://github.com/Finch-API/finch-api-go/commit/8f4dafc03d1cfe122aa17bef8684fd5d6b5a9d26))

## 0.10.3 (2024-01-29)

Full Changelog: [v0.10.2...v0.10.3](https://github.com/Finch-API/finch-api-go/compare/v0.10.2...v0.10.3)

### Bug Fixes

* **api:** update `employer_size` parameter to `employee_size` ([#92](https://github.com/Finch-API/finch-api-go/issues/92)) ([eb47c23](https://github.com/Finch-API/finch-api-go/commit/eb47c23a8e191df096dd042134487816de718bae))

## 0.10.2 (2024-01-29)

Full Changelog: [v0.10.1...v0.10.2](https://github.com/Finch-Api/finch-api-go/compare/v0.10.1...v0.10.2)

### Bug Fixes

* **api:** fix authentication_type enum ([#90](https://github.com/Finch-Api/finch-api-go/issues/90)) ([6e1fe9f](https://github.com/Finch-Api/finch-api-go/commit/6e1fe9f57f521caa6bc1ee50fd3d22029780c667))

## 0.10.1 (2024-01-26)

Full Changelog: [v0.10.0...v0.10.1](https://github.com/Finch-API/finch-api-go/compare/v0.10.0...v0.10.1)

### ⚠ BREAKING CHANGES

* fix oauth casing ([#88](https://github.com/Finch-API/finch-api-go/issues/88))

### Features

* **api:** add employer_contributions field ([#87](https://github.com/Finch-API/finch-api-go/issues/87)) ([cdc1578](https://github.com/Finch-API/finch-api-go/commit/cdc1578ced2aaa364fe2aa697ffa0ce53957799a))


### Bug Fixes

* **test:** avoid test failures when SKIP_MOCK_TESTS is not set ([#85](https://github.com/Finch-API/finch-api-go/issues/85)) ([f84fbe8](https://github.com/Finch-API/finch-api-go/commit/f84fbe85d4250a1613dcf1e3f017fcbd1f23bbb4))
* use content-type application/json for request to POST /auth/token ([#89](https://github.com/Finch-API/finch-api-go/issues/89)) ([d38b7f3](https://github.com/Finch-API/finch-api-go/commit/d38b7f3f02cdd64822ae5d8e9b3b50da3f0fbf03))


### Chores

* **ci:** rely on Stainless GitHub App for releases ([#86](https://github.com/Finch-API/finch-api-go/issues/86)) ([dbdd306](https://github.com/Finch-API/finch-api-go/commit/dbdd30614d8e4b9712d25635113bf5981677288e))
* **internal:** speculative retry-after-ms support ([#83](https://github.com/Finch-API/finch-api-go/issues/83)) ([8ace10b](https://github.com/Finch-API/finch-api-go/commit/8ace10b006dcc51463276aa8db205b216322c643))


### Refactors

* fix oauth casing ([#88](https://github.com/Finch-API/finch-api-go/issues/88)) ([02f4485](https://github.com/Finch-API/finch-api-go/commit/02f44851d8fe2ceedc6cc9393d5cc6aebae35517))

## 0.10.0 (2024-01-12)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/Finch-API/finch-api-go/compare/v0.9.0...v0.10.0)

### Features

* **webhooks:** add Unwrap method ([#80](https://github.com/Finch-API/finch-api-go/issues/80)) ([d21ede6](https://github.com/Finch-API/finch-api-go/commit/d21ede66329ce72c6fa7331030a85bbe35604503))


### Documentation

* **readme:** improve api reference ([#82](https://github.com/Finch-API/finch-api-go/issues/82)) ([ddfcab7](https://github.com/Finch-API/finch-api-go/commit/ddfcab7a06e1c13148e44fd2147c44f6134cdbe0))

## 0.9.0 (2024-01-11)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/Finch-API/finch-api-go/compare/v0.8.0...v0.9.0)

### Features

* **client:** hook up sandbox auth ([#79](https://github.com/Finch-API/finch-api-go/issues/79)) ([45aa9b5](https://github.com/Finch-API/finch-api-go/commit/45aa9b56a7adaea7c1233c518e3a514e75bb7b7f))
* remove redundant endpoint, add sandbox client options (not yet used) ([#78](https://github.com/Finch-API/finch-api-go/issues/78)) ([0257f91](https://github.com/Finch-API/finch-api-go/commit/0257f91a93ffc58944ad4d6ee9ec686bddbb92f9))


### Chores

* **internal:** rename unreleased connection status type ([#76](https://github.com/Finch-API/finch-api-go/issues/76)) ([feee2c9](https://github.com/Finch-API/finch-api-go/commit/feee2c9c69be7e1f541ff4b84ab9a6a96f16bac6))

## 0.8.0 (2024-01-09)

Full Changelog: [v0.7.1...v0.8.0](https://github.com/Finch-API/finch-api-go/compare/v0.7.1...v0.8.0)

### Features

* **api:** add method to create access token ([#72](https://github.com/Finch-API/finch-api-go/issues/72)) ([85b3dcd](https://github.com/Finch-API/finch-api-go/commit/85b3dcd71b0394a4a218054dad22f94690f10df2))
* **api:** add sandbox APIs ([#74](https://github.com/Finch-API/finch-api-go/issues/74)) ([3b861cf](https://github.com/Finch-API/finch-api-go/commit/3b861cf1bc2cf46128788ed6668b82def2b28cb8))


### Chores

* add .keep files for examples and custom code directories ([#73](https://github.com/Finch-API/finch-api-go/issues/73)) ([58e49c4](https://github.com/Finch-API/finch-api-go/commit/58e49c440527ee5b9ae41d34ca71d3987bad1bde))
* **internal:** bump license ([#69](https://github.com/Finch-API/finch-api-go/issues/69)) ([8a16565](https://github.com/Finch-API/finch-api-go/commit/8a165656fe271b949a9b92befef7a935ee58adc9))
* **test:** update examples ([#75](https://github.com/Finch-API/finch-api-go/issues/75)) ([905b7d3](https://github.com/Finch-API/finch-api-go/commit/905b7d3c860cb2a0b24916a5947c58f5c5787284))


### Refactors

* remove excess whitespace ([#71](https://github.com/Finch-API/finch-api-go/issues/71)) ([1c52791](https://github.com/Finch-API/finch-api-go/commit/1c52791b3dfbcf12cfee73db7b06053601f9ba0f))

## 0.7.1 (2023-12-19)

Full Changelog: [v0.7.0...v0.7.1](https://github.com/Finch-API/finch-api-go/compare/v0.7.0...v0.7.1)

### Documentation

* **options:** fix link to readme ([#66](https://github.com/Finch-API/finch-api-go/issues/66)) ([6a72906](https://github.com/Finch-API/finch-api-go/commit/6a729066438dde47f7f6c7f0c0e32bc4d6c912f9))
* replace &lt;br&gt; tags with newlines ([#68](https://github.com/Finch-API/finch-api-go/issues/68)) ([ac7713b](https://github.com/Finch-API/finch-api-go/commit/ac7713bde69a1b84c2dcc8922f4dc7b7b00d9906))

## 0.7.0 (2023-12-17)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/Finch-API/finch-api-go/compare/v0.6.0...v0.7.0)

### Features

* add GetAuthURL and GetAccessToken helpers ([#63](https://github.com/Finch-API/finch-api-go/issues/63)) ([c68b690](https://github.com/Finch-API/finch-api-go/commit/c68b6908ff6743be1ba25db78f82265d68f7400f))
* **api:** add `/jobs` endpoints ([#59](https://github.com/Finch-API/finch-api-go/issues/59)) ([9396454](https://github.com/Finch-API/finch-api-go/commit/9396454105e37487cb8d853c1c1351d1cd453d6c))
* **api:** add `client_type` and `connection_type` to introspection ([#60](https://github.com/Finch-API/finch-api-go/issues/60)) ([a812d5c](https://github.com/Finch-API/finch-api-go/commit/a812d5c4a5d222950380dbe3638581e65737fff8))
* **api:** add `lp` tax payer type enum value ([#62](https://github.com/Finch-API/finch-api-go/issues/62)) ([2c2291b](https://github.com/Finch-API/finch-api-go/commit/2c2291b11c5ddaa908e715699e68f62b5a37c471))


### Chores

* **ci:** run release workflow once per day ([#65](https://github.com/Finch-API/finch-api-go/issues/65)) ([425bca8](https://github.com/Finch-API/finch-api-go/commit/425bca89ec446c2b5ec32620c3075b822e1c8a29))

## 0.6.0 (2023-11-21)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/Finch-API/finch-api-go/compare/v0.5.0...v0.6.0)

### Features

* **api:** updates ([#56](https://github.com/Finch-API/finch-api-go/issues/56)) ([1ca1c45](https://github.com/Finch-API/finch-api-go/commit/1ca1c45bee3150532d0128ed85c3965d041c215c))

## 0.5.0 (2023-11-17)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/Finch-API/finch-api-go/compare/v0.4.0...v0.5.0)

### Features

* **api:** updates ([#53](https://github.com/Finch-API/finch-api-go/issues/53)) ([674fc9e](https://github.com/Finch-API/finch-api-go/commit/674fc9e3b583adeae9ebb064450c0352377c27c8))


### Bug Fixes

* make options.WithHeader utils case-insensitive ([#51](https://github.com/Finch-API/finch-api-go/issues/51)) ([afe25dd](https://github.com/Finch-API/finch-api-go/commit/afe25dd3305cf611c3f33d92b53ba2dcf6e25d64))


### Chores

* **internal:** update stats file ([#55](https://github.com/Finch-API/finch-api-go/issues/55)) ([83d4172](https://github.com/Finch-API/finch-api-go/commit/83d4172c58f8a4b6fc994b9af903283234125a27))


### Refactors

* do not include `JSON` fields when serialising back to json ([#54](https://github.com/Finch-API/finch-api-go/issues/54)) ([fabb56a](https://github.com/Finch-API/finch-api-go/commit/fabb56a7cfc98b9f45a89ff396cbbefb3fd0d3b4))

## 0.4.0 (2023-11-05)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/Finch-API/finch-api-go/compare/v0.3.0...v0.4.0)

### Features

* **client:** allow binary returns ([#48](https://github.com/Finch-API/finch-api-go/issues/48)) ([6eb6aa0](https://github.com/Finch-API/finch-api-go/commit/6eb6aa090193fdfde038df055397dba5fe6b1108))


### Documentation

* **readme:** improve example snippets ([#50](https://github.com/Finch-API/finch-api-go/issues/50)) ([f4a7114](https://github.com/Finch-API/finch-api-go/commit/f4a7114b04f13d50fa27340eb740215decc8233c))

## 0.3.0 (2023-10-31)

Full Changelog: [v0.2.2...v0.3.0](https://github.com/Finch-API/finch-api-go/compare/v0.2.2...v0.3.0)

### Features

* **github:** include a devcontainer setup ([#46](https://github.com/Finch-API/finch-api-go/issues/46)) ([ea59cb7](https://github.com/Finch-API/finch-api-go/commit/ea59cb7ecb453d85da5328c9d9dc4f6690eeb353))

## 0.2.2 (2023-10-27)

Full Changelog: [v0.2.1...v0.2.2](https://github.com/Finch-API/finch-api-go/compare/v0.2.1...v0.2.2)

### Chores

* **internal:** update gitignore ([#44](https://github.com/Finch-API/finch-api-go/issues/44)) ([946f6dd](https://github.com/Finch-API/finch-api-go/commit/946f6dddfca8108bdcc4c0d4a9c82f41b7ffa2ae))

## 0.2.1 (2023-10-24)

Full Changelog: [v0.2.0...v0.2.1](https://github.com/Finch-API/finch-api-go/compare/v0.2.0...v0.2.1)

## 0.2.0 (2023-10-24)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/Finch-API/finch-api-go/compare/v0.1.0...v0.2.0)

### Features

* **client:** adjust retry behavior ([#40](https://github.com/Finch-API/finch-api-go/issues/40)) ([a2b73b7](https://github.com/Finch-API/finch-api-go/commit/a2b73b75383178d72bff7efa04bfc8e9c98143b4))

## 0.1.0 (2023-10-17)

Full Changelog: [v0.0.8...v0.1.0](https://github.com/Finch-API/finch-api-go/compare/v0.0.8...v0.1.0)

### Features

* make webhook headers case insensitive ([#38](https://github.com/Finch-API/finch-api-go/issues/38)) ([fb58c4e](https://github.com/Finch-API/finch-api-go/commit/fb58c4e203d49a3e0a0f25dbf1ab6b84ef654ae8))


### Documentation

* **api.md:** improve formatting of webhook helpers ([#36](https://github.com/Finch-API/finch-api-go/issues/36)) ([d1dd461](https://github.com/Finch-API/finch-api-go/commit/d1dd461e561d24f55671d2b5404672c3790b7350))
* organisation -&gt; organization (UK to US English) ([#39](https://github.com/Finch-API/finch-api-go/issues/39)) ([bbcc80c](https://github.com/Finch-API/finch-api-go/commit/bbcc80ca9510d749967c2b74d8606b61104fa211))

## 0.0.8 (2023-10-12)

Full Changelog: [v0.0.7...v0.0.8](https://github.com/Finch-API/finch-api-go/compare/v0.0.7...v0.0.8)

### Chores

* **internal:** rearrange client arguments ([#33](https://github.com/Finch-API/finch-api-go/issues/33)) ([81460c8](https://github.com/Finch-API/finch-api-go/commit/81460c8360093c145986476bcd9dfb3e38a98d9e))

## 0.0.7 (2023-10-11)

Full Changelog: [v0.0.6...v0.0.7](https://github.com/Finch-API/finch-api-go/compare/v0.0.6...v0.0.7)

### Bug Fixes

* correct benfits to benefits ([#31](https://github.com/Finch-API/finch-api-go/issues/31)) ([9d357a8](https://github.com/Finch-API/finch-api-go/commit/9d357a8560149d857698c8add0591cd1004ee9c4))

## 0.0.6 (2023-10-11)

Full Changelog: [v0.0.5...v0.0.6](https://github.com/Finch-API/finch-api-go/compare/v0.0.5...v0.0.6)

## 0.0.5 (2023-10-03)

Full Changelog: [v0.0.4...v0.0.5](https://github.com/Finch-API/finch-api-go/compare/v0.0.4...v0.0.5)

### Features

* **api:** add `/forward` endpoint and other updates ([#25](https://github.com/Finch-API/finch-api-go/issues/25)) ([f16872d](https://github.com/Finch-API/finch-api-go/commit/f16872dc5025c7c6d88673f65a1b5a14a3fa3e65))


### Bug Fixes

* prevent index out of range bug during auto-pagination ([#23](https://github.com/Finch-API/finch-api-go/issues/23)) ([24b06e2](https://github.com/Finch-API/finch-api-go/commit/24b06e23bcfcaf13d5bdeae175348b71bdd0176d))


### Chores

* **docs:** adjust some docstrings ([#26](https://github.com/Finch-API/finch-api-go/issues/26)) ([0dd63e2](https://github.com/Finch-API/finch-api-go/commit/0dd63e219897c5dce9fd26d2aa7380979c9aeb4c))
* **docs:** adjust some docstrings ([#28](https://github.com/Finch-API/finch-api-go/issues/28)) ([23fa6ec](https://github.com/Finch-API/finch-api-go/commit/23fa6ecd1e5633600590d0cea21c8f08eb1fa1f0))
* **tests:** update test examples ([#27](https://github.com/Finch-API/finch-api-go/issues/27)) ([13f64a3](https://github.com/Finch-API/finch-api-go/commit/13f64a3dd1124e7c54f7def4e96df55b41c8c26a))

## 0.0.4 (2023-09-26)

Full Changelog: [v0.0.3...v0.0.4](https://github.com/Finch-API/finch-api-go/compare/v0.0.3...v0.0.4)

### Features

* **ci:** add reviewers ([#15](https://github.com/Finch-API/finch-api-go/issues/15)) ([92f2916](https://github.com/Finch-API/finch-api-go/commit/92f2916b6be96031166746330d98063607709878))
* improve retry behavior on context deadline ([#18](https://github.com/Finch-API/finch-api-go/issues/18)) ([c1c477e](https://github.com/Finch-API/finch-api-go/commit/c1c477ed80d2578565b40eda8618701826518740))

## 0.0.4 (2023-09-25)

Full Changelog: [v0.0.3...v0.0.4](https://github.com/Finch-API/finch-api-go/compare/v0.0.3...v0.0.4)

### Features

* **ci:** add reviewers ([#15](https://github.com/Finch-API/finch-api-go/issues/15)) ([83dc2d2](https://github.com/Finch-API/finch-api-go/commit/83dc2d2969ba5199211e8087055bdb8b79e63ebd))
* improve retry behavior on context deadline ([#18](https://github.com/Finch-API/finch-api-go/issues/18)) ([6170dbc](https://github.com/Finch-API/finch-api-go/commit/6170dbce545d451301ca84d0fb49c56ea925c2bf))

## 0.0.3 (2023-09-22)

Full Changelog: [v0.0.2...v0.0.3](https://github.com/Finch-API/finch-api-go/compare/v0.0.2...v0.0.3)

### Features

* retry on 408 Request Timeout ([#7](https://github.com/Finch-API/finch-api-go/issues/7)) ([6075488](https://github.com/Finch-API/finch-api-go/commit/6075488158fc1ea2db4bf1d280a3479ad916bd79))


### Bug Fixes

* **core:** improve retry behavior and related docs ([#9](https://github.com/Finch-API/finch-api-go/issues/9)) ([7107ce5](https://github.com/Finch-API/finch-api-go/commit/7107ce5b8411fa3b44c1de961ddaab4b8f17cf24))


### Chores

* **api:** remove deprecated & unused ATS API ([#12](https://github.com/Finch-API/finch-api-go/issues/12)) ([213405d](https://github.com/Finch-API/finch-api-go/commit/213405d52af5efd4593568a245ba27432e820033))


### Documentation

* **api.md:** rename Top Level to client name ([#14](https://github.com/Finch-API/finch-api-go/issues/14)) ([3030c73](https://github.com/Finch-API/finch-api-go/commit/3030c7354f856084304a175ac637c99e0517b0f7))
* **README:** fix variable names in some examples ([#13](https://github.com/Finch-API/finch-api-go/issues/13)) ([4b23c28](https://github.com/Finch-API/finch-api-go/commit/4b23c2831725bbe854d88fb2e399dd16fa88a8b0))

## 0.0.2 (2023-09-14)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/Finch-API/finch-api-go/compare/v0.0.1...v0.0.2)

### Features

* **init:** initial commit ([9f33baa](https://github.com/Finch-API/finch-api-go/commit/9f33baa96aa2259baea5cd2a5fe00c5b57d17bb2))


### Bug Fixes

* **core:** add null check to prevent segfault when canceling context ([#4](https://github.com/Finch-API/finch-api-go/issues/4)) ([39c282f](https://github.com/Finch-API/finch-api-go/commit/39c282fc067c85eb13f4aeef8a5623385bc466e4))


### Chores

* **internal:** improve reliability of cancel delay test ([#5](https://github.com/Finch-API/finch-api-go/issues/5)) ([381707d](https://github.com/Finch-API/finch-api-go/commit/381707dbc242db4653506bc4769c5193246d2645))
