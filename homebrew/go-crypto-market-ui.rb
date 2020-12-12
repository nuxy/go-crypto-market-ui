class GoCryptoMarketUi < Formula
  desc     "Command-line utility to track cryptocurrencies in realtime."
  homepage "https://github.com/nuxy/go-crypto-market-ui"
  head     "https://github.com/nuxy/go-crypto-market-ui.git"
  url      "https://github.com/nuxy/go-crypto-market-ui/archive/0.0.2.tar.gz"
  sha256   "23ea91d95f386a711c423525fb91b4f80c6f739753bed4672d9f9b9107a47f45"
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
