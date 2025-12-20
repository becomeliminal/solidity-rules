subinclude("///go//build_defs:go")

package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "gotool",
    version = "1.25.2",
)

# CGO packages using go_module for RBE compatibility
go_module(
    name = "hid",
    install = ["."],
    module = "github.com/karalabe/hid",
    version = "v1.0.1-0.20240306101548-573246063e52",
)

go_module(
    name = "usb",
    install = ["."],
    module = "github.com/karalabe/usb",
    version = "v0.0.2",
)

#### CLI tools ####

go_repo(
    name = "go-humanize",
    install = ["."],
    module = "github.com/dustin/go-humanize",
    version = "v1.0.0",
)

go_repo(
    name = "go-flags",
    install = ["."],
    module = "github.com/thought-machine/go-flags",
    version = "v1.6.0",
)

go_repo(
    name = "go-logging",
    install = ["."],
    module = "gopkg.in/op/go-logging.v1",
    version = "v1.0.0-20160211212156-b2cb9fa56473",
)

go_repo(
    name = "go-cli-init",
    install = ["flags"],
    module = "github.com/peterebden/go-cli-init/v5",
    requirements = [
        "github.com/dustin/go-humanize",
        "github.com/thought-machine/go-flags",
        "gopkg.in/op/go-logging.v1",
        "golang.org/x/term",
        "golang.org/x/sys",
    ],
    version = "v5.2.1",
)

#### golang.org/x/xxx ####

go_repo(
    name = "golang.org_x_crypto",
    install = ["..."],
    module = "golang.org/x/crypto",
    version = "v0.32.0",
)

filegroup(
    name = "x_crypto",
    exported_deps = ["///third_party/go/golang.org_x_crypto//:installs"],
)

