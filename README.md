# go-crypto-market-ui

Command-line utility to track cryptocurrencies in realtime.

![Dashboard](https://raw.githubusercontent.com/nuxy/go-crypto-market-ui/master/preview/dashboard.png)

## Features

- Configuration based, self contained binary (no dependencies).
- Keyboard-interactive [dashboard terminal](#examples) interface.
- [CoinMarketCap API](https://coinmarketcap.com/api/documentation/v1) v1 integration (Quotes and Metadata).
- Localized UI via templates (locales: [de](https://github.com/nuxy/go-crypto-market-ui/tree/master/locales/de.yaml), [fr](https://github.com/nuxy/go-crypto-market-ui/tree/master/locales/fr.yaml), [es](https://github.com/nuxy/go-crypto-market-ui/tree/master/locales/es.yaml), [it](https://github.com/nuxy/go-crypto-market-ui/tree/master/locales/it.yaml), [ja](https://github.com/nuxy/go-crypto-market-ui/tree/master/locales/ja.yaml), [ko](https://github.com/nuxy/go-crypto-market-ui/tree/master/locales/ko.yaml), [ru](https://github.com/nuxy/go-crypto-market-ui/tree/master/locales/ru.yaml), [vi](https://github.com/nuxy/go-crypto-market-ui/tree/master/locales/vi.yaml), [zh](https://github.com/nuxy/go-crypto-market-ui/tree/master/locales/zh.yaml))
- Configurable money and currency formatting.
- Extendable to support new market API [services](https://github.com/nuxy/go-crypto-market-ui/tree/master/lib/services) and [languages](https://github.com/nuxy/go-crypto-market-ui/tree/master/locales).

## Dependencies

- [Go](https://golang.org)

## Installation

Install the package using [go get](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them).

    $ go get github.com/nuxy/go-crypto-market-ui

## Building the source

Package build changes that includes static imports using [pkger](https://github.com/markbates/pkger).

    $ pgker -import /locales

Install the new build using [gmake](https://www.gnu.org/software/make).

    $ make install

Cross-compile to support [Windows](https://golang.org/dl/go1.15.6.windows-amd64.msi), [OSX](https://golang.org/dl/go1.15.6.darwin-amd64.pkg), [FreeBSD](https://golang.org/dl/go1.15.6.freebsd-amd64.tar.gz), [etc](https://golang.org/dl) ..

    $ make buildall

## Running the terminal

Once compiled it should be as easy as..

    $ crypto-market-ui

## Service API key

Like all cryptocurrency markets the [CoinMarketCap](https://coinmarketcap.com/api/documentation/v1) service requires a valid **API key** in order to query endpoint data.  You can generate an API Key by creating a [FREE account](https://pro.coinmarketcap.com/signup).  While this type of account has a rate limit ([10K call requests per /mo](https://coinmarketcap.com/api/documentation/v1/#section/Errors-and-Rate-Limits)) it should be enough to allow you to run this program consistently as long as you define a conservative **Refresh Rate** value.

## Examples

### Configuration Setup

![Setup](https://raw.githubusercontent.com/nuxy/go-crypto-market-ui/master/preview/setup.png)

### Help Menu

![Help](https://raw.githubusercontent.com/nuxy/go-crypto-market-ui/master/preview/help.png)

## Known Limitations

As of now you must configure cryptocurrencies manually so consider this package a pre-release.  This is due to complexities in how form values are currently being handled which shouldn't be rushed.  Note, this feature is in active development which I'm hoping to have ready by the next release (target 12/31), so be patient.

That being said, adding new cryptocurrency symbols that includes _units_ and _price_ is fairly trivial.  Upon program initialization a configuration file will be created in `$HOME/.crypto-market-ui.yaml` that follows the YAML structure below:

```yaml
- service: CoinMarketCap
  apiKey: 123e4567-e89b-12d3-a456-426614174000
  currency: USD
  language: en
  refreshRate: 60
  symbols:
    ALGO:
    - units: 10000
      price: 0.1234
    BAL:
    - units: 50
      price: 9.99
    BAT:
    - units: 420
      price: 0.7890

...
```

Edit this file by adding the symbols you want to track and restart the program.

## Contributions

If you fix a bug, or have a code you want to contribute, please send a pull-request with your changes. (Note: Before committing your code please ensure that you run [golint](https://github.com/golang/lint)).  Also, please adhere to the code style set forth in this package (`gofmt` is NOT recommended since _I prefer variable alignment_).

## Versioning

This package is maintained under the [Semantic Versioning](https://semver.org) guidelines.

## License and Warranty

This package is distributed in the hope that it will be useful, but without any warranty; without even the implied warranty of merchantability or fitness for a particular purpose.

_go-crypto-market-ui_ is provided under the terms of the [MIT license](http://www.opensource.org/licenses/mit-license.php)

## Author

[Marc S. Brooks](https://github.com/nuxy)
