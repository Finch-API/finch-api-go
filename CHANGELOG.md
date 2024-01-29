# Changelog

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
