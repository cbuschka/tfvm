package inventory

const defaultStateJSON = `
{
 "lastUpdateTime": "2021-09-14T22:44:18+02:00",
 "terraformReleases": [
  {
   "version": "0.1.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.1.0/terraform_0.1.0_darwin_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.1.0/terraform_0.1.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.1.0/terraform_0.1.0_linux_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.1.0/terraform_0.1.0_windows_386.zip"
    }
   ]
  },
  {
   "version": "0.1.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.1.1/terraform_0.1.1_darwin_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.1.1/terraform_0.1.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.1.1/terraform_0.1.1_linux_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.1.1/terraform_0.1.1_windows_386.zip"
    }
   ]
  },
  {
   "version": "0.2.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.2.0/terraform_0.2.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.2.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.2.1/terraform_0.2.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.2.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.2.2/terraform_0.2.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.3.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.3.0/terraform_0.3.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.3.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.3.1/terraform_0.3.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.3.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.3.5/terraform_0.3.5_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.3.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.3.6/terraform_0.3.6_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.3.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.3.7/terraform_0.3.7_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.4.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.4.0/terraform_0.4.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.4.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.4.1/terraform_0.4.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.4.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.4.2/terraform_0.4.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.5.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.5.0/terraform_0.5.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.5.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.5.1/terraform_0.5.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.5.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.5.3/terraform_0.5.3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_linux_arm.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.0/terraform_0.6.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.1/terraform_0.6.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.2/terraform_0.6.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.3/terraform_0.6.3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.4/terraform_0.6.4_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.5/terraform_0.6.5_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.6/terraform_0.6.6_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.7/terraform_0.6.7_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.8/terraform_0.6.8_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.9/terraform_0.6.9_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.10/terraform_0.6.10_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.11/terraform_0.6.11_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.12",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.12/terraform_0.6.12_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.13",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.13/terraform_0.6.13_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.14",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.14/terraform_0.6.14_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.15",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.15/terraform_0.6.15_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.6.16",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.6.16/terraform_0.6.16_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.0/terraform_0.7.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.1/terraform_0.7.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.2/terraform_0.7.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.3/terraform_0.7.3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.4/terraform_0.7.4_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.5/terraform_0.7.5_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.6/terraform_0.7.6_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.7/terraform_0.7.7_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.8/terraform_0.7.8_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.9/terraform_0.7.9_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.10/terraform_0.7.10_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.11/terraform_0.7.11_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.12",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.12/terraform_0.7.12_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.7.13",
   "builds": [
    {
     "os": "darwin",
     "arch": "386",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_darwin_386.zip"
    },
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.7.13/terraform_0.7.13_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.8.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.0/terraform_0.8.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.8.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.1/terraform_0.8.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.8.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.2/terraform_0.8.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.8.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.3/terraform_0.8.3_openbsd_amd64.zip"
    }
   ]
  },
  {
   "version": "0.8.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.4/terraform_0.8.4_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.8.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.5/terraform_0.8.5_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.8.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.6/terraform_0.8.6_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.8.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.7/terraform_0.8.7_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.8.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.8.8/terraform_0.8.8_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.9.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.0/terraform_0.9.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.9.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.1/terraform_0.9.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.9.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.2/terraform_0.9.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.9.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.3/terraform_0.9.3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.9.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.4/terraform_0.9.4_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.9.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.5/terraform_0.9.5_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.9.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.6/terraform_0.9.6_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.9.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.7/terraform_0.9.7_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.9.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.8/terraform_0.9.8_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.9.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.9/terraform_0.9.9_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.9.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.10/terraform_0.9.10_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.9.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.9.11/terraform_0.9.11_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.10.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta1/terraform_0.10.0-beta1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.10.0-beta2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-beta2/terraform_0.10.0-beta2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.10.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0-rc1/terraform_0.10.0-rc1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.10.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.0/terraform_0.10.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.10.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.1/terraform_0.10.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.10.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.2/terraform_0.10.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.10.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.3/terraform_0.10.3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.10.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.4/terraform_0.10.4_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.10.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.5/terraform_0.10.5_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.10.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.6/terraform_0.10.6_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.10.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.7/terraform_0.10.7_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.10.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.10.8/terraform_0.10.8_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-beta1/terraform_0.11.0-beta1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0-rc1/terraform_0.11.0-rc1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.0/terraform_0.11.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.1/terraform_0.11.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.2/terraform_0.11.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.3/terraform_0.11.3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.4/terraform_0.11.4_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.5/terraform_0.11.5_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.6/terraform_0.11.6_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.7/terraform_0.11.7_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.8/terraform_0.11.8_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.9-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9-beta1/terraform_0.11.9-beta1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.9/terraform_0.11.9_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.10/terraform_0.11.10_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.11/terraform_0.11.11_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.12-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12-beta1/terraform_0.11.12-beta1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.12",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.12/terraform_0.11.12_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.13",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.13/terraform_0.11.13_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.14",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.14/terraform_0.11.14_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.15-oci",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15-oci/terraform_0.11.15-oci_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.11.15",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.11.15/terraform_0.11.15_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.0-alpha1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha1/terraform_0.12.0-alpha1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.0-alpha2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-alpha2/terraform_0.12.0-alpha2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.0-alpha3",
   "builds": [
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_darwin_amd64.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_freebsd_386.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_freebsd_amd64.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_freebsd_arm.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_linux_386.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_linux_amd64.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_linux_arm.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_openbsd_386.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_openbsd_amd64.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_solaris_amd64.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_windows_386.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha3/terraform_0.12.0-alpha3_terraform_0.12.0-alpha3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.0-alpha4",
   "builds": [
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_darwin_amd64.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_freebsd_386.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_freebsd_amd64.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_freebsd_arm.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_linux_386.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_linux_amd64.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_linux_arm.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_openbsd_386.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_openbsd_amd64.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_solaris_amd64.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_windows_386.zip"
    },
    {
     "os": "terraform",
     "arch": "0.12",
     "download_path": "/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_openbsd_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta1/terraform_0.12.0-beta1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.0-beta2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-beta2/terraform_0.12.0-beta2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0-rc1/terraform_0.12.0-rc1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.0/terraform_0.12.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.1/terraform_0.12.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.2/terraform_0.12.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.3/terraform_0.12.3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.4/terraform_0.12.4_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.5/terraform_0.12.5_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.6/terraform_0.12.6_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.7/terraform_0.12.7_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.8/terraform_0.12.8_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.9/terraform_0.12.9_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.10/terraform_0.12.10_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.11/terraform_0.12.11_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.12",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.12/terraform_0.12.12_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.13",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.13/terraform_0.12.13_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.14",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.14/terraform_0.12.14_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.15",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.15/terraform_0.12.15_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.16",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.16/terraform_0.12.16_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.17",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.17/terraform_0.12.17_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.18",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.18/terraform_0.12.18_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.19",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.19/terraform_0.12.19_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.20",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.20/terraform_0.12.20_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.21",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.21/terraform_0.12.21_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.22",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.22/terraform_0.12.22_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.23",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.23/terraform_0.12.23_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.24",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.24/terraform_0.12.24_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.25",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.25/terraform_0.12.25_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.26",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.26/terraform_0.12.26_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.27",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.27/terraform_0.12.27_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.28",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.28/terraform_0.12.28_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.29",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.29/terraform_0.12.29_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.30",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.30/terraform_0.12.30_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.12.31",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.12.31/terraform_0.12.31_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.13.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta1/terraform_0.13.0-beta1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.13.0-beta2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta2/terraform_0.13.0-beta2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.13.0-beta3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-beta3/terraform_0.13.0-beta3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.13.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0-rc1/terraform_0.13.0-rc1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.13.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.0/terraform_0.13.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.13.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.1/terraform_0.13.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.13.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.2/terraform_0.13.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.13.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.3/terraform_0.13.3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.13.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.4/terraform_0.13.4_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.13.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.5/terraform_0.13.5_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.13.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.6/terraform_0.13.6_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.13.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.13.7/terraform_0.13.7_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.0-alpha20200910",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200910/terraform_0.14.0-alpha20200910_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.0-alpha20200923",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_linux_arm.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20200923/terraform_0.14.0-alpha20200923_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.0-alpha20201007",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-alpha20201007/terraform_0.14.0-alpha20201007_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta1/terraform_0.14.0-beta1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.0-beta2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-beta2/terraform_0.14.0-beta2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0-rc1/terraform_0.14.0-rc1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.0/terraform_0.14.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.1/terraform_0.14.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.2/terraform_0.14.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.3/terraform_0.14.3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.4/terraform_0.14.4_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.5/terraform_0.14.5_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.6/terraform_0.14.6_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.7",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.7/terraform_0.14.7_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.8",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.8/terraform_0.14.8_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.9",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.9/terraform_0.14.9_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.10",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.10/terraform_0.14.10_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.14.11",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.14.11/terraform_0.14.11_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.0-alpha20210107",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210107/terraform_0.15.0-alpha20210107_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.0-alpha20210127",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210127/terraform_0.15.0-alpha20210127_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.0-alpha20210210",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-alpha20210210/terraform_0.15.0-alpha20210210_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.0-beta1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta1/terraform_0.15.0-beta1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.0-beta2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-beta2/terraform_0.15.0-beta2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.0-rc1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc1/terraform_0.15.0-rc1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.0-rc2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0-rc2/terraform_0.15.0-rc2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.0/terraform_0.15.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.1/terraform_0.15.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.2/terraform_0.15.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.3/terraform_0.15.3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.4/terraform_0.15.4_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "0.15.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/0.15.5/terraform_0.15.5_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.0.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.0/terraform_1.0.0_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.0.1",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.1/terraform_1.0.1_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.0.2",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_darwin_amd64.zip"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_darwin_arm64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.2/terraform_1.0.2_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.0.3",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_darwin_amd64.zip"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_darwin_arm64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.3/terraform_1.0.3_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.0.4",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_darwin_amd64.zip"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_darwin_arm64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.4/terraform_1.0.4_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.0.5",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_darwin_amd64.zip"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_darwin_arm64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.5/terraform_1.0.5_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.0.6",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_darwin_amd64.zip"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_darwin_arm64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.0.6/terraform_1.0.6_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210616",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210616/terraform_1.1.0-alpha20210616_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210630",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_darwin_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210630/terraform_1.1.0-alpha20210630_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210714",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_darwin_amd64.zip"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_darwin_arm64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210714/terraform_1.1.0-alpha20210714_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210728",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_darwin_amd64.zip"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_darwin_arm64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210728/terraform_1.1.0-alpha20210728_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210811",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_darwin_amd64.zip"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_darwin_arm64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210811/terraform_1.1.0-alpha20210811_windows_amd64.zip"
    }
   ]
  },
  {
   "version": "1.1.0-alpha20210908",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_darwin_amd64.zip"
    },
    {
     "os": "darwin",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_darwin_arm64.zip"
    },
    {
     "os": "freebsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_freebsd_386.zip"
    },
    {
     "os": "freebsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_freebsd_amd64.zip"
    },
    {
     "os": "freebsd",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_freebsd_arm.zip"
    },
    {
     "os": "linux",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_linux_386.zip"
    },
    {
     "os": "linux",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_linux_amd64.zip"
    },
    {
     "os": "linux",
     "arch": "arm",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_linux_arm.zip"
    },
    {
     "os": "linux",
     "arch": "arm64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_linux_arm64.zip"
    },
    {
     "os": "openbsd",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_openbsd_386.zip"
    },
    {
     "os": "openbsd",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_openbsd_amd64.zip"
    },
    {
     "os": "solaris",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_solaris_amd64.zip"
    },
    {
     "os": "windows",
     "arch": "386",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_windows_386.zip"
    },
    {
     "os": "windows",
     "arch": "amd64",
     "download_path": "/terraform/1.1.0-alpha20210908/terraform_1.1.0-alpha20210908_windows_amd64.zip"
    }
   ]
  }
 ]
}`
