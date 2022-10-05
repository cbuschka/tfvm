package state

const defaultStateJSON = `
{
 "lastUpdateTime": "2022-10-05T07:46:45.75092885+02:00",
 "terraformReleases": [
  {
   "version": "1.3.0-alpha20220817",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_darwin_amd64.zip",
     "sha256_checksum": "254ceea0ceda640289038b8d2e79e8609b72faee990859ffbc1700c385dd1e56"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_darwin_arm64.zip",
     "sha256_checksum": "b6186f4b1a3a057af66f793c9623282c227b2dd4ca890a7a5007ae0c71fd108a"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_freebsd_386.zip",
     "sha256_checksum": "6f235ac452e3dc3af82412860f9f7689be9602b9250f8a58a57cf989b081555"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_freebsd_amd64.zip",
     "sha256_checksum": "4760b079d3c5f0ff8b8caf94c208e20d3406ab0b7bc58a13362a8f71194eac87"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_freebsd_arm.zip",
     "sha256_checksum": "7a2ba43b4d4adcd4ffdafcc585a2e1ae08062c923c59523bab2a7409adf31f0d"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_linux_386.zip",
     "sha256_checksum": "9ba2fa82e03af79e344e04e0e9357cd610a251585767b29fc7a2172a9336538d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_linux_amd64.zip",
     "sha256_checksum": "f75df826d052daf47adee872adea1e41cf2fa38288eaec47211cc461f277884e"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_linux_arm.zip",
     "sha256_checksum": "a8e5208e212194b3d547d4af9ba6dfb0441b6f55a3ee069ae8c44d37ff2a8213"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_linux_arm64.zip",
     "sha256_checksum": "1d55ebb0b59d5578eb48e350f7eb9a4ff7414168bfe743e7dba173db2bf334eb"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_openbsd_386.zip",
     "sha256_checksum": "8ef67c997118303a7154ee11ff7083a6ffe43e7dfd95c59fb77ba1366f5cb74b"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_openbsd_amd64.zip",
     "sha256_checksum": "d447c2fc84fa47bf383149c91630f6a5cbfe32d2e6cc13abe1881eacac34201"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_solaris_amd64.zip",
     "sha256_checksum": "1ac6d255b27df16208aa7fbdb2aaea11b14dcfa5722b5b28f06d38e82f8f0332"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_windows_386.zip",
     "sha256_checksum": "6b8f6b8cf0905520cab6423aa0120f773d8c7dde076cb3b02ca8a6469d5b3735"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220817/terraform_1.3.0-alpha20220817_windows_amd64.zip",
     "sha256_checksum": "bb99cad0b364816f6335265333f8c63fe0bcd8d193dba8721b46d7dab900c56c"
    }
   ]
  },
  {
   "version": "0.8.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_darwin_amd64.zip",
     "sha256_checksum": "ba53c7424bec5db7c01e0a5178ba5e295eb13669fb04fdae41576098baf88b75"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_freebsd_386.zip",
     "sha256_checksum": "f2ef71d88f0801c9274fe95cee925dd3346d0bc6fceb54a861274ca836835c0b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_freebsd_amd64.zip",
     "sha256_checksum": "e1ca4c377ccfb8307b85f9979189a498820d0536baee264e23c4288fc61f334"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_freebsd_arm.zip",
     "sha256_checksum": "ef2464e60799f8a416da446a586a3784132000d294ef054b3699191a318854c3"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_linux_386.zip",
     "sha256_checksum": "33862079d184fa384d9a2a442fa97ec53cbf94a17f92890ba97cc750ed676f13"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_linux_amd64.zip",
     "sha256_checksum": "7ca424d8d0e06697cc7f492b162223aef525bfbcd69248134a0ce0b529285c8c"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_linux_arm.zip",
     "sha256_checksum": "7258b72a89648b1617346e4be38bbe400166967a54afb02ceb3d4bd1d7b97351"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_openbsd_386.zip",
     "sha256_checksum": "f2e4f3a67bbf9cd5b131b45d3f29fa20bbf73367cb0012307e8220eb55c7232f"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_openbsd_amd64.zip",
     "sha256_checksum": "fe282f9dc8dbf110baddf2a724402371b83cfa68e7bc4c024c830263c3e49717"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_solaris_amd64.zip",
     "sha256_checksum": "2cf67e4ae28f98a429042a6ed98c50ddebe6cacff381170e1d4296bd90e2874e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_windows_386.zip",
     "sha256_checksum": "4e7f89322434cfb91aa6c9448f12b189b18e3d4f631dda2acc11a2fcc0db1a95"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_windows_amd64.zip",
     "sha256_checksum": "dcd2cacfec4f9d6e9834d9c2cfef9b7b868e191001d88607f088dc93406054bd"
    }
   ]
  },
  {
   "version": "0.9.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_darwin_amd64.zip",
     "sha256_checksum": "33d9bbe1516a4085998c74d5a265aa0354d29a11eb56a21611dbcc806aec9c6f"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_freebsd_386.zip",
     "sha256_checksum": "682955e3117fe65e1efad0ff158f87b12ea09b81da95e7fdee5acab7fd90b00e"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_freebsd_amd64.zip",
     "sha256_checksum": "8664b4b11ea4dd8cc9c2c4985e2a09f73c65d3e453376ca415e8289a6045617e"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_freebsd_arm.zip",
     "sha256_checksum": "ec5c8a6c3c0a0c3ef5eec0db5315f259531546b7bc7fb4e59ebbe73477bb7154"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_linux_386.zip",
     "sha256_checksum": "4fe31a7d815cbf898c3ec708bdeb4c39891bd7eeedb7a604c5738961ecc73e1d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_linux_amd64.zip",
     "sha256_checksum": "a916228ea4c19c91c9a5dee2905885f517eb7c31ba4dbf5d79f9f36606973313"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_linux_arm.zip",
     "sha256_checksum": "268bc3f0cdb69e6ddf49d0ec8b3df2cf86d7cddf57276768726eb8eba8b6b506"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_openbsd_386.zip",
     "sha256_checksum": "37e9dc0ae263732a68ded77a31da92abfac678a92dfc24b1060e68b44c937b85"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_openbsd_amd64.zip",
     "sha256_checksum": "a0d2d02d991b4ef467bb02d862fe398153c27986b7ca128282aa9917f8d5a8f8"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_solaris_amd64.zip",
     "sha256_checksum": "ad51479defd40c158c9551a1d48107d6b031db548192d5e8c224ef37feda33bf"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_windows_386.zip",
     "sha256_checksum": "c42f6b4499ac2372c11710a001bd4935e1ee6bd83b7bfe2f728b849d52441610"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_windows_amd64.zip",
     "sha256_checksum": "17cccc1dec8a0d8a8684a1533b1ce200202bb7176f78fcb1ca9d82d4ead53f03"
    }
   ]
  },
  {
   "version": "0.6.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_darwin_386.zip",
     "sha256_checksum": "8f4d0325e992cf6ea9ca70b17f6710c1995a64884459ee115e48d721fc65f007"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_darwin_amd64.zip",
     "sha256_checksum": "ba540f36d1dc3ed9d3db9832db3a2b3f6cfea5d9f80b663281c1d28260d298ed"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_freebsd_386.zip",
     "sha256_checksum": "14dd5a370b7580ef29b7df31ca5efbae9707b190268419e92a646fdcdfbe9a61"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_freebsd_amd64.zip",
     "sha256_checksum": "a3de4aa3f2388a1b467b3ad32e3f0a6466b469bf18b801134b69766ed28156f"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_freebsd_arm.zip",
     "sha256_checksum": "42b303349595f395799dbe7bc8f397d79adbb619a1e950c277ec35b57ec5b101"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_linux_386.zip",
     "sha256_checksum": "7c5a6e4f71cfd47a5840ca318c6cae36ce9b4ed3821831adfd7a9da7630fd26d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_linux_amd64.zip",
     "sha256_checksum": "29618f11e9af410e4ed69d79084c33ad5faa3f475651556e039bb27ff69b3c74"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_linux_arm.zip",
     "sha256_checksum": "967f62e5348e330029f4bceef228e1f1018ca5db1b422b6cc18d901f4b7afc50"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_openbsd_386.zip",
     "sha256_checksum": "5b25694a3b9510928959ab5bc75f64cdddd63620c9884c835572b6e6056b8bc6"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_openbsd_amd64.zip",
     "sha256_checksum": "8be6da35a286e014c9a2bbea42738338aa2f35e8a97eb274de0b5d56e9ec4a17"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_windows_386.zip",
     "sha256_checksum": "b7b4f7414f1ac04e470ec3b0868e7052ed562377df54083a639b7d762d23a53f"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_windows_amd64.zip",
     "sha256_checksum": "43d1d16afc5848cd7de7b90ae35e92b8a5a845a7133e270c62037ef1d8dab195"
    }
   ]
  },
  {
   "version": "0.15.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_darwin_amd64.zip",
     "sha256_checksum": "bc4a4665af56c8e8bcf62788224f8fb91eeb7fe3b064ebcf3f3ab7bc5a90ea43"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_freebsd_386.zip",
     "sha256_checksum": "3aeb52bed1fda2685b84d6141b17bd7999e09182f66bbeb6910f54c8a512f429"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_freebsd_amd64.zip",
     "sha256_checksum": "a71e6582ad86d5d8265828a7560dc2a14425876227ab23bfe4dd8a5b8a4d03f0"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_freebsd_arm.zip",
     "sha256_checksum": "8386dd22e3e0709db77fbbe2ebdcdda9211f45c08c3a3624da36c7db5772b7d9"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_linux_386.zip",
     "sha256_checksum": "3e77bf4a2eac37fca0dc5eccaf180866b31ff8efa5525b64e370c73fd3c33e0a"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_linux_amd64.zip",
     "sha256_checksum": "b96d9bd0df800ce9e6d8b3b10c2aba491980a123c4c04a32049769247c9af76a"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_linux_arm.zip",
     "sha256_checksum": "2eddcd4bdebfeec332f81f730f2eb39b083ba3464b4e416ee1d42d9edf64d828"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_linux_arm64.zip",
     "sha256_checksum": "db2631b84556abbebb18d5afd6d4fd831cdcd00ae7d11b910030c4e3c0855478"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_openbsd_386.zip",
     "sha256_checksum": "7c17b904fda70336957a255d64840aa0abc6886cd04bd61159c48311c2423cc2"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_openbsd_amd64.zip",
     "sha256_checksum": "ce05b515e282d6944c2c8157046d4eed123d4a671ae07ae2c678b179a025189f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_solaris_amd64.zip",
     "sha256_checksum": "7488b4e42c211410abfa4ed72cef9f285a0a2dfabc4de08a0d89ab9f75371164"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_windows_386.zip",
     "sha256_checksum": "29cc7991b299f82529f04627440a6edd586a43a59bbe6270704ee658bcb234ff"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_windows_amd64.zip",
     "sha256_checksum": "753bcc13c00546f4aa434323bf11c9addd18c5c8c42cb5217aaf6fc4797f79bb"
    }
   ]
  },
  {
   "version": "0.14.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_darwin_amd64.zip",
     "sha256_checksum": "5c0110a4dc44ec01edd159c69bf60cebd18540eaf168aacd8b37d828b70e265d"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_freebsd_386.zip",
     "sha256_checksum": "43811b7d91efcac9410f367a97edd379ab85bedacf963b7591e4de8c60fd3b82"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_freebsd_amd64.zip",
     "sha256_checksum": "8a4eabb258bed10f7a8abd9f5cff4a028dabbc413e21c3b09149f4679f524c69"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_freebsd_arm.zip",
     "sha256_checksum": "2d793cf8fe4751e2c7dcaba70195ac7968a44f49787770f155db9fa922780e02"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_linux_386.zip",
     "sha256_checksum": "27aa3d1aacaa1aa64ed743b44995b47e3eebe2ded14b221f4dee9f98dd8350ff"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_linux_amd64.zip",
     "sha256_checksum": "171ef5a4691b6f86eab524feaf9a52d5221c875478bd63dd7e55fef3939f7fd4"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_linux_arm.zip",
     "sha256_checksum": "4f9f628e58998f50d8c8bc96b4f8feb3bf13fc7ae54c47c234526ee48f7514d9"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_linux_arm64.zip",
     "sha256_checksum": "602813ce2afe1e874807a66fc2e28b7cb90d9381b76a5d7b7f0ec9aa768eb2bc"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_openbsd_386.zip",
     "sha256_checksum": "a83c4f0282ef761962b93a5e42ab3a40720999be75a38421514088272f61e0"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_openbsd_amd64.zip",
     "sha256_checksum": "a29cea17067fb3b29f5b3a411e22901210b8a2a71d23452d5733d317ffdd0e6e"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_solaris_amd64.zip",
     "sha256_checksum": "33fdfa45b0be5191e5eae7bf1367b988a148a71d42bbe1ed6daef827bbf1ebc4"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_windows_386.zip",
     "sha256_checksum": "e91e106a0c5ba1c3d98fbf23a18305be0ebceda4112d7d682c370b582225c8d2"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_windows_amd64.zip",
     "sha256_checksum": "acec7890c57074ab9d74bbb19076bcd3b6a11266bb19f6bc6298ca6900c5a4c0"
    }
   ]
  },
  {
   "version": "1.0.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_darwin_amd64.zip",
     "sha256_checksum": "1de66203cf7f62ad990b6d8b583bc2caaadf8594150323f4d632a869448b85b9"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_darwin_arm64.zip",
     "sha256_checksum": "bf330b9d9bf24e87abf155de3828be2afa5a61d07df4d8cfe3d61e32bf71e687"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_freebsd_386.zip",
     "sha256_checksum": "c72979629efb56bac49e773f57050c1a6216358c069e4dd4f2ae0c08f77c458d"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_freebsd_amd64.zip",
     "sha256_checksum": "ff4781fdfdd7da24b9ba823ca661e69417967057346704fe0edadcbc0f2ce3a1"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_freebsd_arm.zip",
     "sha256_checksum": "959816a4820a4623caacedea169d4bf838349a3ce5bba5e33652267e92f2bb1d"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_linux_386.zip",
     "sha256_checksum": "b094927dd3993c6ddb1aca763b94465d4196fee3d6ea92ec40e8ba05e28ccda3"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_linux_amd64.zip",
     "sha256_checksum": "99c4866ffc4d3a749671b1f74d37f907eda1d67d7fc29ed5485aeff592980644"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_linux_arm.zip",
     "sha256_checksum": "10587a94d08b00bcf37792c125c9b33f7c672b27b97f65bd9acb02d78cd70fd9"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_linux_arm64.zip",
     "sha256_checksum": "521a52dc6996c105b059096f7f92c0a382aa0903053ff85755fb2c631d0d1190"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_openbsd_386.zip",
     "sha256_checksum": "4143d4d249c607e1cc2e8a7ecd1a156b1dd9a5fb34b28e54110c81d14a8ac3bd"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_openbsd_amd64.zip",
     "sha256_checksum": "858a886d18f4299e8d26e176a7ed36c223116470a690e32c0e02bef26a90de4b"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_solaris_amd64.zip",
     "sha256_checksum": "8e3f9b149f5e75ab71dde53d8bd4b0949587cc2e7698df44175bff7955a91b66"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_windows_386.zip",
     "sha256_checksum": "5d763847534bc41932681a18d01eade116c69c90eb22cf98da80f74109d883c9"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_windows_amd64.zip",
     "sha256_checksum": "ec87e0835177406d16e8b560089241e67805ef206b3e082e80045a2ab48dd68"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210630",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_darwin_amd64.zip",
     "sha256_checksum": "e1aad1616af609c2c87ccfe6d617c698a44abd1b0147afc0a1958a08bd4c2c9e"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_freebsd_386.zip",
     "sha256_checksum": "92766b3375a6b307da7a1b53e7d58b2d35d0279b25d24bd0a66ba17f6dac3e7a"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_freebsd_amd64.zip",
     "sha256_checksum": "bcf4d8c852abce38b61c20979cdfdafe41f443a3e7ea9bc2899268b6e32c9c59"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_freebsd_arm.zip",
     "sha256_checksum": "e125f329bc6c0d1a934a9e61df835217efd25e9dd3232cf401557fe50b6bde19"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_linux_386.zip",
     "sha256_checksum": "27b5f8d5f86ca0c12c965ba1c5abc2c203893dda265595dc9fc231bfdcf16cac"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_linux_amd64.zip",
     "sha256_checksum": "5c62781f675361f3dbaa80ad65829f1a01302f7c15ac007892ae246e0f6dead7"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_linux_arm.zip",
     "sha256_checksum": "1f6d99a5a8707661c1064ab114d44c4103c8816466a08a5bbdb7809999c5abc5"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_linux_arm64.zip",
     "sha256_checksum": "c2318730fb7f53d1890cd2e2f06baf8c2a49f445c7aaeb373c14868eb1656945"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_openbsd_386.zip",
     "sha256_checksum": "7b7f6effb98d676c0e7d341c084cb62c53169ae5a7ba7ea62b301ae1cb38895b"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_openbsd_amd64.zip",
     "sha256_checksum": "b10a5f6b96ded6b017b1c1ff6815e0ea55f5395f796f84b2a2592aca8fbbd88"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_solaris_amd64.zip",
     "sha256_checksum": "828e1b9981ce0d65bae484268cc701f562dd25e2a0658dcd4f18c8a626759824"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_windows_386.zip",
     "sha256_checksum": "edf69c43f26438124a3cc354d86b66995fc1037b90fb36954063862880e856df"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_windows_amd64.zip",
     "sha256_checksum": "25fb7a73ddcd3055fa6f95b1928f60f40729117e89b96ff9dbde957c498ece75"
    }
   ]
  },
  {
   "version": "0.14.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_darwin_amd64.zip",
     "sha256_checksum": "948d4550b7cd0f9152741c4a5e4fe80167b1cbb7513f939ffef1d50f94c4fb0c"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_freebsd_386.zip",
     "sha256_checksum": "4c1e04ed2ec76c5c6de794e9d5de0c22a457e84d1163296422d8ed575c8d1f3c"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_freebsd_amd64.zip",
     "sha256_checksum": "594fe33b0ef9f38526df8d8a1da2ef0dc21f4af161e58d36bf8f969eaac05933"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_freebsd_arm.zip",
     "sha256_checksum": "bf5596c7b520bb4f6d98a493b22c322150d9803784ea58ce16ee61f579f5f831"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_linux_386.zip",
     "sha256_checksum": "5de8033d94bf49ed22d0a9f3aa55b8d1658eb507563e105b1647a4fa5d42bd67"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_linux_amd64.zip",
     "sha256_checksum": "42f1f4fb47696b3442eca12bce7cce6de0b477b299503ddad6b8bc3777a54b5"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_linux_arm.zip",
     "sha256_checksum": "961b7fec553ea857757873943dc2cc0ed6214e0b4ac2796f2c1c15a8b9876869"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_linux_arm64.zip",
     "sha256_checksum": "113cbde951cb888cffeebdd9a32e35dee51309ea27edbd358e046875aa77da23"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_openbsd_386.zip",
     "sha256_checksum": "b0f96e7a073ea42d946005f7be926d31895ca3e903db1eb9fe3177b52d929452"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_openbsd_amd64.zip",
     "sha256_checksum": "57761864988522ce30ae7b82c6e70b20aa187c22ed2fb74a8291d887cffdaaf0"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_solaris_amd64.zip",
     "sha256_checksum": "ba95bee61a5c07c6460f66525b5d1ea49a7a92e5bdbc710d23f2263eed75a1dd"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_windows_386.zip",
     "sha256_checksum": "39d0a13046446171d7362d49faa3ea045b6c957fd153d4ff48485a9aab7099a7"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_windows_amd64.zip",
     "sha256_checksum": "ee052394a21c0f63e5feda7ba190c6019e286f0fe60a8c3209458b9def0ef17e"
    }
   ]
  },
  {
   "version": "0.14.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_darwin_amd64.zip",
     "sha256_checksum": "6fcb3898b33887fdd3c8f14cf92783a52bdab224164db972e65301f30baac3df"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_freebsd_386.zip",
     "sha256_checksum": "e0b581408b52913d58478a23bc9eba3fc3c7e56ebd6c8f7cd10195a9b60b46f2"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_freebsd_amd64.zip",
     "sha256_checksum": "3d2583d905e4229c79f8164ac6ca19891fc6e2b739769be6e83d7d4dbf81a0c"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_freebsd_arm.zip",
     "sha256_checksum": "d1e36d27088f22af27691720ecd08846e755956e888651868067334f82930434"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_linux_386.zip",
     "sha256_checksum": "d0e0b50625236fa29a1980fee4e81658b83116baece5f3deb19901359a6a523d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_linux_amd64.zip",
     "sha256_checksum": "322fe01fb54118ba25c66f4a406d25e3d8625de062d78544579e9385a3974496"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_linux_arm.zip",
     "sha256_checksum": "7475073aec8ccf837197a2d8a12130a3538b05f29825d942c6eaaf674c078565"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_linux_arm64.zip",
     "sha256_checksum": "269a714e0daeeabef789525e4f9976ade39afa6490bd3aaa99c8d21ff85c7083"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_openbsd_386.zip",
     "sha256_checksum": "ba504cc3fe63687c1c1d8e296e7ae42e3b9728928f462e32af0ff2d08fc5916e"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_openbsd_amd64.zip",
     "sha256_checksum": "170488953854d0cbdb5242329ac43c0192ff297ad4c82910cc4d3443007da8e1"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_solaris_amd64.zip",
     "sha256_checksum": "9cd581a459a0d5e844467c493709ca954f90b501db3bf6d34f0c9b5b649bdb0e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_windows_386.zip",
     "sha256_checksum": "93d179d5e8695bbd48d254cecdc186c69875ec07f6cad8a9675a3c78ae36b422"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_windows_amd64.zip",
     "sha256_checksum": "d5bea3f5aa34c83f79153e84916e6014bb4ef1455f2cc1808f0e539c2a163561"
    }
   ]
  },
  {
   "version": "0.14.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_darwin_amd64.zip",
     "sha256_checksum": "4b2acb55c6350cba92769c852d4502dff3e185726fc5293e3ab0bb64393846c4"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_freebsd_386.zip",
     "sha256_checksum": "3f34d609a8c54a84eddf4f519e4f53ca59e7016236862093737eff6091f8092b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_freebsd_amd64.zip",
     "sha256_checksum": "6cba2ba82edabe2f329000283323988c17d4445e3fd1f19534bf1c22d2319cc"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_freebsd_arm.zip",
     "sha256_checksum": "5a706e3096ce5855efc52b6ffb02796450f86b9489604e7ff715409159470745"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_linux_386.zip",
     "sha256_checksum": "7292b67b78178da2076b79b6a7c67352fe44d3df33c0009038fea778828f3bb1"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_linux_amd64.zip",
     "sha256_checksum": "45d4a12ca7b5c52983f43837d696f45c5ed9ebe536d6b44104f2edb2e1a39894"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_linux_arm.zip",
     "sha256_checksum": "3594abff30bd487e814ea458321fb003541b361c8441d65650a44d08bcce0d0b"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_linux_arm64.zip",
     "sha256_checksum": "e69e69796e7dbb9462d52a22295541fbd7ffa5a9b979ba4bee963309bc65a73d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_openbsd_386.zip",
     "sha256_checksum": "e45b18003c7077a259aa9bdd71aa94943501eff4c9e5fed5ed756143d05653a3"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_openbsd_amd64.zip",
     "sha256_checksum": "3b26cb7ea73fd5364fc0c728a12c305512b597cf5fb969ce980f345957e5f68c"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_solaris_amd64.zip",
     "sha256_checksum": "7a92ed3c81539629847dec9178a9eb0f62a388bdbb6955b48ab7dd70feec795e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_windows_386.zip",
     "sha256_checksum": "c73fd3ff67f535ca2d1192484281bc22edc209938a2592db9660e0ffa95932d5"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_windows_amd64.zip",
     "sha256_checksum": "6b40c552cdc91af4b9aa854f6e8e41d9fcc6b549f4cef1df6c1b9d1a6c7d0535"
    }
   ]
  },
  {
   "version": "0.11.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_darwin_amd64.zip",
     "sha256_checksum": "183078bf230e517e6f41e47d6e7d3b61093c6bb5a2b85958c01a4cf3949b7c14"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_freebsd_386.zip",
     "sha256_checksum": "65d0b63983dc6d7a0541bf2a4edba9e614a764f732a71aa0aa11128f66dbd0a5"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_freebsd_amd64.zip",
     "sha256_checksum": "d042fc43ecd270d7f7f629f4ee22638ffe4c530ee63335c08b44b04ca6e21094"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_freebsd_arm.zip",
     "sha256_checksum": "e4a79ed7b18b8989a6f44baa00e1ba5c972b6a7c008621e1e1a79f30f5da95b5"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_linux_386.zip",
     "sha256_checksum": "bcf425cc3833edc1e4eac817d4d32cdef252b479fd9a9e5a78325c05fe7e0d80"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_linux_amd64.zip",
     "sha256_checksum": "6b8a7b83954597d36bbed23913dd51bc253906c612a070a21db373eab71b277b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_linux_arm.zip",
     "sha256_checksum": "557e5f37633c88128f323cfbaf53b4bd05c74f2b91a11dacf889bdd56af40c5d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_openbsd_386.zip",
     "sha256_checksum": "eb022a74a7c31347999a757c92e1cc4afcb91dd12bf19dc1be01aaaab70f1497"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_openbsd_amd64.zip",
     "sha256_checksum": "bad48cd97c25ea290cb3585dd9b56481b167a078f8e2c0e125cfab7303df5114"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_solaris_amd64.zip",
     "sha256_checksum": "99a749d5520e12486b90eef7dbcb755a577ac126206f90ea928e592557283ec5"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_windows_386.zip",
     "sha256_checksum": "fec2aafec4fc942530e5c5bef5d3f634133dd3d88aab79d29940725c7faa6bb4"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_windows_amd64.zip",
     "sha256_checksum": "b4ff550aee6288ffd0911380a927cf216678121c3d30feb22aaf84446c3595bb"
    }
   ]
  },
  {
   "version": "0.15.0-alpha20210210",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_darwin_amd64.zip",
     "sha256_checksum": "8585395617e78abe64cde98aec5495856f812d42ade11b3c9a6d50f5c76e9f06"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_freebsd_386.zip",
     "sha256_checksum": "edb0a5ac1982c7db66a100c27402bcbd042cfb8544352b17fce7661bafee0203"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_freebsd_amd64.zip",
     "sha256_checksum": "b1dd692e5b0d8f271f9376c6f27ec1aff3e218cd05475e4f76717a995cb7553"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_freebsd_arm.zip",
     "sha256_checksum": "47781c93369ae972ec8f5debaeeb8d98af301ccaf1b635027ee72066c1cec8a5"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_linux_386.zip",
     "sha256_checksum": "24672e40ede483df76a12997ff4ef3bd07598c11d6ecf4d7a96da0b2f69f832e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_linux_amd64.zip",
     "sha256_checksum": "d08ab725461b27ecf8835b618d0cdf2159400ecd04422f8eee61a9ca54e55ea2"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_linux_arm.zip",
     "sha256_checksum": "dfb23d8f4358bc39f8fd75b4a45ac8cdb3b48c2504cb09ab2e9ee39987c41c28"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_linux_arm64.zip",
     "sha256_checksum": "c88569b94f6d8706ec812ffe29564cec1b319cae746aed9d807cccdfcc76497d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_openbsd_386.zip",
     "sha256_checksum": "28ed9feee7ad585269e28e5666ec8c01123ccec80fa486297dc0ff36d62a7b4a"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_openbsd_amd64.zip",
     "sha256_checksum": "45a67dade3481c85973297cfe0f4b6c1ad20fd8bc14942720f26a138cd90c80f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_solaris_amd64.zip",
     "sha256_checksum": "3bf89b4cd9cd1f643ef86b49ddde9d6d430e830f271646363ac502cd8f2e769f"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_windows_386.zip",
     "sha256_checksum": "4194fd153c4b59dc4fb6586d6cff6bbd4a6dabbc57662ec563000eb14851070"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_windows_amd64.zip",
     "sha256_checksum": "ca111056e6a3a37b9c1d1648adb0bba10e18095e4130f0598af8680e157ef1b4"
    }
   ]
  },
  {
   "version": "1.2.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_darwin_amd64.zip",
     "sha256_checksum": "f8eecc764b57a938aa115a3ce2baa0d245479f17c28a4217bcf432ee23c2c5d"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_darwin_arm64.zip",
     "sha256_checksum": "d6b900682d33aff84f8f63f69557f8ea8537218748fcac6f12483aaa46959a14"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_freebsd_386.zip",
     "sha256_checksum": "1ed98d6c1cb87934b06945643f0305d743f5a71c455d1db5f33c17e928b6cc17"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_freebsd_amd64.zip",
     "sha256_checksum": "88c0d0b14e227f207861fe9e69d7524c673de22adac741a93a6cfda4aa7f8db7"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_freebsd_arm.zip",
     "sha256_checksum": "d50572247338cf74dace04a1e5374b238d9f79bca4d2c9a93e3d85f972e8b1a0"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_linux_386.zip",
     "sha256_checksum": "2a63acc4d367565643dc108fd48fbbedee63c8720fa6aa85857b67728cc33835"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_linux_amd64.zip",
     "sha256_checksum": "3e9c46d6f37338e90d5018c156d89961b0ffb0f355249679593aff99f9abe2a2"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_linux_arm.zip",
     "sha256_checksum": "188d6a4cbaea522a61981b3376c1be63a28db2b30b209a16c030a54494bb96a4"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_linux_arm64.zip",
     "sha256_checksum": "26c05cadb05cdaa8ac64b90b982b4e9350715ec2e9995a6b03bb964d230de055"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_openbsd_386.zip",
     "sha256_checksum": "c826954254906b896d4bd17c7c8126a14ad8fe7c9fdd591a7001c5d6bf48d891"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_openbsd_amd64.zip",
     "sha256_checksum": "e63b6bce8c5ee5dc90967f22085f6809b9bf9bc2609de5e2a42ff34b94d59bf4"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_solaris_amd64.zip",
     "sha256_checksum": "c7bdb56e5e7841bdb7d1a7b105e6053cb34f95575679114aff3b14af5ab77d20"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_windows_386.zip",
     "sha256_checksum": "b52ed7b02e0300e03273fb7f801290d093d7f5b8eb3d7151add60682aefe1784"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.8/terraform_1.2.8_windows_amd64.zip",
     "sha256_checksum": "977924884ec36dc06027168bd62ca48cbcb9a5018fb6368978046cfe5692d37"
    }
   ]
  },
  {
   "version": "0.12.0-alpha4",
   "builds": [
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_darwin_amd64.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_freebsd_386.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_freebsd_amd64.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_freebsd_arm.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_linux_386.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_linux_amd64.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_linux_arm.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_openbsd_386.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_openbsd_amd64.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_solaris_amd64.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_windows_386.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_windows_amd64.zip",
     "sha256_checksum": "n/a"
    }
   ]
  },
  {
   "version": "0.15.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_darwin_amd64.zip",
     "sha256_checksum": "54aaf48e4e2daa5407b8d96da4c4e84305e77579b9af118d0c8df4a2512e2bd7"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_freebsd_386.zip",
     "sha256_checksum": "6248d189bfbc9b92e78b6f00a5e28d46a71a3e0988eaf670e9c68e3866b11068"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_freebsd_amd64.zip",
     "sha256_checksum": "97a3b2c7fd86e9855dd178ff49f7bfb41a2fc27c55c22792dd9c0a5982ec66f8"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_freebsd_arm.zip",
     "sha256_checksum": "c6531feeae9a65166dcba062aadeddec4a2ed80124ae0c86f74dcc13550ea943"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_linux_386.zip",
     "sha256_checksum": "8f080d88dc3d0dd9ab03038368028c9df42d997cfb19a8c46238ba353dd4079a"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_linux_amd64.zip",
     "sha256_checksum": "a29ad95e83a7ea4d4007e962182dfc4ca393d66f9963fcc06fc79884f8e35bae"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_linux_arm.zip",
     "sha256_checksum": "6b3dbd2851a4b85f76f69295b43bb810026cee18a9ddcb3f37f2c4903eec626b"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_linux_arm64.zip",
     "sha256_checksum": "7dfc1688e7df1eabed6509aae5ec71ffa41cd591dad20401b2f451a311b794b"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_openbsd_386.zip",
     "sha256_checksum": "a6dab2a0f718561f1580234521fab39d8fbe2e37dc289e67de4f83d9aed5d7ac"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_openbsd_amd64.zip",
     "sha256_checksum": "13ff6deded271428a59e609a147c1a276d541286c70571a0e754feb6bd34b972"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_solaris_amd64.zip",
     "sha256_checksum": "6153fbd7d4a7a5c89841f69c81c6e8b85db2cecd99fed91166d4bdc0fd83a003"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_windows_386.zip",
     "sha256_checksum": "a4d642a778dadd1026a502476e3a36337fab00baae17718966a7e9fabec8a9be"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_windows_amd64.zip",
     "sha256_checksum": "367b13e513983930575a70f712d89eaa1213dfa25e1b0c9807bd64af181ac3df"
    }
   ]
  },
  {
   "version": "0.12.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_darwin_amd64.zip",
     "sha256_checksum": "dbb3c0ffb37a5e659e05b8c223a717f89ffda7761d23eaf596c31b9745557288"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_freebsd_386.zip",
     "sha256_checksum": "b7b6559a69a48931b3d482466c70209d546aae9944c29b4132b75d42d2d2e9a0"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_freebsd_amd64.zip",
     "sha256_checksum": "586c05893b7c05812b0b10d6695d2e39338bc8fdc88ba399a22b7cf25c4491b3"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_freebsd_arm.zip",
     "sha256_checksum": "56bc688325e9ceddafd14b199fd72c545421a88c8d5dd4ca6b768ed11d6597ce"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_linux_386.zip",
     "sha256_checksum": "2710dd7792ee5ff594969b98ac33d4397653d6df2ca1d544cbe5b7f748bd4336"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_linux_amd64.zip",
     "sha256_checksum": "69712c6216cc09b7eca514b9fb137d4b1fead76559c66f338b4185e1c347ace5"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_linux_arm.zip",
     "sha256_checksum": "10943e9f8cab8657adb4325ac6aba023784242b6ab5e5af68c3634095efec9df"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_openbsd_386.zip",
     "sha256_checksum": "56969afd5d93de6aaf5fb57c8f5a053f9e392f8992190661c22b75c8cc524c22"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_openbsd_amd64.zip",
     "sha256_checksum": "549bec5fb77018ac2f44c45823c24755ba3999e2a24ba4e185ed7f5bdf3acb87"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_solaris_amd64.zip",
     "sha256_checksum": "16615d658592936c573288f5c2cceca5759b6d778a9d1d3f8c2bad8e5f11d8ba"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_windows_386.zip",
     "sha256_checksum": "251462474f205ed747f3fc41225308d240c53341d7fff0ffbbf9811bc7d9bf7d"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_windows_amd64.zip",
     "sha256_checksum": "a6f4ade992b270ae478b27cb823a5d2b3fa435409add299258aee1941ff4353c"
    }
   ]
  },
  {
   "version": "1.1.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_darwin_amd64.zip",
     "sha256_checksum": "5e7e939e084ae29af7fd86b00a618433d905477c52add2d4ea8770692acbceac"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_darwin_arm64.zip",
     "sha256_checksum": "a36b6e2810f81a404c11005942b69c3d1d9baa8dd07de6b1f84e87a67eedb58f"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_freebsd_386.zip",
     "sha256_checksum": "f5ce887f376ce5025f12b4ee8c6f197b8d48bbeae6de4b64ea6233f547debd57"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_freebsd_amd64.zip",
     "sha256_checksum": "2e208a0631ba309cc90c7fe8a91786002170676765ed77aeb8b2757ff156b619"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_freebsd_arm.zip",
     "sha256_checksum": "1249a3afface2f44aa438bd3f8936da4353eb96f4d0c540269893be1fef45d05"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_linux_386.zip",
     "sha256_checksum": "9bc98bc3e87887bf464dc0a8e6ac91b9132eaab27e99cf393980be3760e8ded"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_linux_amd64.zip",
     "sha256_checksum": "e4add092a54ff6febd3325d1e0c109c9e590dc6c38f8bb7f9632e4e6bcca99d4"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_linux_arm.zip",
     "sha256_checksum": "81013a3d554a2a40ac9f59dd1dbadf3e4f43d6e04c14d746a8eeffc45788f8cb"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_linux_arm64.zip",
     "sha256_checksum": "2f72982008c52d2d57294ea50794d7c6ae45d2948e08598bfec3e492bce8d96e"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_openbsd_386.zip",
     "sha256_checksum": "c7accdd5caef9f77349e81dc1a52f8f3e003c6f387002ba2f7f012f0db64fd59"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_openbsd_amd64.zip",
     "sha256_checksum": "a7be1f3c0ecb104ad6363bfa8f9f5d93d81d204ec63c3b370ac810cd79a9720a"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_solaris_amd64.zip",
     "sha256_checksum": "685cb1a69152423656800bb02cfe2a4427bd172b727fe6e952b3f98a5e9b142"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_windows_386.zip",
     "sha256_checksum": "2013a810e0e52effd28210d97b1a8d88b57216ddbcb2d654674b50bc2d7129c4"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.7/terraform_1.1.7_windows_amd64.zip",
     "sha256_checksum": "6d971320bec87372f4416263217ba80869e50d423c8094084462122629aa4c4c"
    }
   ]
  },
  {
   "version": "0.12.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_darwin_amd64.zip",
     "sha256_checksum": "7168dfa057d9aed7ea3f111d87294f263e341c8b848e776bc13d169ddf2926c7"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_freebsd_386.zip",
     "sha256_checksum": "e1c32fc6ea4d224ff7e9cb2bcb98353cca7174214e0b52c33dc67f2a192e8f69"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_freebsd_amd64.zip",
     "sha256_checksum": "b7ed3bbb36a2cc8a71ceac703ae834362a3051ffcc0885298e6ab8b59bbcb325"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_freebsd_arm.zip",
     "sha256_checksum": "1bd3063e91cf956501141f0fb2743f2087360ed37476f924b4e879a6838b389"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_linux_386.zip",
     "sha256_checksum": "faa1e4aa3207e3969179caa283d5abe36e9f799e157c01e16619041166fb3a86"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_linux_amd64.zip",
     "sha256_checksum": "6544eb55b3e916affeea0a46fe785329c36de1ba1bdb51ca5239d3567101876f"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_linux_arm.zip",
     "sha256_checksum": "653c376dcf6240cfc52029af48858b148b09cdc148813777de1ab121b3e4bf7d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_openbsd_386.zip",
     "sha256_checksum": "86bd4d548b2d5a946b857d6b91671d85d7ba0a7342dae120e710921afd4c6820"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_openbsd_amd64.zip",
     "sha256_checksum": "a432094715c62c1e40000026bce6005804d6063b7aca30f03d4e48022c30ce74"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_solaris_amd64.zip",
     "sha256_checksum": "ea336616a5f7391cee874f6f7c8b061f41c97bfe75044683da8d19815b9be3e3"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_windows_386.zip",
     "sha256_checksum": "98cf29bf299223ab0081d9a35c182aa6415e0582a6c783ab0029f8a11d30b85b"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_windows_amd64.zip",
     "sha256_checksum": "8a60bb49f75d766d8915c75f0b8ae382b96959587e090c3096202570c9f520b1"
    }
   ]
  },
  {
   "version": "1.1.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_darwin_amd64.zip",
     "sha256_checksum": "29ad0af72d498a76bbc51cc5cb09a6d6d0e5673cbbab6ef7aca57e3c3e780f46"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_darwin_arm64.zip",
     "sha256_checksum": "d6fefdc27396a019da56cce26f7eeea3d6986714cbdd488ff6a424f4bca40de8"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_freebsd_386.zip",
     "sha256_checksum": "7beacbe7079a19ae48960b22190020debac7980a3f561373e61f8ead57f15d91"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_freebsd_amd64.zip",
     "sha256_checksum": "8c1dc41138580eacb2eceeb64630df4f09d95010b402c99058f9aaee57c4ce16"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_freebsd_arm.zip",
     "sha256_checksum": "368a60bdf8989448118c2fd3e3faa8e57ee3bee7babc5376ed74995c0a0b361b"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_linux_386.zip",
     "sha256_checksum": "6523d11f849d09f2822d0dbaaecc820f3536834c8597ed4d0c9e1749aefba795"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_linux_amd64.zip",
     "sha256_checksum": "fbd37c1ec3d163f493075aa0fa85147e7e3f88dd98760ee7af7499783454f4c5"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_linux_arm.zip",
     "sha256_checksum": "c2dac55b0ba4e625d0afe77eb4eb3f65ea4e3b5ed930de6217ceeb46c19d55e8"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_linux_arm64.zip",
     "sha256_checksum": "10b2c063dcff91329ee44bce9d71872825566b713308b3da1e5768c6998fb84f"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_openbsd_386.zip",
     "sha256_checksum": "212dc29da50c54151e77b6225c67bbaf2bb99e92e68572162e29099c45fa39ac"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_openbsd_amd64.zip",
     "sha256_checksum": "f7e0d31b6e74720b7f2e8634f7a1261a352eafb8705f2ce4fd5aadcdc3b7206f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_solaris_amd64.zip",
     "sha256_checksum": "d38e0a1fc2dc136824e62eaf4e341a5dbb0401c251da4c32061281486de8f788"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_windows_386.zip",
     "sha256_checksum": "30b4254b690d612f9e94f1c2dab887bc90a7c03b217c063b3bd3f1e70d94cc2c"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.8/terraform_1.1.8_windows_amd64.zip",
     "sha256_checksum": "a14ccf89a3e534d2bfd7fbc1229865d061ac68309a3ff05852f2e829010a517d"
    }
   ]
  },
  {
   "version": "0.15.0-rc2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_darwin_amd64.zip",
     "sha256_checksum": "2fd60ccbd3d01fb5adb57863c8cb0df98e0743d6a7a9a38dffeea631cef95d9f"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_freebsd_386.zip",
     "sha256_checksum": "dd126a14dcdacbc9c5c3cf8cf5c0c91232e75168b06349c6db3528573207a2e0"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_freebsd_amd64.zip",
     "sha256_checksum": "3875c5bc3551d15f4ac2985d54538cc37a8f61a159e0c19a24a158244a7a3119"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_freebsd_arm.zip",
     "sha256_checksum": "c1152d2fb204ee4c671e9ba5e4e32a1aa167aa2f50c2fd7f5bbb6f0305c74ac6"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_linux_386.zip",
     "sha256_checksum": "9381e30c18b2cffc7ed8fb8142723139e1db6b3d11b3b119896523a119f35b41"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_linux_amd64.zip",
     "sha256_checksum": "74d2f116af4294b1f7ac76268ed21122c70ab5030968d46b5a2a4acf3ab35a6c"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_linux_arm.zip",
     "sha256_checksum": "c6b4bf5658042065fbae59c2df57e23f6be3da1ce0fa119f6e14a45b49aca741"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_linux_arm64.zip",
     "sha256_checksum": "2a5534d227dfb669b3d7fcaf39d2b24057ab78b058b3ec4cac1e5a6c05af2b5d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_openbsd_386.zip",
     "sha256_checksum": "57e8da6516878f496b982c9ca7387320b3d17283b5b56cb267eae199db905c9e"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_openbsd_amd64.zip",
     "sha256_checksum": "d47692072b9fde50a63b47ed51819eb44ba6a001c43658b4142824ba4ad351b1"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_solaris_amd64.zip",
     "sha256_checksum": "c63645c59fb1b74dddf6ef295a671d044ded25a1b849be244b2fe40a1c8862cf"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_windows_386.zip",
     "sha256_checksum": "9a5076f5bd53da948dda858790d540985a1b7f191e6c7a44887d62aa0e80199e"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_windows_amd64.zip",
     "sha256_checksum": "f6caefd16f6cf8a096edc06f6437104ded877eb775f89329d1502d4edd34ce7b"
    }
   ]
  },
  {
   "version": "0.6.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_darwin_386.zip",
     "sha256_checksum": "5b1bf924c8d201a6efad13b59b3dd444d9a52cb0c426aa42a58705690b303562"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_darwin_amd64.zip",
     "sha256_checksum": "d5c50b38bdba7dd11ccd31ebe04de9bb4a1f31a8b30ba967c863e3754d1bfd8b"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_freebsd_386.zip",
     "sha256_checksum": "eb60543f6ca068b4a94b666593b694a30e7896d2537bc0f89d5cceffaad71a3a"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_freebsd_amd64.zip",
     "sha256_checksum": "2bd3d067bc711f51dec17ce96e9ae991f852beafb09871d44ecceb04894f766a"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_freebsd_arm.zip",
     "sha256_checksum": "358698388590d40e6f6ef391311c889670c4446bacbf84bbc7b6221f8843695e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_linux_386.zip",
     "sha256_checksum": "32db19ec45c46b5ab7894c75062fe163b4c9003945ee234979e5f4c8db448923"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_linux_amd64.zip",
     "sha256_checksum": "160fcdb7f0d00948d52912df0626a2e49db958b6df2c6108cbd8b3527ce1144"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_linux_arm.zip",
     "sha256_checksum": "1beba1345a2137e61ce27c57f67f9e21da3a2f02fac3718af4f515cb54392ac0"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_openbsd_386.zip",
     "sha256_checksum": "431ab57a658ed6fdf2ee309af9a01fadee37df368f510d552434f0793424e26f"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_openbsd_amd64.zip",
     "sha256_checksum": "ebb1f0702472c231e8e330f2ad4affc13e286e9f6e6a3b36a46b86d3f91a745e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_windows_386.zip",
     "sha256_checksum": "67e2a58e36f5066351fe92f7db5dfe4a9f3a329e5d8b1ccc1200bdf12705566"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_windows_amd64.zip",
     "sha256_checksum": "de363846cba24799bad3bd550f07ccec0df13e0aa6bfd508a86ff6f826e62e91"
    }
   ]
  },
  {
   "version": "0.2.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_darwin_386.zip",
     "sha256_checksum": "c4997e229924cc8c6e04167b0e6d2e7d9cf591cc8818589d104850e7389da304"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_darwin_amd64.zip",
     "sha256_checksum": "1b4581e41e05145d2e9707cad5313636120a80b04cb796a503b3bfe59b6901d2"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_freebsd_386.zip",
     "sha256_checksum": "55f325ee5e0a89e4a10a55c8d1c4b2fb2eb6b23889fc4ee01d64533ef72ae7e3"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_freebsd_amd64.zip",
     "sha256_checksum": "87f610128534b5ba0255d7126065bcdf73cd8b439f523d25fe232f76fdb3e6eb"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_freebsd_arm.zip",
     "sha256_checksum": "e7e628ee52a7db37f9c284f7459f91456f0054e46b244673c9573b09912be886"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_linux_386.zip",
     "sha256_checksum": "bb39a78e65a27c199575df1271bb5aae43be4fe376059cf5a36a7938af09b58e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_linux_amd64.zip",
     "sha256_checksum": "9ebc457e85b4c209532005c98c32e85eea6227534980095931b6a7e2c00c37ef"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_linux_arm.zip",
     "sha256_checksum": "f6588825f225ca73d3794c832e880d6975981c781ee41bd0ec717cecbd39e321"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_openbsd_386.zip",
     "sha256_checksum": "470843716d1c5960b6f9e3122b3e93b68e2ff638e873825603c9cec96dfa5548"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_openbsd_amd64.zip",
     "sha256_checksum": "b56d6c912db30021de8b780bde82b773e16b3e7b1012a4cf976c5e1c61d5fef3"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_windows_386.zip",
     "sha256_checksum": "623e36f9108a67b950dae05b30a321d18b5c4942056c4f18424696fb791cb743"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_windows_amd64.zip",
     "sha256_checksum": "4b68b0ad9dc5575351756b4ec3a3c5ff2e9203a037c1022569d945345dfce804"
    }
   ]
  },
  {
   "version": "0.8.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_darwin_amd64.zip",
     "sha256_checksum": "4f4410be73200f95f84e359409481c8c48bc70e659fc5f7ea3f33a1db574ff65"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_freebsd_386.zip",
     "sha256_checksum": "450ce0eca92eb9691a872128d6466fd2560a8a0a4e66dfa0a8d9785633ee64c3"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_freebsd_amd64.zip",
     "sha256_checksum": "149124bc0bcc8575045223f106b1ee833d2f455e299bd5ce95aa38f32574907c"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_freebsd_arm.zip",
     "sha256_checksum": "f6d83a0173f5a71c0fd07ae8b69c06cae9deb18f0040ac16276fc4d7ee3de21e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_linux_386.zip",
     "sha256_checksum": "bcaafdf0e62d4e13abc6d5d4dcf414310b1b4c78dd4907a7611cf135acc12e1b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_linux_amd64.zip",
     "sha256_checksum": "9f99258396dad71746fb96036bf421e05fb3f36f402b89c3f16b453f3f0ee9cf"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_linux_arm.zip",
     "sha256_checksum": "67bf62d384c94e436aa4f4791b2e645f4f1b1046f90250b0a99a1a8acc45ad6b"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_openbsd_386.zip",
     "sha256_checksum": "a79f6c6e299e975553639cebcf63778c97e4112e1ab674b5fdded148aca94800"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_openbsd_amd64.zip",
     "sha256_checksum": "4366c4556a978afe503cb183f35775c87c0888f4dafff589cbcfc9c159f84161"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_solaris_amd64.zip",
     "sha256_checksum": "2b72fb297907d6846db96c5a47b3371eab8a9cc329d307d3148ac7bcf05c2ad0"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_windows_386.zip",
     "sha256_checksum": "697692ffc969fa4bac0cf0c70d0ab97efbc3be38530cd63638e5d5b90a5dc3c7"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_windows_amd64.zip",
     "sha256_checksum": "2624b22fb2cf68679952190a8fa836f29a5bbae6de9de4b3c20a41ce116e4cb0"
    }
   ]
  },
  {
   "version": "0.11.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_darwin_amd64.zip",
     "sha256_checksum": "c7bbc03a40c089077e77befb3405c3fdf456f46e7b3bdafc50e48bfcc6f7b5a5"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_freebsd_386.zip",
     "sha256_checksum": "148287694cb044610dce7cd68e0047d9ea814182286682e4a9fd054964796037"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_freebsd_amd64.zip",
     "sha256_checksum": "867150cf3a5b5972847f3d0c6f7ee5cd5d772c77c7fea94c37e26df97956a6ef"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_freebsd_arm.zip",
     "sha256_checksum": "c87e27b655b85e6e283e2cf3fa1cbed9b2d7bed78a476807ebb874525304c1eb"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_linux_386.zip",
     "sha256_checksum": "55eb262a2f92e5fd0918639c3f21782805168d0bac32fa20a5c7c2c3c5dc883e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_linux_amd64.zip",
     "sha256_checksum": "e28701040fbc210c94ff0d1fb786dd7a1bbeb54021ad305eb396f139ac5d07c6"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_linux_arm.zip",
     "sha256_checksum": "9e2c7e74dee5999b8096b799f5c2a65b504de6c9dbc902b1279f85e9ed6de90"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_openbsd_386.zip",
     "sha256_checksum": "6810041fe9b931e8959fd66a63c97fb25fd0551bd9750bb1e5402b3a827ac55a"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_openbsd_amd64.zip",
     "sha256_checksum": "91b1f6bb8cd4de7f7763357ba10d036d3e50f90a21119b0be281a0114a93006d"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_solaris_amd64.zip",
     "sha256_checksum": "3dfc343840e79d2f119935feaa3c6bd3e97822ee9804146e6eee888199963e10"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_windows_386.zip",
     "sha256_checksum": "897fc71af09e8f1f9d4e408f0d50ee7803717825fdf667d018dd023add998b9d"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_windows_amd64.zip",
     "sha256_checksum": "af5be448538bc2535352892dc4867092e062338a0aaea3ddc0a444931d833bac"
    }
   ]
  },
  {
   "version": "0.12.27",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_darwin_amd64.zip",
     "sha256_checksum": "3941e8b3f81257e54997cd717cec5dfbf3a254643a47e3ac8c687f26c0b8814f"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_freebsd_386.zip",
     "sha256_checksum": "d444ba2d4a17a2889d2fe25fd1c65fbab2e01deae6d9c9f0f40d2ab0effb87d5"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_freebsd_amd64.zip",
     "sha256_checksum": "7f58a48c17df2e890e395c8ca69cc150ce94b2ca1c1b2d7871e2cdc4eda71042"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_freebsd_arm.zip",
     "sha256_checksum": "5dafb9c2c0cb599664a61e7acdab27ff5812f1e5c31ccfdabbb1dca36a3c57d"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_linux_386.zip",
     "sha256_checksum": "525fa15ce44a0610e5fa6ba8e35bd87bfdc293644a8fd64222c93f4f7508b453"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_linux_amd64.zip",
     "sha256_checksum": "6842303730775ceaeebf73d5f187fb77c194e3825fa1dfd6f879bc4dae3b798b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_linux_arm.zip",
     "sha256_checksum": "dbc40eca5d7b2becee1bf93fdd0d2150806a04e587075ed41e13130306111f34"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_openbsd_386.zip",
     "sha256_checksum": "7277c91d5b8f8eb9f62654ad0b7a5799aa2b956c24e4c4377d2787421a78d0df"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_openbsd_amd64.zip",
     "sha256_checksum": "95c35626deddb0361c8ab7f808ed92b484a81021585a5ec512dec2cf6dc10fac"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_solaris_amd64.zip",
     "sha256_checksum": "d6ac1b0c5682e13948b2a48408163aa9483918ddf86183b111efa18bf7f4309"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_windows_386.zip",
     "sha256_checksum": "f2327077cdec5046b2334882b5658ae154a00548c6f9bc172ea36802ed28d8c3"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_windows_amd64.zip",
     "sha256_checksum": "1b690cb8e74a35c743fa053f63f0c719c25d0400a8539627ec01570df47ce5cd"
    }
   ]
  },
  {
   "version": "0.11.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_darwin_amd64.zip",
     "sha256_checksum": "5a8f9118bf99285aa41c60b150fb628ec6a1bc49293663fd2255eedc5934f379"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_freebsd_386.zip",
     "sha256_checksum": "9ce87f3f8fb36c806c09d538ef85c01140d1b55a1b1ca2116c1f300a99da2830"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_freebsd_amd64.zip",
     "sha256_checksum": "3af76f4fbb0656b1e39665867b5184dcf3e85b075454404805ce5d5ddfa9ffa5"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_freebsd_arm.zip",
     "sha256_checksum": "bc0a9996a868f8a2bd3be5c6112ef5d6fd0fb22f491eba8636523af9148e2617"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_linux_386.zip",
     "sha256_checksum": "540c8f526342f501d9ce6973f33d62fb64273744e7f695a8302acc67883dfd42"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_linux_amd64.zip",
     "sha256_checksum": "f6bcdc1e159ee0aea2892cda996cd8ae1b3350a57e963a1d7e70d10be73c4120"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_linux_arm.zip",
     "sha256_checksum": "2fb3b019f73fc5310f6e4a0828641d9a1fe30c3d7e0049d63f1807370500c3a1"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_openbsd_386.zip",
     "sha256_checksum": "6272555a0f2b869fbbdb6577b2d9474d7ce2de949e1c63f35d9155c4fbc12b9d"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_openbsd_amd64.zip",
     "sha256_checksum": "7f458949e16f372e7cbad94164ee9ba0249a84ab2da937d12e32e54f9e94cd93"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_solaris_amd64.zip",
     "sha256_checksum": "cf9e7c07680695c300b58bd80d80974d076cd9f6f9e0a0e3a722303c6d097fc4"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_windows_386.zip",
     "sha256_checksum": "c9de9cf770c3f8a6d945c89c6b0ea98f7806e98d2bba3c8364411e2b5155442e"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_windows_amd64.zip",
     "sha256_checksum": "d58f5304cd38ab644e765f931ca3b75aed5eba891a029727a2928695679fd88"
    }
   ]
  },
  {
   "version": "0.12.31",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_darwin_amd64.zip",
     "sha256_checksum": "c1a6ca04026cebe7849610037ebc960609484c25f47a58344efaa7fcd5be1e56"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_freebsd_386.zip",
     "sha256_checksum": "21dacb2316fcedac3696c06774ef63cc8808562f24dbb8a58aec99c41bc102d9"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_freebsd_amd64.zip",
     "sha256_checksum": "50a5ae42d8ed3540d5026a534ca03fec52907abe0e098200fb7701d63bfa5ef9"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_freebsd_arm.zip",
     "sha256_checksum": "151fb293ac76976deb3406d3398a5cfcdcd8e99c65bdf0531212d425e8fa6b3d"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_linux_386.zip",
     "sha256_checksum": "abdd13b29cac315900246bf8ebff19040dbc1ae6ce99c1348e1bc59ba04241f1"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_linux_amd64.zip",
     "sha256_checksum": "e5eeba803bc7d8d0cae7ef04ba7c3541c0abd8f9e934a5e3297bf738b31c5c6d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_linux_arm.zip",
     "sha256_checksum": "f5894ebdc8aceb12da840045508c362eb0f923e4b5115bb601fefa8f1f1d3cb9"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_linux_arm64.zip",
     "sha256_checksum": "737205d76f576e47a5d7a41e506b2ec8eb7a0f6a0b3b3b0a8de59551c2f2b77f"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_openbsd_386.zip",
     "sha256_checksum": "fa837f579c7e05f2f1758325f3538dc0cb446ef5775a0aa8a9a758605086a532"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_openbsd_amd64.zip",
     "sha256_checksum": "a99bccba655b1392049fc7a462797b96ea128598058de588185b4c8aec09cef4"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_solaris_amd64.zip",
     "sha256_checksum": "eb2c4c98971b71f918b4b11c2d38c12b3222a3d1b85abec35762ac0e8c17373e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_windows_386.zip",
     "sha256_checksum": "8b71e3f6ffc7733ac351c35f4f04acc87a962f5c46a647b5ac7c433d7725619d"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_windows_amd64.zip",
     "sha256_checksum": "f5de5733253eb5a32c1bc8d220fdc0ae8230892d356908112268289c808cbb25"
    }
   ]
  },
  {
   "version": "0.7.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_darwin_386.zip",
     "sha256_checksum": "2150c3f97f5f96109646aae95eb09f1731cbd350938cf1e64b7470c27ba850b0"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_darwin_amd64.zip",
     "sha256_checksum": "69c8d2b07f04e9bf0beb4a333dd189d8616d22fe46692bdb5aef10493ac5e5c6"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_freebsd_386.zip",
     "sha256_checksum": "b325e5e368d316535aa87e158173622fb48190a3c849f864a2c3299f37bc968"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_freebsd_amd64.zip",
     "sha256_checksum": "8998482a0871c811579a643bab0b851e44fe3af22cba1cb61017d0f60a03a9f5"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_freebsd_arm.zip",
     "sha256_checksum": "823c4ab25db3c7b5df028fa7aaf5664d382308146327fd4ac8761daa906e79b"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_linux_386.zip",
     "sha256_checksum": "6fa6a6db863829a8e3ff64cb2a870ad622c14880576520587f15d36c873be2b1"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_linux_amd64.zip",
     "sha256_checksum": "1a1aa419d614ea38d93cc751e00874ace9414184c75e66480384c464e8a634db"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_linux_arm.zip",
     "sha256_checksum": "95e925b7a65ded50f2ac91d770d2817ed442884d79afdc87c6d5dfda1d8d2eba"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_openbsd_386.zip",
     "sha256_checksum": "45fc6778b51c1054523a8f1138707eb7982de069866a7d69965c8ac371bdeddf"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_openbsd_amd64.zip",
     "sha256_checksum": "eb9588953a6ba7b068697a60117db13b7bde00e8416de279dd886bd5e406f6ca"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_solaris_amd64.zip",
     "sha256_checksum": "a52e1cc516fd83a3f1e06dffc57d52045487eb68f3db50a1af6837eff280ced6"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_windows_386.zip",
     "sha256_checksum": "3dba63da427becf64699c8c702466bb88927d3b70ad77dac90b11f524091acaf"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_windows_amd64.zip",
     "sha256_checksum": "3bca261be81e5c3ee602920d5e092b1e2ddba46f4287b30848fef05157fc1b1b"
    }
   ]
  },
  {
   "version": "0.12.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_darwin_amd64.zip",
     "sha256_checksum": "e19691d775849888a0695a07e52a884dc617ca2100759eca5bbe4d0f428a7bc3"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_freebsd_386.zip",
     "sha256_checksum": "2fa6f8c67bd739113f0c3a5dfe3d683952333d400bc3560b581de4eaec7851d2"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_freebsd_amd64.zip",
     "sha256_checksum": "996750c4de58bc290d2475a321ef691c50a6449affc242f76d2abdf2eb8a33ef"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_freebsd_arm.zip",
     "sha256_checksum": "ecd8da9787681407f9014496b463b2904002dd5cdeda09a163942868f08ed372"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_linux_386.zip",
     "sha256_checksum": "cd3219d838cd07065fc19135911d296a69adebddef669d6473e3fc6ba2d46b3d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_linux_amd64.zip",
     "sha256_checksum": "231562f26262de233e8e8fac668196af21b7aff355bb04f3ee1606cca239c0a"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_linux_arm.zip",
     "sha256_checksum": "f975c058eae0596e2567e3f3341bf131ff0193dc342e66a5e5794000e9f7c997"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_openbsd_386.zip",
     "sha256_checksum": "a92bfd61f4776c7681cb7046142be32f8c1e130434bff72b3afbe4faa16d52af"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_openbsd_amd64.zip",
     "sha256_checksum": "11ac05d48dd3022fec6298acdf7f671f79177776cc73305b4e52d995b7e8422a"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_solaris_amd64.zip",
     "sha256_checksum": "87ef94d5b026e372fdc42fece4b96e125b97b1405b4cdd4bead78275c50c43a2"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_windows_386.zip",
     "sha256_checksum": "c0056084d9dbae8955360f1ad49984134c10c8510fa4161a7cf9808eb77c14ef"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_windows_amd64.zip",
     "sha256_checksum": "3dfe07bb605d6bd098b2eb374c677e9ffdcda7f114c8a4f62466d2864040e623"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210922",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_darwin_amd64.zip",
     "sha256_checksum": "22f9ddf073127d051e275b674a24cba2e3034f7e13df08ac8da73bc980c82944"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_darwin_arm64.zip",
     "sha256_checksum": "49698f86bb2dd8c58c7e028105890d631f3c705ba210c0719b1e50b8a85a5501"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_freebsd_386.zip",
     "sha256_checksum": "350da05fb0865f9e43f6816981df61fb487d0dcce064d3ffbebd3e25b8ed4e0d"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_freebsd_amd64.zip",
     "sha256_checksum": "5d44f1ded789e45c6058f55167066b7831e21fb495c4bf11b3dfb08b2071a874"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_freebsd_arm.zip",
     "sha256_checksum": "a3a42bb575db7112fa00c6f9ba17fa811058a41ea8f21d067292091727f3b4a0"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_linux_386.zip",
     "sha256_checksum": "b96eab1f5bd966f6cccfd3fe53f7fc9f5a56897aaf47d96fda9d83816e385273"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_linux_amd64.zip",
     "sha256_checksum": "572a98650f8481d67bfce7ffea12baf675a9ea389661037e7d9b48a9a2c08ce6"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_linux_arm.zip",
     "sha256_checksum": "92cba0db171552850a91cf19726e19a2604154045510174ef9d06ffff25a5e63"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_linux_arm64.zip",
     "sha256_checksum": "59efd7944d2b1b9c8e82dfde212bf8b2cc70aad928429e0eb0a5302afb7600e8"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_openbsd_386.zip",
     "sha256_checksum": "ba09c12c5af64376898625c134edf0260f20fe1bbbe157826a82eb60fcaaad0b"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_openbsd_amd64.zip",
     "sha256_checksum": "211d8d73e7f17cbb66c191928784e5b1bdb385a32f3d08fef29fd62e797775cd"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_solaris_amd64.zip",
     "sha256_checksum": "fb83f2530edf284fedcaf8dbb401148cc6d5b84b1162a0893af51d4cff46a72a"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_windows_386.zip",
     "sha256_checksum": "89b525b6e47ffdecaaf0633c6cf9589e55101fff0e408e0e5b0893a9459e64c1"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210922/terraform_1.1.0-alpha20210922_windows_amd64.zip",
     "sha256_checksum": "938f6eabecbaabeae643e1e83eb98f82bbbc31c72cb21b0a3f08e05a75ee15aa"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210616",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_darwin_amd64.zip",
     "sha256_checksum": "941fbe1607ed501b278a9e4c42fb8c5dc11366aa09f5c30139bb1a209c4c2c6c"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_freebsd_386.zip",
     "sha256_checksum": "a479d4e4fbbb2d3cf80ffd3de60e6dc1632161946072c1b521a638d15e0425cd"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_freebsd_amd64.zip",
     "sha256_checksum": "d91618a4c90fe851e54bab7348cfcc87fc3fa0cb0fa4784dd62417ac2528164e"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_freebsd_arm.zip",
     "sha256_checksum": "7ea655ba616386035ade07c5dfad0feb7a124544b3f3990937c8491f59175c13"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_linux_386.zip",
     "sha256_checksum": "2e6a3480a9b82cc73e8632da027352530c8dcfcc8eaca548b69b88c5b2371298"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_linux_amd64.zip",
     "sha256_checksum": "42ce17b679ae6dfa083c2026b271d99101bb533561e48572b4bf237f32e99a5b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_linux_arm.zip",
     "sha256_checksum": "b13d73d8a5458e05a222637840e1676c574ded1b9e58be1fa0ee5cd8055b86f8"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_linux_arm64.zip",
     "sha256_checksum": "abceb7711cdd76b61698c8aad05ebd532f5b93c67d88cfdf623d15cbbcdcf45c"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_openbsd_386.zip",
     "sha256_checksum": "edd7d4c7f8e66672398f640b558a68c60a34499324a03d3fe0b4cf7291721ccd"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_openbsd_amd64.zip",
     "sha256_checksum": "12d051610ebc4ad794b7a8651b91fb579f06c2e6a40d237f35fe3ea7fdc0ad75"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_solaris_amd64.zip",
     "sha256_checksum": "f79dbe588eaa55a66fca464875e67fdb2b301afd5717ca2119ae9571c5141931"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_windows_386.zip",
     "sha256_checksum": "74bad95273c48ad9c28777044104a3e1e2dc15eb70cf8735a0927b02a732d586"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_windows_amd64.zip",
     "sha256_checksum": "3ae8fda2564d899513655e94e1661599ca34d1e7eb79417a2d35bf3ece693491"
    }
   ]
  },
  {
   "version": "1.2.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_darwin_amd64.zip",
     "sha256_checksum": "3e04343620fb01b8be01c8689dcb018b8823d8d7b070346086d7df22cc4cd5e6"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_darwin_arm64.zip",
     "sha256_checksum": "e596dcdfe55b2070a55fcb271873e86d1af7f6b624ffad4837ccef119fdac97a"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_freebsd_386.zip",
     "sha256_checksum": "f156c5c423ea45e2c87503e3fe33c128973cc6859295a7f968fd86f2095763e2"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_freebsd_amd64.zip",
     "sha256_checksum": "9554583d4f2031c3fc67be7c119197697363d5a89ef47cfe923697c50594a193"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_freebsd_arm.zip",
     "sha256_checksum": "b3452f61117405e61e3676577ad35082b7aac7754a5482cdde31c8291cec0229"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_linux_386.zip",
     "sha256_checksum": "33942dd11c11f310fe7555ecc428a8bbbbd681197f1ceb82d63b72919b8e83eb"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_linux_amd64.zip",
     "sha256_checksum": "705ea62a44a0081594dad6b2b093eefefb12d54fa5a20a66562f9e082b00414c"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_linux_arm.zip",
     "sha256_checksum": "d93ac2b204a20e95e33b9293eddd33e5f4545c5582d9711379ac6d926f63b03a"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_linux_arm64.zip",
     "sha256_checksum": "11cfa2233dc708b51b16d5b923379db67e35c22b1b988773e5b31a7c2e251471"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_openbsd_386.zip",
     "sha256_checksum": "f54e9350e2be51e5e3e8a407f9bc9e7f8702eded9a3004077d9166650e01c688"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_openbsd_amd64.zip",
     "sha256_checksum": "db2d6bc12256d79b5d3e7739c3a20761ae7e20a4c5b03cdf6b056f83e16a9b31"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_solaris_amd64.zip",
     "sha256_checksum": "2e10d8f34999dab591e83effb10fcff2dcaf3df76f783ae5b8dc7b5bc3f0d20b"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_windows_386.zip",
     "sha256_checksum": "849f89babd7261292048a7829ccb2f77410a10453fe35e94509d9981b209ccfd"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.4/terraform_1.2.4_windows_amd64.zip",
     "sha256_checksum": "2e2a8bac8097f54ca015ac0ab1ac4f328e48313e6d63347cd846533b452eb8a0"
    }
   ]
  },
  {
   "version": "1.2.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_darwin_amd64.zip",
     "sha256_checksum": "d3bb52ecd91f19f5bd40a8049648cf836808155c47adc4a2140193a1874153be"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_darwin_arm64.zip",
     "sha256_checksum": "1b726e9d6318a03c67f6b2437d7024954f40a13d254d032498e182860094ecae"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_freebsd_386.zip",
     "sha256_checksum": "b9b5b8513951a4df972f47130e044159e71d72919d5d02a06281475fa03c14e8"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_freebsd_amd64.zip",
     "sha256_checksum": "a95638f7a45b1d4b3732c29fdb2ca7c645f0a9f16f3e53804b9e1f5793c41"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_freebsd_arm.zip",
     "sha256_checksum": "6983c3136c48a86758f0b882d6f4adfc701e41ed4c65b3e931c4ada28ef23225"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_linux_386.zip",
     "sha256_checksum": "3f0ecbc46b880b39719c3fb390dbfd53def388e3378e7c1a2563e54ab529c12c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_linux_amd64.zip",
     "sha256_checksum": "ea7dd9274d03eb1df1cb081a210b99c19586d8c2a09b29cc4a8e822e9eaa0c26"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_linux_arm.zip",
     "sha256_checksum": "9ab91f0ec51c29fa000774b497f56ff334835f5255fe97ca648a2f2aae4375d5"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_linux_arm64.zip",
     "sha256_checksum": "2f8c236065dbda1a43061f5825e46eccb54a43b8f31ff37a5f358fb46b27c29d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_openbsd_386.zip",
     "sha256_checksum": "2ae74fc42adffb14f657f42bf8c63adf3c1f3c4c0b76a39b0e98c3ab616d4f62"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_openbsd_amd64.zip",
     "sha256_checksum": "a7d65439762c94bce182e26e4aa349171f2c73718c8fac13615d7138e66cd726"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_solaris_amd64.zip",
     "sha256_checksum": "6e65a7d16965636d59dae1595d79105e0139ca1b3e741aea1142710e494fc9e1"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_windows_386.zip",
     "sha256_checksum": "143d8a309708973c5011024103cfdc52533ec76ecad2bd05777758018bf191"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-beta1/terraform_1.2.0-beta1_windows_amd64.zip",
     "sha256_checksum": "4749d11ce67b97c88ccdc602b9a510f4bca90540c22fd8a526f9a6466a8509f1"
    }
   ]
  },
  {
   "version": "1.1.0-beta2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_darwin_amd64.zip",
     "sha256_checksum": "851347d7777c4ed874ff7a321a7cbd8478bd323edf204783b35d00f077a8c826"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_darwin_arm64.zip",
     "sha256_checksum": "7b0a670952ea91d507169be4e6b34b0efa318f8d8201b6eb4d81e48aad423452"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_freebsd_386.zip",
     "sha256_checksum": "1738946cd5a2fc3e11899e5c8dbc2a7f9b238fd473452d172d641cf151e0bda4"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_freebsd_amd64.zip",
     "sha256_checksum": "9bb2a3527a94dc0e4e3007fc0227f60f8b1bdf3cd31d1a4eb261ec6d8784442e"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_freebsd_arm.zip",
     "sha256_checksum": "ccf9dcd40e0c94102eb3fbd2ba7720d09e57a2965c3b3115d272a30b2d2208b6"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_linux_386.zip",
     "sha256_checksum": "3376ab93a0e98045401014e8151b22ed2219a77a4c8a00dc869fa377aceca090"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_linux_amd64.zip",
     "sha256_checksum": "65e5135a00911f7f5c199ba506d60f531df05277fe0c336b83298667c16bff4"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_linux_arm.zip",
     "sha256_checksum": "b83e9514ecf0831adbbca1238e7821c3d699f7f7444fbe17288ebfa4cbaf9f9f"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_linux_arm64.zip",
     "sha256_checksum": "69155e5b28f796e634dbf3d6f6b965aa697ca16dc12193c79775be1c5568ed2b"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_openbsd_386.zip",
     "sha256_checksum": "f64b348de084ebc65fc6e3b925a4da650718b39061c710019578a345f6adef2c"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_openbsd_amd64.zip",
     "sha256_checksum": "51e38f3e75af999ded36521a66ab311eb5c809cff52a30c5e631c850cb7c29fc"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_solaris_amd64.zip",
     "sha256_checksum": "53415dc5dbfd525d37b0a2bfcf2dd54a8ba7f1e42a6112a9c284c105509b8a99"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_windows_386.zip",
     "sha256_checksum": "8bb55f5f682a7b2a9abc65ad98ac9a84a4c8eb3f076cc89b4930565ed77026a6"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-beta2/terraform_1.1.0-beta2_windows_amd64.zip",
     "sha256_checksum": "8af1f49028d871af0c5a96547711eed6918c71166ff2fc6ce13402b44c1d8715"
    }
   ]
  },
  {
   "version": "0.11.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_darwin_amd64.zip",
     "sha256_checksum": "af78baf9b1a249544cc0b17d6b7abb32cc513a554d1f7dcc85c873e2af93586"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_freebsd_386.zip",
     "sha256_checksum": "afadd7985db3246de75267a058e4e1e108723b5ff43d99d50496fae4167508f9"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_freebsd_amd64.zip",
     "sha256_checksum": "e334d36ccd100183e99eb09dc8c9bf760f2a24e715bd9e20f45e26eb780372fd"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_freebsd_arm.zip",
     "sha256_checksum": "fe514032a7a44a8ed55274838de51af9d9a136700a3a1fa642ca48962909ef0d"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_linux_386.zip",
     "sha256_checksum": "531722fae5e88613dd897660a558e8dedaf3409145d87853696392406a3b9628"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_linux_amd64.zip",
     "sha256_checksum": "131c440263382c79c7f783b70ff35cd1d03eb31c44f7738d153d95a0b8436ac9"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_linux_arm.zip",
     "sha256_checksum": "ec4188d32a8bf0d95616d5290c33c1b4528a13c91916784eb1eb80bae830ec4e"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_openbsd_386.zip",
     "sha256_checksum": "b791ffb32346163d92d89626fe09845e0f7b945e84579093dda59a34b1ceef2a"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_openbsd_amd64.zip",
     "sha256_checksum": "670bde1382583b83093a7be1c43d189a67ef2150f5dc1580f0e23535dccafbc8"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_solaris_amd64.zip",
     "sha256_checksum": "e7d0799a778d2afbbf54df6283a262a184c6935e55e0676636a8ddf6fa4c0850"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_windows_386.zip",
     "sha256_checksum": "47f6b634c78c27ae0f765817cc80a4411f91bdf7f4e593f8d83abcb58290a710"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_windows_amd64.zip",
     "sha256_checksum": "c16be99bae3b252875ff58749adc7d4d32e25d2fe3d858d0090018f3b82d7aa"
    }
   ]
  },
  {
   "version": "0.8.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_darwin_amd64.zip",
     "sha256_checksum": "6bec1c06dbeb89ea7fdc2036be972372aa6847d3883786ab285386750a7ceb6"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_freebsd_386.zip",
     "sha256_checksum": "57bae3236347730d70bb599a80b63b00090c079b7106ee0b2240645c629a5a82"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_freebsd_amd64.zip",
     "sha256_checksum": "f550471ed699cef7fcf39d5f33324b095c673612862d22503015b5f2634eedb3"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_freebsd_arm.zip",
     "sha256_checksum": "594a77eb01f00af6a305d421070592be2829f19cc7addc715ec6ffe611a8875"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_linux_386.zip",
     "sha256_checksum": "ac64cc040139b912fd0271871d1bea068033f44ee67838751b7e31f96439a92f"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_linux_amd64.zip",
     "sha256_checksum": "a366fd2d7d8908d23acc23ab151fc692615a147f8832971bb43e42995554c652"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_linux_arm.zip",
     "sha256_checksum": "b8a60d61d2c1deea3c9b6a363b06d8494121bd1f475b7aa9003dffa575b3e528"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_openbsd_386.zip",
     "sha256_checksum": "7054a7e19699177124981aa277c1c7d3df9696b172698efad4a3b5ec66189e7d"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_openbsd_amd64.zip",
     "sha256_checksum": "e7da5d00ecabf945b9c2613aeb8036b7d2e2a9740ffc61862402cd72cad3145e"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_solaris_amd64.zip",
     "sha256_checksum": "3aea9bd9e75f7a47b734b81ff213bb3377a6d62328b02e338c6dc427fc811ece"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_windows_386.zip",
     "sha256_checksum": "dbc605a20d1d59cf4184e23d641cc59523aec8908e3356e032044f9777246f03"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_windows_amd64.zip",
     "sha256_checksum": "e1166534f77c89f742a5ae2549f2b9dfe2491bd36ba771ed2a88d5e16508e7e2"
    }
   ]
  },
  {
   "version": "0.12.23",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_darwin_amd64.zip",
     "sha256_checksum": "ca1a0bc58b4e482d0bdcaee95d002f4901094935fd4b184f57563a5c34fd18d9"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_freebsd_386.zip",
     "sha256_checksum": "d3160c5eefed2aa8f9990435855a5ecbe656f710f4128e62cff098797cc2b467"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_freebsd_amd64.zip",
     "sha256_checksum": "47428cd8d019a90c4adaa64ccf5651164281158b93c41da5bc658ae46033a826"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_freebsd_arm.zip",
     "sha256_checksum": "e35804f155e1abb1adf79dcafd672a6afc77fba83345f0f41e8f0374e260fe6f"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_linux_386.zip",
     "sha256_checksum": "204e03b6c5665fb3d1050f402a567dfad17250679363561add9e7dd8cc6e49d1"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_linux_amd64.zip",
     "sha256_checksum": "78fd53c0fffd657ee0ab5decac604b0dea2e6c0d4199a9f27db53f081d831a45"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_linux_arm.zip",
     "sha256_checksum": "ddfe78706ad66d7c39f03086613f2fcf80a282130c49b180496ccef5bfa3b8be"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_openbsd_386.zip",
     "sha256_checksum": "18eaab0b56acce2ab1e0510d3209ff61832c851d1a1f2e93e66d60cbcd3f4f83"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_openbsd_amd64.zip",
     "sha256_checksum": "f22b0140c7b6e8c46e729b9470fd289a0cdd0b16bf2a9a7557915f250fb38fd6"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_solaris_amd64.zip",
     "sha256_checksum": "59aad2d8e65e45991d646005ebab7c70623e20353051ed28b33cf62c1937708e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_windows_386.zip",
     "sha256_checksum": "6c9484b7839268620f15ee1694a86cc9067a3e4b6bcf9dc0eb80056cd575aab2"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_windows_amd64.zip",
     "sha256_checksum": "79f89c2622a3ec656f13e92326d1ad98893659eada8f556957e8e0f2e4d688d2"
    }
   ]
  },
  {
   "version": "0.14.0-alpha20200923",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_darwin_amd64.zip",
     "sha256_checksum": "7d7e888fdd28abfe00894f9055209b9eec785153641de98e6852aa071008d4ee"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_freebsd_386.zip",
     "sha256_checksum": "f8b6cf9ade087c17826d49d89cef21261cdc22bd27065bbc5b27d7dbf7fbbf6c"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_freebsd_amd64.zip",
     "sha256_checksum": "a5ba9945606bb7bfb821ba303957eeb40dd9ee4e706ba8da1eaf7cbeb0356e63"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_freebsd_arm.zip",
     "sha256_checksum": "df3a5a8d6ffff7bacf19c92d10d0d500f98169ea17b3764b01a789f563d1aad7"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_linux_386.zip",
     "sha256_checksum": "86119a26576d06b8281a97e8644380da89ce16197cd955f74ea5ee664e9358b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_linux_amd64.zip",
     "sha256_checksum": "1e5f7a5f3ade7b8b1d1d59c5cea2e1a2f8d2f8c3f41962dbbe8647e222be8239"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_linux_arm.zip",
     "sha256_checksum": "e9fd0f3e2254b526a0e81e0cfdfc82583b0cd343778c53ead21aa7d52f776d7"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_openbsd_386.zip",
     "sha256_checksum": "66a947e7de1c74caf9f584c3ed4e91d2cb1af6fe5ce8abaf1cf8f7ff626a09d1"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_openbsd_amd64.zip",
     "sha256_checksum": "def1b73849bec0dc57a04405847921bf9206c75b52ae9de195476facb26bd85e"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_solaris_amd64.zip",
     "sha256_checksum": "48f1826ec31d6f104e46cc2022b41f30cd1019ef48eaec9697654ef9ec37a879"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_windows_386.zip",
     "sha256_checksum": "17e0b496022bc4e4137be15e96d2b051c8acd6e14cb48d9b13b262330464f6cc"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_windows_amd64.zip",
     "sha256_checksum": "2696c86228f491bc5425561c45904c9ce39b1c676b1e17734cb2ee6b578c4bcd"
    }
   ]
  },
  {
   "version": "0.10.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_darwin_amd64.zip",
     "sha256_checksum": "70885c572f7bc54361c77d4839303210579db5875636711f621f6763574c1237"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_freebsd_386.zip",
     "sha256_checksum": "49679e1adc9940de0be33df93c5f33fe02f52ded411f12f28d487bfe6138b992"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_freebsd_amd64.zip",
     "sha256_checksum": "fdaf5b70c08e12a3b461428f6f77152d07416db38994118dc24a83a1ff053e3c"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_freebsd_arm.zip",
     "sha256_checksum": "80362bcd51d0eab55301276e57ca99907470634c3be81b1d8ee4fd37a6eb5c55"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_linux_386.zip",
     "sha256_checksum": "fab709d53c01c2a5df497ad4274d2322d314d6dfae9f5d47c607c9a6b49f785b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_linux_amd64.zip",
     "sha256_checksum": "cff83f669d0e4ac315e792a57659d5aae8ea1fcfdca6931c7cc4679b4e6c60e3"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_linux_arm.zip",
     "sha256_checksum": "73f2ba6a07ff837d0826831d25b310cd0a36f15ac63e48115024f7a00476766a"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_openbsd_386.zip",
     "sha256_checksum": "755e1344d0e2f249a569d2541eb3af21dbc421f0d3b065dfa530cb32ff810ff8"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_openbsd_amd64.zip",
     "sha256_checksum": "30eb745b90434ab1279dc950b6c186723715b47f1d96a05e233325534db67c79"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_solaris_amd64.zip",
     "sha256_checksum": "90bbfc8f945793ea29d06b7a13346e0350f30e648f3d80edda02cafde10a1e1e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_windows_386.zip",
     "sha256_checksum": "ad440091b09f27029e49e46c312be4ef1e102c75ce4c34f0062fef136fe0c4fd"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_windows_amd64.zip",
     "sha256_checksum": "27e40fa8d944b8e1c68b1968a8890625a3ef18e4e752782eaa4dc6e99c6ded3a"
    }
   ]
  },
  {
   "version": "0.12.0-alpha1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_darwin_amd64.zip",
     "sha256_checksum": "2797b82e22c5557da604b6b727cb8112844a92c81b16840980a43ed78d9e0512"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_freebsd_386.zip",
     "sha256_checksum": "7c305b9087f6d2e263c48be2b76af5eed18edc75135dcedbe44ff024f8ca9142"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_freebsd_amd64.zip",
     "sha256_checksum": "7f0bba655d113d1f1fd2ece71ae7c76e6e13056c5e2c99bfe039382dca01bbc2"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_freebsd_arm.zip",
     "sha256_checksum": "36c9e4e7c3973bb509090c10e64f3792ba0589531d16bf9aefc653d6eae1a47e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_linux_386.zip",
     "sha256_checksum": "19f2c3b287fd820abcc6c590dcba67772d74051d969f9a6ac0d52d1d42164930"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_linux_amd64.zip",
     "sha256_checksum": "75251f20455829358b38e28d5f3b726cd544f606e85c8161a34dea3bf924a9da"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_linux_arm.zip",
     "sha256_checksum": "728bcaa5feadd62771ef1f105bbca0bea3ee891edade8b34707949c7bd6dd731"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_openbsd_386.zip",
     "sha256_checksum": "2a033576c7e1642395a317870181d79829c5b6892d524a0e3f56a38149f33211"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_openbsd_amd64.zip",
     "sha256_checksum": "12e4e99ccd338ab0e7f8e10ce658d669d15b7cebb9859c326480387dcf7ea826"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_solaris_amd64.zip",
     "sha256_checksum": "7a3705820ca4ef64e66ab5583e1df82097ed747bdd6fd80c0f7a86f5add76f41"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_windows_386.zip",
     "sha256_checksum": "a2a50de9ec4b8fe0ef4a9143007a885ab612703934f11f3fe051fb8832220591"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_windows_amd64.zip",
     "sha256_checksum": "19da49c0983325c91c353e48bdb0beb95ab9d5f35117f47b001648351dd0a01c"
    }
   ]
  },
  {
   "version": "0.9.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_darwin_amd64.zip",
     "sha256_checksum": "ece7ad727eac202b571c64018ec3d09b4d7693aea7033db81e239d96d11d48b9"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_freebsd_386.zip",
     "sha256_checksum": "aba1620cba744838589d308478a5dadb84ecd54cd4e730035182f71ad0e18b6d"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_freebsd_amd64.zip",
     "sha256_checksum": "3f49f7cb1b845ae97662ea6535ee97e70cd8096c5c98d7b4bb161f79e75f32c6"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_freebsd_arm.zip",
     "sha256_checksum": "dcd5f8454f6e63f2aa5dcfb9da4a52737689040fd4e56170be0b6ecb8fa87606"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_linux_386.zip",
     "sha256_checksum": "4adddba02c54df7dfad760b5e08eeca9756f49c0df2d86bf157c47e85b9f4ed0"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_linux_amd64.zip",
     "sha256_checksum": "df6aecbddca9ea4e0f32290ed5efb92c1b4173d8afd1d95470424f635c41d242"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_linux_arm.zip",
     "sha256_checksum": "53f7bbde44262a32765731553a61dc8e103d49a036c85eacc3f2bdbe9ecf7777"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_openbsd_386.zip",
     "sha256_checksum": "9282ae091706faaa1dbd1f108debb44d2cc8b73a2d5c7c74765a5f457074dc4e"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_openbsd_amd64.zip",
     "sha256_checksum": "3e0b76e5162d38b590e51a65f8491f27ed31f9d7f93d75a8d8ae8523317e28d6"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_solaris_amd64.zip",
     "sha256_checksum": "dbd7b97dbf5612341100b97d3615146cde31a6b7874b5b95bd51b42e765a7db1"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_windows_386.zip",
     "sha256_checksum": "5a1e7dfdb84ecc29afcaf2b79f7e284d106a1f568b571015e59ad3924d674bfd"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_windows_amd64.zip",
     "sha256_checksum": "7054e17eb5436d7e456f7f9368a483723a054507827b21a9938df9d8ac199934"
    }
   ]
  },
  {
   "version": "1.2.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_darwin_amd64.zip",
     "sha256_checksum": "94d1efad05a06c879b9c1afc8a6f7acb2532d33864225605fc766ecdd58d9888"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_darwin_arm64.zip",
     "sha256_checksum": "452675f91cfe955a95708697a739d9b114c39ff566da7d9b31489064ceaaf66a"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_freebsd_386.zip",
     "sha256_checksum": "1bedf7564838493f7cd9cb72544996c27dcfbbae9bf5436ef334e865515e6f24"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_freebsd_amd64.zip",
     "sha256_checksum": "353b21367e5eb9804cfba3140e786c5c149c10098b2a54aa5be3ec30c8425be0"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_freebsd_arm.zip",
     "sha256_checksum": "47aa169b52c4b566f37d9f39f41cfc34ee2e4152641a9109c2767f48007b2457"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_linux_386.zip",
     "sha256_checksum": "3d6c0dc8836dbfcfc82e6ba69891f21bfad6a09116e6ddf7a14187b8ee0acce5"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_linux_amd64.zip",
     "sha256_checksum": "9fd445e7a191317dcfc99d012ab632f2cc01f12af14a44dfbaba82e0f9680365"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_linux_arm.zip",
     "sha256_checksum": "ed49a5422ca51cbc90472a754979f9bbba5f0c39f6a0abe570e525bbae4e6540"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_linux_arm64.zip",
     "sha256_checksum": "322755d11f0da11169cdb234af74ada5599046c698dccc125859505f85da2a20"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_openbsd_386.zip",
     "sha256_checksum": "426d39f1b87bf5dbda3ebb4585483288dba09c36731d5cae146f29df0119036c"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_openbsd_amd64.zip",
     "sha256_checksum": "5b0c59ffe5f83363b20f74df428490b95ff81f53348f8c8394519768085f3eef"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_solaris_amd64.zip",
     "sha256_checksum": "64e70edf5af0e77f54d111ae318282aebcdaa33e8dd545b93881fd421dc4d982"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_windows_386.zip",
     "sha256_checksum": "f26acca0060c42c0e6fb81d268fbf4ab9baac3d5f34c8263ecdb48c0a78f905b"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.6/terraform_1.2.6_windows_amd64.zip",
     "sha256_checksum": "1e3c884cf32879646f97b8b6a253686710eb6e445d44097580a27511a49db88b"
    }
   ]
  },
  {
   "version": "0.7.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_darwin_386.zip",
     "sha256_checksum": "5ddc51310093ea9de3511fbe89a72ab9c55cff1119e6825fa0bad4268972b17b"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_darwin_amd64.zip",
     "sha256_checksum": "2a441124efd097007414545714927a9239980a5b0707384b0ee07badbae781cf"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_freebsd_386.zip",
     "sha256_checksum": "9c761a35e7af2a6b80d2b11a8734970d5a8671620367c7655fca185c102cfb27"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_freebsd_amd64.zip",
     "sha256_checksum": "479bebf320fa50b393f49fd4b7d7625782c6c08252f81754aa540e723e7e51c4"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_freebsd_arm.zip",
     "sha256_checksum": "cc2f618cdbb84c0faafba8763711a8102db50f7777d725476886df4563a3cb7"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_linux_386.zip",
     "sha256_checksum": "71d50efba9b7007f00b044fe83306b567a081ceeb28e125d1edeaec290f30adc"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_linux_amd64.zip",
     "sha256_checksum": "b337c885526a8a653075551ac5363a09925ce9cf141f4e9a0d9f497842c85ad5"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_linux_arm.zip",
     "sha256_checksum": "f0ce914530844c50188423a949a9d3e61ce7ac145b960c7f1a7dfbcd97495d43"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_openbsd_386.zip",
     "sha256_checksum": "987f8e99bacc360e5868cfecc5dab34e38963788c3516c9e1d4640521ba22815"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_openbsd_amd64.zip",
     "sha256_checksum": "f133554d8e282d73561b767c93a985322af92fe8e960816edc8e50ddd18b0015"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_solaris_amd64.zip",
     "sha256_checksum": "51f98ea03e758dcc3d95eff5a23ddb72b04a599fbc832b4bb575135ec9a14c1d"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_windows_386.zip",
     "sha256_checksum": "77d6342114fef22d3703bf3c32424b35bf70cc34d71172af2f847fb218fa2f82"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_windows_amd64.zip",
     "sha256_checksum": "e52251f4c7ef255bfa42697f9d2d29fdd6cf79ce374df02119143299137a3c97"
    }
   ]
  },
  {
   "version": "0.11.9-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_darwin_amd64.zip",
     "sha256_checksum": "a95ac475acd068a876a1068fa90cb2e9370e1c28e8c7fc57b7db016629b533be"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_freebsd_386.zip",
     "sha256_checksum": "1d25771a8fa9623b8c5b443de2f39385364a875b46638385fe8d96ae892e8353"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_freebsd_amd64.zip",
     "sha256_checksum": "475a5bb5e38855e26f151c59971d8c106526ca7b7935157b35336d0c320f82f6"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_freebsd_arm.zip",
     "sha256_checksum": "15394512474ba06781faa4cb5acfb2ea76d2793bfb3064dcef7d2e8a9793bb92"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_linux_386.zip",
     "sha256_checksum": "f2e7970787cb833ec37a238260cc822c5cc5442bf35618eea0e45f677721189a"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_linux_amd64.zip",
     "sha256_checksum": "6f235d05a3e0b3cd79264834448064b52079c4268ed0c2de25e10c79386f746e"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_linux_arm.zip",
     "sha256_checksum": "7862e932c5c32b767be470312c9e610fec8262b623c4bf56f18c3b1614d66eec"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_openbsd_386.zip",
     "sha256_checksum": "e4bc9b67c4b15a11ab02b285a679d8994dad5e03ab5e715d61ac4cfe6fdacca4"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_openbsd_amd64.zip",
     "sha256_checksum": "e1fa53283905878a6b276fb2a6349d2d93bdcc2e6d860ffd2bc812a0beb41670"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_solaris_amd64.zip",
     "sha256_checksum": "b8914bd82409de3be22ebc611bcc4002623adc8402dc7f5db610de177d2c58f4"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_windows_386.zip",
     "sha256_checksum": "ff905570254350188f6b9bbe2e86689d9105c30ff65f00d6d94f2c9c45dfbd61"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_windows_amd64.zip",
     "sha256_checksum": "4d0ee636b7b445f6fd3acb6f40c94530641beb2f06b292312a61881fc31c1479"
    }
   ]
  },
  {
   "version": "0.4.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_darwin_386.zip",
     "sha256_checksum": "8ce42f0b712b58e2d619ae6625f5cebad363afa8b237ca617c55e687379e2ace"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_darwin_amd64.zip",
     "sha256_checksum": "8bb2eaa5b4eae89963e5ed1598689d95d220c0cafb59bbd5f266f8e326ac944"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_freebsd_386.zip",
     "sha256_checksum": "b411c01e7efced45f8b632d0ed4fb96df2ecb6f39e4a4121719752376a8b54ba"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_freebsd_amd64.zip",
     "sha256_checksum": "e3d531415352eb8f8c082e0c24e54f936b29cd99de97215a297f7da58a4a1511"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_freebsd_arm.zip",
     "sha256_checksum": "beef40413aadd9890b41217a8a8a934c0900b78c8e249265d9c8c5a4668cb9b0"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_linux_386.zip",
     "sha256_checksum": "46c4aba13f396baab90945eee2f83ae3c1c74f7c83bce97311f0ad514acfeb8e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_linux_amd64.zip",
     "sha256_checksum": "85ed7c905ce76a4de166acf7aa400bfe48894a97244d9f927804b226dd373f7a"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_linux_arm.zip",
     "sha256_checksum": "e493e413cfc92b6b2c13a1ec69104be0c9918872f7eb5323730af9135b0b9963"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_openbsd_386.zip",
     "sha256_checksum": "892d2a41e8379d9cae7fc7edc1ac0d15e9148f4812b97dba022cfa0615640e62"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_openbsd_amd64.zip",
     "sha256_checksum": "449bb8e28bb358841267d20235eea035365e66404f001758fd6608b71a881d71"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_windows_386.zip",
     "sha256_checksum": "c0081f9ebbd49a44b1746b05b81ec3dc43667d499bfc27a241c404e4a5f80db6"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_windows_amd64.zip",
     "sha256_checksum": "a8441e9e77e4710520b70c93adccd7ae8b279be5f869959968a7fa182cb96a56"
    }
   ]
  },
  {
   "version": "0.12.18",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_darwin_amd64.zip",
     "sha256_checksum": "dd6983cfe0d0922de51e019a80db2a270e8a2636b9f05b49bf7dcbfe7bd90ea9"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_freebsd_386.zip",
     "sha256_checksum": "7dc3d0566755ba168cb26fcc2fe64c06447aef3af320b810470c67701c214a13"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_freebsd_amd64.zip",
     "sha256_checksum": "5fca08e6a64ed4906e4b503c25467d98dbdd9547dd5fabbb45ab42bcdc083889"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_freebsd_arm.zip",
     "sha256_checksum": "9458de071f218e4c2db418d6c09fa19cee39245fbed92ba2fc1c7e8c842c1763"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_linux_386.zip",
     "sha256_checksum": "f35eff98a6160979d6688ded7c251c18c69dd7884d493122e6e00cd77c6a4efb"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_linux_amd64.zip",
     "sha256_checksum": "e7ebe3ca988768bffe0c239d107645bd53515354cff6cbe5651718195a151700"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_linux_arm.zip",
     "sha256_checksum": "ab1d5709d6c15913b30f04f77ddf3c3a8b76e7c452eb0e9054ab26e5a6caab81"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_openbsd_386.zip",
     "sha256_checksum": "499699165abef9cc1fa769d48fadecf762dd651cb7683165c2a533658fa4d138"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_openbsd_amd64.zip",
     "sha256_checksum": "2b1ec25dd716a118c8831f470fec2868bef908d260417ca5e4eb7295b5db02c0"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_solaris_amd64.zip",
     "sha256_checksum": "36d5b028f507e389be12b70543e53e27e550e5c40e75b8f132fe5f1ff5065639"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_windows_386.zip",
     "sha256_checksum": "473eaf44c2589ac04641918d87abab758cb8d8a6e6f9372c8a2649bc5fa393f0"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_windows_amd64.zip",
     "sha256_checksum": "c009292722e8ede525cb71f978cf481753496f5466f23373fc728d8370ab26b1"
    }
   ]
  },
  {
   "version": "1.0.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_darwin_amd64.zip",
     "sha256_checksum": "32c5b3123bc7a4284131dbcabd829c6e72f7cc4df7a83d6e725eb97905099317"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_freebsd_386.zip",
     "sha256_checksum": "21099aa90427c2908606d39b8a517e18db67ee9f45340cd7f0a41a8f27bbdd52"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_freebsd_amd64.zip",
     "sha256_checksum": "68d74b5a9fea3b548d697b9c8c414bc5af602a90a222ceeccca097fa1f8d0536"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_freebsd_arm.zip",
     "sha256_checksum": "7df585093402b7b3b820e190b0e919cf05ba180b1ec9e8f6443f7e1b5aac41fd"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_linux_386.zip",
     "sha256_checksum": "ffb5ee921984bc3c97d948d6de0b26cad21ee40c871fb91cf8021054d0c1849d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_linux_amd64.zip",
     "sha256_checksum": "da94657593636c8d35a96e4041136435ff58bb0061245b7d0f82db4a7728cef3"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_linux_arm.zip",
     "sha256_checksum": "4b95512b651837da3bef8da25c0621ba345f99a2418ee730691e27cb28d841fd"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_linux_arm64.zip",
     "sha256_checksum": "af9c579d59e5180e563a12a69f5acb95fcc61ec2a39ac24aaf13aee33463760c"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_openbsd_386.zip",
     "sha256_checksum": "e2a25e2a6553533cbaa6a0ad813f8318797afa36922744259db6c73c9c99e1d9"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_openbsd_amd64.zip",
     "sha256_checksum": "20753ea36722e5ccde103b90537ad17b1e862aa0e8e2e0542e6eae320e72a034"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_solaris_amd64.zip",
     "sha256_checksum": "405a9e51df62f49c88becbaee7489e2d36592e55a0ceb3a10f725da5ed337430"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_windows_386.zip",
     "sha256_checksum": "e81795f4cec024cced3626522c6be596b794999616d748267e0712086fc12349"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_windows_amd64.zip",
     "sha256_checksum": "df8b9f06b9c9edaa65841de6c6c38f60098df84bceb9db884d24d19dbe2e5fe1"
    }
   ]
  },
  {
   "version": "0.15.0-alpha20210127",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_darwin_amd64.zip",
     "sha256_checksum": "5053459451746b22df24fb42c4f767e247827c7bf1aa425e8636d5cec54ec28c"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_freebsd_386.zip",
     "sha256_checksum": "f057cf5b6ca5f8bb91d4478e6700561d79afc96f53c4a4e5d70a341b7133e9c6"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_freebsd_amd64.zip",
     "sha256_checksum": "15b254fec360b9bd9aad0cb8da75e8cc28c96a11dd530eba70db986f5539251f"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_freebsd_arm.zip",
     "sha256_checksum": "d31730fabe40e9535afb46f01acbc7623b27320d764c5bb7b65bf142fd3cd12a"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_linux_386.zip",
     "sha256_checksum": "99ae4870e8f4c90dd39822e2839670feea595f0c59925c47ec5b024387097b6f"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_linux_amd64.zip",
     "sha256_checksum": "1bb2dfa88fa066df7c9d3c58b96a9db6ecd3302db6bf63d693a4f39268e1e156"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_linux_arm.zip",
     "sha256_checksum": "ff58683c9aa6cfdbccccc80527f54b53979130b693efa54b1e1c14c9748f2fb"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_linux_arm64.zip",
     "sha256_checksum": "6dd63878d77f39cdd096792f37fe68648cc8ee1f626b2f18bfa62ce19172b99"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_openbsd_386.zip",
     "sha256_checksum": "79d67dc919b2296841aaba9bf955c2fce4d706b1762283bc19df8300618d37b8"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_openbsd_amd64.zip",
     "sha256_checksum": "33249089e85505beb2f3b8264716cdd8235d3ce5e716c4437a1ad4f69431bc1c"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_solaris_amd64.zip",
     "sha256_checksum": "bbbe905ab2d793037665e1302391bc147b3a5b2203eb7a16f0170efc179d849b"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_windows_386.zip",
     "sha256_checksum": "e9e9f7c632c43d6f6454edd7732816e8c3474d2b26f35d92eaca8305909328ba"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_windows_amd64.zip",
     "sha256_checksum": "2086fb096eade5d484ea1da0953c45b77eeccb9313aa76afe7983848cab71061"
    }
   ]
  },
  {
   "version": "1.0.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_darwin_amd64.zip",
     "sha256_checksum": "3a97f2fffb75ac47a320d1595e20947afc8324571a784f1bd50bd91e26d5648c"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_darwin_arm64.zip",
     "sha256_checksum": "aaff1eccaf4099da22fe3c6b662011f8295dad9c94a35e1557b92844610f91f3"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_freebsd_386.zip",
     "sha256_checksum": "ffa6094a06ceff98824e34ea6147f223c7f234747a6710498acf2d26a60cf91"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_freebsd_amd64.zip",
     "sha256_checksum": "94989fa7d8bf601992104c5c908ff98ea3250a47a8934346bf8886d814ddf45d"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_freebsd_arm.zip",
     "sha256_checksum": "6c483a304a473f10f8995c0b37c8facbdf8ddee87e26c09bbbb9a8ea89fb5ae8"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_linux_386.zip",
     "sha256_checksum": "d77373bc8ad49fcd366ba4f5ed4adf618ec0164daf6a7791259bc8f4b6c2bdbc"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_linux_amd64.zip",
     "sha256_checksum": "6a454323d252d34e928785a3b7c52bfaff1192f82685dfee4da1279bb700b733"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_linux_arm.zip",
     "sha256_checksum": "10c67773334e2f4ee6c39815327f2cc1df242d97df997666d730fbc2b97558db"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_linux_arm64.zip",
     "sha256_checksum": "2047f8afc7d0d7b645a0422181ba3fe47b3547c4fe658f95eebeb872752ec129"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_openbsd_386.zip",
     "sha256_checksum": "91643f790489d8222e9077dd8a83451f9ca3f8b0574c8afbadcffe1afd4528d7"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_openbsd_amd64.zip",
     "sha256_checksum": "662c916245e478b33157b62438a55ff059aea1d6958d1fe98911c58ef309b5a7"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_solaris_amd64.zip",
     "sha256_checksum": "eab79ac9217e93920c7f5ffdd65394de97fdbc43efcaece84cc414b7f8af258c"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_windows_386.zip",
     "sha256_checksum": "d5588112f7cd73cefcd9d7dff046e7e18eaaf6acba9601fd3363d9a7892cba85"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_windows_amd64.zip",
     "sha256_checksum": "4e292555e71491c859e82b6995548dd43cb43c1b826359f60ac624ccd1111de4"
    }
   ]
  },
  {
   "version": "1.1.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_darwin_amd64.zip",
     "sha256_checksum": "302aebb4e308d95bb6dc7307d109b7f4805d953ec0b331f7c2a643a59f2046bd"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_darwin_arm64.zip",
     "sha256_checksum": "4b94db55087bb167e8b50ba949fceb1bbe9933f1d743768eadf8e55dc86d7061"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_freebsd_386.zip",
     "sha256_checksum": "4858ac29764ad1c88362f272b0216f87dbf25cfb7bbf95e2e531680ff43c2b5c"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_freebsd_amd64.zip",
     "sha256_checksum": "e1b95ed17343572ab9d0e585e187018bd74510c0f412892ac91150e96ccb19b9"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_freebsd_arm.zip",
     "sha256_checksum": "749570b4ff4569896fb27ba8552de8efbc65db03483733e58b5b85299807a324"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_linux_386.zip",
     "sha256_checksum": "739fa1febf8b026aee7b151a73ad1321b4694bde99377b6969fc7bef5bb4a61e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_linux_amd64.zip",
     "sha256_checksum": "a78c3cba635b504ae2f261f5d112243a98ae6203f018af32385b4e2e1e673093"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_linux_arm.zip",
     "sha256_checksum": "2fcbd1c031236f19e3749917f6e97f96e9ba497cef47b31f511a6bd6837a3c8"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_linux_arm64.zip",
     "sha256_checksum": "498eace0d865e193f33c5dc5ca1b4f2b5abfd3d32fe97a6d63b2d1b236d65503"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_openbsd_386.zip",
     "sha256_checksum": "ee86099cd6027a891fa6fd497d93788646c5c0fc20863f180c0fb515fbc726e6"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_openbsd_amd64.zip",
     "sha256_checksum": "473dfc227d0d33c6393e45e0f04025889a0a508807bf3389fbbbb3fa0db33c8b"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_solaris_amd64.zip",
     "sha256_checksum": "26fa7638a7871faf56b626df5181b827df6935d7277a2ff681818d24ce1058a1"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_windows_386.zip",
     "sha256_checksum": "a7540c4851d4d6c328c6aa9ba1ac59288dfda46ef941907ce344d1e323f94d27"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-beta1/terraform_1.1.0-beta1_windows_amd64.zip",
     "sha256_checksum": "a91b46936f55268881942b2eee876c6c649f5e8846eaa536cbee2341a118d22d"
    }
   ]
  },
  {
   "version": "0.14.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_darwin_amd64.zip",
     "sha256_checksum": "10c5c0969438f973a25cffcae3e567697822459da3ff177d707f3ae2d5100962"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_freebsd_386.zip",
     "sha256_checksum": "7300b0fbfb39995347261b9742f70ced05e7dd25d0a06fa3d16875e21799fc89"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_freebsd_amd64.zip",
     "sha256_checksum": "23dd99e4de8ab5fdbf7fa09618a465438dcbf234f271898a90f38648c9227847"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_freebsd_arm.zip",
     "sha256_checksum": "906d627491f0cc9b239dd54e28fd7b47e86f934e78435a7e4d0c98ae5f7b1c07"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_linux_386.zip",
     "sha256_checksum": "166a7a243f257e30cf74a87c0f8101705e1f4fcab40102730daee529baae6e9e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_linux_amd64.zip",
     "sha256_checksum": "e0e07ed96d103fe8e9459f68a6f35bfb15ef82aac8fb5534764b105bc0ac4468"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_linux_arm.zip",
     "sha256_checksum": "544a5feea655f16ef02a70b1b7271653e1b29c2fb2d6d65212399dc5f43ca7f6"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_linux_arm64.zip",
     "sha256_checksum": "95e4b0a2a3de00c11223d3184d529f1259911843aa904f19871a8dd36dabf65c"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_openbsd_386.zip",
     "sha256_checksum": "5ca7374563b343d80373d55afaded63b0b6c1e81998a6bc77cbafe71a9e6e13c"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_openbsd_amd64.zip",
     "sha256_checksum": "90c146568598acef232d4ccf7a0e1a8a72e24cdf9ff3b9a192b665d03c526bb6"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_solaris_amd64.zip",
     "sha256_checksum": "704871593ea8d86a1ed302727e347577374d7b9dce85c8445fbe86d1f0733437"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_windows_386.zip",
     "sha256_checksum": "81e3fc40e21fa4892f6396aebc5251fb49881babdf013a897eeb2c21b726fddc"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_windows_amd64.zip",
     "sha256_checksum": "5c0e1b5fc8acf244fc4754461483768aa00ec12d02c0dfeb821accda3a6b0c1b"
    }
   ]
  },
  {
   "version": "0.8.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_darwin_amd64.zip",
     "sha256_checksum": "84ecdd2adf61629a6bd4c1316df8f76290afad689630225d415666b422214a83"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_freebsd_386.zip",
     "sha256_checksum": "c42545cbb1a7da22d06516005de1b7940085df66fa278342f44d2d372f0df608"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_freebsd_amd64.zip",
     "sha256_checksum": "bc166413a000b30f060fd7f6cc07e0f2795cf7745d508009719b68782175e81"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_freebsd_arm.zip",
     "sha256_checksum": "c46c5a6c447264c45e53466c7e5acdf824c68e9cfac316e61e0f30d4eeddbcb5"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_linux_386.zip",
     "sha256_checksum": "5e907965f030408a0aff69b1bac7e9a74290618f30e5ef96d53f22b5315f2035"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_linux_amd64.zip",
     "sha256_checksum": "aa4c64b0595b8e49fc12da79d3efb7fb3f3653230349819699b9db1a76c51f3b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_linux_arm.zip",
     "sha256_checksum": "e4952e8c80322cb7c05f7628be0e9858bf93088b650b468ed44fc2cb742d6b35"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_openbsd_386.zip",
     "sha256_checksum": "97d5324bc0f450ec6630df8ee02ffd2b1db54c26ab7da651f8b9af88cb766d6d"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_openbsd_amd64.zip",
     "sha256_checksum": "d040f275847b571ccec55324adb374a51eeef8fad5573940f931ed02b68abd1b"
    }
   ]
  },
  {
   "version": "0.12.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_darwin_amd64.zip",
     "sha256_checksum": "e1ddcd5f40d3e9b2758d8bc4858117f5df94169fec16495dada96d3ab358ff34"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_freebsd_386.zip",
     "sha256_checksum": "2d483a2b227bcfce756a6dfea21eb6eb71b72e3288123c8f37f9adc0bbb80dbb"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_freebsd_amd64.zip",
     "sha256_checksum": "70058cd7e87ff49108d7f8ba9aeda2183b02450e3abefc6368ca1c33fd6d841f"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_freebsd_arm.zip",
     "sha256_checksum": "ca5d942f51d0bfc713e49188a3f20640cdf587bc2ca9cdf8b0fc23e161e50aaf"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_linux_386.zip",
     "sha256_checksum": "40698e3178f93dc24f8f60e186eef338c7af897c6d6b946f210896d9a96f355a"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_linux_amd64.zip",
     "sha256_checksum": "d61f8758a25bc079bb0833b81f998fbc4cf03bb0f41b995e08204cf5978f700e"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_linux_arm.zip",
     "sha256_checksum": "38eac5f44190ebcf770db6e03e46c9edef39fae776137f216af18904a12d61b0"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_openbsd_386.zip",
     "sha256_checksum": "d64d22550d9abf3de0272c1d664b1b2e6a523e93953bd2687fa0575a8a870e2b"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_openbsd_amd64.zip",
     "sha256_checksum": "30bed8a3c52e62ef8e750891e4e8378327632c0102c7bf9290d8f968cb82372d"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_solaris_amd64.zip",
     "sha256_checksum": "16b881f77d85270e2ff3c1ede9e1a0d816f02dab0bbb03a26dc7bb24e60725a0"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_windows_386.zip",
     "sha256_checksum": "9335858ba7a02dc06377b9f4164e9eace21972b5580b1b4a7c88f6f2077f2dc3"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_windows_amd64.zip",
     "sha256_checksum": "7fdcf41741df594f745d0eb271ff92fb74323789aa7ed5f92d5bb924cf1b14b6"
    }
   ]
  },
  {
   "version": "1.0.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_darwin_amd64.zip",
     "sha256_checksum": "e2493c7ae12597d4a1e6437f6805b0a8bcaf01fc4e991d1f52f2773af3317342"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_darwin_arm64.zip",
     "sha256_checksum": "9f0e1366484748ecbd87c8ef69cc4d3d79296b0e2c1a108bcbbff985dbb92de8"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_freebsd_386.zip",
     "sha256_checksum": "e93cea64d708ac725f4f032a3a2ef9a4d69b453815f97c1daaf03a81ad11c1de"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_freebsd_amd64.zip",
     "sha256_checksum": "3edb0d8dcf913823e101c018b81e816e45dcc7f49fed616db232e1c9abf24cde"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_freebsd_arm.zip",
     "sha256_checksum": "3ad4e84db39c3b1c2f2e4b09ffe2cd9224cc30fdefc49f2d2da1db08c579974c"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_linux_386.zip",
     "sha256_checksum": "7760107b10ac5d6feb7af8a7f078aea82729de607336acbda299a8efc18a213c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_linux_amd64.zip",
     "sha256_checksum": "a73459d406067ce40a46f026dce610740d368c3b4a3d96591b10c7a577984c2e"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_linux_arm.zip",
     "sha256_checksum": "627f677397413d46e65f5f15e59af272ce1140abe34c8f193dd3ba7221c39b62"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_linux_arm64.zip",
     "sha256_checksum": "1aaef769f4791f9b28530e750aadbc983a8eabd0d55909e26392b333a1a26e4"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_openbsd_386.zip",
     "sha256_checksum": "e16cbe87d23b6f17317b5bd29dc0971dd0836d9c1da43b44b99f070f4c426ee9"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_openbsd_amd64.zip",
     "sha256_checksum": "df94cdf1760f4180653c386b17e5c27b53f375c1f0269da35494e27eb3dc5cc4"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_solaris_amd64.zip",
     "sha256_checksum": "7f31a1e20568b00dad64f6d6a7eb862321447f2a803708b33c0175563cd85479"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_windows_386.zip",
     "sha256_checksum": "cf2428950271cc8fcc460f492b8764c78b59661b89c62b449e974b91f5c5260f"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.8/terraform_1.0.8_windows_amd64.zip",
     "sha256_checksum": "4ba5058d051f15b48dc7a86d65f2b76e3fc637a0c98f321b01204df437ad07d8"
    }
   ]
  },
  {
   "version": "0.8.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_darwin_amd64.zip",
     "sha256_checksum": "79e94dfaf439fdbba2a2fe03dd7c90b24efa699b6661155aa9329df43e68ba51"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_freebsd_386.zip",
     "sha256_checksum": "5c655f8531ebcfe364329cb604a77cb5501df71a7d83b0789c406481a90c54e5"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_freebsd_amd64.zip",
     "sha256_checksum": "5703873b60cba7bc8d30ead9dbd9f40459f98089feec7d0d80e98aad91a96850"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_freebsd_arm.zip",
     "sha256_checksum": "e3557bdd04779a3589912b24dcf81be9c6240838b2c3624d1536771621ca70b9"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_linux_386.zip",
     "sha256_checksum": "e9c54d3ccc44479b5e971cc7c5435c3cdb3dc2e56d0999a09652d9e11db82432"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_linux_amd64.zip",
     "sha256_checksum": "297d35d0b4311445cd87ef032d3dec917bcc7a8b163ead28a4c45d966a2f75cc"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_linux_arm.zip",
     "sha256_checksum": "fc1b0f543e77ec9d609f9293d19219b6c341fdedb247da8ba199a64253ed14a9"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_openbsd_386.zip",
     "sha256_checksum": "4840f57b8e99d7cea13d7af2497f40a01bd3f03653d18463277c66160b86424d"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_openbsd_amd64.zip",
     "sha256_checksum": "36464984a5fdc90d9d3b1897fc2f68b4e4c30b4423edc4e4cfdb920508897ad8"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_solaris_amd64.zip",
     "sha256_checksum": "54238d0990b1e2fc7501af960242fb4e1fb8c093c7158fb08753e685f38ea1ab"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_windows_386.zip",
     "sha256_checksum": "ebb6119a9009bbb28d59177d6db4a737408755c6738e37021f031cbe1194e7f3"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_windows_amd64.zip",
     "sha256_checksum": "4e1f5ec48c943969a2989f2b8060cd2065cc6d7e1987c0d01875ecbd219b98a3"
    }
   ]
  },
  {
   "version": "1.2.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_darwin_amd64.zip",
     "sha256_checksum": "bd2c639edb2f607b82ee6c72b0d8d9e51c201440f6b87ec25a5fd6b564d3a9a9"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_darwin_arm64.zip",
     "sha256_checksum": "42c742deed49595dcf5a445892343462ff6292cb6bd70ce2bea9e1ee99c87cf9"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_freebsd_386.zip",
     "sha256_checksum": "7d9920dd01f4e3ff95241ff98735e01e8504c12eb1baaf0448533c2dddbe5237"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_freebsd_amd64.zip",
     "sha256_checksum": "5b10475cf7a7839d18d6b2ccc11f983b987ab0040c1d0c04c2ee81c3e0a8762e"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_freebsd_arm.zip",
     "sha256_checksum": "4b9101a539be89318604cb30c89b04ca45702547068ab63a936b3cea46496ba8"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_linux_386.zip",
     "sha256_checksum": "b1550765a1c8600ea49c6d50e5ab26335c0c63c9b567c9eb87562db6b74befa0"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_linux_amd64.zip",
     "sha256_checksum": "5bc19bf44c70c3c35b850b5db48a1bf3fac77ba08aa012519ed19b455632c74"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_linux_arm.zip",
     "sha256_checksum": "98df04dafff85db719fe2d1c6158ac9b166784ce733d56fc688480d53ec36383"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_linux_arm64.zip",
     "sha256_checksum": "28dbd713c7c11d098fc02331cf7b1205cada48c398eaa8f6d5f5bf286a2aae05"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_openbsd_386.zip",
     "sha256_checksum": "325cd742188cf9071de7e9e9a8d01575407e14c1e9aa175ce7e78410869d70cb"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_openbsd_amd64.zip",
     "sha256_checksum": "ad4122f07f76c4930660d80cfa85a4a277a5ce97100ee64401b629e5c0702cf"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_solaris_amd64.zip",
     "sha256_checksum": "9fc893a62b86763cf7d86b6ad28030335c202f69283f19d40bee79986f6b3058"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_windows_386.zip",
     "sha256_checksum": "9ae4ceb3f3beb952d8ad7f2fadd2d7ba34f93a8d83e93fa91b84db9d2a77497f"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc1/terraform_1.2.0-rc1_windows_amd64.zip",
     "sha256_checksum": "6d381c67b9a1a9a75a45cd82a2ba4345195eeade11b123ef9e9794931fbd2454"
    }
   ]
  },
  {
   "version": "0.4.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_darwin_386.zip",
     "sha256_checksum": "ada7760fc8985120d3c5d370805fe3f7f8de75b01a6850f61d0899f5d9109c0d"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_darwin_amd64.zip",
     "sha256_checksum": "eba9a10b11d572bc5146c1d01353193ba45af2683a0977db09e7b18dff079398"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_freebsd_386.zip",
     "sha256_checksum": "d208dfb5bc4762596d8254674393a855d711ebb2847708ca48f47544628fbab"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_freebsd_amd64.zip",
     "sha256_checksum": "19a1feeb5c228311cf1440861f54df9be508c6277b1eb9f484d43d89a9e8c790"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_freebsd_arm.zip",
     "sha256_checksum": "613e86461ddcac62475b8892013ae7f6de93c304f7b50e1915da0311842cc873"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_linux_386.zip",
     "sha256_checksum": "3507ca1da6b84084aae56867a9639d20f843062ba5ec7ed9619132eebb200633"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_linux_amd64.zip",
     "sha256_checksum": "57b23ac14dc37b97ff48fa64712e4b00e82e1b603c1d0d4789acf9dd4771f4de"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_linux_arm.zip",
     "sha256_checksum": "78943f4dc174d5d73982f9c48f4fc5d92c6bdf7be8e8578059bbb2dab986b997"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_openbsd_386.zip",
     "sha256_checksum": "a24b3116eb5373d3e58920ac429c6ede776b5f13ed2f3c10440c56687b6c724f"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_openbsd_amd64.zip",
     "sha256_checksum": "eef16d10e37a01770f3f5283bd1593eac5326afba4f9c058b5058d60ce5194dc"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_windows_386.zip",
     "sha256_checksum": "9708d6ba0ebfbec524d8cdee264cad0e997a806db555a5aaf35eb41184fb8521"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_windows_amd64.zip",
     "sha256_checksum": "ba25998053482f7d94f2c0f94f284a0c4431745120fae2273fc876686ed0e2dc"
    }
   ]
  },
  {
   "version": "0.6.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_darwin_386.zip",
     "sha256_checksum": "683207c5fb1285501baa556abc24786a80146d82ad74868fcd775da46aedecad"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_darwin_amd64.zip",
     "sha256_checksum": "9cf892c073a9fce0e9f136162f82c5b2d373c32cc2c5bd5c5eb16631262fad89"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_freebsd_386.zip",
     "sha256_checksum": "19daa5e64dc655647ebbd8f0b56f61a98655d3010e1b23898a998a5b7c756ff1"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_freebsd_amd64.zip",
     "sha256_checksum": "f1ee7701b1b48e689fe92ff93b060f9c36fc0e8bda4ebe88d042298a72d430b6"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_freebsd_arm.zip",
     "sha256_checksum": "68233eb41f2b559a5f6a6c0e633b738ce6911e48f2d299d9f80afaf8f3925cb"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_linux_386.zip",
     "sha256_checksum": "6d4236048c2051ee6198e76eec77a91698f37abca713e857ee3c4df40af2e197"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_linux_amd64.zip",
     "sha256_checksum": "c7d3e76de165be9f47eff8f54b41bb873f6f1881d2fb778a54bb8aaf69abfae6"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_linux_arm.zip",
     "sha256_checksum": "9c5549a46c67ccfb80b5bd2af2892d345826f2dc54529b6ab4c0c13a1e8ede72"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_openbsd_386.zip",
     "sha256_checksum": "a788c13418a4a1e8e870f7312208de2716bc76907db224ad0ae8ec43c897386e"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_openbsd_amd64.zip",
     "sha256_checksum": "a2855b2014f61570092bf0f26bc0ac3942dbde0963208e8034c1edab11319106"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_solaris_amd64.zip",
     "sha256_checksum": "2bcbea7f07b792e0652ca682f64be5c7f267150454cd8da1c22f59f8d615a2ac"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_windows_386.zip",
     "sha256_checksum": "6e4d56fbca03959120e8b84b73cf97ed3380b0a45696f9c4edc6f2c9a4ea2a36"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_windows_amd64.zip",
     "sha256_checksum": "b696d19ab13d23912f69e24269ea9dc590d08e0e2b4de98d0abb7c90cfa0f372"
    }
   ]
  },
  {
   "version": "0.9.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_darwin_amd64.zip",
     "sha256_checksum": "31ca22b9b8e840789314085ea3a9a666af261b17c0f86b68dfedf1eb50345cbd"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_freebsd_386.zip",
     "sha256_checksum": "4fafc95493beb47f9dfd7c37b71a26416db3a08ae562a1a819ba2a50e31cad35"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_freebsd_amd64.zip",
     "sha256_checksum": "8e3b872ddafbb38951480c5a17dfa18400bfa1cac6303238d65f685b7bc4c0af"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_freebsd_arm.zip",
     "sha256_checksum": "4bbaa3f12348ea08a95b81f55efcf444576b4615f481f886ca6e097d5dd17240"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_linux_386.zip",
     "sha256_checksum": "5c4b95c0d4967051763ded642f868d4fd7783eb49d6a07f57c1d2e0de0e1cdf4"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_linux_amd64.zip",
     "sha256_checksum": "804d31cfa5fee5c2b1bff7816b64f0e26b1d766ac347c67091adccc2626e16f3"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_linux_arm.zip",
     "sha256_checksum": "4f9ae0f2135dd98d752adb36b251c616486ec1dcf709a9ceda8e04787b287023"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_openbsd_386.zip",
     "sha256_checksum": "8e77c160b8518526f7c2f26bd875de8cf208a394a4c0fecfa6ebacf2d6b2cdce"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_openbsd_amd64.zip",
     "sha256_checksum": "f960533a5a6c9830f0b552bc3e4599fee09baa64bcb1ff2dfa39e66d87dbac09"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_solaris_amd64.zip",
     "sha256_checksum": "9d9a2cec776d641744c3049f19729bc20c1e351eb85c6e8fd2fa2eae394d1add"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_windows_386.zip",
     "sha256_checksum": "de4bd25dab3df6fabc99c8873f5123d25483ca3f99f91d533ce9b5874b1fda22"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_windows_amd64.zip",
     "sha256_checksum": "d3db33df518bd09d85858cad5aed5aebf92bea28815f500dfcaa231ba762e5e7"
    }
   ]
  },
  {
   "version": "0.11.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_darwin_amd64.zip",
     "sha256_checksum": "1b5a0c916f547c396959b8c303f3bfa7a2e936c78f002bf42e532c9254fd6d75"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_freebsd_386.zip",
     "sha256_checksum": "5ccfe22d1f451191c894e3e3e659a7b000e508db23d53b2930a84a30627e665a"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_freebsd_amd64.zip",
     "sha256_checksum": "1699ea0e39cd59aa49cd47148948e0f38736a1d1a7faecbcdd08e6542cc569e4"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_freebsd_arm.zip",
     "sha256_checksum": "d65d8b7cc1796203f2e1e655b94c64d395b4e139cb9e10949b2aa24a07e65355"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_linux_386.zip",
     "sha256_checksum": "19e49835c8bda0d1cc5bfc83b61404d5238609f554ee43193194a9fee2c6dee3"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_linux_amd64.zip",
     "sha256_checksum": "5d674e7b83945c37f7f14d0e4f655787dad86ba15b26e185604aa0c3812394ab"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_linux_arm.zip",
     "sha256_checksum": "44b2b584d31f6fffcd996081e024a5fa507bfc3479c81aa71358ff28ff65316"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_openbsd_386.zip",
     "sha256_checksum": "5c9d2dee56be7023a82e6169ec120b98d2b849822acc9dae2ba70ce42a8b0130"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_openbsd_amd64.zip",
     "sha256_checksum": "3eb9650a81cf43761a89b1d6b5690702c0b0e2c444eb6f1e61ccc11027abbbd8"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_solaris_amd64.zip",
     "sha256_checksum": "ff375cd0b0cda0a3dcaa69c1fe0d8c8eedfbe7d6d47f425b22a1f4576b6a5b7d"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_windows_386.zip",
     "sha256_checksum": "12e2817e7eca372995e13a582a7db2a9238abc6e77250344a38d6e4dab5df3b6"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_windows_amd64.zip",
     "sha256_checksum": "d2fa6ce1f5019011678593fe1b85578e94e59bf9e89238258aafd803323db668"
    }
   ]
  },
  {
   "version": "0.15.0-beta2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_darwin_amd64.zip",
     "sha256_checksum": "f66f7fa95ad276198bea21cd960d377be3f9dbe16e91ee628c2f29684bafacee"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_freebsd_386.zip",
     "sha256_checksum": "a77015791dd6ce12026147b20abd6a144417749e70c9ebd02da04732711e373f"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_freebsd_amd64.zip",
     "sha256_checksum": "e1f69f03e305ce8e714fb45debf41521342c9309c4d333a417cb477b436a23e5"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_freebsd_arm.zip",
     "sha256_checksum": "1525d613f47dfd647cd5d837e0d508878391faa9b1e9944cc401fd9fbd04bd8b"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_linux_386.zip",
     "sha256_checksum": "4717567a39895ad1b5b31d39894083cbb2f7ebbd26d2f4756e28be6b0fcb4177"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_linux_amd64.zip",
     "sha256_checksum": "ce6cc6e74c1d144ef45b394cb9ae558c0829aa315435caeb623e9e5b686c40c1"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_linux_arm.zip",
     "sha256_checksum": "1dc2626918b9bc48fcc8ef4eee1f070d74ff05160ebeb93519b29827e9cefc7"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_linux_arm64.zip",
     "sha256_checksum": "ff019be63734ebcb50b1645b923518ee598ca76213a17fcd3fa2c0910ec253c0"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_openbsd_386.zip",
     "sha256_checksum": "a4a23bfd6f02a905f2f49f4677d896f6bdffbbab92b21c5607b7e83161054df"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_openbsd_amd64.zip",
     "sha256_checksum": "e6dcc3aca1e834e4847dba50eac13bfedf89ab9427589a733383ebcbb3494e3e"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_solaris_amd64.zip",
     "sha256_checksum": "570739593b648f1ed257f150da92d494eca0c64b574df608e937813b3c0c5ed4"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_windows_386.zip",
     "sha256_checksum": "44984653ce31ab626a8714bcb636b12b2db7cec20b6ecbe38a703febc08dd20e"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_windows_amd64.zip",
     "sha256_checksum": "4982235821a7c029df8e4c3e28870373aff17e7d056a5a5fa00c3a0ac814ea6b"
    }
   ]
  },
  {
   "version": "1.2.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_darwin_amd64.zip",
     "sha256_checksum": "46206e564fdd792e709b7ec70eab1c873c9b1b17f4d33c07a1faa9d68955061b"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_darwin_arm64.zip",
     "sha256_checksum": "e61195aa7cc5caf6c86c35b8099b4a29339cd51e54518eb020bddb35cfc0737d"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_freebsd_386.zip",
     "sha256_checksum": "7724af1cc0efc3e4241aaab0ad3df884dbe958ccd7cb037adda9358ec45f1aff"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_freebsd_amd64.zip",
     "sha256_checksum": "24cf2247d427a27ad1d5b879dd7c914d3e3dc548f83f44e6783ef968939aa2f1"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_freebsd_arm.zip",
     "sha256_checksum": "453132a8856de71d59ead9f409fa4576a2c68db2f8b9d21d87f07b9cd387f94a"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_linux_386.zip",
     "sha256_checksum": "f9028c4834cd3f2d20975e8ae9897e48f811bb8d494b734a87dab0d3968039c7"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_linux_amd64.zip",
     "sha256_checksum": "e0fc38641addac17103122e1953a9afad764a90e74daf4ff8ceeba4e362f2fb"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_linux_arm.zip",
     "sha256_checksum": "647a69acf01b3cac829f4436e8095d927ee6732cd6c076e24674a3c09326aad2"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_linux_arm64.zip",
     "sha256_checksum": "6da7bf01f5a72e61255c2d80eddeba51998e2bb1f50a6d81b0d3b71e70e18531"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_openbsd_386.zip",
     "sha256_checksum": "63e75a8f5fd9aeb1a17fab6d24828ab9ff08452099f45f7b1c679dfd94a6487c"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_openbsd_amd64.zip",
     "sha256_checksum": "16bf79c0e6fa65412788115abea4805d968f1afa1ba59c4cc40908618f59f7dc"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_solaris_amd64.zip",
     "sha256_checksum": "ab11e20e822a0989fd21f1a589a66a90d7e7422cf9fe55012c455ae798c5425e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_windows_386.zip",
     "sha256_checksum": "27268a9f1ece39dcec92a361c1deb96962ae6e5ece52d8d972ea3110d5af63f2"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.9/terraform_1.2.9_windows_amd64.zip",
     "sha256_checksum": "1425bbe982251dde58104dab3d41f48a51d8735122bdb3790b3b3686c57ebfa2"
    }
   ]
  },
  {
   "version": "1.2.0-alpha-20220328",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_darwin_amd64.zip",
     "sha256_checksum": "f13a13b4656e41e6ca18e5ea894707219ebf4d9dc061a02eaebffc3ec05d26e8"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_darwin_arm64.zip",
     "sha256_checksum": "91d9760779d58e9b0f2089858241afa80a24f476187a392affcbf594e62b3c81"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_freebsd_386.zip",
     "sha256_checksum": "2b74bc35bcc0dcdb54a345fe3159fefb3e58af24439395b588b241c6300e3391"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_freebsd_amd64.zip",
     "sha256_checksum": "300e3522fe0a837fa31fec5e947703f0d39e43cb333b641648209bb24a368720"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_freebsd_arm.zip",
     "sha256_checksum": "376109c7a849bce2b692970ba9aab9b2c0fe6f3766a9f73c9611bdb7154c1b2"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_linux_386.zip",
     "sha256_checksum": "1435cbcf2fa627c8bca9aef94c174441522bf53e862ee58496bf361be5031b43"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_linux_amd64.zip",
     "sha256_checksum": "a03d2410a41830efe0b83233c540a33b2d14caf6310dd25fd09e01930f18252b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_linux_arm.zip",
     "sha256_checksum": "e953fb8c2123e115e130bfe8393143e6e611b961dc270084cb0a24c2affb351d"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_linux_arm64.zip",
     "sha256_checksum": "2a0f36d460408abb27f997c2b004710e3fca5881471d9b238189f8a2e2139a53"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_openbsd_386.zip",
     "sha256_checksum": "7beef6a245bd77fb7fc279dbfe3f2f7d6cca255f388150878638633cb8acb5d6"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_openbsd_amd64.zip",
     "sha256_checksum": "dbd053d65e5bc7b3216e82ec99e4b5727e9483ea283aa87551c01b74f37f6d6b"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_solaris_amd64.zip",
     "sha256_checksum": "a2a76ba0b48c0d8b664afae9deddaa944b7e05f34f1ab9e1bf1fbf911ec8c719"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_windows_386.zip",
     "sha256_checksum": "97472a57f748ac1c68cfca72ac7669a0dc4436ac2d40e3aeadb447e32c5bcb19"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha-20220328/terraform_1.2.0-alpha-20220328_windows_amd64.zip",
     "sha256_checksum": "9393de613d5483d171416703fccfba86142b46a788156f70dc4fc1cbc635bd07"
    }
   ]
  },
  {
   "version": "0.11.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_darwin_amd64.zip",
     "sha256_checksum": "cb5ae1fa5bed45d81d79d427cd1dd84ed7c04f712c72b420003e28f522a77a78"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_freebsd_386.zip",
     "sha256_checksum": "f87227243704702539eda35e631c4ffc3c988016b068a2c89193ff049fc6604b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_freebsd_amd64.zip",
     "sha256_checksum": "84853ba6f57a5b733012f213ca47e9e2c20c99d041f8f982ec6a6fc38f79e828"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_freebsd_arm.zip",
     "sha256_checksum": "48fd91ab28af8003432f658a0e59649643148b7838467608623b8651ae5f820b"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_linux_386.zip",
     "sha256_checksum": "609b0c7d95fc4cc913f54a4ad196c4a90430525affae4f51224e9443c8f14afa"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_linux_amd64.zip",
     "sha256_checksum": "43543a0e56e31b0952ea3623521917e060f2718ab06fe2b2d506cfaa14d54527"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_linux_arm.zip",
     "sha256_checksum": "434ca3420ed4b671aab744559f913d94c0ba1fbae0bacf30aa5fd7e57c4bc7f3"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_openbsd_386.zip",
     "sha256_checksum": "240cb6c3eec87a13e62e830a8e4e19bdc05d6dd8de8d6c9b07b23f5c8d94293c"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_openbsd_amd64.zip",
     "sha256_checksum": "c3bbba3e810aff383918701084fa1260c29b46ae10178909ebb69ca2ff93bc8b"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_solaris_amd64.zip",
     "sha256_checksum": "3b04f6b3fabfdde85ba63cc95b278f9d6ab471d8b25181bb4c4dd1d2fde8543e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_windows_386.zip",
     "sha256_checksum": "e832fe40647cdeea6eebed8abfaf0c775fcd951d1b030844597ad1eaf5d6b205"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_windows_amd64.zip",
     "sha256_checksum": "7bbb3d631fa0050431cc73e7fc9892ef60128d838ed8b4afc1a36f1398c717a2"
    }
   ]
  },
  {
   "version": "0.12.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_darwin_amd64.zip",
     "sha256_checksum": "5cb59cdc4a8c4ebdfc0b8715936110e707d869c59603d27020e33b2be2e50f21"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_freebsd_386.zip",
     "sha256_checksum": "5b252c3f0ea6d24e376da41cd5ef695d772f91074ecb32f2adfe0380ed1590ab"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_freebsd_amd64.zip",
     "sha256_checksum": "d69ad2f39d6542d2a4db1560d5741f77233fec2ab21ddd5a7f170b4f4d30ef4"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_freebsd_arm.zip",
     "sha256_checksum": "5c994884eb6f3a0ecb1f0e7878211b54857792ccef2a72c91153a8f637fa74a"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_linux_386.zip",
     "sha256_checksum": "104d065c15011555a8feb6a53a07cd0c9ea3d5d99721beebc8e55e8d00790f20"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_linux_amd64.zip",
     "sha256_checksum": "a0fa11217325f76bf1b4f53b0f7a6efb1be1826826ef8024f2f45e60187925e7"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_linux_arm.zip",
     "sha256_checksum": "a6e5e9576a3c7b6063010ffa2a377902299a2d938662770d9b0d6622c57d6e72"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_openbsd_386.zip",
     "sha256_checksum": "c31f0221f8b455fb07b6f2fac3d48162b41d80bce9510e37f7f065ebd595a187"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_openbsd_amd64.zip",
     "sha256_checksum": "67e5471d87fd6b35b1a99ac7038083fd97a6bfced518ac4c1901d5efe0c93132"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_solaris_amd64.zip",
     "sha256_checksum": "5d832aed370e0eaf61db76d15cb6083c0dd5217442774a5442f3f9fa16780302"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_windows_386.zip",
     "sha256_checksum": "74a566cd857d277bd23b0af890ad91265fed62275813a6d0e67409730930aa57"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_windows_amd64.zip",
     "sha256_checksum": "ce5b0eae0b443cbbb7c592d1b48bad6c8a3c5298932d35a4ebcba800c3488e4e"
    }
   ]
  },
  {
   "version": "0.10.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_darwin_amd64.zip",
     "sha256_checksum": "5aae5125140b6cb39532360bd725fd33a9224b8358140291ff1d34a086dd646b"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_freebsd_386.zip",
     "sha256_checksum": "249b6bbaaf51ea42038881e2219387495a4688d0e6433fe08c36cb6924e89f7a"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_freebsd_amd64.zip",
     "sha256_checksum": "c0186f4b2de692f71551dad1506a55b444f7561021495ecddfe05715046f7946"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_freebsd_arm.zip",
     "sha256_checksum": "8f537521cadb8cb6899d4157d17c16480e0e5244b9e5da26fa4d56cf31e9a3d0"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_linux_386.zip",
     "sha256_checksum": "39adb75277f4c3d9536b7af4c6107db79a81dbb801783c9e6e66d4f062e51eed"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_linux_amd64.zip",
     "sha256_checksum": "632e10925ae5de013d0d7e76e1c5ae8d5495182acf0b577e35d130de125d2c04"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_linux_arm.zip",
     "sha256_checksum": "bc51c801ff6c6da83332bfe73b97085b41d61694ee7792e7319e1432abb12148"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_openbsd_386.zip",
     "sha256_checksum": "f8b529a540d0cc1771d389f40997992d8f9f6de7ed62f42c5f2515b0649a1524"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_openbsd_amd64.zip",
     "sha256_checksum": "c5e218b454bf502f48e6a5b5bc1d71fb79a5e97448ff5ac57406c235ee28f365"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_solaris_amd64.zip",
     "sha256_checksum": "8f44b5ae7da0abf33a87a0a8f4cd0ddd2ab4eabb1b6da9c61abff89fc80925d8"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_windows_386.zip",
     "sha256_checksum": "547a3cccd527341966a85cce67b32dc562147f990d62414a22ac79f895d5bec8"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_windows_amd64.zip",
     "sha256_checksum": "846b297c2bbe483ad9a5d6cdba25049f738f88cbbf5b167859356e677a947bbf"
    }
   ]
  },
  {
   "version": "0.7.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_darwin_386.zip",
     "sha256_checksum": "2a2c04772bc2c76de9045874d617b78bc19d207b9d63aa16b4052cf1e3040035"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_darwin_amd64.zip",
     "sha256_checksum": "eb6255c4c14c61458ea4598a0e3176695c296e9f1650ad56a24a1cb75d8fef35"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_freebsd_386.zip",
     "sha256_checksum": "d85080e35da1fc6c33938c974a5f1f4d0f63b17f7285ed29901823c5c24efb9d"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_freebsd_amd64.zip",
     "sha256_checksum": "c6ead3b4485f183abe880c4f30f42f7482a239147559f38c03d3ae9f9e94f99b"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_freebsd_arm.zip",
     "sha256_checksum": "a0d6cd5df04252bef7a256fcc15a9880f0ab515781dcc97ebef5d331c6f1123c"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_linux_386.zip",
     "sha256_checksum": "8618805f26e4b1e448a502e7bfaec5f579f41a36d3bc1671449ef5e415c15b86"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_linux_amd64.zip",
     "sha256_checksum": "478c4fe311392804ffc449995e8d7f903abab56f7483f317c1f120d8c21b1a81"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_linux_arm.zip",
     "sha256_checksum": "edd4885d44169ec0f8408e0fd09f86e9f3eab820618c662f84d9ec72432c0b97"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_openbsd_386.zip",
     "sha256_checksum": "33cff5402a41241d26e01e819e328a183e35274af1c575e71c830f6be68874bd"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_openbsd_amd64.zip",
     "sha256_checksum": "67afe320e70478509a0866f163ecf6231da49d5107566e0f03f18076bd75e9e7"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_solaris_amd64.zip",
     "sha256_checksum": "efd5c826f2d5e20ae3f9297831df2ae15f0dfa5cbc1836ffb6c6530688aa6db2"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_windows_386.zip",
     "sha256_checksum": "7fa16ba4ccf6e230d98cbd908a058ee1a75f84476b03332b57f72d3f1059e502"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_windows_amd64.zip",
     "sha256_checksum": "bf4c2a0bbb692e1f30a56992d5f853565a081b2e8b11b8c7d5490584f651c51"
    }
   ]
  },
  {
   "version": "0.9.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_darwin_amd64.zip",
     "sha256_checksum": "71f53879c2fc33af57238cdb67a344d576ae3ae88f8db112122d433bd762788d"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_freebsd_386.zip",
     "sha256_checksum": "2c7d2ede942203c01269c665c1fc29e0bd36895c141582764c885c40bed79880"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_freebsd_amd64.zip",
     "sha256_checksum": "5482d6bee3e2f1b94aac28a8f68cb62320e4102d05100887581235aa148013a9"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_freebsd_arm.zip",
     "sha256_checksum": "201fba5e06e4a9a892b212b1747534deccd794b36c0322273a622e8fcca7c4ac"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_linux_386.zip",
     "sha256_checksum": "c063e59a67a92f562cbe6fcf4aadb3150e367c594db143c92c6f112466ce1a31"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_linux_amd64.zip",
     "sha256_checksum": "7ec24a5d57da6ef7bdb5a3003791a4368489b32fa93be800655ccef0eceaf1ba"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_linux_arm.zip",
     "sha256_checksum": "53d13d0246c7b9d1df5e2f279b60f94ff28dbb87e0271b22c43827e6a505f0f7"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_openbsd_386.zip",
     "sha256_checksum": "43d111445313f9a2a0cf62b05d906aaea5e6615261de022b977ffc0b578f4309"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_openbsd_amd64.zip",
     "sha256_checksum": "ea151bd594fff6c678247fef4d2368c2254c15329f8ce103dfb14b83aa8db547"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_solaris_amd64.zip",
     "sha256_checksum": "93f796e23c20927a1ee09483d944d47ce87d1945b5cd1e77902f3987d293188d"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_windows_386.zip",
     "sha256_checksum": "8e3402e91cf9b51a43b17e40d1b3953fcab9f0d9439ffd3189f5ea9499e2b458"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_windows_amd64.zip",
     "sha256_checksum": "179858a1e9f1a8b9ba7b4c053b5504aedea56394900ce828ce25c90e968dae4"
    }
   ]
  },
  {
   "version": "1.1.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_darwin_amd64.zip",
     "sha256_checksum": "6fb2af160879d807291980642efa93cc9a97ddf662b17cc3753065c974a5296d"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_darwin_arm64.zip",
     "sha256_checksum": "f69e0613f09c21d44ce2131b20e8b97909f3fc7aa90c443639475f5e474a22ec"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_freebsd_386.zip",
     "sha256_checksum": "c04d5422cd44350e653064d80651dbf7c12cd4bf82bb013017fe2d9e230f0726"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_freebsd_amd64.zip",
     "sha256_checksum": "3506199574b49312a353c15f1e7683c1e5f3f41f46b5521070a8e73bb4ac38e5"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_freebsd_arm.zip",
     "sha256_checksum": "72cc644b324efaf7691a64da3097f02cfafb1d0d5c053f63d6bd6ad3faa4112b"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_linux_386.zip",
     "sha256_checksum": "c0ca09a32aad1ff2ebd6eb468b2a3b7f58c2cb75989def314208fecf36442dfe"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_linux_amd64.zip",
     "sha256_checksum": "763378aa75500ce5ba67d0cba8aa605670cd28bf8bafc709333a30908441acb5"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_linux_arm.zip",
     "sha256_checksum": "40e7b0c745832dbeaeca0e74b6a93267942236573486390806d63e0f87427d7a"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_linux_arm64.zip",
     "sha256_checksum": "6697e9a263e264310373f3c91bf83f4cbfeb67b13994d2a8f7bcc492b554552e"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_openbsd_386.zip",
     "sha256_checksum": "5b7e01ad7440c0c75e28bda0829c4fd184e08ca2cb8bc74dc2795c042f874f86"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_openbsd_amd64.zip",
     "sha256_checksum": "d66f197ce000110fa25c7799afbdd0074003169004da10b87dd8a39bbacfd16f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_solaris_amd64.zip",
     "sha256_checksum": "41fa21c6d295395d95c11496a3caf8d0cf78a9eb8d6d07721c48849f206dc38d"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_windows_386.zip",
     "sha256_checksum": "e1498d7e3e7254b24102e7d2d683f84b465bf2b6ca24fb648e80055fb9298ad8"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0/terraform_1.1.0_windows_amd64.zip",
     "sha256_checksum": "a7cb68e34c49336f8bd9792dba3ee350a9eccffb0b9c3ce7275e18e35ad4776b"
    }
   ]
  },
  {
   "version": "0.2.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_darwin_386.zip",
     "sha256_checksum": "758e6de14e78f46f450d256fe4f51c77c334e3787eab270d30356e00d7988ef5"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_darwin_amd64.zip",
     "sha256_checksum": "32c1c5d2df88c612207e9b5edea6f0f4c3bbdc8f2ae5f8c577ede2055548136b"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_freebsd_386.zip",
     "sha256_checksum": "a4179cef2df08c83f2908eb78daab072ea774f4bc4eee78fcf6a16651787db49"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_freebsd_amd64.zip",
     "sha256_checksum": "101567b3e25ba2c1464c6e9b54cca4db3cc43c29c6302e22131bf5332199cb0c"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_freebsd_arm.zip",
     "sha256_checksum": "c3ce9b6bebd8fcef039c76b24bf094089b4d0fb62e286ecc6caed088ea8533db"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_linux_386.zip",
     "sha256_checksum": "7eb62e5fa73a4128b9a59f6c9930f23cbb27e75253382db5a15d04ce61484ad5"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_linux_amd64.zip",
     "sha256_checksum": "b329fe2a6ad54b40c76d4511420bb36d5507a13c19a544d747e5d1691895905e"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_linux_arm.zip",
     "sha256_checksum": "7bd52a861f58e9e681103098749f87c7e62aee081b07b9fdeffaa89a38cd0ead"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_openbsd_386.zip",
     "sha256_checksum": "6c82b2645e27eac18765c2f9c8658c0302fc3bc0197e356ec58d5a168d6a0e94"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_openbsd_amd64.zip",
     "sha256_checksum": "56c3f7587d3556a4bf4bdb72ff6bc9d01a6e884d7709f0596651cf4061e73d67"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_windows_386.zip",
     "sha256_checksum": "4af9fea7b9041c11c4cdfa9ba96700851c580e38af89124afb023ca7c182c855"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_windows_amd64.zip",
     "sha256_checksum": "455fd6cf3a40b9d1ec4e7ae529e3f5b881f2c9222ac234c33a4c2ed72650f823"
    }
   ]
  },
  {
   "version": "0.8.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_darwin_amd64.zip",
     "sha256_checksum": "b80dedb16ab6583afcf66e9b03d3714fbfa44b827094420956d807b710e4fd6"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_freebsd_386.zip",
     "sha256_checksum": "875e0b409e5daf7fcf6b6dd0fb6f2de48d8e0e356103ac4759a53a85a9483f61"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_freebsd_amd64.zip",
     "sha256_checksum": "3b1e11a3ed4a9b9f8c52b20ef84184d42ca78b3812f9f484d3053aa15d0ab5c4"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_freebsd_arm.zip",
     "sha256_checksum": "d6b061203f40fe77496d46118cec6df5aabbef9e35dee74ecb9e62a0ddcc2027"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_linux_386.zip",
     "sha256_checksum": "400097b37049cf194e0b702ca9f98c81c8af881554ab70ae9c59f0dab293dd9b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_linux_amd64.zip",
     "sha256_checksum": "2b4f330e70b757a640ba8d4e1eada86445240b5f8cd43194d878e0c05175f6c0"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_linux_arm.zip",
     "sha256_checksum": "beb588b0a4e373115608d8bc086c0974a0684ca4add9d36755c2ac1eca4cbe5d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_openbsd_386.zip",
     "sha256_checksum": "44b82199d975729faa21e6b6cc30cbe7ed2b6907c2c17d7620aa5a79c81f7a85"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_openbsd_amd64.zip",
     "sha256_checksum": "5dbdfaf92f53248cc5cdf5183ece86557e078a77e187f35f99668d2b775fc26f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_solaris_amd64.zip",
     "sha256_checksum": "39827ebf13a4400fe0193ac06261e569740ffb99c07677f499b5334ca4717685"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_windows_386.zip",
     "sha256_checksum": "27998b13ca5572e79382dc00f67dfba1625d0933103678bcd429d1e0c50ee129"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_windows_amd64.zip",
     "sha256_checksum": "1056a8d8644c1787b2be29f2e12619123040cacd66b9acef43473702eb5ba7d8"
    }
   ]
  },
  {
   "version": "1.3.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_darwin_amd64.zip",
     "sha256_checksum": "6502dbcbd7d1a356fa446ec12c2859a9a7276af92c89ce3cef7f675a8582a152"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_darwin_arm64.zip",
     "sha256_checksum": "6a3512a1b1006f2edc6fe5f51add9a6e1ef3967912ecf27e66f22e70b9ad7158"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_freebsd_386.zip",
     "sha256_checksum": "5d05afef644c3e7e3f3232139f283f61753b09316a3c7ebb12acfe8690773726"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_freebsd_amd64.zip",
     "sha256_checksum": "2b12c2bd12b5f657b3dbfb2c3ee9a5451df5c1d0ebb1173a911c91b7da4d3807"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_freebsd_arm.zip",
     "sha256_checksum": "bf69ff4c8bd90505a6e52d2c04140b61ac9686bf6ded4f493ed214bee4679fd8"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_linux_386.zip",
     "sha256_checksum": "1e90bef5ef5f8272ee020df721c4159f10cc08b7170efeed537dede1cbff78bb"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_linux_amd64.zip",
     "sha256_checksum": "380ca822883176af928c80e5771d1c0ac9d69b13c6d746e6202482aedde7d457"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_linux_arm.zip",
     "sha256_checksum": "92fd3389c72f508808dbf558f7e65eda6748cd853e8141fd5a62561764e5d2a6"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_linux_arm64.zip",
     "sha256_checksum": "a15de6f934cf2217e5055412e7600d342b4f7dcc133564690776fece6213a9a"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_openbsd_386.zip",
     "sha256_checksum": "87a21b47d7c1120a849e4e3b8ac84ac95399108d7ae0f7de4113b22bee3c1ca6"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_openbsd_amd64.zip",
     "sha256_checksum": "caf5f83d5ae1741553d2740816cde707fd58d9be71612d7d4a0616c65999cd88"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_solaris_amd64.zip",
     "sha256_checksum": "540da852acb62964dc86bcbe8ce996cd76f7bc1d2cca8d145e82f6c1ad76944a"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_windows_386.zip",
     "sha256_checksum": "1447daafa69357c5ddc5d1ed88f9120e3e1425bc12b9d5fa576d053510db8e84"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0/terraform_1.3.0_windows_amd64.zip",
     "sha256_checksum": "75d6be7c73c15cc55202ce725a58b62af1f0a5d98224095a166583911bd93d15"
    }
   ]
  },
  {
   "version": "0.6.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_darwin_386.zip",
     "sha256_checksum": "89c28a9e6773456f731bd2f098e283cced3aa42dab2ab0b6645a581e7245bd95"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_darwin_amd64.zip",
     "sha256_checksum": "c519d3d18d5a2b0605bff6e0ca7bb677ea85c833f8e8dbb4af6a48e0ebf76cad"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_freebsd_386.zip",
     "sha256_checksum": "6c073ffcf44fdf0c825ebe0e3fd43327b2cbd24f678ac7c8cdd07320792f112d"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_freebsd_amd64.zip",
     "sha256_checksum": "d69fc7e9bd4f4a04274223db92b6b76ddd6cb769d6d808aae7d600b50f1a5f6c"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_freebsd_arm.zip",
     "sha256_checksum": "7072a2468247dafb6a95b6cefa9af46caf4a5f46beb43ac8c201d7b8db13aea6"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_linux_386.zip",
     "sha256_checksum": "e159da9f08f12e101a4320adf8e4cc2f2d9220e8d77b3bb34c24e00cc4835c5f"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_linux_amd64.zip",
     "sha256_checksum": "7346bd793136d17f646ebbd9d0e797c61b5a32a032632d01ae8d32e3d6952704"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_linux_arm.zip",
     "sha256_checksum": "ff181a31bad9c4ea57608ea2526c4cfb3dd25c2e193fc4bb6bed28185e7a9303"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_windows_386.zip",
     "sha256_checksum": "1fbf2c4a84898c40d8ed7be32b0b33b645c3fc81d1181275c3ce4b56c72db1e7"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_windows_amd64.zip",
     "sha256_checksum": "7f1a30613d9dc4e43a9faab0a16040194013b8713bd9ca5ff1e49bb40371ca25"
    }
   ]
  },
  {
   "version": "0.12.29",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_darwin_amd64.zip",
     "sha256_checksum": "fdcda98ff7b7e65832248f64ef6c2922e05036de25d40c5cdcd732c5117150aa"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_freebsd_386.zip",
     "sha256_checksum": "f90d1ea2954e1ebfac08252cac2cc036dadea3e9c730998337bcaed6561ab9df"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_freebsd_amd64.zip",
     "sha256_checksum": "30cc4c76d1534538ed28be5692d76a72a466760fbb0e648987dbca98abf3a2c3"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_freebsd_arm.zip",
     "sha256_checksum": "9e541926404cad30c7b03fbd8bd9768344fcacec1c849aec3c8b7c31e617d3ae"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_linux_386.zip",
     "sha256_checksum": "8979fe530f451bdb1979c2f09af110acbe17af11b1eb7bb2e29555bdb6b177e2"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_linux_amd64.zip",
     "sha256_checksum": "872245d9c6302b24dc0d98a1e010aef1e4ef60865a2d1f60102c8ad03e9d5a1d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_linux_arm.zip",
     "sha256_checksum": "2826086d167b792eeb11d132548bce7ec30bccc5591d3b34803e9b960b3b736d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_openbsd_386.zip",
     "sha256_checksum": "2fef1b9065d01e1e803459b54500bcb6696351946db3e9e39bd5641acf408e37"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_openbsd_amd64.zip",
     "sha256_checksum": "66aa96663696c0911535a781c4a28ab371451a81ee5111fc0cbcc188485e4521"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_solaris_amd64.zip",
     "sha256_checksum": "c7321899a6d8b7810510ee9ace89a2c304b0c8ad1613112bf970fd24a31fe9a9"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_windows_386.zip",
     "sha256_checksum": "3cd8f9f7d91d7a053a7792d4ee54a69966a5ec2d69ed887e233cb80ea9f8104f"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_windows_amd64.zip",
     "sha256_checksum": "662b0418dac59352fd64525171f98c2eef305002bdd5e168e579063ab68c0349"
    }
   ]
  },
  {
   "version": "0.9.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_darwin_amd64.zip",
     "sha256_checksum": "83b5596c2a510925f90a6572d237b864bc4cf277609ebac294c8f400261e657c"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_freebsd_386.zip",
     "sha256_checksum": "b917b7914e2938d2082459d2287ffae928f01b0c400ea35348f7ced7bec922fe"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_freebsd_amd64.zip",
     "sha256_checksum": "53afdf7a69729b66ab9e541a7cd2923053205e37169fe15a2a6cdb16f2e88f64"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_freebsd_arm.zip",
     "sha256_checksum": "bf4c4af7a2c2239215ce8f1f2af3be27b070b79ab55422df8cdcdabcbdf40b6a"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_linux_386.zip",
     "sha256_checksum": "e89de3f2eaa66d033e4c4995a27bf2ba303788c2ddd5b62ba77fd31dcbd38fce"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_linux_amd64.zip",
     "sha256_checksum": "cbb5474c76d878fbc99e7705ce6117f4ea0838175c13b2663286a207e38d783"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_linux_arm.zip",
     "sha256_checksum": "107d74c36733eb1ef5da1333076bac6f4718b0a4322184519906fa5a1830f860"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_openbsd_386.zip",
     "sha256_checksum": "368935e6ebf133ced4633856d1f698f1424d47d750cb932ffd6ba89498726cbd"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_openbsd_amd64.zip",
     "sha256_checksum": "6463ee1e6d3137abe3a7eab995a58cbed84963f0c9784a8c891036a20eda2a97"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_solaris_amd64.zip",
     "sha256_checksum": "d760c4b74d601bdf6e0d6ce82a60e52834f7381c847b07c9ebe4f0f4f5c7abf0"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_windows_386.zip",
     "sha256_checksum": "280c2c5a09a4fe65a2d73564f723c3ce3128d2f414bc29130186959acf48a118"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_windows_amd64.zip",
     "sha256_checksum": "c3f82782c7164935af3c4404166fe61bbc45c806ea2d51e9b8a9b4c753221de4"
    }
   ]
  },
  {
   "version": "0.12.0-alpha3",
   "builds": [
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_darwin_amd64.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_freebsd_386.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_freebsd_amd64.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_freebsd_arm.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_linux_386.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_linux_amd64.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_linux_arm.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_openbsd_386.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_openbsd_amd64.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_solaris_amd64.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_windows_386.zip",
     "sha256_checksum": "n/a"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_windows_amd64.zip",
     "sha256_checksum": "n/a"
    }
   ]
  },
  {
   "version": "0.12.30",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_darwin_amd64.zip",
     "sha256_checksum": "f107ed316be1b86a63df4e47a1fb8ab8c9ffdbbc606dcdf90043f91bdb21826d"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_freebsd_386.zip",
     "sha256_checksum": "85d085376d4d3143b47f262e13a82ee3f83ee73657aa7ec29cccac8eb2a8f404"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_freebsd_amd64.zip",
     "sha256_checksum": "87a6cfeebe124ff6fe616de71e7a7bc1923ed9a50a370dbcdf2875242bf0fa0f"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_freebsd_arm.zip",
     "sha256_checksum": "d805860ca66d1722d8041c20e51c69a5166aba42825b181c7b323631f6c3d37d"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_linux_386.zip",
     "sha256_checksum": "e2b0eea5d7dd211e3905bc7efae0a87978af5784101e6d2639dcff6306683d55"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_linux_amd64.zip",
     "sha256_checksum": "a646b61232ac0c400ec8cc2c062f4e36b5a9e8515f11f7f5f61eb03fe058f18d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_linux_arm.zip",
     "sha256_checksum": "c3c9033fccfe8878f6192da9d236f0b24b5cc1a8eb49e6d507a2d980091908c7"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_linux_arm64.zip",
     "sha256_checksum": "e7a4d8710257dcb42dce1f5bc97d1540bb37dd39e7f2d14dac710feaf430a7f"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_openbsd_386.zip",
     "sha256_checksum": "edda944b91862da401593245d234130e2b1642c92bffb3518c80a896b6b1f395"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_openbsd_amd64.zip",
     "sha256_checksum": "613c3fd68aa59a2fe46a41cf9484040d779090de75556269fd89c9211478777b"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_solaris_amd64.zip",
     "sha256_checksum": "951cf62fe9aad4d988c348188d0857c0761f341ed30109bdab30af731c19ca34"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_windows_386.zip",
     "sha256_checksum": "cebd2b303f8bb73e1e0d434671770e06f31c3a5b6987c693aca471c64a622ad1"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_windows_amd64.zip",
     "sha256_checksum": "628f3d90da89258199a47894f1e7a3fa957c6af8c343316ace38d2737734ab2a"
    }
   ]
  },
  {
   "version": "0.7.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_darwin_386.zip",
     "sha256_checksum": "6e4f6a9d450b03ed52108e818798fad19ddaa442bb8db98349789d0b330cddcc"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_darwin_amd64.zip",
     "sha256_checksum": "ab5e9ffe690f52ff13b8f095937119d67d3f0a07744be851657555236245dd98"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_freebsd_386.zip",
     "sha256_checksum": "ef68d9b1199aa726162f61478d032ae753eb189cbf3d9ffbd58ee0b234c73b6e"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_freebsd_amd64.zip",
     "sha256_checksum": "cb4a6c1166489490ca207a96930060f2608e342ad6d4119e196f39c38e866ea1"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_freebsd_arm.zip",
     "sha256_checksum": "527c7fe997bc32cb409601b4c3224d679ac01d9bdff9d214bb011b60c3b35850"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_linux_386.zip",
     "sha256_checksum": "d1ea1e8285da801a29d005a0e50c6d16425544160095ff7a503fd52e6aa39185"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_linux_amd64.zip",
     "sha256_checksum": "133766ed558af04255490f135fed17f497b9ba1e277ff985224e1287726ab2dc"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_linux_arm.zip",
     "sha256_checksum": "6ea8e54dafa88a75d199a9d6e61cccb1d753d8fc98041127d2c09ff86e16ff7e"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_openbsd_386.zip",
     "sha256_checksum": "e04c00074f841f3b4e1f5e803c1d55eeb8aa67d063b331fabf420b9b13e29d1d"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_openbsd_amd64.zip",
     "sha256_checksum": "a0046a3db9612c6f1f171eb0da0c6ec3201a0535a43704c071daeb9297ffc42b"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_solaris_amd64.zip",
     "sha256_checksum": "3cb873118e07255f96bfc8b754adf038d5cdfd32c63c270c51e92aa728612294"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_windows_386.zip",
     "sha256_checksum": "e63cc458ce18dc916038a9e5281fd2e99cfa454ea56ad7b4709fd4eb6dd0b1ee"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_windows_amd64.zip",
     "sha256_checksum": "e80f22482e6da22dfb7d0caa064265d0673e0d120e81dc3d15eaa489647ec852"
    }
   ]
  },
  {
   "version": "1.0.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_darwin_amd64.zip",
     "sha256_checksum": "b67f5e25a73866b83880fd6fbc5e57add3ed893a31eda26b748aea4afc7255c4"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_freebsd_386.zip",
     "sha256_checksum": "95e37242259b41cce51ab397e3486a91c63c3c9c9d0db124ccac12a61b89ba21"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_freebsd_amd64.zip",
     "sha256_checksum": "84081a743bae06df77d36ae30391031a19e5290656ae285e0c5a0087b716cb00"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_freebsd_arm.zip",
     "sha256_checksum": "28eb8abb84d3463389822ab6b09fc2beabf13ab90001643e4511452f8af67d5c"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_linux_386.zip",
     "sha256_checksum": "9d08dd8ecd4f2946c22ce5733f9751ce3f71502b11952465420b0b65db60723f"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_linux_amd64.zip",
     "sha256_checksum": "8be33cc3be8089019d95eb8f546f35d41926e7c1e5deff15792e969dde573eb5"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_linux_arm.zip",
     "sha256_checksum": "9a08693fc466f28adc40dde7b23f2da6ce2947493ca6b78c68fb7a3c716ff4b5"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_linux_arm64.zip",
     "sha256_checksum": "ca7a3c5d03912d7a4b425866c6972e61a3f4595a2e19d800f797ef1fa5d206fe"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_openbsd_386.zip",
     "sha256_checksum": "bfcfa0aed8efc8e3387d6476c41762196a3dffa2193bc6f783a6cd150ada5605"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_openbsd_amd64.zip",
     "sha256_checksum": "e802ef42850e9bae0ded58ee0c3bf4c96d9fd81b5bac312e41db48c62ff63b16"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_solaris_amd64.zip",
     "sha256_checksum": "2b08849f6c8817806e5e2fcaf16703239976f6c416b9ea4ba8a4edcf0cdcbcc8"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_windows_386.zip",
     "sha256_checksum": "eb55cf1dac2cb2c0cbbdc4a4b77331fe93c91ada305cf82b77ffe8e905c253e5"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_windows_amd64.zip",
     "sha256_checksum": "82ff30315a9d43c477adbf31dead4b3ff9f96415754f8e7dc22a481982e08098"
    }
   ]
  },
  {
   "version": "0.10.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_darwin_amd64.zip",
     "sha256_checksum": "1ad6bad0349a3bcda8264746a3db0a39875c2cd93e3418393cc082bbb4812541"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_freebsd_386.zip",
     "sha256_checksum": "5bba703527849d24911a497e2c43ca2df687584beb9241e1744ae879180a4049"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_freebsd_amd64.zip",
     "sha256_checksum": "f56a74cc187cbbfda929b58d758b50f1e25aefab425f476fb6c99ade7b9da842"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_freebsd_arm.zip",
     "sha256_checksum": "75df067654d220d4cfc067d31d72734be881198a5f432c78a77c3ffc7fb7a27a"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_linux_386.zip",
     "sha256_checksum": "c920c8eb52ce800e7271fc30166759488f6e22406eb48346937f19904b581465"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_linux_amd64.zip",
     "sha256_checksum": "6c1b5ce1a78bc7bb895055052d9074e519f51293770871acfe2dbd289e2f2aaa"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_linux_arm.zip",
     "sha256_checksum": "e68eb1d7124611038f8f24d21f9d4fbdc97f693c4b846fe2e72a6f89cb5db095"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_openbsd_386.zip",
     "sha256_checksum": "2f73089865f207e77d620f96119125f94342bf8f3376c715faac2079d03ff716"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_openbsd_amd64.zip",
     "sha256_checksum": "9538e27cd266bc966ec02cd45ec2cdf49197bb07bb9680c0f8308b9353e0573d"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_solaris_amd64.zip",
     "sha256_checksum": "856ad31e3e8fd7b9ed69ed70dc1a4790e09944fd45f74c2be820cd2c95d369cf"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_windows_386.zip",
     "sha256_checksum": "e8459cf9c9b4ec1e732ea02648b619e8ba127396bf72c999af4f22c02abc2dd4"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_windows_amd64.zip",
     "sha256_checksum": "b117dfde4c64c626a0551b88eda94c0d08015ff221935c0c82855203d7378c99"
    }
   ]
  },
  {
   "version": "0.12.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_darwin_amd64.zip",
     "sha256_checksum": "9dbee9dea660ea64352f8ddf2539e60d1c414210e9c4a29c8585926fef366be1"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_freebsd_386.zip",
     "sha256_checksum": "ee57936ea0a4bdac65837464abac7c8a5f68387693e6e479128f96c66facb930"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_freebsd_amd64.zip",
     "sha256_checksum": "571e391f4cb581ff6c3815fcfeac0aaa0d7d8f25ff090e192e2e227a7df49598"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_freebsd_arm.zip",
     "sha256_checksum": "a49547b309cb7e5baecca91032991d2140c26d171357604460d7da862905e65f"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_linux_386.zip",
     "sha256_checksum": "ed46f577a29dab97bd1a6f2176b31920f8a6bc81186e7ae8b01eb597d69341f9"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_linux_amd64.zip",
     "sha256_checksum": "42ffd2db97853d5249621d071f4babeed8f5fdba40e3685e6c1013b9b7b25830"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_linux_arm.zip",
     "sha256_checksum": "ea7bc7ed4452f3e2fc74cde8cdc9b42daea0c38b6255326e2911d7ae16bfb166"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_openbsd_386.zip",
     "sha256_checksum": "9df507d4e4e9a34871eeeb09427b64456a8381af2eed37e43f907bdb935f71f8"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_openbsd_amd64.zip",
     "sha256_checksum": "a41a6a22a689ca63ffaa9ee77443cca536208ebda2ddd7afc8af0138c18751c1"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_solaris_amd64.zip",
     "sha256_checksum": "dc38629a2f834cbe52632951946d02d2beeb978d8b0e1fdd547c0b7ac5f66ea3"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_windows_386.zip",
     "sha256_checksum": "c56c8ce625333006b0d52334a8ffc95d181e50bd6737bb0c1ac77bb4038a51ed"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_windows_amd64.zip",
     "sha256_checksum": "737ec6a684669125579858700a294799aba7deb7a72393eda64bea99aff8b38d"
    }
   ]
  },
  {
   "version": "0.3.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_darwin_386.zip",
     "sha256_checksum": "2514928f49d3d9285e43706ffb4f3f1e63d2763fa1384fd397dbdb86a8d61e54"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_darwin_amd64.zip",
     "sha256_checksum": "aecdc8119cd637e3e60967c97f9912735400814546b8e925152203fb6e99c732"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_freebsd_386.zip",
     "sha256_checksum": "5b80ceb8fa7a5034135234c73e090096276f4ae4289ee78afbdf8705323c974b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_freebsd_amd64.zip",
     "sha256_checksum": "6183cc4784dd6b436157a8fbbac2525cc054f64a0a666c07a1384838ccdaf101"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_freebsd_arm.zip",
     "sha256_checksum": "3fd50924bf629c0f91eafa754b06dbd5ef03a5d6366eb17b61a7d2829c8bf52"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_linux_386.zip",
     "sha256_checksum": "e7ef409c65439c7f1195aacb629a01dab2a6de67f79dec40093bec42d8be9d6c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_linux_amd64.zip",
     "sha256_checksum": "b63b36d76d9ea31cb6971edf23899ceb7291800618177adf987b7660b486527b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_linux_arm.zip",
     "sha256_checksum": "6f636f6bc892833cbb9c54f0826841151d3295c2c6dc83caecd714b6afbcb47d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_openbsd_386.zip",
     "sha256_checksum": "86f2f444f9ec92d363a01b0acfb062d5078191788686d5c85b621f08fad38658"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_openbsd_amd64.zip",
     "sha256_checksum": "6e91575d95dd8a561602b0bba1a1debe6f8f22467385909259f86149114723dd"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_windows_386.zip",
     "sha256_checksum": "8cdc90b3d30cfa7f9548a5d1e433cf27dfa1128a8cc828f96f8f6c59a7feac36"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_windows_amd64.zip",
     "sha256_checksum": "324ff1e3e7d86b927b8b621f125932b8f12b88edbccfea484a002b2f9043685a"
    }
   ]
  },
  {
   "version": "0.12.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_darwin_amd64.zip",
     "sha256_checksum": "2da57018c25ada511b7131d85257f534030eddf23b347663af4c4ca89d3d9220"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_freebsd_386.zip",
     "sha256_checksum": "1cf3917c39e0a551f07b6adee13d17d2c09bd5570fd95e99e329d653f8573c33"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_freebsd_amd64.zip",
     "sha256_checksum": "7a2e2b56dce91900c5ec39019509e686fa3050bdf1f7e2e3da3c626b8a771aa6"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_freebsd_arm.zip",
     "sha256_checksum": "c0d9d3889ab151071107bb5ae53dce3ae8b29ce8a8201480e611e4e1e24e1012"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_linux_386.zip",
     "sha256_checksum": "99c1b8a5f825adbb88cf9cbc1591696904fae3c11addaf5563548153c9598a55"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_linux_amd64.zip",
     "sha256_checksum": "3ca1a3126e17bbe8cf8aa677ff39f8b7a5cff71de91fa9286c3c7cacebba78e6"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_linux_arm.zip",
     "sha256_checksum": "316cd80547f024ea1515e1da8861fc3c3e86760f2ae43e88faca2412e56beab7"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_openbsd_386.zip",
     "sha256_checksum": "d8298cc4f552f19f184e57ca0cef08d9ede7c101ec209700f057ad787f1272d6"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_openbsd_amd64.zip",
     "sha256_checksum": "af1313f7b788dfc4500b809d314228eb4a73f98efaa798d4834fed1940914702"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_windows_386.zip",
     "sha256_checksum": "d997e1419a34affc437b0f89aba24b474084a79b4ba76b20d0a8282f18615191"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_windows_amd64.zip",
     "sha256_checksum": "63e24b1a2179cde24cfeda227d6396330711cac7fe05a27900056d28cfd95158"
    }
   ]
  },
  {
   "version": "1.1.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_darwin_amd64.zip",
     "sha256_checksum": "214da2e97f95389ba7557b8fcb11fe05a23d877e0fd67cd97fcbc160560078f1"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_darwin_arm64.zip",
     "sha256_checksum": "39e28f49a753c99b5e2cb30ac8146fb6b48da319c9db9d152b1e8a05ec9d4a13"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_freebsd_386.zip",
     "sha256_checksum": "4163bc7f77d4fb8db6dfa4129628674752bdf2a6a1b9d613c6f8e48f5e185049"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_freebsd_amd64.zip",
     "sha256_checksum": "40c8c98978af2ade2aee551991b227e77003ef0a60d0a42a0adb2de5d001c080"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_freebsd_arm.zip",
     "sha256_checksum": "4c9484af969d20718533ed41cdac87548beeec165ae467cc74e0e19b7f5b6d2f"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_linux_386.zip",
     "sha256_checksum": "9aa333f3c105b2636a335eb48af8ea7aa434cb9e52246943fa2a55c9fbab494c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_linux_amd64.zip",
     "sha256_checksum": "734efa82e2d0d3df8f239ce17f7370dabd38e535d21e64d35c73e45f35dfa95c"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_linux_arm.zip",
     "sha256_checksum": "afba01532a637df6209dbc38e9dab75add13ff39fdb3254f4ad2c26b0e652df0"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_linux_arm64.zip",
     "sha256_checksum": "88e2226d1ddb7f68a4f65c704022a1cfdbf20fe40f02e0c3646942f211fd746"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_openbsd_386.zip",
     "sha256_checksum": "e38dedbc9b5a7fd6335c77463042b3948e9339bff3c2b8f702d176b8eaaf4f78"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_openbsd_amd64.zip",
     "sha256_checksum": "a2d12bc2348f270efe92fdacd3b8db4ca32182a05397bf1802a3226e8c4f9f58"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_solaris_amd64.zip",
     "sha256_checksum": "46dcaa93437ed85ed3e8bfbdbaa14155ff02b3bd955988c523acfd6c0e4733ef"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_windows_386.zip",
     "sha256_checksum": "6aa818b9ad9e7de02f8c45bc7fb8f6a0b0d603f22d48711059d74a013696bb00"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.2/terraform_1.1.2_windows_amd64.zip",
     "sha256_checksum": "4c40b48b78e1c7b049fcff2f0cbf287b3df1c6d602fd6635189c45414ed16f94"
    }
   ]
  },
  {
   "version": "0.6.12",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_darwin_386.zip",
     "sha256_checksum": "5932c2125b6e63f8890045ea1243200dbe29d9fcb2c0f24368aaa17d5fc5cdc5"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_darwin_amd64.zip",
     "sha256_checksum": "eaa50e05a88ef83a9ba18a3768932f4d530ce1b710b29ae29992f94addac0bfb"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_freebsd_386.zip",
     "sha256_checksum": "7552f81abeaefd36f5fe2b1c1bdd319913650cfd7d927639fdd0e7a4752da851"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_freebsd_amd64.zip",
     "sha256_checksum": "bdf5ad9c8b24feada0629ba50ea95e463231442227d4cfc016ecf1a6448771b0"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_freebsd_arm.zip",
     "sha256_checksum": "3c6a48e131dd15be573fe753dfa89b4dba4011128570e978a505abb27b11e5d8"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_linux_386.zip",
     "sha256_checksum": "c936e073988cef78dd1c1d29873276f85a7ec44329074bc353a022054fbd5e06"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_linux_amd64.zip",
     "sha256_checksum": "37513aba20f751705f8f98cd0518ebb6a4a9c2148453236b9a5c30074e2edd8d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_linux_arm.zip",
     "sha256_checksum": "9d9db3dd01ab7c4d32a29f1b3f3f3e3f1bf823c017c720e91a27d17b74d13f1"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_openbsd_386.zip",
     "sha256_checksum": "45c58ab39e33d9b454287b2462f76e4ff9da53413aa0eb9184c16b0c25bbd5be"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_openbsd_amd64.zip",
     "sha256_checksum": "b8ba5faa943567514f5026ed89bf31db82f1398f1cabf2601dea7b0ea966f392"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_solaris_amd64.zip",
     "sha256_checksum": "75fc480b817392f6ced2c6e333778fe7921dde13890da60588c34d3a51a49e94"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_windows_386.zip",
     "sha256_checksum": "af538377662c9a7996eb7f57b10526d10ad42b5dfe3c40c74aef2a75fddcc467"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_windows_amd64.zip",
     "sha256_checksum": "90b42c37d3cae73badc1be1984a7836547b5877f4d51c75bb77a18d1de0a3cd4"
    }
   ]
  },
  {
   "version": "0.12.12",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_darwin_amd64.zip",
     "sha256_checksum": "51507dedba7fcc2638c5c2c40206ec604155e2d3067a132b618f4e99ea9f1db9"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_freebsd_386.zip",
     "sha256_checksum": "a57401f598cb7b32e6abc59ae480cfa2830d54e9c1078f0ae27609f3e64d6ad1"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_freebsd_amd64.zip",
     "sha256_checksum": "9cf417e37ce9d2a2347db158725149639b37bcbc42682a495dde546eee11b63f"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_freebsd_arm.zip",
     "sha256_checksum": "d5a3088244c66f73215c5f229431afa5c7eb98f4ccd565501b2103664c5c721e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_linux_386.zip",
     "sha256_checksum": "e22cd136ba4ef6b0204bfe2b0ac40a1dfdc086d0557c44ddf3e60e95f243f4cf"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_linux_amd64.zip",
     "sha256_checksum": "67bc7a49c0946ad48b14cc6e95482fdd3e7e9f7dc6811f4ce6ff531fc565bd3a"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_linux_arm.zip",
     "sha256_checksum": "300baf222721e48e355a160191870bb5a772f1f263bb0c2842e8977909959545"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_openbsd_386.zip",
     "sha256_checksum": "75ff5cc86efd5088f71ada298b4aebea4ac3c8b1bb9b43003923f63e61e5abf2"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_openbsd_amd64.zip",
     "sha256_checksum": "a929b0e1e3a996f82e8629578c93488ea61e4ae61b144ae1a6d5fb109ad85068"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_solaris_amd64.zip",
     "sha256_checksum": "54ad3b833074fb85ca123a874b6724136cf493df0f7d17b852c5cf799e30e436"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_windows_386.zip",
     "sha256_checksum": "3750cae1bcfb06b6d36936c8b0fdf3e2d114b20d65a6088fe4bf14a9caa0ae96"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_windows_amd64.zip",
     "sha256_checksum": "fe06fbc99b6403192bb6249d151fa701801ff1b1b9102c2400996e8a104714c5"
    }
   ]
  },
  {
   "version": "0.13.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_darwin_amd64.zip",
     "sha256_checksum": "7af2f9c03e8687c87e7798178a2dac9a3061955eb19f0f69501475e017b8d8f6"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_freebsd_386.zip",
     "sha256_checksum": "62044072f0b73f26c9b3b77d6cf925ac53cc0ecef419d15b3993ee2ed15e98a7"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_freebsd_amd64.zip",
     "sha256_checksum": "e131098931ed49f493c11e1bb1780532efc00af210cd488c247b1213296a1b5b"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_freebsd_arm.zip",
     "sha256_checksum": "8d6f4343f3775be50b98f4f70fe6214a291f02c9b2ee9f1453da2519b99a302a"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_linux_386.zip",
     "sha256_checksum": "532e761483d68f056e1a038f482dfa7af75fd6a879d5cd33e7c780ca617ea1a8"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_linux_amd64.zip",
     "sha256_checksum": "6c1c6440c5cb199e85926aea65773450564f501fddcd7876f453ba95b45ba746"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_linux_arm.zip",
     "sha256_checksum": "5383a0a8dc716fafefc30f08c4c6b73e06c28ba8c2a396fe7ae2d1cb9787c95c"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_openbsd_386.zip",
     "sha256_checksum": "45ae994baf3300805727c62f7a898756397efde2388502d87ef7369beeef7821"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_openbsd_amd64.zip",
     "sha256_checksum": "33637472710a1c2fe77e6fd3047bbd77599764992ab9cff5b7935ecb6ed7f2f5"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_solaris_amd64.zip",
     "sha256_checksum": "2702fd31585c51909ea95a696f49783c40a3921a9a6e8ed703d703836e27dbec"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_windows_386.zip",
     "sha256_checksum": "43bc61096c51da8b23652e94cb86952f0008ebb30edf86e81965ae09037da04"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_windows_amd64.zip",
     "sha256_checksum": "56582718c6694c00e27f6c14ef2684059798961d19e389ab24d9ded0c7d898a1"
    }
   ]
  },
  {
   "version": "0.11.15-oci",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_darwin_amd64.zip",
     "sha256_checksum": "4ac1b0a1a7ee9e04165ce035300eddf9119124046d63fc4bfeffcc88fc6365bb"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_freebsd_386.zip",
     "sha256_checksum": "db0cf57dc21bdb360d4179cd1c8611c8cab61e6dee9bfbc04cbe536f105c5f21"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_freebsd_amd64.zip",
     "sha256_checksum": "56e7c083a12c7fa16f0faab7f1f82d82420b3f94d0b77b161ee99386fa7445b7"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_freebsd_arm.zip",
     "sha256_checksum": "3a292aa0bd10ebfdb31e5b39ae4fa34d7594416ba89cd8501b97ef25b5807be2"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_linux_386.zip",
     "sha256_checksum": "c82fc0d9c1a06bf82d5369ef45a1676d372b17fbf80c460d63e61a78cdadb3e6"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_linux_amd64.zip",
     "sha256_checksum": "6f4edcb271ef332e7757193da6f84cb5fb48e1d8e6208cda69d556b4a903c30a"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_linux_arm.zip",
     "sha256_checksum": "666264d01400cd6cffa7ca96572b4df3cdff97297d70eb0b2c613c1d529e66b5"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_openbsd_386.zip",
     "sha256_checksum": "66311bbadf694db31dc32b4b6327ce247545c028ba205be45c2235428cf582cf"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_openbsd_amd64.zip",
     "sha256_checksum": "2f9cae7af7416bf153127d4f437f891417d03b939e580144113e6da7472c4341"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_solaris_amd64.zip",
     "sha256_checksum": "14bc42df6bedfc5fb953bc1082a892510223d8a9786835a4ec53be15c1423a0c"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_windows_386.zip",
     "sha256_checksum": "79065368a2bb6365927896e6d4c7999f2ccf5494c0e518fef1bc1eabcfa5f36d"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_windows_amd64.zip",
     "sha256_checksum": "c047f0e7dc9aabc28beee1a903f964c0af1ac2d7304ba5fa7e27c8261888b250"
    }
   ]
  },
  {
   "version": "0.13.0-beta3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_darwin_amd64.zip",
     "sha256_checksum": "8e49d45da847120ea1e162d0b3fcd6b322e8dff419c6cc5cb535a3041a650391"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_freebsd_386.zip",
     "sha256_checksum": "da0f82f6628a7d76ca4615f5f88544a813a58220655294312c40bee9c45a142b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_freebsd_amd64.zip",
     "sha256_checksum": "4ca87852c545bb7baa3d11b3b50d945ce89fe4015196672b3daa6bf865fec9e6"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_freebsd_arm.zip",
     "sha256_checksum": "71958f63e48c4ef8f1401f4d0b808f5748e34e908b2841078299c5de2c6571de"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_linux_386.zip",
     "sha256_checksum": "39b6a2edfd4617c6beb2d3c4dbfa81a2e3e4313130f19b169e8e423a8d5ae648"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_linux_amd64.zip",
     "sha256_checksum": "f842254756616d1ef90c06e611a1d4cc209ee250473d3ab123b1f6bfb03bd724"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_linux_arm.zip",
     "sha256_checksum": "6cb7f1f2360ff51dbdc07b0db1d5747567d5c434477c71d7ba5ce0b283ff0095"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_openbsd_386.zip",
     "sha256_checksum": "2efc3d6705b7bc3cb7c6687f719c7fe696556b844311b52779c0a1e8ba28e58e"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_openbsd_amd64.zip",
     "sha256_checksum": "a5470d76c788a35238accd702e6a13e364d3a6f959c9fe70494bfc47092c1315"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_solaris_amd64.zip",
     "sha256_checksum": "9bdb8def773f9105d185a82580cedd1ab15c43bda4459afe2b7802910d34efd1"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_windows_386.zip",
     "sha256_checksum": "10cb47612534c91c6a6edcf5aa6c7a8723251ca12fde7707f258028aef8218c9"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_windows_amd64.zip",
     "sha256_checksum": "c632a9e27196d7f8dfe9ee9269f1c572fb2a8929367aa6080abd41aeb0d3ed11"
    }
   ]
  },
  {
   "version": "0.14.0-alpha20201007",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_darwin_amd64.zip",
     "sha256_checksum": "f2689edb2dbe46fafbdf92062a37c695d398d5224756c57db62fae8097f54a0a"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_freebsd_386.zip",
     "sha256_checksum": "f726ed8797b6eb337126e7b94f6d711f14d84efb6202754f782899f5f16ede47"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_freebsd_amd64.zip",
     "sha256_checksum": "ed0920b347bb7cc3095689b1ed27d61c59042b1812d2c08d6e1ea9f468b40d51"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_freebsd_arm.zip",
     "sha256_checksum": "47b64150d5b747a18b555c0dd5a07bc0880bcd1e8b6b50d8e913f5a1aee13f49"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_linux_386.zip",
     "sha256_checksum": "c82a4e644940650ccf6cff6b69a6fb858202a7d2fd09401fb1e9e811badb6f0e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_linux_amd64.zip",
     "sha256_checksum": "7528187b5023e87ef5e179342b66b2514687ddf35c6ac279f549000272983aff"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_linux_arm.zip",
     "sha256_checksum": "37bee5b92b0e3c7a0cc60105c4920a0617438fb320517aba074b224fa3d16c41"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_linux_arm64.zip",
     "sha256_checksum": "3c7dc81d571afa1cbf286d17d92a780870f5d3528f7db46cbddea881f6387529"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_openbsd_386.zip",
     "sha256_checksum": "a3c14a64b9bb6d2efadc35e06814af4ba630dcf401d9f60be98276fc23778899"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_openbsd_amd64.zip",
     "sha256_checksum": "81dc450b4ad031cce685a7383edcf98c2fff46540674cbcc7bae1fbdb95bbfaa"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_solaris_amd64.zip",
     "sha256_checksum": "d44301caade04a8a663a8ad80c795df968bb953591780052a4f72c79d453c359"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_windows_386.zip",
     "sha256_checksum": "d31f371a505bd3b7617efd2d317a14fe32aa0b295d089b938d3aae3869fecfde"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_windows_amd64.zip",
     "sha256_checksum": "72fd63d5d6674a384e5a0d8840b875eff1dc3c0067d0553508f79fa26312f04c"
    }
   ]
  },
  {
   "version": "1.1.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_darwin_amd64.zip",
     "sha256_checksum": "c2b2500835d2eb9d614f50f6f74c08781f0fee803699279b3eb0188b656427f2"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_darwin_arm64.zip",
     "sha256_checksum": "a753e6cf402beddc4043a3968ff3e790cf50cc526827cda83a0f442a893f2235"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_freebsd_386.zip",
     "sha256_checksum": "f4244bc94a3a00fddc43c4e00c302ef40a2c41f74ec54b9512b44c6c58428afb"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_freebsd_amd64.zip",
     "sha256_checksum": "d49f3642439c6099329cc3e132b94229f572133eed3a248512531691a014f294"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_freebsd_arm.zip",
     "sha256_checksum": "c178c9a431941eabfcac00956eb69027acdb2475c7bf6e6f36b4cdeb918bdcd6"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_linux_386.zip",
     "sha256_checksum": "1277c4b0d957be9f1e89bc9b49af2b2830349fc22a2943562bf40fabe818e86d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_linux_amd64.zip",
     "sha256_checksum": "fca028d622f82788fdc35c1349e78d69ff07c7bb68c27d12f8b48c420e3ecdfb"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_linux_arm.zip",
     "sha256_checksum": "390d8d177b26fcefe80d6a5b8d333b9ce29058a99f9d52f9b3e10e92ebcdf3cf"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_linux_arm64.zip",
     "sha256_checksum": "3c1982cf0d16276c82960db60c998d79ba19e413af4fa2c7f6f86e4994379437"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_openbsd_386.zip",
     "sha256_checksum": "61fe9de27583e13797f2cda03e7b2094d62b18ac99823136ce95b18054d6f2"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_openbsd_amd64.zip",
     "sha256_checksum": "d1b36b1b81e2e4f5a1f444898a835035cdcaf3262fd12904677936bf275d5cf6"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_solaris_amd64.zip",
     "sha256_checksum": "3c0d52dd3f15c5343e06817324aa1f6de391b0556623239ff32fa4723e9fd4f8"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_windows_386.zip",
     "sha256_checksum": "d6af8634e1ad02952fb33215b46da2fd9f7537e55dc97d589cf9dcd1e7203fcd"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.4/terraform_1.1.4_windows_amd64.zip",
     "sha256_checksum": "32a9ae97b361aed822e7e151368c47ebf10ec67a2f040acf0415e4f28be54e0b"
    }
   ]
  },
  {
   "version": "0.5.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_darwin_386.zip",
     "sha256_checksum": "191a4150c86f9b12d922e3300a5dc884f3fcfcbf0158d3c3a9d2c82e8738037b"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_darwin_amd64.zip",
     "sha256_checksum": "5915d7668b07ea3770f1bc8126764f90723eade0245e0634af3b051ae2ceb7e5"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_freebsd_386.zip",
     "sha256_checksum": "c60f23cafd309faddb592adf42a85a8b672f8d88c050a1cf8d8365c5939d0744"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_freebsd_amd64.zip",
     "sha256_checksum": "5c84d810a5e2d813ed299650cb1b3e91210a1d189a08e8c531d4cd9d5e948ed8"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_freebsd_arm.zip",
     "sha256_checksum": "87e8c4d758be1f040ce902952ef7f30330c34efde798282b230e8f8ab05f0d9b"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_linux_386.zip",
     "sha256_checksum": "97d80d595c87387205a84a4140eae6e85ca784abb7ba88b7f5cef2b7c90f377b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_linux_amd64.zip",
     "sha256_checksum": "5f651ffd0f8d7386ed5d44d50ef0053ee974d1a73bb9becf7c99c886615a98f7"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_linux_arm.zip",
     "sha256_checksum": "d28d10b1d08e8d525963dbc662ed0071f34d5133bcc8397240e5666ad38eee9c"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_openbsd_386.zip",
     "sha256_checksum": "559562e2dd1774ea95d981acba11b8c5f47a3078aafe90e194848b9834017123"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_openbsd_amd64.zip",
     "sha256_checksum": "e7b1b2f43411c3f950cdb82945bfa70b37e835e4c510462b720c7cac4c0ccc2"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_windows_386.zip",
     "sha256_checksum": "927a556566ed4c3cc9a2c64015927bf9038aa6892258451fdf9d99791a2889a5"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_windows_amd64.zip",
     "sha256_checksum": "fb0b9a1c15252c33fcc523ccf9cddaecc76ef122a129a0badea4b2817512c279"
    }
   ]
  },
  {
   "version": "0.15.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_darwin_amd64.zip",
     "sha256_checksum": "96537262e38008a421d329ce51c1bc2a1926f0b4e68270c92a81a8a42fa2c513"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_freebsd_386.zip",
     "sha256_checksum": "4c92c50a36346b196f21a8b1925692164ac7e17458434fd721445dda9ebb7b1f"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_freebsd_amd64.zip",
     "sha256_checksum": "71af8ec9ede7bc8533c92a8f85c59749fece88d45b7cfb9835bfc4e91ed0b5cd"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_freebsd_arm.zip",
     "sha256_checksum": "bace7c9f1d60348897ce28c223d53bf568b9edabb446172e0527e566726e9996"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_linux_386.zip",
     "sha256_checksum": "1d7a7fe31b9c11881710b8ddf3d2353ebe3afc5f93a51177e70148f39abc8fd3"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_linux_amd64.zip",
     "sha256_checksum": "69c5db7bd6d4a5d3dd060678e5c3d9442e32610ed05879b4325e6aa4807d0529"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_linux_arm.zip",
     "sha256_checksum": "f9b034762d0acdb04b4d7a095c1c6a4088e6e064cc510a6d6e2479e957406204"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_linux_arm64.zip",
     "sha256_checksum": "a83897c82bbfc01a390fbd9acd3cd0cc91f179705385b6ae276a0d8de98a8cba"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_openbsd_386.zip",
     "sha256_checksum": "7a6c01f9c9a7036e32d6387f714ed11a008c42fc5bf059074874dd562f07e999"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_openbsd_amd64.zip",
     "sha256_checksum": "3dfaaa63bc5b3f1c79d4100fa89b3a603057dcca329b3afdabfb0b6ac1783c27"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_solaris_amd64.zip",
     "sha256_checksum": "ecaee209578923dfbc96079b3806fc0009c18dea36a8b09bc05642eb839a9b29"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_windows_386.zip",
     "sha256_checksum": "fa3ffcfe85cca5e83faa2c95a9e1518a214a2b68b34a7627fe448ef95c50f09a"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_windows_amd64.zip",
     "sha256_checksum": "fa55cc82adc4b91fa4f718941da79c3b695aa597e244dae40b574d966bb40692"
    }
   ]
  },
  {
   "version": "0.7.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_darwin_386.zip",
     "sha256_checksum": "30bfdc9ce8529dd2a48f8377e42d91e2e9198151680888bfb9b90c4819ed899f"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_darwin_amd64.zip",
     "sha256_checksum": "e65095c09cd94d60f0a6bc470ad29b249051448533344722755cc617bdd277a4"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_freebsd_386.zip",
     "sha256_checksum": "b43ba2b79fa5a8fe3a553d18f2cb3deace7d7b79f04cef5ccf48bb980058ff98"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_freebsd_amd64.zip",
     "sha256_checksum": "11c5ca0d1a5473e6855a2cdd232def09f99650d4c396fd143cddb2736d4238dc"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_freebsd_arm.zip",
     "sha256_checksum": "e8121021eb45469a07a2786cfc93ca5aafb1455ce3d017b72931967d1fb656f0"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_linux_386.zip",
     "sha256_checksum": "933c7a0cccd072719c929fb2733db458c38c05493508b6c7375eaf608f16fbb3"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_linux_amd64.zip",
     "sha256_checksum": "a6da76d6228349855f7c503b769fb231e6b1009add5e5b2586ecb7624e9ecf15"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_linux_arm.zip",
     "sha256_checksum": "b36c0090218622ea351f671e93621890cd4e673397b7386f15390de865400e57"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_openbsd_386.zip",
     "sha256_checksum": "218b6c60426f4ad9d4df374224699470e7e36dd710bed6fc5c0d727ca4b55168"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_openbsd_amd64.zip",
     "sha256_checksum": "68cee72c0e2e3c5d8d1ba98ae3bc4e974e0e4f1eccf990f5ac8104bdcca08e64"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_solaris_amd64.zip",
     "sha256_checksum": "e08d705ee5152c04bd128dba97bfd36e47da0112085dbf826d22747650775c72"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_windows_386.zip",
     "sha256_checksum": "6bac894a144d8dc5556855726611e4f48c41f208908e2836471399c1b716b4d4"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_windows_amd64.zip",
     "sha256_checksum": "caa7ae92e05f7ea24d66dee8a2af3788d277697666fd4d45f0d7e1f7be1e2186"
    }
   ]
  },
  {
   "version": "1.3.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_darwin_amd64.zip",
     "sha256_checksum": "a4583ddee28cb4c3bcade0b093755c9894d56546753edd1765669e993b764f83"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_darwin_arm64.zip",
     "sha256_checksum": "615c3c56077fafdc68786fd7732f176330381c838b380f1695b9861f6ee2505f"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_freebsd_386.zip",
     "sha256_checksum": "e49676f53ff8042eaf1a814ca165bcb937ae3223afd3c1155cd54897b2634505"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_freebsd_amd64.zip",
     "sha256_checksum": "49aac735708131efcf8fb00b1cd1c55d341cac07168fda0fbd1f5024e3260797"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_freebsd_arm.zip",
     "sha256_checksum": "37bd0666bd947ec8913b0b682cc1102f36da7604ddec89f70a3475d0427e8c8c"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_linux_386.zip",
     "sha256_checksum": "70b2662b1c991bc5046a57a330c83a5b102a16becf53255b57fad7462bfe8061"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_linux_amd64.zip",
     "sha256_checksum": "a4d84c1b6a86de6e5aa27daf592689135dd85027d06302314b5b61f08dddb434"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_linux_arm.zip",
     "sha256_checksum": "5a205194ebf24fbad3df3f0d57f105852ca2708716d85e503af14ba5167d7830"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_linux_arm64.zip",
     "sha256_checksum": "3976ad1941bc846925e3a42d2cb61bad6b31b5d20ea73eaaaafcf9760352eca5"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_openbsd_386.zip",
     "sha256_checksum": "96212fa2bb376653e0de58a1b57e2b3ee5e14c0c4bdcf437ab630b08ef30436b"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_openbsd_amd64.zip",
     "sha256_checksum": "8475f20bf8b1550216679f8f655d2fbc6f7f9bb31939f79c438bc5c28e346918"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_solaris_amd64.zip",
     "sha256_checksum": "f995b9336788f3b53d09de1436aac0c5d8689d5438c1d58567b9327db14e4bb2"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_windows_386.zip",
     "sha256_checksum": "dfc88daf6e74ae6b94013e64e832375af407c5ec4c92f847b15b02a231b08de1"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-rc1/terraform_1.3.0-rc1_windows_amd64.zip",
     "sha256_checksum": "3a9d94b4571c1e4f11f9b18e9876da9d27bbef0fe9b14fefe59b66ac43750f98"
    }
   ]
  },
  {
   "version": "1.2.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_darwin_amd64.zip",
     "sha256_checksum": "d7c9a677efb22276afdd6c7703cbfee87d509a31acb247b96aa550a35154400a"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_darwin_arm64.zip",
     "sha256_checksum": "96e3659e89bfb50f70d1bb8660452ec433019d00a862d2291817c831305d85ea"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_freebsd_386.zip",
     "sha256_checksum": "d4ea426a82ea65da078d944e973dbd532e80071b07d8ea7b07453756b72fd399"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_freebsd_amd64.zip",
     "sha256_checksum": "a19bbb7d5e2ad8267935fcfa72cc4895d5a7d3434137fb03c1dde881bbbb07eb"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_freebsd_arm.zip",
     "sha256_checksum": "6a630f3a741116b938fd16c22dc7a04a70d45e574293e1fb4ee0cec7785a2c34"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_linux_386.zip",
     "sha256_checksum": "c6672d137f10cce0800ff449e9f976dc1fc46af44d3ca3eef97ade5748f4d67d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_linux_amd64.zip",
     "sha256_checksum": "8cf8eb7ed2d95a4213fbfd0459ab303f890e79220196d1c4aae9ecf22547302e"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_linux_arm.zip",
     "sha256_checksum": "6d90bc9e711fdaef1f1c0bff099e7d7910837562e5c61b00ebd3a6b4c88e01ff"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_linux_arm64.zip",
     "sha256_checksum": "972ea512dac822274791dedceb6e7f8b9ac2ed36bd7759269b6806d0ab049128"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_openbsd_386.zip",
     "sha256_checksum": "d8afc03c5fb101a6ca3ed842f01d24bd2079846dac1b26bbf4e45614a811180c"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_openbsd_amd64.zip",
     "sha256_checksum": "d00af7e9dd746134ed2806179587cfd16240da9056dc687460056e63da7042b9"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_solaris_amd64.zip",
     "sha256_checksum": "1f4d6ef0e97de3e4b90e9ee07e2a7bf6796d138bb488524a04f8c60e92a7d897"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_windows_386.zip",
     "sha256_checksum": "75f4cfffc0604a518ed984b809dbbddfa4bb48d25ff6b898706745c0b7d87552"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.1/terraform_1.2.1_windows_amd64.zip",
     "sha256_checksum": "dc7d9ddda5e17aacc2fe7b4dc66c9bbba2c06894f58f5f4d312bdf06adce254a"
    }
   ]
  },
  {
   "version": "0.12.28",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_darwin_amd64.zip",
     "sha256_checksum": "893050bcfc5e7445acd3a30f1500227b989b29cbd958ca64a8233589194a198d"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_freebsd_386.zip",
     "sha256_checksum": "7f443c80d4793d3181cc579ef4c37a7cf76b63b316670b1abd7598ebe421dd44"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_freebsd_amd64.zip",
     "sha256_checksum": "9d3381adc015cda84687ccc4aed7714645857f901c301d7f5d20ecc78844f3fd"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_freebsd_arm.zip",
     "sha256_checksum": "c8c76a215c7f088a547fdb318ff08a48ad45a6995a4b5b36cfcf05d01ed7c86d"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_linux_386.zip",
     "sha256_checksum": "1c489282d86b16f2d5f89f38071c6dabd948a4ca7cc4e42915e604b82564f3a6"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_linux_amd64.zip",
     "sha256_checksum": "be99da1439a60942b8d23f63eba1ea05ff42160744116e84f46fc24f1a8011b6"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_linux_arm.zip",
     "sha256_checksum": "ee1835429b681460e0adcc6f873a1d1390264647bae77151ba8211610dc86d83"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_openbsd_386.zip",
     "sha256_checksum": "c5c32dc1896188a17f93e0e7aa3f160fe842ad0eb4e4761345e728b1ac26e8e2"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_openbsd_amd64.zip",
     "sha256_checksum": "d1818c9d5414b8fa5ed9618bf00abd605b016be7812e5b71661337693a47d165"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_solaris_amd64.zip",
     "sha256_checksum": "a904775f596895ff84dcc76ffff68bebf0c115b39e21f9a612431fcadbc814da"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_windows_386.zip",
     "sha256_checksum": "a6c8c2f4d20c38f3467cbc65e3d34d446751a675f4b1bfdf6be155baf14f8963"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_windows_amd64.zip",
     "sha256_checksum": "cd524b5b6b7cd9eec9c4e49aa37cbcb34ed1395876c212f6dde84c6e57d6ce1c"
    }
   ]
  },
  {
   "version": "0.9.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_darwin_amd64.zip",
     "sha256_checksum": "8d55db3e114a72ec2cefb2e928af485c10f61c2df8121847972f73ca301fe5c6"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_freebsd_386.zip",
     "sha256_checksum": "7355c652dc31ebb55886bd2ce3bb72430edfaa1e8d35567f44da417af4229380"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_freebsd_amd64.zip",
     "sha256_checksum": "b4e3966d01e4c4026cfacec764d9a6ee278c6fe904378fb3ca27fabb3bc17a3c"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_freebsd_arm.zip",
     "sha256_checksum": "5948a5074200b512323c233efaa43428cfdc3615b4261425b8a09114c72d9974"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_linux_386.zip",
     "sha256_checksum": "31017b0054c83899e94923ac10f8deabcfa222746c012f707af316817a5a2234"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_linux_amd64.zip",
     "sha256_checksum": "77f0d01182d665f7f3c63c326aa699b452fba043c2e2f9050c4bd114f98a1207"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_linux_arm.zip",
     "sha256_checksum": "f3f05c1fdfdb11314f8a4e9909e0cb69f55c16b20e08ba91937fadcba858a7d9"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_openbsd_386.zip",
     "sha256_checksum": "608ce5b355628317730d5f9ed3006b426af3bbecf6d68b89de93f94752917e16"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_openbsd_amd64.zip",
     "sha256_checksum": "d13946ec9ecde1f6656b9deea3764b7f15a2cd8dde01acf3a1a6eb45563df0d3"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_solaris_amd64.zip",
     "sha256_checksum": "335801b1bb0203baf464638e2e0e0c5ad29945ba8e7e1ed762100c925dec174"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_windows_386.zip",
     "sha256_checksum": "285d441907cc243f9679e459ae114c22aeb45ffee39439cf7e151c6ebe8fcf68"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_windows_amd64.zip",
     "sha256_checksum": "8174d714eaa1c5bd13dd116adba62809d6860383ba90638dc330a0133dbb519"
    }
   ]
  },
  {
   "version": "0.2.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_darwin_386.zip",
     "sha256_checksum": "65a3ebbad375daabe857046bdbbd207f315c89a27ca89d1e1630b0f4c6fa03a2"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_darwin_amd64.zip",
     "sha256_checksum": "28076fa5b074d2b2457f857fe8f2182a8ef7a35c15b8c3b18a129df60790ea7"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_freebsd_386.zip",
     "sha256_checksum": "2476b8f2011dee827c8b74e75a0f4eff34c5be49ec2ecda54f4b5a860ec1dfe1"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_freebsd_amd64.zip",
     "sha256_checksum": "7caee00b360f9e2d40e96092a0b5569ffa5a7f8e2b0038df166d22cf4ec9ee26"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_freebsd_arm.zip",
     "sha256_checksum": "17be0969487f26536d91102f573dbc31a2502d00986976b3b4dd0c28c6c5c715"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_linux_386.zip",
     "sha256_checksum": "cc75654bdc8b9f04540533a762f6720e98e17c6ea9d3b59da03246ca6ef2f36e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_linux_amd64.zip",
     "sha256_checksum": "10c66a00368b8d3e9658f05b31d77654bea7871c05fd83e1cc00be127e2d0577"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_linux_arm.zip",
     "sha256_checksum": "1115b5af1e56e09c5f60a087824768ebfc9f2131f45e46cea80058a401d2bf3c"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_openbsd_386.zip",
     "sha256_checksum": "972cd25169c4683216a7e8c26abda734664a1513eef5b0af5d8bde35104001c4"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_openbsd_amd64.zip",
     "sha256_checksum": "f86d91ffbb5454fc6d6c8b8e351aeb64cbe7c9c6f15b72503af513acb5d242fc"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_windows_386.zip",
     "sha256_checksum": "fe713e6c057ca35273c7d8dda79f756ec1f23c0aa638a6333a6df436d1c3c31d"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_windows_amd64.zip",
     "sha256_checksum": "1fac6cd266d94936600c5d70ba27df9b46d780242e4b19cd5826b81f27fa7443"
    }
   ]
  },
  {
   "version": "0.7.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_darwin_386.zip",
     "sha256_checksum": "6c26b8b9a1ee07ad0cc62c7ddd2f619762bc24d6d7dcc72e373042eafc4b4c4d"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_darwin_amd64.zip",
     "sha256_checksum": "e0057e4f32e6490361611e3eb34e35f8b5314d861aa26fd9e89e1a7c4ab773bf"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_freebsd_386.zip",
     "sha256_checksum": "2a22c8310b861960e8e1f37a258a7b64e0f50e828eaeadc14602788c05addb51"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_freebsd_amd64.zip",
     "sha256_checksum": "c1b19402cbba5471d1fec10dbab91abcf6142d2cc5d51a9318273e7611d67842"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_freebsd_arm.zip",
     "sha256_checksum": "2ba6a4587b6a43ed1c8ff388f3933ddf95fc03af8f4293e96bce9f45653eda1b"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_linux_386.zip",
     "sha256_checksum": "bc8a9322d05f6e52f73306395892d82f70c7cc0d1097967aeb4028113d5ad1b3"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_linux_amd64.zip",
     "sha256_checksum": "4e985f222ec99616e8c730d737c9b400f9d73bf0c436661ec888b2406d3a6f39"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_linux_arm.zip",
     "sha256_checksum": "586f4331936d38174fd9241ad89efe2406d3368a005578e66e74f900fd4cf272"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_openbsd_386.zip",
     "sha256_checksum": "145a7d0e407918eba28472bd1e1516c9a6ebc8ba9a4d1e88414333514a21efa"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_openbsd_amd64.zip",
     "sha256_checksum": "9fd02776e67b6bb65a3da5a1f68bb3fb776d9547ff9b0b955ad5848f9ee7787"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_solaris_amd64.zip",
     "sha256_checksum": "c3adcf4d03daf112ebeee7f175ad968e128722a4874bdf705cb9394c2834f204"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_windows_386.zip",
     "sha256_checksum": "e7126ca567572351becd3ea2a770f2da56323af2a98fa2467f9e1ec3953ae021"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_windows_amd64.zip",
     "sha256_checksum": "67b6176204a3542fed3834e3e32f5d70681a911ce65e384727f5e4405d7a7c82"
    }
   ]
  },
  {
   "version": "1.3.0-alpha20220608",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_darwin_amd64.zip",
     "sha256_checksum": "f3ae274c586e3ca7c1dbaecf1c475ff70fe942de91caeb47540b39e906589d0c"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_darwin_arm64.zip",
     "sha256_checksum": "798d0b02ce243a59d4982fcc0a1d00ed9f3fcdd22cc194503866291f217047ee"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_freebsd_386.zip",
     "sha256_checksum": "65c2952f4483f682b3da813351f4b44b7e56a2e1e84e8ccf5ccf35c4ad94fd5e"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_freebsd_amd64.zip",
     "sha256_checksum": "8b37c732e10aec5f39c106436b68dd5abd5563f88fe85f644c5174b514538fca"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_freebsd_arm.zip",
     "sha256_checksum": "87f2dc430ef3c5fd8de40689c679a23ca0b640c0e36401e81bb804bfb047bd3"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_linux_386.zip",
     "sha256_checksum": "3694b911eedfbe66f685dcaec6131738b3e267f5744ada919735d280a9043511"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_linux_amd64.zip",
     "sha256_checksum": "13c7993f2e8c744ad922a9e108c3e470ebba1badb1526be38584f974062db071"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_linux_arm.zip",
     "sha256_checksum": "1627409331ffac7906839c1ce5f86c4572d4fe07996a840fc71be189c401a2d2"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_linux_arm64.zip",
     "sha256_checksum": "5ebca8858618491eccb5cc00e67cc097df2800f14c3975981b4d1393e5a8a4bf"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_openbsd_386.zip",
     "sha256_checksum": "7b103328e325d8bb9aaff19fe8d990322a686e7044e5453c9a65913fd38b2914"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_openbsd_amd64.zip",
     "sha256_checksum": "87f4d59c5b13b3149360826ade9155bc3d46482c8a5f92780e1f1103383e3fe8"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_solaris_amd64.zip",
     "sha256_checksum": "28d245c8a7028a2d8c8b9102aba9b8de7964350be261dfd617b3f369860fd15d"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_windows_386.zip",
     "sha256_checksum": "41adfe1200f1c4ed64b9749931f2d61c4d0e8324f74fb1eff649286d0220c27a"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220608/terraform_1.3.0-alpha20220608_windows_amd64.zip",
     "sha256_checksum": "dc5501564ed47439e28d6e7cd8217da237435349680def9c36428f5f247b1c08"
    }
   ]
  },
  {
   "version": "1.2.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_darwin_amd64.zip",
     "sha256_checksum": "acc781e964be9b527101b00eb6e7e63e7e509dd1355ff8567b80d0244c460634"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_darwin_arm64.zip",
     "sha256_checksum": "e4717057e1cbb606f1e089261def9a17ddd18b78707d9e212c768dc0d739a220"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_freebsd_386.zip",
     "sha256_checksum": "e5bec128884f6db261b5e4b36cce4cf52a3d98373ab6b540e028c6d95938e0fc"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_freebsd_amd64.zip",
     "sha256_checksum": "392d27a645037d5d5b8d61ccbff01707e674836f4165613e540fb78a850bea1e"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_freebsd_arm.zip",
     "sha256_checksum": "22b1b276a73b91b829a92326d5f486d806e100249ffca1b80f8d605208997f77"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_linux_386.zip",
     "sha256_checksum": "bedd408479680b5d06ddcd6efad7e1657ad693bc41c7b288f0716689e70ecfde"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_linux_amd64.zip",
     "sha256_checksum": "dfd7c44a5b6832d62860a01095a15b53616fb3ea4441ab89542f9364e3fca718"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_linux_arm.zip",
     "sha256_checksum": "c2557f3f878330a66a02de3c08ae8d0824b658c66fcb36b1eb07939bf2bc0468"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_linux_arm64.zip",
     "sha256_checksum": "80d064008d57ba5dc97e189215c87275bf39ca14b1234430eae2f114394ea229"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_openbsd_386.zip",
     "sha256_checksum": "25955cd2ad5e6f714d211973f8b5c974652d01eca63e086a3e71aa2c47279921"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_openbsd_amd64.zip",
     "sha256_checksum": "1852dfedfeb0e2147a85500365c8c63637bfcf12937b7f0179399086663c8cff"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_solaris_amd64.zip",
     "sha256_checksum": "465d82d8e14aa203d806169f3c4823086274cc8fbf9a3d6afc92689a7dbccdd7"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_windows_386.zip",
     "sha256_checksum": "e39212d075008017f04dfb2a834cecaa92c4d2645c0580b47c7a2e03e113e817"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.7/terraform_1.2.7_windows_amd64.zip",
     "sha256_checksum": "3866eb55549514ffa73b6d460807ab8172e1a063f8dc21ab9044c45668e4dce9"
    }
   ]
  },
  {
   "version": "1.2.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_darwin_amd64.zip",
     "sha256_checksum": "2962b0ebdf6f431b8fb182ffc1d8b582b73945db0c3ab99230ffc360d9e297a2"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_darwin_arm64.zip",
     "sha256_checksum": "601962205ad3dcf9b1b75f758589890a07854506cbd08ca2fc25afbf373bff53"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_freebsd_386.zip",
     "sha256_checksum": "f1f65fab8f71c0b0773283ad62fefb00d7511344e14712d51729eb341cabddfb"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_freebsd_amd64.zip",
     "sha256_checksum": "9345a279c19bdf740c49300e3fbc01c9220d82bd1c68b7f3ce97a9c4058b46ba"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_freebsd_arm.zip",
     "sha256_checksum": "d91a80399a894741c94956fbd5a486d5656994c7b860b5abe9681a4ffa6e6215"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_linux_386.zip",
     "sha256_checksum": "cb73b88213d866673c85a078d41e3385b17d6007cd51db5c7a72ab6948ad509f"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_linux_amd64.zip",
     "sha256_checksum": "728b6fbcb288ad1b7b6590585410a98d3b7e05efe4601ef776c37e15e9a83a96"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_linux_arm.zip",
     "sha256_checksum": "a28b2e9163d6fc351bc791e827b01f7a356b9d76d89fa94f94bfe6dc50756d53"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_linux_arm64.zip",
     "sha256_checksum": "a48991e938a25bfe5d257f4b6cbbdc73d920cc34bbc8f0e685e28b9610ad75fe"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_openbsd_386.zip",
     "sha256_checksum": "66c91f38ba6213150bd01e5386130f1692f366469e3ea39dbd1951d2ea9086c9"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_openbsd_amd64.zip",
     "sha256_checksum": "fb34a52b2cc41d5c66980fcf524f7f74a7e1cf0ad6839e375c2d89357ad4566c"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_solaris_amd64.zip",
     "sha256_checksum": "f31cdd00838ec2ddd27e53fb5b71e718d73a504b0f0263026d7e5291a9f6618d"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_windows_386.zip",
     "sha256_checksum": "811c2911b8e9f19b7d73a11b4858234854c870271864bb22ec21b582fd06e695"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.3/terraform_1.2.3_windows_amd64.zip",
     "sha256_checksum": "19773a16263d0873bc86b1109412abb80733524bf6ef7e7290278b2fad33bff6"
    }
   ]
  },
  {
   "version": "0.11.12-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_darwin_amd64.zip",
     "sha256_checksum": "8c1f4f975bf4bba6725e6e1cbc69c4a7764a3fb6dc6aebf2e718456eed6405a9"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_freebsd_386.zip",
     "sha256_checksum": "37fd4ee1c666141d14dcb6786d17dc06489e90d06511c754d0df1c12bbd91ed"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_freebsd_amd64.zip",
     "sha256_checksum": "1a768629f297ba076db9589988c9538a6e70cc94115a27f7940b3e0770735f92"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_freebsd_arm.zip",
     "sha256_checksum": "5fc0a6c810049dbd7188f44e7504fe232a12b2b715a2a7c77f4ef87863a38e79"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_linux_386.zip",
     "sha256_checksum": "e1b017dba7fd50dbf93e6437ae0065c8eab6808a64d069078203081c054db61"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_linux_amd64.zip",
     "sha256_checksum": "5a8c4f734ae58a3b2583d9b93c619fe7ab0e7b76c5a7d51ecc956166fb279258"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_linux_arm.zip",
     "sha256_checksum": "de9c6aba1b5c03a7b643cd3f0d6505ac6d790a216d0cd1acaf4bfb614595c34a"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_openbsd_386.zip",
     "sha256_checksum": "822ec1b5ff637b5e7bcf2c143a5ab66bd31ba426c50043484b308cef14b96395"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_openbsd_amd64.zip",
     "sha256_checksum": "591dea1266574eed7e88b3fce1c479deb2b30feb6945276af7b09cd30e0d14f7"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_solaris_amd64.zip",
     "sha256_checksum": "89af4ea457bad56a496554becc2de8b66f2a5da99fd562623d3547c49f814ab1"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_windows_386.zip",
     "sha256_checksum": "59b14213ecf25324f6f081a7044848a0d9a276397c6d2f18fbdf27d6aa485055"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_windows_amd64.zip",
     "sha256_checksum": "20b50587c42597a5675c0ecfb69f09164966650a13d8a72337eb8c2df8002846"
    }
   ]
  },
  {
   "version": "0.12.26",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_darwin_amd64.zip",
     "sha256_checksum": "79fb293324012bc981006e1527267987666dd80cff80b11f93fb0ab2e321c450"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_freebsd_386.zip",
     "sha256_checksum": "486853c9806e1f5e75cfa253db4b9194138dda9e0410f3b92a2bb486d0e0ce97"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_freebsd_amd64.zip",
     "sha256_checksum": "83e6187c43a09417519202d238193222e6750cad9226c5a8131bcb3e83f164d3"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_freebsd_arm.zip",
     "sha256_checksum": "ef5f444ab6a500e5558f95889beea01ef8ee053043f5c927a9fc9ca12c383cd1"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_linux_386.zip",
     "sha256_checksum": "7352682128af08d921fc3cad78d1c83a656d57c5876857c60bc48e3e22364421"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_linux_amd64.zip",
     "sha256_checksum": "607bc802b1c6c2a5e62cc48640f38aaa64bef1501b46f0ae4829feb51594b257"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_linux_arm.zip",
     "sha256_checksum": "8036807501c16a5b643fa68cbee78ca6b2fa3ede43b56af7e9ff95bcf1302130"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_openbsd_386.zip",
     "sha256_checksum": "8101453b7cf08356bd58fa1f099d737ffae835ec86f35b1ae61ae80d28acef3f"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_openbsd_amd64.zip",
     "sha256_checksum": "cf3fad28811968421a8ba7b95666a587e6f240d86a06c22d41de5015f0bc7cfa"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_solaris_amd64.zip",
     "sha256_checksum": "7cca1f906b3c41a31290ac75d14fb18b6ad1741379311750d9458da07f14ccdf"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_windows_386.zip",
     "sha256_checksum": "a793b82a0b61e87a4174aff54efcfdf522a2a20b21d73875e9581156f9ae036a"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_windows_amd64.zip",
     "sha256_checksum": "f232bf25dc32e618fbb692b98857d10a84e16e531e9ce5e87e060c1369bde092"
    }
   ]
  },
  {
   "version": "0.12.13",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_darwin_amd64.zip",
     "sha256_checksum": "744dfa3c4f566cabddf2fa6b3b19fab06d512f3c654c09906e8acaaaa2388cfb"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_freebsd_386.zip",
     "sha256_checksum": "1087ebd0ce7632794545763e25392e78ac482d7940aac5936dba5301a70e858"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_freebsd_amd64.zip",
     "sha256_checksum": "7815105500561ee306ffd542b23010a0d96176190c0f0022f7d0361e730661f5"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_freebsd_arm.zip",
     "sha256_checksum": "37b1d9f6782489ea80131c8a0edf31e5fe86c307d4e1618d01aff132caaa0ef4"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_linux_386.zip",
     "sha256_checksum": "8677e3209ebb40f06f028b29c6d2ce85f79fa04ea6cd753dba9dc6186b346ef0"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_linux_amd64.zip",
     "sha256_checksum": "63f765a3f83987b67b046a9c31acff1ec9ee618990d0eab4db34eca6c0d861ec"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_linux_arm.zip",
     "sha256_checksum": "7e6f52d841e962a5c18ed73cec565338937bd6170299897f6398495462e798f9"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_openbsd_386.zip",
     "sha256_checksum": "ef028883ea9844fd4fe6492c74fb6f6fec1770234e8cc2e1cbabae857f7eca3b"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_openbsd_amd64.zip",
     "sha256_checksum": "5683f475399523f1f2f714bc205426031bfc1ada32b38a98ea9f9e4c7533cb15"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_solaris_amd64.zip",
     "sha256_checksum": "e668c991170bf6126947fa2c794d1cc945a82ab1b103977b192acc094a89948e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_windows_386.zip",
     "sha256_checksum": "94b2e8d3fd881625c8a7b296a25cffebcd8bf6645c810d94cbb3cfb5e3be9b1"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_windows_amd64.zip",
     "sha256_checksum": "d68127371734d169c9a9d020cc6079a6a985ed7876b36b899b553aa8f0e04abc"
    }
   ]
  },
  {
   "version": "0.1.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.1.0/terraform_0.1.0_darwin_amd64.zip",
     "sha256_checksum": "309aed0ed61586e2682f58b77781f8e9805745a5edd1aebcddf883c9f624a0b9"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.1.0/terraform_0.1.0_linux_386.zip",
     "sha256_checksum": "7925b3f1f759074c521dce21f836405f1b88347adf45233591a1d8f1a4985460"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.1.0/terraform_0.1.0_linux_amd64.zip",
     "sha256_checksum": "f9236171f522c9bd1c99c9f5e0ecde7f8ba881afd9f5aef8f70bdeb0a209fc22"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.1.0/terraform_0.1.0_windows_386.zip",
     "sha256_checksum": "ab7393f86fad7f870d398aed6984c88001b7cc36366973489cf46701dc90cf3e"
    }
   ]
  },
  {
   "version": "0.14.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_darwin_amd64.zip",
     "sha256_checksum": "eda23614cd1dce1e96e7adf84f445c2783132c072fbd987f1f8858f34c361e41"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_freebsd_386.zip",
     "sha256_checksum": "3ee858041aec16f6594de04b3bc0dffd983863de1b5864c959407a9423484bdc"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_freebsd_amd64.zip",
     "sha256_checksum": "cf5e05b014ecc4046f0010e5a2acae808fdab4944462e2078c95940cd9d8595"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_freebsd_arm.zip",
     "sha256_checksum": "99109533473961a28d2888b403300570869da0d6101e934dfbaecce579ff2ae2"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_linux_386.zip",
     "sha256_checksum": "ead9667ec1a9d70f61f2a493694031830cef35b2638b027ba260e2ff374326db"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_linux_amd64.zip",
     "sha256_checksum": "aa7b6cb6f366ffb920083b2a9739079ee560240ca31b580fe422af4af28cbb5a"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_linux_arm.zip",
     "sha256_checksum": "a014681c881d751114c9f2bfb5250e73d904195742c6bbe56d7eddb51e5ceaa6"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_linux_arm64.zip",
     "sha256_checksum": "c30d1dff656f8c77497c3e5ef5324f1896356254c3437be58f21b962d307eeda"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_openbsd_386.zip",
     "sha256_checksum": "6dbaeab15692d7da5a7ede3548ab1c412fbe23a7d6953b7c7162ec1b955b2f11"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_openbsd_amd64.zip",
     "sha256_checksum": "37eee016a45e7fd0dce74715cfe3a52003054d58d9e6ecf0e5cf263b49605507"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_solaris_amd64.zip",
     "sha256_checksum": "8fbedfc82daf2e77ee7970d39b2fbd4b6903769885f8aaa2fd34c38c73208955"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_windows_386.zip",
     "sha256_checksum": "784e838aad10ba13c4c50b0a738591e914ee5f6b1a092da89ba09e4e5b1e1877"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_windows_amd64.zip",
     "sha256_checksum": "f9b13a80c8926a773ccb052da98cca6cb45863ba4140466af7e24c0b61a073c2"
    }
   ]
  },
  {
   "version": "0.12.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_darwin_amd64.zip",
     "sha256_checksum": "3c5b0aa8f3acf477a5d5ae997174bd16d49bb0789915b5a40a6deb39692a5c8d"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_freebsd_386.zip",
     "sha256_checksum": "4aebbc80f23b81ce9ce73417be9e15fe755a86e53c08806bbe9eada34f7821b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_freebsd_amd64.zip",
     "sha256_checksum": "a3b0b665459fc034ed49c0c93c447629c8341192666addf7500c4f123d7b644d"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_freebsd_arm.zip",
     "sha256_checksum": "6b44417ee601a1e0309bae64ddee0349c92798aafc6728f82b69dbcc25f442f0"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_linux_386.zip",
     "sha256_checksum": "5b153e04d57da499866434d44ffd34c81c6c776a520db3f9e7582349513cee42"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_linux_amd64.zip",
     "sha256_checksum": "c9a30d3e3abf751b3b3e323897e9c7cb411d5c4bb7473a3284a2a2b4b89f93ed"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_linux_arm.zip",
     "sha256_checksum": "55414bdb1b9fe4e773f27c89a03a3a6c869d29b370fd6750d54d8814215dcd56"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_openbsd_386.zip",
     "sha256_checksum": "968f5ed83c61f3368d3f5a1ebaf84bd8129f8fbd054201e56e5d543cdc5d41fb"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_openbsd_amd64.zip",
     "sha256_checksum": "fa9c0f4274cfe87d6acbad1c552e4d066a7a54c3f476d9db31907ff4ab04bc4d"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_solaris_amd64.zip",
     "sha256_checksum": "c09131582b780c9170108859d600b5fe810687b9e5fc06ce429943139d8e53d8"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_windows_386.zip",
     "sha256_checksum": "b7ed313163268b0eeb4ad7fd43de2af2d0073703c62365a279d752a32007ffb7"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_windows_amd64.zip",
     "sha256_checksum": "54969c83bf6ef5e9b675575b59f70e9b64c2ba2a90958727dc7fbb8dae845729"
    }
   ]
  },
  {
   "version": "1.3.0-alpha20220622",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_darwin_amd64.zip",
     "sha256_checksum": "81cf0bdef5d6cc90b3f9fb498abb53e029fe00c570bf6ba0b00d2d02ea60eaed"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_darwin_arm64.zip",
     "sha256_checksum": "281373c9f7cef1b488436d5dcacbb578f056a655db5302a4cd36760fb538d002"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_freebsd_386.zip",
     "sha256_checksum": "fcbcc5f7b1204e7e7cb9cdaf641a828ddcabf6311f5962017c075e6adb5538b1"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_freebsd_amd64.zip",
     "sha256_checksum": "3773edaff630d7cdc3f0b06c26eece5440ffcc86d263470a7f3ed2dda4dda4e1"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_freebsd_arm.zip",
     "sha256_checksum": "ec41d762891f818f240f94609d8a783a65217eb13ae65f52ca2fec093a7cf231"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_linux_386.zip",
     "sha256_checksum": "e50d86e51d977739d0d19721400e4f8c5a9ae6a39c32b50244a050fa2f1288ba"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_linux_amd64.zip",
     "sha256_checksum": "23a2cbe69147b5052c141e4dbe2c71c326f11a70ef70a6511d5ee4828a229c8d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_linux_arm.zip",
     "sha256_checksum": "5e8657ae4d6d1cdb4d99f1d44d38da65b1242ea250034f3fd5177bc5df1fcbe7"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_linux_arm64.zip",
     "sha256_checksum": "62eb3f95ade32f1bf3d78a0f248cbcc0035f366fe56b48d47d36c39bacf0ce7f"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_openbsd_386.zip",
     "sha256_checksum": "a21e22cdd7ef158200c121d70558a209e8123d1d06a1dd1ffd51d306db1f7042"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_openbsd_amd64.zip",
     "sha256_checksum": "ec5d214ca795c1a56002f7a122dcdc0eed2d8af95dfc7b18bddf25737fe647bc"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_solaris_amd64.zip",
     "sha256_checksum": "c896c8aed50c4fdf73c393ab8299c500d7c73660512b5ebc67ca079b7d36e03"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_windows_386.zip",
     "sha256_checksum": "6ec9f89c8e6a753f6d8d15a85fca4da4468c9860cf55dbcdbae8695a34978201"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220622/terraform_1.3.0-alpha20220622_windows_amd64.zip",
     "sha256_checksum": "96db1771d93221147a46c08e8f9eeabeb02b3a6b262b7d078b078ac112ca1af7"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20211029",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_darwin_amd64.zip",
     "sha256_checksum": "ab58023b6daa9fd6e5f07282ab48ebd6ad0c5d240224c237f6bf8355fb91574e"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_darwin_arm64.zip",
     "sha256_checksum": "3cf2ba01b18d20fe80ac327be61270aca61fc1e50a884734226639122b1738d2"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_freebsd_386.zip",
     "sha256_checksum": "676a70b8c2072f48da91416414285b1e4924b784068578daca39e0e1a041b0d7"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_freebsd_amd64.zip",
     "sha256_checksum": "360a82be3fbd4465135cb425aace017dcd4fd1c8acb8a07c357361014d071b82"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_freebsd_arm.zip",
     "sha256_checksum": "a97998450fe27bf7505cf86be732006cf9a1e36728fc7f5e9bef60f6c030fa33"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_linux_386.zip",
     "sha256_checksum": "3362627b763bfe4ee06502c879de5431a6684c2fbde49a4802ae1030210877ca"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_linux_amd64.zip",
     "sha256_checksum": "2755bf95146d872b45694c5783b4d0fc28806f86de57bf57acea7049b406758"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_linux_arm.zip",
     "sha256_checksum": "b44e192f7e26aee29130650bd433e29b5eeb53a7ebdb42bb29c2cfc13b7df752"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_linux_arm64.zip",
     "sha256_checksum": "294e1a252f492db24a34ee9fee6610d2ddaf8b1652ed63221beb4d2e7681a6b9"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_openbsd_386.zip",
     "sha256_checksum": "36ee72b47fbbb0dc502ae911d84297c3236b6ea03c790db65d15241761697502"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_openbsd_amd64.zip",
     "sha256_checksum": "6091e60dc7f838135c16c8b2034f6611c8154442544f13fddc451df4594f4a48"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_solaris_amd64.zip",
     "sha256_checksum": "ff6780868fe82d9aaedc463efb1db549b6a968efe683ee81492707377b6f4c17"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_windows_386.zip",
     "sha256_checksum": "ce5abd1c56e23026347ce1e68f86246ca3751164a586a022dd720c3f3ecd3cbf"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211029/terraform_1.1.0-alpha20211029_windows_amd64.zip",
     "sha256_checksum": "faefa64237991b3ea1dc60a1d5d60ae4f50c056c4bf01c6d8f75da89daa9dc3a"
    }
   ]
  },
  {
   "version": "0.10.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_darwin_amd64.zip",
     "sha256_checksum": "6d7c51b8b8eee81b07c6b594077e0af95be518ed88b312bd3989c37b2924c2e6"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_freebsd_386.zip",
     "sha256_checksum": "4a7638e045c5157236bab32737b74d61f513580ca74a7a6a161205be3036bc7c"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_freebsd_amd64.zip",
     "sha256_checksum": "7910fa8d741e28fa46ea99ac5305c439b7c37f81d16738aec2d40a505c8b141b"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_freebsd_arm.zip",
     "sha256_checksum": "b85a313297c42b06db61ceb5d7ba3c624ca20643ceb0aff76d0336a8ac70cdfb"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_linux_386.zip",
     "sha256_checksum": "69a5601f1c15d55d9309c45913f9cb9f010a0fd4a7cd0f7a98ed03e58c7eb31e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_linux_amd64.zip",
     "sha256_checksum": "f316c6ff8b2abe257250d19cbe0e3cf745dedfa67b37bb4afaf95e0291efeade"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_linux_arm.zip",
     "sha256_checksum": "ee768c98e6504177ad3c4dfa29604632015335cc62529473a01b71748a42669c"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_openbsd_386.zip",
     "sha256_checksum": "2401ea2b7eedec78f33a380012b684939a31485b1f3be2ea2734756123bb9cb3"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_openbsd_amd64.zip",
     "sha256_checksum": "94c470fbb3b0bbf6eba7ac5d19c8270bd5ab8eee74a73a7b82fb5946399e251d"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_solaris_amd64.zip",
     "sha256_checksum": "128d9f74a195659d30dcfe4f9601d055b34c888e34a196342e5a35d6ef545063"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_windows_386.zip",
     "sha256_checksum": "28411169c231625ecb25a291907eba68345b8bab838813e26f1ba938275c8abb"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_windows_amd64.zip",
     "sha256_checksum": "df0e14d5428adfebd71e7cc12ecb05e7b770febd81e8444ee02084ace0d572d9"
    }
   ]
  },
  {
   "version": "0.12.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_darwin_amd64.zip",
     "sha256_checksum": "e0afcf6f6401e9eaab0be588b55b5226549253854acc1d0cde331b8ca54727e0"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_freebsd_386.zip",
     "sha256_checksum": "3a6e74da436f697ce659675960bcbbaab9c33632e20fd972d87abe204ef990f7"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_freebsd_amd64.zip",
     "sha256_checksum": "e932552f9e25d02a202b61f3402529d12d86f178214b3aa1e987f0dd678f75a"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_freebsd_arm.zip",
     "sha256_checksum": "2ce9a11513e03eade549ea50db2261f71b0b981b1962179cd1e3a6af61869418"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_linux_386.zip",
     "sha256_checksum": "3604424599cfc4bf14bae03e75512023571b30195a13e59cbafabe657baf20cf"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_linux_amd64.zip",
     "sha256_checksum": "babb4a30b399fb6fc87a6aa7435371721310c2e2102a95a763ef2c979ab06ce2"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_linux_arm.zip",
     "sha256_checksum": "dd723c2056882a5dfd90563098bb26f21218bcb64b710be0f07dc6f74071c168"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_openbsd_386.zip",
     "sha256_checksum": "601f36a48dea1e9b30e3311b18b8091b6b6c408b74ce4b5a58d688d8ff213559"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_openbsd_amd64.zip",
     "sha256_checksum": "7484909eb1dcdf726f33d73493b69dce7f307e10851bc01ca8f84adc54a327c9"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_solaris_amd64.zip",
     "sha256_checksum": "7913de2a398164f0e69b17f45c0a1a130507eeb07d2c7072b6c4e39ca4ced5c4"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_windows_386.zip",
     "sha256_checksum": "6d3d82660f463895052f37c52163445433744c74851152f163d910ec344f484"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_windows_amd64.zip",
     "sha256_checksum": "6b720bf935f3c121430612ca10850dcd457f4d74d2b756baadcec50fb6feac20"
    }
   ]
  },
  {
   "version": "0.11.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_darwin_amd64.zip",
     "sha256_checksum": "6514a8fe5a344c5b8819c7f32745cd571f58092ffc9bbe9ea3639799b97ced5f"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_freebsd_386.zip",
     "sha256_checksum": "807e1c0001a3b066a35a78cfad4a4deccb825d2832e237ebfe0c819a6648b785"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_freebsd_amd64.zip",
     "sha256_checksum": "8884917de66fde86937c45ad62e6f8c8dd4d0f05fba84ad2ecdf3e2be72ad851"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_freebsd_arm.zip",
     "sha256_checksum": "aaacf7fdaf2c64e43eee30b1d469302077bc66a0795f090336ea0ebf3e8fc430"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_linux_386.zip",
     "sha256_checksum": "62749ecc930dccc0469a32f652534a6647f52d7c0c924106e15491d518acc190"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_linux_amd64.zip",
     "sha256_checksum": "6b8ce67647a59b2a3f70199c304abca0ddec0e49fd060944c26f666298e23418"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_linux_arm.zip",
     "sha256_checksum": "bd0bb7ede0864b7c71f0b575ba79f5ea5d656d8d22d554dadb544e781cf66788"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_openbsd_386.zip",
     "sha256_checksum": "1662c7c14aa414a81ce53044ec18f6fc0188965b7ad97155d4cea60fcf19518"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_openbsd_amd64.zip",
     "sha256_checksum": "d1a115f78963f4d38192d23e686135c2a15c8cbe581cf2bead9322acc489770b"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_solaris_amd64.zip",
     "sha256_checksum": "8403d87f17122bed8d0964dc73a546da2e100712924f69ef6129d0b6e378812e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_windows_386.zip",
     "sha256_checksum": "4e4b50de1143d72d348e6300d78aed596acae09714516d697bd4f0d9443aa7d3"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_windows_amd64.zip",
     "sha256_checksum": "5fd003ef20f7a6a85ced4ad30bf95698afd4d0bfd477541583ff014e96026d6c"
    }
   ]
  },
  {
   "version": "0.6.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_darwin_386.zip",
     "sha256_checksum": "74f9370b3464a64b7e080913375f3eb2b3510f072e1725d59a487a4468eb4179"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_darwin_amd64.zip",
     "sha256_checksum": "71fd8ff20f657a4c7d82794756d55c55b0686516a8253356b8edd1a728230577"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_freebsd_386.zip",
     "sha256_checksum": "76ecb40dca75e9df78b2c99f363b45c6c52d61408e7824922fe090a748e0f76"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_freebsd_amd64.zip",
     "sha256_checksum": "783f176a631b51ba7a39a534d554397a30d28a3873367dd31fe827911f14c7bd"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_freebsd_arm.zip",
     "sha256_checksum": "9679549e15bbd5c95b530498193de2db66552402efe450d44cb7c9984c9203c4"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_linux_386.zip",
     "sha256_checksum": "5260aba337b536e2dbee57f7df02b4b41dc3cb4610d8603590ae2d6ae5eec59b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_linux_amd64.zip",
     "sha256_checksum": "fd61718820c3f2334276517a89694cebe82db354b584ea90c376f1c6d34bb92d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_linux_arm.zip",
     "sha256_checksum": "6ba5d99776d5bc1accce7817da4f0093a1871dba56cb0d797b962f919814904"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_openbsd_386.zip",
     "sha256_checksum": "dfbba500dfd628efd889bd11b58b09b9fd78e7c2a9351200d36f865688dddfc0"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_openbsd_amd64.zip",
     "sha256_checksum": "6521d4123dbde3a96dc954f06bdc936811607daaa60badbd039a45bc87bcdc49"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_windows_386.zip",
     "sha256_checksum": "4a2fbb7b5dd7ad6400f853f24db2860cafefdf319c87d559355121cd180739d4"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_windows_amd64.zip",
     "sha256_checksum": "357dd1df7443fa1078747e123dd56abba793ebe47b3670556fb11a2547ad6750"
    }
   ]
  },
  {
   "version": "0.12.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_darwin_amd64.zip",
     "sha256_checksum": "d97db2217c6050926eedf517b7b0427b1b5f1bda989742cfd33d8fe56c95bb05"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_freebsd_386.zip",
     "sha256_checksum": "72cded29e9b93c9dc92eea7205d34836c1c518490b9d2c3c9e21cc495f7ad8d8"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_freebsd_amd64.zip",
     "sha256_checksum": "56fe5888fff713aefd51b8b0026af20c3304f57435b99eaf25f9f64854849a4d"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_freebsd_arm.zip",
     "sha256_checksum": "e8a986a58bc9788be1b23cf07a313c35658c42d99531d52724b5c1b868631d16"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_linux_386.zip",
     "sha256_checksum": "864abab4637f52d569c8cdcbde378badd60d28ac54a315cb0703596837904fe2"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_linux_amd64.zip",
     "sha256_checksum": "2215208822f1a183fb57e24289de417c9b3157affbe8a5e520b768edbcb420b4"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_linux_arm.zip",
     "sha256_checksum": "3a91d603e381dd38e37a2aa7339dfe8826b30514dd6bbce24ebcb8e6a3bbec22"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_openbsd_386.zip",
     "sha256_checksum": "6af7a57f0b74a9868ce3ae87f5d81e9b7422a709686e480d6979998b5091e3d0"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_openbsd_amd64.zip",
     "sha256_checksum": "3ba6dce00091d87bf7bd15321b62d47044303c20e60277e5a55495a654fa6b9f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_solaris_amd64.zip",
     "sha256_checksum": "de8d25deef682e3bcd6d4571a893878d2f62aed793a7ffeb7e637727b4c5fd61"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_windows_386.zip",
     "sha256_checksum": "7d950a1582c3c3773ae1c2376f4ac6792a218a76d9bdeb4012fc04f41da0476e"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_windows_amd64.zip",
     "sha256_checksum": "985fd7ee23030d43dd73fa649522e95a9dcfed7d29b8d42d37e92993c22b9b76"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210811",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_darwin_amd64.zip",
     "sha256_checksum": "fd2c4c8de79d7f1de30450f568996324ce68b9f07474eec14ada5112d7f80f5e"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_darwin_arm64.zip",
     "sha256_checksum": "de1719528bad276ba987f0d41cddade327555d7626a9b32a7745e16ce02b53"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_freebsd_386.zip",
     "sha256_checksum": "4d7a1de71496b73aab0bb4fb92b13cdc8027af9f9a93a923a468d7219a349e7a"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_freebsd_amd64.zip",
     "sha256_checksum": "7c3fdc2d7b6ac38a983df306c0b981aef882f2e9baeb3e46dc1c66648a097612"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_freebsd_arm.zip",
     "sha256_checksum": "66fd86ff5628b98eb89674442e39b18285a890fa5ec56f79d40e47365957be60"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_linux_386.zip",
     "sha256_checksum": "8d3d411642ad2a840cbd885ccded678622e28be7444c95ab62a7daa66efc9e6e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_linux_amd64.zip",
     "sha256_checksum": "cbee56500a6aae48c10f4765bd056863a213a1b860c2d5e5980587aa41a53d79"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_linux_arm.zip",
     "sha256_checksum": "30911fadf0a53e3d309a8ae39eefea51c30d34eeb92e51609e8a4961112eb1"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_linux_arm64.zip",
     "sha256_checksum": "4ee4e06bd66d56bee25eb0003b5186489b310c21038874fd1ea14e9be9a3b49a"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_openbsd_386.zip",
     "sha256_checksum": "8bdd666d65fa831bc14e27722eee2bbea8f2dcd98fc8e54c0872bfe64661efdc"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_openbsd_amd64.zip",
     "sha256_checksum": "183293b304db804143196a07a070086cef3e95984b90ac9c83f809a0d6d823a3"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_solaris_amd64.zip",
     "sha256_checksum": "7b271fc8930c4207ceb4062dd518362d02568ee80d47c3c85dafdf7004fb2f35"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_windows_386.zip",
     "sha256_checksum": "dfd4fada18bbf25e4e2f1e01872f966853a89ef7f7acc3cc924e307c5c39e2b6"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_windows_amd64.zip",
     "sha256_checksum": "7feb87d778ba601e7beb5cd02705b556210ddee51420d7aaf58cb6daae83feb5"
    }
   ]
  },
  {
   "version": "0.14.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_darwin_amd64.zip",
     "sha256_checksum": "30115a2ee5f61178527089d8e5da20053927b364b08dc7aee6894a162ccbd793"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_freebsd_386.zip",
     "sha256_checksum": "32d054ff869990b67fc36e871cd1601091c809f6a68f9864c9c41c6b3173c357"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_freebsd_amd64.zip",
     "sha256_checksum": "d4262e2f128fb0443b5d8e321b9b3f5347a9493cd1f7c457e727bba3e260ee08"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_freebsd_arm.zip",
     "sha256_checksum": "f95a6a77a7e2b30553ba3d0eb6a1209c0350dca4a575c67e66c2a7741ca2f05"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_linux_386.zip",
     "sha256_checksum": "3eabff7c64323684e083fd34cf4f9824992589070c97a2d4208b0a0e4801f559"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_linux_amd64.zip",
     "sha256_checksum": "4c4c6730374f25bd70e61b83250eb52f39e340188b0f0216f7243b90396ba8b6"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_linux_arm.zip",
     "sha256_checksum": "3a698865df13693d12a6da8a6e12ae38b59e728e84151858c78cdfe3b8974ca9"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_linux_arm64.zip",
     "sha256_checksum": "a3186e8b9aef8c8add897430622804563df53e225b49f69397de20b059fb8f66"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_openbsd_386.zip",
     "sha256_checksum": "917975b050708e29419f696fea301a0dc2df9653df92f6f4bc232e230b951951"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_openbsd_amd64.zip",
     "sha256_checksum": "c83ee477e4269f4e95ef1bfe391fa2de88cf9c4461ea790dd6ff370817eaea7d"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_solaris_amd64.zip",
     "sha256_checksum": "4c2b90337e5aa2e65090cccff011a62405be8ce6bff522b780ed18af71dbf06a"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_windows_386.zip",
     "sha256_checksum": "5b0d17dc517c6fbde85d29292d271bd3a206679b45944223c35c728d09a8204a"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_windows_amd64.zip",
     "sha256_checksum": "5f02874f0d030e2c76176f4668e28589dd849bafbbd5179a4d83077b516a2ce0"
    }
   ]
  },
  {
   "version": "0.15.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_darwin_amd64.zip",
     "sha256_checksum": "27d5ae2431240dff928e6483170b570782a8dd1a2b7c32ce1793097af1acb9bd"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_freebsd_386.zip",
     "sha256_checksum": "c74e27d96047b8bdbfcf429cb9f7df53fba0c925cbf1a5b9171fa14f89b1ed8c"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_freebsd_amd64.zip",
     "sha256_checksum": "2827929e870ef3cccc9fe244e1e0d561792c7e3371744fd45198cd868da4215f"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_freebsd_arm.zip",
     "sha256_checksum": "3cf86ccb7127e4dd64d6624273628058987d85bdf32dcca46ca9fd92f2d099be"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_linux_386.zip",
     "sha256_checksum": "155d6844ae554021f8ca58b57f1e6dda7d88b6cf1d02113f3a3480e178c46b9"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_linux_amd64.zip",
     "sha256_checksum": "3b144499e08c245a8039027eb2b84c0495e119f57d79e8fb605864bb48897a7d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_linux_arm.zip",
     "sha256_checksum": "3122e7ac6353d488d766ac84379139e3d14f564fbe61f4b069f7584ca9c29e01"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_linux_arm64.zip",
     "sha256_checksum": "bc5a9d734010e55fb374a4cab9eb0539139d34fd84f58d2242573f275f67fc13"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_openbsd_386.zip",
     "sha256_checksum": "ca3cc8f1ff44cf8e73d1ebaf1c28ec1eb0fdf0b1962fa6b01f9d2452f96be615"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_openbsd_amd64.zip",
     "sha256_checksum": "7e6dad933a0dc127c89febb8a647d0ad2b96a7d88314d2ad7f8c04aeb0a3c952"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_solaris_amd64.zip",
     "sha256_checksum": "a023e319be1e82344ef21d298b689633d64ac60c4884a8c1538cd6ba75b26a99"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_windows_386.zip",
     "sha256_checksum": "dd095b70975179fb84f785ac7dec3ee9bbfae296432723661dec9eb3386780f1"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_windows_amd64.zip",
     "sha256_checksum": "568c237fe533681d2f7acd67554938c0081c5ce3e82de08a92eaba29ee36397a"
    }
   ]
  },
  {
   "version": "1.2.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_darwin_amd64.zip",
     "sha256_checksum": "d196f94486e54407524a0efbcb5756b197b763863ead2e145f86dd6c80fc9ce8"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_darwin_arm64.zip",
     "sha256_checksum": "77dd998d26e578aa22de557dc142672307807c88e3a4da65d8442de61479899f"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_freebsd_386.zip",
     "sha256_checksum": "5bb0f99344d79d51a2c9e0e714a8cd864da9f58424ad3653943b681d570593e8"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_freebsd_amd64.zip",
     "sha256_checksum": "729c656cfcc26ab5e476a5d71316b8cf5b9dfbe9743bdd8102c3049550665756"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_freebsd_arm.zip",
     "sha256_checksum": "fa4c7a5ca24a43164cd56c2230cd12578e4986d75422aff30effc71b1b5f64b0"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_linux_386.zip",
     "sha256_checksum": "8dfc636d2749d59bd2df54a0612454c72254a491e138f823f5ddf26732105d1c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_linux_amd64.zip",
     "sha256_checksum": "281344ed7e2b49b3d6af300b1fe310beed8778c56f3563c4d60e5541c0978f1b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_linux_arm.zip",
     "sha256_checksum": "3c9f27b88fc5a56b7e7494e067e3df2136b62f726fe4986b2404d38fcec6f824"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_linux_arm64.zip",
     "sha256_checksum": "544420eb29b792444014988018ae77a7c8df6b23d84983728695ba73e38f54a"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_openbsd_386.zip",
     "sha256_checksum": "d3bf669f4313c2a55a0a0cdacc3a0a9c6967c1c77628ccd10350e76965898ea1"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_openbsd_amd64.zip",
     "sha256_checksum": "fb5bc3e9a76ee4c0c49568d0f9e52695d1ab9945fb50acb3a2e928146c56d4d9"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_solaris_amd64.zip",
     "sha256_checksum": "be8acdf988dc8ef2e90f5530d29b05f578cd6b9b153c1551c0b504c291abebce"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_windows_386.zip",
     "sha256_checksum": "bbbe1e05a4f0226198a143d9f6bfb1678720d8c6eee455eee52117fd4d138217"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.5/terraform_1.2.5_windows_amd64.zip",
     "sha256_checksum": "94e15cc757b6542a0df436004146e81a2f80d9bba49b5b2f7fe5b91f27a19b4b"
    }
   ]
  },
  {
   "version": "0.10.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_darwin_amd64.zip",
     "sha256_checksum": "8170d52bd55bd80744aacd96ae8d87b39e29ed3d2d2853c9cb66ca62b5e295c6"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_freebsd_386.zip",
     "sha256_checksum": "ff937433dfaca0ffc1c79f55c5e6299828a23e5b25ee9b7b9a1ff3d3fd4c7181"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_freebsd_amd64.zip",
     "sha256_checksum": "9de2fcc51b979859ba20f3109f7be897961da7a23f2b48a9dcd61f423b269e7"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_freebsd_arm.zip",
     "sha256_checksum": "e0a2c170e475a5233cff1c22c4da5b73ec7d1859a6bb7a365ac08c3da040b8a6"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_linux_386.zip",
     "sha256_checksum": "234b30fc7e52281474e11f3a3e87752c26164114d9dac0a812bc1a49cf933969"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_linux_amd64.zip",
     "sha256_checksum": "9e14c850054faaaa3f40a1c28d8991fc4a42f0dc2d53564d8f60b12243c7ae5d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_linux_arm.zip",
     "sha256_checksum": "1b7641978b6b2cbabad20ee4b9284101e2e47d7266fc0ce6dee2ed963406d1c7"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_openbsd_386.zip",
     "sha256_checksum": "94488df15b5898002aea2f37c181a52e548ff017c8d40ba5c7e884bc7da50afb"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_openbsd_amd64.zip",
     "sha256_checksum": "f505cff7554217270c0e1281ffea89a70d22a5e003dd434468755908eca789e6"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_solaris_amd64.zip",
     "sha256_checksum": "8dd5b865c3d45ad25fb3a50a2718a32537e52fee264d38d93946812bfde8c0e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_windows_386.zip",
     "sha256_checksum": "af350ed57730d186e9f6dfb909c384bcccc73816726ffc4b02fbac18bfd9a5e0"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_windows_amd64.zip",
     "sha256_checksum": "b7f0004d8110fa0746010e55e6736d96658f84a99ed07458aa1557d8acb8977e"
    }
   ]
  },
  {
   "version": "0.15.0-alpha20210107",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_darwin_amd64.zip",
     "sha256_checksum": "91d1c9424968c0efbc9ac5958e14bfe103981c885da1ace12114288884b8c855"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_freebsd_386.zip",
     "sha256_checksum": "f5f17ced7d5d84b429bbe08e21a82db5d81e57211c28e504a662ed49e526aac9"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_freebsd_amd64.zip",
     "sha256_checksum": "aea650320ba730a2ef66c83abca760159f38c41b2ff888e3add731b53dc4f426"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_freebsd_arm.zip",
     "sha256_checksum": "11ed1976b356861d6551c0c3f67add392bc1b58707e993d2c335eb474d9b99cc"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_linux_386.zip",
     "sha256_checksum": "b506c9e02cf2cdd21b79a0206339d1500165fa302bf3826e4e14b72b608122b9"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_linux_amd64.zip",
     "sha256_checksum": "4dacde757b7fe81fdda36779cce5cd9d5123d4fbbec546551f9d2dac71dde39c"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_linux_arm.zip",
     "sha256_checksum": "60a1fed9b17237ba436e7830a43791cba637413045af00ca2e79d12f3ac6721f"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_linux_arm64.zip",
     "sha256_checksum": "35bcef6df77f51e6e7db1a9d9d7aa3af82e76da7b96d11bdd56e9e758b195e4d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_openbsd_386.zip",
     "sha256_checksum": "6da5e6c047734c5f009eb4e0a39cdf56be9fe0421a0b897d82c4409a98ab8a0"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_openbsd_amd64.zip",
     "sha256_checksum": "f2b0a6b203b6b0631f8dd92f5e8fe8a79feb89a369b97dc06911da9008f8e74"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_solaris_amd64.zip",
     "sha256_checksum": "d1d42c719768d589a5e6535abe16ba8d4a869bf507915a34fe6e44902e259853"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_windows_386.zip",
     "sha256_checksum": "eb90c0e7a31d91a688dc8c5afd11e3108a6e3565becd91748483d0138a1aacb9"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_windows_amd64.zip",
     "sha256_checksum": "a4f98805f71089344473f46b9ac1e7ddfd33f5b6c8db1ecc7d22451d09ab0c38"
    }
   ]
  },
  {
   "version": "0.11.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_darwin_amd64.zip",
     "sha256_checksum": "edbdde7ca769a5c7ca1c048bd5729b1f70d556b4ee61287dff5057660bc1f64d"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_freebsd_386.zip",
     "sha256_checksum": "1146672ba59a06f94875e1912be603d8b3d1e32146f9a22209e61960107db06b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_freebsd_amd64.zip",
     "sha256_checksum": "ebe866038bd971ecb930639ab1581655e67e16b6a713370fc75bdbbde118e6ad"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_freebsd_arm.zip",
     "sha256_checksum": "7482a7f0281fa21d4d74284df77479fe31232ff5ef1982f9b370394aaf4c6171"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_linux_386.zip",
     "sha256_checksum": "fb300a0ec351db96b84b48f199c1df39b6546752f8ac62ecdcc48fb391d3f48f"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_linux_amd64.zip",
     "sha256_checksum": "aed5c7388a3c54dc816986903d4dea32e182a002d746295e1016f6db741f472d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_linux_arm.zip",
     "sha256_checksum": "2ae8de10ea99e9d1931cc9e469fabd827ff9790555826017c42d850719e51d09"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_openbsd_386.zip",
     "sha256_checksum": "e474cb22e0e66af79246a6a5cf3d9f2fcad8a408fa6dfb23bd0ca8014d16fb37"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_openbsd_amd64.zip",
     "sha256_checksum": "c84bfd7d1a32fc5d569d907b6b3286ccd118d247e52ccf4ad31e2e9a18e25770"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_solaris_amd64.zip",
     "sha256_checksum": "5daf7c65dcca06b2b902671c461e9d6bbcadd19d17f701c4fd3291c77d0dcea9"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_windows_386.zip",
     "sha256_checksum": "f199e81c7f339c418ed7b0dbc892082961958d7fed5e0213d3ad295296ba8239"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_windows_amd64.zip",
     "sha256_checksum": "8002f35d514a8fcf616e3be3de16076ff737adc01ee7dd0f36eca13970990749"
    }
   ]
  },
  {
   "version": "0.13.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_darwin_amd64.zip",
     "sha256_checksum": "d16d3094b0f9f56d7e05b4c09a923141a483f51e58613ae64507b0f7ba45bb34"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_freebsd_386.zip",
     "sha256_checksum": "d2483da0d4acd2ea73ce17e8d4cf2d5ddecfefcb955dc6182df7f4d5757d5f24"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_freebsd_amd64.zip",
     "sha256_checksum": "de6e178a211f009615b3896e619aafd5c3164a8ce13cb6ff67bc951d87e90bbb"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_freebsd_arm.zip",
     "sha256_checksum": "d23a177fe7030ff1b1ed62e80b3ec48442c6f323a88965b0fd52f3a69ad8febb"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_linux_386.zip",
     "sha256_checksum": "16225f17149505a42751599acd8e12d01690324f680e7fa2099afed581876916"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_linux_amd64.zip",
     "sha256_checksum": "a92df4a151d390144040de5d18351301e597d3fae3679a814ea57554f6aa9b24"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_linux_arm.zip",
     "sha256_checksum": "28f90802515d2ef8468eafbe4b0125fe812ab383d7707429adb2c8f9ff8aab7d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_openbsd_386.zip",
     "sha256_checksum": "cdea05b679ea1409aae15b82f0fcb3d49d5d0eadccebb0ca37945c7ded7e9685"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_openbsd_amd64.zip",
     "sha256_checksum": "663d9f7445dc4c4ed2accafde473b73a10c88b35b0159f9dda43226edb155c22"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_solaris_amd64.zip",
     "sha256_checksum": "20aaa0142c4d7d5042434e285d3e4888a9231fe946ff13b6883c98bcc4fdb7f4"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_windows_386.zip",
     "sha256_checksum": "cdcc3994072250a137a2c46a94f74bd8038cb09e57be5279904c726f93a70caa"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_windows_amd64.zip",
     "sha256_checksum": "87fd61dcd5c997fb06bb552795f4c0e2a60b5e90db8e4caf0e665f4bc33b8119"
    }
   ]
  },
  {
   "version": "1.1.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_darwin_amd64.zip",
     "sha256_checksum": "85fa7c90359c4e3358f78e58f35897b3e466d00c0d0648820830cac5a07609c3"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_darwin_arm64.zip",
     "sha256_checksum": "9cd8faf29095c57e30f04f9ca5fa9105f6717b277c65061a46f74f22f0f5907e"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_freebsd_386.zip",
     "sha256_checksum": "1b6c7d9cc3e72f761693ebcca96cf42b50abdec15ed825afe88273de39a63c04"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_freebsd_amd64.zip",
     "sha256_checksum": "89f04b966f6b106538d06cee6036999c232c381c30cca0820b348e66cd5aab85"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_freebsd_arm.zip",
     "sha256_checksum": "faf69dfa38a592209fe99b357e46ccb42378d9dafafbcb3980bb3b26ed44d5ea"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_linux_386.zip",
     "sha256_checksum": "22da6e06a13d49c60ca1229ae25d7d7260b7872db258e66a6284690d99b00b97"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_linux_amd64.zip",
     "sha256_checksum": "7b8dc444540918597a60db9351af861335c3941f28ea8774e168db97dd74557"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_linux_arm.zip",
     "sha256_checksum": "debb5332cba6d1eb7f1560feaa0ba1212e55379a3582887bd806b8dd681ff6b5"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_linux_arm64.zip",
     "sha256_checksum": "d6fd14da47af9ec5fa3ad5962eaef8eed6ff2f8a5041671f9c90ec5f4f8bb554"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_openbsd_386.zip",
     "sha256_checksum": "6680b28bb6a9998b411af1645062efb0249c6ed534162e662eb683a2c6965502"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_openbsd_amd64.zip",
     "sha256_checksum": "688ea72768be8836c4a040ae029cd7b83f0a62e22587f469dda26507e1383365"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_solaris_amd64.zip",
     "sha256_checksum": "d1d30e32ad3a50680db8fd7935c5b4d0155c80d8103c6c779406f2e9ba3bfadd"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_windows_386.zip",
     "sha256_checksum": "a1e361373c30a4bc336aa725f638f78d35947c6409270248e7f7ac2cd42420db"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.1/terraform_1.1.1_windows_amd64.zip",
     "sha256_checksum": "a426b6ee78d0c17d26729dfb996d0e296030986449afae1fee09c9b0a9d73bde"
    }
   ]
  },
  {
   "version": "0.12.0-beta2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_darwin_amd64.zip",
     "sha256_checksum": "6be99d150329e55ae636e40500e96a6243a6a00d74126eef9fdb47f17a1070d7"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_freebsd_386.zip",
     "sha256_checksum": "d830952aa6edcc413dfd15c316291605d94fb3325b9df3ffbed0104ac1c65eeb"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_freebsd_amd64.zip",
     "sha256_checksum": "43469064142129df8941598a20043662ab50c9cc1543a78fba59517e08f03757"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_freebsd_arm.zip",
     "sha256_checksum": "53f0d900933541908cc937cb9547c58564e4a8f5a157e1c39320959ee1cc0664"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_linux_386.zip",
     "sha256_checksum": "cac407bcc3e3e268887c5ab931f7ffde8ab868bca5d972c4d9c98031bee4251e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_linux_amd64.zip",
     "sha256_checksum": "ad7515000955fe4a32757dc92e36f7b4d046bc51f3d683cf1e691bb7a6dc09a4"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_linux_arm.zip",
     "sha256_checksum": "5cae86b6ffb46d2806f67a4dc40c9cd9dc08a4092cc1d5b2356e2c03a1269eda"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_openbsd_386.zip",
     "sha256_checksum": "693bf11378aaa0566dbba2c32346f61f44792f1766a6f0bb6377b9dc2658e896"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_openbsd_amd64.zip",
     "sha256_checksum": "e3f45b366b1608b5fcf3905ce9085b97616a43dadb213d934a118bcca44eebf1"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_solaris_amd64.zip",
     "sha256_checksum": "572bcc6e8e4e90ff6b6393cd55d2b9432e27baced1205e664013e83ac7619ffe"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_windows_386.zip",
     "sha256_checksum": "d0210c7eb94df73be086108f0117b0cd96d586d0e0c093656e4e031f67c3fe03"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_windows_amd64.zip",
     "sha256_checksum": "2b7d461d273229389d768152316c883c983488f5ff0d1fa70b9a62ff0062edf4"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210908",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_darwin_amd64.zip",
     "sha256_checksum": "1ab3aa88a42557af41bd9412f14335479785e36af95408a75c0b3bf0af867a27"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_darwin_arm64.zip",
     "sha256_checksum": "cecb24cfd7583a581bb45a46bdfc6617c2afd199ffd78754d903eef883333498"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_freebsd_386.zip",
     "sha256_checksum": "6e253c8a702a9714a00bfb72b2c0bc6639c555b9cbd513d22737d1c88da37609"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_freebsd_amd64.zip",
     "sha256_checksum": "2741feb9a545df0510a7085bca6415ea8e354a9dd644630621864bf13e722f13"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_freebsd_arm.zip",
     "sha256_checksum": "c34ad2ea6e556aa81267a03f2adf56a35d0f8a48da2e4383a626f11958c46d6d"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_linux_386.zip",
     "sha256_checksum": "bf1ab6e52ae0f0dfd1adb765ae01fc444c2441c70768b9a250166e90d82c5c38"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_linux_amd64.zip",
     "sha256_checksum": "92e0db8311f55d6cdfc655260a14de9ead8068a2fad147d96b8ad6855eaf5bb7"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_linux_arm.zip",
     "sha256_checksum": "650bc89c22e32347303279623b2302c92667317a1c0fb2e0637d1fd26a2f1dfa"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_linux_arm64.zip",
     "sha256_checksum": "62b80fedbbb87910bb2649d97b7cabe17025d9597837671f736ad7ee288d673e"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_openbsd_386.zip",
     "sha256_checksum": "677c91543ee7314e1ad2cf63072c39c8bb809ff7e7c900ba8f0ac42b76c693f2"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_openbsd_amd64.zip",
     "sha256_checksum": "eb1cd19fe2b119341c15edf7dc3e958eaa794f9341efe7d0d09242385e8c886f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_solaris_amd64.zip",
     "sha256_checksum": "accdf66092e65852dd5be4e6ebf67dfc78becb9f57d25e1a90188d2a08e6333f"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_windows_386.zip",
     "sha256_checksum": "e71e8ae63783c38fe9a2062d4c2064953ea185baacad39b1c16bd0958e92682a"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_windows_amd64.zip",
     "sha256_checksum": "7ecab59c020ba16c0496f0ff6cdd142d8b600a1d7d2f6f3f98c945f5814152bb"
    }
   ]
  },
  {
   "version": "0.13.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_darwin_amd64.zip",
     "sha256_checksum": "3f7b636638c587c3b03a4b3d1d4cab77c872ad5f7e55405ddaf38dfb94e0e89"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_freebsd_386.zip",
     "sha256_checksum": "66c3e173bf5d94e250c166439d00ce9a83c02af5b6ab00ba48dca01065c3540"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_freebsd_amd64.zip",
     "sha256_checksum": "c8ee2e9913af67c27dad8cab8aa1db200e9baf2314d8b61d5145ae14ecbc2813"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_freebsd_arm.zip",
     "sha256_checksum": "70c354ee536bf43585954b3530ffc5d8b6e76d50ec364cdba6ba9a88cc403bad"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_linux_386.zip",
     "sha256_checksum": "e497be04adfd5f03737ef0da5f705c5bac91a7178ba8786eef7e182c4883908b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_linux_amd64.zip",
     "sha256_checksum": "f7b7a7b1bfbf5d78151cfe3d1d463140b5fd6a354e71a7de2b5644e652ca5147"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_linux_arm.zip",
     "sha256_checksum": "602e055078fa6be51983a58c6685fd48ee4dca9150c657da65604e9b8bfceced"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_linux_arm64.zip",
     "sha256_checksum": "e00d6140d3c92d337835ec968bf47233606943b9ce51467c5119ce8cfd97cc62"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_openbsd_386.zip",
     "sha256_checksum": "396826cf18b04525615453e907c45d2142161a8545363fad80fc7018ce63db1d"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_openbsd_amd64.zip",
     "sha256_checksum": "8f153c999f25724f587ca7fecbd8c306df8dd5fed5e423f3e12de43fdfd84337"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_solaris_amd64.zip",
     "sha256_checksum": "2141d3be34303910bccf07bb0a988da07cebd12b0f2126475f8ce96b39640d3a"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_windows_386.zip",
     "sha256_checksum": "b82d1ffd89f16b75e54c56f1f7bc8c98331ebe802c8fe768926cd47890d70eb2"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_windows_amd64.zip",
     "sha256_checksum": "1e020a208757f23da415b4fee74ae084984c2cfc20210b5c62513c6d9f0a174e"
    }
   ]
  },
  {
   "version": "1.1.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_darwin_amd64.zip",
     "sha256_checksum": "507605a8e3f73f006d5446ba656d69562184bf1631fb96dce3b0abde73dba635"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_darwin_arm64.zip",
     "sha256_checksum": "ccae178553fba390a8a8c73e33285581f9bf510d56010ef7fbd67d067e1f92a8"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_freebsd_386.zip",
     "sha256_checksum": "c7f718db03a37d8574201879e6781bc79808a08a7c45af7ce16ce895306cc69"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_freebsd_amd64.zip",
     "sha256_checksum": "c67f130f85a1e86176a3b139ab32ad6fc37c9616c07e5872ed5f5f3a95a10d11"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_freebsd_arm.zip",
     "sha256_checksum": "732cfd9a5855115e3320765f2083873181cb5c3834716ec396d86c2bdff493c2"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_linux_386.zip",
     "sha256_checksum": "580605a6c35ca785133521a1814317a20f59318e12fe3291a6fafb7314d2f7b4"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_linux_amd64.zip",
     "sha256_checksum": "51113913f04769561fc83d8025d64ed7520338982509d2284033b7763e7c6ffe"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_linux_arm.zip",
     "sha256_checksum": "85fcf0d67257480945be00c556ff51787ea86eab8d531d3b6d7219c491a7508b"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_linux_arm64.zip",
     "sha256_checksum": "be94253ab74d39a5f58659d5538c54c3ea5316af834c46293f96477696436c16"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_openbsd_386.zip",
     "sha256_checksum": "f0ac6d08c2f08fbf74638bf6c47e228a7ba8d9a6683b9d5f3207738b547807ba"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_openbsd_amd64.zip",
     "sha256_checksum": "b7cd2eef59878fe312914af5b02515d3604efe7b52605a3dd06739bcba340ec7"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_solaris_amd64.zip",
     "sha256_checksum": "f27c13ffaa64760f7115f94938eea1f709d588773f7ac778287548fec6310888"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_windows_386.zip",
     "sha256_checksum": "23961df8844093b2973b491f055327a120ca55275ab30f97e79e37998f58ec56"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-rc1/terraform_1.1.0-rc1_windows_amd64.zip",
     "sha256_checksum": "d218cc55382e4db3d9598d7b4a34c1cf8dd31006e90507745d6cca426708a091"
    }
   ]
  },
  {
   "version": "0.12.25",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_darwin_amd64.zip",
     "sha256_checksum": "179fc99ccea5ed3617e9e7026dcfa59a5916ea91162afd7a2acd8350906a0d68"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_freebsd_386.zip",
     "sha256_checksum": "7dbdaa1bcf5355572dd222c34ea11cbe6615d1104cd42fe79faa3a05a5043fc1"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_freebsd_amd64.zip",
     "sha256_checksum": "25108eb73776c0d308742646aa6aa9a44385ef24f41ab726d7ba136502f94c1e"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_freebsd_arm.zip",
     "sha256_checksum": "ee8b2656f15e5afee0842a8c0f2970b2729ee19dac3228be645af19779121291"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_linux_386.zip",
     "sha256_checksum": "51a4653dfdbee000bc47f05606ca8f328a1d31ebd6a3b6569230bcd303883571"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_linux_amd64.zip",
     "sha256_checksum": "e95daabd1985329f87e6d40ffe7b9b973ff0abc07a403f767e8658d64d733fb0"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_linux_arm.zip",
     "sha256_checksum": "20979f0f53a0931cd23332cd2c52ed4ec8257797e8050e5b794ad53c791a7e85"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_openbsd_386.zip",
     "sha256_checksum": "f98f01fbcc1e8e55cc2fdbbb20879b6d030e3ac2f193b92677ae13201a96b7aa"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_openbsd_amd64.zip",
     "sha256_checksum": "b68f8543ef7df431f1318902244daf8baceb41aa0515521d83a83d2c3b45ed2f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_solaris_amd64.zip",
     "sha256_checksum": "6b2a402295801c2a5e9f82a71fea2f7f5dbb0e07d97332a8c828e7004653c6d9"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_windows_386.zip",
     "sha256_checksum": "daa4d3ce5e3e0dbcfe83be4310e7c8fb50098953e2e8fbfabfe2bc526390bcca"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_windows_amd64.zip",
     "sha256_checksum": "81356460648abc8e6b76974c518be7989c6fd6f497bb3c604988fd876b363321"
    }
   ]
  },
  {
   "version": "0.14.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_darwin_amd64.zip",
     "sha256_checksum": "126e1c9e058f12c247a194db5a9567e59ec755cbc0211cd5d58c8b7d37412b2c"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_freebsd_386.zip",
     "sha256_checksum": "f200a6948769ac921b62ace414803be76a2573bc5df423820fa37251e46d601b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_freebsd_amd64.zip",
     "sha256_checksum": "374c0332dcbc5370941e972f5cec39011bc8441996df4a2fc354f1feb7abc4b4"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_freebsd_arm.zip",
     "sha256_checksum": "79538fe678364ae3509bee8b1ba6acaa61ede190a577e788d4b41bbecc6d8c73"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_linux_386.zip",
     "sha256_checksum": "bc22e833a32c8c6424c983e44d1430b7af534a379e75d480647671d7f0351c90"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_linux_amd64.zip",
     "sha256_checksum": "63a5a45edde435fa3f278c86ce96346ee7f6b204ea949734f26f963b7dbc1074"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_linux_arm.zip",
     "sha256_checksum": "adc27133844680a3b9ac2b18d3fb2a61894d58508b93696f57e8c9e6ca3873ee"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_linux_arm64.zip",
     "sha256_checksum": "512791e0a6185b7a895cfe112c289d5d2fb816112406772134343bbe4f1f4a63"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_openbsd_386.zip",
     "sha256_checksum": "bfd4cae6403b7a2523b16d17ff774e3a5eb7ab5da18ffea7234052ce331d700c"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_openbsd_amd64.zip",
     "sha256_checksum": "340d1e42ce25a3c9129f85af0ee72232f6d124908babcfa58c02960181a46e0f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_solaris_amd64.zip",
     "sha256_checksum": "20c0b546265b3f4201dd873395afdb959e6090b89a53f8bced272a52edae9b00"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_windows_386.zip",
     "sha256_checksum": "c6d450ea3ffc53030d325d6c4d1db4b9194c0fb8b3094251b05a5a9ca0948d8c"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_windows_amd64.zip",
     "sha256_checksum": "8f9984006386d32fb5220d012dbeb9c6e654a6fa1476775114b97b343ec4fa4a"
    }
   ]
  },
  {
   "version": "0.6.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_darwin_386.zip",
     "sha256_checksum": "fb7cf062b856e0ded770fd11856113568155954aa37ec46dfa854131b9ff7d0b"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_darwin_amd64.zip",
     "sha256_checksum": "fe54fa09af11a1375a2b85912fe416d494a52137be7c5b0b4aaae35d75b0d588"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_freebsd_386.zip",
     "sha256_checksum": "5862129c4922e34b1672bc0dd7b088ce6091a616171aa7ba51e0714faea8549b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_freebsd_amd64.zip",
     "sha256_checksum": "2268c6dfca0c04618ba70e2db2da989d5adac3b7d1a3873e260ec512b20e2ee7"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_freebsd_arm.zip",
     "sha256_checksum": "712a8d65869c0eaf127d89823d03d3f217c54909f62bfd5fb829e6284170e09e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_linux_386.zip",
     "sha256_checksum": "c12b89c42af88cebdfcfa00d8482e3c9b2ee09d419c09859423503f35def071a"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_linux_amd64.zip",
     "sha256_checksum": "2e5de08f545e117eb9825b697c5ad290ee3fdcaae7d6de4b0e99830e58b38b2e"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_linux_arm.zip",
     "sha256_checksum": "e16349208eb372330aac5045739653fbbf371dc01ec16b85d064caeabbe1bb01"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_openbsd_386.zip",
     "sha256_checksum": "55eca01309e816796e34792c8eea9c0a08f3caa3b1d5034ecf1738d37ee34f2f"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_openbsd_amd64.zip",
     "sha256_checksum": "1472437b677c08d24b1c325f69ef449bfbecd39906082a38039fada0ee4ee00c"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_windows_386.zip",
     "sha256_checksum": "abf7bd57f143b2de902e4230980414c48fe96e0c2fcbad0fd91aeb2c98a0dbb8"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_windows_amd64.zip",
     "sha256_checksum": "1447b66c032f71fc71014342f37e839a991ba796cfde77e74f62903bfef225c1"
    }
   ]
  },
  {
   "version": "0.12.17",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_darwin_amd64.zip",
     "sha256_checksum": "b0ab66e77bac3abcd8b36afa5e567ab4fef103fc21c4a223c954c4ea60f5d244"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_freebsd_386.zip",
     "sha256_checksum": "add4579dbd81bc17eea97457bf910cb2dfae22fec664d25d745464e659251732"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_freebsd_amd64.zip",
     "sha256_checksum": "e357f3eb43dd8ff03d41c57ecb90c31c0dd01c43c51c48f8e118d352d9a52d5"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_freebsd_arm.zip",
     "sha256_checksum": "ccd2353e806238e620e685c2dbe28e2584e13edc40772971b85c165ab4044a90"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_linux_386.zip",
     "sha256_checksum": "64d4a2a7b3c99a150cacf07c4f11682637e9da187c6f0699e8743c4f61e9c1c6"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_linux_amd64.zip",
     "sha256_checksum": "8124c7dfe5036377de0637378ad32cf530477403c29ab1ccefbaa50a05d059c2"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_linux_arm.zip",
     "sha256_checksum": "a49872d6959a7beab61b9772a30ef77330ed0b9414ddde1399ffbb68c08c2e29"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_openbsd_386.zip",
     "sha256_checksum": "76aa1486719eaac5a93e914d09f53578b066c3c767448dd55dfb6e442a787fec"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_openbsd_amd64.zip",
     "sha256_checksum": "321da2dc4339bf03c3a7645ef62034d2dda5cc4a8175573b351d7f3ac02776e9"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_solaris_amd64.zip",
     "sha256_checksum": "74c22205ba1b5b677c748b808783159f8d4d64c99c66a9190a593115de77e648"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_windows_386.zip",
     "sha256_checksum": "4e63d256bdb857deac3974a7aa3057d3cabfaff295423e1bc13b1c2b7ab74bc3"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_windows_amd64.zip",
     "sha256_checksum": "323ffff9c874eff1e44b7bf54178558d85fca7efb1c876852c5fb4677663e490"
    }
   ]
  },
  {
   "version": "1.3.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_darwin_amd64.zip",
     "sha256_checksum": "5f5967e12e75a3ca1720be3eeba8232b4ba8b42d2d9e9f9664eff7a68267e873"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_darwin_arm64.zip",
     "sha256_checksum": "a525488cc3a26d25c5769fb7ffcabbfcd64f79cec5ebbfc94c18b5ec74a03b35"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_freebsd_386.zip",
     "sha256_checksum": "d20ce9b9f7273f858ba0e2202370a0bab0e7f29f6a7fc2480a96ea0e519b1b95"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_freebsd_amd64.zip",
     "sha256_checksum": "92c96c585b3c6d4b701273ab2b9a0b91fc61566d9d166c1c9f4916c33ffd7ebd"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_freebsd_arm.zip",
     "sha256_checksum": "8ffd8b96ee51d746671cb02c806b2ecaebb77bf24e88117d7e0eac1351fc77b6"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_linux_386.zip",
     "sha256_checksum": "e4a14c4729896d8e4d40dbbbd2bc6e46ef8e5f779a4d93e5bd499f043c3a26af"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_linux_amd64.zip",
     "sha256_checksum": "847b14917536600ba743a759401c45196bf89937b51dd863152137f32791899"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_linux_arm.zip",
     "sha256_checksum": "88e9cbd70a456c54db92ca777ad5ed324c002533ca383675cc27c6a9105a7502"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_linux_arm64.zip",
     "sha256_checksum": "7ebb3d1ff94017fbef8acd0193e0bd29dec1a8925e2b573c05a92fdb743d1d5b"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_openbsd_386.zip",
     "sha256_checksum": "51e7b5b3703b34dfe74add09ba2a94feb7b2ec8a5e5a4ad3043bd9ac1b924389"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_openbsd_amd64.zip",
     "sha256_checksum": "d48efd0fc028ec460c11feeabd79ebdacf3969e1a66d3581a6bab736532da26a"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_solaris_amd64.zip",
     "sha256_checksum": "3d3af7fe05dc67fa4db5104ae6d855cd1d03a071910e5487268541915b5d823b"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_windows_386.zip",
     "sha256_checksum": "5b9d78bf3520e4cb0d5c4159295645cdb7ef374f555a59079230b51e6272f857"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.1/terraform_1.3.1_windows_amd64.zip",
     "sha256_checksum": "77e8a03f2866203732efa456cc006e86018a4ff67fa5842cd308acf4c525ccd"
    }
   ]
  },
  {
   "version": "1.0.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_darwin_amd64.zip",
     "sha256_checksum": "ae0b07ba099d3d9241e5e8bcdfc88ada8fcbbe302cb1d8f822da866a25e55330"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_darwin_arm64.zip",
     "sha256_checksum": "3de4b9f167392622ef49d807e438a166e6c86c631afa730ff3189cf72cc950e2"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_freebsd_386.zip",
     "sha256_checksum": "e4ca4e2a702f11f8189cf7c2eee5ef9a951a66cd17cb4f630234eadc284ea1e6"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_freebsd_amd64.zip",
     "sha256_checksum": "f801871361bacf78776feee8ef1f8f2d241b7cb45216900153e766204679cf9"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_freebsd_arm.zip",
     "sha256_checksum": "3a58993c95bbb14b0ca65f18c5708f1868a73e08dca6f25236cbbc1b178c1c81"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_linux_386.zip",
     "sha256_checksum": "f9b72dc8bd095b39eb453523312d7cdd784e6bf1f1b803f23e2c172679be9a26"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_linux_amd64.zip",
     "sha256_checksum": "7ce24478859ab7ca0ba4d8c9c12bb345f52e8efdc42fa3ef9dd30033dbf4b561"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_linux_arm.zip",
     "sha256_checksum": "e34b5274e2fb76d7e6779697304c8f843ee52b523cf212d0bc868c6f4c533ad5"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_linux_arm64.zip",
     "sha256_checksum": "f6c56ebb17d6109d908c2936bbdab74f6f7813c542db85ef6cef3dd020359eb2"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_openbsd_386.zip",
     "sha256_checksum": "a93231cd2f1c15a363c1419d53317f1f6a2ebe366d613c030741c04b6f570ced"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_openbsd_amd64.zip",
     "sha256_checksum": "842b5d91128f48532471c5b5f566d7da9fe1a2f1eb3bc4f404881d4de4504a75"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_solaris_amd64.zip",
     "sha256_checksum": "d4197498b95356d58d390417fed38fddd26ec2542452e87c4d997103586a2373"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_windows_386.zip",
     "sha256_checksum": "a6b5b51ffb8bf5b4f6d6ee894a2c3970797c565d9485ff2e20fda89073906863"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_windows_amd64.zip",
     "sha256_checksum": "37de2cd8153286e41b029a719f03b747058cda09576e3297d3d24e1d30e27a12"
    }
   ]
  },
  {
   "version": "0.14.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_darwin_amd64.zip",
     "sha256_checksum": "31a7691f891f4e4c5d77b5eab9ecc86516df638a0b7cbde120c9e14bef68f7ac"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_freebsd_386.zip",
     "sha256_checksum": "b036d053f4bfed0054469d69e2290260b591d42f479c4d2d3a349bc6e54e0d98"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_freebsd_amd64.zip",
     "sha256_checksum": "61436a868e504e8343c5d9cf1645dd36b077bf8d13fa7840275bb36dab564316"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_freebsd_arm.zip",
     "sha256_checksum": "3ee02c465c8d0b7440e0a7623b029dba1ec3839c2553f44b36ea4be6e7f3a69c"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_linux_386.zip",
     "sha256_checksum": "12c922897a07366686f5b3f3002cdbec9df97008c6ad39fd2f929d752d2b8a9c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_linux_amd64.zip",
     "sha256_checksum": "6f380c0c7a846f9e0aedb98a2073d2cbd7d1e2dc0e070273f9325f1b69e668b2"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_linux_arm.zip",
     "sha256_checksum": "f38fc630f76bd859300b4242172c2f4ed7d0aae6e6ea36dd40bd41627fee2ee"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_linux_arm64.zip",
     "sha256_checksum": "27a3911e33168729fdef5e41b8233f9fedc28ad1d3e06a18c163cdcab74137e7"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_openbsd_386.zip",
     "sha256_checksum": "8dea2d96f13b0da15d2e97af4b22d3fe2d9f1ee6dbd1ed0dbdd0b1ac52bbc00f"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_openbsd_amd64.zip",
     "sha256_checksum": "966e1adab77142379a7f415845e5c131712e413a2d315f69d298990da6effcfc"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_solaris_amd64.zip",
     "sha256_checksum": "1ae028e58fa317eb7fb07c965cc5f67ff784267c166f8c57061ecce74c687327"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_windows_386.zip",
     "sha256_checksum": "979bde8be96f42b7c62d065a12d4ca8d912c75044be250963c0aee061e6b2f47"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_windows_amd64.zip",
     "sha256_checksum": "cccde2f6a3aa8265fde7740561352dbba275f0f5f29323641c3caea6d0e6f286"
    }
   ]
  },
  {
   "version": "0.14.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_darwin_amd64.zip",
     "sha256_checksum": "e728f9c5f64b9a7507f7038ad243743b4bcad0057fe7cc83021eb825cc2b6b9c"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_freebsd_386.zip",
     "sha256_checksum": "6c25e113424f2e773b327581ed61b885c6dd73c99402021174a1f48505e48fec"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_freebsd_amd64.zip",
     "sha256_checksum": "ef5f95e0a83901187a0d7238692c2bcd7ba9aa1cdb7a9bcb6021f601d45911f2"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_freebsd_arm.zip",
     "sha256_checksum": "11fb5b88e636eadab54673f41ab94fcdde8eee40c129923366af5088406edc18"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_linux_386.zip",
     "sha256_checksum": "40c78d378bfabb3f3905440ab3d49d4bd718634fbaea3ab58772af89905569a9"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_linux_amd64.zip",
     "sha256_checksum": "7fd7173f7a360ad5e4d5ea5035670cf426cf7a08d0486bc0fe7c9d76b447722"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_linux_arm.zip",
     "sha256_checksum": "57e5b3e312b7f75c04b0071c1d39ed4de58c71e1e47e7bdf7ef9ad5be7041246"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_linux_arm64.zip",
     "sha256_checksum": "290c437ab37a1f54717daaef359f4935a1fdcdf036bc43fca25e8565c00262d7"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_openbsd_386.zip",
     "sha256_checksum": "68d07a6f6e607b26fbdc6bb93f1680daaaa5a597f32d3d873f935d8512c121"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_openbsd_amd64.zip",
     "sha256_checksum": "680a88209ef1c321298bc1e9cb3df1b05ec0e41b6f4d7d3664036dd16c090990"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_solaris_amd64.zip",
     "sha256_checksum": "a7bcc0967163430b3568544244028d79613c973ec72b0650ca3f96175abc730a"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_windows_386.zip",
     "sha256_checksum": "464240a880b2338b1500b144b6b75eba21b1e63e0a22769e20a42d0752bc7e66"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_windows_amd64.zip",
     "sha256_checksum": "1364434d84981681a1d2f4f7d2d313f2abf23a971b9acd02630dd401000214f2"
    }
   ]
  },
  {
   "version": "0.13.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_darwin_amd64.zip",
     "sha256_checksum": "80af0420732cd08941aa4c0d2b4693056b24366724faa11b107bf052f7de203"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_freebsd_386.zip",
     "sha256_checksum": "d36d8762dfb147598b827235f0713633b1b297e57bbd62807d4a3bec5dd84a84"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_freebsd_amd64.zip",
     "sha256_checksum": "6119c1c59221773bc68c3d501aa2c6183425f8fb1c8614f4e0154ec8556fb450"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_freebsd_arm.zip",
     "sha256_checksum": "57147cb714cee206ffb775fcafea36e5ad92c2a19a88b69aab0fae9ce055dec7"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_linux_386.zip",
     "sha256_checksum": "b856291b48032615595b6dfb32853863a3d989e3dcc5497b9eaa6881f2e59409"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_linux_amd64.zip",
     "sha256_checksum": "9ed437560faf084c18716e289ea712c784a514bdd7f2796549c735d439dbe378"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_linux_arm.zip",
     "sha256_checksum": "92a317950e7a308533db8a4a1b33fe7cdba07f3eb33d81bb7def6a2555bcfe62"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_openbsd_386.zip",
     "sha256_checksum": "d22c64809c68ce55143d88250585f1ea24d2800a60bce6cb8f1803b854151ef3"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_openbsd_amd64.zip",
     "sha256_checksum": "f1b6f3f3f6cc0385e96dcb63efd6457ebe759ff5aecf32d8470c50299eb4a52d"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_solaris_amd64.zip",
     "sha256_checksum": "1483c8ebb99b2dbcccffa04d30b75b1eeeac41b3b129ab320542600ae30e707"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_windows_386.zip",
     "sha256_checksum": "53a8b2162d8a5b77b18472b344af578845983c83eaa910b999ab1cc9186388ac"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_windows_amd64.zip",
     "sha256_checksum": "8af85914d8804c521152167749ca680d7d51447127deb2c7853835b6c62aa9ed"
    }
   ]
  },
  {
   "version": "1.0.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_darwin_amd64.zip",
     "sha256_checksum": "90e58796d84db0a16b5ad40140182061533c38210680980e099812c43b43ff7a"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_darwin_arm64.zip",
     "sha256_checksum": "eace5976af85f9eaf87ac20f6b6899562b8f18500af2fe4eee9e20b61d510b99"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_freebsd_386.zip",
     "sha256_checksum": "a75b3b483176a13b919de3deaa9845dfda6e3dcd6124b3522a1028abcf7de2fa"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_freebsd_amd64.zip",
     "sha256_checksum": "ff287e4253f55896a5fd6ba04ec66c92dc4600e43d83cca49e0f97b3dc5052bf"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_freebsd_arm.zip",
     "sha256_checksum": "4b77e9ec2160600befbed624007189f51184153be5e5a7733fdb82ba2954d119"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_linux_386.zip",
     "sha256_checksum": "82aad0b1aa8f79517b9471ba241c6875647c222c53fa068dfaa4934a7878b557"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_linux_amd64.zip",
     "sha256_checksum": "7329f887cc5a5bda4bedaec59c439a4af7ea0465f83e3c1b0f4d04951e1181f4"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_linux_arm.zip",
     "sha256_checksum": "49977b07651b6662f4d268f49e90b4a057954901347d4c18a54a49cf29b96408"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_linux_arm64.zip",
     "sha256_checksum": "291bcbb4dac589e72aa7c958e9392c77e3afb27694b85a3bf013656f4fd9bf95"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_openbsd_386.zip",
     "sha256_checksum": "2070dc858e5342c75b8f40bb93b6049b57479f86e94a379fd3482dc0814b0040"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_openbsd_amd64.zip",
     "sha256_checksum": "4e65744d44baafb1912beb0b3814a1edb6bf13b5d13cbd35348e8a07f579ef78"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_solaris_amd64.zip",
     "sha256_checksum": "2456bed0cfe43ee686c329c655e7ae99f0b1212c21bc3ba093a41448d1ba2b58"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_windows_386.zip",
     "sha256_checksum": "371d7fb4e20ff34297a38a04f87b2aaba1cace15c348014421d2827543d65ef9"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_windows_amd64.zip",
     "sha256_checksum": "bb515146c49619f78cc141449765fe0290a9d5f955fef8ed59215e163495a3db"
    }
   ]
  },
  {
   "version": "0.3.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_darwin_386.zip",
     "sha256_checksum": "885c566166d7640c730c770471d33e7e2ca8e28d44efffdc02b0932d9b37cb72"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_darwin_amd64.zip",
     "sha256_checksum": "6c8eb551381eb331c0ef3f5615a60529bc45de1c702b02ed4dfa523cffa26084"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_freebsd_386.zip",
     "sha256_checksum": "47080995f7a9aa97087fd7b274ffd53ee7b4f7ab0829d29aac812dfababd7fc3"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_freebsd_amd64.zip",
     "sha256_checksum": "ab9a87c1b2697d75f4c1d406e5a9343c40cb650c7cd044d61668e59ef7a85752"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_freebsd_arm.zip",
     "sha256_checksum": "8fef668badb8257d01496b970a6e48e85e52a46ba5e1cabc176002e52d1c75a7"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_linux_386.zip",
     "sha256_checksum": "8f10ec850287c4a650d73638c79b65d448d59095d68ae24e5e740779a442a581"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_linux_amd64.zip",
     "sha256_checksum": "6e0ee4d7658731f8ce62a8c4326fdbfd39e0f7ce3648f29c96d257bdff88bdc9"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_linux_arm.zip",
     "sha256_checksum": "e5df095b6db2fa49677f9c5ddca2a694e92f45c6ae761194c82a2ed91cf1cff9"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_openbsd_386.zip",
     "sha256_checksum": "d84c62e6ce599fd3924ad547aecfee808fdbd6c7660bde701bcd7a4af28dfca9"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_openbsd_amd64.zip",
     "sha256_checksum": "304b02c1e0684b56b335ef10c7f3d40847a6ace0b968c84b2d7b2e49347eafaa"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_windows_386.zip",
     "sha256_checksum": "f57b62a61fa9e5d53de63c1db9d5b5bccc0db7afcd829929734a0de4c54ab017"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_windows_amd64.zip",
     "sha256_checksum": "fde98446559affba5c763841ce2854f934cee9f7d66b4bdd0d22373982e1459d"
    }
   ]
  },
  {
   "version": "0.11.13",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_darwin_amd64.zip",
     "sha256_checksum": "e9988443da39e5d81a5f7f1b6a5d97b25e2a1151d9be76cdc2e380df97e57856"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_freebsd_386.zip",
     "sha256_checksum": "b07e4c8d991197607af6f80c4bac96c7dbb4658ea61b5bb0fe57353823ea88d2"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_freebsd_amd64.zip",
     "sha256_checksum": "2b17f2aef4635ec61b4796e72b7bb10f761431b1ed8a0083a4cd1714e5cfbf5"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_freebsd_arm.zip",
     "sha256_checksum": "14f97a8a608ac018e7bf290856f34dc59294b9848081be7dd7350f58dc099fe4"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_linux_386.zip",
     "sha256_checksum": "f69a15bfc052c3f4cd6334871e8df381a3c7362d4b6a0399e0ff412fc97cf37e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_linux_amd64.zip",
     "sha256_checksum": "5925cd4d81e7d8f42a0054df2aafd66e2ab7408dbed2bd748f0022cfe592f8d2"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_linux_arm.zip",
     "sha256_checksum": "f77e0bc95e1882b79980e105a5d5731aa9756006f9388dad1ee8459530fc2556"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_openbsd_386.zip",
     "sha256_checksum": "4ff5d427a0eb76d46bbe0bebb55de44551539d3843fc937992f09e788ce3484c"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_openbsd_amd64.zip",
     "sha256_checksum": "30ca2286157bc3eb29e384b0741fe7d0b82c5d0b44f4bc6da684d93048c24b4"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_solaris_amd64.zip",
     "sha256_checksum": "949dae47e2b95669966a6e35b4b13d1fffb197282317d07e09bbb549e790a109"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_windows_386.zip",
     "sha256_checksum": "e4a68ef9728925a0ed8a5485a46af31da5593c884ff3c7d90597c18edeccae0d"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_windows_amd64.zip",
     "sha256_checksum": "b758f90ffe713217eee42c6f00c5fe47958f7286935b67508055af217a33d120"
    }
   ]
  },
  {
   "version": "0.6.13",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_darwin_386.zip",
     "sha256_checksum": "f9659e0cd7302a9351b98a2bf4a925880fc9423e416227f50234eddcd25b17fa"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_darwin_amd64.zip",
     "sha256_checksum": "5f285ea0bf7f6bd704ef262330f88dc195ffa6ed118490d54961958dfe2dab24"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_freebsd_386.zip",
     "sha256_checksum": "fec5734842597629962efbcd44f48ecf613af520e67a70f87258fb7704a1d9f5"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_freebsd_amd64.zip",
     "sha256_checksum": "436b0497e420bd6e031aab571d5f986620683611203845849d21b95e9940b7ed"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_freebsd_arm.zip",
     "sha256_checksum": "e75c152b82a48e7a36eb050dbbcadb4833bc0e97949cc6f89570a1fc23118a91"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_linux_386.zip",
     "sha256_checksum": "cb03081f2347ce164d7745aadbc0fa9d0fd0d26d3381f87065645c9d3e1f4865"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_linux_amd64.zip",
     "sha256_checksum": "8eeddca53958d5c871de93a624eb0f5971f97cbaf107e0b2d1cc289f2ac21b79"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_linux_arm.zip",
     "sha256_checksum": "8e7613fc1fac74321a1f165a63632f0690d1ec9698355e6aea65f31dac82091c"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_openbsd_386.zip",
     "sha256_checksum": "fab966c1a1ed25522da566a3c3c4c80066d34dd92bb89224cfcddfbaee2bcc11"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_openbsd_amd64.zip",
     "sha256_checksum": "75545567a873f8f93b6a47b44c01fe5dbfd307c36cbca098b63f4cdbb27c3612"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_solaris_amd64.zip",
     "sha256_checksum": "e3c9b71c381eb858a06a5c33556aafb70ba4a31352f5d5566263dd04dc358bad"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_windows_386.zip",
     "sha256_checksum": "c86c1732beedf50410d13fb8d78abc9dbef32b55e72a30f098456a8a04d9a234"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_windows_amd64.zip",
     "sha256_checksum": "f63dd182f991ea0222e2769641819400aa99363b1bee05d0eab2e0d44a7d190d"
    }
   ]
  },
  {
   "version": "0.14.0-beta2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_darwin_amd64.zip",
     "sha256_checksum": "fc9c78035efa97c36b2abf590d562fe99ffb9d0fb3224c3b0fb6f80fff4d2754"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_freebsd_386.zip",
     "sha256_checksum": "92f897f142e650e73439bdb1edb5a4156a5fc5e39af71cfdd0734727f7e1bef"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_freebsd_amd64.zip",
     "sha256_checksum": "8e3d294e525d04c729ea929296f68d07cebe22f2b3ae6dec33b5bc4c58f626f"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_freebsd_arm.zip",
     "sha256_checksum": "d7241284ec4ef2544ce02db1e37bd247a8bf3e9a96446a60ca6bdc515647b742"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_linux_386.zip",
     "sha256_checksum": "ee5bdbd95f9759b034792397b36ce86919863594849b2b11706d279148acfcb7"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_linux_amd64.zip",
     "sha256_checksum": "4c2a46e964926624780dd23fa2251964c611c47895ff162001f14dbed8ec4822"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_linux_arm.zip",
     "sha256_checksum": "7e46f9a2f504cbb798c3c6255b845dacf2557383303a3d38a672407bca8113b3"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_linux_arm64.zip",
     "sha256_checksum": "5eda9586ca07312c9c6e7f850f5ef40ed66d84da90f2d3aca6d6f015241782c"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_openbsd_386.zip",
     "sha256_checksum": "14a4d417e9a879b7e095ec54cb33e647e778d8afc7f27b02449a935d4bdaea03"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_openbsd_amd64.zip",
     "sha256_checksum": "65e419c2ec248c01f436e9144d69c25c27be13b7b6ebdfe2a6b1982440447920"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_solaris_amd64.zip",
     "sha256_checksum": "23f0bb9adb55ce5a0adb6d3292e1f9dd122f5f33b7029c2c69dcb1114e829842"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_windows_386.zip",
     "sha256_checksum": "2a4f476ff9dbe7c2b2f55fab534773c5924f66df5590f4419365297f3ded1ae7"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_windows_amd64.zip",
     "sha256_checksum": "844c1b91e2b78327641d3d78638215028a621c016ed0d724809f64276035abdb"
    }
   ]
  },
  {
   "version": "0.8.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_darwin_amd64.zip",
     "sha256_checksum": "10253ac843b7a170844d629cbdbd2287bf687cdd3d2938e4ab9140d10534cf38"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_freebsd_386.zip",
     "sha256_checksum": "a6e9e5d873617778727f17ed49108abe8e3b16fa6a2a1bf42461fa659260d9f1"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_freebsd_amd64.zip",
     "sha256_checksum": "ee6d3bf57892b8e34de011b12271257d3e8cf3b5f4378f55b1f3c52a6c829366"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_freebsd_arm.zip",
     "sha256_checksum": "a60538e287275639fe19eba70f82d08ff599a4f55bb9e01fa67d74b12f3d3fd9"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_linux_386.zip",
     "sha256_checksum": "951ec7f4e5ae441a90d27574abec5e196ea9b3c9fc8c2e2d5c026bcdce428123"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_linux_amd64.zip",
     "sha256_checksum": "4b4324e354c26257f0b830eacb0e7cc7e2ced017d78855f74cb9377f1abf1dd7"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_linux_arm.zip",
     "sha256_checksum": "14ecf7c5196df7e7d3278f1588e0e04e2d27c5ead856e4349ef09e4fe4a90a01"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_openbsd_386.zip",
     "sha256_checksum": "bf5b00825a784e4fc814137b62db563dc83f3b86d202d22f53adb85051e4dc5a"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_openbsd_amd64.zip",
     "sha256_checksum": "a4378c33fe9992e73a589c5a91d480d83b02afae58880f730180237b2c088328"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_solaris_amd64.zip",
     "sha256_checksum": "7091a6595e3a257a9eddca0c9399dad3a233ba7d84aad573d392492013581a2b"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_windows_386.zip",
     "sha256_checksum": "19326143e64b07ebe8241505cd5eb0dd901111bb6590be5855365b4be448821b"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_windows_amd64.zip",
     "sha256_checksum": "5b2d482edcd6478bbb43d25b4f6a142aeffc8cd9be2d1ba28f35a1e569a11c6d"
    }
   ]
  },
  {
   "version": "0.11.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_darwin_amd64.zip",
     "sha256_checksum": "98c168b06e8b4058c66e044e3744d49956ce7bc3664dc1679a33f8fffc84564d"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_freebsd_386.zip",
     "sha256_checksum": "7130cc270e6b710895420d4402466868a7d5636da8ce42a2d96e6338e09baa18"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_freebsd_amd64.zip",
     "sha256_checksum": "29e56bee0dea0ed2755a46535d2dd5d3630ea75df2e089c1bd446098544ab3a2"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_freebsd_arm.zip",
     "sha256_checksum": "9010bc576db9993f2251a68fa6233c6ae448e8716c50977a9174cc5117c64e57"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_linux_386.zip",
     "sha256_checksum": "cc40c4bc677054929777b5f6161702db5a185f1e0bbfa5a8f8c5b08196d0e0df"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_linux_amd64.zip",
     "sha256_checksum": "84ccfb8e13b5fce63051294f787885b76a1fedef6bdbecf51c5e586c9e20c9b7"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_linux_arm.zip",
     "sha256_checksum": "8cecb3854eebd0fd6da3d7a2d45727a6597d6f98ffc916f2dd63b43350d321f9"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_openbsd_386.zip",
     "sha256_checksum": "ee0a27a5857b93fa40ad13c1ab93fce63ea35dd1a9fde1529b1fa4fa46fd3bec"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_openbsd_amd64.zip",
     "sha256_checksum": "a625513c22945849d283c4004b2f35800b6b48d447c27911322cfd18da613b83"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_solaris_amd64.zip",
     "sha256_checksum": "d3baf4758392598c1e59acdcc75b1f7179e218a75d94a8c521e35dede43f4489"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_windows_386.zip",
     "sha256_checksum": "94ed500e27f6db531c696d087be3bd72539a31f1834f0a7d32d0cb9daf8c8e43"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_windows_amd64.zip",
     "sha256_checksum": "b14dbb9b9b100ddcd516ad426f31a23016ce351d7481407bc91bceb342307826"
    }
   ]
  },
  {
   "version": "0.6.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_darwin_386.zip",
     "sha256_checksum": "beb7cf7b50aab11558c378511b7ae3acac15197d44831aa539870fd8d6ac916"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_darwin_amd64.zip",
     "sha256_checksum": "e2eee073432487aabd69003b3a293caa6e087d4b435d29f6406079333e2dca73"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_freebsd_386.zip",
     "sha256_checksum": "805ffc68561c0c6804f9a1d1039ad0a5c274f4239a9cdef5118e38851a7c243b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_freebsd_amd64.zip",
     "sha256_checksum": "92b5440b42e7c4e5bbcb72b4802724ad6c6ac58b9301a31304dc481f6293035d"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_freebsd_arm.zip",
     "sha256_checksum": "c293cd8d06b2ce2b4815d432b178ea7b7e7a070f42cf531ec340dc511e1008e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_linux_386.zip",
     "sha256_checksum": "90bdb676ba475e75a179c5ac27f5ffeb3d8e2cf90d3233d4c55f814f5fa85d46"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_linux_amd64.zip",
     "sha256_checksum": "196ce55e56879d8b696fc5219d30125fc84ba7e0249c92cb8fa93498b0be5c6d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_linux_arm.zip",
     "sha256_checksum": "fffec9f2795d8a9367fc273cc4f062276370e0414323d4aa802f37c821c8c08f"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_openbsd_386.zip",
     "sha256_checksum": "3105081429082bfda821324a18c92ce23cf76603cf71a33f2567dd4848b97b2"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_openbsd_amd64.zip",
     "sha256_checksum": "ff0c6487a25c6bad1522b9b52306d06da68ce948e8030d29b3981b25a5fd35b8"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_windows_386.zip",
     "sha256_checksum": "8ca208bf182a3dfef9f3ff0b947a9d27a13c380c08f428a45d34903e73499c2f"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_windows_amd64.zip",
     "sha256_checksum": "efd8769ba21b34108462022c5cf6440c938b23f58ea74c8e61ce7d8474240de6"
    }
   ]
  },
  {
   "version": "0.6.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_darwin_386.zip",
     "sha256_checksum": "7ab9321ba20ee72387a3212d2de8102135dbc66368bd90126295d1e8487066b3"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_darwin_amd64.zip",
     "sha256_checksum": "a06768862d1c3ee928d26961302c5134c9c8f716e567c4cf52fce85951f61bee"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_freebsd_386.zip",
     "sha256_checksum": "7ecb0a94742d1e96ea0a27cb784c3251358e2e8980a87f6484efd41f696290e8"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_freebsd_amd64.zip",
     "sha256_checksum": "b906ecba74cb21205b29275124221443b55f697319e66cf1f647fa69a1e0f8fa"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_freebsd_arm.zip",
     "sha256_checksum": "653f22bc2aa9ac6e7b1b1e8074a258bd56ee44060374541e4f01374703581c37"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_linux_386.zip",
     "sha256_checksum": "f569250d7e5aca6efe5763b9a1564d6ee816e8148bad99a9c38a89490470a7f4"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_linux_amd64.zip",
     "sha256_checksum": "3f7e135cb106c331f71667e188b602623d98fa37b10eb93c5330c4a63ebee244"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_linux_arm.zip",
     "sha256_checksum": "49f478f4f8e0033ea904cd4d5e445d935c9a9425049ab1ceb7120cb81debdf96"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_openbsd_386.zip",
     "sha256_checksum": "be9226e5bf7baa3b0806edd536d7e72aefd92e11155e98981a90445898022ab7"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_openbsd_amd64.zip",
     "sha256_checksum": "450a01a4b08b0e42b78e2e9f0a93404dab760d408e6603657fa6679a0f7592a2"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_windows_386.zip",
     "sha256_checksum": "13c0e48a0a709adf93c4afda3b80b0754c0038ac8aea41e15643ea573c976182"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_windows_amd64.zip",
     "sha256_checksum": "c2eef7c79cc418c72afabc08c817a8faa799eb8d869c88c0de0219f541fa7dc5"
    }
   ]
  },
  {
   "version": "0.12.14",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_darwin_amd64.zip",
     "sha256_checksum": "2a4538ccf212865cb2c275dc079926f409b3809cb589638f560d5ab389babe00"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_freebsd_386.zip",
     "sha256_checksum": "22587108318e05c210d03cc89fd4ae1da491b36528eed97920bec7ca2de5043b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_freebsd_amd64.zip",
     "sha256_checksum": "b34a35a08e52db147ab7fbeb2242990b7e267ea1de0a285192a71d45a2b2a502"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_freebsd_arm.zip",
     "sha256_checksum": "e70e4fcd2242aac176c9aaf805372b9ac93130356cdae82a848c8419593baa42"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_linux_386.zip",
     "sha256_checksum": "43dfb1163f1a7d5139f61e997f87b46e658c7cd391c937242deae96908c3bd6c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_linux_amd64.zip",
     "sha256_checksum": "8db6b396eb2626f18bf5e98af824645bb96d8290e2ab761b2dace3a2574b8cee"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_linux_arm.zip",
     "sha256_checksum": "baee3536ab5fb2f02a4d426024d46440f911fe2824ee8c5ca80b8cd81ba96c18"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_openbsd_386.zip",
     "sha256_checksum": "60f3a2a2eb797b977f23671a3b874a1c36e48c4dcc9364f64133cb4444b196a9"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_openbsd_amd64.zip",
     "sha256_checksum": "9975c36a33c13291ec100343fb6610f7d5d08133ea41c333a03cf57423cde21a"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_solaris_amd64.zip",
     "sha256_checksum": "6423222ceaca689ba8b45ea6c7b5bb2bef55191691ad65a7713f90b0544fdb12"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_windows_386.zip",
     "sha256_checksum": "44ec4f2110cc7f2158c15d89e3bc161cc422536f16d5c0f42bff73694640b6a6"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_windows_amd64.zip",
     "sha256_checksum": "7409fd42a95dada5ba1ff7ca6ddba531ac04bf85269c725e9120fbe3199d842a"
    }
   ]
  },
  {
   "version": "0.11.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_darwin_amd64.zip",
     "sha256_checksum": "c328b8d60840b96641f519deb85601cb1f2cce458c7bdb7786712471234ac0c5"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_freebsd_386.zip",
     "sha256_checksum": "b41782a4ccb84f1b5a935f53ae59548106cf86f95b27b167a7e7dbad42cf73bb"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_freebsd_amd64.zip",
     "sha256_checksum": "fd7ca41be687f3aacdbb42de7aeb80662c98606a9cce9fab6c43e310726eb905"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_freebsd_arm.zip",
     "sha256_checksum": "63d49e05629563803cf6eb2eb1e2b6a85356a901fadf2d2958678c78a1878f55"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_linux_386.zip",
     "sha256_checksum": "f0af9efe1c3bc87bb6e3e4ac22168ee32cbce8ffd00480cda89f4c5e90fd4bf7"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_linux_amd64.zip",
     "sha256_checksum": "817be651ca41b999c09250a9fcade541a941afab41c0c663bd25529a4d5cfd31"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_linux_arm.zip",
     "sha256_checksum": "5f62de74f1329d91ec7350b31afe9f9c98761c9a46454f015f283c83088ec66e"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_openbsd_386.zip",
     "sha256_checksum": "fe96c763c6bf4c38bbeba413038cf7f1bd1867e6bbf8b28666e4c2ae34a13017"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_openbsd_amd64.zip",
     "sha256_checksum": "7c9cd246a2f88fe56f898d9fffa1fa0f85406afa6cad3d538f3d7ca8124876a4"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_solaris_amd64.zip",
     "sha256_checksum": "56249c1b9ecc4a8a35f4ba0edc5940eac59f891338342b8bba3a23fbff3167b"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_windows_386.zip",
     "sha256_checksum": "8b82283e5cbad072b0f1436041a662ab5a9bca83845cd246a9a9081b542a299f"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_windows_amd64.zip",
     "sha256_checksum": "a762b329798b872f44df3b5db33122469a3cf1ad28c1915fee17605ec8245508"
    }
   ]
  },
  {
   "version": "0.12.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_darwin_amd64.zip",
     "sha256_checksum": "2c2d9d435712f4be989738b7899917ced7c12ab05b8ddc14359ed4ddb1bc9375"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_freebsd_386.zip",
     "sha256_checksum": "51212822e93583cbf8a0e1c9492b0a5add8456d45a99592abe207f5d030c4330"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_freebsd_amd64.zip",
     "sha256_checksum": "a6294e2ba470c354bbf5ccf7af5be8536339d2d8f89482537d311774091c596e"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_freebsd_arm.zip",
     "sha256_checksum": "f544ee6abe76e5187cdd224549dacc606f4fc75a93bbaaea0c99910cba388b14"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_linux_386.zip",
     "sha256_checksum": "93ac24326034eb269d68f2face60bcec5d3af90fc7dfc68b058c66dc5c139e25"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_linux_amd64.zip",
     "sha256_checksum": "43806e68f7af396449dd4577c6e5cb63c6dc4a253ae233e1dddc46cf423d808b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_linux_arm.zip",
     "sha256_checksum": "9f47411805f71de0ad83c03fae1e0b8c36d47e23112c32d1767d99d0d2426a2"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_openbsd_386.zip",
     "sha256_checksum": "52b4861a31002ceaee8638c3847d4158df8a2ee94c9052cb0ad2eb582c98e83b"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_openbsd_amd64.zip",
     "sha256_checksum": "8514a8aca55a713a3986ccc748bc77737da6d30cd136f7cec9d0959cbd465f0f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_solaris_amd64.zip",
     "sha256_checksum": "f8467d89073b18f50c2415abc2efa4a478f729308010c5a48188e2d4cae9384"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_windows_386.zip",
     "sha256_checksum": "c356c4d17e15980c9320cb89c825e3d9867359d4ce91b16504e5ffdb3d0687f4"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_windows_amd64.zip",
     "sha256_checksum": "3d0c41c514841ef38645acdefe3f70f8376e7d8de55828c50087083bd34e9c20"
    }
   ]
  },
  {
   "version": "1.2.0-rc2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_darwin_amd64.zip",
     "sha256_checksum": "7563d6c8cd17ef0e6116c29c56f86e06adcb204e7208b06f1961ea55268a92ff"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_darwin_arm64.zip",
     "sha256_checksum": "eb624f4cd5adcdd2cf75261bb686649b4cca916e1b4546b518448d68f8c36f9a"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_freebsd_386.zip",
     "sha256_checksum": "48978b75ec6dc38aa11ae157c5c1fd042f24f9b0058c6b4cad87679b274678ee"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_freebsd_amd64.zip",
     "sha256_checksum": "abec15165917df9e19fba34ce347f04933ce42768786cfc9a81c530157131eaf"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_freebsd_arm.zip",
     "sha256_checksum": "d8442115aaa49be005daf3908538b265a6e041ab1486d725bd68882453c8192c"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_linux_386.zip",
     "sha256_checksum": "fe7b831b3fa17e8b590339254d5738f3f4a1872da1a4f010630b39cd919a750c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_linux_amd64.zip",
     "sha256_checksum": "77310cb7fd25fe1ca6ccc89fdb319d9c99de083784143cb15c139781a53d2c13"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_linux_arm.zip",
     "sha256_checksum": "8060e0e1e75211f0e4493e43b5deb37277563f9423237d53f7620df58d62ecfd"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_linux_arm64.zip",
     "sha256_checksum": "e383e2a00abc121d2f95909fa890236a1aa89397a45cdcc8c34b0f980e2dc815"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_openbsd_386.zip",
     "sha256_checksum": "f7477f129c7e10452a12655ed7473f36897cfe184fb6dceae141ed1f3da758de"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_openbsd_amd64.zip",
     "sha256_checksum": "85fc89f99a28889a4f32349c0f6a82b47c4b63c5a5113dc189f5fa579b23b6e9"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_solaris_amd64.zip",
     "sha256_checksum": "fc05a2ea69e6aa1e694d6b13fff82ef95eb00198418bf0e85a2476fa59234017"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_windows_386.zip",
     "sha256_checksum": "880c165c07361a921027d0b15fc0ca48d2c1f1e5e4c8a3b80b2d2d2c362401f9"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-rc2/terraform_1.2.0-rc2_windows_amd64.zip",
     "sha256_checksum": "ce5c99f19f790eb197934241188544082c67ea3d1eb32c4780697fc28f40edfd"
    }
   ]
  },
  {
   "version": "0.8.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_darwin_amd64.zip",
     "sha256_checksum": "55ab547539e68c9375c144062460457fcfdb3f5b9f412d3bb162f73298602d78"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_freebsd_386.zip",
     "sha256_checksum": "d72a7a3e5b6ef31e4564009952bc9a5dab273c7dc3f3522f8022540d2eacb3e7"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_freebsd_amd64.zip",
     "sha256_checksum": "b9f17ef543dab4a3a333c1ab228afc3bfaa723fe9d2b574538da8d0b3c079105"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_freebsd_arm.zip",
     "sha256_checksum": "20a4c33679c86efb2b9fa3b140a3397042731419ae3278c7d0e3b16f796d2df"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_linux_386.zip",
     "sha256_checksum": "d483dd90baf752e18af030738a68540e8237ceee5a9ce6f3882cfa7efda7ac1b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_linux_amd64.zip",
     "sha256_checksum": "403d65b8a728b8dffcdd829262b57949bce9748b91f2e82dfd6d61692236b376"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_linux_arm.zip",
     "sha256_checksum": "7591914ba8a64db00b8fc8ac9379643a519a2f796ffe2cc424790a7d1bd99763"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_openbsd_386.zip",
     "sha256_checksum": "933f660009db3ec44d2dc6630526033de5d866a343cbce3eeecbf7c3bae36235"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_openbsd_amd64.zip",
     "sha256_checksum": "1d2ae43502a1550fee2eba358210f8c584f2d9df46b99944e3353ba8fb7f446c"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_solaris_amd64.zip",
     "sha256_checksum": "a694fdc84aa47df5057852138ccfe00d41a5290d31f0ad2b4d9d5057fdfe62e4"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_windows_386.zip",
     "sha256_checksum": "efbf851427d078270d983e12c26dab925555207982a39a6fed81513294b32545"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_windows_amd64.zip",
     "sha256_checksum": "3085671f53e18c3cb5a8c3938b64e5e9f002f8adcfce083233cdb97e814d1daf"
    }
   ]
  },
  {
   "version": "0.7.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_darwin_386.zip",
     "sha256_checksum": "9111c2c83a0f1c8367ed65cd6f22be8f8ad267745b4e17cdb8afa7162ea4b850"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_darwin_amd64.zip",
     "sha256_checksum": "960e0e79c9dcaa51fa349f923e62f46fd4b49a91dcb06677ab096918f6074e2e"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_freebsd_386.zip",
     "sha256_checksum": "32cbeb8482fc65c363f1a3905c1d66dd887d3d05e50c4fa34ed905ddaf60135"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_freebsd_amd64.zip",
     "sha256_checksum": "a15f2dc3c393b2b16a3aeaa76721f3138f1bd91faf81aff4af695dd9a16648db"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_freebsd_arm.zip",
     "sha256_checksum": "b067b9d2a08a2109e62961fe031d6961225e79974e8ba0515315e84157b58f36"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_linux_386.zip",
     "sha256_checksum": "2033813e2afa82dc7eb0062b53536440da33671b02ec0d557aebdc06c574cf67"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_linux_amd64.zip",
     "sha256_checksum": "ac1d0302818ae17f1042dc26407e7ff94cd1e34ed260dae9d72c75a4d0b59cfc"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_linux_arm.zip",
     "sha256_checksum": "2730a808a7d7fd72d7f4e8af280d8f40cb1efd707cbc9c91d43b0429a4d0cf9b"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_openbsd_386.zip",
     "sha256_checksum": "784783eaf58e06a73cc15f51d42d5ff132c635d010ec0419feeb44b4cc4e8268"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_openbsd_amd64.zip",
     "sha256_checksum": "2bd02514a6ea39191d695ff72993c2e6f9657d431c5ee13ebb99474d0081327f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_solaris_amd64.zip",
     "sha256_checksum": "6dc9423391ada1a7382e5174832da26a5b5eba77e128d0c557e63152fd2cff92"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_windows_386.zip",
     "sha256_checksum": "aab289de857d84a2b383f88958c056bdcae92a0b9f50844698cc4d3dcba5d3d0"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_windows_amd64.zip",
     "sha256_checksum": "db6731e48539765f2ea4585ac947b75265c0a8358e9217163c19e1ed9449c8b4"
    }
   ]
  },
  {
   "version": "0.5.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_darwin_386.zip",
     "sha256_checksum": "4645fb5f46337120ec31dc2e8f5a0741bbbbdd4e4c46abf98ca4f71a57254e58"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_darwin_amd64.zip",
     "sha256_checksum": "8033564434ed964fc630fe5ff8b4830945d38a528ad5b14e7a88e23f85591f05"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_freebsd_386.zip",
     "sha256_checksum": "fb01c79aec9908506f76cc722b9114d28d0792db6004466a6f50354a53f59edd"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_freebsd_amd64.zip",
     "sha256_checksum": "ec9a8338df6022201dcb752d86d9a84df8b5f8bd8927989df584fcba655b46f7"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_freebsd_arm.zip",
     "sha256_checksum": "75684eb3bb822146b5d88536a6392e6e40d293518d6239c85daf767f5b4dfcaf"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_linux_386.zip",
     "sha256_checksum": "8a6331314050c29c5919f9fb189ca6d4b2f7c116c8b0830299dd239b4f67559d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_linux_amd64.zip",
     "sha256_checksum": "f643a0113c6c0c3f0825a693fe116369cfbe44a9c6432b2923b8f4b1cbf1cee8"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_linux_arm.zip",
     "sha256_checksum": "51d3fa7b20fb928b7d76dbb4bd8a501d00ff3290af2b60acf106dcfcaf2929d6"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_openbsd_386.zip",
     "sha256_checksum": "57ebb371e8112980986dd9964d0318c9ef477bd9f727894f8ffbd3b73bf3c0b9"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_openbsd_amd64.zip",
     "sha256_checksum": "7cdd48124a290d0ccb7ff8e8a9230e54df623a8a3a2b42cf424505f12dcf7f70"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_windows_386.zip",
     "sha256_checksum": "4faa3577c5978b247bda6347a4080175703ad85d5e79abcea37438d18d466394"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_windows_amd64.zip",
     "sha256_checksum": "dfaba26ee034c0fd8337329df939b68458911104f037c9e748aa334b518ee44c"
    }
   ]
  },
  {
   "version": "0.10.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_darwin_amd64.zip",
     "sha256_checksum": "d39ce30b7aa77834d3000173d95df476c0fcfea8114825d8276c38277d3a7436"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_freebsd_386.zip",
     "sha256_checksum": "1936f07ab731b4d953fe46b4162f2140306fb5cf7efc4d96af37d5af298acfbe"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_freebsd_amd64.zip",
     "sha256_checksum": "9a0b5305c0e325737fe7629bc1c37b5e52967060faed863647241d766a427466"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_freebsd_arm.zip",
     "sha256_checksum": "3100470bc1323db9f66e59c3fe8e0368fd6e8ed6baaa6f3f7c2b9610089d0558"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_linux_386.zip",
     "sha256_checksum": "64332f29b3a53cfd6436e507b5b7dc4aa14b2242fcaf0033065bef3c72c720cf"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_linux_amd64.zip",
     "sha256_checksum": "acec7133ffa00da385ca97ab015b281c6e90e99a41076ede7025a4c78425e09f"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_linux_arm.zip",
     "sha256_checksum": "3f64454aac602887bd760aa1f488286646634739d6b104e62d62303b313bf64f"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_openbsd_386.zip",
     "sha256_checksum": "2b58e6623287013d97fb60bb1eeedcb852da032c81f6346bb4cfaa9b84431b66"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_openbsd_amd64.zip",
     "sha256_checksum": "2579dbc2ec4be501e4dc854f553e87cbb4f3398129703b130e97d9822383566b"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_solaris_amd64.zip",
     "sha256_checksum": "9682c40fd85f4a544d73ff9a21b0b6ea0324b3c3e036f4b480deecad416432a2"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_windows_386.zip",
     "sha256_checksum": "f31912b4419c987853817fbd5905bb41415ea7ddcda6aa8cd0444e32f6bf507e"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_windows_amd64.zip",
     "sha256_checksum": "d84e3f99847c82914055fb4d7619bb9c79b0d788d2ec28f22f606a9aaad71c87"
    }
   ]
  },
  {
   "version": "0.7.12",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_darwin_386.zip",
     "sha256_checksum": "98ee6b9ec94969e347a77c44fd2345a7e48d0c286068958931c8ad01e589ff46"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_darwin_amd64.zip",
     "sha256_checksum": "bfd79badf239509b09c5f036bd5cb1d688297644f26ffaf39d89c1abf9a2936d"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_freebsd_386.zip",
     "sha256_checksum": "fd1a5cd2a9b4dfb14c44b390afcbd2854e9870f933434dee1036a07ae411d79d"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_freebsd_amd64.zip",
     "sha256_checksum": "78441c40f82536fd4686224a89e96c9e8862110e34c2d17c39a1eefa8ea099db"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_freebsd_arm.zip",
     "sha256_checksum": "bf431dc0795c77a389dd0c1daee2a1b0e6e297c33d7a080f2e12f6316844ce4e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_linux_386.zip",
     "sha256_checksum": "d96326ef699d0e4cc6b422b7099efed88c5bab774baf7b2ceffe9c834918d0bd"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_linux_amd64.zip",
     "sha256_checksum": "ba75875a9c61b010e217786f42c7b214b7bc4227922261604c67834638ad65f7"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_linux_arm.zip",
     "sha256_checksum": "49cfaca9ae63777ec87c37d37923abd4d76f9efc2238eb2969d205e9b8441db5"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_openbsd_386.zip",
     "sha256_checksum": "fbe6d043d15d4c569fa80b5de34c28cae84823063a48ea8df8511db8198cbccd"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_openbsd_amd64.zip",
     "sha256_checksum": "52c79636e019961130d6331ff838cd579c89c9824933908ebc36ba0edff3621d"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_solaris_amd64.zip",
     "sha256_checksum": "3b1697d6d9961a89bb37e57943fbfd60ab1b4b3b4cf198d5a6c99f4befdacd91"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_windows_386.zip",
     "sha256_checksum": "b3b9c8632ae0223adee96425f51d71ca4945a63a2fdadc44c4528b47a0b3e5a3"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_windows_amd64.zip",
     "sha256_checksum": "364c92b78cbd26d881ca0df7ef3e29f984d1af48794010c57b84187b3ced012b"
    }
   ]
  },
  {
   "version": "1.0.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_darwin_amd64.zip",
     "sha256_checksum": "92f2e7eebb9699e23800f8accd519775a02bd25fe79e1fe4530eca123f178202"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_darwin_arm64.zip",
     "sha256_checksum": "f38af81641b00a2cbb8d25015d917887a7b62792c74c28d59e40e56ce6f265c"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_freebsd_386.zip",
     "sha256_checksum": "e9b7ba8e291771360082d0a06090d1347ae3b70268733074c0ad17f0d349cc8c"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_freebsd_amd64.zip",
     "sha256_checksum": "9f2d6189992a477b73c7ede602bd83bb444c67b46e40e910469bf7c33baf1d90"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_freebsd_arm.zip",
     "sha256_checksum": "eef96b2f4eda03bbe96f94559d06494048f5fe359bb627b26e5ca57a52c3a662"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_linux_386.zip",
     "sha256_checksum": "143d01fb4bacf8dd97918fb2e1ad38d4489fbe9f5669a32d0adb69ce3e9aecea"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_linux_amd64.zip",
     "sha256_checksum": "eeb46091a42dc303c3a3c300640c7774ab25cbee5083dafa5fd83b54c8aca664"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_linux_arm.zip",
     "sha256_checksum": "cce11dd382af930ef20234d84695d187bf869e366e7d91337068719ff1a7c843"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_linux_arm64.zip",
     "sha256_checksum": "30c650f4bc218659d43e07d911c00f08e420664a3d12c812228e66f666758645"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_openbsd_386.zip",
     "sha256_checksum": "fb92819ea05d0464f2beff937c662dec3521ab8921cfefe6fe182cb639aaebf4"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_openbsd_amd64.zip",
     "sha256_checksum": "45d7a676c9f34e284aaea079aaff2451ef9966827ecaf3e257d13a39f97e8fa5"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_solaris_amd64.zip",
     "sha256_checksum": "8050b841e22205eab23b7d26102c50711a38917d74d5aa5ef475af58768d98fb"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_windows_386.zip",
     "sha256_checksum": "5c36b396efb6f3024d5d371a53a964c45b213aadb68131500f2f465b00e8066"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.11/terraform_1.0.11_windows_amd64.zip",
     "sha256_checksum": "b2903d4866bda5579e46eeadff569d1eee4468ad9c03eb688769978f6cff8ae3"
    }
   ]
  },
  {
   "version": "0.14.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_darwin_amd64.zip",
     "sha256_checksum": "3077741547eaa8885aa0f8fb9ed160b6f069a55c8e8f908a316416a13c4407ca"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_freebsd_386.zip",
     "sha256_checksum": "59d9d4feb2ffb7f98813cc3037c84d38c05049e05cf255cc2cf46a8fab47d8c3"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_freebsd_amd64.zip",
     "sha256_checksum": "df5d09134cf45e3ba2f4468e2c447060b935bd5395f615b16df337622d7b09ca"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_freebsd_arm.zip",
     "sha256_checksum": "138e7812862065d591ffe843da7a2261e1150e6dcfeefc28dd924b26a03a7764"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_linux_386.zip",
     "sha256_checksum": "b1b07704529eece3b0f8d54f5d8d7f5cf80707d0cf9900c86a5dc9fefa6d7090"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_linux_amd64.zip",
     "sha256_checksum": "497dde94e0d8f5a9799fbb0d026430539f6c0cee536d15207800ae06edcbf5bb"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_linux_arm.zip",
     "sha256_checksum": "2f0791cc307417cf6dfb875fd3f64e793fcc9bc913a28994a6136f0ee3aa2689"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_linux_arm64.zip",
     "sha256_checksum": "a6dde2a1720eaed74220542777734bd950eba0ceae33c862bb395db238e9b1e9"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_openbsd_386.zip",
     "sha256_checksum": "b904e9e4c23b97de8c09475aab3ed83aeb5d0a929830494777085bf26988624d"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_openbsd_amd64.zip",
     "sha256_checksum": "557f8f7c186200e234a6c567ae586bd475068e6bcda7a9ae0c4ab1315850799e"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_solaris_amd64.zip",
     "sha256_checksum": "a5fa83cf8ea11d39d6cd17d3f176ae4e29f3ece1a060b7ac0a13e99c14ca3a8d"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_windows_386.zip",
     "sha256_checksum": "a17a1c140f09513a8bfff8a5f2aeb4c5c20d43e6a0065a20e49878e41cf495a2"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_windows_amd64.zip",
     "sha256_checksum": "939701cde46a65e2247114a87db8d64c4a23466c4be854ced4b7b08e443d7939"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20211006",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_darwin_amd64.zip",
     "sha256_checksum": "4f5ca3423ca5d3bd27f7006b07a6259132c2a9fc0d0809736b8509847b987a07"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_darwin_arm64.zip",
     "sha256_checksum": "b4225ec2ca33ca4086be9b90066f4b93df987dff924234ea20f612a756579099"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_freebsd_386.zip",
     "sha256_checksum": "80354f91559510bbe378482c6ece88bb39a785bae3fdf428f615e4d847d93107"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_freebsd_amd64.zip",
     "sha256_checksum": "a662eff709d162c26af3d0e21df60fc96393baacda33913c7cfc0f0ee77c3aa0"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_freebsd_arm.zip",
     "sha256_checksum": "3eb6ad22bc71ad771e0d44de1439b626166ca64c800c0ab8b8760ed7b8786c75"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_linux_386.zip",
     "sha256_checksum": "d1b79676514a9f7146d3680fd66fb3146b0546a17245250f004966065b3a61ec"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_linux_amd64.zip",
     "sha256_checksum": "2787d3b6d454e76c45520ff998437ba3e125686c0c577f1c6db50f38c5d8d75b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_linux_arm.zip",
     "sha256_checksum": "6aa6f47ee7129854a2f96f6e13852646e5697bc2af3b763f9ceb45a66ccf943b"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_linux_arm64.zip",
     "sha256_checksum": "b988ebacc10bc154537d2a1979aef50385f6c80a766589ebab3a28dc1f6a268d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_openbsd_386.zip",
     "sha256_checksum": "3201f4faab4a0a04a3e3b731b669b7f912c69ea5034e2a6de0f1a375726799b4"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_openbsd_amd64.zip",
     "sha256_checksum": "4052b9e2544bab7f0103160bbe2624d7e66d7d5a7b16b460a9210752dea4a560"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_solaris_amd64.zip",
     "sha256_checksum": "803662aba2481ad2b19dfe85ce5b59a76770f782d117d4cb8fcb8b084dc3b7f7"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_windows_386.zip",
     "sha256_checksum": "902eac37afe2a62d3c91c222003472121e0e648baac414e04ca95674d8cf10af"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211006/terraform_1.1.0-alpha20211006_windows_amd64.zip",
     "sha256_checksum": "c4f3da90998a00072f2968ae0923446ad9fc2d0472fdbc0cf3b5b0d08bddf39b"
    }
   ]
  },
  {
   "version": "1.2.0-alpha20220413",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_darwin_amd64.zip",
     "sha256_checksum": "65220850ab69b2b953526a15e7a9865ca36afc1c4161fd26694327c797697a6d"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_darwin_arm64.zip",
     "sha256_checksum": "479b394738050446235784288879190fadad9e8f463b512d351c22830d9399d9"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_freebsd_386.zip",
     "sha256_checksum": "fa6e0f06b35ab9bc3a1a2c7e16b348e010dda1149babc16d7929febea95216a2"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_freebsd_amd64.zip",
     "sha256_checksum": "65068bf9fba0142ec3133a986aad68eb2b1ab8ccf9613483da6fc16a7e4d803a"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_freebsd_arm.zip",
     "sha256_checksum": "8d964767cd81a30d70572df95a170e2a44b0bac5ceedaa940ac5bfec76c82577"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_linux_386.zip",
     "sha256_checksum": "5d42e55d8229937108e5f6e66d14110660da7101599bcab31ec26c770624d35d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_linux_amd64.zip",
     "sha256_checksum": "a00b063118db28b2046aa08795aeac03817bdf7c1d2326d6166bca8ef10527ce"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_linux_arm.zip",
     "sha256_checksum": "f94a2cfc83c4bd1a1e878a683cd1a5447b19fafc545a9efcf099e276467da691"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_linux_arm64.zip",
     "sha256_checksum": "24ff333ff88817cecb662aa15aec520c700fa934837c29e7680fbe87f2773f"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_openbsd_386.zip",
     "sha256_checksum": "c060b2acf13e6094ef1d9fca9643a67fdd42f5e0eb5caedba626368dbc2dbe2"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_openbsd_amd64.zip",
     "sha256_checksum": "6d8da47825ce28ec7db88abbdb2bc2bad8d82105c28555c7f87483b8a28810ec"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_solaris_amd64.zip",
     "sha256_checksum": "b68d59bf6023a0f6905187098f14c13da94725ba359e19b6c2480071d7f3036f"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_windows_386.zip",
     "sha256_checksum": "ed3790c92c3babdcb89db6fde22aae414e0a9f66f13a81d5ff2a8bfda31c4fc4"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0-alpha20220413/terraform_1.2.0-alpha20220413_windows_amd64.zip",
     "sha256_checksum": "2ae82a46653a41d21ff48d21d5ef859230aaae65df76c04003ac0d5ce39df04f"
    }
   ]
  },
  {
   "version": "0.10.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_darwin_amd64.zip",
     "sha256_checksum": "1584dc21ad5ac1dc0d9a2876542a85d092778d00a0622622c28f8740abadddb9"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_freebsd_386.zip",
     "sha256_checksum": "6903daf9cd4755537508f7fa2ba1e81574f6f45fb15e672fe962927883e4cd7f"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_freebsd_amd64.zip",
     "sha256_checksum": "903b2524b915a46d5bbb8ddba3b21f0460ed16073d79ed1c6f002ed76d4f0f36"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_freebsd_arm.zip",
     "sha256_checksum": "651bc639768fab1cd0a7a4806d2123839b5f7b5537b7f28c1fe4da65323b23ce"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_linux_386.zip",
     "sha256_checksum": "97cca04d14136480604e73f7cb2efa409d27ca6651291c67bc3e8c1a2ae499d0"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_linux_amd64.zip",
     "sha256_checksum": "f991039e3822f10d6e05eabf77c9f31f3831149b52ed030775b6ec5195380999"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_linux_arm.zip",
     "sha256_checksum": "9167ded74935738f6f036670ba01f82d8e7d3378ed6145ab56db00c0006305c0"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_openbsd_386.zip",
     "sha256_checksum": "f4538f9156df1eb8279fa6e49b969a81cec8faad6178358c8c8f888160a9ac6d"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_openbsd_amd64.zip",
     "sha256_checksum": "e90520f1298c45edcd2e70c8cfad9f70cff802c46079c1438684a44cab16a50f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_solaris_amd64.zip",
     "sha256_checksum": "607c162164e0173beb13a4396701a0bb739bab9c77db08f69118261c908d715f"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_windows_386.zip",
     "sha256_checksum": "a118202c72a8404d406e8250e2163dcf218dd31fe16f846ac3996854fa3c7613"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_windows_amd64.zip",
     "sha256_checksum": "d9697b0153bed48af4b077c711197f8fca1163e94d5a1c067b077a02201be25f"
    }
   ]
  },
  {
   "version": "0.3.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_darwin_386.zip",
     "sha256_checksum": "702ec65ddbcf9b25c9957335884f7118d27bc209c6be426f9fc48882a0310f90"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_darwin_amd64.zip",
     "sha256_checksum": "d583d58719951a5c3a06eec38390fe31bef7645af7fee3e915293aab7a910885"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_freebsd_386.zip",
     "sha256_checksum": "ba7f5fbf2dee9430becfd668b453934e3ba7f7f14e2b29f65aceac4e11838499"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_freebsd_amd64.zip",
     "sha256_checksum": "3969ec88222b0f7970d4b965d5fd86500f114de281706fcfbc083519061a7d83"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_freebsd_arm.zip",
     "sha256_checksum": "e330a6fc6983b5a65bb6548b18052d88216e4e757cc1e3743a22cbdba31bbd9c"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_linux_386.zip",
     "sha256_checksum": "530ebdb835701d1808f1c0ea041e6c1819c68732e78a8f65e5eaf530825000cc"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_linux_amd64.zip",
     "sha256_checksum": "fab2246ede46c35a45dd0c2a15d089dc97b6541327c3d33c4849658f3068ab18"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_linux_arm.zip",
     "sha256_checksum": "8725eac3cf6d8bf5245d6322f07b6b4ec9597f5102784e14d1ec3fc8bf378508"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_openbsd_386.zip",
     "sha256_checksum": "b511951e9ed479a2d4c49e1f00a58b2d78c2ee33ae486c198ae00da25c6aa9e"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_openbsd_amd64.zip",
     "sha256_checksum": "6fad14ae3a8020501ab04218e2346234128f1a7c90c66891daa6eb0e36bab3b3"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_windows_386.zip",
     "sha256_checksum": "36814330ac097c1263d2478d2d47b00b5dd2eb1c5cc3d30b4875d78395273403"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_windows_amd64.zip",
     "sha256_checksum": "79f3f73e3ef954bc1d302ec1149aabc4d22e180ea760977fff5ab5465ec4f315"
    }
   ]
  },
  {
   "version": "0.13.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_darwin_amd64.zip",
     "sha256_checksum": "cf24555d0089947d690dbb4860bc7f4206da5b71092f150c4785185b2ed837cd"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_freebsd_386.zip",
     "sha256_checksum": "f43fcbb696c2be2a9b8001a6442c51474be40fa4dba6fe99aef2330f41d539d1"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_freebsd_amd64.zip",
     "sha256_checksum": "37d79093a0ef14bde0b065639a651109f39a7bba2f77b7b7829d502e13df9724"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_freebsd_arm.zip",
     "sha256_checksum": "3bd3d934019bb1106f32ee072e509e8c836c2deff4035f1a342f3058d7452ff5"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_linux_386.zip",
     "sha256_checksum": "cc7ec15fcf4ed6885449040a5227a113d67382c5a0545df65572dc9f0cee95c6"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_linux_amd64.zip",
     "sha256_checksum": "4712b8c73e8aedb2fad973df483671f30c6796c5723c867baaa22253672b4371"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_linux_arm.zip",
     "sha256_checksum": "cf35c0fb830480ec031f62f8a05e060a835e9068a8448189fde64a3e913fcdbb"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_openbsd_386.zip",
     "sha256_checksum": "768f1058b4370cf8b03fdb7bed2fb729a8d7f5f80ed9ccf4768d671d23db2f33"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_openbsd_amd64.zip",
     "sha256_checksum": "31e4d08788ef2c30ed164914ff849756fb26efc092980f62204970bd3d4655df"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_solaris_amd64.zip",
     "sha256_checksum": "563059efff168ab75e5ddc922aee1233d7e164dc6b9c1c64cbc8326ff4ea8557"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_windows_386.zip",
     "sha256_checksum": "50f77ba9e0691bf2fc68182b5b9ae250e4748f173bcb9e6352ef5003c54349e4"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_windows_amd64.zip",
     "sha256_checksum": "8dc2ec27b462eca3ee6fbc4cb7529c951375354836f3d4e5f878cd447d24bce5"
    }
   ]
  },
  {
   "version": "1.0.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_darwin_amd64.zip",
     "sha256_checksum": "cf7c7750c6380a12d6a4534c63844c803f14f5c5a8f71e32f7ad237ae81ae7a9"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_darwin_arm64.zip",
     "sha256_checksum": "8b0a06a75fc6ae50cb75d2a3fb64d66fbe443bcda6d12d9e637cd3a701a29567"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_freebsd_386.zip",
     "sha256_checksum": "355908794fe4014c8c27e6796d1fd294f610e7e8f614377a14d761f993c8dc74"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_freebsd_amd64.zip",
     "sha256_checksum": "d26dacd69ea3d3792f47d2b3dd1094df1fb54bcbdefc95d4aaf2b5f225856789"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_freebsd_arm.zip",
     "sha256_checksum": "98068675cfdb2b15e1191732fad01a398e80c0b9bd71bc56d9c067331f5baf25"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_linux_386.zip",
     "sha256_checksum": "7dec4881e811cd9789e92db09db2865d4ad29922e0a3a925f273797f3add08a5"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_linux_amd64.zip",
     "sha256_checksum": "5c0be4d52de72143e2cd78e417ee2dd582ce229d73784fd19444445fa6e1335e"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_linux_arm.zip",
     "sha256_checksum": "4829e9ab71b1c13ea43e5c1085bc78fa81c1d001686d96276010ebe76b46c66d"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_linux_arm64.zip",
     "sha256_checksum": "4e0132f0023077030df92a5b2f2b34431310da198ecf22aaf363b6adc374430f"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_openbsd_386.zip",
     "sha256_checksum": "8436ad26b562a5f0210a4e66b7eded32d33f864768ca604b0a9465bceefb8ee7"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_openbsd_amd64.zip",
     "sha256_checksum": "97b50413eeb166f26390cb6dcde42939a83d5480ea992dc2df4735dd3ad7b8a"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_solaris_amd64.zip",
     "sha256_checksum": "1c63aab88efaeba82846c996ecf15830a4a85432672c9113fd798f2982c006a8"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_windows_386.zip",
     "sha256_checksum": "c3d82422eb237c3175db68cd12e7bf2c17ae1b320a49cb06dd451db004f64412"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_windows_amd64.zip",
     "sha256_checksum": "643ce086596cc3ffbf4389850588f72d775f25bc732f4b388fd8219f7be4cc1b"
    }
   ]
  },
  {
   "version": "0.6.15",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_darwin_386.zip",
     "sha256_checksum": "d7d97ffdaf1d2853f23425568734126e50486defe81178cee74f3e9f4f717cbb"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_darwin_amd64.zip",
     "sha256_checksum": "9cb305ac00b85e2575da3c71504f3fdd3f7ef61f35457af999c7b88802143311"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_freebsd_386.zip",
     "sha256_checksum": "7716620b30511ec3d58789df7be10c6a935d7587b9ce2fcde16f8fd593b61138"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_freebsd_amd64.zip",
     "sha256_checksum": "a00ff5f5246326f80603348df578abe48327b05bf2477b08d49967e011a2ddd8"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_freebsd_arm.zip",
     "sha256_checksum": "5db358ee040fe43d2902bd5169c5afd6445ba1fb2ba5c0bdd2e68e3e44cc83ef"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_linux_386.zip",
     "sha256_checksum": "dd2c26d3fafe9fc83f429425cf00bceae857c6e5a96d79d8cd49a396cbd17fcd"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_linux_amd64.zip",
     "sha256_checksum": "2a81faa54ed6c5e7c065444617fc999f0ab6d433e4e03a0ad599892e74ffff6b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_linux_arm.zip",
     "sha256_checksum": "1a86b16b8ac1a20672b4cdb8eb6afff0241fa1be3a1a18c38f1d71d57a610d2f"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_openbsd_386.zip",
     "sha256_checksum": "fceba6d9d101f329c2bd44db6ef0359b18bac3291b0e3ab0e967256f5f6a23a7"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_openbsd_amd64.zip",
     "sha256_checksum": "22db1c3f957ae0035efd9910c73a76b98d4b3aae780f9c8b731db16521f5e16"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_solaris_amd64.zip",
     "sha256_checksum": "f09a49a7c513e844000aaedc7cad6f2a4f33ee02376276024cd3df50ec73fce5"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_windows_386.zip",
     "sha256_checksum": "e8d0dd1ce202ac80467333bc45f0bf2d7168d138871737db39868d6dc834305f"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_windows_amd64.zip",
     "sha256_checksum": "3f3c74a3c4ca3dc2e77f333217bf6316e79581b9661d8567b5822cc715b22ba"
    }
   ]
  },
  {
   "version": "0.5.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_darwin_386.zip",
     "sha256_checksum": "18268585adce424895f03440d6a5ccc1342f49f345d543723c1be76de173e651"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_darwin_amd64.zip",
     "sha256_checksum": "9d3388266510a03ea5f5ba2a721ab2affc854777c973d821f16e7dcd514adb7b"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_freebsd_386.zip",
     "sha256_checksum": "d33135f9a2fb37e3096584d31aab4f34bee8f8c6164a1ddb3403c849ed60c1f6"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_freebsd_amd64.zip",
     "sha256_checksum": "c6d064757e34189f5d9acef8f6c09100d57be3c3ce32ac0ceda0724e34b010fb"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_freebsd_arm.zip",
     "sha256_checksum": "f291d843236c8927e165d1de7081a906b91fcb21999ec2552dfc8f9985a202de"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_linux_386.zip",
     "sha256_checksum": "20287208868a215b4de3d64ca0a0b2bf991fd96bd63db25f4de388577634a5d2"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_linux_amd64.zip",
     "sha256_checksum": "c5c36be4e2cd2168142dd2b2950231d4b68ab0c14880aec1bbae1ed04cf2a16b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_linux_arm.zip",
     "sha256_checksum": "692d4678959ce35fdfc4ff50736329148dc17a765cdb7cf49888eed663baab65"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_openbsd_386.zip",
     "sha256_checksum": "e4cd0219f2297746997bb4a42e7e24761da0065cf3f37a44e49bdddd7fe98ac3"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_openbsd_amd64.zip",
     "sha256_checksum": "55e50c7210cf36cf3e52e6a8a7484014f38ca7f792150879a6bdfe1f4b65c0f8"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_windows_386.zip",
     "sha256_checksum": "521e97a362e5ec6f3f63dbe7ac22f929a1f78477645664961022a5f30e884e2b"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_windows_amd64.zip",
     "sha256_checksum": "3724c68ee1d39ac8842702a14d07a8d3786d5994c7ba21487f6f7e8ce64bb43e"
    }
   ]
  },
  {
   "version": "0.13.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_darwin_amd64.zip",
     "sha256_checksum": "ccbfd3af8732a47b6bd32c419e1a52e41eb8a39ff7437afffbef438b5c0f92c3"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_freebsd_386.zip",
     "sha256_checksum": "55d1fd2a709946c1fcfe1152f437843d5e155e9f1fadd85c875fcbe69cb44fd8"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_freebsd_amd64.zip",
     "sha256_checksum": "f4b400a61de88af157ee9955fe226dd4b35474c57bb0e6b8407b395771b86648"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_freebsd_arm.zip",
     "sha256_checksum": "86e49667a5d9c476e964d43d7335e50643032a7cbd3dfdc0e262286152221403"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_linux_386.zip",
     "sha256_checksum": "5b52a34da61a9b411b519c622078e8e6cc07aa8685e428fe86c4cec51d6dca6f"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_linux_amd64.zip",
     "sha256_checksum": "35c662be9d32d38815cde5fa4c9fa61a3b7f39952ecd50ebf92fd1b2ddd6109b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_linux_arm.zip",
     "sha256_checksum": "c5274bcc6cc597467d9c90ee6f22281496d91da0bc1f899dd185f9c7bedaf207"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_openbsd_386.zip",
     "sha256_checksum": "9a16d371c9025bef9b6d8ca92260a5623f3ae71c51b967791ea58482686f57a5"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_openbsd_amd64.zip",
     "sha256_checksum": "76e7d480c6979459b4a111d39012def1929ac310b1f84e5756a6209bc8064c2c"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_solaris_amd64.zip",
     "sha256_checksum": "4c293e365f34b226d741ac6b7cbed59645b4311618d6e398baef7ef6b896ebd7"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_windows_386.zip",
     "sha256_checksum": "f039b0180abff3486022a3e35986aa959b7650a438277b99bfa13a16a11538d0"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_windows_amd64.zip",
     "sha256_checksum": "e4aba639b2fb946c5c17b982c22c8ff3a7a3c07725978284ffc1cc651961da2c"
    }
   ]
  },
  {
   "version": "0.10.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_darwin_amd64.zip",
     "sha256_checksum": "60924d17e40be4b055629719a1f633736cca70c4506b8f7e32fa17e0d6e57477"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_freebsd_386.zip",
     "sha256_checksum": "5462fae3b795ed4a6420705daa56cdf4c437aeb4a5e61ecc94285ab0c6d66ca5"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_freebsd_amd64.zip",
     "sha256_checksum": "4e3e1bcfc271180c79232cb9a06fc958d6ae631041b5afebdb9017e18c79e307"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_freebsd_arm.zip",
     "sha256_checksum": "343dc6e85e80c0454076ad20ce3846862780b66542c4e76a767e02d4e9cfc883"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_linux_386.zip",
     "sha256_checksum": "423d905a67f1d867a0626e720d33a75d17aeb8e6eac20e4d7e2ca5ccec8e48ce"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_linux_amd64.zip",
     "sha256_checksum": "8fb5f587fcf67fd31d547ec53c31180e6ab9972e195905881d3dddb8038c5a37"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_linux_arm.zip",
     "sha256_checksum": "68c0a94e23150fa0dffc4826983f06c4dfde86c2142393e8f6b2cee1bcac34a0"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_openbsd_386.zip",
     "sha256_checksum": "de2dd10e2adacf3a2d5010487e773fceb7df025551393876099d114adfeba516"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_openbsd_amd64.zip",
     "sha256_checksum": "2e100a3da07f658e6d22fcb2ae478bdb336388b2b4e302a083056284ca3364c9"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_solaris_amd64.zip",
     "sha256_checksum": "bf10b3146e376ad112b5914028a8364374b301dc04d1e4320498925065a09b22"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_windows_386.zip",
     "sha256_checksum": "9c7d40a3a30c3bd611fd8dd72b2aac4c0e83be898dc097cb2e7ed7a5721736c3"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_windows_amd64.zip",
     "sha256_checksum": "872fa58641020c242a579587838bb0842361e1daf0ae1f97fc3e9bab9a4d63af"
    }
   ]
  },
  {
   "version": "0.13.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_darwin_amd64.zip",
     "sha256_checksum": "d5fbb589bc35c2655d0705c26117135cbb25e4259f120415009e0e6427ea97c8"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_freebsd_386.zip",
     "sha256_checksum": "f85d2440ce66422982228e7868cffcc3ace6685548bf3d17bff0629b91f1a32e"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_freebsd_amd64.zip",
     "sha256_checksum": "69583586db251e969c8d8f962b6bf5759cb29a9ac6e817da645217a2f89e60cb"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_freebsd_arm.zip",
     "sha256_checksum": "fa064b9609639f2b2c1f577472fcd8b71978ee5248bd35a7d3ab01939ccf668a"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_linux_386.zip",
     "sha256_checksum": "90df3c9e8a9facf93631c6d338b867db391b435a39a1ca0a02ae4908f01dc2ea"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_linux_amd64.zip",
     "sha256_checksum": "4a52886e019b4fdad2439da5ff43388bbcc6cce9784fde32c53dcd0e28ca9957"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_linux_arm.zip",
     "sha256_checksum": "9af803e215be58ed6cf89e9d22b826d7b3634bef70dc94abc13fdd2bd85b1f80"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_linux_arm64.zip",
     "sha256_checksum": "59b85b7b4b4b2939a4c05db905100dd8808f8df8b5a0e8a00a26e7dcf0c4a2be"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_openbsd_386.zip",
     "sha256_checksum": "f440fc7ee83c559ec9708a819f6b9029cba049b226e4aaa4a88ed5e94635acff"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_openbsd_amd64.zip",
     "sha256_checksum": "21368d031177036be0cf5c3b6481a248cae214a551b565e81413ca4ae47c39e9"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_solaris_amd64.zip",
     "sha256_checksum": "a80c162f5feacfbde6e5c77ce5d99d5832ef03c6677415f07506dd12341985af"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_windows_386.zip",
     "sha256_checksum": "6c0223a0e5f510596e0cbca2d0deb1ff055be206d71f0be485e5eef2e7a97a34"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_windows_amd64.zip",
     "sha256_checksum": "62e6a61afbd043faa4ad011a37f4073a181fe4150036eb209162c548d328afa6"
    }
   ]
  },
  {
   "version": "0.10.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_darwin_amd64.zip",
     "sha256_checksum": "cb8b8c7abc291467bd432cbadb993b6972538c0d438cd6933d29c5c0702574d2"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_freebsd_386.zip",
     "sha256_checksum": "c18807de9da06b8e736f1ce9d8766649d1d1e1675ae5fc76f51e13c2781e170d"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_freebsd_amd64.zip",
     "sha256_checksum": "e9d8c248590d07962a88520435b92f0d8feadbd9ca8f17934b4dde2372bc29"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_freebsd_arm.zip",
     "sha256_checksum": "a3ae2349b4003ac58fd3ff40f9b6606c18432e75c6565a4b2c097482df10b4f0"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_linux_386.zip",
     "sha256_checksum": "85514545520d8f7eb47fea492215fdff23de578029595b44e72a36adc6964753"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_linux_amd64.zip",
     "sha256_checksum": "66820a4d12c25a6d0c83dd877d308aa0c8fce11ce63072f5753167e917464770"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_linux_arm.zip",
     "sha256_checksum": "15cb72abf497ea0d2e4bb4d06387b93ac8394c73ae4bae5da53d240fa2fac556"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_openbsd_386.zip",
     "sha256_checksum": "9afded76df3f2424fdd20b3596eddfe4f14ae60c99a8f34bfc04067e81ad04fa"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_openbsd_amd64.zip",
     "sha256_checksum": "6272eff9a8f8e6ec668b7a29e032632661ffaddb8e94ae1b4a9f65ac02b2077e"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_solaris_amd64.zip",
     "sha256_checksum": "b3ee1ddd161a19f63530b07a4c096203a15087780484db48bb46263feefbda23"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_windows_386.zip",
     "sha256_checksum": "77491fcec1906654c1acd37902f3250115e3c62634a13f1c0c0f3d7e46124cc8"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_windows_amd64.zip",
     "sha256_checksum": "7710b187dcc84b696a59f196093537c377c1771a1d5239b38f87fb7a805f3760"
    }
   ]
  },
  {
   "version": "0.6.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_darwin_386.zip",
     "sha256_checksum": "e90beadcafc96f5811052c1b081066c6a796520fd4c34e27294435fd236397d6"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_darwin_amd64.zip",
     "sha256_checksum": "43912f5d3eac34a73eaa182a78e13e8392ff4b81f053be4a61cd78db53c505a7"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_freebsd_386.zip",
     "sha256_checksum": "15f28384f43ae6a3b7dc682567a994271cb99c4fd2a17f823730887bc2816f4d"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_freebsd_amd64.zip",
     "sha256_checksum": "e4db1f8aeb60049b35ee9cee5449ab38ff203099c79bd368ae9680d5dde0c3e7"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_freebsd_arm.zip",
     "sha256_checksum": "aca12970d2d261fa5423d4f76cae7d05a8273e2afa8da50580f5b446ba835a7c"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_linux_386.zip",
     "sha256_checksum": "27fbf1818a7aa7848a24d685198c213103f04566e4cf5d4f4f2846bf1da441bb"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_linux_amd64.zip",
     "sha256_checksum": "497e9f9ace4c3da9afd8240222f6f29a7209b3a9ac5b143cefe117e41d5985f4"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_linux_arm.zip",
     "sha256_checksum": "b9e51a93ebdb9262caaba14a67252c3fe0c3366b84d3dcdff7fd6f1edf81d400"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_openbsd_386.zip",
     "sha256_checksum": "30ab9b8b711badce7adbb50414a8e85c00fbef8127a74ce7fc698ba5eeb5ce9a"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_openbsd_amd64.zip",
     "sha256_checksum": "f2d3e96a894cae46cdada7eb2e201e24c8460dc8d756c4f806cd6438fd3d7f4"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_windows_386.zip",
     "sha256_checksum": "f49845bd54bbbd2c8a7e64c21978d4093ff51d336d13200d65484a7550334d55"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_windows_amd64.zip",
     "sha256_checksum": "793fa65210c0da21441876b97a9fd823d829a5a02c2f7a5e363134a36cd95f0a"
    }
   ]
  },
  {
   "version": "0.6.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_darwin_386.zip",
     "sha256_checksum": "fe97b6a0dae5850c0442d0e4bbc507672fb3a4050897b6e0b4b3ac6e4ed461e5"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_darwin_amd64.zip",
     "sha256_checksum": "76a11f1ccd4af7881fab07ba7008a05ddf5ddeb25da2683c258619c9223d8162"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_freebsd_386.zip",
     "sha256_checksum": "fa0f805a124a4f2266e6c3c7721a054f9cf5b4ac084597c4e3756d0d9ea30393"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_freebsd_amd64.zip",
     "sha256_checksum": "bc391d142d8c7aa0cda32681a52c653f6a7060a3f3fbf6db76f41ccaf4398744"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_freebsd_arm.zip",
     "sha256_checksum": "ebfa350f49a4eba83f9e6fbd47ac060602a93c25679d26248295a4d3806703cf"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_linux_386.zip",
     "sha256_checksum": "291b0e1fcba6d81836d12e98c02c8b853ecaf9d0f63c1cb8165366b7a53ca356"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_linux_amd64.zip",
     "sha256_checksum": "36e92a21ca384f48802c43f7334f756267632cc9bb939a11a625e2c414de0360"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_linux_arm.zip",
     "sha256_checksum": "638c84363de77c0d8218325792bf5591fea9c4822379ac168b6a9171e6cd2ab8"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_openbsd_386.zip",
     "sha256_checksum": "f24b2e9142ccf4f5c82592923d88484d5aa6cb4b5c4e1bacc099ab6fcee4f62e"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_openbsd_amd64.zip",
     "sha256_checksum": "27b7816e9021e0adf777cfeaaf84ac1a25079e9e4737898e5f8db5ba11b118fd"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_windows_386.zip",
     "sha256_checksum": "286978840c37a42ad5ee81a038805aae4138ec6b3dfa6d81422ea3848815d5e2"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_windows_amd64.zip",
     "sha256_checksum": "46b961a6edea8f37b8838fd5b3bed8aabcbff4facf426a6034ac5445c4106824"
    }
   ]
  },
  {
   "version": "0.12.15",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_darwin_amd64.zip",
     "sha256_checksum": "c1ec56b36e8395a454b7d0ba421aa42c54d2f91c913893447d20aecf1437623f"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_freebsd_386.zip",
     "sha256_checksum": "d7b378e094445edcf76f1c0edd46382dc0f6a543b5fd91f3a67b4e5d55bbf419"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_freebsd_amd64.zip",
     "sha256_checksum": "72f62e67ee7bf3b5c4bd68f1e653115665b053b80fe7b6b3c1b97f2a507020a5"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_freebsd_arm.zip",
     "sha256_checksum": "7a7836240ad85231336553a980071f3a54399d1cbba009e362abe3bfdcc09cb9"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_linux_386.zip",
     "sha256_checksum": "7de34b846cae49b4606b9fa5aeea02a6f5225a43946a1759b38b83e20f9fbfcd"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_linux_amd64.zip",
     "sha256_checksum": "2acb99936c32f04d0779c3aba3552d6d2a1fa32ed63cbca83a84e58714f22022"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_linux_arm.zip",
     "sha256_checksum": "9c6dbd295ba994bf7df49771622794b2c977786e149eb4f85170efedafc2d7d9"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_openbsd_386.zip",
     "sha256_checksum": "49cb3bca7813f75b3af2e6f9a63ef5aa36d600a078d1f525aaf25f739881a013"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_openbsd_amd64.zip",
     "sha256_checksum": "99f9a9a0b91e56b5b5b64dfd98342ec2ad2c3bef4498a342eb46861bca74d252"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_solaris_amd64.zip",
     "sha256_checksum": "eeecf25c91ba271206077dda78478d74ea13e20954eafa3ed62744ac47930a47"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_windows_386.zip",
     "sha256_checksum": "5ee83859b3805d4f7ab3395d4f32756666f7c74f5b3c9a30ef9a3f03e053325"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_windows_amd64.zip",
     "sha256_checksum": "90abb738d565f040fe627d02aeacc3e839341d80cc5cedb7d915079bec62d9f"
    }
   ]
  },
  {
   "version": "0.12.24",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_darwin_amd64.zip",
     "sha256_checksum": "72482000a5e25c33e88e95d70208304acfd09bf855a7ede110da032089d13b4f"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_freebsd_386.zip",
     "sha256_checksum": "3ec0a88f8138a2947b5636e6227165d0af42000ae8a5e62ca8316086d80dac93"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_freebsd_amd64.zip",
     "sha256_checksum": "3cfd883b3e80d83a6c5f1bf13bca7956b1c7469262354504370516b6fa75f998"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_freebsd_arm.zip",
     "sha256_checksum": "5f7bad18516d6723ba4660f1466305151d339c0f05b28fae30912020034eb51b"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_linux_386.zip",
     "sha256_checksum": "a525cf02c3e267141a8ceb40d136fa6306067216381369b46a2fab63f21ef783"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_linux_amd64.zip",
     "sha256_checksum": "602d2529aafdaa0f605c06adb7c72cfb585d8aa19b3f4d8d189b42589e27bf11"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_linux_arm.zip",
     "sha256_checksum": "fb0e6ae5ea4b416dad4ef5b661115b3951f404fd8a58724f10583de786616d23"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_openbsd_386.zip",
     "sha256_checksum": "c670699ce8569c5aab71f5d7f9fbdca3dc793d8a0d2726853906abcaaf597324"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_openbsd_amd64.zip",
     "sha256_checksum": "c140cdb490167f6e45ecc0a6e987b3b90a5678ed088a82b980f29a95563f3c6e"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_solaris_amd64.zip",
     "sha256_checksum": "18a1489b88e50f358b2ff19a6d4f15941d0158f499e10c37cd67da16796debc3"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_windows_386.zip",
     "sha256_checksum": "2741f6027c7e07fbb51201985ede0e789eeba7777f4ddc065d064499c5dd8197"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_windows_amd64.zip",
     "sha256_checksum": "fd1679999d4555639f7074ad0a86d1a627a6da5e52dacdb77d3ecbcd1b5bca0a"
    }
   ]
  },
  {
   "version": "0.4.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_darwin_386.zip",
     "sha256_checksum": "32749590114402047e4ab9e933c3c8827bdac3c320385ad420bc7e3b546f8be"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_darwin_amd64.zip",
     "sha256_checksum": "317e2b9721394c1f6cc6710f13598cd91e8816b82fdc3781485556cadf1311dd"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_freebsd_386.zip",
     "sha256_checksum": "4f3e3cbd4e5ab5c89c99af95a1ccbe52e1d1da22208e8206e12b23719295db72"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_freebsd_amd64.zip",
     "sha256_checksum": "fa4bce052de38cfd9e7c07d2a7e20004469f9490f748e2fbdba1d30af59fbf0b"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_freebsd_arm.zip",
     "sha256_checksum": "7e184e2da12f9834ce4915f2189d3a65217e01cb81653be8920a6a398f97d08e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_linux_386.zip",
     "sha256_checksum": "b2e315fd434d87860fc907d438786bc12ac3b10a99037d9422eb202ad4d82833"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_linux_amd64.zip",
     "sha256_checksum": "9fc236f86bc6fb002e6058e805ba49c78437cf0c688d26ecfbb21113c9613ec7"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_linux_arm.zip",
     "sha256_checksum": "9377b60ff4405f545e899e449d8667b6d82d149eec32d03882c3320e8b9bd3c7"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_openbsd_386.zip",
     "sha256_checksum": "8c91afb20ea72651fa7b38d59cc9186ede617f6b0b6f708d70b6eab374dfca11"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_openbsd_amd64.zip",
     "sha256_checksum": "bf7717d2f7964734ef92afef445cc641fefe73eec385395b40146b39ffc08c9b"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_windows_386.zip",
     "sha256_checksum": "ee5e7713ad3733d53aa3c2d8832629b88daa22beb42c64b5dd1972fd59abfafa"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_windows_amd64.zip",
     "sha256_checksum": "6a4e6e374f57983ecc205d75e8e7b251a2202a67c4dab84bf495dcff18a1d254"
    }
   ]
  },
  {
   "version": "1.3.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_darwin_amd64.zip",
     "sha256_checksum": "2f407536ec57a1c6ccb6936dee202e676c4d5f1bc085bcbf6a4e8fad1180d879"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_darwin_arm64.zip",
     "sha256_checksum": "3eecf0ac0d1d091f17c9dfedd3c0f7e45de79b30092fa84f0014094b288234df"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_freebsd_386.zip",
     "sha256_checksum": "5012359d543b61f73e34302fb4b81febcbf2ba5df4121e4459d50cbfa371fc34"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_freebsd_amd64.zip",
     "sha256_checksum": "13af370a6a68736e3ca4881d3ab205346101cde500e97f7f8df0e7d9034c87a0"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_freebsd_arm.zip",
     "sha256_checksum": "e96d4dc05296b055816255648ab67b9c20d182a2aa853b9cf0a27f2452e4c330"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_linux_386.zip",
     "sha256_checksum": "3dc4a61c3b70e87faee6dd883e2ebb6ddcec49b095aa1d6431aa974d5251c26b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_linux_amd64.zip",
     "sha256_checksum": "cca0257ad3703aeb7989ed89c9d6d02677b4114f177183990a61446c0e9aa45"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_linux_arm.zip",
     "sha256_checksum": "c1fcc1a29d78e7c228bf7acfa6e547c669f546efad82869e998afb38d16083e6"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_linux_arm64.zip",
     "sha256_checksum": "e7bad31ef48c070fcdc29b73312c6a7ebeb5a8ed1d26eb5f8c5b91e380e6c547"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_openbsd_386.zip",
     "sha256_checksum": "c617adc282aa0002ef87bab99d6dbdc1ae65c4d8834c72ce5931b59e2d7a5192"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_openbsd_amd64.zip",
     "sha256_checksum": "ca02ccb0396cb6e7038adef5047d4f1d7b4991db7b2f2c61195e7435a15748f9"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_solaris_amd64.zip",
     "sha256_checksum": "4107de55fce2bc748ddbecdb78cbdb57c90cb672f130088b4357d4e4447182fb"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_windows_386.zip",
     "sha256_checksum": "d352e74307923e8bfa2768822a614f71a1d659d39f3b0be7435932bb9b6a247c"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-beta1/terraform_1.3.0-beta1_windows_amd64.zip",
     "sha256_checksum": "d6828d058bfb3469d56d4a39177710d4ec75ead399fd5d6892c493b80b2eee87"
    }
   ]
  },
  {
   "version": "1.3.0-alpha20220803",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_darwin_amd64.zip",
     "sha256_checksum": "faa45834e3e9378650e7783bde03048472bf77fe48bfee0539de4f7cf9d0ab11"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_darwin_arm64.zip",
     "sha256_checksum": "821f91c0b80f02d68bb9990c7e2ad125246b04bee8029404b34d33de304f40f"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_freebsd_386.zip",
     "sha256_checksum": "2675ade94cc82ad42d6352e50659a56fb7a4438cd4c1cda852f1857c3ff2fc92"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_freebsd_amd64.zip",
     "sha256_checksum": "7da45aeae7d37a7c2c98b200451a8fd48402de0c02ef8277bf010832b118a68"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_freebsd_arm.zip",
     "sha256_checksum": "a9984b95a690cea75f76484fc72b956c059cd96956c850612d4f46b8d7818233"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_linux_386.zip",
     "sha256_checksum": "394344f5a6a4e8e22448d75412309e4560b9ae911333a04967856f76cddaae4e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_linux_amd64.zip",
     "sha256_checksum": "b65659bde4472402426e1958b7139d50f8eb6b47db554ee8e0f47831ca58d064"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_linux_arm.zip",
     "sha256_checksum": "9c6f114780690ed2fff15ddd989bdfd2f725bb2410a945f97b82610116ccfd6d"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_linux_arm64.zip",
     "sha256_checksum": "8bafd932bcc0d44fdd92e43bbcdf4a6f7a9a6644749b078b3b8ad0433edc3213"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_openbsd_386.zip",
     "sha256_checksum": "2cf901c8545640e2496257bc797313b6f9d9c7e82956612348524b5973d22bb3"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_openbsd_amd64.zip",
     "sha256_checksum": "e61b4246e44004c6ba5e67fea706c50df0a58afe88ede7757470993891a19a60"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_solaris_amd64.zip",
     "sha256_checksum": "2ccf78a9365d2eb20771f4a19315836e5f4032466d236e779c6f853fe71017ce"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_windows_386.zip",
     "sha256_checksum": "ce56a69c5853b046aba8b37f4a84c9e7c42d7a0751f535ebc7779267476d12b9"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220803/terraform_1.3.0-alpha20220803_windows_amd64.zip",
     "sha256_checksum": "f2ecbc8638562e9d82d205fd2328527fe56536673131d64e4482ecb0f17fbf24"
    }
   ]
  },
  {
   "version": "0.9.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_darwin_amd64.zip",
     "sha256_checksum": "f2f4e12bcb6e8bbd8876194221fbb79860ad700926d47a42654a354d70b06022"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_freebsd_386.zip",
     "sha256_checksum": "896135830162b0f4375fea35fe658740778a9805e688964caa7a59702d6c9fe6"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_freebsd_amd64.zip",
     "sha256_checksum": "ffda09bb1a4ff38e34934adffb94de7da388c91093860b7711b778271a35d9ee"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_freebsd_arm.zip",
     "sha256_checksum": "c6b37af049c525020ebcdabdc26040b7439ee3491567b5d971f68dec1b1e2b16"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_linux_386.zip",
     "sha256_checksum": "8b485df05371a71ced45f704fe506b133cc585a3f0a90950743e5c602a1cd159"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_linux_amd64.zip",
     "sha256_checksum": "f951885f4e15deb4cf66f3b199964e3e74a0298bb46c9fe42e105df2ebcf3d16"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_linux_arm.zip",
     "sha256_checksum": "75f8b11040d47c38f46ed593e155c4f3c58c0564120d3d6c8bfb9d7f0d931936"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_openbsd_386.zip",
     "sha256_checksum": "39e923a8e0cc8a7793cfae8a6d41d7423fe4649f2a5749d488726a5d7866d860"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_openbsd_amd64.zip",
     "sha256_checksum": "be93ecbb6afc1c02bcb624647bf0f5350c1096d1d36f2bb4d05ddfe889458e1a"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_solaris_amd64.zip",
     "sha256_checksum": "cc0ee5904dbbee2b1a2e696add59e8d777c173936d3244f156456cb8c5aadd81"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_windows_386.zip",
     "sha256_checksum": "2e2701ac76c07b08d051b4b4557ff2a2c7d5d4b72415afa4cc99f356358df019"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_windows_amd64.zip",
     "sha256_checksum": "1b1632d2f574d3f6033fbfabacc4cef16973c5e32585820c70f9a2e42e526b4a"
    }
   ]
  },
  {
   "version": "0.9.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_darwin_amd64.zip",
     "sha256_checksum": "657d522fc08b6f6fba0c913c9d474a80b1c9c1c6e9a497445455a8ff22fd72b3"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_freebsd_386.zip",
     "sha256_checksum": "d5817739829a8f3f411fbca0a9a300f3c1c0c6d5d385a99b062643889bb86d3"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_freebsd_amd64.zip",
     "sha256_checksum": "525364034c1b13a3a03a1fadfb800663689199155fd478f0c02efa26d5471478"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_freebsd_arm.zip",
     "sha256_checksum": "e3894ee30756444ef7af5648057c5ea053b7fb16309cdbca6e40aceb8a0665b6"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_linux_386.zip",
     "sha256_checksum": "813839f28e04df0de1c62969f018709b4f9457f9db2448351e6b16534f1e3f6e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_linux_amd64.zip",
     "sha256_checksum": "5327b8fa93eb14ef57f4eff94f5380c2fbcefe0f9c135c8ab33425bec4054a6e"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_linux_arm.zip",
     "sha256_checksum": "fe631e30d576d6f3a4d7c03d97a77705790023c77dc4a3fa4373538d1920ff97"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_openbsd_386.zip",
     "sha256_checksum": "4ddbe79cf0e89f057360430dc5ba34930708b1eb5e13a8f8da3e0fa865644d1d"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_openbsd_amd64.zip",
     "sha256_checksum": "80d426a5abb3165bd427bd55570db12e51cb4699847ae62ba3106a6480b05b06"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_solaris_amd64.zip",
     "sha256_checksum": "ae77ad4394c4963fdea0a77f24d2839e52cc42008d83c50ada042f6e0cbae7a6"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_windows_386.zip",
     "sha256_checksum": "c37719e1d53efc4909ed24660f1354ac6ce57b4034bc499320833e4a7103cd3b"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_windows_amd64.zip",
     "sha256_checksum": "96f03e81b49de1b8d2dd4f385680b520bd63d27d318ccd771d5997c41bef5510"
    }
   ]
  },
  {
   "version": "0.7.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_darwin_386.zip",
     "sha256_checksum": "4c1d1df0cbb2252d3f99616214adf6aa6db52bc53965205f5d31912068be20e"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_darwin_amd64.zip",
     "sha256_checksum": "87cae476176b2f4416e5e0eb6c46ff218dd62201c31d3a3dfc16c08849d01b03"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_freebsd_386.zip",
     "sha256_checksum": "79efd4ca311155ac928c57972d41a22177c6fb9ce51349a94e6495d0323f3f7e"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_freebsd_amd64.zip",
     "sha256_checksum": "b913d9bd8b60db29583fd5b4ad76f8b967ae65104bbf56e781b34dd87efa94a5"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_freebsd_arm.zip",
     "sha256_checksum": "fa0d94afed69933673fa9db15c6ca46b008c0059b8e030a1c24f469fc7932226"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_linux_386.zip",
     "sha256_checksum": "83ace94b800db15d32fbb9a6d6f2fc25ed5c0936fa2cb700221dd31bf071a156"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_linux_amd64.zip",
     "sha256_checksum": "7def82015b3a9a1bab13b4548e4c8d4ac960322a815cff7d9ebf76ef74a4cb34"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_linux_arm.zip",
     "sha256_checksum": "9fc85ba1b4a1ace60b1ef5612373f9e9a75159a98793af1753e694830e84a670"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_openbsd_386.zip",
     "sha256_checksum": "d27df20bd79550cdd29c3c569e4b4d5827f0201ee4aafa5b99c08eef87793e97"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_openbsd_amd64.zip",
     "sha256_checksum": "493789384cbb264cb863fcaa989a89e9f7800f460ddd4a7057e00b98c86d4774"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_solaris_amd64.zip",
     "sha256_checksum": "aff343e50d6dd1d1f462c2b10843ed233afb4207c2b8cc995670af560c8e0e93"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_windows_386.zip",
     "sha256_checksum": "c660426f7a7d1173f66e67ba59a9447679f19fd5b309911a7dbaffab4696b120"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_windows_amd64.zip",
     "sha256_checksum": "325f8b931364bde66e19369faa42eb5001ba86024c53654385a206b841c658d5"
    }
   ]
  },
  {
   "version": "0.1.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.1.1/terraform_0.1.1_darwin_amd64.zip",
     "sha256_checksum": "1387eca09fcad8571f02d2f34b79d7cff5f420da8cc52e9b0841696461c99b38"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.1.1/terraform_0.1.1_linux_386.zip",
     "sha256_checksum": "8cc22e4e5782128a509a29db8ec5128439a42aeee58d5f80e69dabc7efbf576f"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.1.1/terraform_0.1.1_linux_amd64.zip",
     "sha256_checksum": "3685ff506135fd29be81c7239fcefb22f0fc5783e28bd308f3101f06f1a6355"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.1.1/terraform_0.1.1_windows_386.zip",
     "sha256_checksum": "b1e3a7d729a29ed3725d3e0feae9d1318e7afdb0bfe5b10b800a979ce5b1f5f4"
    }
   ]
  },
  {
   "version": "0.14.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_darwin_amd64.zip",
     "sha256_checksum": "96d0b1c807415ba295a70e8afed04e233778673103587f321164ebb96be123d8"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_freebsd_386.zip",
     "sha256_checksum": "a5e764bc419afde3fc9e351967b1a6fe175f8597f5ca0dfc066fc60633490e3b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_freebsd_amd64.zip",
     "sha256_checksum": "415544f3e06b90f59110dd3266ce5859498bc21f6fce9e9487ed8a13981922c3"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_freebsd_arm.zip",
     "sha256_checksum": "929e0d5cee00042c514d54021eb8ff3f7f9ab059f4aec9fd602ca1a045f27a42"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_linux_386.zip",
     "sha256_checksum": "c785f2d28d708af8c40e612b5e8bdbb51f01c3e216d5d4a048dcd703770ac26b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_linux_amd64.zip",
     "sha256_checksum": "47e097cfbfb64e97492934f50e646cb84df952eb76897182557811b45603dbf0"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_linux_arm.zip",
     "sha256_checksum": "a58df5be281c7e2c0627cf15f755634201220f0c55ce30c7b37d3fc088bd3ef2"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_linux_arm64.zip",
     "sha256_checksum": "5bcdaf46927edcc46f063faef02878620b137a84a4523004c70b6ab05b20a15c"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_openbsd_386.zip",
     "sha256_checksum": "b62e6ee9ab4bd6c3dc80e1a4db691dc56219c4ae6ec5d9d5b451cf273826693b"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_openbsd_amd64.zip",
     "sha256_checksum": "ccd99a2ab25232fe9d23ae8baf244070b082939c00f66eb17d7ff452327241ae"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_solaris_amd64.zip",
     "sha256_checksum": "feba081708f65bc76067cd56417be8babd2a2a2305440828376ba8e595e87717"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_windows_386.zip",
     "sha256_checksum": "66b508be761319aab846cf9bce3e7f92ebe603b0c3cdda06d1c4d69561bd69f2"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_windows_amd64.zip",
     "sha256_checksum": "8814aacbab1e15c42a1f9f978401500ef8c33e4075f8a29a74e3a67a1a535e26"
    }
   ]
  },
  {
   "version": "0.6.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_darwin_386.zip",
     "sha256_checksum": "3a8daadbf1b2afb39efe4c652e6ca01c59aff53ed1cbb2f876c4c543709df374"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_darwin_amd64.zip",
     "sha256_checksum": "9009582111ba938bd7e22767f533c712fb763dffa9f390b40b17f18742bfac59"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_freebsd_386.zip",
     "sha256_checksum": "c2e0c7524f522c77fc1ed83e9cf7d6231ab0b49e5bef51c0f88fcb92c1c86748"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_freebsd_amd64.zip",
     "sha256_checksum": "9172ca4aba084548439d535f8902f05fe4c69b6adf5c9cf4330a85f9430828b9"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_freebsd_arm.zip",
     "sha256_checksum": "626a93c0817f8ca27e7a6b5c425a42b9e0d7dd4406316ea3cfdf44ae487c0456"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_linux_386.zip",
     "sha256_checksum": "a39f4d835bd1d1ec8a3dc3156699381b1e4b57c6e428921783d7587eeed36223"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_linux_amd64.zip",
     "sha256_checksum": "d7c07e2bf587257673bae710c776430a8cc5a755a9ad4a2cbae066d0cd02a862"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_linux_arm.zip",
     "sha256_checksum": "a80518f28c508b9fdaeb3534e02ad7b7e78e48bb19a3ff8df465f7fd1515ae60"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_openbsd_386.zip",
     "sha256_checksum": "8f4c787933fa55753186e29e8652ee17007fa61f4e4972beaf8dd66e495ed9c2"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_openbsd_amd64.zip",
     "sha256_checksum": "76a1c8a5622e0603dc1962f0aceb432dc36dc9529838b6b827729297db70b9c"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_solaris_amd64.zip",
     "sha256_checksum": "5528244caf580c5dd11b021ae3b7566d1339c0a000cae4e3ae14fff86da0da82"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_windows_386.zip",
     "sha256_checksum": "9d835f0ab6ae05c46be0839d2f4eae9e14183679e36e9532a28db5d0aa397460"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_windows_amd64.zip",
     "sha256_checksum": "6d0b816d5d0d5e8f0549fb983b9b6a566186aec86806c7ae85fbcbd12b2a55aa"
    }
   ]
  },
  {
   "version": "0.12.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_darwin_amd64.zip",
     "sha256_checksum": "cb10093fe8b14771047314b547c7710e363199c40e129bb7e3b4886e3f3b3ca6"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_freebsd_386.zip",
     "sha256_checksum": "a9715bc447f462f8390e0434922a0d0113c1570ca177b0d91a840bfa75ccfaad"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_freebsd_amd64.zip",
     "sha256_checksum": "ca4bfda0e37c482258ca06adba58b45a91e695338f52c0fb309b8787f66461f6"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_freebsd_arm.zip",
     "sha256_checksum": "d2191f12e8cd23c33bc7baacae71ff54b3cf5c8d1f9760a3e1de9224ccb3b0bf"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_linux_386.zip",
     "sha256_checksum": "59ec95e4c6f807699eaf50ebf29baead2cee430a35caf69bc57b7defbe9d0b14"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_linux_amd64.zip",
     "sha256_checksum": "b16b32c39de9597b381c0ba189c08a37ea1c56cbca23122a3675a0481158944e"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_linux_arm.zip",
     "sha256_checksum": "19475addb3b6f210c9f08fc87f48b258095c4e91050a038aa036b9cc96a791a8"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_openbsd_386.zip",
     "sha256_checksum": "700f19b308798b3d3af35920bd1befb6073b002e3a185c8b3dc4b7dc58089a7b"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_openbsd_amd64.zip",
     "sha256_checksum": "b7e2f0261163f8307acf837a316f4ba974f7a875485b9f5b33a92d061c6a84a2"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_solaris_amd64.zip",
     "sha256_checksum": "10a462886264f704d7721b2bf6db578afee3cdc63d9a1c042cad721c95e104f"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_windows_386.zip",
     "sha256_checksum": "ae6390fe0bb1661251cf173e085f26838391d02c107b05df21adbf95408f0f3e"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_windows_amd64.zip",
     "sha256_checksum": "5017ffd23e6cdda3f14820fd26255643fd8eb543df4ce16ea1f2fb7b0d4b57b2"
    }
   ]
  },
  {
   "version": "0.3.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_darwin_386.zip",
     "sha256_checksum": "55ec1ef9fa479e2cc9ac480fc11ee759d49f23ebebcd04899ea9103b09a48655"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_darwin_amd64.zip",
     "sha256_checksum": "65b4c5bfc34bb0464b691b31ac554132c87ac0c5d7acef936c039777a27dccad"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_freebsd_386.zip",
     "sha256_checksum": "d87e4201cd92ea8b65a9acda7060b802efd8c0f13fde6e200853322d08d40e21"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_freebsd_amd64.zip",
     "sha256_checksum": "e027365d3d9d28568c5d9a488d1a862790f61281c466770dbdbeafed9e98a7cd"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_freebsd_arm.zip",
     "sha256_checksum": "cba0151e13083ff2aa4e473c1ff7a4a87a387570e28450e9410bb0ecade2e3e0"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_linux_386.zip",
     "sha256_checksum": "828b7bc13361fd72d3a7c85a1309711409a86e32832099e542d4b7b56508e57d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_linux_amd64.zip",
     "sha256_checksum": "4c5890b098d07e4902e76331b576762632d6443e50030acb818b9a3d642a91c5"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_linux_arm.zip",
     "sha256_checksum": "bd4f9177fc9c0c8a5e26ac4780e0bf59468d3efcd47944cf2919ee0550bb7830"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_openbsd_386.zip",
     "sha256_checksum": "ce11b16d2ecb5375733b904a2cfe5b354f91119f4b50d3218294cff024074eef"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_openbsd_amd64.zip",
     "sha256_checksum": "fc5a78405d50776d21f3def23d087e94e3297954eb8d4c6749824ba8066f62fc"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_windows_386.zip",
     "sha256_checksum": "59dacab10d70e92e3a58dbf2f6e8fac5af225b9b12e58bc622691540e9ad0125"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_windows_amd64.zip",
     "sha256_checksum": "a8c5571647f66784e81cc45fe8804464a19a02959f011893ef3b8f1b063fcc0a"
    }
   ]
  },
  {
   "version": "0.13.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_darwin_amd64.zip",
     "sha256_checksum": "cbb76aed9c01a8c0fbee4e3a10112ab7836440fa63d93414a1dc45ef59bc0ea2"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_freebsd_386.zip",
     "sha256_checksum": "35124251a059b97e8aadb1263337ff4d5687f6d10c0a4bfbb6edbfd6112f0fbd"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_freebsd_amd64.zip",
     "sha256_checksum": "470d349391e33efc0d1654ac7e30d0f99d78c7592d7791ab0d4414fe3601667d"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_freebsd_arm.zip",
     "sha256_checksum": "9b6fce1bbd57fa439674a163940f6fae254c0a011f2b5bbe378758a89e793c78"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_linux_386.zip",
     "sha256_checksum": "3dfab74569cadaafd170d8603f8eecaac372b9b1a3b6b6357e2646b168b38d74"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_linux_amd64.zip",
     "sha256_checksum": "55f2db00b05675026be9c898bdd3e8230ff0c5c78dd12d743ca38032092abfc9"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_linux_arm.zip",
     "sha256_checksum": "9902e3e8260368ffab0756caf64e4433dbfd31792f58d78eb793d5ff7674534b"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_linux_arm64.zip",
     "sha256_checksum": "b0bb7c6a8ff6a81af7bcc8b6261f393e0a0b53900647a472a700bc25c7f4a262"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_openbsd_386.zip",
     "sha256_checksum": "3dda14fd5774f7de8ad95c4ada8dacfd6730ce6aaf94562096faa2f47e684371"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_openbsd_amd64.zip",
     "sha256_checksum": "995c4f114f36fa66ca0d76f57609bbd82d6af9c931ee022e9409fd1129bbc134"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_solaris_amd64.zip",
     "sha256_checksum": "e30e2409f9a91b90f86d6806a48321533117b06f2b3b291d5ab30f19ce068e90"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_windows_386.zip",
     "sha256_checksum": "6aef930a0c3a0f798876040937a005027cd2be7bb6ac8be1ad09216cfb60d56c"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_windows_amd64.zip",
     "sha256_checksum": "62e0586c868ab9cea12847d8e040c8e4d7a97df68f043ad5019ace6863627947"
    }
   ]
  },
  {
   "version": "0.6.14",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_darwin_386.zip",
     "sha256_checksum": "1203104b5b989ce29d5fac565a8753547f8feb9baec1d0cf43d4c0fa52b6f1e7"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_darwin_amd64.zip",
     "sha256_checksum": "9334f55a549d5cb3c583430be15e73b407bd7e115dc53db290381a482da17788"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_freebsd_386.zip",
     "sha256_checksum": "872a68e7f6aee7c06f6fe4f62ac24909a7f3c5a91686cb5688e49b50dab751a4"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_freebsd_amd64.zip",
     "sha256_checksum": "4062260c21c6b587541953b66919e5933e4575a8c0a03a8e1cc1b4b86eb9158e"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_freebsd_arm.zip",
     "sha256_checksum": "be14951a6d96d525668d28078adafcb5b38052083b8515850a20197313ea383f"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_linux_386.zip",
     "sha256_checksum": "3f43b75a5376b21576ff341af97f57857b5c8cc306331a4917f7d76895a97b13"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_linux_amd64.zip",
     "sha256_checksum": "6d93290f980df15a453e907ea9a2ca8f8fed41143c220953c911b5174c3e3ab0"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_linux_arm.zip",
     "sha256_checksum": "2a2f7255aaed452e710d36bd1bebc2f44a320215e575bffa74cb737a1abd20ba"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_openbsd_386.zip",
     "sha256_checksum": "bd39f4917f7ed015bea2717962931e65eacaa46a25dfb08a73a8106d144f867c"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_openbsd_amd64.zip",
     "sha256_checksum": "e69cd348a042bfb45745409a32d8ff42b99d3dcde33514b70b35d41d2293dcb9"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_solaris_amd64.zip",
     "sha256_checksum": "3293133a4b0a014e9d147d7b4b0b5a83e73635c40a7307977dc573a8d3b27277"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_windows_386.zip",
     "sha256_checksum": "46a2e0a28b862cdaf919ad5596dc85866c5616c2510b6de2a6a1c9f95a8c7979"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_windows_amd64.zip",
     "sha256_checksum": "3cc26f06df35eff3c222665883720ee738bf766a9025c50f1e2e45c1973aa381"
    }
   ]
  },
  {
   "version": "1.1.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_darwin_amd64.zip",
     "sha256_checksum": "c54022e514a97e9b96dae24a3308227d034989ecbafb65e3293eea91f2d5edfb"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_darwin_arm64.zip",
     "sha256_checksum": "856e435da081d0a214c47a4eb09b1842f35eaa55e7ef0f9fa715d4816981d640"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_freebsd_386.zip",
     "sha256_checksum": "de15651d67c8eb2d0ba64776f851a2b9fbf288c22fc8e363f2f2199ec4d4d104"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_freebsd_amd64.zip",
     "sha256_checksum": "c7ad1dba43a1262bfb23e40992f8254879a4ba0f7401c02ed728bf4ca19310f5"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_freebsd_arm.zip",
     "sha256_checksum": "1f44fa6cdcc76fb3c559a6e6186d128acb6981c9eb46b5a5f441046c3f52205f"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_linux_386.zip",
     "sha256_checksum": "e7baa3c9fbd74231f97d44a62e98540fd9fecfbd5214e49d53f03e73e214275d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_linux_amd64.zip",
     "sha256_checksum": "b215de2a18947fff41803716b1829a3c462c4f009b687c2cbdb52ceb51157c2f"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_linux_arm.zip",
     "sha256_checksum": "a245b86fe0d11348543c0991321fd7231ce31f1ce3f36bd7337f5929ddb34348"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_linux_arm64.zip",
     "sha256_checksum": "ad5a1f2c132bedc5105e3f9900e4fe46858d582c0f2a2d74355da718bbcef65d"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_openbsd_386.zip",
     "sha256_checksum": "5e24857a354fda6cf70a35ce0ff6fee254b998f637cb787c17a9a3edd1c083dd"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_openbsd_amd64.zip",
     "sha256_checksum": "aa82f1f61728ba321cc0940dbda84fc1ca9f7acd15831a825190b65dcc5cc808"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_solaris_amd64.zip",
     "sha256_checksum": "effa34306bf003b9a31aa5d889a815d152b0ae496268489e0cade71c0bdcc60c"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_windows_386.zip",
     "sha256_checksum": "e9c41f7141ae511f05e13067ccac4158d78a1214d535d0bda4a549529d4ef0e3"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.3/terraform_1.1.3_windows_amd64.zip",
     "sha256_checksum": "b2890c05c04cfa3cad3c3d5316e6670b963a347ad42ba02fede55268148bce5a"
    }
   ]
  },
  {
   "version": "0.14.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_darwin_amd64.zip",
     "sha256_checksum": "8a5ec04afcc9c2653bb927844eb76ad51e12bcaec0638103512d7b160dd530ea"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_freebsd_386.zip",
     "sha256_checksum": "ffab3dd13699c687c18a92b6cc3cd1fa2acb348d01a45acfc8a5c77ce2be2891"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_freebsd_amd64.zip",
     "sha256_checksum": "9b737d344950ca923420600b819f552aa443bb02e8057510c7998e2bf8c333ad"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_freebsd_arm.zip",
     "sha256_checksum": "fb80439e07bdf2f54ff9ab71e80e33f3d6c4eb6cc31f4b8b4e7dbbe2d3450517"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_linux_386.zip",
     "sha256_checksum": "4d9512c0bdfae7f64da9e7a9bd61b3e8284cbda9316ef8e10113a6aa41e27a51"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_linux_amd64.zip",
     "sha256_checksum": "6b66e1faf0ad4ece28c42a1877e95bbb1355396231d161d78b8ca8a99accc2d7"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_linux_arm.zip",
     "sha256_checksum": "28a349b39440ef60de452ab9ac04122f60fbf79536c9b812e8a8ae23442a9869"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_linux_arm64.zip",
     "sha256_checksum": "a621f1dc411953b955aaf2d7d46b2f350bd3a85a2284ec994ae41419844120b"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_openbsd_386.zip",
     "sha256_checksum": "de8e110c8672791a8143bad9c4e8b0166c376a7f98324feade9392bd6c238fc"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_openbsd_amd64.zip",
     "sha256_checksum": "339185dd1617545f667df504a2c098414d0cfab78c83b536075c1962e53ac861"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_solaris_amd64.zip",
     "sha256_checksum": "1d6e124fb77525137e1c5af49777227a337ef464b04c073e836a2b5bc8bca4eb"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_windows_386.zip",
     "sha256_checksum": "d7e0c139803e74edc94ae2e60f125456db557908554089d0df684982ef461526"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_windows_amd64.zip",
     "sha256_checksum": "1cc49c7522d3a6a583ad627aea2d2b4fb182312f4f97d70d445e2345e4a4f4d4"
    }
   ]
  },
  {
   "version": "0.7.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_darwin_386.zip",
     "sha256_checksum": "c5fd08d0db6a1e8b4807f4cbc7f07f6bd0806824d74ac760f02fdd46bed1433d"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_darwin_amd64.zip",
     "sha256_checksum": "4720e4b2878b3b0d3d781f68ff363707ed42fe39cb89e2e34c6c11f8e0f76b04"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_freebsd_386.zip",
     "sha256_checksum": "5e45ce00a36fc75c2106fb7fa0ef3cf59799167dd33aa2a0b9390f4d7a964da0"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_freebsd_amd64.zip",
     "sha256_checksum": "2304f3503d1dcdc70bbb748d3b41f37f76708b19fbd3be35ddda58cf9b6462f6"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_freebsd_arm.zip",
     "sha256_checksum": "ea52aeac398698708aea36e0566f854efc4a4b4d5bbe63d654ca71e338e2e5ce"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_linux_386.zip",
     "sha256_checksum": "e22a756ffaec39a6635f253ebd8389ed19dd13c0e7151cdf11779cffa411754b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_linux_amd64.zip",
     "sha256_checksum": "a196c63b967967343f3ae9bb18ce324a18b27690e2d105e1f38c5a2d7c02038d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_linux_arm.zip",
     "sha256_checksum": "b48034c82e38598a944235d77bd6a114ff19b0cbed4bf851891f1e9543ba0f02"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_openbsd_386.zip",
     "sha256_checksum": "a41e7904612101b377a99c7cfcb1cf7596206cb1684971c37003b5544934a4b8"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_openbsd_amd64.zip",
     "sha256_checksum": "ca2627816c8ee80862a5be5061d97d8f34567b44beaa89c246bb924bb1ceb4e9"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_solaris_amd64.zip",
     "sha256_checksum": "7d79216e7b9f3892140be438114868807449138c748c38a38c8be49f1f778e18"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_windows_386.zip",
     "sha256_checksum": "8d1909a03899fac1dd2e7f91b8336a8e8a09445009d059de2373b4a067874aa0"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_windows_amd64.zip",
     "sha256_checksum": "c891dffeb8f82e1b96c8d2e9777158b99442754866a876fdeffed3557651bcc3"
    }
   ]
  },
  {
   "version": "0.6.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_darwin_386.zip",
     "sha256_checksum": "cef10c71b7337cfe85c895a0a42d4f512c99f7944a06102b578f73aba32b45a7"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_darwin_amd64.zip",
     "sha256_checksum": "9802b1d56576bea86e34fd3800e100eb043ab6de5a5fa40f7f05a0a44f364dd2"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_freebsd_386.zip",
     "sha256_checksum": "2eb8079176fd173803e06159c1901707b00045dff4e5ea175166bcc957726733"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_freebsd_amd64.zip",
     "sha256_checksum": "a8d28c82dfa9e0f6503b6c2a840d8e373f9cf54437db08d935be92152719ba07"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_freebsd_arm.zip",
     "sha256_checksum": "2806bbadfe70139f60beeda381b2976f75ce8682eaef294b77fcbeeb332e2d02"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_linux_386.zip",
     "sha256_checksum": "8b3bf0e1ac3180f3846abea9bfb3a30eb098460f7923989b5eb4bb8cd4ae80fc"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_linux_amd64.zip",
     "sha256_checksum": "f451411db521fc4d22c9fe0c80105cf028eb8df0639bac7c1e781880353d9a5f"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_linux_arm.zip",
     "sha256_checksum": "7a906ea4137f590afd650aae4bc121e52ae40d5a7f7e57a2ba5c0c3bd316768c"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_openbsd_386.zip",
     "sha256_checksum": "9f04586b066c7a3496e07aaafb1dd1bb470e049ad812657f1c1c59259be69a93"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_openbsd_amd64.zip",
     "sha256_checksum": "a57fbcff72ad17c3261b272b4fc9ae15e059ff186d4aa9064ada43613ada3f23"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_solaris_amd64.zip",
     "sha256_checksum": "738fad4fc55d0120d4bc0ff47fb57a45376a45cf7d8f939c03e5becb97f30cd"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_windows_386.zip",
     "sha256_checksum": "8564488b4a8e8e40d37c63424006315b89c13096357411a21ce84f26ec92767f"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_windows_amd64.zip",
     "sha256_checksum": "4832e6ca2b60bb3d08ef69497bc11890d3e986f0ce76d20bed1a2e9520fc93ba"
    }
   ]
  },
  {
   "version": "0.10.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_darwin_amd64.zip",
     "sha256_checksum": "a37f190cfcac21fa2343ec7e3112137d27fb9286c9f5c128547c6221502442c9"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_freebsd_386.zip",
     "sha256_checksum": "5cb7f922aa7869f78493e7e1060b867d656b4348f0c751e738e9d3677d27c69d"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_freebsd_amd64.zip",
     "sha256_checksum": "9027f2b119190f4316bb85a574cc6d2a40f9ab7e37905feb098d0fbeadf9a2a5"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_freebsd_arm.zip",
     "sha256_checksum": "fbeeb27650f924fac8844f088be5a32ae34494cd2391cee9d33595c8807539db"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_linux_386.zip",
     "sha256_checksum": "265aa867dd1c2a043f72495325e15f1f44859b17a6ddf6f5b40c27ea557a375e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_linux_amd64.zip",
     "sha256_checksum": "fbb4c37d91ee34aff5464df509367ab71a90272b7fab0fbd1893b367341d6e23"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_linux_arm.zip",
     "sha256_checksum": "f512585594deb565b9580402cc89cb26398c29b4a0ca8c2681b0cc9ab08fa19"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_openbsd_386.zip",
     "sha256_checksum": "c50ccab00edb2355d71329558af62b0b3d6011ee58977cdf4bb4e1d1409809b9"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_openbsd_amd64.zip",
     "sha256_checksum": "d02bb17c08b5ace5454aa4223c7511f322e0373c015b1c806cf0552bfc0b758a"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_solaris_amd64.zip",
     "sha256_checksum": "28d30ec6f4618cb1b8f0a4db9e6ae9fd4f7a20085e9573d9bf097e43bfe11ee9"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_windows_386.zip",
     "sha256_checksum": "46ec8e6f4c7307a0c1edbf64904fdddf7ec9e575f579d3d45fa4f72c9fb5a2d7"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_windows_amd64.zip",
     "sha256_checksum": "6cfd7df0b0d394ef55339e023cb3cb8379277575b891130ca499ea59819913a6"
    }
   ]
  },
  {
   "version": "1.1.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_darwin_amd64.zip",
     "sha256_checksum": "7d4dbd76329c25869e407706fed01213beb9d6235c26e01c795a141c2065d053"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_darwin_arm64.zip",
     "sha256_checksum": "723363af9524c0897e9a7d871d27f0d96f6aafd11990df7e6348f5b45d2dbe2c"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_freebsd_386.zip",
     "sha256_checksum": "1bbc51d01710834a0a43d112a85c6879cf7d748cd4d73eec1a17eda8c794a718"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_freebsd_amd64.zip",
     "sha256_checksum": "db605c43318329413b606d77905c50e6cdeaf0447b3e394b69ce45fd97c1ce04"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_freebsd_arm.zip",
     "sha256_checksum": "950be78ea038a8448a943e54823e9c253973022377937656758d5fe98e317442"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_linux_386.zip",
     "sha256_checksum": "c1f3c5a1415838a41c646ed19b76daf8224a9d034e530f96ceba78b7dba27cc3"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_linux_amd64.zip",
     "sha256_checksum": "30942d5055c7151f051c8ea75481ff1dc95b2c4409dbb50196419c21168d6467"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_linux_arm.zip",
     "sha256_checksum": "7b58756f4123ff6fe8680aea0bde91d12ecad28b1a1d25e33465c812f88fe387"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_linux_arm64.zip",
     "sha256_checksum": "2fb6324c24c14523ae63cedcbc94a8e6c1c317987eced0abfca2f6218d217ca5"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_openbsd_386.zip",
     "sha256_checksum": "f51db36fefab2bac9e1149cd4a87aa645788380e07fab48ffa2a1c834281e044"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_openbsd_amd64.zip",
     "sha256_checksum": "3fad034ec87e8cd1b39046d4354760fc5040f2965a82f54b710079b8dbcbfef3"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_solaris_amd64.zip",
     "sha256_checksum": "829c1c65fef51e1bcfecd2930f98859e7bc38d82751cf6b6ba11dd2cbb110650"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_windows_386.zip",
     "sha256_checksum": "986c2bba3a01875b4f03c5d24d899a2ca55337309e40d7a74c0ae112b2d8b31d"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.5/terraform_1.1.5_windows_amd64.zip",
     "sha256_checksum": "d56b2699c2cdb61408ccceba6bf1acea7a792f5a5024f50afb203de363f3d869"
    }
   ]
  },
  {
   "version": "1.2.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_darwin_amd64.zip",
     "sha256_checksum": "f608b1fee818988d89a16b7d1b6d22b37cc98892608c52c22661ca6cbfc3d216"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_darwin_arm64.zip",
     "sha256_checksum": "d4df7307bad8c13e443493c53898a7060f77d661bfdf06215b61b65621ed53e9"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_freebsd_386.zip",
     "sha256_checksum": "18d1fee983ff3d4f5b781dad7c7a690484b175d04b9c71889132310b747b4ad5"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_freebsd_amd64.zip",
     "sha256_checksum": "ec923c2cc6c928d425f3dc3465937bf4526f8b03759dba4ac53dc67075ef306f"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_freebsd_arm.zip",
     "sha256_checksum": "3c8a85d2564f820a97f2bd6902a526c8289725ffaf8acd14ef8eeb5eeb657f06"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_linux_386.zip",
     "sha256_checksum": "7dd896d6ec7a57e959e82f8a3e1ee6786674e95b708dfb4dc399b8b814529c19"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_linux_amd64.zip",
     "sha256_checksum": "b87de03adbdfdff3c2552c8c8377552d0eecd787154465100cf4e29de4a7be1f"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_linux_arm.zip",
     "sha256_checksum": "63d985bb18da6ab4e6baa5778726ad97403f3853ac3caef22fdf39baa6c7123e"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_linux_arm64.zip",
     "sha256_checksum": "ee80b8635d8fdbaed57beffe281cf87b8b1fd1ddb29c08d20e25a152d9f0f871"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_openbsd_386.zip",
     "sha256_checksum": "72e15b83012d14085cc24664fc56fe94b4ae252e9793aae272d9c7416770f70e"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_openbsd_amd64.zip",
     "sha256_checksum": "83117ff4a0e132001c39313a20acb92e591acbdfc6bfd453209a8f191ab061d4"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_solaris_amd64.zip",
     "sha256_checksum": "ad7fe5b33146f6bbec40f166b6221d60e4f10c03a778d4ce3731a42758f36750"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_windows_386.zip",
     "sha256_checksum": "66d7d530d1b80d39feb9ae8a26f444fe91df63cf295b6f7cfcbdbc3c96a5beab"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.0/terraform_1.2.0_windows_amd64.zip",
     "sha256_checksum": "de88c66fbcdca0561d4e6d8d8090d697e97de9bbd97c212ceed21d942e33f937"
    }
   ]
  },
  {
   "version": "1.1.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_darwin_amd64.zip",
     "sha256_checksum": "c902b3c12042ac1d950637c2dd72ff19139519658f69290b310f1a5924586286"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_darwin_arm64.zip",
     "sha256_checksum": "918a8684da5a5529285135f14b09766bd4eb0e8c6612a4db7c121174b4831739"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_freebsd_386.zip",
     "sha256_checksum": "a5890d9c9f08c9160b37e3156ff2a1bc33de1db68ee942f12c4f60e8e74c8e02"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_freebsd_amd64.zip",
     "sha256_checksum": "c204f1ca8162feb59d39bf905d8a1d7687a72b2884d81214ced8ac327908352e"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_freebsd_arm.zip",
     "sha256_checksum": "c27e4b9d88598a55fe5dd0e79746e6b77eb582e12aaf4689935d0c16aa9ceebe"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_linux_386.zip",
     "sha256_checksum": "a29a5c069e1712753ed553f7c6e63f1cd35caefee73496210461c05158b836b4"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_linux_amd64.zip",
     "sha256_checksum": "9d2d8a89f5cc8bc1c06cb6f34ce76ec4b99184b07eb776f8b39183b513d7798a"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_linux_arm.zip",
     "sha256_checksum": "800eee18651b5e552772c60fc1b5eb00cdcefddf11969412203c6de6189aa10a"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_linux_arm64.zip",
     "sha256_checksum": "e8a09d1fe5a68ed75e5fabe26c609ad12a7e459002dea6543f1084993b87a266"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_openbsd_386.zip",
     "sha256_checksum": "b7b509b5a0bae6d1f7e2a61d6e4deccba41e691204148f9451efe353e15ece2d"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_openbsd_amd64.zip",
     "sha256_checksum": "c702a8b31d90c9ced4b95e7facc8d7828f2a31453acc9fc258b9fffeda5ded52"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_solaris_amd64.zip",
     "sha256_checksum": "704190dfb5cd923c2949787505f72227b2b090674f1c8ce941ca180d82d7a4ff"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_windows_386.zip",
     "sha256_checksum": "fd2b9bc7506a85f5293d0e2d12ab5ac3be34b5915f2ae7ae7dfdc178e0abad94"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_windows_amd64.zip",
     "sha256_checksum": "ab4df98d2256a74c151ea7ccfd69a4ad9487b4deba86a61727fb07a1348311cc"
    }
   ]
  },
  {
   "version": "0.12.22",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_darwin_amd64.zip",
     "sha256_checksum": "13d0dd4a4c7cb5dea403c1a02dd9200ff9de086e8ddd832ffea2219c59d33fe1"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_freebsd_386.zip",
     "sha256_checksum": "721be38bcb2832067f503ee110166d65d9e2b5f813b6bf4413687f5f4359fa64"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_freebsd_amd64.zip",
     "sha256_checksum": "8635fdc6e6f95e98e657495eed337b6d160117e695bcf615ca0f0cc822ff08bc"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_freebsd_arm.zip",
     "sha256_checksum": "b31125ea3ce1c795f452b2a9bfa03d237c7a74207c00ae90b4b12ef081f434ef"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_linux_386.zip",
     "sha256_checksum": "15db45b6998e50015666a354424cec196c892db9533c78118a00985ce6b57899"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_linux_amd64.zip",
     "sha256_checksum": "98faa10a7da2055561339ce67436b37c0c760084e51f219102d87b155be088c5"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_linux_arm.zip",
     "sha256_checksum": "5dd7687e3470de3dd08e090d703e86b0aa9fcc797f25c1e03cb06ccd0d2705ef"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_openbsd_386.zip",
     "sha256_checksum": "7f657c2bfa52b104e5e75f4a5249d5d9ff07ae9b079a391cb014c8f19cc8f652"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_openbsd_amd64.zip",
     "sha256_checksum": "fe8437c16cfd6ed0d3d1f37779dee920772d878003c5dccec5ebff122d3c3c98"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_solaris_amd64.zip",
     "sha256_checksum": "f191d51631cc56c265ebf2605a05f2abf8839bf57fe2a0ff21c290f5937560e3"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_windows_386.zip",
     "sha256_checksum": "23705e8aac3b3b73d5f5dc0246522591b5755aa9c86db9352d8b946da7f192ec"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_windows_amd64.zip",
     "sha256_checksum": "e45fb885722d9cfaa3f45bc2dd7d9f0d5093e1d8e6ce74c58cbc48587ec4b558"
    }
   ]
  },
  {
   "version": "0.15.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_darwin_amd64.zip",
     "sha256_checksum": "dd7220e6a76e4c9555576c0500bea94c6a5cb65f286b3e74e8ea7cc34bbce5be"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_freebsd_386.zip",
     "sha256_checksum": "72eacd82ced71d667bde02dcdf06dfe7fb30b762d7989651e2a63dde7df873f2"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_freebsd_amd64.zip",
     "sha256_checksum": "13e24c1ddced0316717701a8e5488b4de55b3005ce88061e4701fe4bcfcbe7cd"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_freebsd_arm.zip",
     "sha256_checksum": "1b7ce3dbc83d22a6a9fb98eb0009498745e43d54737f805b0c25b81a37121e0b"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_linux_386.zip",
     "sha256_checksum": "5e88ed699a22ebe5b9276a77bd1c91937a2053994f8997a3fdb3e4e5a815fe4c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_linux_amd64.zip",
     "sha256_checksum": "1ff798791abf518fb0b5d9958ec8327b7213f1c91fb5235923e91cc96c59ef2c"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_linux_arm.zip",
     "sha256_checksum": "e39b377edf4cf029ecd8ae38c248a3f6a3921d8bfaa1ca0dcc5c960c1b4945de"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_linux_arm64.zip",
     "sha256_checksum": "39aa72b66b5802a2909a17ed14ffe7bcf3b6466a91710989c4e6cb854d594538"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_openbsd_386.zip",
     "sha256_checksum": "9f77cf310225b229f67281dd53319c85bd6a624c9b3c0f0efe81926f64515fdd"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_openbsd_amd64.zip",
     "sha256_checksum": "2df858639d23ce9a3c64a2a3516feedc370566f8ed917e26d1362ea73b2033e2"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_solaris_amd64.zip",
     "sha256_checksum": "9d8c1e348b91b3409b5953a26a1465305c0b46edc76f034ebd1aa03fe1beaff4"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_windows_386.zip",
     "sha256_checksum": "5f31c483516e98bf1fb21eb1455ea55b8ff7fc98660854cf695d3be7d6763ad6"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_windows_amd64.zip",
     "sha256_checksum": "c04be77792cfd57e2583830e13712958d4739043ca5d6715857f187491a2ab39"
    }
   ]
  },
  {
   "version": "0.9.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_darwin_amd64.zip",
     "sha256_checksum": "180afdeb14f4049f3374fe02b9143ad428ebd31dd89c6595775d7ba439d7fbf0"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_freebsd_386.zip",
     "sha256_checksum": "733aa19cb6a257bd35dfeb7b79520cfd261d88b022089a02035d3e7f3efca095"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_freebsd_amd64.zip",
     "sha256_checksum": "8a8ea146c973a2d254f90cb92f0a79c5d41d1c8f974d7eb20599f293e53caae3"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_freebsd_arm.zip",
     "sha256_checksum": "2d3d56064281063ba08d7440ce92157d0b8d817013579cb587e6aa71260c94ed"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_linux_386.zip",
     "sha256_checksum": "5285a7c1d2ca329632eff5e4d853ff4a0a21dde0e940e0fc885a8acb659f4f3"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_linux_amd64.zip",
     "sha256_checksum": "f34b96f7b7edaf8c4dc65f6164ba0b8f21195f5cbe5b7288ad994aa9794bb607"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_linux_arm.zip",
     "sha256_checksum": "69b59f4d03e20b90b77822e011bd42e38c3a23ddcaacd9837e81fa3a8c9f7e57"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_openbsd_386.zip",
     "sha256_checksum": "ec9a06458672ad7f0223d9ac0af991d6d39b331ce2b66a17995e43da43592015"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_openbsd_amd64.zip",
     "sha256_checksum": "2a849ea234e560044ded088312d6a49a3cb802348af8c12751b8262cc585a802"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_solaris_amd64.zip",
     "sha256_checksum": "499a9c1f7c4b2ce4e104dfe76ebdc6013752c7d7edf8032d5c2f3ea667977365"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_windows_386.zip",
     "sha256_checksum": "2af313c55c72fa10c1536dfb62895bc8a9351c383ceb24cd6f4565737d03bc83"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_windows_amd64.zip",
     "sha256_checksum": "27a67d3127d762a95024b354a96f36650c340e837566bd296c7ec3ecc15be94a"
    }
   ]
  },
  {
   "version": "0.9.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_darwin_amd64.zip",
     "sha256_checksum": "73ec3c66a77e0c0879e6397fe2b4c4910b24464971fd0c27795b0fa09143f9ad"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_freebsd_386.zip",
     "sha256_checksum": "c1ae0f9333ad8586add3cad8a16b4f88a13a09994564043742d5da1dc5601441"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_freebsd_amd64.zip",
     "sha256_checksum": "cbcb8b7e48e5657624def1d397d16b3c4a28a8c64345b7c2e218a7d5b7e47d76"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_freebsd_arm.zip",
     "sha256_checksum": "fbb356dded39a48009a6dcab46617bfde75ddc43a22ac75f2a4706fda631a8b0"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_linux_386.zip",
     "sha256_checksum": "1e589601c5bc50d283afbee78d47af0cc1966445a45ff178cd1ef25aac210752"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_linux_amd64.zip",
     "sha256_checksum": "cc1cffee3b82820b7f049bb290b841762ee920aef3cf4d95382cc7ea01135707"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_linux_arm.zip",
     "sha256_checksum": "705d13707387864a93e1528658c3affea020cb751b08549b8b35b4772d2a5326"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_openbsd_386.zip",
     "sha256_checksum": "74150634c3eb7ada34605da6315f12d83b5a304805d67c1e06e9c6f7c60a9417"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_openbsd_amd64.zip",
     "sha256_checksum": "48d95958c22fceeb8b836f254bb2af0785bcab639a990339fd54cefa9a5ded6a"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_solaris_amd64.zip",
     "sha256_checksum": "625203ac3d89965f7240c8a8a9d83ed471abe3c6019ffc3ec3a2fd2f5bddbbf8"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_windows_386.zip",
     "sha256_checksum": "e5a7c38c9ce64458bd7cd0c258bc00df83921ece30f668e135d2bfc051aec642"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_windows_amd64.zip",
     "sha256_checksum": "d7c220da2f0b9a52d686cbb7278d27fba3af8dd02d785dfe33c4c6e96df27100"
    }
   ]
  },
  {
   "version": "1.0.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_darwin_amd64.zip",
     "sha256_checksum": "80ae021d6143c7f7cbf4571f65595d154561a2a25fd934b7a8ccc1ebf3014b9b"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_darwin_arm64.zip",
     "sha256_checksum": "cbab9aca5bc4e604565697355eed185bb699733811374761b92000cc188a7725"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_freebsd_386.zip",
     "sha256_checksum": "b5dd00d3e8a072da806e2223a49a9f4292e5c978c25ec5b0dc45d7e38904c43b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_freebsd_amd64.zip",
     "sha256_checksum": "330f815036288ee1f135e26d2dff4f262429f8039fc9945f9e081ab1b355daf3"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_freebsd_arm.zip",
     "sha256_checksum": "76e1d92c0580aef177618552e4bd3103ed351c876a8c157cb4737b29297524c4"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_linux_386.zip",
     "sha256_checksum": "1e171065bde9ce7758f75b3740674b42db1ff1f6ffc0692d9905813a9c381263"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_linux_amd64.zip",
     "sha256_checksum": "bc79e47649e2529049a356f9e60e06b47462bf6743534a10a4c16594f443be7b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_linux_arm.zip",
     "sha256_checksum": "be8d9de8c34e3e843ff8cbc713b41ce0c2bd97b0020c80e934443976b4429ae2"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_linux_arm64.zip",
     "sha256_checksum": "4e71a9e759578020750be41e945c086e387affb58568db6d259d80d123ac80d3"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_openbsd_386.zip",
     "sha256_checksum": "c813de2db415f5df576366f9dbbaeb62f0bea56d5b5ad0b9e47f469a18cb28a4"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_openbsd_amd64.zip",
     "sha256_checksum": "7b97be6c3eb56fe0f0f1eaa8104dd24308cb86785c8404c2df566d3907dcc634"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_solaris_amd64.zip",
     "sha256_checksum": "257fb6e57c2e3353445b52b5ab9cb7caf5e1c6680cbe8f158300705ee7a6e215"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_windows_386.zip",
     "sha256_checksum": "64b930ce06f85c214d1830a213a7528a5d988bd09779729d5b33956b962ec708"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.7/terraform_1.0.7_windows_amd64.zip",
     "sha256_checksum": "88b8a4ae66367d9662a75599421e0bccbb70e6af92d64c62e91aff883c9eccc6"
    }
   ]
  },
  {
   "version": "0.7.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_darwin_386.zip",
     "sha256_checksum": "2e0fb8ad963fb30fca5ebd14a98dfb97381fde2f0fd72d9858599b4cd090cd55"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_darwin_amd64.zip",
     "sha256_checksum": "5c315498c58700d5e0eeba205c1e07e5299d04dd0f7fb7e87e4c38a8c9903774"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_freebsd_386.zip",
     "sha256_checksum": "6cdd0e97edae3628927f3ccf7caa9d2e5613bee2cebef0320cbdeb7e9eab7e1f"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_freebsd_amd64.zip",
     "sha256_checksum": "44c75f355148570c79f80a26fc1f57eb131b29e42dc42b7d1ca36a2857f47210"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_freebsd_arm.zip",
     "sha256_checksum": "47da7b2a2a011d158d9aa0793c6fb77eb86a04803e811692d650954e01e53800"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_linux_386.zip",
     "sha256_checksum": "efe67405000ae7134ced760312129501679535ea5f67bf23c38f698daccb7105"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_linux_amd64.zip",
     "sha256_checksum": "86ccdc048295dcc47f0a3611e39226b8961e08bdfa2969f495c89591c62fd2ed"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_linux_arm.zip",
     "sha256_checksum": "11f81b02aeb4b45abd4d4e8eb7d8dc797f7774e659379c75194ccde7ed930d35"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_openbsd_386.zip",
     "sha256_checksum": "505237ebdd3250398cd7c1f96bfb13393a31485ca0f4d1eda837d444fe5b8081"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_openbsd_amd64.zip",
     "sha256_checksum": "ad3860459b920911d8b98f63b8e11fcc13b3baa3a60e8a84c045b0064f2ea0b2"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_solaris_amd64.zip",
     "sha256_checksum": "88535638a4edf559773862393c29557cb64301e2a938443adfe4c43ca7d93c99"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_windows_386.zip",
     "sha256_checksum": "1f95c2ee23cbf2acedb157ad5db668bde65720eab63b3439ff2383f5cf3dfdc3"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_windows_amd64.zip",
     "sha256_checksum": "a14213e119eaa884ab18d482de295dfc794257e57345c93a09c5a0686844960d"
    }
   ]
  },
  {
   "version": "1.3.0-alpha20220706",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_darwin_amd64.zip",
     "sha256_checksum": "6e3a1d8dce6854bd026a0fa669a78119be7bda3434996d8e8b9a1a496555c15"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_darwin_arm64.zip",
     "sha256_checksum": "6d3e88fee2ce4084b3a59955a82743c7a2b0144913015d1374939d6ece3e085c"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_freebsd_386.zip",
     "sha256_checksum": "833deb976bd906bddb59181ae95e58e20f32def2379b396e744fc0e626f617bc"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_freebsd_amd64.zip",
     "sha256_checksum": "c0ac5c01b1b6fb47e6100b9e2303680a2975e21868a39f07b052b02acc6b6eab"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_freebsd_arm.zip",
     "sha256_checksum": "4e980d5fba875e7354991a417122fa7a6e25e84056ae76fec6aa2c105003d4e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_linux_386.zip",
     "sha256_checksum": "e353f1c27cf5f49c131c84c9c0761dfb0f1cd5bbcd1811e7ac1c8627693fbacf"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_linux_amd64.zip",
     "sha256_checksum": "f3d0450399d3f78c171c36561eb4d27d37926b192a0a7401903ead171aef8229"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_linux_arm.zip",
     "sha256_checksum": "7e1b7bedd85bdb85921c7fb88f504d5e4efec8ce0a9c8b5668f17520a789a06c"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_linux_arm64.zip",
     "sha256_checksum": "f03071b048a3203210d1c796befa642dea8132e81ae7d035e13acde89018d7eb"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_openbsd_386.zip",
     "sha256_checksum": "bbd6a348ab11b5d3688f5717871a7922c2634175175db5ac70254e0487f5fce7"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_openbsd_amd64.zip",
     "sha256_checksum": "5719c0c2446bef12c785df7685ff92bbe968ef5898a72b6955f24f43163b820b"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_solaris_amd64.zip",
     "sha256_checksum": "a4b54cbdbeb8f38aa90598086a5df5e62b85427fbef72b1cc91639a3cafd2024"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_windows_386.zip",
     "sha256_checksum": "5671a73d95fd7072e3b821d9ef005b01cb1f29153e222c2196a6366b3d60488c"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.3.0-alpha20220706/terraform_1.3.0-alpha20220706_windows_amd64.zip",
     "sha256_checksum": "650dc5cb193a6bfc1da63ece68b68ed4656f688ab5e607a596ceb089c4251ede"
    }
   ]
  },
  {
   "version": "0.13.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_darwin_amd64.zip",
     "sha256_checksum": "fe5d1b6e22892c5dcc8b44d2a26ea1e29d90af6fcb1472f3881ca3c08c8a8084"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_freebsd_386.zip",
     "sha256_checksum": "6e32abdbd5d1b5ca9cabff39fed603f00b70eff68797ea76bcea2ea45d4dd314"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_freebsd_amd64.zip",
     "sha256_checksum": "fd556e3ff585c0cefd794783d5d5c05fba5ed793211991c5d98add0c6a4b67ab"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_freebsd_arm.zip",
     "sha256_checksum": "b015f1a9cdbbe1df060c62ddb6dffe79d12dda64921cbc95eb59d0ef6715d2e2"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_linux_386.zip",
     "sha256_checksum": "9e031e17121f524d8166d6f728ff3ccb64fed82e18d63f431c2fd5deb2707084"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_linux_amd64.zip",
     "sha256_checksum": "f7b842d1c06045b496fd00db83520e83a974a294b070dbaf88cb1013c5f02caf"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_linux_arm.zip",
     "sha256_checksum": "c15548aa1ec77e968af78c02847672b252145ef7aea1350de9cdbec55fd12924"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_openbsd_386.zip",
     "sha256_checksum": "e5db1c231ee0abd891d3b0ce737ccce8bd17e9eb2fb98b07d2acc50815ebd3bb"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_openbsd_amd64.zip",
     "sha256_checksum": "8ab1de5f9dcfa97e2b8608ca6ab8b357585bb4273ddc2ff44cb9b818905a6879"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_solaris_amd64.zip",
     "sha256_checksum": "d4adbc7427261bb7c052861a96369c19c48878f72e10c9900fc0e26544a4daa6"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_windows_386.zip",
     "sha256_checksum": "af1babe4ad7aa2e88fbf12e3ee176846db859617c572f3480de59f9284460457"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_windows_amd64.zip",
     "sha256_checksum": "fa7b17147baf4aca478ae83b25c14c1baa064a118a21e0e7b34e3d70e1732b6d"
    }
   ]
  },
  {
   "version": "0.12.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_darwin_amd64.zip",
     "sha256_checksum": "f0cc23bc6ec1a5adc4043108ff5c79c2bddcdc70b056bd207defca1ae386d477"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_freebsd_386.zip",
     "sha256_checksum": "a3a279e864ab02e73b3a276e4a04c4fdc16b84dbb0eeae640f15dae20a4faa12"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_freebsd_amd64.zip",
     "sha256_checksum": "5fb6b0c82a095ba4ee578a729df0169b24318815dd4cbc60e882f10fb8c558bd"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_freebsd_arm.zip",
     "sha256_checksum": "eb9a47ea1d992ddd205acc11256d25e9673967e0951d9b6345a310c9580e0894"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_linux_386.zip",
     "sha256_checksum": "10106f3dff1ed7fe78a766d25803d6352c59f87773276117436a251c11e10b4c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_linux_amd64.zip",
     "sha256_checksum": "d9a96b646420d7f0a80aa5d51bb7b2a125acead537ab13c635f76668de9b8e32"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_linux_arm.zip",
     "sha256_checksum": "2b2069512379d4a31c87159266ea76a951df12d446523cba3d9d5236d18cab6"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_openbsd_386.zip",
     "sha256_checksum": "f73af4a311fed8df99b8427f3cd0fe6321afdae3ecdfd22caa55c920e46641ad"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_openbsd_amd64.zip",
     "sha256_checksum": "95bf8b42974fe5e2c7face74fd8759233e4d982562159915a519b68e615ace8b"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_solaris_amd64.zip",
     "sha256_checksum": "5ef5c45482de7793c2b1a8b06e646cdfa13969c16053e1a8c143beb01ee2142"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_windows_386.zip",
     "sha256_checksum": "e0738cb2b7462569f1bea30c53d9cf4ca0a56c5b8060cccc5ce5cd04fb5f5776"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_windows_amd64.zip",
     "sha256_checksum": "ad4867345c404b21bdb39d8ee8041c2a7897b74590867bd98f12a29e4b3f0a52"
    }
   ]
  },
  {
   "version": "0.12.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_darwin_amd64.zip",
     "sha256_checksum": "f0e09af8ce413ec9a949c00ea6645cd8169a03412e545a3375adf91c3ad8c7ad"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_freebsd_386.zip",
     "sha256_checksum": "db9efdeaa6fa9c46caee6859f97e6213329acc384a58551a8260086b27ef9734"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_freebsd_amd64.zip",
     "sha256_checksum": "a3de6ff4e83ae23312469592d0b01d4e78cb0bb98cea66c67fd849c77f20312d"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_freebsd_arm.zip",
     "sha256_checksum": "f7666020c0d5442887c7cfedb7d8b2726dfa084211a7f31fe65ea8af73df6674"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_linux_386.zip",
     "sha256_checksum": "e00bcc5b157cf014ba160dd59d5db6e2c61939150c01bd116b7a9aaf6df0eb11"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_linux_amd64.zip",
     "sha256_checksum": "75e4323b8514074f8c2118ea382fc677c8b1d1730eda323ada222e0fac57f7db"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_linux_arm.zip",
     "sha256_checksum": "b67a5ba46a32c8e88d22a80861ddde64f7aec4403703c1172c34ad2dfa5c8c82"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_openbsd_386.zip",
     "sha256_checksum": "663983fb34a576596397bfbfef473cc452bd18ed132589cea0e3ed354a0db10c"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_openbsd_amd64.zip",
     "sha256_checksum": "24345acaed9c69ee6a6e98a48041304a0816c445124fc2c6184df17bb7e7c9e"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_solaris_amd64.zip",
     "sha256_checksum": "a9781a8c1d7900fafec392edcf972c6843eb2bf52b5521bf226de37353766102"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_windows_386.zip",
     "sha256_checksum": "299b574557e67ecf00c234d6d0b445a4c294ea4a1f71bcf153b8a128557335f5"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_windows_amd64.zip",
     "sha256_checksum": "8405fcf493bac25e377abd8402296a84e555e015db718d944baa8de83655311e"
    }
   ]
  },
  {
   "version": "0.11.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_darwin_amd64.zip",
     "sha256_checksum": "f5e04d3886e9a427490d1aa857a61b5a87d08dc26fb8637e3eaa72b30562c330"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_freebsd_386.zip",
     "sha256_checksum": "10a0917f86c589fc1088beccd1d2f5339b0062593833a188a0787daa84691e4d"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_freebsd_amd64.zip",
     "sha256_checksum": "ab5d2d3a0c3095c7d6017e2caab20d3149161e14db69e9de2d6901cd50869ea9"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_freebsd_arm.zip",
     "sha256_checksum": "4e6fbeab6cf77656ac1d7c216eeb0456db554b952d3dac24f0fd71489055aa71"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_linux_386.zip",
     "sha256_checksum": "9073990224ecd359ab9e511cd22a9feae203528e0b4d32b41856b46e0a490e0c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_linux_amd64.zip",
     "sha256_checksum": "4e3d5e4c6a267e31e9f95d4c1b00f5a7be5a319698f0370825b459cb786e2f35"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_linux_arm.zip",
     "sha256_checksum": "5fecc1f1f5b35ee3d46f6867a58c1452bfde09c27258cd8c7c73850a43b015f2"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_openbsd_386.zip",
     "sha256_checksum": "c54348b25ff03728914e62760e862dbda12a923d9da6e942a26637bee260cf0f"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_openbsd_amd64.zip",
     "sha256_checksum": "81c04d763e81d0e358c4716d4d7b0a7459ca5ecc08419cb5616f914f23b8774b"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_solaris_amd64.zip",
     "sha256_checksum": "288bafc239b5b52db61d827fbaf3ee727342f90a925cf9044f00fbab4b2cb56a"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_windows_386.zip",
     "sha256_checksum": "ff4895614441c30f116c751570d284b6a4f1c49fcfebbbb44592d1e85f36e5c3"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_windows_amd64.zip",
     "sha256_checksum": "b72152b537907c3297a17a6928def5c0acbcba8c0d92cbb5b12c24be59380ed3"
    }
   ]
  },
  {
   "version": "0.14.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_darwin_amd64.zip",
     "sha256_checksum": "363d0e0c5c4cb4e69f5f2c7f64f9bf01ab73af0801665d577441521a24313a07"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_freebsd_386.zip",
     "sha256_checksum": "5a3e0c7873faa048f59d563a2a98caf7f04045967cbb5ad6cf05f5991e20b8d1"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_freebsd_amd64.zip",
     "sha256_checksum": "4b7f2b878a9854652493b2c94ac586586f2ab53f93e3baa55fc2199ccd5a042d"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_freebsd_arm.zip",
     "sha256_checksum": "3c201a9a3e1d2776d0cfc0163e52484f3dbbbd73eb08d5bac491ca87a9aa3b7"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_linux_386.zip",
     "sha256_checksum": "b262998c85a7cad1c24b90f3d309d592bd349d411167a2939eb482dc2b99702d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_linux_amd64.zip",
     "sha256_checksum": "2899f47860b7752e31872e4d57b1c03c99de154f12f0fc84965e231bc50f312f"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_linux_arm.zip",
     "sha256_checksum": "a971a5f5da82ea896a2e91fd828c90ea9c28e3de575d03a7ce25a5840ed7ae2b"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_linux_arm64.zip",
     "sha256_checksum": "d3cab7d777eec230b67eb9723f3b271cd43e29c688439e4c67e3398cdaf6406b"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_openbsd_386.zip",
     "sha256_checksum": "67b153c8c754ca03e3f8954b201cf27ec31387c8d3c77d302d647417bc4a23f4"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_openbsd_amd64.zip",
     "sha256_checksum": "62fbc3f596490e33e6493a8e186ae50e7b6077ac2a842392991d918189187fc"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_solaris_amd64.zip",
     "sha256_checksum": "f66920ffedd7e81cd116d185ada479ba466f5514f8b20194cc180d3c6184e060"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_windows_386.zip",
     "sha256_checksum": "f8bf1fca0ef11a33955d225198d1211e15827d43488cc9174dcda14d1a7a1d19"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_windows_amd64.zip",
     "sha256_checksum": "5d25f9afc71fc49d5f3e8c7ccc3ccd83a840c56e7a015f55f321fc970a73050b"
    }
   ]
  },
  {
   "version": "0.14.0-alpha20200910",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_darwin_amd64.zip",
     "sha256_checksum": "2e65f929c74134f2a40ae1f092097c24159186e5ac58fbf19841a21b9f575893"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_freebsd_386.zip",
     "sha256_checksum": "a4856edba45e82498b3c7f55757aec1b4e392c96945accb03423de2062eb93be"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_freebsd_amd64.zip",
     "sha256_checksum": "d2d1b77373467f60fde5547c5a3d2425a2b26c92cb287f25e40de893ba837171"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_freebsd_arm.zip",
     "sha256_checksum": "a26233710219a7ef06797337ddcfdd7e1749fb9e4269fd7061a971fb415c789e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_linux_386.zip",
     "sha256_checksum": "425b95a85e5e6059addd5bff99394a04581f4fbe5c4514571cfde4ad4306fbfe"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_linux_amd64.zip",
     "sha256_checksum": "d18e0ccbe20a594a8c477526c1850c625b5c06dfd3d49c36c1ff35f1b83e777d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_linux_arm.zip",
     "sha256_checksum": "614cc913b0a77613dce0c8b50da503ddcbc6e126ad9087e1fa598e13c7fab0d5"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_openbsd_386.zip",
     "sha256_checksum": "b7e9d4a65943138b1142e97fa676db9cf5b18cd57ff4f7e7fc58f5d1fcc02737"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_openbsd_amd64.zip",
     "sha256_checksum": "21b607495d0ee3b56f65508787e415c0fb014a2e3e73dd7b43d87f733dc8b2ee"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_solaris_amd64.zip",
     "sha256_checksum": "1837334ca0ea367919bc20fa191dbadfa3352153fc1ecb82fb26d41257ad3f81"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_windows_386.zip",
     "sha256_checksum": "81f58157dad910f14c29789a5a238d1ef4f7107887b7059d407f1b60e0da2c7b"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_windows_amd64.zip",
     "sha256_checksum": "b80158c470f773414035cb26aa8887399ec6a20302c5f39712ad61359f0f05e1"
    }
   ]
  },
  {
   "version": "0.10.0-beta2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_darwin_amd64.zip",
     "sha256_checksum": "6138b4177e392e759bebc378cfe3a8dbbab6eae43214269464a005597aed85c6"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_freebsd_386.zip",
     "sha256_checksum": "bcd26a7664e90b32ff1683808f194d15f90b8533c2130361ed3bfb3b9d81aff4"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_freebsd_amd64.zip",
     "sha256_checksum": "31c0b44dba461fe20c50f9f4dba061ba7d48c86e26415bd3119158c5be15ce32"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_freebsd_arm.zip",
     "sha256_checksum": "f0bdf03936b06f7289c3f80eb802d54805c52d78df284ce9d57a31272d344238"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_linux_386.zip",
     "sha256_checksum": "92b7200dc41e3ae5045fb8ab0095307fad47b7a45144a5f05c99160d349a7cff"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_linux_amd64.zip",
     "sha256_checksum": "31eb44cd3ffe12ec386668a2702c7040a69c63b3bd833f046758c1a91b3a79df"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_linux_arm.zip",
     "sha256_checksum": "ac9bad13d3a0bb36d1d6962c206475c7003a61067316dc37f1266fca0706143b"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_openbsd_386.zip",
     "sha256_checksum": "dd525b2a1ecc8a2f5b46883e89b1467fd49e055d77b9de4a7a40e68c4183b90"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_openbsd_amd64.zip",
     "sha256_checksum": "ae9d9b8ef0b95883ec72247d7b6ede6dc63cd07c3f84ec9179bed3c3faa121f0"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_solaris_amd64.zip",
     "sha256_checksum": "be30a959cc979aff8f8b0e503118db1c9df87162ae20c6408cb2c2c137889c4"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_windows_386.zip",
     "sha256_checksum": "ab112b6f28d8864030d16865583aeaa49e717140e4b303e745338f11fcf9557b"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_windows_amd64.zip",
     "sha256_checksum": "a5f881b4dec585f539ca39200d92a0ffb2b56a063c5b16cc6f2da70b14fdaa19"
    }
   ]
  },
  {
   "version": "0.13.0-beta2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_darwin_amd64.zip",
     "sha256_checksum": "4729fa267c5be7f1d0c19a85d36e2b4577f303866fc25403650acb4243c2021f"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_freebsd_386.zip",
     "sha256_checksum": "98d5e28665f1e61840fafbb220a63ab7fd3c3e37412c730ecfab3743c99d7749"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_freebsd_amd64.zip",
     "sha256_checksum": "3b5b8f40ec5845689fb5d63f23442be413c749205c5943db2dbd14cc3180199b"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_freebsd_arm.zip",
     "sha256_checksum": "fe6aa78a3ee209774bfb5c7dc8293ad206dee45718594e9a30c59a985e54f799"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_linux_386.zip",
     "sha256_checksum": "a03c4114d2a77c5d47909e9c4d238e45129a8d4876adf0c244708c8232d9b1f3"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_linux_amd64.zip",
     "sha256_checksum": "ae0c4cb78b8e406f26b77863bc09a47dfb9b7e7fbcf4cd672d737eb7ea6a52b6"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_linux_arm.zip",
     "sha256_checksum": "84ef5fa08520a9c6e0a1ab8901cd5faf5b908df4a6a18892d58613570869e8ae"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_openbsd_386.zip",
     "sha256_checksum": "eec5ba434e92d2649c32f041cf0650f848fda49c52cffc89644503a51f30b8a1"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_openbsd_amd64.zip",
     "sha256_checksum": "926b1ddf70373a9b9b7b446b66a92029160631fb6d76d88ab6ed42a460e34569"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_solaris_amd64.zip",
     "sha256_checksum": "7144a41e0b978eaec7d6c475353861e958b5bcd07981c40e309ffda9c24db3aa"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_windows_386.zip",
     "sha256_checksum": "98e5f3d9a260efc59187135bd455ad0bcbbb968e67b637a144dd41815ed2d4cb"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_windows_amd64.zip",
     "sha256_checksum": "aa8c67d95605f658e30c3eee0f4a906ed9e70679bee6f12b8ded2d38188456c1"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20211020",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_darwin_amd64.zip",
     "sha256_checksum": "16686005a31b8404f5baedb260ebb11186868ce6c42f345b45e19b83e170f9c1"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_darwin_arm64.zip",
     "sha256_checksum": "282d76d3c33b2e9bf94212a2485fe5a31fc1a4e6d6339ac6cfa923db001edb7b"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_freebsd_386.zip",
     "sha256_checksum": "d9c8c763c4d6634fae8729e1dc1990ae1b40b25542acf0c0bd1b3039411d6b3b"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_freebsd_amd64.zip",
     "sha256_checksum": "33e6cb65e5fb1c4d6b4226dab70e1db8979366a8dd20eb515b675012837dcc08"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_freebsd_arm.zip",
     "sha256_checksum": "5e63cf4f360efe285c2718dd2d59bb12e1626ab6cd747a45acd8fde30202f4b9"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_linux_386.zip",
     "sha256_checksum": "5601d720c44586f9e46da344d9e7ba92fafb4722fee08611548129a725f26bf2"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_linux_amd64.zip",
     "sha256_checksum": "c8769b93162cce9439a0f0487f085402aebf316401c47916b8879bccb7f2ce05"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_linux_arm.zip",
     "sha256_checksum": "503d1d8c8776cd05ce24dcba65e9cfe509cef14142776210fef8186b265f6245"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_linux_arm64.zip",
     "sha256_checksum": "54486f3cedc30c5c1d64f5f8570d96e8691bfc96bd9bae8a5de873d9fbe7e067"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_openbsd_386.zip",
     "sha256_checksum": "46831a8ebb0f177ee3c8c5767494000f4bd90cfe67beaed0ae7fdd6124df7f0c"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_openbsd_amd64.zip",
     "sha256_checksum": "93482219eebca754c5165e099c37f3b6a2f126dd9c23f7247a3e4572fb9f7d41"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_solaris_amd64.zip",
     "sha256_checksum": "3322b8ad47c03effcc6cafa51908ac8b5bb7f65c761c4cb49b79f976d9c62d56"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_windows_386.zip",
     "sha256_checksum": "d18f37a949b0b56142989afc33adfc11a6f52567870305552c2114959c5e8778"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20211020/terraform_1.1.0-alpha20211020_windows_amd64.zip",
     "sha256_checksum": "681fd2dcb6dc328b872b28bbcdb5d8c5bb4cb6d97789a27098afae907cb35758"
    }
   ]
  },
  {
   "version": "0.15.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_darwin_amd64.zip",
     "sha256_checksum": "d13b507e6f51fc58d880313775262954369fa6c98e163fa71e21b7d2a613c32a"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_freebsd_386.zip",
     "sha256_checksum": "e3cf339577aadd2fe5c6842e9560dd7346a04f74ff8f8c34e6a581739a9160b2"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_freebsd_amd64.zip",
     "sha256_checksum": "1b2586b115360d8bd9f70a99a8cb39383185564c4fadc097468e1645f35e12dd"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_freebsd_arm.zip",
     "sha256_checksum": "f39c09fe23d409dacae71c051ef5a36bcb4819e255e33f429ddbf45aece8de71"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_linux_386.zip",
     "sha256_checksum": "7f5dda9e9247447194091731a1c1fffb4f3b727fcc26e8970d80f78ae8d2af1a"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_linux_amd64.zip",
     "sha256_checksum": "6d4780cbbe49f2d49eb49e8d7a90f5b8d5beec7e64d80d3f8eb0c6e19156a26b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_linux_arm.zip",
     "sha256_checksum": "8ab0c49cba6e443acae61239a4c371698cb45856094f69f8e9f21d3c0daa4d28"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_linux_arm64.zip",
     "sha256_checksum": "646e3ae3db87bde44d6076c01e5e9911fb8d9c7369ff9ac2260aed261ebb4194"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_openbsd_386.zip",
     "sha256_checksum": "7b5b5062ddab1ba43f9ea85a86457a94cae421ee3af5584da68170c6dce27983"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_openbsd_amd64.zip",
     "sha256_checksum": "8b214b02639d55d647a174e31e6126f06ca6640dcdee76c135243dfe7b348605"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_solaris_amd64.zip",
     "sha256_checksum": "46f22a85839c270f275dd8fec495d2b2616710bd42e3b4ba2e75c47b5096cacc"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_windows_386.zip",
     "sha256_checksum": "cffe9b549cf23c99c3b4f63cb955023799dca5a3d39d379182355328a1035885"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_windows_amd64.zip",
     "sha256_checksum": "40e19e371e4c2b54bea3c5984aa7fc410f33099d28d31f4d688607cffd8e6125"
    }
   ]
  },
  {
   "version": "0.12.19",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_darwin_amd64.zip",
     "sha256_checksum": "5238fe45d051cac90f0fc0701796c5244ef88218d0fe4eceec31cee43899a434"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_freebsd_386.zip",
     "sha256_checksum": "8020780ad13eb5b2a1a16cdc1f223263716126b532ecba9088b2deece71a84c1"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_freebsd_amd64.zip",
     "sha256_checksum": "3850ba419a787e7a99ad4f731a14173976bbe8b64aa1c7ba6bd5a5fd44b65acd"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_freebsd_arm.zip",
     "sha256_checksum": "5d1b0982f6cffca4148d01ec2bf877105679e7329a6091af5b2d6848963a8b3a"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_linux_386.zip",
     "sha256_checksum": "3cbaad33ca5096547d7270da978201d7a2baba6a54ebcc30c34a09b745528a8"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_linux_amd64.zip",
     "sha256_checksum": "a549486112f5350075fb540cfd873deb970a9baf8a028a86ee7b4472fc91e167"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_linux_arm.zip",
     "sha256_checksum": "5351aad393735df1bb6e85db385cc4f106c1e8913eab025f5d777e05f1e09e8"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_openbsd_386.zip",
     "sha256_checksum": "a19d35d2e9335c345c3e8dff0ca25b4310da43101bb3483017c1d39495b1d9a9"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_openbsd_amd64.zip",
     "sha256_checksum": "7ec79bbf140ba9a8aa543301ab27a6f2f44506b3f2107bcbc56239b8eed075fd"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_solaris_amd64.zip",
     "sha256_checksum": "56075ea59e6d62dab0f8291e6cd5aee9aa1a49cc9802e624e360bc2f2a9b4761"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_windows_386.zip",
     "sha256_checksum": "42c82b37910db3efe73c5b3ac2d983f06d29f708ce5633fcbb51521f2848c93a"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_windows_amd64.zip",
     "sha256_checksum": "9599da2ba4c649492982507f2a6e4f650c5294261c11038d388261a85bc4a9d"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210714",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_darwin_amd64.zip",
     "sha256_checksum": "e35af55d525ab8c18ed880b2875caf280fb720b3d845f23021e7c0efb07048b8"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_darwin_arm64.zip",
     "sha256_checksum": "63524ccb1ceca43180205ba15e4d5af6831ddfb1797e93d7a0e583f56962968c"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_freebsd_386.zip",
     "sha256_checksum": "fbdd1824e8868895ed607c234589d3570e49edc533da23b39532dad78df4ca68"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_freebsd_amd64.zip",
     "sha256_checksum": "7bb9e44521a241e0d168dd6e78c30f8afa0185a61138363d5310338d554b03ae"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_freebsd_arm.zip",
     "sha256_checksum": "3618f35eee99afa865f6bcd57ceccd28b17869dd38d6da5b2253f12763d3bcdd"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_linux_386.zip",
     "sha256_checksum": "ec2cb82eba21d47d43e3507390c8bb419343a3860f66c38cc73ec69c96ed3022"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_linux_amd64.zip",
     "sha256_checksum": "707310efcf16f4b2cbf025b90c246c5176b64b67054663a5479fe1c7d7e0c193"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_linux_arm.zip",
     "sha256_checksum": "d7c60f648b87e61b724dce654afa341d1a1642ce9879ec5f81a3a0c738c70a1d"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_linux_arm64.zip",
     "sha256_checksum": "ae0d9cffeef2f4ff36f0b22a2d5f11f7110b076f8eb59f1e8e8288ad1051a786"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_openbsd_386.zip",
     "sha256_checksum": "805f3b411ca5e89be23d6d9f93a59d5777a7cd70c62c0e615a8c72d2b2396f15"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_openbsd_amd64.zip",
     "sha256_checksum": "f8ba798895d4a48756f19ddefd6f4e530124b26fca166722a074525a6e730ce9"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_solaris_amd64.zip",
     "sha256_checksum": "3dc7e72641da262702ce828f2f507cc0890436ba6e6041be3ada91ab02f16f49"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_windows_386.zip",
     "sha256_checksum": "9149c5089112244a5e75547be9b3bb56b88c2befa96a0aedaad7d39062584926"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_windows_amd64.zip",
     "sha256_checksum": "5dc8a31128b82b03ec7c4ed840cf30355d986705c535dc5467a2c579b3d0eda7"
    }
   ]
  },
  {
   "version": "0.7.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_darwin_386.zip",
     "sha256_checksum": "9f576232b582deed7b20f9a9183b8ba6573badc99db1caf609a80bc9ba9ff914"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_darwin_amd64.zip",
     "sha256_checksum": "21c8ecc161628ecab88f45eba6b5ca1fbf3eb897e8bc951b0fbac4c0ad77fb04"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_freebsd_386.zip",
     "sha256_checksum": "9d4ee0bdc43ee1b29c7973746c89146df28cc81a61edffc88565d82c189079e8"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_freebsd_amd64.zip",
     "sha256_checksum": "362e6fc15a54f1a80eeea21ed1fd0a6da090eeccbcffae1027349c20b224af61"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_freebsd_arm.zip",
     "sha256_checksum": "723da17c7408bd07f87a7cfffa10246d3f56374f4811c752a5f2483b0c4c196e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_linux_386.zip",
     "sha256_checksum": "7ff291402b2fab895ab1d7a75965a0c0bc8f7b1bf529f904852b6c8e5f3f6604"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_linux_amd64.zip",
     "sha256_checksum": "8950ab77430d0ec04dc315f0d2d0433421221357b112d44aa33ed53cbf5838f6"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_linux_arm.zip",
     "sha256_checksum": "ddfe78fb3fab9491429a545bd114353eb7b3454fada17f39b97a08837e6112ca"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_openbsd_386.zip",
     "sha256_checksum": "34a5a318f6dd87d734852e4e906fdd2c05aedd3d554b885f77413925c52a3f88"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_openbsd_amd64.zip",
     "sha256_checksum": "6a97da3404ad8a285bca4c4c20c855e6a03a4624cd50632bcdd83707eaf7363a"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_solaris_amd64.zip",
     "sha256_checksum": "780324925a15051b68102d011422e8228283d5ba749abfa95be46e9643a5d2aa"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_windows_386.zip",
     "sha256_checksum": "ca9433509cb5f8febc19fcbd23be820511b563ebe3f17ade5946660388d42582"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_windows_amd64.zip",
     "sha256_checksum": "b50acc1269909804075dd12df436f0029c54ffbed2939537d1fc0db395c703c8"
    }
   ]
  },
  {
   "version": "0.13.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_darwin_amd64.zip",
     "sha256_checksum": "dfdc8ef005df19d7ec0fcb5f151e51b144233ca425c39dabf94c037e80780b05"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_freebsd_386.zip",
     "sha256_checksum": "655eb3e321e644b5d3911e86a764cd954fa04746e72a04f442a0d27b5cb59475"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_freebsd_amd64.zip",
     "sha256_checksum": "cc4cf629c224057215903d8c1ae09d967ddd47b41fded99c1d3df989fba53931"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_freebsd_arm.zip",
     "sha256_checksum": "1a6285a9debd0b9ca3b3cfee4b7ba91a5d554c38dc9673c48c4c71b619783126"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_linux_386.zip",
     "sha256_checksum": "bd8aa7a89139223d915a0031616e732fe921caed3c2a790f635afc5d37a1802a"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_linux_amd64.zip",
     "sha256_checksum": "cd867252689979ca33517b9e391fe39fceecf0b6df91450a6abb75833cbd2c7f"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_linux_arm.zip",
     "sha256_checksum": "c38ffb6495b409bbe5ec612bd9c1e6a76c18c1f4391aab2917456a3c2ebbbf9a"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_openbsd_386.zip",
     "sha256_checksum": "3668de1eff0ba8a93faae36b720944fa4f8b53dd0d14586b85836bfbab2447e4"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_openbsd_amd64.zip",
     "sha256_checksum": "7343494dbf7a8fa7b03c089110666d729a13d7abe594aeb9d1e30af9bb497e39"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_solaris_amd64.zip",
     "sha256_checksum": "41af6fdb6a620b317f097480f870cf9c93363c9d24e8833dd3ed8ba22864157e"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_windows_386.zip",
     "sha256_checksum": "368665190a8209c23a32bfce60975aef6bf5326e7e121f4a37d9e1f92f8a32f6"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_windows_amd64.zip",
     "sha256_checksum": "3a3154a244760b4020ecec92b91a4bed3e972d030671895980c9c5983ab6b73"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210728",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_darwin_amd64.zip",
     "sha256_checksum": "16f1d4e53b29dacc61184c8e52217c6f8c1197b3803a698cdc2ed1848a1b20f6"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_darwin_arm64.zip",
     "sha256_checksum": "3de8ab1cff4e8db772ba4177e3cf9b25957f1aea19517f7a70718c18f0ccc7f9"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_freebsd_386.zip",
     "sha256_checksum": "fdf67b94caac9eb22adc7b6424092d3b9e49829ce1b6e2d754ba000fbec91db9"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_freebsd_amd64.zip",
     "sha256_checksum": "15819b5e95f05f91adff628f505c59a05bbc6aa8ea12e3364ff09598aa7386"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_freebsd_arm.zip",
     "sha256_checksum": "3f20d1740215640a770de2f8969da060520571cd7a398ceaa9a127fa56a98482"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_linux_386.zip",
     "sha256_checksum": "291fcf1782c041580844271675ee577307a599c7527ff61f12d251fb59b5ac95"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_linux_amd64.zip",
     "sha256_checksum": "89686fea25dd9b4f6acf6d3381e089b59352f0c17873b9c63cf7f56a9b7c6c19"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_linux_arm.zip",
     "sha256_checksum": "c6e1a3a1dfcb705f8fc1998129f13d1a721922c6b229fbaf04e20a573fc0d199"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_linux_arm64.zip",
     "sha256_checksum": "fe08fdeecd6307e26b6a343b21a6011556fd94a77affb7a670a9ee641dcf7c04"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_openbsd_386.zip",
     "sha256_checksum": "ede31c7441de5979012eb42260169f8ba9df2c1c97c8cd4e66ce8f3f718e0c49"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_openbsd_amd64.zip",
     "sha256_checksum": "1b13950f52f1cef9a970c64c2957f41ef1e61628ccb29e64869f87390c0aace7"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_solaris_amd64.zip",
     "sha256_checksum": "74eac96ebb7e0696afc0dc712b2a8e975694d190f3b1465b993a54d339e9167a"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_windows_386.zip",
     "sha256_checksum": "9dbfa83da11b49896e8c059c9587093dbf160f97ea6c94ba8346f571c3db7a53"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_windows_amd64.zip",
     "sha256_checksum": "45ee611e13072cf3420e7cfc3d4874c55e31e0e278850f2bce533f45d9c8e9c5"
    }
   ]
  },
  {
   "version": "0.6.16",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_darwin_386.zip",
     "sha256_checksum": "c50270c498807c4fab828fbdae1a02720ba123cc2a8acbe6fdb783ad7e168afd"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_darwin_amd64.zip",
     "sha256_checksum": "23feb79263126877e6128a03c600cd626f6691a118a474694c5ad45cc5da9366"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_freebsd_386.zip",
     "sha256_checksum": "4dd5e25f055a09ce11b473197053e2b1181a4a93982bd2260e929e89ea9eb842"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_freebsd_amd64.zip",
     "sha256_checksum": "8afedcbb3d045048b4f788ffcc210bf69c5bf5214bfaca8ea60f127f5e0a95a4"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_freebsd_arm.zip",
     "sha256_checksum": "6d8738a71ad3383847f3bf2cedae94280ce91944bd4cccb19f4fa14d0e917546"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_linux_386.zip",
     "sha256_checksum": "7ed928cafcf1164ea802912eaf8a40dbd382676b3d9e2373dd1323a5e4a8a11c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_linux_amd64.zip",
     "sha256_checksum": "e10987bca7ec15301bc2fd152795d51cfc9fdbe6c70c9708e6e2ed81eaa1f082"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_linux_arm.zip",
     "sha256_checksum": "ca31a61ff9c8c792699b2524a3c58b60fd0801e6155275885af775b426ab595a"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_openbsd_386.zip",
     "sha256_checksum": "1ecf2db7cc5a807f97d4215e62daf694028e90bd849f29a6a77c114219b5377c"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_openbsd_amd64.zip",
     "sha256_checksum": "5e40ee5089ae7527f1fbb82f34edbbf90bfd8cf22fa8a83a8d3c2933aa486f9"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_solaris_amd64.zip",
     "sha256_checksum": "370057977b5589fb86fda42498861a44a986d8f23d8eb828613d498b6fb10c69"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_windows_386.zip",
     "sha256_checksum": "4032ee9f298005ee525aeef5c8763787f6867ab377367597083a39bb33a06ad5"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_windows_amd64.zip",
     "sha256_checksum": "6acbc7e7025104d265cbfee982e5e7af0427255fa70bafa5bcbc595b56c6cce7"
    }
   ]
  },
  {
   "version": "0.11.15",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_darwin_amd64.zip",
     "sha256_checksum": "b0d2c9f9068be007f9b8eff211a679e1f07368b640245871bb02a6c2cdf28c07"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_freebsd_386.zip",
     "sha256_checksum": "2a9ec79271961e8babf350eada350cb056181118699a916873b544b21d92304e"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_freebsd_amd64.zip",
     "sha256_checksum": "14a7c4c6d467469bfbbf6f2e3183d8952331d928fd825a5e120c1c4fc2db5b7b"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_freebsd_arm.zip",
     "sha256_checksum": "5a5f9e8390d99c50879f6888470869afd46fd7596b5a147532103011498ebcf2"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_linux_386.zip",
     "sha256_checksum": "83d55cb7979dac785e2fe409ab9142fd7c4f1c934090776c77c53f3683d404a3"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_linux_amd64.zip",
     "sha256_checksum": "e6c8c884de6c353cf98252c5e11faf972d4b30b5d070ab5ff70eaf92660a5aac"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_linux_arm.zip",
     "sha256_checksum": "45941c547e3c8f9822173fc2a4b20836db5693e91cbd4da195f414295f05ec53"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_linux_arm64.zip",
     "sha256_checksum": "99be59d7baf6744b496b77b92fe35fbfcd0395a1c2be328a87328406bc81d708"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_openbsd_386.zip",
     "sha256_checksum": "cc8dbe1496de13c965f5a72497990c814d67eca20fd6b25b0a34d908a7f46375"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_openbsd_amd64.zip",
     "sha256_checksum": "409ef675dfba7c10a09f62d006b720afef0ee085dd51f25189adbae02c66f96c"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_solaris_amd64.zip",
     "sha256_checksum": "8c22ff5fdc36dc1d49e35316af85dd824713f1d9680e33605c8bc19eef57c77f"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_windows_386.zip",
     "sha256_checksum": "1e064e78d090c739fe1d5314e9aef53e10153cc739aa54eaf2748f353c60478c"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_windows_amd64.zip",
     "sha256_checksum": "a5f15fb56558f2055d40da93c9bda8f90e4763922a0e9b7dd7d3f432e0222056"
    }
   ]
  },
  {
   "version": "1.1.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_darwin_amd64.zip",
     "sha256_checksum": "bbfc916117e45788661c066ec39a0727f64c7557bf6ce9f486bbd97c16841975"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_darwin_arm64.zip",
     "sha256_checksum": "dddb11195fc413653b98e7a830ec7314f297e6c22575fc878f4ee2287a25b4f5"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_freebsd_386.zip",
     "sha256_checksum": "4c09759e7cf88e016d03ba6c651d11e63b63152b3fff2abd5eb721402ce47cc4"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_freebsd_amd64.zip",
     "sha256_checksum": "d18ed70f2c52218c1a53cc169aaa40a9091893f6f894f2f75da98b2df78b8434"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_freebsd_arm.zip",
     "sha256_checksum": "70bbe087b6eea763aac7581b69cead24827874af718ee00b90d9def0dbe3de60"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_linux_386.zip",
     "sha256_checksum": "e5a8bcbe048c4c3f510106d434208ceab22c627b314007145a2d0ec1c3086825"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_linux_amd64.zip",
     "sha256_checksum": "3e330ce4c8c0434cdd79fe04ed6f6e28e72db44c47ae50d01c342c8a2b05d331"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_linux_arm.zip",
     "sha256_checksum": "2b977ed400747f03374262f171d3deba710f303a0be2daafe9ae2fd65ed851f3"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_linux_arm64.zip",
     "sha256_checksum": "a53fb63625af3572f7252b9fb61d787ab153132a8984b12f4bb84b8ee408ec53"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_openbsd_386.zip",
     "sha256_checksum": "2cb9414ad9131abe0889a8a9c9bfe678e19dc300f6eb662f7e20aee91a1f24ca"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_openbsd_amd64.zip",
     "sha256_checksum": "cb41a51d99562871cdac7160e0d441d6c3e8104a9a487be6e1f1cb7c967c10d0"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_solaris_amd64.zip",
     "sha256_checksum": "d34e265c7d8df94e41dc36ef0aaa8ab85b5b4292bb6a114f7cc9eaad87371dd0"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_windows_386.zip",
     "sha256_checksum": "9d240d2eb60ca403f6130b5252e697056d1f1ada46e0cc09b80acf52bd1fbdac"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.6/terraform_1.1.6_windows_amd64.zip",
     "sha256_checksum": "9ff13cab10ba1441e3e587758f01ca6054ddcee6920770f16790e261a1d6aa16"
    }
   ]
  },
  {
   "version": "1.0.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_darwin_amd64.zip",
     "sha256_checksum": "fb791c3efa323c5f0c2c36d14b9230deb1dc37f096a8159e718e8a9efa49a879"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_darwin_arm64.zip",
     "sha256_checksum": "aa5cc13903be35236a60d116f593e519534bcabbb2cf91b69cae19307a17b3c0"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_freebsd_386.zip",
     "sha256_checksum": "aaa7791eb7c100281d8848a1a6674fc7d81b98107ff9e5b5ebb76a0c983fe477"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_freebsd_amd64.zip",
     "sha256_checksum": "195752d63d9219c3b9f2feaf9baba0f8f5cddc94f3d0c36f3ff98aff5a7c4bb"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_freebsd_arm.zip",
     "sha256_checksum": "df72cd7b6eb4918532f0b431132d2f761205ae54dacce77fa968cd9f74a60ac8"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_linux_386.zip",
     "sha256_checksum": "e3eaf5eb688377e7b5eae2ec3734058ac35f7e52a934c66b75676b8355e1d27d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_linux_amd64.zip",
     "sha256_checksum": "f06ac64c6a14ed6a923d255788e4a5daefa2b50e35f32d7a3b5a2f9a5a91e255"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_linux_arm.zip",
     "sha256_checksum": "208772e1955c5b0b41bd8c28010c170c77bf09c457840f87e930cdc8116cd149"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_linux_arm64.zip",
     "sha256_checksum": "457ac590301126e7b151ea08c5b9586a882c60039a0605fb1e44b8d23d2624fd"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_openbsd_386.zip",
     "sha256_checksum": "6dd8f6502ae8f7c07f4fbdbe4aed1838704563915d53457460feee2bc40860b5"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_openbsd_amd64.zip",
     "sha256_checksum": "ad82264758d2e1ca5a31fcfa5c44ff5dd3f2c08bf245aea35e739976ac8ef9fa"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_solaris_amd64.zip",
     "sha256_checksum": "705ecbb40196e572b2b18729a7c34cfd865f7b5646951fc9018b3e4cb357b23"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_windows_386.zip",
     "sha256_checksum": "83590b1a77ad59c0e5aea20e27eef7d35e61d951842de10f3a5fce52c1470fdb"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.9/terraform_1.0.9_windows_amd64.zip",
     "sha256_checksum": "5518b138557ede4c0e98ca2635222096879e9753ea4cc9bbb0306299468b8cda"
    }
   ]
  },
  {
   "version": "0.11.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_darwin_amd64.zip",
     "sha256_checksum": "d5f7ffcfd34fe58ed25fe48650f1c9ac1d9e15983af43deaeffc6d0a88ba346"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_freebsd_386.zip",
     "sha256_checksum": "94f8d88149f577c45ba763f1f975e417b5476ac8946b6ecc62535a9eaf4667ae"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_freebsd_amd64.zip",
     "sha256_checksum": "fb868eaf385e683a4fb6dd0c9a3a377df4607df1b684f2d37d45062ce91308de"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_freebsd_arm.zip",
     "sha256_checksum": "6ffc07d24f098b93689ea14e18ee07625ef2ab7ceff422f14996c723ba3dc84f"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_linux_386.zip",
     "sha256_checksum": "c44a1fd6d13148db2a51cf682cf7362f41f9cd1c524f0b8511922556077e3f6f"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_linux_amd64.zip",
     "sha256_checksum": "402b4333792967986383670134bb52a8948115f83ab6bda35f57fa2c3c9e9279"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_linux_arm.zip",
     "sha256_checksum": "26fbb7318aea085f0400b987980b40f1a5b11a1c79206eb96d40e07d966d4ea7"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_openbsd_386.zip",
     "sha256_checksum": "533309cfae42625e1ed87f06ca31e9dd76c869652240414d45bb826aff19f211"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_openbsd_amd64.zip",
     "sha256_checksum": "a199aa3830910fbdd0ef10eafd782eec055daaad94d100946be808535c5d2071"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_solaris_amd64.zip",
     "sha256_checksum": "6707ebc0b1b55167dffe1925aab702802d686b6f3a9aec6ae7b6bff143a00201"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_windows_386.zip",
     "sha256_checksum": "3081a94822c7d088575c75dc0c4be2543301c4834d8da731bd53171e082a6d86"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_windows_amd64.zip",
     "sha256_checksum": "a958adf0c4177ecccfdc34df626e752027b52278f04fd527244f46ef220a54d5"
    }
   ]
  },
  {
   "version": "0.8.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_darwin_amd64.zip",
     "sha256_checksum": "275104513600bf50a28942131d928d2be405c75f9f36a9c722718500075856a1"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_freebsd_386.zip",
     "sha256_checksum": "1509419bd9aa6554b2810edec093b9d8b1d04574a41ccc5c4ba38687581279a"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_freebsd_amd64.zip",
     "sha256_checksum": "9a042c62605171e29f72a26f310afcf712776b90232f0c8cbcc2547bfee3dcae"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_freebsd_arm.zip",
     "sha256_checksum": "51c8c47622f65a8ab674d364bca173ddb76ec53ba57a930a8508c6ab4e97921c"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_linux_386.zip",
     "sha256_checksum": "beafc708c1f2ec65447206e75abd0ef3367981d6293a6e412c0db41fab66303d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_linux_amd64.zip",
     "sha256_checksum": "da98894a79b7e97ddcb2a1fed7700d3f53c3660f294fb709e1d52c9baaee5c59"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_linux_arm.zip",
     "sha256_checksum": "df984febce2cbf14d1fec358e9f6c1ce76436e29b86ac2e2fe5fb29731e26280"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_openbsd_386.zip",
     "sha256_checksum": "b8a187357a3575c08e2c512c6675ce0c6d39928307991b8e0df856e2868df50f"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_openbsd_amd64.zip",
     "sha256_checksum": "64be3aa8dd6040d134b71e82b47499ce6e87e3df1a83e0a47aba37edfe13f0b8"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_solaris_amd64.zip",
     "sha256_checksum": "9c6555b73343239413df85eb61d36fb04d019b5fbfe1ef6a5639d56a3acb8b77"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_windows_386.zip",
     "sha256_checksum": "fa5c5fd00b8fe55034344c4d7f633a62a591d83a524f33f551ca5547497b5e93"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_windows_amd64.zip",
     "sha256_checksum": "61a7b5cfd3587665ba99b855800c3e7d662b917be68b24dee1f11a59653ab1a2"
    }
   ]
  },
  {
   "version": "0.7.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_darwin_386.zip",
     "sha256_checksum": "febadcfafde9656c06cf0d795a8027b4ce97c3bf797663cfbcf8d6ff0387c25a"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_darwin_amd64.zip",
     "sha256_checksum": "9daaec788ee0540d7b3a92f2dcf86656f3c567e2c267c64c03aa712901796470"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_freebsd_386.zip",
     "sha256_checksum": "3147e9d929d136369a52e6950f78e63f52e85494fd62d09965225481c0afbfb7"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_freebsd_amd64.zip",
     "sha256_checksum": "77349280b2ae79f849c89df2d8b1e30213d1c79c0038ecf83aa81fa67b85695d"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_freebsd_arm.zip",
     "sha256_checksum": "73da85740e5be0d7cc6bc5b6fc6adb395a3308049abf8911f53516cdd97abbb1"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_linux_386.zip",
     "sha256_checksum": "25b6eb661af4023cb12f6087e1fed4e117d2219b2dc08b823c1f78574d5f66af"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_linux_amd64.zip",
     "sha256_checksum": "b3394910c6a1069882f39ad590eead0414d34d5bd73d4d47fa44e66f53454b5a"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_linux_arm.zip",
     "sha256_checksum": "cc74511b5c9c61fb0cbf4bfe19a51b9f83fefad5b0727923a2ed75f8d63fccca"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_openbsd_386.zip",
     "sha256_checksum": "8dca8f7f4a8f09843e35b9fdeabc01ef83a6e070874dff578eb85990f0268154"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_openbsd_amd64.zip",
     "sha256_checksum": "68efa6eb8a9eebb0dda6abb3d902dd94189a53542edcdc62269567cc8ccfa4dc"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_solaris_amd64.zip",
     "sha256_checksum": "aaf005f1826a895c5baf037ae60d3b6132a4f74b06c921db796c2862e3653670"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_windows_386.zip",
     "sha256_checksum": "517455fc10f30441bb149e9dd3af436526dc2a7abd81d07d8cb6a67f457347ce"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_windows_amd64.zip",
     "sha256_checksum": "f45521b9e077f080c1078117efb668e29ff2b15c20ec71cbc3c5312ef6cf6a1c"
    }
   ]
  },
  {
   "version": "0.15.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_darwin_amd64.zip",
     "sha256_checksum": "2cfa2f896aaf2c2b2fdadea6881f2796fe0d85ad0a42f471aadfb113bc32d11b"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_freebsd_386.zip",
     "sha256_checksum": "a737c07bb56425928332a61fadcf08df22c2d611b341e70613afcaaec16d2141"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_freebsd_amd64.zip",
     "sha256_checksum": "6266df57cceb81e8bfc894a7431ecf1e3b7409af8850c8df99363cd6ad10928e"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_freebsd_arm.zip",
     "sha256_checksum": "142b5653aeb2f36a2aa4ad4a8e26ddd31f4618d1318a15d590836adee0b2919e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_linux_386.zip",
     "sha256_checksum": "b1a1428fb3e4c30a09c5b5833d3f271e4d71cf1beb35c79dbb2bc4f487a0108b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_linux_amd64.zip",
     "sha256_checksum": "5ce5834fd74e3368ad7bdaac847f973e66e61acae469ee86b88da4c6d9f933d4"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_linux_arm.zip",
     "sha256_checksum": "d022162bc1c0f7e73dbbf810c4591b72234ee409626b2f04108686d83d9f1bc9"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_linux_arm64.zip",
     "sha256_checksum": "71df9fdeedd823465dae434f479e29a745a792b2315030925b39838f519f8bf3"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_openbsd_386.zip",
     "sha256_checksum": "2e1fc4d68ac65eddf03d0665ea3a670c5bd45ee369c061b24175aaaf98589962"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_openbsd_amd64.zip",
     "sha256_checksum": "b419e89f4cb110fd7982144a620356e1ffeed954506a1954f431199661b08b40"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_solaris_amd64.zip",
     "sha256_checksum": "35d934b53fe4d5192cd84fdb498c11415828e3bf3e5e4818eab39bb8c1d222c5"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_windows_386.zip",
     "sha256_checksum": "d068b7bea4c9b9f59b32f2aea4fc7ee405f20feadc13b378d5627389865d1747"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_windows_amd64.zip",
     "sha256_checksum": "960fc4f9689a03f11a411f2604033479c27b3a3177f1c5ca21b9ca7fda254c24"
    }
   ]
  },
  {
   "version": "0.11.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_darwin_amd64.zip",
     "sha256_checksum": "6b6e8253b678554c67d717c42209fd857bfe64a1461763c05d3d1d85c6f618d3"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_freebsd_386.zip",
     "sha256_checksum": "b23e8bea87b6f7d97c5f92d9b78e32c65ed2e7548aa748888ba20a3d7d87a9f8"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_freebsd_amd64.zip",
     "sha256_checksum": "7bca8e846377a34cf0cc370d3c525d33e06756845afee68bb336fe345aefa615"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_freebsd_arm.zip",
     "sha256_checksum": "f85f3835282ee7f62e7d87cbee6f9c12d3dae4a77cacc77216d8a06b4a898dd1"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_linux_386.zip",
     "sha256_checksum": "eb29c6d9bb0d502cfd5d79cb57ff24775b6322163ec4bbc885be67faad2ef617"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_linux_amd64.zip",
     "sha256_checksum": "94504f4a67bad612b5c8e3a4b7ce6ca2772b3c1559630dfd71e9c519e3d6149c"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_linux_arm.zip",
     "sha256_checksum": "9252c06fa031059eef07cb6ee3e837b8c0324a1861779941c4b642765ca9e8f0"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_openbsd_386.zip",
     "sha256_checksum": "78c7a2f4979661e0cb555bea4f068a9cee6d4efe3d0fe002e6c4610d7aa348a"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_openbsd_amd64.zip",
     "sha256_checksum": "b5f011aa2338a4fe857d6a260f0e06ad0711453b51fdd9043880d2babaa4411e"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_solaris_amd64.zip",
     "sha256_checksum": "3ed21fb0c1a35f0593abb94369b062543c511772c1868e40361df1ebf71a330c"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_windows_386.zip",
     "sha256_checksum": "f97dbb8a43081b67bc3d7176115adee0a578d724a6b0885c506f62ad39c9b2b0"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_windows_amd64.zip",
     "sha256_checksum": "87252cc67486ef2ff2b8501c0bf9e795a53585b4dc5c09a8aa876c2564f77991"
    }
   ]
  },
  {
   "version": "0.15.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_darwin_amd64.zip",
     "sha256_checksum": "c7e413ad9d00a5ba177a32b0d46ff177239bd53a1aab67e7c5efad2e1e25978e"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_freebsd_386.zip",
     "sha256_checksum": "cc8bab95baf66d77bbb900e77ede50f7ae3a9eea701cde287eeb4c44dc129eaf"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_freebsd_amd64.zip",
     "sha256_checksum": "9da5e4c4d71ea21835c15e0f508dc533fea6f404d914d6825d842fab5700fe0c"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_freebsd_arm.zip",
     "sha256_checksum": "a909bd1c99f056787788c7e1635340172a44ee298165f8b99c98010b4d7c1b00"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_linux_386.zip",
     "sha256_checksum": "22b23f06f0169905d15d4bd8ad1d72edef00a06e17509880775a95e8aa2c5cbf"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_linux_amd64.zip",
     "sha256_checksum": "ddf9b409599b8c3b44d4e7c080da9a106befc1ff9e53b57364622720114e325c"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_linux_arm.zip",
     "sha256_checksum": "6d32766dc0f8eef754860a90a84fec623695a175ec246c4faac1ab5959110b4e"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_linux_arm64.zip",
     "sha256_checksum": "8bbbaf8b48f857b4044983cc757c5f05da5ab877449b9d9847c680b8955edc85"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_openbsd_386.zip",
     "sha256_checksum": "ced78f9579addd137f3cc39da3c2cec26473c58abc2d5560836361d0f008dc98"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_openbsd_amd64.zip",
     "sha256_checksum": "3ea4913faa7ec13bd2c37e05ff64442416f3cd44e71e759a8cf475500836e1b2"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_solaris_amd64.zip",
     "sha256_checksum": "b9e8bd65611d77e72b83f3bd873cc698f773850c63dd6c5705e4b025f6764bc6"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_windows_386.zip",
     "sha256_checksum": "3665e3a6d51d2ab1d694f981dda060147759129216380f0775b596d9d6dc8b07"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_windows_amd64.zip",
     "sha256_checksum": "1e7c7bfe5182c9d4700faaefa35472dd43b03793fd8502531783b21f70e80084"
    }
   ]
  },
  {
   "version": "1.2.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_darwin_amd64.zip",
     "sha256_checksum": "bd224d57718ed2b6e5e3b55383878d4b122c6dc058d65625605cef1ace9dcb25"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_darwin_arm64.zip",
     "sha256_checksum": "4750d46e47345809a0baa3c330771c8c8a227b77bec4caa7451422a21acefae5"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_freebsd_386.zip",
     "sha256_checksum": "e8393d6c8a5e203a65ec25d74a6a44d092a1825cf51f22f4a35baa73854beef1"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_freebsd_amd64.zip",
     "sha256_checksum": "3ef3e1fd92d23e4248d5d3cc295f16148778529792cbe3d40f1c4c05e599770a"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_freebsd_arm.zip",
     "sha256_checksum": "1e9b3a292b5367dc0c6d351df40a8ad07cc77abb2148a960819902d155dda8b0"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_linux_386.zip",
     "sha256_checksum": "153043ee6c3a49e7929cbe92a7296e45c7920658e1a1614ae6617843b5f65583"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_linux_amd64.zip",
     "sha256_checksum": "2934a0e8824925beb956b2edb5fef212a6141c089d29d8568150a43f95b3a626"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_linux_arm.zip",
     "sha256_checksum": "74850f769f861c3b8b70bd298500d5fe3093815035f00ed3d9ed67097f8f9cb9"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_linux_arm64.zip",
     "sha256_checksum": "9c6202237d7477412054dcd36fdc269da9ee66ecbc45bb07d0d63b7d36af7b21"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_openbsd_386.zip",
     "sha256_checksum": "a8da7ee4e01c6e7e4976178c6a426b09d2147e3674710d04e6bbfa4cf4713a07"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_openbsd_amd64.zip",
     "sha256_checksum": "9b253671eb59e99cf7a7e90126a377f882eb3cd53880fb706a70845a89ca2dd1"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_solaris_amd64.zip",
     "sha256_checksum": "76a15309343ccb086d35b8db392e019c45d1518dee12152acddb4e1382ecb91d"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_windows_386.zip",
     "sha256_checksum": "ad292884c4618f1765065add2f61758ba604400ebf84ed987b1b1bacd0c48fb1"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_windows_amd64.zip",
     "sha256_checksum": "753aea8cc3c06d547454e3758537b50f1f09f7b247d2ec5b45ba1e2c12d49127"
    }
   ]
  },
  {
   "version": "0.12.21",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_darwin_amd64.zip",
     "sha256_checksum": "f89b620e59439fccc80950bbcbd37a069101cbef7029029a12227eee831e463f"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_freebsd_386.zip",
     "sha256_checksum": "2fdbe03bd682e65f50a2b3885484be33bc837a82ffdf6fc6f981d4d4b79873d5"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_freebsd_amd64.zip",
     "sha256_checksum": "e85be1602551dda4b35b56f65f1fff8e8e0032e3901612d3d230a4923a189963"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_freebsd_arm.zip",
     "sha256_checksum": "cd3c6dda03d559f169e2c820c4d6e4c6e4b4ec00e4ba8e62c5677e04cefa8726"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_linux_386.zip",
     "sha256_checksum": "4f2a4d749755e186fe71aa831151a676fd11fd35c7c2bf02e221a988d3928cf1"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_linux_amd64.zip",
     "sha256_checksum": "ca0d0796c79d14ee73a3d45649dab5e531f0768ee98da71b31e423e3278e9aa9"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_linux_arm.zip",
     "sha256_checksum": "37c06dbe2172faffa81b95720764e865d3901414ec148019508d1b0cefcf6320"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_openbsd_386.zip",
     "sha256_checksum": "25b7f79531861735cb4c14b438aa845040d80363e4a4a1c0bd23c90de7f3c336"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_openbsd_amd64.zip",
     "sha256_checksum": "4dd26f371698d6285a3b96576cb9b83f89b8af915ec92a16b9e161dd551c07fd"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_solaris_amd64.zip",
     "sha256_checksum": "98d48daa46e63f9e6d3e152f061eb1289be6ca4327ef47ee1d629f47de13456c"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_windows_386.zip",
     "sha256_checksum": "b662994294709bd7fdea67b5522310c5fd1ef1446fb78f8868ed672247800355"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_windows_amd64.zip",
     "sha256_checksum": "254e5f870efe9d86a3f211a1b9c3c01325fc380e428f54542b7750d8bfd62bb1"
    }
   ]
  },
  {
   "version": "0.3.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_darwin_386.zip",
     "sha256_checksum": "f4a42243abb2b3d6fde815f43480217982bc676f0510c46c05f544546c4d2d89"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_darwin_amd64.zip",
     "sha256_checksum": "dda41425c7eb06c5e8b3f5ad4904e993aa8a9ab6b61f954ee2e259667cb6ff57"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_freebsd_386.zip",
     "sha256_checksum": "a3751daa08f0e30087d1ce58cc4200d30d647e8687fc6d285f57d5f5e6d61aaf"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_freebsd_amd64.zip",
     "sha256_checksum": "d1a54b62edd9e575dda3b1ff91fdf22c5aed380e04807228dd4bab8d54b81ba8"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_freebsd_arm.zip",
     "sha256_checksum": "84427ce689b5070ec168009d7e296aa2fd69423bf2d8bf9acb1cd502a01d828"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_linux_386.zip",
     "sha256_checksum": "f773e9f75c055789e67dde80db4306d2f922cfa2f7b0d9c6413412778c594ffd"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_linux_amd64.zip",
     "sha256_checksum": "bd7b3e352a186010471818f8c50578e4afa9a2d2cf71acffb1810292db725e33"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_linux_arm.zip",
     "sha256_checksum": "8f742a7753cb5ed48e7dd04dec4875b107265c82bfe9df1e362012aa6cd60839"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_openbsd_386.zip",
     "sha256_checksum": "a6001660534adf695923cb1e9a4fb85134472c9b59af4ca0bed7e4991e7763f0"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_openbsd_amd64.zip",
     "sha256_checksum": "adf2e37a3b30ab2b0665b0cb30721fd9fe97566aa4e65c4b5928910d40e5e0a"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_windows_386.zip",
     "sha256_checksum": "e5ac6a68224e2a9b4774237cf37da39722239b1a058082fb830288c20691225a"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_windows_amd64.zip",
     "sha256_checksum": "76ff60d10fe74556d7d20c557c15e853c5900036cf14576584a459fd84c78e30"
    }
   ]
  },
  {
   "version": "0.7.13",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_darwin_386.zip",
     "sha256_checksum": "ca3f64bd815365faa2980ba7f6a17c1a1e7fa92b6e96155d37de3e171eb7417f"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_darwin_amd64.zip",
     "sha256_checksum": "c1e004ad2bff4e92edb13cf32a18b67b5178fc3597a844beeda09cc4f9c30b65"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_freebsd_386.zip",
     "sha256_checksum": "ce42c97be8d05230b25045a8da5f04b0f96d8353fbef8fbbd4127812eb81b7d9"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_freebsd_amd64.zip",
     "sha256_checksum": "c5bde373f337c66f3f735d7cb060a9e0ec35d9d90d1170db7c2504ce03637569"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_freebsd_arm.zip",
     "sha256_checksum": "6a8b27a66694f4f5c68f2d3eab7ca97dec2cfff3aa3f0ebc95ab2361b005a4dc"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_linux_386.zip",
     "sha256_checksum": "69eeeec837837ff5c7c796a86f7ac5b1fb442220c228b4358b450cd906f8a3c3"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_linux_amd64.zip",
     "sha256_checksum": "5a4f762a194542d38406b9b92c722b57f7910344db084e24c9c43d7719f4aa18"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_linux_arm.zip",
     "sha256_checksum": "2c0a23bde70ae7da93f5d9c7536ee825040ed11fe82954b761d243a451b31c6b"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_openbsd_386.zip",
     "sha256_checksum": "5a49a4c3f15f6470dfef85c91671a3d54bbaa5820938d82896de252247a70736"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_openbsd_amd64.zip",
     "sha256_checksum": "79e58d9ce0dd3594ecbfdee08cb551b57c6173e52167ac58afa187bf4346f44f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_solaris_amd64.zip",
     "sha256_checksum": "6af3527daa76318b79fee8c6c03c4d2d54d9ed0d0b916bfbc25b58a8e63f0182"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_windows_386.zip",
     "sha256_checksum": "55586a3dac85f56c25666ea339b4f94a21fc23e450652f6f9b068b5fabdcba4"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_windows_amd64.zip",
     "sha256_checksum": "f2dc1d85da946c472388087565f769d3e112e8f1ebfba44e0421a0f17a9db56b"
    }
   ]
  },
  {
   "version": "0.12.20",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_darwin_amd64.zip",
     "sha256_checksum": "9e2ef974618402b70d4491f50701621e1a9f1cb32862592f0af3fee12324d378"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_freebsd_386.zip",
     "sha256_checksum": "8b3ef60931ecd64b54f69a963b7838a6fa9da9feb6fec97ed6deb42e88d67b77"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_freebsd_amd64.zip",
     "sha256_checksum": "79ff462ebd5975e7d2da5da240858b58252d36ccb97c7d03ec40591f6f8a09d6"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_freebsd_arm.zip",
     "sha256_checksum": "dd0b3305f6cb2f2582f6fc138b141aea0ba2df3834cfed33459b904c348c0955"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_linux_386.zip",
     "sha256_checksum": "68164945a4e829c28e1ccfc6cf4700c5e04686b7ddad425ab6f858a99e60e861"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip",
     "sha256_checksum": "46bd906f8cb9bbb871905ecb23ae7344af8017d214d735fbb6d6c8e0feb20ff3"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_linux_arm.zip",
     "sha256_checksum": "6ef8bf5bf39e6d9f75794b73154b10ae466750533bb84050d2da9e5c928faab7"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_openbsd_386.zip",
     "sha256_checksum": "5c85a1812e73db284b6706a1d228926d26844df75e0c104d68cb008e461a67a4"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_openbsd_amd64.zip",
     "sha256_checksum": "220108046f82a2052aff9ca16e782c09b79ca0d08143eb02c6721e8b662ecac2"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_solaris_amd64.zip",
     "sha256_checksum": "a5febb1b0940ec4aeaec8df6f08d1d552360b8447ec4cc623eadd8aaef4459b1"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_windows_386.zip",
     "sha256_checksum": "3e7c6aca476f7c82a96106d854f72851715483939667d99ebb0118a944592b6b"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_windows_amd64.zip",
     "sha256_checksum": "7213442bb7b7b554ec0895a2917a71e0e2945ffbae5424bb92696d9213f50bb6"
    }
   ]
  },
  {
   "version": "1.0.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_darwin_amd64.zip",
     "sha256_checksum": "e7595530a0dcdaec757621cbd9f931926fd904b1a1e5206bf2c9db6b73cee04d"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_darwin_arm64.zip",
     "sha256_checksum": "eecea1343888e2648d5f7ea25a29494fd3b5ecde95d0231024414458c59cb184"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_freebsd_386.zip",
     "sha256_checksum": "c426e7c07567354089ac0c15e07d570d763ebaac2a8147ba92df9ddcb889095a"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_freebsd_amd64.zip",
     "sha256_checksum": "98f5d74eaf80757fbb888ab5024065c757afb014ac4b48eabf8c8fb19af5abf8"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_freebsd_arm.zip",
     "sha256_checksum": "23e0cb2a77a53e80a8166eb158eef3464213877177cdee319de0b5d6d7f7ff4e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_linux_386.zip",
     "sha256_checksum": "ad2c506fb1fa1e1e9f814afb1e5a31988f5497922dade8a52b973b559870f107"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_linux_amd64.zip",
     "sha256_checksum": "a221682fcc9cbd7fde22f305ead99b3ad49d8303f152e118edda086a2807716d"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_linux_arm.zip",
     "sha256_checksum": "4bcf91e002f0111e13abb40e7867405c019911f7748830ab79b22ff5972aaf66"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_linux_arm64.zip",
     "sha256_checksum": "b091dbe5c00785ae8b5cb64149d697d61adea75e495d9e3d910f61d8c9967226"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_openbsd_386.zip",
     "sha256_checksum": "a933a188eb0d2a6738002ede35a8cc546e748bcaf61b4684bff2b115bb10a21f"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_openbsd_amd64.zip",
     "sha256_checksum": "7fbd5aaef330b4d59bcfcf9abda64c1dee3d8210be08eb9f83aa11d23bbdb5d8"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_solaris_amd64.zip",
     "sha256_checksum": "1aa9cb6c3d7b54630e6ff4c95ea603589a0c1689fe66e05f4ad1b551da2fc5db"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_windows_386.zip",
     "sha256_checksum": "62f7f31a2d485aab566ddff1f3b10a3ac9fe15b30d658e7750cf817289b74a65"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.10/terraform_1.0.10_windows_amd64.zip",
     "sha256_checksum": "1c7f6f471290ef740b999cdd6175ca2ef52075384ff525bfbd92da503d2bb701"
    }
   ]
  },
  {
   "version": "0.11.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_darwin_amd64.zip",
     "sha256_checksum": "ff5c3c4bcfe84e011b96a2232704b2db196383ce5d4a32e47956c883ddc94bac"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_freebsd_386.zip",
     "sha256_checksum": "f73267ebfb45bf56e3f92c04f89aed14c07f3e1601614f3f8edbc20d34d5a023"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_freebsd_amd64.zip",
     "sha256_checksum": "97e9841b55f7ed12e8c793cca13e84c92534c06a1c626b801f91496212c77453"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_freebsd_arm.zip",
     "sha256_checksum": "4556a90639a2dfff1235945170d83c453219f733e4f5458d7403c38782f776c6"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_linux_386.zip",
     "sha256_checksum": "637b0299461cf358ea534014ee2a929b5c059d501bbda2b8b5b86918fc3cb68d"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_linux_amd64.zip",
     "sha256_checksum": "f728fa73ff2a4c4235a28de4019802531758c7c090b6ca4c024d48063ab8537b"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_linux_arm.zip",
     "sha256_checksum": "fd718355933319dfd6d71f7bd7ff99c837047277dc06408c69bca095c32c5c23"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_openbsd_386.zip",
     "sha256_checksum": "195c4f2c08788740eabec19282be042b9fc540488632bc106b863eee7ea1b3da"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_openbsd_amd64.zip",
     "sha256_checksum": "5aca120b85307ea77c0d1070bd7a6ea1d6df72641ada0a5d36a502dde4a3fb19"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_solaris_amd64.zip",
     "sha256_checksum": "bf8dadc6d20ae0f22fab1cc86154d58a1d757e94d1c94511b36627c0a9ba803c"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_windows_386.zip",
     "sha256_checksum": "91b9a021b85d5ffdeafad155fb323f48ea47b535e44e9bd600e3ca2997b22ca8"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_windows_amd64.zip",
     "sha256_checksum": "f8e12b552eb2570bfa0ad52d18ef327b0704d54ef70317e0192a6c07f9813b6d"
    }
   ]
  },
  {
   "version": "0.9.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_darwin_amd64.zip",
     "sha256_checksum": "4140c52917da91a276db34f01e5efc27d07b6e1deeede4137625fccf7bfabb83"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_freebsd_386.zip",
     "sha256_checksum": "a2fa14befad367b2e583c6e902170aff8744d0a39a6814a21bb6b93694719ed3"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_freebsd_amd64.zip",
     "sha256_checksum": "53d221ad7373e844894d2a0488dc81c4755cdb273a2f00997cc564072b388ea8"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_freebsd_arm.zip",
     "sha256_checksum": "db9853a978324eeeff2eccd6faf49fb9fe7ee1ad8f605f20907876ae99fde5e0"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_linux_386.zip",
     "sha256_checksum": "ba4fe982ff9bad14223d1e984e25861b71a98f4f4d81ac157777b03ae1ce138c"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_linux_amd64.zip",
     "sha256_checksum": "b3b18a719258dcc02b7b972eedf417be0b497e4129063711bca82877dbe65553"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_linux_arm.zip",
     "sha256_checksum": "5967734c4193cc693e8188a52d3f3847d0ff24c9b303d06dcdbd66662eb88771"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_openbsd_386.zip",
     "sha256_checksum": "488a7caa205fb8b7dc55aff141b0fe3b8d4d0b5948a9fbc94dfddcb66470f1e7"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_openbsd_amd64.zip",
     "sha256_checksum": "d352b740cc3b2577bbe2fe2ff29dac9b1cf28657a1461e34a3ac7f482b42c7e"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_solaris_amd64.zip",
     "sha256_checksum": "92e61f0caec282ccb3c487f6a17f2f58e56a484dfb03a0eea9e8a64785b7e5d0"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_windows_386.zip",
     "sha256_checksum": "8949e4736852fd5e71bdda35d0e4a053ed8827fb6d20409d117f79cf68d3267"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_windows_amd64.zip",
     "sha256_checksum": "12508576ffd0c01901794d16adb7fc9389be6a2911cbdfa6c5611576b1352843"
    }
   ]
  },
  {
   "version": "0.11.12",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_darwin_amd64.zip",
     "sha256_checksum": "316fa873b26463f3e015db11dba00eab1839338f930f1352dbab2d0bcd0828a5"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_freebsd_386.zip",
     "sha256_checksum": "42c821df7d6d99ad223b7b972db3b13a5cb78d5660fa4ed3f2f5aabcca522a17"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_freebsd_amd64.zip",
     "sha256_checksum": "71f6064aead8f19d1d1c3c3a0713a15bbfdfa56821f721c621b346bea9217f01"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_freebsd_arm.zip",
     "sha256_checksum": "d9a28b972da7c953aa965701bb4674d9607e339c4a098502f661b58df0212d4e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_linux_386.zip",
     "sha256_checksum": "94dabed58871730862f329cf8503bb88702b92a69f563d27e323c91e17dcc7a"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_linux_amd64.zip",
     "sha256_checksum": "d3bb9c958c56a178528ef3b18e27a24cfd96c9aa6da3c7b6dc8d7dd8a4b9dab9"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_linux_arm.zip",
     "sha256_checksum": "e84d4111d587c922e31af7fc9ef3d1869b172113ef329ce3b54f8b471c5ca467"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_openbsd_386.zip",
     "sha256_checksum": "cf8e736831361e1004b482557adf4342950e62929b6445c6d28c68b429fbb4e0"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_openbsd_amd64.zip",
     "sha256_checksum": "522fe82632c217bac16f3a3448226395722c5c117fdde706c3cb2145fcf4c825"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_solaris_amd64.zip",
     "sha256_checksum": "e12d07effffcfd6adbdd533bc3eead92b4cec8fe3634aed4f6b4b12f8e12314f"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_windows_386.zip",
     "sha256_checksum": "a600447a98b5078e29253e34a33a029ac1199f641125e7ff3f5d3645cc3a0470"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_windows_amd64.zip",
     "sha256_checksum": "8bcde36ba8cc415322abf68a2a47219e0a3bd001a66a7728dbfb332b65cdd269"
    }
   ]
  },
  {
   "version": "0.9.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_darwin_amd64.zip",
     "sha256_checksum": "b6de7307c989455c4b1f351c2df1ad1a0308edc71868bb432ad74f3980f8a6a3"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_freebsd_386.zip",
     "sha256_checksum": "1d3c4a57e3f25b572ee8101cb5cea5df20610cdc572f22966247422558e02e74"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_freebsd_amd64.zip",
     "sha256_checksum": "e6c636853af94f79d1cf3a40d13f8e0bf0730cbe48814490f462641903013005"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_freebsd_arm.zip",
     "sha256_checksum": "7267ffaa00583488585ba952ba9d2b5f3d4e3333fa3d17b7bfd2f86924a53f78"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_linux_386.zip",
     "sha256_checksum": "4e05088a88242c95c9551f6824c2894dbe96483b64529a4d274c1d17b7a049b7"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_linux_amd64.zip",
     "sha256_checksum": "965cf0f9d373a550163aef4732438c201c7271ea5f96a902863248cb7f39f1c"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_linux_arm.zip",
     "sha256_checksum": "f2de255e7dd8eab3d47772f562fa27b4b94579bfd623fd28923b95beffe552de"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_openbsd_386.zip",
     "sha256_checksum": "e02b603e60ac5fe8dc1e9dcd8a8d3affbd7d9432252d7f79ea04f95dcba9376d"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_openbsd_amd64.zip",
     "sha256_checksum": "f05093f637c90739dc856d8f967f5449bb5b221bc65dbacbab9571ba7bb53f43"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_solaris_amd64.zip",
     "sha256_checksum": "9c37e83fd078e5707167547df6445326f03a1a4f4b3fdda9fb8b17f90bd5a014"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_windows_386.zip",
     "sha256_checksum": "fedc695b774b9f23e208451e7c2f2078575a45a4f6d8daf18d295fdc4e60e86b"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_windows_amd64.zip",
     "sha256_checksum": "24a7ecfc89c986dc4be20a4f50427054115874352f6b0d7337157f3ca62cf9af"
    }
   ]
  },
  {
   "version": "0.10.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_darwin_amd64.zip",
     "sha256_checksum": "3f05acdf0a9e04ba7e3bda18521feb0b310462dcce62c454854a40519b1695ed"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_freebsd_386.zip",
     "sha256_checksum": "e92667ac24deab9706149998274b90b11c8f2acb14c3069853340f48bdf8ab46"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_freebsd_amd64.zip",
     "sha256_checksum": "1ff0806d25172d8d8d9f3d5847552a49315710982842f10ae73d38c6e001ad7f"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_freebsd_arm.zip",
     "sha256_checksum": "95c5e238939138a38c53c4e66b2f6d90b52356d5200641a1c1b099027dea4a6e"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_linux_386.zip",
     "sha256_checksum": "a13ee4fb98fe0b219678934d72efaafd4b59fe10036a7de4cb769d51a9246ee5"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_linux_amd64.zip",
     "sha256_checksum": "b786c0cf936e24145fad632efd0fe48c831558cc9e43c071fffd93f35e3150db"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_linux_arm.zip",
     "sha256_checksum": "f0dd85b3fbdc0e873dd7cc59b96ce2d70467d8326575d79e853718b0baa79cc0"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_openbsd_386.zip",
     "sha256_checksum": "1a3360b38e5a62437d96143a4784c4341c016036a9fe4d67c796206021dc6b7"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_openbsd_amd64.zip",
     "sha256_checksum": "16cf47a865a562b45d8ccc2ee76acd3a24aeedbf29cf9961348bbc05cee61402"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_solaris_amd64.zip",
     "sha256_checksum": "5036abc58b46fb481a08eca471070f74aab4bfae84a8abe1256f48931c9b4387"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_windows_386.zip",
     "sha256_checksum": "2518afa45d8b364802b5b58d504bbc76a47597990dffaad7292ac8ba60ec7189"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_windows_amd64.zip",
     "sha256_checksum": "6bbaf59aaf80b476ecbf765369678f5d8c296cb3a30e4d783ad560f0e4b81a76"
    }
   ]
  },
  {
   "version": "0.12.16",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_darwin_amd64.zip",
     "sha256_checksum": "2f893e326b25705aff2594d9f28a4a0c9d50f44a0e7e7129633f02c11a2e47d"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_freebsd_386.zip",
     "sha256_checksum": "a430469db0b661f198cde8c90e969637df6cf8d2368630008a1ea0df33391eda"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_freebsd_amd64.zip",
     "sha256_checksum": "5d417f94b2819e0b4f1c0814f5a26fe6c443de4df98b2a577c99b100004f2564"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_freebsd_arm.zip",
     "sha256_checksum": "2228663ea703e49a15dd24a154755a1168675ef80b2b982af07a845ae59f333c"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_linux_386.zip",
     "sha256_checksum": "8d7b567e4637d39e3a91768c2055eb57179a8703a34860ff269c35c6e331fc8e"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_linux_amd64.zip",
     "sha256_checksum": "fcc719314660adc66cbd688918d13baa1095301e2e507f9ac92c9e22acf4cc02"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_linux_arm.zip",
     "sha256_checksum": "d3c9af9aa42a2eceffc4a349bea46bb5718fd38b77ed37e946aa7c42865f5835"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_openbsd_386.zip",
     "sha256_checksum": "a20ddd90ef94e61d42d026a5a952f1af860b1b364d3e7949dfc2358316668fc1"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_openbsd_amd64.zip",
     "sha256_checksum": "f459339fcf82864acc98de0852c0f8f7bf12bb48f55f5c19766112bd13b9be90"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_solaris_amd64.zip",
     "sha256_checksum": "aeed2bf25b82937044accddb072db718cb739292e8a506ec201b3f626bca94fa"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_windows_386.zip",
     "sha256_checksum": "87c519e132f5f0901a1fd953c63bf331f4e7558411a8cd6390ecffedcb04dbd6"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_windows_amd64.zip",
     "sha256_checksum": "af330eadf2129c1bbdde65f21b78b1fda52e4ba204cca112fb5a160db73385bb"
    }
   ]
  },
  {
   "version": "0.11.14",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_darwin_amd64.zip",
     "sha256_checksum": "829bdba148afbd61eab4aafbc6087838f0333d8876624fe2ebc023920cfc2ad5"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_freebsd_386.zip",
     "sha256_checksum": "f09d763c3918b13b0929b04371f002da7f922fc88465244ccc87e1c7d7bbb245"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_freebsd_amd64.zip",
     "sha256_checksum": "1dbf6275e1c0b55a1082571180c4ed0faf279dba9c7649c4f37fca46f6d9f590"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_freebsd_arm.zip",
     "sha256_checksum": "6f3724b182a89e47b813c32b2891c96540405ecf7b8680163bbbb4ac82670649"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_linux_386.zip",
     "sha256_checksum": "b6b2c61b80a35646df2cb7d443efeba3f4dedcdecbabab3b2626c2ea8976e87"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_linux_amd64.zip",
     "sha256_checksum": "9b9a4492738c69077b079e595f5b2a9ef1bc4e8fb5596610f69a6f322a8af8dd"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_linux_arm.zip",
     "sha256_checksum": "73ab704f5395a7221cc943d7fb1e07d4256397d148d83f1c301cf0b8a11902f0"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_openbsd_386.zip",
     "sha256_checksum": "8bd4a06473fadf2a40d670334228b30e9134f31acffea29ead5c159fb83f1cb4"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_openbsd_amd64.zip",
     "sha256_checksum": "417466ac49fcbe12153e0a6cf7c429f3fabc4ef4df18c6f0711d5c61da9b7a0f"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_solaris_amd64.zip",
     "sha256_checksum": "4bc1b6e00aaadbe04778c023ec53d9b73fe2662031393e51b7354699947a9bb8"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_windows_386.zip",
     "sha256_checksum": "f2eb847761cba796f306880288083b4c68f5ae9dd86c6cff47023eecc9895f8f"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_windows_amd64.zip",
     "sha256_checksum": "bfec66e2ad079a1fab6101c19617a82ef79357dc1b92ddca80901bb8d5312dc0"
    }
   ]
  },
  {
   "version": "0.12.0-alpha2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_darwin_amd64.zip",
     "sha256_checksum": "859fa4459f8cc8b4cda026b71cd7c8011fafc765e570fbdf3abe9fbcad44d59c"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_freebsd_386.zip",
     "sha256_checksum": "2a0400e68cf225700b631daea4241092630cfcf394fb77e70829dceb2189defc"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_freebsd_amd64.zip",
     "sha256_checksum": "b6024d02e7566fb12b0ff8184d46fb12b6324c6044a2f2404d986dedeaf3fa61"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_freebsd_arm.zip",
     "sha256_checksum": "fdc54dff88487734b7c9cd6420c849cbc250ef9f2bd3451e9960bb44c467750"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_linux_386.zip",
     "sha256_checksum": "50649ecafa606afaa5a76331cd915fd83b01e5400cfb5a5b9fa3a7d9f57d218b"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_linux_amd64.zip",
     "sha256_checksum": "14d222094f10fc605131ed0629c34008943c10eb782d4c02db05032d7928c048"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_linux_arm.zip",
     "sha256_checksum": "d15418e1b7cc10a2763a4892a18fdae2926b17a022d68a9eb14efe7e047b059f"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_openbsd_386.zip",
     "sha256_checksum": "a4a9c3fd4b5f19b5f5e5cccf7e4f5ffbe458d5758df5cfc604691cde7aa30d00"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_openbsd_amd64.zip",
     "sha256_checksum": "4aed295b91de1aa5ca927d123b64cbcde043f57f918a1090e7f3768b3340de30"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_solaris_amd64.zip",
     "sha256_checksum": "4c3bbc53c6801dc07d0eaa65b4b1b81edad007c45ac06cd557318b6c79582896"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_windows_386.zip",
     "sha256_checksum": "f2c7f9ea62cfb5a7515119d2e61d6b1eac016ab7ab7daa8e54b777b667828342"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_windows_amd64.zip",
     "sha256_checksum": "n/a"
    }
   ]
  }
 ]
}`
