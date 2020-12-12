class GoCryptoMarketUi < Formula
  desc     "Command-line utility to track cryptocurrencies in realtime."
  homepage "https://github.com/nuxy/go-crypto-market-ui"
  head     "https://github.com/nuxy/go-crypto-market-ui.git"
  url      "https://github.com/nuxy/go-crypto-market-ui/archive/0.0.2.tar.gz"
  sha256   "642808f1b01085df0d8123f8c9411c0ef56136a11a108f69f17adf26e89f5f1e"
  license  "MIT"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["PATH"]   = "#{ENV["PATH"]}:#{buildpath}/bin"

    (buildpath/"src/github.com/nuxy/go-crypto-market-ui").install buildpath.children

    cd "src/github.com/nuxy/go-crypto-market-ui" do
      system "make", "install"

      bin.install "#{buildpath}/bin/crypto-market-ui"
    end
  end
end
