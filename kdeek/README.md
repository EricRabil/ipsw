# kdeek
Why wait for KDKs when you can make them?

## Disclaimer
I'm really sorry if this is messy, I made this in a 9 hour sprint fueled by a wawa hoagie, cold brew, some ice cream, and chex mix. I will do another pass over this and clean it up.

## Usage example
```
WORK_DIR=/tmp/kraft VOLUME_GROUP=F63F059E-BDAA-4C89-B78C-C32AC13B8D76 KDK_OUTPUT_PATH=/Users/ericrabil/Eric.kdk KERNEL_PATH=/Users/ericrabil/kernel.release.t8103.k ./make-kdk.sh all
```

- `WORK_DIR` (optional) where to work! defaults to /tmp/kdeek
- `VOLUME_GROUP` the preboot volume group to use
- `KDK_OUTPUT_PATH` where to generate the KDK
- `KERNEL_PATH` the path to the kernel to include in the custom kernelcache

## TODO
Rewrite in Go, probably