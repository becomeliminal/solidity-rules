subinclude("///foundry//build_defs:foundry")
subinclude("//build_defs:solidity")

# Foundry v1.5.0. hashes are the sha256 of foundry_v1.5.0_<os>_<arch>.tar.gz
# for linux_amd64, linux_arm64, darwin_amd64, and darwin_arm64. The foundry
# plugin accepts the download when it matches any one of them.
foundry(
    name = "foundry",
    version = "1.5.0",
    hashes = [
        "sha256:5cd98f9092bcc28be087939491f786b2bf3ed55e492996a409e29519b8ab4dc8",
        "sha256:8138e1615568bfcca5999773830892d93a569370eb0ae4b7dd97db46e2af47f9",
        "sha256:493b04c731b02b85d95f3bf9aff9cdddd0acd52c277902f3313fbe55e9b154e6",
        "sha256:17748f7b1815a66d91cce88946989b049358a874262cb848284609dd9094df06",
    ],
    visibility = ["PUBLIC"],
)

# One solc per version the tests compile with. SolcPackage is this package.
solc_versions(
    versions = [
        "0.7.6",
        "0.8.20",
        "0.8.23",
    ],
)
