# Changelog

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