go_repo(
    name = "x_errors",
    install = ["..."],
    module = "golang.org/x/xerrors",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_repo(
    name = "x_exp",
    install = ["constraints", "rand", "slog", "slog/...", "slices", "maps"],
    module = "golang.org/x/exp",
    version = "v0.0.0-20240103183307-be819d1f06fc",
)

go_repo(
    name = "x_mod",
    install = ["..."],
    module = "golang.org/x/mod",
    version = "v0.5.1",
)

go_repo(
    name = "x_net",
    install = ["..."],
    module = "golang.org/x/net",
    version = "v0.0.0-20210614182718-04defd469f4e",
)

go_repo(
    name = "x_sync",
    install = ["..."],
    module = "golang.org/x/sync",
    version = "v0.6.0",
)

go_repo(
    name = "x_sys",
    install = ["..."],
    module = "golang.org/x/sys",
    version = "v0.7.0",
)

go_repo(
    name = "x_term",
    install = ["."],
    module = "golang.org/x/term",
    version = "v0.0.0-20210615171337-6886f2dfbf5b",
)

go_repo(
    name = "x_text",
    install = ["..."],
    module = "golang.org/x/text",
    version = "v0.16.0",
)

go_repo(
    name = "x_time",
    install = ["rate"],
    module = "golang.org/x/time",
    version = "v0.5.0",
)

go_repo(
    name = "x_tools",
    install = ["..."],
    module = "golang.org/x/tools",
    version = "v0.1.4",
)

go_repo(
    name = "x_perf",
    install = ["benchfmt"],
    module = "golang.org/x/perf",
    version = "v0.0.0-20240108191414-4ad5199aa6b5",
)

go_repo(
    name = "goldmark",
    install = ["..."],
    module = "github.com/yuin/goldmark",
    version = "v1.3.9",
)

#### Common deps ####

go_repo(
    name = "testify",
    install = ["..."],
    module = "github.com/stretchr/testify",
    requirements = [
        "github.com/stretchr/objx",
        "github.com/pmezard/go-difflib",
        "github.com/davecgh/go-spew",
        "gopkg.in/yaml.v3",
    ],
    version = "v1.8.4",
)

go_repo(
    name = "objx",
    install = ["."],
    module = "github.com/stretchr/objx",
    version = "v0.5.0",
)

go_repo(
    name = "ginkgo",
    install = ["..."],
    module = "github.com/onsi/ginkgo",
    version = "v1.16.5",
)

go_repo(
    name = "BurntSushi-toml",
    install = ["."],
    module = "github.com/BurntSushi/toml",
    version = "v1.2.1",
)

go_repo(
    name = "difflib",
    install = ["difflib"],
    module = "github.com/pmezard/go-difflib",
    version = "v1.0.0",
)

go_repo(
    name = "spew",
    install = ["spew"],
    module = "github.com/davecgh/go-spew",
    version = "v1.1.1",
)

go_repo(
    name = "yaml.v2",
    install = ["."],
    module = "gopkg.in/yaml.v2",
    version = "v2.2.2",
)

go_repo(
    name = "yaml.v3",
    install = ["."],
    module = "gopkg.in/yaml.v3",
    version = "v3.0.1",
)

go_repo(
    name = "errors",
    install = ["."],
    module = "github.com/pkg/errors",
    version = "v0.9.1",
)

go_repo(
    name = "snappy",
    install = ["."],
    module = "github.com/golang/snappy",
    version = "v0.0.5-0.20220116011046-fa5810519dcb",
)

go_repo(
    name = "xxhash",
    install = ["..."],
    module = "github.com/cespare/xxhash/v2",
    version = "v2.3.0",
)

go_repo(
    name = "mapstructure",
    install = ["."],
    module = "github.com/mitchellh/mapstructure",
    version = "v1.4.1",
)

go_repo(
    name = "color",
    install = ["."],
    module = "github.com/fatih/color",
    version = "v1.13.0",
)

go_repo(
    name = "go-colorable",
    install = ["."],
    module = "github.com/mattn/go-colorable",
    version = "v0.1.8",
)

go_repo(
    name = "go-isatty",
    install = ["."],
    module = "github.com/mattn/go-isatty",
    version = "v0.0.13",
)

go_repo(
    name = "google-uuid",
    install = ["."],
    module = "github.com/google/uuid",
    version = "v1.6.0",
)

go_repo(
    name = "moremath",
    install = ["mathx", "stats"],
    module = "github.com/aclements/go-moremath",
    version = "f10218a38794e5bdebc62889bae58d4d8a458de8",
)

go_repo(
    name = "cmp",
    install = ["cmp", "cmp/..."],
    module = "github.com/google/go-cmp",
    version = "v0.7.0",
)

go_repo(
    name = "protobuf",
    install = ["..."],
    module = "github.com/golang/protobuf",
    version = "v1.5.4",
)

go_repo(
    name = "google.golang.org.protobuf",
    install = ["..."],
    module = "google.golang.org/protobuf",
    version = "v1.34.1",
)

#### Ethereum deps ####

go_repo(
    name = "automaxprocs",
    install = ["..."],
    module = "go.uber.org/automaxprocs",
    version = "v1.5.3",
)

go_repo(
    name = "go-stub",
    install = ["..."],
    module = "github.com/prashantv/gostub",
    version = "v1.1.0",
)

go_repo(
    name = "bls12-381-util",
    install = ["..."],
    module = "github.com/protolambda/bls12-381-util",
    version = "d8552aa452c7a9de2f0b536f27683c91e06c6ad4",
)

go_repo(
    name = "bls12-381",
    install = ["..."],
    module = "github.com/kilic/bls12-381",
    version = "v0.1.0",
)

go_repo(
    name = "ztyp",
    install = ["..."],
    module = "github.com/protolambda/ztyp",
    version = "v0.2.2",
)

go_repo(
    name = "zrnt",
    install = ["..."],
    module = "github.com/protolambda/zrnt",
    version = "v0.32.2",
)

go_repo(
    name = "messagediff",
    install = ["..."],
    module = "github.com/protolambda/messagediff",
    version = "v1.4.0",
)

go_repo(
    name = "sha256-simd",
    install = ["..."],
    module = "github.com/minio/sha256-simd",
    version = "v1.0.0",
)

go_repo(
    name = "cpuid",
    install = ["..."],
    module = "github.com/klauspost/cpuid/v2",
    version = "v2.2.3",
)

go_repo(
    name = "billy",
    install = ["."],
    module = "github.com/holiman/billy",
    version = "v0.0.0-20240216141850-2abb0c79d3c4",
)

go_repo(
    name = "cli.v1",
    install = ["."],
    module = "gopkg.in/urfave/cli.v1",
    version = "v1.20.0",
)

go_repo(
    name = "cli.v2",
    install = ["..."],
    module = "github.com/urfave/cli/v2",
    requirements = ["github.com/BurntSushi/toml"],
    version = "v2.24.1",
)

go_repo(
    name = "smetrics",
    install = ["."],
    module = "github.com/xrash/smetrics",
    version = "1d8dd44e695e5d58393aed584b0ae4b5cb272bb9",
)

go_repo(
    name = "md2man",
    install = ["md2man"],
    module = "github.com/cpuguy83/go-md2man/v2",
    version = "v2.0.3",
)

go_repo(
    name = "blackfriday",
    install = ["."],
    module = "github.com/russross/blackfriday/v2",
    version = "v2.1.0",
)

go_repo(
    name = "bloomfilter",
    install = ["..."],
    module = "github.com/holiman/bloomfilter/v2",
    version = "v2.0.3",
)

go_repo(
    name = "hamming",
    install = ["."],
    module = "github.com/steakknife/hamming",
    version = "0.2.5",
)

go_repo(
    name = "cors",
    install = ["."],
    module = "github.com/rs/cors",
    version = "v1.7.0",
)

go_repo(
    name = "eventsource",
    install = ["."],
    module = "github.com/donovanhide/eventsource",
    version = "v0.0.0-20210830082556-c59027999da0",
)

go_repo(
    name = "fastssz",
    install = ["."],
    module = "github.com/ferranbt/fastssz",
    version = "v0.1.2",
)

go_repo(
    name = "flock",
    install = ["."],
    module = "github.com/gofrs/flock",
    version = "v0.8.1",
)

go_repo(
    name = "fastcache",
    install = ["."],
    module = "github.com/VictoriaMetrics/fastcache",
    version = "v1.6.0",
)

go_repo(
    name = "fsnotify",
    install = ["."],
    module = "github.com/fsnotify/fsnotify",
    version = "v1.6.0",
)

go_repo(
    name = "goja",
    install = ["."],
    module = "github.com/dop251/goja",
    version = "v0.0.0-20210614154742-14a1ffa82844",
)

go_repo(
    name = "regexp2",
    install = ["..."],
    module = "github.com/dlclark/regexp2",
    version = "v1.10.0",
)

go_repo(
    name = "go-sourcemap",
    install = ["."],
    module = "github.com/go-sourcemap/sourcemap",
    version = "v2.1.3",
)

go_repo(
    name = "golang-lru",
    install = ["..."],
    module = "github.com/hashicorp/golang-lru",
    version = "v0.5.4",
)

go_repo(
    name = "go-nat-pmp",
    install = ["."],
    module = "github.com/jackpal/go-nat-pmp",
    version = "v1.0.2",
)

go_repo(
    name = "goupnp",
    install = ["..."],
    module = "github.com/huin/goupnp",
    version = "v1.0.0",
)

go_repo(
    name = "goutil",
    install = ["codegen"],
    module = "github.com/huin/goutil",
    version = "1ca381bf315033e89af3286fdec0109ce8d86126",
)

go_repo(
    name = "golang-set",
    install = ["."],
    module = "github.com/deckarep/golang-set",
    version = "v1.7.1",
)

go_repo(
    name = "golang-set-v2",
    install = ["."],
    module = "github.com/deckarep/golang-set/v2",
    version = "v2.6.0",
)

go_repo(
    name = "go-kzg-4844",
    install = ["..."],
    module = "github.com/crate-crypto/go-kzg-4844",
    version = "v1.1.0",
)

go_repo(
    name = "gnark-crypto",
    install = ["ecc/...", "field/hash", "field/generator/config", "field/generator/internal/addchain", "field/pool", "internal/parallel", "internal/generator/config"],
    module = "github.com/consensys/gnark-crypto",
    version = "v0.10.0",
)

go_repo(
    name = "bavard",
    install = ["."],
    module = "github.com/consensys/bavard",
    version = "v0.1.13",
)

go_repo(
    name = "tmplfunc",
    install = ["..."],
    module = "rsc.io/tmplfunc",
    version = "v0.0.3",
)

go_repo(
    name = "bitset",
    install = ["."],
    module = "github.com/bits-and-blooms/bitset",
    version = "v1.13.0",
)

go_repo(
    name = "addchain",
    install = ["..."],
    module = "github.com/mmcloughlin/addchain",
    version = "v0.4.0",
)

go_repo(
    name = "profile",
    install = ["."],
    module = "github.com/mmcloughlin/profile",
    version = "v0.1.1",
)

go_repo(
    name = "subcommands",
    install = ["."],
    module = "github.com/google/subcommands",
    version = "v1.2.0",
)

go_repo(
    name = "go-bexpr",
    install = ["..."],
    module = "github.com/hashicorp/go-bexpr",
    version = "v0.1.13",
)

go_repo(
    name = "pointerstructure",
    install = ["."],
    module = "github.com/mitchellh/pointerstructure",
    version = "v1.2.1",
)

go_repo(
    name = "go-jwt-v4",
    install = ["."],
    module = "github.com/golang-jwt/jwt/v4",
    version = "v4.5.2",
)

go_repo(
    name = "go-bip39",
    install = ["..."],
    module = "github.com/tyler-smith/go-bip39",
    version = "v1.1.0",
)

go_module(
    name = "goleveldb",
    install = ["..."],
    module = "github.com/syndtr/goleveldb",
    strip = ["leveldb/testutil"],
    version = "64b5b1c73954",
    deps = [":snappy"],
)

go_repo(
    name = "gofuzz",
    install = ["."],
    module = "github.com/google/gofuzz",
    version = "v1.2.0",
)

go_repo(
    name = "gorilla_websocket",
    install = ["."],
    module = "github.com/gorilla/websocket",
    version = "v1.4.2",
)

go_repo(
    name = "httprouter",
    install = ["."],
    module = "github.com/julienschmidt/httprouter",
    version = "v1.3.0",
)

go_repo(
    name = "lumberjack.v2",
    install = ["."],
    module = "gopkg.in/natefinch/lumberjack.v2",
    version = "v2.2.1",
)

go_repo(
    name = "liner",
    install = ["."],
    module = "github.com/peterh/liner",
    version = "v1.2.1",
)

go_repo(
    name = "go-runewidth",
    install = ["."],
    module = "github.com/mattn/go-runewidth",
    version = "v0.0.13",
)

go_repo(
    name = "uniseg",
    install = ["."],
    module = "github.com/rivo/uniseg",
    version = "v0.2.0",
)

go_repo(
    name = "pebble",
    install = ["..."],
    module = "github.com/cockroachdb/pebble",
    version = "v1.1.2",
)

go_repo(
    name = "fifo",
    install = ["."],
    module = "github.com/cockroachdb/fifo",
    version = "v0.0.0-20240606204812-0bbfbd93a7ce",
)

go_repo(
    name = "asciigraph",
    install = ["."],
    module = "github.com/guptarohit/asciigraph",
    version = "v0.5.6",
)

go_repo(
    name = "hdrhistogram-go",
    install = ["."],
    module = "github.com/codahale/hdrhistogram",
    version = "v0.0.0-20161010025455-3a0bb77429bd",
)

go_repo(
    name = "hdrhistogram",
    install = ["."],
    module = "github.com/HdrHistogram/hdrhistogram-go",
    version = "v1.1.2",
)

go_repo(
    name = "zstd",
    install = ["."],
    module = "github.com/DataDog/zstd",
    version = "v1.5.0",
)

go_repo(
    name = "cockroachdb-errors",
    install = ["..."],
    module = "github.com/cockroachdb/errors",
    version = "v1.8.9",
)

go_repo(
    name = "gogo-protobuf",
    install = ["..."],
    module = "github.com/gogo/protobuf",
    version = "v1.3.2",
)

go_repo(
    name = "cockroachdb-redact",
    install = ["..."],
    module = "github.com/cockroachdb/redact",
    version = "v1.1.3",
)

go_repo(
    name = "cockroachdb-logtags",
    install = ["."],
    module = "github.com/cockroachdb/logtags",
    version = "dabe8e521a4ffbce9a63bffceb5bcbc5159effef",
)

go_repo(
    name = "cockroachdb-tokenbucket",
    install = ["."],
    module = "github.com/cockroachdb/tokenbucket",
    version = "cc333fc44b06d2f30c5a42d29645699e6ae89859",
)

go_repo(
    name = "cockroachdb-datadriven",
    install = ["."],
    module = "github.com/cockroachdb/datadriven",
    version = "v1.0.2",
)

go_repo(
    name = "pretty",
    install = ["."],
    module = "github.com/kr/pretty",
    version = "v0.3.0",
)

go_repo(
    name = "text",
    install = ["."],
    module = "github.com/kr/text",
    version = "v0.2.0",
)

go_repo(
    name = "fmtsort",
    install = ["fmtsort"],
    module = "github.com/rogpeppe/go-internal",
    version = "v1.8.1",
)

go_repo(
    name = "sentry-go",
    install = [".", "internal/debug", "internal/crypto/randutil", "internal/ratelimit"],
    module = "github.com/getsentry/sentry-go",
    version = "v0.12.0",
)

go_repo(
    name = "cobra",
    install = ["."],
    module = "github.com/spf13/cobra",
    version = "v1.3.0",
)

go_repo(
    name = "pflag",
    install = ["."],
    module = "github.com/spf13/pflag",
    version = "v1.0.5",
)

go_repo(
    name = "json-iterator",
    install = ["..."],
    module = "github.com/json-iterator/go",
    version = "v1.1.12",
)

go_repo(
    name = "concurrent",
    install = ["..."],
    module = "github.com/modern-go/concurrent",
    version = "1.0.3",
)

go_repo(
    name = "reflect2",
    install = ["..."],
    module = "github.com/modern-go/reflect2",
    requirements = ["github.com/modern-go/concurrent"],
    version = "v1.0.1",
)

go_repo(
    name = "prometheus-client",
    install = ["..."],
    module = "github.com/prometheus/client_golang",
    version = "v1.18.0",
)

go_repo(
    name = "client_model",
    install = ["..."],
    module = "github.com/prometheus/client_model",
    version = "v0.5.0",
)

go_repo(
    name = "perks",
    install = ["..."],
    module = "github.com/beorn7/perks",
    version = "v1.0.1",
)

go_repo(
    name = "procfs",
    install = ["..."],
    module = "github.com/prometheus/procfs",
    version = "v0.0.3",
)

go_repo(
    name = "prometheus-common",
    install = ["expfmt", "model"],
    module = "github.com/prometheus/common",
    requirements = ["github.com/matttproud/golang_protobuf_extensions/v2"],
    version = "v0.45.0",
)

go_repo(
    name = "golang_protobuf_extensions_v2",
    install = ["..."],
    module = "github.com/matttproud/golang_protobuf_extensions/v2",
    version = "v2.0.0",
)

go_repo(
    name = "tablewriter",
    install = ["."],
    module = "github.com/olekukonko/tablewriter",
    version = "v0.0.5",
)

go_repo(
    name = "toml",
    install = ["."],
    module = "github.com/naoina/toml",
    requirements = ["github.com/naoina/go-stringutil"],
    version = "v0.1.1",
)

go_repo(
    name = "go-stringutil",
    install = ["."],
    module = "github.com/naoina/go-stringutil",
    version = "v0.1.0",
)

go_repo(
    name = "go-toml-v2",
    install = ["."],
    module = "github.com/pelletier/go-toml/v2",
    version = "v2.2.3",
)

go_repo(
    name = "uint256",
    install = ["."],
    module = "github.com/holiman/uint256",
    version = "v1.3.2",
)

go_repo(
    name = "verkle",
    install = ["..."],
    module = "github.com/ethereum/go-verkle",
    version = "v0.2.2",
)

go_repo(
    name = "gballet-go-verkle",
    install = ["..."],
    module = "github.com/gballet/go-verkle",
    version = "v0.1.1-0.20231031103413-a67434b50f46",
)

go_repo(
    name = "go-ipa",
    install = ["..."],
    module = "github.com/crate-crypto/go-ipa",
    version = "v0.0.0-20240223125850-b1e8a79f509c",
)

go_repo(
    name = "btcec",
    install = ["..."],
    module = "github.com/btcsuite/btcd/btcec/v2",
    requirements = [
        "github.com/decred/dcrd/dcrec/secp256k1/v4",
        "github.com/btcsuite/btcd/chaincfg/chainhash",
    ],
    version = "v2.2.0",
)

go_repo(
    name = "btcd-chaincfg-chainhash",
    install = ["."],
    module = "github.com/btcsuite/btcd/chaincfg/chainhash",
    version = "v1.0.1",
)

go_repo(
    name = "decred-secp256k1",
    install = ["..."],
    module = "github.com/decred/dcrd/dcrec/secp256k1/v4",
    requirements = ["github.com/decred/dcrd/crypto/blake256"],
    version = "v4.0.1",
)

go_repo(
    name = "decred-blake256",
    install = ["."],
    module = "github.com/decred/dcrd/crypto/blake256",
    version = "v1.0.0",
)

go_repo(
    name = "btcd-chaincfg",
    install = ["chaincfg/chainhash"],
    module = "github.com/btcsuite/btcd",
    version = "v0.22.0-beta",
)

go_repo(
    name = "graphql-go",
    install = ["..."],
    module = "github.com/graph-gophers/graphql-go",
    version = "v1.3.0",
)

go_repo(
    name = "go-minisign",
    install = ["."],
    module = "github.com/jedisct1/go-minisign",
    version = "v0.0.0-20230811132847-661be99b8267",
)

go_repo(
    name = "go-libpcsclite",
    install = ["."],
    module = "github.com/gballet/go-libpcsclite",
    version = "v0.0.0-20190607065134-2772fd86a8ff",
)

go_repo(
    name = "derivationpath",
    install = ["derivationpath"],
    module = "github.com/status-im/keycard-go",
    version = "v0.2.0",
)

go_repo(
    name = "influxdb1-client",
    install = ["..."],
    module = "github.com/influxdata/influxdb1-client",
    version = "v0.0.0-20220302092344-a9ab5670611c",
)

go_repo(
    name = "influxdb-client-go",
    install = [
        ".",
        "api",
        "api/...",
        "domain",
        "log",
    ],
    module = "github.com/influxdata/influxdb-client-go/v2",
    version = "v2.4.0",
)

go_repo(
    name = "influxdb-line-protocol",
    install = ["..."],
    module = "github.com/influxdata/line-protocol",
    version = "v0.0.0-20200327222509-2487e7298839",
)

go_repo(
    name = "deepequal",
    install = ["."],
    module = "github.com/influxdata/influxdb-client-go/v2/internal/gzip",
    version = "v2.4.0",
)

go_repo(
    name = "kylelemons-godebug",
    install = ["..."],
    module = "github.com/kylelemons/godebug",
    version = "v1.1.0",
)

go_repo(
    name = "opentracing",
    install = ["..."],
    module = "github.com/opentracing/opentracing-go",
    version = "v1.2.0",
)

go_repo(
    name = "gopter",
    install = ["..."],
    module = "github.com/leanovate/gopter",
    version = "v0.2.9",
)

# gopsutil - needed by go-ethereum/metrics
go_repo(
    name = "gopsutil",
    install = ["..."],
    module = "github.com/shirou/gopsutil",
    requirements = [
        "github.com/tklauser/go-sysconf",
        "golang.org/x/sys",
    ],
    version = "v3.21.5+incompatible",
)

go_repo(
    name = "go-sysconf",
    install = ["..."],
    module = "github.com/tklauser/go-sysconf",
    requirements = ["github.com/tklauser/numcpus"],
    version = "v0.3.13",
)

go_repo(
    name = "numcpus",
    install = ["."],
    module = "github.com/tklauser/numcpus",
    version = "v0.7.0",
)

go_repo(
    name = "azure-sdk-azblob",
    install = [".", "container"],
    module = "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob",
    version = "v1.2.0",
)

go_repo(
    name = "azure-sdk-azcore",
    install = ["..."],
    module = "github.com/Azure/azure-sdk-for-go/sdk/azcore",
    version = "v1.7.0",
)

go_repo(
    name = "azure-sdk-internal",
    install = ["..."],
    module = "github.com/Azure/azure-sdk-for-go/sdk/internal",
    version = "v1.3.0",
)

go_repo(
    name = "azure-storege-blob-go",
    install = ["azblob"],
    module = "github.com/Azure/azure-storage-blob-go",
    version = "v0.7.0",
)

go_repo(
    name = "azure-pipeline-go",
    install = ["pipeline"],
    module = "github.com/Azure/azure-pipeline-go",
    version = "v0.2.3",
)

go_repo(
    name = "cloudflare-go",
    install = ["."],
    module = "github.com/cloudflare/cloudflare-go",
    version = "v0.17.0",
)

# CGO package - needs go_module
go_module(
    name = "go-duktape.v3",
    install = ["."],
    module = "gopkg.in/olebedev/go-duktape.v3",
    version = "v3.0.0-20210326210528-650f7c854440",
)

go_repo(
    name = "docker-reexec",
    install = ["pkg/reexec"],
    module = "github.com/docker/docker",
    requirements = ["golang.org/x/sys"],
    version = "v20.10.7",
)

go_repo(
    name = "fileutil",
    install = ["fileutil"],
    module = "github.com/prometheus/tsdb",
    version = "v0.10.0",
)

go_repo(
    name = "github.com_shirou_gopsutil_v3",
    install = ["..."],
    module = "github.com/shirou/gopsutil/v3",
    version = "v3.24.5",
)

go_repo(
    name = "influxdb-client",
    install = ["client", "query", "models", "pkg/escape"],
    module = "github.com/influxdata/influxdb",
    version = "v1.8.3",
)

go_repo(
    name = "mmap-go",
    install = ["."],
    module = "github.com/edsrzf/mmap-go",
    version = "v1.2.0",
)

go_repo(
    name = "memsizeui",
    install = ["..."],
    module = "github.com/fjl/memsize",
    version = "v0.0.0-20190710130421-bcb5799ab5e5",
)

go_repo(
    name = "notify",
    install = ["."],
    module = "github.com/rjeczalik/notify",
    version = "v0.9.2",
)

go_repo(
    name = "stack",
    install = ["."],
    module = "github.com/go-stack/stack",
    version = "v1.8.0",
)

go_repo(
    name = "stun",
    install = ["..."],
    module = "github.com/pion/stun/v2",
    version = "v2.0.0",
)

go_repo(
    name = "pion-transport.v3",
    install = ["..."],
    module = "github.com/pion/transport/v3",
    version = "v3.0.1",
)

go_repo(
    name = "pion-transport",
    install = ["..."],
    module = "github.com/pion/transport/v2",
    version = "v2.2.1",
)

go_repo(
    name = "pion-logging",
    install = ["..."],
    module = "github.com/pion/logging",
    version = "v0.2.2",
)

go_repo(
    name = "oapi-codegen",
    install = [
        "pkg/runtime",
        "pkg/types",
    ],
    module = "github.com/deepmap/oapi-codegen",
    version = "v1.6.0",
)

go_repo(
    name = "dtls",
    install = ["..."],
    module = "github.com/pion/dtls/v2",
    version = "v2.2.7",
)

go_repo(
    name = "zap",
    install = ["..."],
    module = "go.uber.org/zap",
    version = "v1.17.0",
)

go_repo(
    name = "multierr",
    install = ["."],
    module = "go.uber.org/multierr",
    version = "v1.6.0",
)

go_repo(
    name = "atomic",
    install = ["."],
    module = "go.uber.org/atomic",
    version = "v1.7.0",
)

go_repo(
    name = "influxql",
    install = ["."],
    module = "github.com/influxdata/influxql",
    version = "b6ab99450c93",
)

go_repo(
    name = "treeprint",
    install = ["."],
    module = "github.com/xlab/treeprint",
    version = "v1.1.0",
)

go_repo(
    name = "go-ieproxy",
    install = ["."],
    module = "github.com/mattn/go-ieproxy",
    version = "v0.0.1",
)

GO_ETHEREUM_DEPS = [
    ":azure-storege-blob-go",
    ":automaxprocs",
    ":azure-sdk-azblob",
    ":bls12-381-util",
    ":billy",
    ":bloomfilter",
    ":cors",
    ":cli.v1",
    ":cli.v2",
    ":color",
    ":cloudflare-go",
    ":derivationpath",
    ":docker-reexec",
    ":eventsource",
    ":fastssz",
    ":flock",
    ":fileutil",
    ":fastcache",
    ":fsnotify",
    ":google.golang.org.protobuf",
    ":goja",
    ":golang-lru",
    ":gopsutil",
    ":github.com_shirou_gopsutil_v3",
    ":go-minisign",
    ":go-duktape.v3",
    ":graphql-go",
    ":go-nat-pmp",
    ":goupnp",
    ":go-kzg-4844",
    ":go-bexpr",
    ":go-jwt-v4",
    ":go-bip39",
    ":go-libpcsclite",
    ":golang-set",
    ":golang-set-v2",
    ":gofuzz",
    ":gorilla_websocket",
    ":goleveldb",
    ":httprouter",
    ":influxdb-client-go",
    ":influxdb1-client",
    ":influxdb-client",
    ":lumberjack.v2",
    ":liner",
    ":mmap-go",
    ":memsizeui",
    ":notify",
    ":pebble",
    ":stack",
    ":stun",
    ":snappy",
    ":tablewriter",
    ":toml",
    ":testify",
    ":uint256",
    ":verkle",
    ":x_sync",
    ":x_crypto",
    ":x_exp",
    ":ztyp",
    ":zrnt",
    ":google-uuid",
    ":hid",
    ":usb",
    ":gnark-crypto",
]

# go-ethereum uses CGO for crypto/secp256k1, so keep as go_module
go_module(
    name = "go-ethereum",
    install = ["..."],
    module = "github.com/ethereum/go-ethereum",
    strip = [
        "cmd",
        "rlp/rlpgen",
        "tests/fuzzers",
    ],
    version = "v1.15.3",
    deps = GO_ETHEREUM_DEPS,
)

# abigen binary from go-ethereum
go_module(
    name = "abigen",
    binary = True,
    install = ["cmd/abigen"],
    module = "github.com/ethereum/go-ethereum",
    strip = [
        "tests/fuzzers",
    ],
    version = "v1.15.3",
    deps = GO_ETHEREUM_DEPS + [":go-ethereum"],
)
